package clusterserviceclass

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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return clusterServiceClassRESTStrategies
}

type clusterServiceClassRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type clusterServiceClassStatusRESTStrategy struct {
	clusterServiceClassRESTStrategy
}

var (
	clusterServiceClassRESTStrategies				= clusterServiceClassRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_					rest.RESTCreateStrategy	= clusterServiceClassRESTStrategies
	_					rest.RESTUpdateStrategy	= clusterServiceClassRESTStrategies
	_					rest.RESTDeleteStrategy	= clusterServiceClassRESTStrategies
	clusterServiceClassStatusUpdateStrategy				= clusterServiceClassStatusRESTStrategy{clusterServiceClassRESTStrategies}
	_					rest.RESTUpdateStrategy	= clusterServiceClassStatusUpdateStrategy
)

func (clusterServiceClassRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceclass object to create")
	}
}
func (clusterServiceClassRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServiceClassRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterServiceClass, ok := obj.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceclass object to create")
	}
	clusterServiceClass.Status = sc.ClusterServiceClassStatus{}
}
func (clusterServiceClassRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateClusterServiceClass(obj.(*sc.ClusterServiceClass))
}
func (clusterServiceClassRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServiceClassRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (clusterServiceClassRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceClass, ok := new.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceclass object to update to")
	}
	oldServiceClass, ok := old.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceclass object to update from")
	}
	newServiceClass.Status = oldServiceClass.Status
	newServiceClass.Spec.ClusterServiceBrokerName = oldServiceClass.Spec.ClusterServiceBrokerName
}
func (clusterServiceClassRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceclass, ok := new.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceclass object to validate to")
	}
	oldServiceclass, ok := old.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceclass object to validate from")
	}
	return scv.ValidateClusterServiceClassUpdate(newServiceclass, oldServiceclass)
}
func (clusterServiceClassStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceClass, ok := new.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceClass object to update to")
	}
	oldServiceClass, ok := old.(*sc.ClusterServiceClass)
	if !ok {
		klog.Fatal("received a non-clusterserviceClass object to update from")
	}
	newServiceClass.Spec = oldServiceClass.Spec
}
func (clusterServiceClassStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return field.ErrorList{}
}
