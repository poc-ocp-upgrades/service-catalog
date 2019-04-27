package v1beta1

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

type ServiceInstanceExpansion interface {
	UpdateReferences(serviceInstance *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error)
}

func (c *serviceInstances) UpdateReferences(serviceInstance *v1beta1.ServiceInstance) (result *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	result = &v1beta1.ServiceInstance{}
	err = c.client.Put().Namespace(serviceInstance.Namespace).Resource("serviceinstances").Name(serviceInstance.Name).SubResource("reference").Body(serviceInstance).Do().Into(result)
	return
}
