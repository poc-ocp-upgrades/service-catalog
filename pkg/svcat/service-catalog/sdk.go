package servicecatalog

import (
	"time"
	apiv1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
	apicorev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type SvcatClient interface {
	Bind(string, string, string, string, string, interface{}, map[string]string) (*apiv1beta1.ServiceBinding, error)
	BindingParentHierarchy(*apiv1beta1.ServiceBinding) (*apiv1beta1.ServiceInstance, *apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, *apiv1beta1.ClusterServiceBroker, error)
	DeleteBinding(string, string) error
	DeleteBindings([]types.NamespacedName) ([]types.NamespacedName, error)
	IsBindingFailed(*apiv1beta1.ServiceBinding) bool
	IsBindingReady(*apiv1beta1.ServiceBinding) bool
	RetrieveBinding(string, string) (*apiv1beta1.ServiceBinding, error)
	RetrieveBindings(string) (*apiv1beta1.ServiceBindingList, error)
	RetrieveBindingsByInstance(*apiv1beta1.ServiceInstance) ([]apiv1beta1.ServiceBinding, error)
	Unbind(string, string) ([]types.NamespacedName, error)
	WaitForBinding(string, string, time.Duration, *time.Duration) (*apiv1beta1.ServiceBinding, error)
	Deregister(string, *ScopeOptions) error
	RetrieveBrokers(opts ScopeOptions) ([]Broker, error)
	RetrieveBroker(string) (*apiv1beta1.ClusterServiceBroker, error)
	RetrieveBrokerByClass(*apiv1beta1.ClusterServiceClass) (*apiv1beta1.ClusterServiceBroker, error)
	Register(string, string, *RegisterOptions, *ScopeOptions) (Broker, error)
	Sync(string, ScopeOptions, int) error
	WaitForBroker(string, time.Duration, *time.Duration) (Broker, error)
	RetrieveClasses(ScopeOptions) ([]Class, error)
	RetrieveClassByName(string, ScopeOptions) (Class, error)
	RetrieveClassByID(string) (*apiv1beta1.ClusterServiceClass, error)
	RetrieveClassByPlan(Plan) (*apiv1beta1.ClusterServiceClass, error)
	CreateClassFrom(CreateClassFromOptions) (Class, error)
	Deprovision(string, string) error
	InstanceParentHierarchy(*apiv1beta1.ServiceInstance) (*apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, *apiv1beta1.ClusterServiceBroker, error)
	InstanceToServiceClassAndPlan(*apiv1beta1.ServiceInstance) (*apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, error)
	IsInstanceFailed(*apiv1beta1.ServiceInstance) bool
	IsInstanceReady(*apiv1beta1.ServiceInstance) bool
	Provision(string, string, string, *ProvisionOptions) (*apiv1beta1.ServiceInstance, error)
	RetrieveInstance(string, string) (*apiv1beta1.ServiceInstance, error)
	RetrieveInstanceByBinding(*apiv1beta1.ServiceBinding) (*apiv1beta1.ServiceInstance, error)
	RetrieveInstances(string, string, string) (*apiv1beta1.ServiceInstanceList, error)
	RetrieveInstancesByPlan(Plan) ([]apiv1beta1.ServiceInstance, error)
	TouchInstance(string, string, int) error
	WaitForInstance(string, string, time.Duration, *time.Duration) (*apiv1beta1.ServiceInstance, error)
	WaitForInstanceToNotExist(string, string, time.Duration, *time.Duration) (*apiv1beta1.ServiceInstance, error)
	RetrievePlans(string, ScopeOptions) ([]Plan, error)
	RetrievePlanByName(string, ScopeOptions) (Plan, error)
	RetrievePlanByClassAndName(string, string, ScopeOptions) (Plan, error)
	RetrievePlanByClassIDAndName(string, string, ScopeOptions) (Plan, error)
	RetrievePlanByID(string, ScopeOptions) (Plan, error)
	RetrieveSecretByBinding(*apiv1beta1.ServiceBinding) (*apicorev1.Secret, error)
	ServerVersion() (*version.Info, error)
}
type SDK struct {
	K8sClient		kubernetes.Interface
	ServiceCatalogClient	clientset.Interface
}

func (sdk *SDK) ServiceCatalog() v1beta1.ServicecatalogV1beta1Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdk.ServiceCatalogClient.ServicecatalogV1beta1()
}
func (sdk *SDK) Core() corev1.CoreV1Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdk.K8sClient.CoreV1()
}
