package v1beta1

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceClassLister interface {
	List(selector labels.Selector) (ret []*v1beta1.ServiceClass, err error)
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
func (s *serviceClassLister) List(selector labels.Selector) (ret []*v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceClass))
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
	List(selector labels.Selector) (ret []*v1beta1.ServiceClass, err error)
	Get(name string) (*v1beta1.ServiceClass, error)
	ServiceClassNamespaceListerExpansion
}
type serviceClassNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceClassNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceClass))
	})
	return ret, err
}
func (s serviceClassNamespaceLister) Get(name string) (*v1beta1.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("serviceclass"), name)
	}
	return obj.(*v1beta1.ServiceClass), nil
}
