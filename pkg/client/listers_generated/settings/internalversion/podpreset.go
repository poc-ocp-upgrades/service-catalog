package internalversion

import (
	settings "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type PodPresetLister interface {
	List(selector labels.Selector) (ret []*settings.PodPreset, err error)
	PodPresets(namespace string) PodPresetNamespaceLister
	PodPresetListerExpansion
}
type podPresetLister struct{ indexer cache.Indexer }

func NewPodPresetLister(indexer cache.Indexer) PodPresetLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &podPresetLister{indexer: indexer}
}
func (s *podPresetLister) List(selector labels.Selector) (ret []*settings.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*settings.PodPreset))
	})
	return ret, err
}
func (s *podPresetLister) PodPresets(namespace string) PodPresetNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return podPresetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type PodPresetNamespaceLister interface {
	List(selector labels.Selector) (ret []*settings.PodPreset, err error)
	Get(name string) (*settings.PodPreset, error)
	PodPresetNamespaceListerExpansion
}
type podPresetNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s podPresetNamespaceLister) List(selector labels.Selector) (ret []*settings.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*settings.PodPreset))
	})
	return ret, err
}
func (s podPresetNamespaceLister) Get(name string) (*settings.PodPreset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(settings.Resource("podpreset"), name)
	}
	return obj.(*settings.PodPreset), nil
}
