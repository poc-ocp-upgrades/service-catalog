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

type FakeServicePlans struct {
	Fake	*FakeServicecatalog
	ns	string
}

var serviceplansResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "", Resource: "serviceplans"}
var serviceplansKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "", Kind: "ServicePlan"}

func (c *FakeServicePlans) Get(name string, options v1.GetOptions) (result *servicecatalog.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(serviceplansResource, c.ns, name), &servicecatalog.ServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}
func (c *FakeServicePlans) List(opts v1.ListOptions) (result *servicecatalog.ServicePlanList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(serviceplansResource, serviceplansKind, c.ns, opts), &servicecatalog.ServicePlanList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicecatalog.ServicePlanList{ListMeta: obj.(*servicecatalog.ServicePlanList).ListMeta}
	for _, item := range obj.(*servicecatalog.ServicePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakeServicePlans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(serviceplansResource, c.ns, opts))
}
func (c *FakeServicePlans) Create(servicePlan *servicecatalog.ServicePlan) (result *servicecatalog.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(serviceplansResource, c.ns, servicePlan), &servicecatalog.ServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}
func (c *FakeServicePlans) Update(servicePlan *servicecatalog.ServicePlan) (result *servicecatalog.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(serviceplansResource, c.ns, servicePlan), &servicecatalog.ServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}
func (c *FakeServicePlans) UpdateStatus(servicePlan *servicecatalog.ServicePlan) (*servicecatalog.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(serviceplansResource, "status", c.ns, servicePlan), &servicecatalog.ServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}
func (c *FakeServicePlans) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(serviceplansResource, c.ns, name), &servicecatalog.ServicePlan{})
	return err
}
func (c *FakeServicePlans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(serviceplansResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &servicecatalog.ServicePlanList{})
	return err
}
func (c *FakeServicePlans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *servicecatalog.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(serviceplansResource, c.ns, name, pt, data, subresources...), &servicecatalog.ServicePlan{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicecatalog.ServicePlan), err
}
