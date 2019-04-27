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

type ClusterServiceClassesGetter interface {
	ClusterServiceClasses() ClusterServiceClassInterface
}
type ClusterServiceClassInterface interface {
	Create(*servicecatalog.ClusterServiceClass) (*servicecatalog.ClusterServiceClass, error)
	Update(*servicecatalog.ClusterServiceClass) (*servicecatalog.ClusterServiceClass, error)
	UpdateStatus(*servicecatalog.ClusterServiceClass) (*servicecatalog.ClusterServiceClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*servicecatalog.ClusterServiceClass, error)
	List(opts v1.ListOptions) (*servicecatalog.ClusterServiceClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ClusterServiceClass, err error)
	ClusterServiceClassExpansion
}
type clusterServiceClasses struct{ client rest.Interface }

func newClusterServiceClasses(c *ServicecatalogClient) *clusterServiceClasses {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceClasses{client: c.RESTClient()}
}
func (c *clusterServiceClasses) Get(name string, options v1.GetOptions) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServiceClass{}
	err = c.client.Get().Resource("clusterserviceclasses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *clusterServiceClasses) List(opts v1.ListOptions) (result *servicecatalog.ClusterServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &servicecatalog.ClusterServiceClassList{}
	err = c.client.Get().Resource("clusterserviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *clusterServiceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("clusterserviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *clusterServiceClasses) Create(clusterServiceClass *servicecatalog.ClusterServiceClass) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServiceClass{}
	err = c.client.Post().Resource("clusterserviceclasses").Body(clusterServiceClass).Do().Into(result)
	return
}
func (c *clusterServiceClasses) Update(clusterServiceClass *servicecatalog.ClusterServiceClass) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServiceClass{}
	err = c.client.Put().Resource("clusterserviceclasses").Name(clusterServiceClass.Name).Body(clusterServiceClass).Do().Into(result)
	return
}
func (c *clusterServiceClasses) UpdateStatus(clusterServiceClass *servicecatalog.ClusterServiceClass) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServiceClass{}
	err = c.client.Put().Resource("clusterserviceclasses").Name(clusterServiceClass.Name).SubResource("status").Body(clusterServiceClass).Do().Into(result)
	return
}
func (c *clusterServiceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("clusterserviceclasses").Name(name).Body(options).Do().Error()
}
func (c *clusterServiceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("clusterserviceclasses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *clusterServiceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServiceClass{}
	err = c.client.Patch(pt).Resource("clusterserviceclasses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
