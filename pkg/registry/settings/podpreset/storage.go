package podpreset

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	scmeta "github.com/kubernetes-incubator/service-catalog/pkg/api/meta"
	settingsapi "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
	"github.com/kubernetes-incubator/service-catalog/pkg/registry/servicecatalog/server"
)

var (
	errNotAPodPreset = errors.New("not a podpreset")
)

func EmptyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &settingsapi.PodPreset{}
}
func NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &settingsapi.PodPresetList{}
}
func NewStorage(opts server.Options) (rest.Storage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := "/" + opts.ResourcePrefix()
	storageInterface, dFunc := opts.GetStorage(&settingsapi.PodPreset{}, prefix, podPresetRESTStrategy, NewList, nil, storage.NoTriggerPublisher)
	store := genericregistry.Store{NewFunc: func() runtime.Object {
		return &settingsapi.PodPreset{}
	}, NewListFunc: NewList, KeyRootFunc: opts.KeyRootFunc(), KeyFunc: opts.KeyFunc(true), ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return scmeta.GetAccessor().Name(obj)
	}, PredicateFunc: Matcher, DefaultQualifiedResource: settingsapi.Resource("podpresets"), CreateStrategy: podPresetRESTStrategy, UpdateStrategy: podPresetRESTStrategy, DeleteStrategy: podPresetRESTStrategy, EnableGarbageCollection: true, Storage: storageInterface, DestroyFunc: dFunc}
	return &store, nil
}
