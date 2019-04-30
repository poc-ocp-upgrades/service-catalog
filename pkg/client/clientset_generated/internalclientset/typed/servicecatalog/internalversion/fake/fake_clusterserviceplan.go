package fake

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakeClusterServicePlans struct{ Fake *FakeServicecatalog }

var clusterserviceplansResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "", Resource: "clusterserviceplans"}
var clusterserviceplansKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "", Kind: "ClusterServicePlan"}

func (c *FakeClusterServicePlans) Get(name string, options v1.GetOptions) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterserviceplansResource, name), &servicecatalog.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) List(opts v1.ListOptions) (result *servicecatalog.ClusterServicePlanList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterserviceplansResource, clusterserviceplansKind, opts), &servicecatalog.ClusterServicePlanList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.ClusterServicePlanList{ListMeta: obj.(*servicecatalog.ClusterServicePlanList).ListMeta}
	for _, item := range obj.(*servicecatalog.ClusterServicePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeClusterServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterserviceplansResource, opts))
}
func (c *FakeClusterServicePlans) Create(clusterServicePlan *servicecatalog.ClusterServicePlan) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterserviceplansResource, clusterServicePlan), &servicecatalog.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) Update(clusterServicePlan *servicecatalog.ClusterServicePlan) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterserviceplansResource, clusterServicePlan), &servicecatalog.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) UpdateStatus(clusterServicePlan *servicecatalog.ClusterServicePlan) (*servicecatalog.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(clusterserviceplansResource, "status", clusterServicePlan), &servicecatalog.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServicePlan), err
}
func (c *FakeClusterServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterserviceplansResource, name), &servicecatalog.ClusterServicePlan{})
	return err
}
func (c *FakeClusterServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(clusterserviceplansResource, listOptions)
	_, err := c.Fake.Invokes(action, &servicecatalog.ClusterServicePlanList{})
	return err
}
func (c *FakeClusterServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterserviceplansResource, name, pt, data, subresources...), &servicecatalog.ClusterServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServicePlan), err
}
