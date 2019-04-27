package apiserver

import (
	"fmt"
	"k8s.io/apiserver/pkg/registry/generic"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/storage"
)

type ErrAPIGroupDisabled struct{ Name string }

func (e ErrAPIGroupDisabled) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("API group %s is disabled", e.Name)
}
func IsErrAPIGroupDisabled(e error) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := e.(ErrAPIGroupDisabled)
	return ok
}

type RESTStorageProvider interface {
	GroupName() string
	NewRESTStorage(apiResourceConfigSource storage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (*genericapiserver.APIGroupInfo, error)
}
