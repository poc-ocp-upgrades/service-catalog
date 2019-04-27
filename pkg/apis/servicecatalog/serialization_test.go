package servicecatalog_test

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"github.com/davecgh/go-spew/spew"
	proto "github.com/golang/protobuf/proto"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	testapi "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testapi"
	apitesting "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testing"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/apimachinery/pkg/util/sets"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	"k8s.io/apimachinery/pkg/api/apitesting/roundtrip"
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testapi.Groups[servicecatalog.GroupName] = serviceCatalogAPIGroup()
}

var codecsToTest = []func(version schema.GroupVersion, item runtime.Object) (runtime.Codec, bool, error){func(version schema.GroupVersion, item runtime.Object) (runtime.Codec, bool, error) {
	c, err := testapi.GetCodecForObject(item)
	return c, true, err
}}

func fuzzInternalObject(t *testing.T, forVersion schema.GroupVersion, item runtime.Object, seed int64) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fuzzer.FuzzerFor(apitesting.FuzzerFuncs, rand.NewSource(seed), api.Codecs).Fuzz(item)
	j, err := meta.TypeAccessor(item)
	if err != nil {
		t.Fatalf("Unexpected error %v for %#v", err, item)
	}
	j.SetKind("")
	j.SetAPIVersion("")
	return item
}
func dataAsString(data []byte) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	dataString := string(data)
	if !strings.HasPrefix(dataString, "{") {
		dataString = "\n" + hex.Dump(data)
		proto.NewBuffer(make([]byte, 0, 1024)).DebugPrint("decoded object", data)
	}
	return dataString
}
func roundTripSame(t *testing.T, group testapi.TestGroup, item runtime.Object, except ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	set := sets.NewString(except...)
	seed := rand.Int63()
	fuzzInternalObject(t, group.InternalGroupVersion(), item, seed)
	version := *group.GroupVersion()
	codecs := []runtime.Codec{}
	for _, fn := range codecsToTest {
		codec, ok, err := fn(version, item)
		if err != nil {
			t.Errorf("unable to get codec: %v", err)
			return
		}
		if !ok {
			continue
		}
		codecs = append(codecs, codec)
	}
	t.Logf("version: %v\n", version)
	t.Logf("codec: %#v\n", codecs[0])
	if !set.Has(version.String()) {
		fuzzInternalObject(t, version, item, seed)
		for _, codec := range codecs {
			roundTrip(t, codec, item)
		}
	}
}
func roundTrip(t *testing.T, codec runtime.Codec, item runtime.Object) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	printer := spew.ConfigState{DisableMethods: true}
	original := item
	item = item.DeepCopyObject()
	name := reflect.TypeOf(item).Elem().Name()
	data, err := runtime.Encode(codec, item)
	if err != nil {
		if runtime.IsNotRegisteredError(err) {
			t.Logf("%v: not registered: %v (%s)", name, err, printer.Sprintf("%#v", item))
		} else {
			t.Errorf("%v: %v (%s)", name, err, printer.Sprintf("%#v", item))
		}
		return
	}
	if !equality.Semantic.DeepEqual(original, item) {
		t.Errorf("0: %v: encode altered the object, diff: %v", name, diff.ObjectReflectDiff(original, item))
		return
	}
	t.Logf("Codec: %+v\n", codec)
	obj2, err := runtime.Decode(codec, data)
	if err != nil {
		t.Errorf("ERROR: %v", err)
		t.Errorf("0: %v: %v\nCodec: %#v\nData: %s\nSource: %#v", name, err, codec, dataAsString(data), printer.Sprintf("%#v", item))
		panic("failed")
	}
	if !equality.Semantic.DeepEqual(original, obj2) {
		t.Errorf("\n1: %v: diff: %v\nCodec: %#v\nSource:\n\n%#v\n\nEncoded:\n\n%s\n\nFinal:\n\n%#v", name, diff.ObjectReflectDiff(item, obj2), codec, printer.Sprintf("%#v", item), dataAsString(data), printer.Sprintf("%#v", obj2))
		return
	}
	obj3 := reflect.New(reflect.TypeOf(item).Elem()).Interface().(runtime.Object)
	if err := runtime.DecodeInto(codec, data, obj3); err != nil {
		t.Errorf("2: %v: %v", name, err)
		return
	}
	if !equality.Semantic.DeepEqual(item, obj3) {
		t.Errorf("3: %v: diff: %v\nCodec: %#v", name, diff.ObjectReflectDiff(item, obj3), codec)
		return
	}
}
func serviceCatalogAPIGroup() testapi.TestGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupVersion, err := schema.ParseGroupVersion("servicecatalog.k8s.io/v1beta1")
	if err != nil {
		panic(fmt.Sprintf("Error parsing groupversion: %v", err))
	}
	externalGroupVersion := schema.GroupVersion{Group: servicecatalog.GroupName, Version: api.Scheme.PrioritizedVersionsForGroup(servicecatalog.GroupName)[0].Version}
	return testapi.NewTestGroup(groupVersion, servicecatalog.SchemeGroupVersion, api.Scheme.KnownTypes(servicecatalog.SchemeGroupVersion), api.Scheme.KnownTypes(externalGroupVersion))
}
func TestSpecificKind(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	group := serviceCatalogAPIGroup()
	for _, kind := range group.InternalTypes() {
		t.Log(kind)
	}
	internalGVK := schema.GroupVersionKind{Group: group.GroupVersion().Group, Version: group.GroupVersion().Version, Kind: "ClusterServiceClass"}
	seed := rand.Int63()
	fuzzer := fuzzer.FuzzerFor(apitesting.FuzzerFuncs, rand.NewSource(seed), api.Codecs)
	roundtrip.RoundTripSpecificKindWithoutProtobuf(t, internalGVK, api.Scheme, api.Codecs, fuzzer, nil)
}
func TestClusterServiceBrokerList(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	kind := "ClusterServiceBrokerList"
	item, err := api.Scheme.New(serviceCatalogAPIGroup().InternalGroupVersion().WithKind(kind))
	if err != nil {
		t.Errorf("Couldn't make a %v? %v", kind, err)
		return
	}
	roundTripSame(t, serviceCatalogAPIGroup(), item)
}

var catalogGroups = map[string]testapi.TestGroup{"servicecatalog": serviceCatalogAPIGroup()}

func TestRoundTripTypes(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	seed := rand.Int63()
	fuzzer := fuzzer.FuzzerFor(apitesting.FuzzerFuncs, rand.NewSource(seed), api.Codecs)
	nonRoundTrippableTypes := map[schema.GroupVersionKind]bool{}
	roundtrip.RoundTripTypesWithoutProtobuf(t, api.Scheme, api.Codecs, fuzzer, nonRoundTrippableTypes)
}
func TestBadJSONRejection(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	badJSONMissingKind := []byte(`{ }`)
	if _, err := runtime.Decode(testapi.ServiceCatalog.Codec(), badJSONMissingKind); err == nil {
		t.Errorf("Did not reject despite lack of kind field: %s", badJSONMissingKind)
	}
	badJSONUnknownType := []byte(`{"kind": "bar"}`)
	if _, err1 := runtime.Decode(testapi.ServiceCatalog.Codec(), badJSONUnknownType); err1 == nil {
		t.Errorf("Did not reject despite use of unknown type: %s", badJSONUnknownType)
	}
}
