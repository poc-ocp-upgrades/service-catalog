package serviceplan

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
	return servicePlanRESTStrategies
}

type servicePlanRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type servicePlanStatusRESTStrategy struct{ servicePlanRESTStrategy }

var (
	servicePlanRESTStrategies								= servicePlanRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_								rest.RESTCreateStrategy	= servicePlanRESTStrategies
	_								rest.RESTUpdateStrategy	= servicePlanRESTStrategies
	_								rest.RESTDeleteStrategy	= servicePlanRESTStrategies
	servicePlanStatusUpdateStrategy							= servicePlanStatusRESTStrategy{servicePlanRESTStrategies}
	_								rest.RESTUpdateStrategy	= servicePlanStatusUpdateStrategy
)

func (servicePlanRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to create")
	}
}
func (servicePlanRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (servicePlanRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to create")
	}
}
func (servicePlanRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateServicePlan(obj.(*sc.ServicePlan))
}
func (servicePlanRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (servicePlanRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (servicePlanRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServicePlan, ok := new.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to update to")
	}
	oldServicePlan, ok := old.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to update from")
	}
	newServicePlan.Spec.ServiceClassRef = oldServicePlan.Spec.ServiceClassRef
	newServicePlan.Spec.ServiceBrokerName = oldServicePlan.Spec.ServiceBrokerName
}
func (servicePlanRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServicePlan, ok := new.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to validate to")
	}
	oldServicePlan, ok := old.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to validate from")
	}
	return scv.ValidateServicePlanUpdate(newServicePlan, oldServicePlan)
}
func (servicePlanStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceClass, ok := new.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to update to")
	}
	oldServiceClass, ok := old.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to update from")
	}
	newServiceClass.Spec = oldServiceClass.Spec
}
func (servicePlanStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServicePlan, ok := new.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to validate to")
	}
	oldServicePlan, ok := old.(*sc.ServicePlan)
	if !ok {
		klog.Fatal("received a non-ServicePlan object to validate from")
	}
	return scv.ValidateServicePlanUpdate(newServicePlan, oldServicePlan)
}
