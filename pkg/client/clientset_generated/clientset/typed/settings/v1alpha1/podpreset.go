package v1alpha1

import (
	"time"
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	scheme "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

type PodPresetsGetter interface {
	PodPresets(namespace string) PodPresetInterface
}
type PodPresetInterface interface {
	Create(*v1alpha1.PodPreset) (*v1alpha1.PodPreset, error)
	Update(*v1alpha1.PodPreset) (*v1alpha1.PodPreset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PodPreset, error)
	List(opts v1.ListOptions) (*v1alpha1.PodPresetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodPreset, err error)
	PodPresetExpansion
}
type podPresets struct {
	client	rest.Interface
	ns	string
}

func newPodPresets(c *SettingsV1alpha1Client, namespace string) *podPresets {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &podPresets{client: c.RESTClient(), ns: namespace}
}
func (c *podPresets) Get(name string, options v1.GetOptions) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1alpha1.PodPreset{}
	err = c.client.Get().Namespace(c.ns).Resource("podpresets").Name(name).VersionedParams(&options, scheme.ParameterCodec).Do().Into(result)
	return
}
func (c *podPresets) List(opts v1.ListOptions) (result *v1alpha1.PodPresetList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PodPresetList{}
	err = c.client.Get().Namespace(c.ns).Resource("podpresets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Do().Into(result)
	return
}
func (c *podPresets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().Namespace(c.ns).Resource("podpresets").VersionedParams(&opts, scheme.ParameterCodec).Timeout(timeout).Watch()
}
func (c *podPresets) Create(podPreset *v1alpha1.PodPreset) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1alpha1.PodPreset{}
	err = c.client.Post().Namespace(c.ns).Resource("podpresets").Body(podPreset).Do().Into(result)
	return
}
func (c *podPresets) Update(podPreset *v1alpha1.PodPreset) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1alpha1.PodPreset{}
	err = c.client.Put().Namespace(c.ns).Resource("podpresets").Name(podPreset.Name).Body(podPreset).Do().Into(result)
	return
}
func (c *podPresets) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.client.Delete().Namespace(c.ns).Resource("podpresets").Name(name).Body(options).Do().Error()
}
func (c *podPresets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().Namespace(c.ns).Resource("podpresets").VersionedParams(&listOptions, scheme.ParameterCodec).Timeout(timeout).Body(options).Do().Error()
}
func (c *podPresets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1alpha1.PodPreset{}
	err = c.client.Patch(pt).Namespace(c.ns).Resource("podpresets").SubResource(subresources...).Name(name).Body(data).Do().Into(result)
	return
}
