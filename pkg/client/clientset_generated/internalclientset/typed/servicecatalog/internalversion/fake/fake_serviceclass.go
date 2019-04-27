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

type FakeServiceClasses struct {
	Fake	*FakeServicecatalog
	ns	string
}

var serviceclassesResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "", Resource: "serviceclasses"}
var serviceclassesKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "", Kind: "ServiceClass"}

func (c *FakeServiceClasses) Get(name string, options v1.GetOptions) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(serviceclassesResource, c.ns, name), &servicecatalog.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceClass), err
}
func (c *FakeServiceClasses) List(opts v1.ListOptions) (result *servicecatalog.ServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(serviceclassesResource, serviceclassesKind, c.ns, opts), &servicecatalog.ServiceClassList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.ServiceClassList{ListMeta: obj.(*servicecatalog.ServiceClassList).ListMeta}
	for _, item := range obj.(*servicecatalog.ServiceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServiceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(serviceclassesResource, c.ns, opts))
}
func (c *FakeServiceClasses) Create(serviceClass *servicecatalog.ServiceClass) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(serviceclassesResource, c.ns, serviceClass), &servicecatalog.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceClass), err
}
func (c *FakeServiceClasses) Update(serviceClass *servicecatalog.ServiceClass) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(serviceclassesResource, c.ns, serviceClass), &servicecatalog.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceClass), err
}
func (c *FakeServiceClasses) UpdateStatus(serviceClass *servicecatalog.ServiceClass) (*servicecatalog.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(serviceclassesResource, "status", c.ns, serviceClass), &servicecatalog.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceClass), err
}
func (c *FakeServiceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(serviceclassesResource, c.ns, name), &servicecatalog.ServiceClass{})
	return err
}
func (c *FakeServiceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(serviceclassesResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &servicecatalog.ServiceClassList{})
	return err
}
func (c *FakeServiceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(serviceclassesResource, c.ns, name, pt, data, subresources...), &servicecatalog.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServiceClass), err
}
