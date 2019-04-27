package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServiceBrokerLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceBroker, err error)
	ServiceBrokers(namespace string) ServiceBrokerNamespaceLister
	ServiceBrokerListerExpansion
}
type serviceBrokerLister struct{ indexer cache.Indexer }

func NewServiceBrokerLister(indexer cache.Indexer) ServiceBrokerLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &serviceBrokerLister{indexer: indexer}
}
func (s *serviceBrokerLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceBroker))
	})
	return ret, err
}
func (s *serviceBrokerLister) ServiceBrokers(namespace string) ServiceBrokerNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return serviceBrokerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServiceBrokerNamespaceLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServiceBroker, err error)
	Get(name string) (*servicecatalog.ServiceBroker, error)
	ServiceBrokerNamespaceListerExpansion
}
type serviceBrokerNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s serviceBrokerNamespaceLister) List(selector labels.Selector) (ret []*servicecatalog.ServiceBroker, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServiceBroker))
	})
	return ret, err
}
func (s serviceBrokerNamespaceLister) Get(name string) (*servicecatalog.ServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("servicebroker"), name)
	}
	return obj.(*servicecatalog.ServiceBroker), nil
}
