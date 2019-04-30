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

type FakeClusterServiceBrokers struct{ Fake *FakeServicecatalogV1beta1 }

var clusterservicebrokersResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "clusterservicebrokers"}
var clusterservicebrokersKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ClusterServiceBroker"}

func (c *FakeClusterServiceBrokers) Get(name string, options v1.GetOptions) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootGetAction(clusterservicebrokersResource, name), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}
func (c *FakeClusterServiceBrokers) List(opts v1.ListOptions) (result *v1beta1.ClusterServiceBrokerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootListAction(clusterservicebrokersResource, clusterservicebrokersKind, opts), &v1beta1.ClusterServiceBrokerList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterServiceBrokerList{ListMeta: obj.(*v1beta1.ClusterServiceBrokerList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterServiceBrokerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeClusterServiceBrokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewRootWatchAction(clusterservicebrokersResource, opts))
}
func (c *FakeClusterServiceBrokers) Create(clusterServiceBroker *v1beta1.ClusterServiceBroker) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootCreateAction(clusterservicebrokersResource, clusterServiceBroker), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}
func (c *FakeClusterServiceBrokers) Update(clusterServiceBroker *v1beta1.ClusterServiceBroker) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateAction(clusterservicebrokersResource, clusterServiceBroker), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}
func (c *FakeClusterServiceBrokers) UpdateStatus(clusterServiceBroker *v1beta1.ClusterServiceBroker) (*v1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootUpdateSubresourceAction(clusterservicebrokersResource, "status", clusterServiceBroker), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}
func (c *FakeClusterServiceBrokers) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewRootDeleteAction(clusterservicebrokersResource, name), &v1beta1.ClusterServiceBroker{})
	return err
}
func (c *FakeClusterServiceBrokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewRootDeleteCollectionAction(clusterservicebrokersResource, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ClusterServiceBrokerList{})
	return err
}
func (c *FakeClusterServiceBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewRootPatchSubresourceAction(clusterservicebrokersResource, name, pt, data, subresources...), &v1beta1.ClusterServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterServiceBroker), err
}
