package podpreset

import (
	"context"
	"fmt"
	api "github.com/kubernetes-incubator/service-catalog/pkg/api"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	apistorage "k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/validation"
)

func NewScopeStrategy() rest.NamespaceScopedStrategy {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return podPresetRESTStrategy
}

type podPresetStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

var (
	podPresetRESTStrategy				= podPresetStrategy{api.Scheme, names.SimpleNameGenerator}
	_			rest.RESTCreateStrategy	= podPresetRESTStrategy
	_			rest.RESTUpdateStrategy	= podPresetRESTStrategy
	_			rest.RESTDeleteStrategy	= podPresetRESTStrategy
)

func (podPresetStrategy) NamespaceScoped() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func (podPresetStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pip := obj.(*settings.PodPreset)
	pip.Generation = 1
}
func (podPresetStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	newPodPreset := obj.(*settings.PodPreset)
	oldPodPreset := old.(*settings.PodPreset)
	newPodPreset.Spec = oldPodPreset.Spec
}
func (podPresetStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pip := obj.(*settings.PodPreset)
	return validation.ValidatePodPreset(pip)
}
func (podPresetStrategy) Canonicalize(obj runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func (podPresetStrategy) AllowCreateOnUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (podPresetStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	validationErrorList := validation.ValidatePodPreset(obj.(*settings.PodPreset))
	updateErrorList := validation.ValidatePodPresetUpdate(obj.(*settings.PodPreset), old.(*settings.PodPreset))
	return append(validationErrorList, updateErrorList...)
}
func (podPresetStrategy) AllowUnconditionalUpdate() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return true
}
func SelectableFields(pip *settings.PodPreset) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return generic.ObjectMetaFieldsSet(&pip.ObjectMeta, true)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pip, ok := obj.(*settings.PodPreset)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a podpreset")
	}
	return labels.Set(pip.ObjectMeta.Labels), SelectableFields(pip), pip.Initializers != nil, nil
}
func Matcher(label labels.Selector, field fields.Selector) apistorage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apistorage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
