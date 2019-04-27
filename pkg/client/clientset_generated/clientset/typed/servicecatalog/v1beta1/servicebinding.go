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

type ServiceBindingsGetter interface {
	ServiceBindings(namespace string) ServiceBindingInterface
}
type ServiceBindingInterface interface {
	Create(*v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
	Update(*v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
	UpdateStatus(*v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ServiceBinding, error)
	List(opts v1.ListOptions) (*v1beta1.ServiceBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBinding, err error)
	ServiceBindingExpansion
}
type serviceBindings struct {
	client	rest.Interface
	ns	string
}

func newServiceBindings(c *ServicecatalogV1beta1Client, namespace string) *serviceBindings {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBindings{client: c.RESTClient(), ns: namespace}
}
func (c *serviceBindings) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBinding{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebindings").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceBindings) List(opts v1.ListOptions) (result *v1beta1.ServiceBindingList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ServiceBindingList{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *serviceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("servicebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceBindings) Create(serviceBinding *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBinding{}
	err = c.client.Post().Namespace(c.ns).Resource("servicebindings").Body(serviceBinding).Do().Into(result)
	return
}
func (c *serviceBindings) Update(serviceBinding *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBinding{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebindings").Name(serviceBinding.Name).Body(serviceBinding).Do().Into(result)
	return
}
func (c *serviceBindings) UpdateStatus(serviceBinding *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBinding{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebindings").Name(serviceBinding.Name).SubResource("status").Body(serviceBinding).Do().Into(result)
	return
}
func (c *serviceBindings) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("servicebindings").Name(name).Body(options).Do().Error()
}
func (c *serviceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("servicebindings").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceBinding{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("servicebindings").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
