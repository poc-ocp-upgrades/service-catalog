package servicecatalog

import (
	"fmt"
	"k8s.io/apimachinery/pkg/version"
)

func (sdk *SDK) ServerVersion() (*version.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serverVersion, err := sdk.ServiceCatalogClient.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("unable to get version, %v", err)
	}
	return serverVersion, nil
}
