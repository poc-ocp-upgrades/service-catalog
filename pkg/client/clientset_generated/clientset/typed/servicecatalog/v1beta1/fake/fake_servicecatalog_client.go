package fake

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServicecatalogV1beta1 struct{ *testing.Fake }

func (c *FakeServicecatalogV1beta1) ClusterServiceBrokers() v1beta1.ClusterServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeClusterServiceBrokers{c}
}
func (c *FakeServicecatalogV1beta1) ClusterServiceClasses() v1beta1.ClusterServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeClusterServiceClasses{c}
}
func (c *FakeServicecatalogV1beta1) ClusterServicePlans() v1beta1.ClusterServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeClusterServicePlans{c}
}
func (c *FakeServicecatalogV1beta1) ServiceBindings(namespace string) v1beta1.ServiceBindingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceBindings{c, namespace}
}
func (c *FakeServicecatalogV1beta1) ServiceBrokers(namespace string) v1beta1.ServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceBrokers{c, namespace}
}
func (c *FakeServicecatalogV1beta1) ServiceClasses(namespace string) v1beta1.ServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceClasses{c, namespace}
}
func (c *FakeServicecatalogV1beta1) ServiceInstances(namespace string) v1beta1.ServiceInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceInstances{c, namespace}
}
func (c *FakeServicecatalogV1beta1) ServicePlans(namespace string) v1beta1.ServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServicePlans{c, namespace}
}
func (c *FakeServicecatalogV1beta1) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
