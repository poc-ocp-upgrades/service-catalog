package instance

import (
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

func TestNewListNilField(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newList := NewList()
	realObj := newList.(*servicecatalog.ServiceInstanceList)
	if realObj.Items == nil {
		t.Fatalf("nil incorrectly set on Items field")
	}
}
