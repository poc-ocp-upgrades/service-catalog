package servicecatalog_test

import (
	"math/rand"
	"reflect"
	"testing"
	"github.com/google/gofuzz"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testapi"
	sctesting "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testing"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/uuid"
)

func doUnstructuredRoundTrip(t *testing.T, group testapi.TestGroup, kind string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	internalObj, err := api.Scheme.New(group.InternalGroupVersion().WithKind(kind))
	if err != nil {
		t.Fatalf("Couldn't create internal object %v: %v", kind, err)
	}
	seed := rand.Int63()
	fuzzer.FuzzerFor(sctesting.FuzzerFuncs, rand.NewSource(seed), api.Codecs).Funcs(func(is *servicecatalog.ServiceInstanceSpec, c fuzz.Continue) {
		c.FuzzNoCustom(is)
		is.ExternalID = string(uuid.NewUUID())
		is.Parameters = nil
	}, func(is *servicecatalog.ServiceInstanceStatus, c fuzz.Continue) {
		c.FuzzNoCustom(is)
		is.DefaultProvisionParameters = nil
	}, func(is *servicecatalog.CommonServicePlanSpec, c fuzz.Continue) {
		c.FuzzNoCustom(is)
		is.DefaultProvisionParameters = nil
		is.ExternalMetadata = nil
		is.ServiceBindingCreateParameterSchema = nil
		is.ServiceBindingCreateResponseSchema = nil
		is.InstanceCreateParameterSchema = nil
		is.InstanceUpdateParameterSchema = nil
	}, func(cs *servicecatalog.CommonServiceClassSpec, c fuzz.Continue) {
		c.FuzzNoCustom(cs)
		cs.DefaultProvisionParameters = nil
		cs.ExternalMetadata = nil
	}, func(bs *servicecatalog.ServiceBindingSpec, c fuzz.Continue) {
		c.FuzzNoCustom(bs)
		bs.ExternalID = string(uuid.NewUUID())
		for bs.SecretName == "" {
			bs.SecretName = c.RandString()
		}
		bs.Parameters = nil
	}, func(ps *servicecatalog.ServiceInstancePropertiesState, c fuzz.Continue) {
		c.FuzzNoCustom(ps)
		ps.Parameters = nil
	}, func(ps *servicecatalog.ServiceBindingPropertiesState, c fuzz.Continue) {
		c.FuzzNoCustom(ps)
		ps.Parameters = nil
	}).Fuzz(internalObj)
	item, err := api.Scheme.New(group.GroupVersion().WithKind(kind))
	if err != nil {
		t.Fatalf("Couldn't create external object %v: %v", kind, err)
	}
	if err := api.Scheme.Convert(internalObj, item, nil); err != nil {
		t.Fatalf("Conversion for %v failed: %v", kind, err)
	}
	data, err := json.Marshal(item)
	if err != nil {
		t.Errorf("Error when marshaling object: %v", err)
		return
	}
	unstr := make(map[string]interface{})
	err = json.Unmarshal(data, &unstr)
	if err != nil {
		t.Errorf("Error when unmarshaling to unstructured: %v", err)
		return
	}
	data, err = json.Marshal(unstr)
	if err != nil {
		t.Errorf("Error when marshaling unstructured: %v", err)
		return
	}
	unmarshalledObj := reflect.New(reflect.TypeOf(item).Elem()).Interface()
	err = json.Unmarshal(data, &unmarshalledObj)
	if err != nil {
		t.Errorf("Error when unmarshaling to object: %v", err)
		return
	}
	if !apiequality.Semantic.DeepEqual(item, unmarshalledObj) {
		t.Errorf("Object changed during JSON operations, diff: %v", diff.ObjectReflectDiff(item, unmarshalledObj))
		return
	}
	newUnstr, err := runtime.DefaultUnstructuredConverter.ToUnstructured(item)
	if err != nil {
		t.Errorf("ToUnstructured failed: %v", err)
		return
	}
	newObj := reflect.New(reflect.TypeOf(item).Elem()).Interface().(runtime.Object)
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(newUnstr, newObj)
	if err != nil {
		t.Errorf("FromUnstructured failed: %v", err)
		return
	}
	if !apiequality.Semantic.DeepEqual(item, newObj) {
		t.Errorf("%v: Object changed, diff: %v", kind, diff.ObjectReflectDiff(item, newObj))
	}
}
func TestRoundTripTypesToUnstructured(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for groupKey, group := range catalogGroups {
		for kind := range group.InternalTypes() {
			t.Logf("Testing: %v in %v", kind, groupKey)
			for i := 0; i < 50; i++ {
				doUnstructuredRoundTrip(t, group, kind)
				if t.Failed() {
					break
				}
			}
		}
	}
}
