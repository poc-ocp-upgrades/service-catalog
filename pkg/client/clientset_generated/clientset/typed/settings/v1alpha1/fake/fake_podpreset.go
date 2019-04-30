package fake

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

type FakePodPresets struct {
	Fake	*FakeSettingsV1alpha1
	ns	string
}

var podpresetsResource = schema.GroupVersionResource{Group: "settings.servicecatalog.k8s.io", Version: "v1alpha1", Resource: "podpresets"}
var podpresetsKind = schema.GroupVersionKind{Group: "settings.servicecatalog.k8s.io", Version: "v1alpha1", Kind: "PodPreset"}

func (c *FakePodPresets) Get(name string, options v1.GetOptions) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewGetAction(podpresetsResource, c.ns, name), &v1alpha1.PodPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}
func (c *FakePodPresets) List(opts v1.ListOptions) (result *v1alpha1.PodPresetList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewListAction(podpresetsResource, podpresetsKind, c.ns, opts), &v1alpha1.PodPresetList{})
	if obj == nil {
		return nil, err
	}
	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodPresetList{ListMeta: obj.(*v1alpha1.PodPresetList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodPresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}
func (c *FakePodPresets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Fake.InvokesWatch(testing.NewWatchAction(podpresetsResource, c.ns, opts))
}
func (c *FakePodPresets) Create(podPreset *v1alpha1.PodPreset) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewCreateAction(podpresetsResource, c.ns, podPreset), &v1alpha1.PodPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}
func (c *FakePodPresets) Update(podPreset *v1alpha1.PodPreset) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewUpdateAction(podpresetsResource, c.ns, podPreset), &v1alpha1.PodPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}
func (c *FakePodPresets) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := c.Fake.Invokes(testing.NewDeleteAction(podpresetsResource, c.ns, name), &v1alpha1.PodPreset{})
	return err
}
func (c *FakePodPresets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	action := testing.NewDeleteCollectionAction(podpresetsResource, c.ns, listOptions)
	_, err := c.Fake.Invokes(action, &v1alpha1.PodPresetList{})
	return err
}
func (c *FakePodPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, err := c.Fake.Invokes(testing.NewPatchSubresourceAction(podpresetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PodPreset{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPreset), err
}
