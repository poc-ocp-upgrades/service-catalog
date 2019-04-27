package v1beta1

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ServicecatalogV1beta1Interface interface {
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
type ServicecatalogV1beta1Client struct{ restClient rest.Interface }

func (c *ServicecatalogV1beta1Client) ClusterServiceBrokers() ClusterServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newClusterServiceBrokers(c)
}
func (c *ServicecatalogV1beta1Client) ClusterServiceClasses() ClusterServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newClusterServiceClasses(c)
}
func (c *ServicecatalogV1beta1Client) ClusterServicePlans() ClusterServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newClusterServicePlans(c)
}
func (c *ServicecatalogV1beta1Client) ServiceBindings(namespace string) ServiceBindingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceBindings(c, namespace)
}
func (c *ServicecatalogV1beta1Client) ServiceBrokers(namespace string) ServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceBrokers(c, namespace)
}
func (c *ServicecatalogV1beta1Client) ServiceClasses(namespace string) ServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceClasses(c, namespace)
}
func (c *ServicecatalogV1beta1Client) ServiceInstances(namespace string) ServiceInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceInstances(c, namespace)
}
func (c *ServicecatalogV1beta1Client) ServicePlans(namespace string) ServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServicePlans(c, namespace)
}
func NewForConfig(c *rest.Config) (*ServicecatalogV1beta1Client, error) {
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
	return &ServicecatalogV1beta1Client{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *ServicecatalogV1beta1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
func New(c rest.Interface) *ServicecatalogV1beta1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &ServicecatalogV1beta1Client{c}
}
func setConfigDefaults(config *rest.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	return nil
}
func (c *ServicecatalogV1beta1Client) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.restClient
}
