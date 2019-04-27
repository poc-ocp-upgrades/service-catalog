package fake

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeClusterServicePlans struct{ Fake *FakeServicecatalogV1beta1 }

var clusterserviceplansResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "clusterserviceplans"}
var clusterserviceplansKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ClusterServicePlan"}

func (c *FakeClusterServicePlans) Get(name string, options v1.GetOptions) (result *v1beta1.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterserviceplansResource, name), &v1beta1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) List(opts v1.ListOptions) (result *v1beta1.ClusterServicePlanList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterserviceplansResource, clusterserviceplansKind, opts), &v1beta1.ClusterServicePlanList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterServicePlanList{ListMeta: obj.(*v1beta1.ClusterServicePlanList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterServicePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeClusterServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterserviceplansResource, opts))
}
func (c *FakeClusterServicePlans) Create(clusterServicePlan *v1beta1.ClusterServicePlan) (result *v1beta1.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterserviceplansResource, clusterServicePlan), &v1beta1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) Update(clusterServicePlan *v1beta1.ClusterServicePlan) (result *v1beta1.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterserviceplansResource, clusterServicePlan), &v1beta1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) UpdateStatus(clusterServicePlan *v1beta1.ClusterServicePlan) (*v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(clusterserviceplansResource, "status", clusterServicePlan), &v1beta1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterserviceplansResource, name), &v1beta1.ClusterServicePlan{})
	return err
}
func (c *FakeClusterServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(clusterserviceplansResource, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ClusterServicePlanList{})
	return err
}
func (c *FakeClusterServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterserviceplansResource, name, pt, data, subresources...), &v1beta1.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServicePlan), err
}
