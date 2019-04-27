package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ServicePlanLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error)
	ServicePlans(namespace string) ServicePlanNamespaceLister
	ServicePlanListerExpansion
}
type servicePlanLister struct{ indexer cache.Indexer }

func NewServicePlanLister(indexer cache.Indexer) ServicePlanLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicePlanLister{indexer: indexer}
}
func (s *servicePlanLister) List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServicePlan))
	})
	return ret, err
}
func (s *servicePlanLister) ServicePlans(namespace string) ServicePlanNamespaceLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return servicePlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type ServicePlanNamespaceLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error)
	Get(name string) (*servicecatalog.ServicePlan, error)
	ServicePlanNamespaceListerExpansion
}
type servicePlanNamespaceLister struct {
	indexer		cache.Indexer
	namespace	string
}

func (s servicePlanNamespaceLister) List(selector labels.Selector) (ret []*servicecatalog.ServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ServicePlan))
	})
	return ret, err
}
func (s servicePlanNamespaceLister) Get(name string) (*servicecatalog.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("serviceplan"), name)
	}
	return obj.(*servicecatalog.ServicePlan), nil
}
