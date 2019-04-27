package meta

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	accessor	= meta.NewAccessor()
	selfLinker	= runtime.SelfLinker(accessor)
)

func GetAccessor() meta.MetadataAccessor {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return accessor
}
