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

type FakeClusterServiceClasses struct{ Fake *FakeServicecatalogV1beta1 }

var clusterserviceclassesResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "clusterserviceclasses"}
var clusterserviceclassesKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ClusterServiceClass"}

func (c *FakeClusterServiceClasses) Get(name string, options v1.GetOptions) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterserviceclassesResource, name), &v1beta1.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) List(opts v1.ListOptions) (result *v1beta1.ClusterServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterserviceclassesResource, clusterserviceclassesKind, opts), &v1beta1.ClusterServiceClassList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterServiceClassList{ListMeta: obj.(*v1beta1.ClusterServiceClassList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterServiceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeClusterServiceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterserviceclassesResource, opts))
}
func (c *FakeClusterServiceClasses) Create(clusterServiceClass *v1beta1.ClusterServiceClass) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterserviceclassesResource, clusterServiceClass), &v1beta1.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) Update(clusterServiceClass *v1beta1.ClusterServiceClass) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterserviceclassesResource, clusterServiceClass), &v1beta1.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) UpdateStatus(clusterServiceClass *v1beta1.ClusterServiceClass) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(clusterserviceclassesResource, "status", clusterServiceClass), &v1beta1.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceClass), err
}
func (c *FakeClusterServiceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterserviceclassesResource, name), &v1beta1.ClusterServiceClass{})
	return err
}
func (c *FakeClusterServiceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(clusterserviceclassesResource, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ClusterServiceClassList{})
	return err
}
func (c *FakeClusterServiceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterserviceclassesResource, name, pt, data, subresources...), &v1beta1.ClusterServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceClass), err
}
