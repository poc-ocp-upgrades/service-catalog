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

type FakeServiceBrokers struct {
	Fake	*FakeServicecatalogV1beta1
	ns	string
}

var servicebrokersResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "servicebrokers"}
var servicebrokersKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ServiceBroker"}

func (c *FakeServiceBrokers) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(servicebrokersResource, c.ns, name), &v1beta1.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBroker), err
}
func (c *FakeServiceBrokers) List(opts v1.ListOptions) (result *v1beta1.ServiceBrokerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(servicebrokersResource, servicebrokersKind, c.ns, opts), &v1beta1.ServiceBrokerList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceBrokerList{ListMeta: obj.(*v1beta1.ServiceBrokerList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceBrokerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServiceBrokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(servicebrokersResource, c.ns, opts))
}
func (c *FakeServiceBrokers) Create(serviceBroker *v1beta1.ServiceBroker) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(servicebrokersResource, c.ns, serviceBroker), &v1beta1.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBroker), err
}
func (c *FakeServiceBrokers) Update(serviceBroker *v1beta1.ServiceBroker) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(servicebrokersResource, c.ns, serviceBroker), &v1beta1.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBroker), err
}
func (c *FakeServiceBrokers) UpdateStatus(serviceBroker *v1beta1.ServiceBroker) (*v1beta1.ServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(servicebrokersResource, "status", c.ns, serviceBroker), &v1beta1.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBroker), err
}
func (c *FakeServiceBrokers) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(servicebrokersResource, c.ns, name), &v1beta1.ServiceBroker{})
	return err
}
func (c *FakeServiceBrokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(servicebrokersResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ServiceBrokerList{})
	return err
}
func (c *FakeServiceBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(servicebrokersResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceBroker), err
}
