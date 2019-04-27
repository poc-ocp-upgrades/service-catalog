package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func RegisterDefaults(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddTypeDefaultingFunc(&ClusterServiceBroker{}, func(obj interface{}) {
		SetObjectDefaults_ClusterServiceBroker(obj.(*ClusterServiceBroker))
	})
	scheme.AddTypeDefaultingFunc(&ClusterServiceBrokerList{}, func(obj interface{}) {
		SetObjectDefaults_ClusterServiceBrokerList(obj.(*ClusterServiceBrokerList))
	})
	scheme.AddTypeDefaultingFunc(&ServiceBinding{}, func(obj interface{}) {
		SetObjectDefaults_ServiceBinding(obj.(*ServiceBinding))
	})
	scheme.AddTypeDefaultingFunc(&ServiceBindingList{}, func(obj interface{}) {
		SetObjectDefaults_ServiceBindingList(obj.(*ServiceBindingList))
	})
	scheme.AddTypeDefaultingFunc(&ServiceBroker{}, func(obj interface{}) {
		SetObjectDefaults_ServiceBroker(obj.(*ServiceBroker))
	})
	scheme.AddTypeDefaultingFunc(&ServiceBrokerList{}, func(obj interface{}) {
		SetObjectDefaults_ServiceBrokerList(obj.(*ServiceBrokerList))
	})
	return nil
}
func SetObjectDefaults_ClusterServiceBroker(in *ClusterServiceBroker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_ClusterServiceBrokerSpec(&in.Spec)
}
func SetObjectDefaults_ClusterServiceBrokerList(in *ClusterServiceBrokerList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ClusterServiceBroker(a)
	}
}
func SetObjectDefaults_ServiceBinding(in *ServiceBinding) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_ServiceBinding(in)
}
func SetObjectDefaults_ServiceBindingList(in *ServiceBindingList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ServiceBinding(a)
	}
}
func SetObjectDefaults_ServiceBroker(in *ServiceBroker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SetDefaults_ServiceBrokerSpec(&in.Spec)
}
func SetObjectDefaults_ServiceBrokerList(in *ServiceBrokerList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ServiceBroker(a)
	}
}
