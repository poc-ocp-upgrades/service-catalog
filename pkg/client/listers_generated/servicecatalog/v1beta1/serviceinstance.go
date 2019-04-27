package v1beta1

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceInstanceLister interface {
	List(selector labels.Selector) (ret []*v1beta1.ServiceInstance, err error)
	ServiceInstances(namespace string) ServiceInstanceNamespaceLister
	ServiceInstanceListerExpansion
}
type serviceInstanceLister struct{ indexer cache.Indexer }

func NewServiceInstanceLister(indexer cache.Indexer) ServiceInstanceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceInstanceLister{indexer: indexer}
}
func (s *serviceInstanceLister) List(selector labels.Selector) (ret []*v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceInstance))
	})
	return ret, err
}
func (s *serviceInstanceLister) ServiceInstances(namespace string) ServiceInstanceNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return serviceInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServiceInstanceNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1beta1.ServiceInstance, err error)
	Get(name string) (*v1beta1.ServiceInstance, error)
	ServiceInstanceNamespaceListerExpansion
}
type serviceInstanceNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ServiceInstance))
	})
	return ret, err
}
func (s serviceInstanceNamespaceLister) Get(name string) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("serviceinstance"), name)
	}
	return obj.(*v1beta1.ServiceInstance), nil
}
