package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "settings.servicecatalog.k8s.io"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

func Resource(resource string) schema.GroupResource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder		= runtime.NewSchemeBuilder(addKnownTypes, addDefaultingFuncs)
	localSchemeBuilder	= &SchemeBuilder
	AddToScheme		= SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scheme.AddKnownTypes(SchemeGroupVersion, &PodPreset{}, &PodPresetList{})
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	scheme.AddKnownTypes(schema.GroupVersion{Version: "v1"}, &metav1.Status{})
	return nil
}
