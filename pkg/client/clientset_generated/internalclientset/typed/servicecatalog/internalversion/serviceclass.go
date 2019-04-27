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

type ServiceClassesGetter interface {
	ServiceClasses(namespace string) ServiceClassInterface
}
type ServiceClassInterface interface {
	Create(*servicecatalog.ServiceClass) (*servicecatalog.ServiceClass, error)
	Update(*servicecatalog.ServiceClass) (*servicecatalog.ServiceClass, error)
	UpdateStatus(*servicecatalog.ServiceClass) (*servicecatalog.ServiceClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*servicecatalog.ServiceClass, error)
	List(opts v1.ListOptions) (*servicecatalog.ServiceClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceClass, err error)
	ServiceClassExpansion
}
type serviceClasses struct {
	client	rest.Interface
	ns	string
}

func newServiceClasses(c *ServicecatalogClient, namespace string) *serviceClasses {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceClasses{client: c.RESTClient(), ns: namespace}
}
func (c *serviceClasses) Get(name string, options v1.GetOptions) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceClass{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceclasses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceClasses) List(opts v1.ListOptions) (result *servicecatalog.ServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &servicecatalog.ServiceClassList{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *serviceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("serviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceClasses) Create(serviceClass *servicecatalog.ServiceClass) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceClass{}
	err = c.client.Post().Namespace(c.ns).Resource("serviceclasses").Body(serviceClass).Do().Into(result)
	return
}
func (c *serviceClasses) Update(serviceClass *servicecatalog.ServiceClass) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceClass{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceclasses").Name(serviceClass.Name).Body(serviceClass).Do().Into(result)
	return
}
func (c *serviceClasses) UpdateStatus(serviceClass *servicecatalog.ServiceClass) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceClass{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceclasses").Name(serviceClass.Name).SubResource("status").Body(serviceClass).Do().Into(result)
	return
}
func (c *serviceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("serviceclasses").Name(name).Body(options).Do().Error()
}
func (c *serviceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("serviceclasses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceClass{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("serviceclasses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
