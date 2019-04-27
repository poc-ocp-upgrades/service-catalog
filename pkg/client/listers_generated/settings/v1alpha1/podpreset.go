package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type PodPresetLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error)
	PodPresets(namespace string) PodPresetNamespaceLister
	PodPresetListerExpansion
}
type podPresetLister struct{ indexer cache.Indexer }

func NewPodPresetLister(indexer cache.Indexer) PodPresetLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &podPresetLister{indexer: indexer}
}
func (s *podPresetLister) List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodPreset))
	})
	return ret, err
}
func (s *podPresetLister) PodPresets(namespace string) PodPresetNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return podPresetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type PodPresetNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error)
	Get(name string) (*v1alpha1.PodPreset, error)
	PodPresetNamespaceListerExpansion
}
type podPresetNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s podPresetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodPreset, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodPreset))
	})
	return ret, err
}
func (s podPresetNamespaceLister) Get(name string) (*v1alpha1.PodPreset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podpreset"), name)
	}
	return obj.(*v1alpha1.PodPreset), nil
}
