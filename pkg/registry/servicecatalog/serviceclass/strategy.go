package serviceclass

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
	return serviceClassRESTStrategies
}

type serviceClassRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type serviceClassStatusRESTStrategy struct{ serviceClassRESTStrategy }

var (
	serviceClassRESTStrategies									= serviceClassRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_									rest.RESTCreateStrategy	= serviceClassRESTStrategies
	_									rest.RESTUpdateStrategy	= serviceClassRESTStrategies
	_									rest.RESTDeleteStrategy	= serviceClassRESTStrategies
	serviceClassStatusUpdateStrategy							= serviceClassStatusRESTStrategy{serviceClassRESTStrategies}
	_									rest.RESTUpdateStrategy	= serviceClassStatusUpdateStrategy
)

func (serviceClassRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to create")
	}
}
func (serviceClassRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (serviceClassRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, ok := obj.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to create")
	}
	serviceClass.Status = sc.ServiceClassStatus{}
}
func (serviceClassRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateServiceClass(obj.(*sc.ServiceClass))
}
func (serviceClassRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (serviceClassRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (serviceClassRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceClass, ok := new.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to update to")
	}
	oldServiceClass, ok := old.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to update from")
	}
	newServiceClass.Status = oldServiceClass.Status
	newServiceClass.Spec.ServiceBrokerName = oldServiceClass.Spec.ServiceBrokerName
}
func (serviceClassRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceclass, ok := new.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to validate to")
	}
	oldServiceclass, ok := old.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to validate from")
	}
	return scv.ValidateServiceClassUpdate(newServiceclass, oldServiceclass)
}
func (serviceClassStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceClass, ok := new.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to update to")
	}
	oldServiceClass, ok := old.(*sc.ServiceClass)
	if !ok {
		klog.Fatal("received a non-serviceclass object to update from")
	}
	newServiceClass.Spec = oldServiceClass.Spec
}
func (serviceClassStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return field.ErrorList{}
}
