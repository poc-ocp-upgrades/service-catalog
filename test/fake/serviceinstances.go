package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	watch "k8s.io/apimachinery/pkg/watch"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1beta1typed "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
)

type ServiceInstances struct {
	v1beta1typed.ServiceInstanceInterface
}

func (c *ServiceInstances) Create(serviceInstance *v1beta1.ServiceInstance) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.Create(serviceInstance)
}
func (c *ServiceInstances) Update(serviceInstance *v1beta1.ServiceInstance) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceCopy := serviceInstance.DeepCopy()
	updatedInstance, err := c.ServiceInstanceInterface.Update(instanceCopy)
	if updatedInstance != nil {
		updatedInstance.ResourceVersion = rand.String(10)
	}
	return updatedInstance, err
}
func (c *ServiceInstances) UpdateStatus(serviceInstance *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceCopy := serviceInstance.DeepCopy()
	updatedInstance, err := c.ServiceInstanceInterface.UpdateStatus(instanceCopy)
	if updatedInstance != nil {
		updatedInstance.ResourceVersion = rand.String(10)
	}
	return updatedInstance, err
}
func (c *ServiceInstances) Delete(name string, options *v1.DeleteOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.Delete(name, options)
}
func (c *ServiceInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.DeleteCollection(options, listOptions)
}
func (c *ServiceInstances) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.Get(name, options)
}
func (c *ServiceInstances) List(opts v1.ListOptions) (result *v1beta1.ServiceInstanceList, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.List(opts)
}
func (c *ServiceInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.Watch(opts)
}
func (c *ServiceInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServiceInstanceInterface.Patch(name, pt, data, subresources...)
}
