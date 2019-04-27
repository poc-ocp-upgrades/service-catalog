package instance

import (
	"context"
	api "github.com/kubernetes-incubator/service-catalog/pkg/api"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scv "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/validation"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	"k8s.io/klog"
)

func NewScopeStrategy() rest.NamespaceScopedStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return instanceRESTStrategies
}

type instanceRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type instanceStatusRESTStrategy struct{ instanceRESTStrategy }
type instanceReferenceRESTStrategy struct{ instanceRESTStrategy }

var (
	instanceRESTStrategies						= instanceRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_				rest.RESTCreateStrategy		= instanceRESTStrategies
	_				rest.RESTUpdateStrategy		= instanceRESTStrategies
	_				rest.RESTDeleteStrategy		= instanceRESTStrategies
	_				rest.RESTGracefulDeleteStrategy	= instanceRESTStrategies
	instanceStatusUpdateStrategy					= instanceStatusRESTStrategy{instanceRESTStrategies}
	_				rest.RESTUpdateStrategy		= instanceStatusUpdateStrategy
	instanceReferenceUpdateStrategy					= instanceReferenceRESTStrategy{instanceRESTStrategies}
	_				rest.RESTUpdateStrategy		= instanceReferenceUpdateStrategy
)

func (instanceRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to create")
	}
}
func (instanceRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (instanceRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, ok := obj.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to create")
	}
	if instance.Spec.ExternalID == "" {
		instance.Spec.ExternalID = string(uuid.NewUUID())
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		setServiceInstanceUserInfo(ctx, instance)
	}
	instance.Status = sc.ServiceInstanceStatus{Conditions: []sc.ServiceInstanceCondition{}, DeprovisionStatus: sc.ServiceInstanceDeprovisionStatusNotRequired}
	instance.Spec.ClusterServiceClassRef = nil
	instance.Spec.ClusterServicePlanRef = nil
	instance.Finalizers = []string{sc.FinalizerServiceCatalog}
	instance.Generation = 1
}
func (instanceRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateServiceInstance(obj.(*sc.ServiceInstance))
}
func (instanceRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (instanceRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (instanceRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceInstance, ok := new.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to update to")
	}
	oldServiceInstance, ok := old.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to update from")
	}
	newServiceInstance.Status = oldServiceInstance.Status
	newServiceInstance.Spec.ClusterServiceClassRef = oldServiceInstance.Spec.ClusterServiceClassRef
	newServiceInstance.Spec.ClusterServicePlanRef = oldServiceInstance.Spec.ClusterServicePlanRef
	planUpdated := newServiceInstance.Spec.ClusterServicePlanExternalName != oldServiceInstance.Spec.ClusterServicePlanExternalName || newServiceInstance.Spec.ClusterServicePlanExternalID != oldServiceInstance.Spec.ClusterServicePlanExternalID || newServiceInstance.Spec.ClusterServicePlanName != oldServiceInstance.Spec.ClusterServicePlanName
	if planUpdated {
		newServiceInstance.Spec.ClusterServicePlanRef = nil
	}
	if newServiceInstance.Spec.UpdateRequests == 0 {
		newServiceInstance.Spec.UpdateRequests = oldServiceInstance.Spec.UpdateRequests
	}
	if !apiequality.Semantic.DeepEqual(oldServiceInstance.Spec, newServiceInstance.Spec) {
		if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
			setServiceInstanceUserInfo(ctx, newServiceInstance)
		}
		newServiceInstance.Generation = oldServiceInstance.Generation + 1
	}
}
func (instanceRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceInstance, ok := new.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to validate to")
	}
	oldServiceInstance, ok := old.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to validate from")
	}
	return scv.ValidateServiceInstanceUpdate(newServiceInstance, oldServiceInstance)
}
func (instanceRESTStrategy) CheckGracefulDelete(ctx context.Context, obj runtime.Object, options *metav1.DeleteOptions) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		serviceInstance, ok := obj.(*sc.ServiceInstance)
		if !ok {
			klog.Fatal("received a non-instance object to delete")
		}
		setServiceInstanceUserInfo(ctx, serviceInstance)
	}
	return false
}
func (instanceStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceInstance, ok := new.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to update to")
	}
	oldServiceInstance, ok := old.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to update from")
	}
	newServiceInstance.Spec = oldServiceInstance.Spec
}
func (instanceStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceInstance, ok := new.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to validate to")
	}
	oldServiceInstance, ok := old.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to validate from")
	}
	return scv.ValidateServiceInstanceStatusUpdate(newServiceInstance, oldServiceInstance)
}
func (instanceReferenceRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceInstance, ok := new.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to update to")
	}
	oldServiceInstance, ok := old.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to update from")
	}
	newClusterServiceClassRef := newServiceInstance.Spec.ClusterServiceClassRef
	newClusterServicePlanRef := newServiceInstance.Spec.ClusterServicePlanRef
	newServiceClassRef := newServiceInstance.Spec.ServiceClassRef
	newServicePlanRef := newServiceInstance.Spec.ServicePlanRef
	newServiceInstance.Spec = oldServiceInstance.Spec
	newServiceInstance.Spec.ClusterServiceClassRef = newClusterServiceClassRef
	newServiceInstance.Spec.ClusterServicePlanRef = newClusterServicePlanRef
	newServiceInstance.Spec.ServiceClassRef = newServiceClassRef
	newServiceInstance.Spec.ServicePlanRef = newServicePlanRef
	newServiceInstance.Status = oldServiceInstance.Status
}
func (instanceReferenceRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceInstance, ok := new.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to validate to")
	}
	oldServiceInstance, ok := old.(*sc.ServiceInstance)
	if !ok {
		klog.Fatal("received a non-instance object to validate from")
	}
	return scv.ValidateServiceInstanceReferencesUpdate(newServiceInstance, oldServiceInstance)
}
func setServiceInstanceUserInfo(ctx context.Context, instance *sc.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance.Spec.UserInfo = nil
	if user, ok := genericapirequest.UserFrom(ctx); ok {
		instance.Spec.UserInfo = &sc.UserInfo{Username: user.GetName(), UID: user.GetUID(), Groups: user.GetGroups()}
		if extra := user.GetExtra(); len(extra) > 0 {
			instance.Spec.UserInfo.Extra = map[string]sc.ExtraValue{}
			for k, v := range extra {
				instance.Spec.UserInfo.Extra[k] = sc.ExtraValue(v)
			}
		}
	}
}
