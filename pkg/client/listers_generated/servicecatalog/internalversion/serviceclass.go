package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceClassLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceClass, err error)
	ServiceClasses(namespace string) ServiceClassNamespaceLister
	ServiceClassListerExpansion
}
type serviceClassLister struct{ indexer cache.Indexer }

func NewServiceClassLister(indexer cache.Indexer) ServiceClassLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceClassLister{indexer: indexer}
}
func (s *serviceClassLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceClass))
	})
	return ret, err
}
func (s *serviceClassLister) ServiceClasses(namespace string) ServiceClassNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return serviceClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServiceClassNamespaceLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceClass, err error)
	Get(name string) (*servicecatalog.ServiceClass, error)
	ServiceClassNamespaceListerExpansion
}
type serviceClassNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceClassNamespaceLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceClass))
	})
	return ret, err
}
func (s serviceClassNamespaceLister) Get(name string) (*servicecatalog.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("serviceclass"), name)
	}
	return obj.(*servicecatalog.ServiceClass), nil
}
