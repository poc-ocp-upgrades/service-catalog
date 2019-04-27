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

type ServiceInstancesGetter interface {
	ServiceInstances(namespace string) ServiceInstanceInterface
}
type ServiceInstanceInterface interface {
	Create(*servicecatalog.ServiceInstance) (*servicecatalog.ServiceInstance, error)
	Update(*servicecatalog.ServiceInstance) (*servicecatalog.ServiceInstance, error)
	UpdateStatus(*servicecatalog.ServiceInstance) (*servicecatalog.ServiceInstance, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*servicecatalog.ServiceInstance, error)
	List(opts v1.ListOptions) (*servicecatalog.ServiceInstanceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceInstance, err error)
	ServiceInstanceExpansion
}
type serviceInstances struct {
	client	rest.Interface
	ns	string
}

func newServiceInstances(c *ServicecatalogClient, namespace string) *serviceInstances {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceInstances{client: c.RESTClient(), ns: namespace}
}
func (c *serviceInstances) Get(name string, options v1.GetOptions) (result *servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceInstance{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceinstances").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceInstances) List(opts v1.ListOptions) (result *servicecatalog.ServiceInstanceList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &servicecatalog.ServiceInstanceList{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceinstances").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *serviceInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("serviceinstances").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceInstances) Create(serviceInstance *servicecatalog.ServiceInstance) (result *servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceInstance{}
	err = c.client.Post().Namespace(c.ns).Resource("serviceinstances").Body(serviceInstance).Do().Into(result)
	return
}
func (c *serviceInstances) Update(serviceInstance *servicecatalog.ServiceInstance) (result *servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceInstance{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceinstances").Name(serviceInstance.Name).Body(serviceInstance).Do().Into(result)
	return
}
func (c *serviceInstances) UpdateStatus(serviceInstance *servicecatalog.ServiceInstance) (result *servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceInstance{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceinstances").Name(serviceInstance.Name).SubResource("status").Body(serviceInstance).Do().Into(result)
	return
}
func (c *serviceInstances) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("serviceinstances").Name(name).Body(options).Do().Error()
}
func (c *serviceInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("serviceinstances").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceInstance{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("serviceinstances").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
