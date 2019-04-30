package fake

import (
	servicecataloginternalversion "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	settingsinternalversion "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)
var localSchemeBuilder = runtime.SchemeBuilder{servicecataloginternalversion.AddToScheme, settingsinternalversion.AddToScheme}
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
