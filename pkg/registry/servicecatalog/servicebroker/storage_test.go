package servicebroker

import (
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

func TestNewListNilItems(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newList := NewList()
	realObj := newList.(*servicecatalog.ServiceBrokerList)
	if realObj.Items == nil {
		t.Fatalf("nil incorrectly set on Items field")
	}
}
