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

type ServiceBindingsGetter interface {
	ServiceBindings(namespace string) ServiceBindingInterface
}
type ServiceBindingInterface interface {
	Create(*servicecatalog.ServiceBinding) (*servicecatalog.ServiceBinding, error)
	Update(*servicecatalog.ServiceBinding) (*servicecatalog.ServiceBinding, error)
	UpdateStatus(*servicecatalog.ServiceBinding) (*servicecatalog.ServiceBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*servicecatalog.ServiceBinding, error)
	List(opts v1.ListOptions) (*servicecatalog.ServiceBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceBinding, err error)
	ServiceBindingExpansion
}
type serviceBindings struct {
	client	rest.Interface
	ns	string
}

func newServiceBindings(c *ServicecatalogClient, namespace string) *serviceBindings {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBindings{client: c.RESTClient(), ns: namespace}
}
func (c *serviceBindings) Get(name string, options v1.GetOptions) (result *servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBinding{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebindings").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *serviceBindings) List(opts v1.ListOptions) (result *servicecatalog.ServiceBindingList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &servicecatalog.ServiceBindingList{}
	err = c.client.Get().Namespace(c.ns).Resource("servicebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *serviceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("servicebindings").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *serviceBindings) Create(serviceBinding *servicecatalog.ServiceBinding) (result *servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBinding{}
	err = c.client.Post().Namespace(c.ns).Resource("servicebindings").Body(serviceBinding).Do().Into(result)
	return
}
func (c *serviceBindings) Update(serviceBinding *servicecatalog.ServiceBinding) (result *servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBinding{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebindings").Name(serviceBinding.Name).Body(serviceBinding).Do().Into(result)
	return
}
func (c *serviceBindings) UpdateStatus(serviceBinding *servicecatalog.ServiceBinding) (result *servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBinding{}
	err = c.client.Put().Namespace(c.ns).Resource("servicebindings").Name(serviceBinding.Name).SubResource("status").Body(serviceBinding).Do().Into(result)
	return
}
func (c *serviceBindings) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("servicebindings").Name(name).Body(options).Do().Error()
}
func (c *serviceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("servicebindings").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *serviceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ServiceBinding{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("servicebindings").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
