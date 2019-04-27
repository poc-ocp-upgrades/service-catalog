package binding

import (
	"fmt"
	"testing"
	sctestutil "github.com/kubernetes-incubator/service-catalog/test/util"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getTestInstanceCredential() *servicecatalog.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &servicecatalog.ServiceBinding{ObjectMeta: metav1.ObjectMeta{Generation: 1}, Spec: servicecatalog.ServiceBindingSpec{InstanceRef: servicecatalog.LocalObjectReference{Name: "some-string"}}, Status: servicecatalog.ServiceBindingStatus{Conditions: []servicecatalog.ServiceBindingCondition{{Type: servicecatalog.ServiceBindingConditionReady, Status: servicecatalog.ConditionTrue}}}}
}
func TestInstanceCredentialUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name				string
		older				*servicecatalog.ServiceBinding
		newer				*servicecatalog.ServiceBinding
		shouldGenerationIncrement	bool
	}{{name: "no spec change", older: getTestInstanceCredential(), newer: getTestInstanceCredential()}}
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	for _, tc := range cases {
		bindingRESTStrategies.PrepareForUpdate(createContext, tc.newer, tc.older)
		expectedGeneration := tc.older.Generation
		if tc.shouldGenerationIncrement {
			expectedGeneration = expectedGeneration + 1
		}
		if e, a := expectedGeneration, tc.newer.Generation; e != a {
			t.Errorf("%v: expected %v, got %v for generation", tc.name, e, a)
		}
	}
}
func TestInstanceCredentialUserInfo(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prevOrigIDEnablement := sctestutil.EnableOriginatingIdentity(t, true)
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=%v", scfeatures.OriginatingIdentity, prevOrigIDEnablement))
	creatorUserName := "creator"
	createdInstanceCredential := getTestInstanceCredential()
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	bindingRESTStrategies.PrepareForCreate(createContext, createdInstanceCredential)
	if e, a := creatorUserName, createdInstanceCredential.Spec.UserInfo.Username; e != a {
		t.Errorf("unexpected user info in created spec: expected %q, got %q", e, a)
	}
	deleterUserName := "deleter"
	deletedInstanceCredential := getTestInstanceCredential()
	deleteContext := sctestutil.ContextWithUserName(deleterUserName)
	bindingRESTStrategies.CheckGracefulDelete(deleteContext, deletedInstanceCredential, nil)
	if e, a := deleterUserName, deletedInstanceCredential.Spec.UserInfo.Username; e != a {
		t.Errorf("unexpected user info in deleted spec: expected %q, got %q", e, a)
	}
}
func TestExternalIDSet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	createdInstanceCredential := getTestInstanceCredential()
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	bindingRESTStrategies.PrepareForCreate(createContext, createdInstanceCredential)
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
	createdInstanceCredential := getTestInstanceCredential()
	createdInstanceCredential.Spec.ExternalID = userExternalID
	creatorUserName := "creator"
	createContext := sctestutil.ContextWithUserName(creatorUserName)
	bindingRESTStrategies.PrepareForCreate(createContext, createdInstanceCredential)
	if createdInstanceCredential.Spec.ExternalID != userExternalID {
		t.Errorf("Modified user provided ExternalID to %q", createdInstanceCredential.Spec.ExternalID)
	}
}
