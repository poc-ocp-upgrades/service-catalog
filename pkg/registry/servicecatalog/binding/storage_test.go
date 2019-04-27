package binding

import (
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

func TestNewListNilItems(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newList := NewList()
	realObj := newList.(*servicecatalog.ServiceBindingList)
	if realObj.Items == nil {
		t.Fatalf("nil incorrectly set on Items field")
	}
}
