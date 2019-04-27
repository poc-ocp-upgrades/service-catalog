package server

import (
	"os"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	genericserveroptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/klog"
)

const (
	certDirectory		= "/var/run/kubernetes-service-catalog"
	storageTypeFlagName	= "storage-type"
)

type ServiceCatalogServerOptions struct {
	GenericServerRunOptions	*genericserveroptions.ServerRunOptions
	AdmissionOptions	*genericserveroptions.AdmissionOptions
	SecureServingOptions	*genericserveroptions.SecureServingOptionsWithLoopback
	AuthenticationOptions	*genericserveroptions.DelegatingAuthenticationOptions
	AuthorizationOptions	*genericserveroptions.DelegatingAuthorizationOptions
	AuditOptions		*genericserveroptions.AuditOptions
	EtcdOptions		*EtcdOptions
	DisableAuth		bool
	StandaloneMode		bool
	ServeOpenAPISpec	bool
	KubeconfigPath		string
}

func NewServiceCatalogServerOptions() *ServiceCatalogServerOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	opts := &ServiceCatalogServerOptions{GenericServerRunOptions: genericserveroptions.NewServerRunOptions(), AdmissionOptions: genericserveroptions.NewAdmissionOptions(), SecureServingOptions: genericserveroptions.NewSecureServingOptions().WithLoopback(), AuthenticationOptions: genericserveroptions.NewDelegatingAuthenticationOptions(), AuthorizationOptions: genericserveroptions.NewDelegatingAuthorizationOptions(), AuditOptions: genericserveroptions.NewAuditOptions(), EtcdOptions: NewEtcdOptions(), StandaloneMode: standaloneMode()}
	registerAllAdmissionPlugins(opts.AdmissionOptions.Plugins)
	opts.SecureServingOptions.ServerCert.CertDirectory = certDirectory
	return opts
}
func (s *ServiceCatalogServerOptions) AddFlags(flags *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_ = flags.String(storageTypeFlagName, "", "The type of backing storage this API server should use")
	flags.MarkDeprecated(storageTypeFlagName, "The value of this flag is now unused and will be removed in the near future")
	flags.Lookup(storageTypeFlagName).Hidden = false
	flags.BoolVar(&s.DisableAuth, "disable-auth", false, "Disable authentication and authorization for testing purposes")
	flags.BoolVar(&s.ServeOpenAPISpec, "serve-openapi-spec", false, "Whether this API server should serve the OpenAPI spec (problematic with older versions of kubectl)")
	flags.StringVar(&s.KubeconfigPath, "kubeconfig", "", "Path to kubeconfig to use over the in-cluster service account token")
	s.GenericServerRunOptions.AddUniversalFlags(flags)
	s.AdmissionOptions.AddFlags(flags)
	s.SecureServingOptions.AddFlags(flags)
	s.AuthenticationOptions.AddFlags(flags)
	s.AuthorizationOptions.AddFlags(flags)
	s.EtcdOptions.addFlags(flags)
	s.AuditOptions.AddFlags(flags)
}
func (s *ServiceCatalogServerOptions) Validate() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	errors := []error{}
	errors = append(errors, s.SecureServingOptions.Validate()...)
	errors = append(errors, s.AuthenticationOptions.Validate()...)
	errors = append(errors, s.AuthorizationOptions.Validate()...)
	etcdErrs := s.EtcdOptions.Validate()
	if len(etcdErrs) > 0 {
		klog.Errorln("Error validating etcd options, do you have `--etcd-servers localhost` set?")
	}
	errors = append(errors, etcdErrs...)
	return utilerrors.NewAggregate(errors)
}
func standaloneMode() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	val := os.Getenv("SERVICE_CATALOG_STANDALONE")
	return val == "true"
}
