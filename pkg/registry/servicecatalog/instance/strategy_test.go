package instance

import (
	"fmt"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	sctestutil "github.com/kubernetes-incubator/service-catalog/test/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
)

func getTestInstance() *servicecatalog.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Generation: 1}, Spec: servicecatalog.ServiceInstanceSpec{PlanReference: servicecatalog.PlanReference{ClusterServiceClassExternalName: "test-clusterserviceclass", ClusterServicePlanExternalName: "test-clusterserviceplan"}, ClusterServiceClassRef: &servicecatalog.ClusterObjectReference{}, ClusterServicePlanRef: &servicecatalog.ClusterObjectReference{}, UserInfo: &servicecatalog.UserInfo{Username: "some-user"}}, Status: servicecatalog.ServiceInstanceStatus{Conditions: []servicecatalog.ServiceInstanceCondition{{Type: servicecatalog.ServiceInstanceConditionReady, Status: servicecatalog.ConditionTrue}}}}
}
func TestInstanceUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name				string
		older				*servicecatalog.ServiceInstance
		newer				*servicecatalog.ServiceInstance
		shouldGenerationIncrement	bool
		shouldPlanRefClear		bool
	}{{name: "no spec change", older: getTestInstance(), newer: getTestInstance()}, {name: "UpdateRequest increment", older: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.UpdateRequests = 1
		return i
	}(), newer: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.UpdateRequests = 2
		return i
	}(), shouldGenerationIncrement: true}, {name: "external plan name change", older: getTestInstance(), newer: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.ClusterServicePlanExternalName = "new-plan"
		return i
	}(), shouldGenerationIncrement: true, shouldPlanRefClear: true}, {name: "external plan id change", older: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.ClusterServiceClassExternalName = ""
		i.Spec.ClusterServicePlanExternalName = ""
		i.Spec.ClusterServiceClassExternalID = "test-clusterserviceclass"
		i.Spec.ClusterServicePlanExternalID = "test-clusterserviceplan"
		return i
	}(), newer: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.ClusterServiceClassExternalName = ""
		i.Spec.ClusterServicePlanExternalName = ""
		i.Spec.ClusterServiceClassExternalID = "test-clusterserviceclass"
		i.Spec.ClusterServicePlanExternalID = "new plan"
		return i
	}(), shouldGenerationIncrement: true, shouldPlanRefClear: true}, {name: "k8s plan change", older: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.ClusterServiceClassExternalName = ""
		i.Spec.ClusterServicePlanExternalName = ""
		i.Spec.ClusterServiceClassName = "test-clusterserviceclass"
		i.Spec.ClusterServicePlanName = "test-clusterserviceplan"
		return i
	}(), newer: func() *servicecatalog.ServiceInstance {
		i := getTestInstance()
		i.Spec.ClusterServiceClassExternalName = ""
		i.Spec.ClusterServicePlanExternalName = ""
		i.Spec.ClusterServiceClassName = "test-clusterserviceclass"
		i.Spec.ClusterServicePlanName = "new plan"
		return i
	}(), shouldGenerationIncrement: true, shouldPlanRefClear: true}}
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	for _, tc := range cases {
		instanceRESTStrategies.PrepareForUpdate(createContext, tc.newer, tc.older)
		expectedGeneration := tc.older.Generation
		if tc.shouldGenerationIncrement {
			expectedGeneration = expectedGeneration + 1
		}
		if e, a := expectedGeneration, tc.newer.Generation; e != a {
			t.Errorf("%v: expected %v, got %v for generation", tc.name, e, a)
			continue
		}
		if tc.shouldPlanRefClear {
			if tc.newer.Spec.ClusterServicePlanRef != nil {
				t.Errorf("%v: expected ServicePlanRef to be nil", tc.name)
			}
		} else {
			if tc.newer.Spec.ClusterServicePlanRef == nil {
				t.Errorf("%v: expected ServicePlanRef to not be nil", tc.name)
			}
		}
	}
}
func TestInstanceUserInfo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prevOrigIDEnablement := sctestutil.EnableOriginatingIdentity(t, true)
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=%v", scfeatures.OriginatingIdentity, prevOrigIDEnablement))
	creatorUserName := "creator"
	createdInstance := getTestInstance()
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	instanceRESTStrategies.PrepareForCreate(createContext, createdInstance)
	if e, a := creatorUserName, createdInstance.Spec.UserInfo.Username; e != a {
		t.Errorf("unexpected user info in created spec: expected %v, got %v", e, a)
	}
	updaterUserName := "updater"
	updatedInstance := getTestInstance()
	updatedInstance.Spec.UpdateRequests = updatedInstance.Spec.UpdateRequests + 1
	updateContext := sctestutil.ContextWithUserName(updaterUserName)
	instanceRESTStrategies.PrepareForUpdate(updateContext, updatedInstance, createdInstance)
	if e, a := updaterUserName, updatedInstance.Spec.UserInfo.Username; e != a {
		t.Errorf("unexpected user info in updated spec: expected %v, got %v", e, a)
	}
	deleterUserName := "deleter"
	deletedInstance := getTestInstance()
	deleteContext := sctestutil.ContextWithUserName(deleterUserName)
	instanceRESTStrategies.CheckGracefulDelete(deleteContext, deletedInstance, nil)
	if e, a := deleterUserName, deletedInstance.Spec.UserInfo.Username; e != a {
		t.Errorf("unexpected user info in deleted spec: expected %v, got %v", e, a)
	}
}
func TestInstanceUpdateForUpdateRequests(t *testing.T) {
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
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	for _, tc := range cases {
		oldInstance := getTestInstance()
		oldInstance.Spec.UpdateRequests = tc.oldValue
		newInstance := getTestInstance()
		newInstance.Spec.UpdateRequests = tc.newValue
		instanceRESTStrategies.PrepareForUpdate(createContext, newInstance, oldInstance)
		if e, a := tc.expectedValue, newInstance.Spec.UpdateRequests; e != a {
			t.Errorf("%s: got unexpected UpdateRequests: expected %v, got %v", tc.name, e, a)
		}
	}
}
func TestExternalIDSet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	createdInstanceCredential := getTestInstance()
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	instanceRESTStrategies.PrepareForCreate(createContext, createdInstanceCredential)
	if createdInstanceCredential.Spec.ExternalID == "" {
		t.Error("Expected an ExternalID to be set, but got none")
	}
}
func TestExternalIDUserProvided(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	userExternalID := "my-id"
	createdInstanceCredential := getTestInstance()
	createdInstanceCredential.Spec.ExternalID = userExternalID
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	instanceRESTStrategies.PrepareForCreate(createContext, createdInstanceCredential)
	if createdInstanceCredential.Spec.ExternalID != userExternalID {
		t.Errorf("Modified user provided ExternalID to %q", createdInstanceCredential.Spec.ExternalID)
	}
}
