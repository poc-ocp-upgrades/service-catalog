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

type ClusterServicePlansGetter interface {
	ClusterServicePlans() ClusterServicePlanInterface
}
type ClusterServicePlanInterface interface {
	Create(*servicecatalog.ClusterServicePlan) (*servicecatalog.ClusterServicePlan, error)
	Update(*servicecatalog.ClusterServicePlan) (*servicecatalog.ClusterServicePlan, error)
	UpdateStatus(*servicecatalog.ClusterServicePlan) (*servicecatalog.ClusterServicePlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*servicecatalog.ClusterServicePlan, error)
	List(opts v1.ListOptions) (*servicecatalog.ClusterServicePlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ClusterServicePlan, err error)
	ClusterServicePlanExpansion
}
type clusterServicePlans struct{ client rest.Interface }

func newClusterServicePlans(c *ServicecatalogClient) *clusterServicePlans {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServicePlans{client: c.RESTClient()}
}
func (c *clusterServicePlans) Get(name string, options v1.GetOptions) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServicePlan{}
	err = c.client.Get().Resource("clusterserviceplans").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *clusterServicePlans) List(opts v1.ListOptions) (result *servicecatalog.ClusterServicePlanList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &servicecatalog.ClusterServicePlanList{}
	err = c.client.Get().Resource("clusterserviceplans").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *clusterServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Resource("clusterserviceplans").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *clusterServicePlans) Create(clusterServicePlan *servicecatalog.ClusterServicePlan) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServicePlan{}
	err = c.client.Post().Resource("clusterserviceplans").Body(clusterServicePlan).Do().Into(result)
	return
}
func (c *clusterServicePlans) Update(clusterServicePlan *servicecatalog.ClusterServicePlan) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServicePlan{}
	err = c.client.Put().Resource("clusterserviceplans").Name(clusterServicePlan.Name).Body(clusterServicePlan).Do().Into(result)
	return
}
func (c *clusterServicePlans) UpdateStatus(clusterServicePlan *servicecatalog.ClusterServicePlan) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServicePlan{}
	err = c.client.Put().Resource("clusterserviceplans").Name(clusterServicePlan.Name).SubResource("status").Body(clusterServicePlan).Do().Into(result)
	return
}
func (c *clusterServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Resource("clusterserviceplans").Name(name).Body(options).Do().Error()
}
func (c *clusterServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Resource("clusterserviceplans").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *clusterServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &servicecatalog.ClusterServicePlan{}
	err = c.client.Patch(pt).Resource("clusterserviceplans").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
