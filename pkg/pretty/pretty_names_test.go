package pretty

import (
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPrettyNames(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := `ServiceInstance (K8S: "k8s" ExternalName: "extern")`
	g := Name(ServiceInstance, "k8s", "extern")
	if g != e {
		t.Fatalf("Unexpected value of PrettyName String; expected %v, got %v", e, g)
	}
}
func TestServiceInstanceName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := `ServiceInstance "namespace/name"`
	serviceInstance := &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: "name", Namespace: "namespace"}}
	g := ServiceInstanceName(serviceInstance)
	if g != e {
		t.Fatalf("Unexpected value of PrettyName String; expected %v, got %v", e, g)
	}
}
func TestClusterServiceBrokerName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	e := `ClusterServiceBroker "brokerName"`
	g := ClusterServiceBrokerName("brokerName")
	if g != e {
		t.Fatalf("Unexpected value of PrettyName String; expected %v, got %v", e, g)
	}
}
func TestClusterServiceClassName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass := &v1beta1.ClusterServiceClass{ObjectMeta: metav1.ObjectMeta{Name: "service-class"}, Spec: v1beta1.ClusterServiceClassSpec{CommonServiceClassSpec: v1beta1.CommonServiceClassSpec{ExternalName: "external-class-name"}}}
	e := `ClusterServiceClass (K8S: "service-class" ExternalName: "external-class-name")`
	g := ClusterServiceClassName(serviceClass)
	if g != e {
		t.Fatalf("Unexpected value of PrettyName String; expected %v, got %v", e, g)
	}
}
func TestClusterServicePlanName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	servicePlan := &v1beta1.ClusterServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "service-plan"}, Spec: v1beta1.ClusterServicePlanSpec{CommonServicePlanSpec: v1beta1.CommonServicePlanSpec{ExternalName: "external-plan-name"}}}
	e := `ClusterServicePlan (K8S: "service-plan" ExternalName: "external-plan-name")`
	g := ClusterServicePlanName(servicePlan)
	if g != e {
		t.Fatalf("Unexpected value of PrettyName String; expected %v, got %v", e, g)
	}
}
