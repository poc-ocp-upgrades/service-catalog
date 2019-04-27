package server

import (
	"k8s.io/apiserver/pkg/admission"
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/broker/authsarcheck"
	siclifecycle "github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/servicebindings/lifecycle"
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/serviceplan/changevalidator"
	"github.com/kubernetes-incubator/service-catalog/plugin/pkg/admission/serviceplan/defaultserviceplan"
)

func registerAllAdmissionPlugins(plugins *admission.Plugins) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultserviceplan.Register(plugins)
	siclifecycle.Register(plugins)
	changevalidator.Register(plugins)
	authsarcheck.Register(plugins)
}
