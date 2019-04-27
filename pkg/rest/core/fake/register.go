package fake

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
)

var (
	Scheme		= runtime.NewScheme()
	ParameterCodec	= runtime.NewParameterCodec(Scheme)
	Codecs		= serializer.NewCodecFactory(Scheme)
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	install.Install(Scheme)
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned, &metav1.Status{}, &metav1.APIVersions{}, &metav1.APIGroupList{}, &metav1.APIGroup{}, &metav1.APIResourceList{})
}
