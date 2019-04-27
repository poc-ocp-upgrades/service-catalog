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

type ServiceClassesGetter interface {
	ServiceClasses(namespace string) ServiceClassInterface
}
type ServiceClassInterface interface {
	Create(*v1beta1.ServiceClass) (*v1beta1.ServiceClass, error)
	Update(*v1beta1.ServiceClass) (*v1beta1.ServiceClass, error)
	UpdateStatus(*v1beta1.ServiceClass) (*v1beta1.ServiceClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ServiceClass, error)
	List(opts v1.ListOptions) (*v1beta1.ServiceClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceClass, err error)
	ServiceClassExpansion
}
type serviceClasses struct {
	client	rest.Interface
	ns	string
}

func newServiceClasses(c *ServicecatalogV1beta1Client, namespace string) *serviceClasses {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceClasses{client: c.RESTClient(), ns: namespace}
}
func (c *serviceClasses) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceClass{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceclasses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceClasses) List(opts v1.ListOptions) (result *v1beta1.ServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ServiceClassList{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *serviceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("serviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceClasses) Create(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceClass{}
	err = c.client.Post().Namespace(c.ns).Resource("serviceclasses").Body(serviceClass).Do().Into(result)
	return
}
func (c *serviceClasses) Update(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceClass{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceclasses").Name(serviceClass.Name).Body(serviceClass).Do().Into(result)
	return
}
func (c *serviceClasses) UpdateStatus(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceClass{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceclasses").Name(serviceClass.Name).SubResource("status").Body(serviceClass).Do().Into(result)
	return
}
func (c *serviceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("serviceclasses").Name(name).Body(options).Do().Error()
}
func (c *serviceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("serviceclasses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceClass{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("serviceclasses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
