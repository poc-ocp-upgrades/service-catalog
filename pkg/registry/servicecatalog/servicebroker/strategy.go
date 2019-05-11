package servicebroker

import (
	"context"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
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
	return serviceBrokerRESTStrategies
}

type serviceBrokerRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type serviceBrokerStatusRESTStrategy struct{ serviceBrokerRESTStrategy }

var (
	serviceBrokerRESTStrategies									= serviceBrokerRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_									rest.RESTCreateStrategy	= serviceBrokerRESTStrategies
	_									rest.RESTUpdateStrategy	= serviceBrokerRESTStrategies
	_									rest.RESTDeleteStrategy	= serviceBrokerRESTStrategies
	serviceBrokerStatusUpdateStrategy							= serviceBrokerStatusRESTStrategy{serviceBrokerRESTStrategies}
	_									rest.RESTUpdateStrategy	= serviceBrokerStatusUpdateStrategy
)

func (serviceBrokerRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to create")
	}
}
func (serviceBrokerRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (serviceBrokerRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, ok := obj.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to create")
	}
	broker.Status = sc.ServiceBrokerStatus{}
	broker.Status.Conditions = []sc.ServiceBrokerCondition{}
	broker.Finalizers = []string{sc.FinalizerServiceCatalog}
	broker.Generation = 1
}
func (serviceBrokerRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateServiceBroker(obj.(*sc.ServiceBroker))
}
func (serviceBrokerRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (serviceBrokerRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (serviceBrokerRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBroker, ok := new.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to update to")
	}
	oldServiceBroker, ok := old.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to update from")
	}
	newServiceBroker.Status = oldServiceBroker.Status
	if newServiceBroker.Spec.RelistRequests == 0 {
		newServiceBroker.Spec.RelistRequests = oldServiceBroker.Spec.RelistRequests
	}
	if !apiequality.Semantic.DeepEqual(oldServiceBroker.Spec, newServiceBroker.Spec) {
		newServiceBroker.Generation = oldServiceBroker.Generation + 1
	}
}
func (serviceBrokerRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBroker, ok := new.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to validate to")
	}
	oldServiceBroker, ok := old.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to validate from")
	}
	return scv.ValidateServiceBrokerUpdate(newServiceBroker, oldServiceBroker)
}
func (serviceBrokerStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBroker, ok := new.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to update to")
	}
	oldServiceBroker, ok := old.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to update from")
	}
	newServiceBroker.Spec = oldServiceBroker.Spec
}
func (serviceBrokerStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBroker, ok := new.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to validate to")
	}
	oldServiceBroker, ok := old.(*sc.ServiceBroker)
	if !ok {
		klog.Fatal("received a non-servicebroker object to validate from")
	}
	return scv.ValidateServiceBrokerStatusUpdate(newServiceBroker, oldServiceBroker)
}
