package v1beta1

import (
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/externalversions/internalinterfaces"
)

type Interface interface {
	ClusterServiceBrokers() ClusterServiceBrokerInformer
	ClusterServiceClasses() ClusterServiceClassInformer
	ClusterServicePlans() ClusterServicePlanInformer
	ServiceBindings() ServiceBindingInformer
	ServiceBrokers() ServiceBrokerInformer
	ServiceClasses() ServiceClassInformer
	ServiceInstances() ServiceInstanceInformer
	ServicePlans() ServicePlanInformer
}
type version struct {
	factory			internalinterfaces.SharedInformerFactory
	namespace		string
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
}

func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}
func (v *version) ClusterServiceBrokers() ClusterServiceBrokerInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceBrokerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) ClusterServiceClasses() ClusterServiceClassInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) ClusterServicePlans() ClusterServicePlanInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServicePlanInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
func (v *version) ServiceBindings() ServiceBindingInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) ServiceBrokers() ServiceBrokerInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBrokerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) ServiceClasses() ServiceClassInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) ServiceInstances() ServiceInstanceInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
func (v *version) ServicePlans() ServicePlanInformer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicePlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
