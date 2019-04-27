package controller

import (
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/controller"
	"testing"
)

var _ controller.Controller = &userProvidedController{}

func TestController(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
}
