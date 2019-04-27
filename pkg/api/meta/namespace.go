package meta

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func GetNamespace(obj runtime.Object) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return selfLinker.Namespace(obj)
}
