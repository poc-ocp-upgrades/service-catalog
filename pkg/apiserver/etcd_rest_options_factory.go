package apiserver

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/server/storage"
)

type etcdRESTOptionsFactory struct {
	deleteCollectionWorkers	int
	enableGarbageCollection	bool
	storageFactory		storage.StorageFactory
	storageDecorator	generic.StorageDecorator
}

func (f etcdRESTOptionsFactory) GetRESTOptions(resource schema.GroupResource) (generic.RESTOptions, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	storageConfig, err := f.storageFactory.NewConfig(resource)
	if err != nil {
		return generic.RESTOptions{}, err
	}
	return generic.RESTOptions{StorageConfig: storageConfig, Decorator: f.storageDecorator, DeleteCollectionWorkers: f.deleteCollectionWorkers, EnableGarbageCollection: f.enableGarbageCollection, ResourcePrefix: resource.Group + "/" + resource.Resource}, nil
}
