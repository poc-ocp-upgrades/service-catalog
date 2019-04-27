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

type FakeServiceBrokers struct {
	Fake	*FakeServicecatalog
	ns	string
}

var servicebrokersResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "", Resource: "servicebrokers"}
var servicebrokersKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "", Kind: "ServiceBroker"}

func (c *FakeServiceBrokers) Get(name string, options v1.GetOptions) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(servicebrokersResource, c.ns, name), &servicecatalog.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceBroker), err
}
func (c *FakeServiceBrokers) List(opts v1.ListOptions) (result *servicecatalog.ServiceBrokerList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(servicebrokersResource, servicebrokersKind, c.ns, opts), &servicecatalog.ServiceBrokerList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.ServiceBrokerList{ListMeta: obj.(*servicecatalog.ServiceBrokerList).ListMeta}
	for _, item := range obj.(*servicecatalog.ServiceBrokerList).Items {
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
func (c *FakeServiceBrokers) Create(serviceBroker *servicecatalog.ServiceBroker) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(servicebrokersResource, c.ns, serviceBroker), &servicecatalog.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceBroker), err
}
func (c *FakeServiceBrokers) Update(serviceBroker *servicecatalog.ServiceBroker) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(servicebrokersResource, c.ns, serviceBroker), &servicecatalog.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceBroker), err
}
func (c *FakeServiceBrokers) UpdateStatus(serviceBroker *servicecatalog.ServiceBroker) (*servicecatalog.ServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(servicebrokersResource, "status", c.ns, serviceBroker), &servicecatalog.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceBroker), err
}
func (c *FakeServiceBrokers) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(servicebrokersResource, c.ns, name), &servicecatalog.ServiceBroker{})
	return err
}
func (c *FakeServiceBrokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(servicebrokersResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &servicecatalog.ServiceBrokerList{})
	return err
}
func (c *FakeServiceBrokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(servicebrokersResource, c.ns, name, pt, data, subresources...), &servicecatalog.ServiceBroker{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceBroker), err
}
