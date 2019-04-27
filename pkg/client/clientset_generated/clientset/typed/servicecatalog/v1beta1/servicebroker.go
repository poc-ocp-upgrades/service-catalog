package v1beta1

import (
	"time"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scheme "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ServiceBrokersGetter interface {
	ServiceBrokers(namespace string) ServiceBrokerInterface
}
type ServiceBrokerInterface interface {
	Create(*v1beta1.ServiceBroker) (*v1beta1.ServiceBroker, error)
	Update(*v1beta1.ServiceBroker) (*v1beta1.ServiceBroker, error)
	UpdateStatus(*v1beta1.ServiceBroker) (*v1beta1.ServiceBroker, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ServiceBroker, error)
	List(opts v1.ListOptions) (*v1beta1.ServiceBrokerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBroker, err error)
	ServiceBrokerExpansion
}
type serviceBrokers struct {
	client	rest.Interface
	ns	string
}

func newServiceBrokers(c *ServicecatalogV1beta1Client, namespace string) *serviceBrokers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBrokers{client: c.RESTClient(), ns: namespace}
}
func (c *serviceBrokers) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBroker{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebrokers").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceBrokers) List(opts v1.ListOptions) (result *v1beta1.ServiceBrokerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ServiceBrokerList{}
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
func (c *serviceBrokers) Create(serviceBroker *v1beta1.ServiceBroker) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBroker{}
	err = c.client.Post().Namespace(c.ns).Resource("servicebrokers").Body(serviceBroker).Do().Into(result)
	return
}
func (c *serviceBrokers) Update(serviceBroker *v1beta1.ServiceBroker) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBroker{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebrokers").Name(serviceBroker.Name).Body(serviceBroker).Do().Into(result)
	return
}
func (c *serviceBrokers) UpdateStatus(serviceBroker *v1beta1.ServiceBroker) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBroker{}
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
func (c *serviceBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBroker{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("servicebrokers").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
