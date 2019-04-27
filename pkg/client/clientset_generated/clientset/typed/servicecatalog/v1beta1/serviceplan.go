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

type ServicePlansGetter interface {
	ServicePlans(namespace string) ServicePlanInterface
}
type ServicePlanInterface interface {
	Create(*v1beta1.ServicePlan) (*v1beta1.ServicePlan, error)
	Update(*v1beta1.ServicePlan) (*v1beta1.ServicePlan, error)
	UpdateStatus(*v1beta1.ServicePlan) (*v1beta1.ServicePlan, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ServicePlan, error)
	List(opts v1.ListOptions) (*v1beta1.ServicePlanList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServicePlan, err error)
	ServicePlanExpansion
}
type servicePlans struct {
	client	rest.Interface
	ns	string
}

func newServicePlans(c *ServicecatalogV1beta1Client, namespace string) *servicePlans {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicePlans{client: c.RESTClient(), ns: namespace}
}
func (c *servicePlans) Get(name string, options v1.GetOptions) (result *v1beta1.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServicePlan{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceplans").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *servicePlans) List(opts v1.ListOptions) (result *v1beta1.ServicePlanList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ServicePlanList{}
	err = c.client.Get().Namespace(c.ns).Resource("serviceplans").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *servicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("serviceplans").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *servicePlans) Create(servicePlan *v1beta1.ServicePlan) (result *v1beta1.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServicePlan{}
	err = c.client.Post().Namespace(c.ns).Resource("serviceplans").Body(servicePlan).Do().Into(result)
	return
}
func (c *servicePlans) Update(servicePlan *v1beta1.ServicePlan) (result *v1beta1.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServicePlan{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceplans").Name(servicePlan.Name).Body(servicePlan).Do().Into(result)
	return
}
func (c *servicePlans) UpdateStatus(servicePlan *v1beta1.ServicePlan) (result *v1beta1.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServicePlan{}
	err = c.client.Put().Namespace(c.ns).Resource("serviceplans").Name(servicePlan.Name).SubResource("status").Body(servicePlan).Do().Into(result)
	return
}
func (c *servicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("serviceplans").Name(name).Body(options).Do().Error()
}
func (c *servicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("serviceplans").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *servicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServicePlan{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("serviceplans").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
