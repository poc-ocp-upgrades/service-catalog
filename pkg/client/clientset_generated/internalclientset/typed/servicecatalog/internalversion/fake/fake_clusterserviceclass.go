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

type FakeClusterServiceClasses struct{ Fake *FakeServicecatalog }

var clusterserviceclassesResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "", Resource: "clusterserviceclasses"}
var clusterserviceclassesKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "", Kind: "ClusterServiceClass"}

func (c *FakeClusterServiceClasses) Get(name string, options v1.GetOptions) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterserviceclassesResource, name), &servicecatalog.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) List(opts v1.ListOptions) (result *servicecatalog.ClusterServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterserviceclassesResource, clusterserviceclassesKind, opts), &servicecatalog.ClusterServiceClassList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.ClusterServiceClassList{ListMeta: obj.(*servicecatalog.ClusterServiceClassList).ListMeta}
	for _, item := range obj.(*servicecatalog.ClusterServiceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeClusterServiceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterserviceclassesResource, opts))
}
func (c *FakeClusterServiceClasses) Create(clusterServiceClass *servicecatalog.ClusterServiceClass) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterserviceclassesResource, clusterServiceClass), &servicecatalog.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) Update(clusterServiceClass *servicecatalog.ClusterServiceClass) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterserviceclassesResource, clusterServiceClass), &servicecatalog.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) UpdateStatus(clusterServiceClass *servicecatalog.ClusterServiceClass) (*servicecatalog.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(clusterserviceclassesResource, "status", clusterServiceClass), &servicecatalog.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterserviceclassesResource, name), &servicecatalog.ClusterServiceClass{})
	return err
}
func (c *FakeClusterServiceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(clusterserviceclassesResource, listOptions)
	_, err := c.Fake.Invokes(action, &servicecatalog.ClusterServiceClassList{})
	return err
}
func (c *FakeClusterServiceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterserviceclassesResource, name, pt, data, subresources...), &servicecatalog.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ClusterServiceClass), err
}
