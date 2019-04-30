package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceInstanceLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceInstance, err error)
	ServiceInstances(namespace string) ServiceInstanceNamespaceLister
	ServiceInstanceListerExpansion
}
type serviceInstanceLister struct{ indexer cache.Indexer }

func NewServiceInstanceLister(indexer cache.Indexer) ServiceInstanceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceInstanceLister{indexer: indexer}
}
func (s *serviceInstanceLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceInstance))
	})
	return ret, err
}
func (s *serviceInstanceLister) ServiceInstances(namespace string) ServiceInstanceNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return serviceInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServiceInstanceNamespaceLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceInstance, err error)
	Get(name string) (*servicecatalog.ServiceInstance, error)
	ServiceInstanceNamespaceListerExpansion
}
type serviceInstanceNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceInstanceNamespaceLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceInstance))
	})
	return ret, err
}
func (s serviceInstanceNamespaceLister) Get(name string) (*servicecatalog.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("serviceinstance"), name)
	}
	return obj.(*servicecatalog.ServiceInstance), nil
}
