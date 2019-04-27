package internalversion

import (
	"time"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scheme "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ServiceBrokersGetter interface {
	ServiceBrokers(namespace string) ServiceBrokerInterface
}
type ServiceBrokerInterface interface {
	Create(*servicecatalog.ServiceBroker) (*servicecatalog.ServiceBroker, error)
	Update(*servicecatalog.ServiceBroker) (*servicecatalog.ServiceBroker, error)
	UpdateStatus(*servicecatalog.ServiceBroker) (*servicecatalog.ServiceBroker, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*servicecatalog.ServiceBroker, error)
	List(opts v1.ListOptions) (*servicecatalog.ServiceBrokerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceBroker, err error)
	ServiceBrokerExpansion
}
type serviceBrokers struct {
	client	rest.Interface
	ns	string
}

func newServiceBrokers(c *ServicecatalogClient, namespace string) *serviceBrokers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBrokers{client: c.RESTClient(), ns: namespace}
}
func (c *serviceBrokers) Get(name string, options v1.GetOptions) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBroker{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebrokers").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceBrokers) List(opts v1.ListOptions) (result *servicecatalog.ServiceBrokerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &servicecatalog.ServiceBrokerList{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebrokers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *serviceBrokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("servicebrokers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceBrokers) Create(serviceBroker *servicecatalog.ServiceBroker) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBroker{}
	err = c.client.Post().Namespace(c.ns).Resource("servicebrokers").Body(serviceBroker).Do().Into(result)
	return
}
func (c *serviceBrokers) Update(serviceBroker *servicecatalog.ServiceBroker) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBroker{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebrokers").Name(serviceBroker.Name).Body(serviceBroker).Do().Into(result)
	return
}
func (c *serviceBrokers) UpdateStatus(serviceBroker *servicecatalog.ServiceBroker) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBroker{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebrokers").Name(serviceBroker.Name).SubResource("status").Body(serviceBroker).Do().Into(result)
	return
}
func (c *serviceBrokers) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("servicebrokers").Name(name).Body(options).Do().Error()
}
func (c *serviceBrokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("servicebrokers").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBroker{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("servicebrokers").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
