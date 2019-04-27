package v1beta1

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceBindingLister interface {
	List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error)
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
func (s *serviceBindingLister) List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceBinding))
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
	List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error)
	Get(name string) (*v1beta1.ServiceBinding, error)
	ServiceBindingNamespaceListerExpansion
}
type serviceBindingNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ServiceBinding, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceBinding))
	})
	return ret, err
}
func (s serviceBindingNamespaceLister) Get(name string) (*v1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("servicebinding"), name)
	}
	return obj.(*v1beta1.ServiceBinding), nil
}
