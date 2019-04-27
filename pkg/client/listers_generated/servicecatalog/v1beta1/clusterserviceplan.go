package v1beta1

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ClusterServicePlanLister interface {
	List(selector labels.Selector) (ret []*v1beta1.ClusterServicePlan, err error)
	Get(name string) (*v1beta1.ClusterServicePlan, error)
	ClusterServicePlanListerExpansion
}
type clusterServicePlanLister struct{ indexer cache.Indexer }

func NewClusterServicePlanLister(indexer cache.Indexer) ClusterServicePlanLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServicePlanLister{indexer: indexer}
}
func (s *clusterServicePlanLister) List(selector labels.Selector) (ret []*v1beta1.ClusterServicePlan, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ClusterServicePlan))
	})
	return ret, err
}
func (s *clusterServicePlanLister) Get(name string) (*v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("clusterserviceplan"), name)
	}
	return obj.(*v1beta1.ClusterServicePlan), nil
}
