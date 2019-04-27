package apiserver

import (
	"fmt"
	"strconv"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/server/resourceconfig"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	utilflag "k8s.io/apiserver/pkg/util/flag"
)

func NewStorageFactory(storageConfig storagebackend.Config, defaultMediaType string, serializer runtime.StorageSerializer, defaultResourceEncoding *serverstorage.DefaultResourceEncodingConfig, storageEncodingOverrides map[string]schema.GroupVersion, resourceEncodingOverrides []schema.GroupVersionResource, defaultAPIResourceConfig *serverstorage.ResourceConfig, resourceConfigOverrides utilflag.ConfigurationMap) (*serverstorage.DefaultStorageFactory, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceEncodingConfig := mergeGroupEncodingConfigs(defaultResourceEncoding, storageEncodingOverrides)
	resourceEncodingConfig = mergeResourceEncodingConfigs(resourceEncodingConfig, resourceEncodingOverrides)
	apiResourceConfig, err := resourceconfig.MergeAPIResourceConfigs(defaultAPIResourceConfig, resourceConfigOverrides, api.Scheme)
	if err != nil {
		return nil, err
	}
	return serverstorage.NewDefaultStorageFactory(storageConfig, defaultMediaType, serializer, resourceEncodingConfig, apiResourceConfig, nil), nil
}
func mergeResourceEncodingConfigs(defaultResourceEncoding *serverstorage.DefaultResourceEncodingConfig, resourceEncodingOverrides []schema.GroupVersionResource) *serverstorage.DefaultResourceEncodingConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceEncodingConfig := defaultResourceEncoding
	for _, gvr := range resourceEncodingOverrides {
		resourceEncodingConfig.SetResourceEncoding(gvr.GroupResource(), gvr.GroupVersion(), schema.GroupVersion{Group: gvr.Group, Version: runtime.APIVersionInternal})
	}
	return resourceEncodingConfig
}
func mergeGroupEncodingConfigs(defaultResourceEncoding *serverstorage.DefaultResourceEncodingConfig, storageEncodingOverrides map[string]schema.GroupVersion) *serverstorage.DefaultResourceEncodingConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	resourceEncodingConfig := defaultResourceEncoding
	for group, storageEncodingVersion := range storageEncodingOverrides {
		resourceEncodingConfig.SetVersionEncoding(group, storageEncodingVersion, schema.GroupVersion{Group: group, Version: runtime.APIVersionInternal})
	}
	return resourceEncodingConfig
}
func getRuntimeConfigValue(overrides utilflag.ConfigurationMap, apiKey string, defaultValue bool) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	flagValue, ok := overrides[apiKey]
	if ok {
		if flagValue == "" {
			return true, nil
		}
		boolValue, err := strconv.ParseBool(flagValue)
		if err != nil {
			return false, fmt.Errorf("invalid value of %s: %s, err: %v", apiKey, flagValue, err)
		}
		return boolValue, nil
	}
	return defaultValue, nil
}
