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

type FakeServiceInstances struct {
	Fake	*FakeServicecatalogV1beta1
	ns	string
}

var serviceinstancesResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "serviceinstances"}
var serviceinstancesKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ServiceInstance"}

func (c *FakeServiceInstances) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(serviceinstancesResource, c.ns, name), &v1beta1.ServiceInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceInstance), err
}
func (c *FakeServiceInstances) List(opts v1.ListOptions) (result *v1beta1.ServiceInstanceList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(serviceinstancesResource, serviceinstancesKind, c.ns, opts), &v1beta1.ServiceInstanceList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceInstanceList{ListMeta: obj.(*v1beta1.ServiceInstanceList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServiceInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(serviceinstancesResource, c.ns, opts))
}
func (c *FakeServiceInstances) Create(serviceInstance *v1beta1.ServiceInstance) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(serviceinstancesResource, c.ns, serviceInstance), &v1beta1.ServiceInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceInstance), err
}
func (c *FakeServiceInstances) Update(serviceInstance *v1beta1.ServiceInstance) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(serviceinstancesResource, c.ns, serviceInstance), &v1beta1.ServiceInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceInstance), err
}
func (c *FakeServiceInstances) UpdateStatus(serviceInstance *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(serviceinstancesResource, "status", c.ns, serviceInstance), &v1beta1.ServiceInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceInstance), err
}
func (c *FakeServiceInstances) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(serviceinstancesResource, c.ns, name), &v1beta1.ServiceInstance{})
	return err
}
func (c *FakeServiceInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(serviceinstancesResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ServiceInstanceList{})
	return err
}
func (c *FakeServiceInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(serviceinstancesResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceInstance), err
}
