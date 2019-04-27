package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return RegisterDefaults(scheme)
}
func SetDefaults_ClusterServiceBrokerSpec(spec *ClusterServiceBrokerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	setCommonServiceBrokerDefaults(&spec.CommonServiceBrokerSpec)
}
func SetDefaults_ServiceBrokerSpec(spec *ServiceBrokerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	setCommonServiceBrokerDefaults(&spec.CommonServiceBrokerSpec)
}
func setCommonServiceBrokerDefaults(spec *CommonServiceBrokerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec.RelistBehavior == "" {
		spec.RelistBehavior = ServiceBrokerRelistBehaviorDuration
	}
}
func SetDefaults_ServiceBinding(binding *ServiceBinding) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if binding.Spec.SecretName == "" {
		binding.Spec.SecretName = binding.Name
	}
}
