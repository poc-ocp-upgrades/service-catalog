package clusterserviceclass

import (
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

func TestNewList(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newList := NewList()
	realObj := newList.(*servicecatalog.ClusterServiceClassList)
	if realObj.Items == nil {
		t.Fatalf("nil incorrectly set on Items field")
	}
}
