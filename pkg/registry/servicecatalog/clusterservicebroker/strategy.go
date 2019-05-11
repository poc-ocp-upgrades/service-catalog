package clusterservicebroker

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
	return clusterServiceBrokerRESTStrategies
}

type clusterServiceBrokerRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type clusterServiceBrokerStatusRESTStrategy struct {
	clusterServiceBrokerRESTStrategy
}

var (
	clusterServiceBrokerRESTStrategies									= clusterServiceBrokerRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_											rest.RESTCreateStrategy	= clusterServiceBrokerRESTStrategies
	_											rest.RESTUpdateStrategy	= clusterServiceBrokerRESTStrategies
	_											rest.RESTDeleteStrategy	= clusterServiceBrokerRESTStrategies
	clusterServiceBrokerStatusUpdateStrategy							= clusterServiceBrokerStatusRESTStrategy{clusterServiceBrokerRESTStrategies}
	_											rest.RESTUpdateStrategy	= clusterServiceBrokerStatusUpdateStrategy
)

func (clusterServiceBrokerRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to create")
	}
}
func (clusterServiceBrokerRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServiceBrokerRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, ok := obj.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to create")
	}
	broker.Status = sc.ClusterServiceBrokerStatus{}
	broker.Status.Conditions = []sc.ServiceBrokerCondition{}
	broker.Finalizers = []string{sc.FinalizerServiceCatalog}
	broker.Generation = 1
}
func (clusterServiceBrokerRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateClusterServiceBroker(obj.(*sc.ClusterServiceBroker))
}
func (clusterServiceBrokerRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServiceBrokerRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServiceBrokerRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newClusterServiceBroker, ok := new.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to update to")
	}
	oldClusterServiceBroker, ok := old.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to update from")
	}
	newClusterServiceBroker.Status = oldClusterServiceBroker.Status
	if newClusterServiceBroker.Spec.RelistRequests == 0 {
		newClusterServiceBroker.Spec.RelistRequests = oldClusterServiceBroker.Spec.RelistRequests
	}
	if !apiequality.Semantic.DeepEqual(oldClusterServiceBroker.Spec, newClusterServiceBroker.Spec) {
		newClusterServiceBroker.Generation = oldClusterServiceBroker.Generation + 1
	}
}
func (clusterServiceBrokerRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newClusterServiceBroker, ok := new.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to validate to")
	}
	oldClusterServiceBroker, ok := old.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to validate from")
	}
	return scv.ValidateClusterServiceBrokerUpdate(newClusterServiceBroker, oldClusterServiceBroker)
}
func (clusterServiceBrokerStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newClusterServiceBroker, ok := new.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to update to")
	}
	oldClusterServiceBroker, ok := old.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to update from")
	}
	newClusterServiceBroker.Spec = oldClusterServiceBroker.Spec
}
func (clusterServiceBrokerStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newClusterServiceBroker, ok := new.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to validate to")
	}
	oldClusterServiceBroker, ok := old.(*sc.ClusterServiceBroker)
	if !ok {
		klog.Fatal("received a non-clusterservicebroker object to validate from")
	}
	return scv.ValidateClusterServiceBrokerStatusUpdate(newClusterServiceBroker, oldClusterServiceBroker)
}
