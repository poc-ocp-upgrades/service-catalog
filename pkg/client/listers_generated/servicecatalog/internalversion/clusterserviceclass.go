package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ClusterServiceClassLister interface {
	List(selector labels.Selector) (ret []*servicecatalog.ClusterServiceClass, err error)
	Get(name string) (*servicecatalog.ClusterServiceClass, error)
	ClusterServiceClassListerExpansion
}
type clusterServiceClassLister struct{ indexer cache.Indexer }

func NewClusterServiceClassLister(indexer cache.Indexer) ClusterServiceClassLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceClassLister{indexer: indexer}
}
func (s *clusterServiceClassLister) List(selector labels.Selector) (ret []*servicecatalog.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ClusterServiceClass))
	})
	return ret, err
}
func (s *clusterServiceClassLister) Get(name string) (*servicecatalog.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("clusterserviceclass"), name)
	}
	return obj.(*servicecatalog.ClusterServiceClass), nil
}
