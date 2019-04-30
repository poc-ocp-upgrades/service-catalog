package v1beta1

import (
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

type ClusterServiceClassLister interface {
	List(selector labels.Selector) (ret []*v1beta1.ClusterServiceClass, err error)
	Get(name string) (*v1beta1.ClusterServiceClass, error)
	ClusterServiceClassListerExpansion
}
type clusterServiceClassLister struct{ indexer cache.Indexer }

func NewClusterServiceClassLister(indexer cache.Indexer) ClusterServiceClassLister {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &clusterServiceClassLister{indexer: indexer}
}
func (s *clusterServiceClassLister) List(selector labels.Selector) (ret []*v1beta1.ClusterServiceClass, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ClusterServiceClass))
	})
	return ret, err
}
func (s *clusterServiceClassLister) Get(name string) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("clusterserviceclass"), name)
	}
	return obj.(*v1beta1.ClusterServiceClass), nil
}
