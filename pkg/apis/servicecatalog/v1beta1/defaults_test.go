package v1beta1_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/testapi"
	versioned "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	groupVersion, err := schema.ParseGroupVersion("servicecatalog.k8s.io/v1beta1")
	if err != nil {
		panic(fmt.Sprintf("Error parsing groupversion: %v", err))
	}
	externalGroupVersion := schema.GroupVersion{Group: servicecatalog.GroupName, Version: api.Scheme.PrioritizedVersionsForGroup(servicecatalog.GroupName)[0].Version}
	testapi.Groups[servicecatalog.GroupName] = testapi.NewTestGroup(groupVersion, servicecatalog.SchemeGroupVersion, api.Scheme.KnownTypes(servicecatalog.SchemeGroupVersion), api.Scheme.KnownTypes(externalGroupVersion))
}
func roundTrip(t *testing.T, obj runtime.Object) runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	codec, err := testapi.GetCodecForObject(obj)
	if err != nil {
		t.Fatalf("%v\n %#v", err, obj)
	}
	data, err := runtime.Encode(codec, obj)
	if err != nil {
		t.Fatalf("%v\n %#v", err, obj)
	}
	obj2, err := runtime.Decode(codec, data)
	if err != nil {
		t.Fatalf("%v\nData: %s\nSource: %#v", err, string(data), obj)
	}
	obj3 := reflect.New(reflect.TypeOf(obj).Elem()).Interface().(runtime.Object)
	err = api.Scheme.Convert(obj2, obj3, nil)
	if err != nil {
		t.Fatalf("%v\nSource: %#v", err, obj2)
	}
	return obj3
}
func TestSetDefaultClusterServiceBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		broker		*versioned.ClusterServiceBroker
		behavior	versioned.ServiceBrokerRelistBehavior
		duration	*metav1.Duration
	}{{name: "neither duration or behavior set", broker: &versioned.ClusterServiceBroker{}, behavior: versioned.ServiceBrokerRelistBehaviorDuration, duration: &metav1.Duration{Duration: 15 * time.Minute}}, {name: "behavior set to manual", broker: func() *versioned.ClusterServiceBroker {
		b := &versioned.ClusterServiceBroker{}
		b.Spec.RelistBehavior = versioned.ServiceBrokerRelistBehaviorManual
		return b
	}(), behavior: versioned.ServiceBrokerRelistBehaviorManual, duration: nil}, {name: "behavior set to duration but no duration provided", broker: func() *versioned.ClusterServiceBroker {
		b := &versioned.ClusterServiceBroker{}
		b.Spec.RelistBehavior = versioned.ServiceBrokerRelistBehaviorDuration
		return b
	}(), behavior: versioned.ServiceBrokerRelistBehaviorDuration, duration: &metav1.Duration{Duration: 15 * time.Minute}}}
	for _, tc := range cases {
		o := roundTrip(t, runtime.Object(tc.broker))
		ab := o.(*versioned.ClusterServiceBroker)
		actualSpec := ab.Spec
		if tc.behavior != actualSpec.RelistBehavior {
			t.Errorf("%v: unexpected default RelistBehavior: expected %v, got %v", tc.name, tc.behavior, actualSpec.RelistBehavior)
		}
	}
}
