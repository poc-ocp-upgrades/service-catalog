package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "servicecatalog.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1beta1"}

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
	SchemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes, addDefaultingFuncs)
	localSchemeBuilder	= &SchemeBuilder
	AddToScheme			= SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &ClusterServiceBroker{}, &ClusterServiceBrokerList{}, &ServiceBroker{}, &ServiceBrokerList{}, &ClusterServiceClass{}, &ClusterServiceClassList{}, &ServiceClass{}, &ServiceClassList{}, &ClusterServicePlan{}, &ClusterServicePlanList{}, &ServicePlan{}, &ServicePlanList{}, &ServiceInstance{}, &ServiceInstanceList{}, &ServiceBinding{}, &ServiceBindingList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	scheme.AddKnownTypes(schema.GroupVersion{Version: "v1"}, &metav1.Status{})
	scheme.AddFieldLabelConversionFunc(serviceCatalogV1Beta1GVK("ClusterServiceClass"), ClusterServiceClassFieldLabelConversionFunc)
	scheme.AddFieldLabelConversionFunc(serviceCatalogV1Beta1GVK("ServiceClass"), ServiceClassFieldLabelConversionFunc)
	scheme.AddFieldLabelConversionFunc(serviceCatalogV1Beta1GVK("ClusterServicePlan"), ClusterServicePlanFieldLabelConversionFunc)
	scheme.AddFieldLabelConversionFunc(serviceCatalogV1Beta1GVK("ServicePlan"), ServicePlanFieldLabelConversionFunc)
	scheme.AddFieldLabelConversionFunc(serviceCatalogV1Beta1GVK("ServiceInstance"), ServiceInstanceFieldLabelConversionFunc)
	scheme.AddFieldLabelConversionFunc(serviceCatalogV1Beta1GVK("ServiceBinding"), ServiceBindingFieldLabelConversionFunc)
	return nil
}
func serviceCatalogV1Beta1GVK(kind string) schema.GroupVersionKind {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: kind}
}
