package apiserver

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	servicecatalogrest "github.com/kubernetes-incubator/service-catalog/pkg/registry/servicecatalog/rest"
	settingsrest "github.com/kubernetes-incubator/service-catalog/pkg/registry/settings/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/pkg/version"
	restclient "k8s.io/client-go/rest"
)

const (
	apiServerName = "service-catalog-apiserver"
)

func restStorageProviders(defaultNamespace string, restClient restclient.Interface) []RESTStorageProvider {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return []RESTStorageProvider{servicecatalogrest.StorageProvider{DefaultNamespace: defaultNamespace, RESTClient: restClient}, settingsrest.StorageProvider{RESTClient: restClient}}
}
func completeGenericConfig(cfg *genericapiserver.RecommendedConfig) genericapiserver.CompletedConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cfg.Serializer = api.Codecs
	completedCfg := cfg.Complete()
	version := version.Get()
	cfg.Version = &version
	return completedCfg
}
func createSkeletonServer(genericCfg genericapiserver.CompletedConfig) (*ServiceCatalogAPIServer, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	genericServer, err := genericCfg.New(apiServerName, genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}
	return &ServiceCatalogAPIServer{GenericAPIServer: genericServer}, nil
}
