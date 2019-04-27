package apiserver

import (
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/storage"
	"k8s.io/klog"
)

type etcdConfig struct {
	genericConfig	*genericapiserver.RecommendedConfig
	extraConfig	*extraConfig
}
type extraConfig struct {
	deleteCollectionWorkers	int
	storageFactory		storage.StorageFactory
}

func NewEtcdConfig(genCfg *genericapiserver.RecommendedConfig, deleteCollWorkers int, factory storage.StorageFactory) Config {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &etcdConfig{genericConfig: genCfg, extraConfig: &extraConfig{deleteCollectionWorkers: deleteCollWorkers, storageFactory: factory}}
}
func (c *etcdConfig) Complete() CompletedConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	completedGenericConfig := completeGenericConfig(c.genericConfig)
	return completedEtcdConfig{genericConfig: completedGenericConfig, extraConfig: c.extraConfig, apiResourceConfigSource: DefaultAPIResourceConfigSource()}
}

type completedEtcdConfig struct {
	genericConfig		genericapiserver.CompletedConfig
	extraConfig		*extraConfig
	apiResourceConfigSource	storage.APIResourceConfigSource
}

func (c completedEtcdConfig) NewServer(stopCh <-chan struct{}) (*ServiceCatalogAPIServer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s, err := createSkeletonServer(c.genericConfig)
	if err != nil {
		return nil, err
	}
	klog.V(4).Infoln("Created skeleton API server")
	roFactory := etcdRESTOptionsFactory{deleteCollectionWorkers: c.extraConfig.deleteCollectionWorkers, enableGarbageCollection: true, storageFactory: c.extraConfig.storageFactory, storageDecorator: generic.UndecoratedStorage}
	klog.V(4).Infoln("Installing API groups")
	providers := restStorageProviders("", nil)
	for _, provider := range providers {
		groupInfo, err := provider.NewRESTStorage(c.apiResourceConfigSource, roFactory)
		if IsErrAPIGroupDisabled(err) {
			klog.Warningf("Skipping API group %v because it is not enabled", provider.GroupName())
			continue
		} else if err != nil {
			klog.Errorf("Error initializing storage for provider %v: %v", provider.GroupName(), err)
			return nil, err
		}
		klog.V(4).Infof("Installing API group %v", provider.GroupName())
		if err := s.GenericAPIServer.InstallAPIGroup(groupInfo); err != nil {
			klog.Fatalf("Error installing API group %v: %v", provider.GroupName(), err)
		} else {
			for _, mappings := range groupInfo.VersionedResourcesStorageMap {
				for _, storage := range mappings {
					go func(store rest.Storage) {
						s, ok := store.(*registry.Store)
						if ok {
							<-stopCh
							s.DestroyFunc()
						}
					}(storage)
				}
			}
		}
	}
	klog.Infoln("Finished installing API groups")
	return s, nil
}
