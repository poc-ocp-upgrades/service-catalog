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

type FakeServiceClasses struct {
	Fake	*FakeServicecatalogV1beta1
	ns	string
}

var serviceclassesResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "serviceclasses"}
var serviceclassesKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ServiceClass"}

func (c *FakeServiceClasses) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(serviceclassesResource, c.ns, name), &v1beta1.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}
func (c *FakeServiceClasses) List(opts v1.ListOptions) (result *v1beta1.ServiceClassList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(serviceclassesResource, serviceclassesKind, c.ns, opts), &v1beta1.ServiceClassList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceClassList{ListMeta: obj.(*v1beta1.ServiceClassList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServiceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(serviceclassesResource, c.ns, opts))
}
func (c *FakeServiceClasses) Create(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(serviceclassesResource, c.ns, serviceClass), &v1beta1.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}
func (c *FakeServiceClasses) Update(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(serviceclassesResource, c.ns, serviceClass), &v1beta1.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}
func (c *FakeServiceClasses) UpdateStatus(serviceClass *v1beta1.ServiceClass) (*v1beta1.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(serviceclassesResource, "status", c.ns, serviceClass), &v1beta1.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}
func (c *FakeServiceClasses) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(serviceclassesResource, c.ns, name), &v1beta1.ServiceClass{})
	return err
}
func (c *FakeServiceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(serviceclassesResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &v1beta1.ServiceClassList{})
	return err
}
func (c *FakeServiceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(serviceclassesResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}
