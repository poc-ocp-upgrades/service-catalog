package controller

import (
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"testing"
)

var _ controller.Controller = &testController{}

func TestController(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
