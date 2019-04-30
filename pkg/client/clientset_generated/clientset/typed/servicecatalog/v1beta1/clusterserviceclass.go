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

type ClusterServiceClassesGetter interface {
	ClusterServiceClasses() ClusterServiceClassInterface
}
type ClusterServiceClassInterface interface {
	Create(*v1beta1.ClusterServiceClass) (*v1beta1.ClusterServiceClass, error)
	Update(*v1beta1.ClusterServiceClass) (*v1beta1.ClusterServiceClass, error)
	UpdateStatus(*v1beta1.ClusterServiceClass) (*v1beta1.ClusterServiceClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ClusterServiceClass, error)
	List(opts v1.ListOptions) (*v1beta1.ClusterServiceClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServiceClass, err error)
	ClusterServiceClassExpansion
}
type clusterServiceClasses struct{ client rest.Interface }

func newClusterServiceClasses(c *ServicecatalogV1beta1Client) *clusterServiceClasses {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceClasses{client: c.RESTClient()}
}
func (c *clusterServiceClasses) Get(name string, options v1.GetOptions) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceClass{}
	err = c.client.Get().Resource("clusterserviceclasses").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *clusterServiceClasses) List(opts v1.ListOptions) (result *v1beta1.ClusterServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ClusterServiceClassList{}
	err = c.client.Get().Resource("clusterserviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *clusterServiceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("clusterserviceclasses").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *clusterServiceClasses) Create(clusterServiceClass *v1beta1.ClusterServiceClass) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceClass{}
	err = c.client.Post().Resource("clusterserviceclasses").Body(clusterServiceClass).Do().Into(result)
	return
}
func (c *clusterServiceClasses) Update(clusterServiceClass *v1beta1.ClusterServiceClass) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceClass{}
	err = c.client.Put().Resource("clusterserviceclasses").Name(clusterServiceClass.Name).Body(clusterServiceClass).Do().Into(result)
	return
}
func (c *clusterServiceClasses) UpdateStatus(clusterServiceClass *v1beta1.ClusterServiceClass) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceClass{}
	err = c.client.Put().Resource("clusterserviceclasses").Name(clusterServiceClass.Name).SubResource("status").Body(clusterServiceClass).Do().Into(result)
	return
}
func (c *clusterServiceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("clusterserviceclasses").Name(name).Body(options).Do().Error()
}
func (c *clusterServiceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("clusterserviceclasses").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *clusterServiceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceClass{}
	err = c.client.Patch(pt).Resource("clusterserviceclasses").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
