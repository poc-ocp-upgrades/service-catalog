package servicecatalog

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "servicecatalog.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}

func Kind(kind string) schema.GroupKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}
func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder	= &SchemeBuilder
	AddToScheme			= SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &ClusterServiceBroker{}, &ClusterServiceBrokerList{}, &ServiceBroker{}, &ServiceBrokerList{}, &ClusterServiceClass{}, &ClusterServiceClassList{}, &ServiceClass{}, &ServiceClassList{}, &ClusterServicePlan{}, &ClusterServicePlanList{}, &ServicePlan{}, &ServicePlanList{}, &ServiceInstance{}, &ServiceInstanceList{}, &ServiceBinding{}, &ServiceBindingList{})
	return nil
}
