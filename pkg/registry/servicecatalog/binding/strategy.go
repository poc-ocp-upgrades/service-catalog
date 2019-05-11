package binding

import (
	"context"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
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
	return bindingRESTStrategies
}

type bindingRESTStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}
type bindingStatusRESTStrategy struct{ bindingRESTStrategy }

var (
	bindingRESTStrategies										= bindingRESTStrategy{ObjectTyper: api.Scheme, NameGenerator: names.SimpleNameGenerator}
	_							rest.RESTCreateStrategy			= bindingRESTStrategies
	_							rest.RESTUpdateStrategy			= bindingRESTStrategies
	_							rest.RESTDeleteStrategy			= bindingRESTStrategies
	_							rest.RESTGracefulDeleteStrategy	= bindingRESTStrategies
	bindingStatusUpdateStrategy									= bindingStatusRESTStrategy{bindingRESTStrategies}
	_							rest.RESTUpdateStrategy			= bindingStatusUpdateStrategy
)

func (bindingRESTStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to create")
	}
}
func (bindingRESTStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (bindingRESTStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	binding, ok := obj.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to create")
	}
	if binding.Spec.ExternalID == "" {
		binding.Spec.ExternalID = string(uuid.NewUUID())
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		setServiceBindingUserInfo(ctx, binding)
	}
	binding.Status = sc.ServiceBindingStatus{UnbindStatus: sc.ServiceBindingUnbindStatusNotRequired}
	binding.Status.Conditions = []sc.ServiceBindingCondition{}
	binding.Finalizers = []string{sc.FinalizerServiceCatalog}
	binding.Generation = 1
}
func (bindingRESTStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return scv.ValidateServiceBinding(obj.(*sc.ServiceBinding))
}
func (bindingRESTStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (bindingRESTStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (bindingRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBinding, ok := new.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to update to")
	}
	oldServiceBinding, ok := old.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to update from")
	}
	newServiceBinding.Status = oldServiceBinding.Status
	newServiceBinding.Spec = oldServiceBinding.Spec
	if !apiequality.Semantic.DeepEqual(oldServiceBinding.Spec, newServiceBinding.Spec) {
		if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
			setServiceBindingUserInfo(ctx, newServiceBinding)
		}
		newServiceBinding.Generation = oldServiceBinding.Generation + 1
	}
}
func (bindingRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBinding, ok := new.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to validate to")
	}
	oldServiceBinding, ok := old.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to validate from")
	}
	return scv.ValidateServiceBindingUpdate(newServiceBinding, oldServiceBinding)
}
func (bindingRESTStrategy) CheckGracefulDelete(ctx context.Context, obj runtime.Object, options *metav1.DeleteOptions) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		serviceInstanceCredential, ok := obj.(*sc.ServiceBinding)
		if !ok {
			klog.Fatal("received a non-ServiceBinding object to delete")
		}
		setServiceBindingUserInfo(ctx, serviceInstanceCredential)
	}
	return false
}
func (bindingStatusRESTStrategy) PrepareForUpdate(ctx context.Context, new, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBinding, ok := new.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to update to")
	}
	oldServiceBinding, ok := old.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to update from")
	}
	newServiceBinding.Spec = oldServiceBinding.Spec
}
func (bindingStatusRESTStrategy) ValidateUpdate(ctx context.Context, new, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newServiceBinding, ok := new.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to validate to")
	}
	oldServiceBinding, ok := old.(*sc.ServiceBinding)
	if !ok {
		klog.Fatal("received a non-binding object to validate from")
	}
	return scv.ValidateServiceBindingStatusUpdate(newServiceBinding, oldServiceBinding)
}
func setServiceBindingUserInfo(ctx context.Context, instanceCredential *sc.ServiceBinding) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceCredential.Spec.UserInfo = nil
	if user, ok := genericapirequest.UserFrom(ctx); ok {
		instanceCredential.Spec.UserInfo = &sc.UserInfo{Username: user.GetName(), UID: user.GetUID(), Groups: user.GetGroups()}
		if extra := user.GetExtra(); len(extra) > 0 {
			instanceCredential.Spec.UserInfo.Extra = map[string]sc.ExtraValue{}
			for k, v := range extra {
				instanceCredential.Spec.UserInfo.Extra[k] = sc.ExtraValue(v)
			}
		}
	}
}
