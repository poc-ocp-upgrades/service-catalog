package clusterservicebroker

import (
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

func TestNewListNilItems(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newList := NewList()
	realObj := newList.(*servicecatalog.ClusterServiceBrokerList)
	if realObj.Items == nil {
		t.Fatalf("nil incorrectly set on Items field")
	}
}
