package integration

import (
	"time"
	"k8s.io/apiserver/pkg/util/feature"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
)

var (
	pollInterval	= 2 * time.Second
	defaultTimeout	= 30 * time.Second
)

func strPtr(s string) *string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &s
}
func truePtr() *bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := true
	return &b
}
func falsePtr() *bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := false
	return &b
}
func enableNamespacedResources() (resetFeaturesFunc func(), err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	previousFeatureGate := feature.DefaultFeatureGate
	newFeatureGate := feature.NewFeatureGate()
	if err := newFeatureGate.Add(map[feature.Feature]feature.FeatureSpec{scfeatures.NamespacedServiceBroker: {Default: true, PreRelease: feature.Alpha}}); err != nil {
		return nil, err
	}
	feature.DefaultFeatureGate = newFeatureGate
	return func() {
		feature.DefaultFeatureGate = previousFeatureGate
	}, nil
}
