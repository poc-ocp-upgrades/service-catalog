package v1beta1

import (
	"encoding/json"
	"testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestConvertServiceClassToProperties(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name	string
		sc	*ServiceClass
		json	string
	}{{name: "nil object", json: "{}"}, {name: "normal object", sc: &ServiceClass{ObjectMeta: metav1.ObjectMeta{Name: "service-class"}, Spec: ServiceClassSpec{CommonServiceClassSpec: CommonServiceClassSpec{ExternalName: "external-class-name", ExternalID: "external-id"}}}, json: `{"name":"service-class","spec.externalID":"external-id","spec.externalName":"external-class-name"}`}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := ConvertServiceClassToProperties(tc.sc)
			if p == nil {
				t.Fatalf("Failed to create Properties object from %+v", tc.sc)
			}
			b, err := json.Marshal(p)
			if err != nil {
				t.Fatalf("Unexpected error with json marshal, %v", err)
			}
			js := string(b)
			if js != tc.json {
				t.Fatalf("Failed to create expected Properties object,\n\texpected: \t%q,\n \tgot: \t\t%q", tc.json, js)
			}
		})
	}
}
func TestConvertServicePlanToProperties(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name	string
		sp	*ServicePlan
		json	string
	}{{name: "nil object", json: "{}"}, {name: "normal object", sp: &ServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "service-plan"}, Spec: ServicePlanSpec{CommonServicePlanSpec: CommonServicePlanSpec{ExternalName: "external-plan-name", ExternalID: "external-id", Free: true}, ServiceClassRef: LocalObjectReference{Name: "service-class-name"}}}, json: `{"name":"service-plan","spec.externalID":"external-id","spec.externalName":"external-plan-name","spec.free":"true","spec.serviceClass.name":"service-class-name"}`}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := ConvertServicePlanToProperties(tc.sp)
			if p == nil {
				t.Fatalf("Failed to create Properties object from %+v", tc.sp)
			}
			b, err := json.Marshal(p)
			if err != nil {
				t.Fatalf("Unexpected error with json marshal, %v", err)
			}
			js := string(b)
			if js != tc.json {
				t.Fatalf("Failed to create expected Properties object,\n\texpected: \t%q,\n \tgot: \t\t%q", tc.json, js)
			}
		})
	}
}
func TestConvertClusterServiceClassToProperties(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name	string
		sc	*ClusterServiceClass
		json	string
	}{{name: "nil object", json: "{}"}, {name: "normal object", sc: &ClusterServiceClass{ObjectMeta: metav1.ObjectMeta{Name: "service-class"}, Spec: ClusterServiceClassSpec{CommonServiceClassSpec: CommonServiceClassSpec{ExternalName: "external-class-name", ExternalID: "external-id"}}}, json: `{"name":"service-class","spec.externalID":"external-id","spec.externalName":"external-class-name"}`}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := ConvertClusterServiceClassToProperties(tc.sc)
			if p == nil {
				t.Fatalf("Failed to create Properties object from %+v", tc.sc)
			}
			b, err := json.Marshal(p)
			if err != nil {
				t.Fatalf("Unexpected error with json marshal, %v", err)
			}
			js := string(b)
			if js != tc.json {
				t.Fatalf("Failed to create expected Properties object,\n\texpected: \t%q,\n \tgot: \t\t%q", tc.json, js)
			}
		})
	}
}
func TestConvertClusterServicePlanToProperties(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name	string
		sp	*ClusterServicePlan
		json	string
	}{{name: "nil object", json: "{}"}, {name: "normal object", sp: &ClusterServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "service-plan"}, Spec: ClusterServicePlanSpec{CommonServicePlanSpec: CommonServicePlanSpec{ExternalName: "external-plan-name", ExternalID: "external-id", Free: true}, ClusterServiceClassRef: ClusterObjectReference{Name: "cluster-service-class-name"}}}, json: `{"name":"service-plan","spec.clusterServiceClass.name":"cluster-service-class-name","spec.externalID":"external-id","spec.externalName":"external-plan-name","spec.free":"true"}`}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := ConvertClusterServicePlanToProperties(tc.sp)
			if p == nil {
				t.Fatalf("Failed to create Properties object from %+v", tc.sp)
			}
			b, err := json.Marshal(p)
			if err != nil {
				t.Fatalf("Unexpected error with json marshal, %v", err)
			}
			js := string(b)
			if js != tc.json {
				t.Fatalf("Failed to create expected Properties object,\n\texpected: \t%q,\n \tgot: \t\t%q", tc.json, js)
			}
		})
	}
}
