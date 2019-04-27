package internalversion

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ServicecatalogInterface interface {
	RESTClient() rest.Interface
	ClusterServiceBrokersGetter
	ClusterServiceClassesGetter
	ClusterServicePlansGetter
	ServiceBindingsGetter
	ServiceBrokersGetter
	ServiceClassesGetter
	ServiceInstancesGetter
	ServicePlansGetter
}
type ServicecatalogClient struct{ restClient rest.Interface }

func (c *ServicecatalogClient) ClusterServiceBrokers() ClusterServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newClusterServiceBrokers(c)
}
func (c *ServicecatalogClient) ClusterServiceClasses() ClusterServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newClusterServiceClasses(c)
}
func (c *ServicecatalogClient) ClusterServicePlans() ClusterServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newClusterServicePlans(c)
}
func (c *ServicecatalogClient) ServiceBindings(namespace string) ServiceBindingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceBindings(c, namespace)
}
func (c *ServicecatalogClient) ServiceBrokers(namespace string) ServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceBrokers(c, namespace)
}
func (c *ServicecatalogClient) ServiceClasses(namespace string) ServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceClasses(c, namespace)
}
func (c *ServicecatalogClient) ServiceInstances(namespace string) ServiceInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceInstances(c, namespace)
}
func (c *ServicecatalogClient) ServicePlans(namespace string) ServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServicePlans(c, namespace)
}
func NewForConfig(c *rest.Config) (*ServicecatalogClient, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ServicecatalogClient{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *ServicecatalogClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
func New(c rest.Interface) *ServicecatalogClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ServicecatalogClient{c}
}
func setConfigDefaults(config *rest.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != scheme.Scheme.PrioritizedVersionsForGroup("servicecatalog.k8s.io")[0].Group {
		gv := scheme.Scheme.PrioritizedVersionsForGroup("servicecatalog.k8s.io")[0]
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs
	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}
	return nil
}
func (c *ServicecatalogClient) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.restClient
}
