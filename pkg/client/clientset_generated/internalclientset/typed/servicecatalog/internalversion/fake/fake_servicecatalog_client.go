package fake

import (
	internalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/typed/servicecatalog/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeServicecatalog struct{ *testing.Fake }

func (c *FakeServicecatalog) ClusterServiceBrokers() internalversion.ClusterServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeClusterServiceBrokers{c}
}
func (c *FakeServicecatalog) ClusterServiceClasses() internalversion.ClusterServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeClusterServiceClasses{c}
}
func (c *FakeServicecatalog) ClusterServicePlans() internalversion.ClusterServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeClusterServicePlans{c}
}
func (c *FakeServicecatalog) ServiceBindings(namespace string) internalversion.ServiceBindingInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceBindings{c, namespace}
}
func (c *FakeServicecatalog) ServiceBrokers(namespace string) internalversion.ServiceBrokerInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceBrokers{c, namespace}
}
func (c *FakeServicecatalog) ServiceClasses(namespace string) internalversion.ServiceClassInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceClasses{c, namespace}
}
func (c *FakeServicecatalog) ServiceInstances(namespace string) internalversion.ServiceInstanceInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServiceInstances{c, namespace}
}
func (c *FakeServicecatalog) ServicePlans(namespace string) internalversion.ServicePlanInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakeServicePlans{c, namespace}
}
func (c *FakeServicecatalog) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
