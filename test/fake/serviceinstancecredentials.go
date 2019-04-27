package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1beta1typed "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
)

type ServiceBindings struct {
	v1beta1typed.ServiceBindingInterface
}

func (c *ServiceBindings) Create(serviceInstance *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.Create(serviceInstance)
}
func (c *ServiceBindings) Update(serviceInstance *v1beta1.ServiceBinding) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.Update(serviceInstance)
}
func (c *ServiceBindings) UpdateStatus(serviceInstance *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceCopy := serviceInstance.DeepCopy()
	_, err := c.ServiceBindingInterface.UpdateStatus(instanceCopy)
	return serviceInstance, err
}
func (c *ServiceBindings) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.Delete(name, options)
}
func (c *ServiceBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.DeleteCollection(options, listOptions)
}
func (c *ServiceBindings) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.Get(name, options)
}
func (c *ServiceBindings) List(opts v1.ListOptions) (result *v1beta1.ServiceBindingList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.List(opts)
}
func (c *ServiceBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.Watch(opts)
}
func (c *ServiceBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceBindingInterface.Patch(name, pt, data, subresources...)
}
