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

type FakeServiceBindings struct {
	Fake	*FakeServicecatalogV1beta1
	ns	string
}

var servicebindingsResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "servicebindings"}
var servicebindingsKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ServiceBinding"}

func (c *FakeServiceBindings) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(servicebindingsResource, c.ns, name), &v1beta1.ServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBinding), err
}
func (c *FakeServiceBindings) List(opts v1.ListOptions) (result *v1beta1.ServiceBindingList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(servicebindingsResource, servicebindingsKind, c.ns, opts), &v1beta1.ServiceBindingList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceBindingList{ListMeta: obj.(*v1beta1.ServiceBindingList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServiceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(servicebindingsResource, c.ns, opts))
}
func (c *FakeServiceBindings) Create(serviceBinding *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(servicebindingsResource, c.ns, serviceBinding), &v1beta1.ServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBinding), err
}
func (c *FakeServiceBindings) Update(serviceBinding *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(servicebindingsResource, c.ns, serviceBinding), &v1beta1.ServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBinding), err
}
func (c *FakeServiceBindings) UpdateStatus(serviceBinding *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(servicebindingsResource, "status", c.ns, serviceBinding), &v1beta1.ServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBinding), err
}
func (c *FakeServiceBindings) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(servicebindingsResource, c.ns, name), &v1beta1.ServiceBinding{})
	return err
}
func (c *FakeServiceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(servicebindingsResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ServiceBindingList{})
	return err
}
func (c *FakeServiceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(servicebindingsResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBinding), err
}
