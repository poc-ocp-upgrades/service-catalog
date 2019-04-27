package clusterservicebroker

import (
	"testing"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	sctestutil "github.com/kubernetes-incubator/service-catalog/test/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func clusterServiceBrokerWithOldSpec() *sc.ClusterServiceBroker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &sc.ClusterServiceBroker{ObjectMeta: metav1.ObjectMeta{Generation: 1}, Spec: sc.ClusterServiceBrokerSpec{CommonServiceBrokerSpec: sc.CommonServiceBrokerSpec{URL: "https://kubernetes.default.svc:443/brokers/template.k8s.io"}}, Status: sc.ClusterServiceBrokerStatus{CommonServiceBrokerStatus: sc.CommonServiceBrokerStatus{Conditions: []sc.ServiceBrokerCondition{{Type: sc.ServiceBrokerConditionReady, Status: sc.ConditionFalse}}}}}
}
func clusterServiceBrokerWithNewSpec() *sc.ClusterServiceBroker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	b := clusterServiceBrokerWithOldSpec()
	b.Spec.URL = "new"
	return b
}
func TestClusterServiceBrokerStrategyTrivial(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if clusterServiceBrokerRESTStrategies.NamespaceScoped() {
		t.Errorf("clusterservicebroker must not be namespace scoped")
	}
	if clusterServiceBrokerRESTStrategies.AllowCreateOnUpdate() {
		t.Errorf("clusterservicebroker should not allow create on update")
	}
	if clusterServiceBrokerRESTStrategies.AllowUnconditionalUpdate() {
		t.Errorf("clusterservicebroker should not allow unconditional update")
	}
}
func TestClusterServiceBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker := &sc.ClusterServiceBroker{Spec: sc.ClusterServiceBrokerSpec{CommonServiceBrokerSpec: sc.CommonServiceBrokerSpec{URL: "abcd"}}, Status: sc.ClusterServiceBrokerStatus{CommonServiceBrokerStatus: sc.CommonServiceBrokerStatus{Conditions: nil}}}
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	clusterServiceBrokerRESTStrategies.PrepareForCreate(createContext, broker)
	if broker.Status.Conditions == nil {
		t.Fatalf("Fresh clusterservicebroker should have empty status")
	}
	if len(broker.Status.Conditions) != 0 {
		t.Fatalf("Fresh clusterservicebroker should have empty status")
	}
}
func TestClusterServiceBrokerUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name				string
		older				*sc.ClusterServiceBroker
		newer				*sc.ClusterServiceBroker
		shouldGenerationIncrement	bool
	}{{name: "no spec change", older: clusterServiceBrokerWithOldSpec(), newer: clusterServiceBrokerWithOldSpec(), shouldGenerationIncrement: false}, {name: "spec change", older: clusterServiceBrokerWithOldSpec(), newer: clusterServiceBrokerWithNewSpec(), shouldGenerationIncrement: true}}
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	for i := range cases {
		clusterServiceBrokerRESTStrategies.PrepareForUpdate(createContext, cases[i].newer, cases[i].older)
		if cases[i].shouldGenerationIncrement {
			if e, a := cases[i].older.Generation+1, cases[i].newer.Generation; e != a {
				t.Fatalf("%v: expected %v, got %v for generation", cases[i].name, e, a)
			}
		} else {
			if e, a := cases[i].older.Generation, cases[i].newer.Generation; e != a {
				t.Fatalf("%v: expected %v, got %v for generation", cases[i].name, e, a)
			}
		}
	}
}
func TestClusterServiceBrokerUpdateForRelistRequests(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		oldValue	int64
		newValue	int64
		expectedValue	int64
	}{{name: "both default", oldValue: 0, newValue: 0, expectedValue: 0}, {name: "old default", oldValue: 0, newValue: 1, expectedValue: 1}, {name: "new default", oldValue: 1, newValue: 0, expectedValue: 1}, {name: "neither default", oldValue: 1, newValue: 2, expectedValue: 2}}
	for _, tc := range cases {
		oldBroker := clusterServiceBrokerWithOldSpec()
		oldBroker.Spec.RelistRequests = tc.oldValue
		newClusterServiceBroker := clusterServiceBrokerWithOldSpec()
		newClusterServiceBroker.Spec.RelistRequests = tc.newValue
		creatorUserName := "creator"
		createContext := sctestutil.ContextWithUserName(creatorUserName)
		clusterServiceBrokerRESTStrategies.PrepareForUpdate(createContext, newClusterServiceBroker, oldBroker)
		if e, a := tc.expectedValue, newClusterServiceBroker.Spec.RelistRequests; e != a {
			t.Errorf("%s: got unexpected RelistRequests: expected %v, got %v", tc.name, e, a)
		}
	}
}
