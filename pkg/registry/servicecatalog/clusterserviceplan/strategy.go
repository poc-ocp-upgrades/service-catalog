package clusterserviceplan

import (
	"context"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scv "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/validation"
	"k8s.io/klog"
)

func NewScopeStrategy() rest.NamespaceScopedStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clusterServicePlanRESTStrategies
}

type clusterServicePlanRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type clusterServicePlanStatusRESTStrategy struct{ clusterServicePlanRESTStrategy }

var (
	clusterServicePlanRESTStrategies								= clusterServicePlanRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_										rest.RESTCreateStrategy	= clusterServicePlanRESTStrategies
	_										rest.RESTUpdateStrategy	= clusterServicePlanRESTStrategies
	_										rest.RESTDeleteStrategy	= clusterServicePlanRESTStrategies
	clusterServicePlanStatusUpdateStrategy							= clusterServicePlanStatusRESTStrategy{clusterServicePlanRESTStrategies}
	_										rest.RESTUpdateStrategy	= clusterServicePlanStatusUpdateStrategy
)

func (clusterServicePlanRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to create")
	}
}
func (clusterServicePlanRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServicePlanRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to create")
	}
}
func (clusterServicePlanRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateClusterServicePlan(obj.(*sc.ClusterServicePlan))
}
func (clusterServicePlanRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServicePlanRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServicePlanRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServicePlan, ok := new.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to update to")
	}
	oldServicePlan, ok := old.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to update from")
	}
	newServicePlan.Spec.ClusterServiceClassRef = oldServicePlan.Spec.ClusterServiceClassRef
	newServicePlan.Spec.ClusterServiceBrokerName = oldServicePlan.Spec.ClusterServiceBrokerName
}
func (clusterServicePlanRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServicePlan, ok := new.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to validate to")
	}
	oldServicePlan, ok := old.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to validate from")
	}
	return scv.ValidateClusterServicePlanUpdate(newServicePlan, oldServicePlan)
}
func (clusterServicePlanStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceClass, ok := new.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to update to")
	}
	oldServiceClass, ok := old.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to update from")
	}
	newServiceClass.Spec = oldServiceClass.Spec
}
func (clusterServicePlanStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServicePlan, ok := new.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to validate to")
	}
	oldServicePlan, ok := old.(*sc.ClusterServicePlan)
	if !ok {
		klog.Fatal("received a non-ClusterServicePlan object to validate from")
	}
	return scv.ValidateClusterServicePlanUpdate(newServicePlan, oldServicePlan)
}
