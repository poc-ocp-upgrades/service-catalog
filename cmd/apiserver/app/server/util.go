package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
	"github.com/go-openapi/spec"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/apiserver/pkg/admission/initializer"
	admissionmetrics "k8s.io/apiserver/pkg/admission/metrics"
	"k8s.io/apiserver/pkg/authorization/authorizerfactory"
	apiopenapi "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericserveroptions "k8s.io/apiserver/pkg/server/options"
	kubeinformers "k8s.io/client-go/informers"
	kubeclientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	scadmission "github.com/kubernetes-incubator/service-catalog/pkg/apiserver/admission"
	"github.com/kubernetes-incubator/service-catalog/pkg/apiserver/authenticator"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset"
	informers "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion"
	"github.com/kubernetes-incubator/service-catalog/pkg/openapi"
	"github.com/kubernetes-incubator/service-catalog/pkg/util/kube"
	"github.com/kubernetes-incubator/service-catalog/pkg/version"
)

const (
	inClusterNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
)

type serviceCatalogConfig struct {
	sharedInformers		informers.SharedInformerFactory
	kubeSharedInformers	kubeinformers.SharedInformerFactory
	client				internalclientset.Interface
	kubeClient			kubeclientset.Interface
}

func buildGenericConfig(s *ServiceCatalogServerOptions) (*genericapiserver.RecommendedConfig, *serviceCatalogConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.StandaloneMode {
		klog.Infof("service catalog is in standalone mode")
	}
	if err := s.SecureServingOptions.MaybeDefaultWithSelfSignedCerts(s.GenericServerRunOptions.AdvertiseAddress.String(), nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, nil, err
	}
	genericConfig := genericapiserver.NewRecommendedConfig(api.Codecs)
	if err := s.GenericServerRunOptions.ApplyTo(&genericConfig.Config); err != nil {
		return nil, nil, err
	}
	if err := s.SecureServingOptions.ApplyTo(&genericConfig.Config.SecureServing, &genericConfig.Config.LoopbackClientConfig); err != nil {
		return nil, nil, err
	}
	if !s.DisableAuth && !s.StandaloneMode {
		if err := s.AuthenticationOptions.ApplyTo(&genericConfig.Config.Authentication, genericConfig.Config.SecureServing, genericConfig.Config.OpenAPIConfig); err != nil {
			return nil, nil, err
		}
		if err := s.AuthorizationOptions.ApplyTo(&genericConfig.Config.Authorization); err != nil {
			return nil, nil, err
		}
	} else {
		klog.Warning("Authentication and authorization disabled for testing purposes")
		genericConfig.Authentication.Authenticator = &authenticator.AnyUserAuthenticator{}
		genericConfig.Authorization.Authorizer = authorizerfactory.NewAlwaysAllowAuthorizer()
	}
	namespace, err := getInClusterNamespace("service-catalog")
	if err != nil {
		return nil, nil, err
	}
	if err := s.AuditOptions.ApplyTo(&genericConfig.Config, genericConfig.ClientConfig, genericConfig.SharedInformerFactory, genericserveroptions.NewProcessInfo("service-catalog-apiserver", namespace), nil); err != nil {
		return nil, nil, err
	}
	if s.ServeOpenAPISpec {
		genericConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(openapi.GetOpenAPIDefinitions, apiopenapi.NewDefinitionNamer(api.Scheme))
		if genericConfig.OpenAPIConfig.Info == nil {
			genericConfig.OpenAPIConfig.Info = &spec.Info{}
		}
		if genericConfig.OpenAPIConfig.Info.Version == "" {
			if genericConfig.Version != nil {
				genericConfig.OpenAPIConfig.Info.Version = strings.Split(genericConfig.Version.String(), "-")[0]
			} else {
				genericConfig.OpenAPIConfig.Info.Version = "unversioned"
			}
		}
	} else {
		klog.Warning("OpenAPI spec will not be served")
	}
	genericConfig.SwaggerConfig = genericapiserver.DefaultSwaggerConfig()
	genericConfig.EnableMetrics = true
	serviceCatalogVersion := version.Get()
	genericConfig.Version = &serviceCatalogVersion
	client, err := internalclientset.NewForConfig(genericConfig.LoopbackClientConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create clientset for service catalog self-communication: %v", err)
	}
	sharedInformers := informers.NewSharedInformerFactory(client, 10*time.Minute)
	scConfig := &serviceCatalogConfig{client: client, sharedInformers: sharedInformers}
	if !s.StandaloneMode {
		clusterConfig, err := kube.LoadConfig(s.KubeconfigPath, "")
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse kube client config: %v", err)
		}
		if clusterConfig == nil {
			clusterConfig, err = restclient.InClusterConfig()
			if err != nil {
				return nil, nil, fmt.Errorf("failed to get kube client config: %v", err)
			}
		}
		clusterConfig.GroupVersion = &schema.GroupVersion{}
		kubeClient, err := kubeclientset.NewForConfig(clusterConfig)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create clientset interface: %v", err)
		}
		kubeSharedInformers := kubeinformers.NewSharedInformerFactory(kubeClient, 10*time.Minute)
		genericConfig.SharedInformerFactory = kubeSharedInformers
		genericConfig.AdmissionControl, err = buildAdmission(genericConfig, s, client, sharedInformers, kubeClient, kubeSharedInformers)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to initialize admission: %v", err)
		}
		scConfig.kubeClient = kubeClient
		scConfig.kubeSharedInformers = kubeSharedInformers
	}
	return genericConfig, scConfig, nil
}
func buildAdmission(c *genericapiserver.RecommendedConfig, s *ServiceCatalogServerOptions, client internalclientset.Interface, sharedInformers informers.SharedInformerFactory, kubeClient kubeclientset.Interface, kubeSharedInformers kubeinformers.SharedInformerFactory) (admission.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pluginNames := enabledPluginNames(s.AdmissionOptions)
	klog.Infof("Admission control plugin names: %v", pluginNames)
	genericInitializer := initializer.New(kubeClient, kubeSharedInformers, c.Authorization.Authorizer, api.Scheme)
	scPluginInitializer := scadmission.NewPluginInitializer(client, sharedInformers, kubeClient, kubeSharedInformers)
	initializersChain := admission.PluginInitializers{scPluginInitializer, genericInitializer}
	pluginsConfigProvider, err := admission.ReadAdmissionConfiguration(pluginNames, s.AdmissionOptions.ConfigFile, api.Scheme)
	if err != nil {
		return nil, fmt.Errorf("failed to read plugin config: %v", err)
	}
	return s.AdmissionOptions.Plugins.NewFromPlugins(pluginNames, pluginsConfigProvider, initializersChain, admission.DecoratorFunc(admissionmetrics.WithControllerMetrics))
}
func enabledPluginNames(a *genericserveroptions.AdmissionOptions) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	allOffPlugins := append(a.DefaultOffPlugins.List(), a.DisablePlugins...)
	disabledPlugins := sets.NewString(allOffPlugins...)
	enabledPlugins := sets.NewString(a.EnablePlugins...)
	disabledPlugins = disabledPlugins.Difference(enabledPlugins)
	resultPlugins := sets.NewString()
	orderedPlugins := []string{}
	for _, plugin := range a.RecommendedPluginOrder {
		if !disabledPlugins.Has(plugin) {
			orderedPlugins = append(orderedPlugins, plugin)
			resultPlugins.Insert(plugin)
		}
	}
	for plugin := range enabledPlugins {
		if !resultPlugins.Has(plugin) {
			orderedPlugins = append(orderedPlugins, plugin)
			resultPlugins.Insert(plugin)
		}
	}
	return orderedPlugins
}
func addPostStartHooks(server *genericapiserver.GenericAPIServer, scConfig *serviceCatalogConfig, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	server.AddPostStartHook("start-service-catalog-apiserver-informers", func(context genericapiserver.PostStartHookContext) error {
		klog.Infof("Starting shared informers")
		scConfig.sharedInformers.Start(stopCh)
		if scConfig.kubeSharedInformers != nil {
			scConfig.kubeSharedInformers.Start(stopCh)
		}
		klog.Infof("Started shared informers")
		return nil
	})
}
func getInClusterNamespace(defaultNamespace string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := os.Stat(inClusterNamespacePath)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultNamespace, nil
		}
		return "", fmt.Errorf("error checking namespace file: %v", err)
	}
	namespace, err := ioutil.ReadFile(inClusterNamespacePath)
	if err != nil {
		return "", fmt.Errorf("error reading namespace file: %v", err)
	}
	return string(namespace), nil
}
