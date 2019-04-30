package fake

import (
	rest "k8s.io/client-go/rest"
	servicecatalogv1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
)

type ServicecatalogV1beta1 struct {
	servicecatalogv1beta1.ServicecatalogV1beta1Interface
}

var _ servicecatalogv1beta1.ServicecatalogV1beta1Interface = &ServicecatalogV1beta1{}

func (c *ServicecatalogV1beta1) ClusterServiceBrokers() v1beta1.ClusterServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServicecatalogV1beta1Interface.ClusterServiceBrokers()
}
func (c *ServicecatalogV1beta1) ClusterServiceClasses() v1beta1.ClusterServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServicecatalogV1beta1Interface.ClusterServiceClasses()
}
func (c *ServicecatalogV1beta1) ServiceInstances(namespace string) v1beta1.ServiceInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceInstances := c.ServicecatalogV1beta1Interface.ServiceInstances(namespace)
	return &ServiceInstances{serviceInstances}
}
func (c *ServicecatalogV1beta1) ServiceBindings(namespace string) v1beta1.ServiceBindingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceBindings := c.ServicecatalogV1beta1Interface.ServiceBindings(namespace)
	return &ServiceBindings{serviceBindings}
}
func (c *ServicecatalogV1beta1) ClusterServicePlans() v1beta1.ClusterServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServicecatalogV1beta1Interface.ClusterServicePlans()
}
func (c *ServicecatalogV1beta1) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServicecatalogV1beta1Interface.RESTClient()
}
