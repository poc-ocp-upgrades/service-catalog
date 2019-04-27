package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceBindingLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceBinding, err error)
	ServiceBindings(namespace string) ServiceBindingNamespaceLister
	ServiceBindingListerExpansion
}
type serviceBindingLister struct{ indexer cache.Indexer }

func NewServiceBindingLister(indexer cache.Indexer) ServiceBindingLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBindingLister{indexer: indexer}
}
func (s *serviceBindingLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceBinding))
	})
	return ret, err
}
func (s *serviceBindingLister) ServiceBindings(namespace string) ServiceBindingNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return serviceBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServiceBindingNamespaceLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceBinding, err error)
	Get(name string) (*servicecatalog.ServiceBinding, error)
	ServiceBindingNamespaceListerExpansion
}
type serviceBindingNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceBindingNamespaceLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceBinding))
	})
	return ret, err
}
func (s serviceBindingNamespaceLister) Get(name string) (*servicecatalog.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("servicebinding"), name)
	}
	return obj.(*servicecatalog.ServiceBinding), nil
}
