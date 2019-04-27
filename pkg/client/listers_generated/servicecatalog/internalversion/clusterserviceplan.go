package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ClusterServicePlanLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ClusterServicePlan, err error)
	Get(name string) (*servicecatalog.ClusterServicePlan, error)
	ClusterServicePlanListerExpansion
}
type clusterServicePlanLister struct{ indexer cache.Indexer }

func NewClusterServicePlanLister(indexer cache.Indexer) ClusterServicePlanLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServicePlanLister{indexer: indexer}
}
func (s *clusterServicePlanLister) List(selector labels.Selector) (ret []*servicecatalog.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ClusterServicePlan))
	})
	return ret, err
}
func (s *clusterServicePlanLister) Get(name string) (*servicecatalog.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("clusterserviceplan"), name)
	}
	return obj.(*servicecatalog.ClusterServicePlan), nil
}
