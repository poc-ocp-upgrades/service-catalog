package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	fakeosb "github.com/pmorie/go-open-service-broker-client/v2/fake"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	"github.com/kubernetes-incubator/service-catalog/test/fake"
	sctestutil "github.com/kubernetes-incubator/service-catalog/test/util"
	corev1 "k8s.io/api/core/v1"
	clientgotesting "k8s.io/client-go/testing"
)

const (
	lastOperationDescription = "testdescr"
)

func TestReconcileServiceInstanceNonExistentClusterServiceClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, noFakeActions())
	instance := &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: testServiceInstanceName, Generation: 1}, Spec: v1beta1.ServiceInstanceSpec{PlanReference: v1beta1.PlanReference{ClusterServiceClassExternalName: "nothere", ClusterServicePlanExternalName: "nothere"}, ExternalID: testServiceInstanceGUID}}
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("nothere is a service class that cannot be referenced by the service instance as it does not exist.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.externalName", instance.Spec.ClusterServiceClassExternalName)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	updatedServiceInstance := assertUpdateStatus(t, actions[1], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorNonexistentClusterServiceClassReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServiceClassReason).msgf("References a non-existent ClusterServiceClass %c or there is more than one (found: 0)", instance.Spec.PlanReference)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceNonExistentClusterServiceClassWithK8SName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, noFakeActions())
	instance := &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: testServiceInstanceName, Generation: 1}, Spec: v1beta1.ServiceInstanceSpec{PlanReference: v1beta1.PlanReference{ClusterServiceClassName: "nothereclass", ClusterServicePlanName: "nothereplan"}, ExternalID: testServiceInstanceGUID}}
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("nothere is a service class that cannot be referenced by the service instance as it does not exist.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorNonexistentClusterServiceClassReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServiceClassReason).msgf("References a non-existent ClusterServiceClass %c", instance.Spec.PlanReference)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceNonExistentClusterServiceBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("The broker referenced by the instance exists when it should not.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorNonexistentClusterServiceBrokerReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServiceBrokerReason).msgf("The instance references a non-existent broker %q", "test-clusterservicebroker")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceWithNotExistingBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	testController.brokerClientManager.RemoveBrokerClient(NewClusterServiceBrokerKey(getTestClusterServiceBroker().Name))
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	fakeKubeClient.AddReactor("get", "secrets", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, nil, errors.New("no secret defined")
	})
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("There was no secret to be found, but does_not_exist/auth-name was found.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorNonexistentClusterServiceBrokerReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServiceBrokerReason).msgf("The instance references a non-existent broker %q", "test-clusterservicebroker")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceNonExistentClusterServicePlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: testServiceInstanceName, Generation: 1}, Spec: v1beta1.ServiceInstanceSpec{PlanReference: v1beta1.PlanReference{ClusterServiceClassExternalName: testClusterServiceClassName, ClusterServicePlanExternalName: "nothere"}, ClusterServiceClassRef: &v1beta1.ClusterObjectReference{Name: testClusterServiceClassGUID}, ExternalID: testServiceInstanceGUID}}
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("The service plan nothere should not exist to be referenced.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=nothere,spec.clusterServiceBrokerName=test-clusterservicebroker,spec.clusterServiceClassRef.name=cscguid")}
	assertList(t, actions[0], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateStatus(t, actions[1], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorNonexistentClusterServicePlanReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServicePlanReason).msgf(`References a non-existent ClusterServicePlan %b on ClusterServiceClass %s %c or there is more than one (found: 0)`, instance.Spec.PlanReference, instance.Spec.ClusterServiceClassRef.Name, instance.Spec.PlanReference)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceNonExistentClusterServicePlanK8SName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: testServiceInstanceName, Generation: 1}, Spec: v1beta1.ServiceInstanceSpec{PlanReference: v1beta1.PlanReference{ClusterServiceClassName: testClusterServiceClassGUID, ClusterServicePlanName: "nothereplan"}, ClusterServiceClassRef: &v1beta1.ClusterObjectReference{Name: testClusterServiceClassGUID}, ExternalID: testServiceInstanceGUID}}
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("The service plan nothere should not exist to be referenced.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorNonexistentClusterServicePlanReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServicePlanReason).msgf("References a non-existent ClusterServicePlan %v", instance.Spec.PlanReference)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceWithParameters(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	type secretDef struct {
		name	string
		data	map[string][]byte
	}
	cases := []struct {
		name					string
		params					[]byte
		paramsFrom				[]v1beta1.ParametersFromSource
		secrets					[]secretDef
		expectedParams				map[string]interface{}
		expectedParamsWithSecretsRedacted	map[string]interface{}
		expectedError				bool
	}{{name: "no params", expectedParams: nil}, {name: "plain params", params: []byte(`{"Name":"test-param","Args":{"first":"first-arg","second":"second-arg"}}`), expectedParams: map[string]interface{}{"Name": "test-param", "Args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}}, expectedParamsWithSecretsRedacted: map[string]interface{}{"Name": "test-param", "Args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}}}, {name: "secret params", paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "secret-key"}}}, secrets: []secretDef{{name: "secret-name", data: map[string][]byte{"secret-key": []byte(`{"A":"B","C":{"D":"E","F":"G"}}`)}}}, expectedParams: map[string]interface{}{"A": "B", "C": map[string]interface{}{"D": "E", "F": "G"}}, expectedParamsWithSecretsRedacted: map[string]interface{}{"A": "<redacted>", "C": "<redacted>"}}, {name: "plain and secret params", params: []byte(`{"Name":"test-param","Args":{"first":"first-arg","second":"second-arg"}}`), paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "secret-key"}}}, secrets: []secretDef{{name: "secret-name", data: map[string][]byte{"secret-key": []byte(`{"A":"B","C":{"D":"E","F":"G"}}`)}}}, expectedParams: map[string]interface{}{"Name": "test-param", "Args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}, "A": "B", "C": map[string]interface{}{"D": "E", "F": "G"}}, expectedParamsWithSecretsRedacted: map[string]interface{}{"Name": "test-param", "Args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}, "A": "<redacted>", "C": "<redacted>"}}, {name: "bad params", params: []byte("bad"), expectedError: true}, {name: "missing secret", paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "secret-key"}}}, expectedError: true}, {name: "missing secret key", paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "other-secret-key"}}}, secrets: []secretDef{{name: "secret-name", data: map[string][]byte{"secret-key": []byte(`bad`)}}}, expectedError: true}, {name: "empty secret data", paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "secret-key"}}}, secrets: []secretDef{{name: "secret-name", data: map[string][]byte{}}}, expectedError: true}, {name: "bad secret data", paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "secret-key"}}}, secrets: []secretDef{{name: "secret-name", data: map[string][]byte{"secret-key": []byte(`bad`)}}}, expectedError: true}, {name: "no params in secret data", paramsFrom: []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "secret-name", Key: "secret-key"}}}, secrets: []secretDef{{name: "secret-name", data: map[string][]byte{"secret-key": []byte(`{}`)}}}}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			for _, s := range tc.secrets {
				fakeKubeClient.PrependReactor("get", "secrets", func(action clientgotesting.Action) (bool, runtime.Object, error) {
					getAction, ok := action.(clientgotesting.GetAction)
					if !ok {
						return true, nil, apierrors.NewInternalError(fmt.Errorf("could not convert get secrets action to a GetAction: %T", action))
					}
					if getAction.GetName() != s.name {
						return false, nil, nil
					}
					secret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Namespace: testNamespace, Name: s.name}, Data: s.data}
					return true, secret, nil
				})
			}
			instance := getTestServiceInstanceWithClusterRefs()
			if tc.params != nil {
				instance.Spec.Parameters = &runtime.RawExtension{Raw: tc.params}
			}
			instance.Spec.ParametersFrom = tc.paramsFrom
			err := reconcileServiceInstance(t, testController, instance)
			if tc.expectedError {
				if err == nil {
					t.Fatalf("Reconcile expected to fail")
				}
			} else {
				if err != nil {
					t.Fatalf("Reconcile not expected to fail : %v", err)
				}
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 0)
			expectedKubeActions := []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}
			for range tc.paramsFrom {
				expectedKubeActions = append(expectedKubeActions, kubeClientAction{verb: "get", resourceName: "secrets", checkType: checkGetActionType})
			}
			kubeActions := fakeKubeClient.Actions()
			if err := checkKubeClientActions(kubeActions, expectedKubeActions); err != nil {
				t.Fatal(err)
			}
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			events := getRecordedEvents(testController)
			if tc.expectedError {
				assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorWithParametersReason, instance)
				expectedEvent := warningEventBuilder(errorWithParametersReason).msg("failed to prepare parameters")
				if err := checkEventPrefixes(events, expectedEvent.stringArr()); err != nil {
					t.Fatal(err)
				}
			} else {
				assertServiceInstanceOperationInProgressWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, tc.expectedParamsWithSecretsRedacted, generateChecksumOfParametersOrFail(t, tc.expectedParams), instance)
				if err := checkEvents(events, []string{}); err != nil {
					t.Fatal(err)
				}
			}
			if tc.expectedError {
				return
			}
			fakeCatalogClient.ClearActions()
			fakeKubeClient.ClearActions()
			instance = updatedServiceInstance.(*v1beta1.ServiceInstance)
			err = reconcileServiceInstance(t, testController, instance)
			if err != nil {
				t.Fatalf("Reconcile not expected to fail : %v", err)
			}
			brokerActions = fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext, Parameters: tc.expectedParams})
			actions = fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance = assertUpdateStatus(t, actions[0], instance)
			assertServiceInstanceOperationSuccessWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, tc.expectedParamsWithSecretsRedacted, generateChecksumOfParametersOrFail(t, tc.expectedParams), instance)
			kubeActions = fakeKubeClient.Actions()
			if err := checkKubeClientActions(kubeActions, expectedKubeActions); err != nil {
				t.Fatal(err)
			}
			events = getRecordedEvents(testController)
			expectedEvent := normalEventBuilder(successProvisionReason).msg("The instance was provisioned successfully")
			if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestReconcileServiceInstanceResolvesReferences(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sc := getTestClusterServiceClass()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(sc)
	sp := getTestClusterServicePlan()
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(sp)
	instance := getTestServiceInstance()
	var scItems []v1beta1.ClusterServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: scItems}, nil
	})
	var spItems []v1beta1.ClusterServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: spItems}, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 3)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.externalName", instance.Spec.ClusterServiceClassExternalName)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	listRestrictions = clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=test-clusterserviceplan,spec.clusterServiceBrokerName=test-clusterservicebroker,spec.clusterServiceClassRef.name=cscguid")}
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateReference(t, actions[2], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ClusterServiceClassRef == nil || updateObject.Spec.ClusterServiceClassRef.Name != "cscguid" {
		t.Fatalf("ClusterServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ClusterServicePlanRef == nil || updateObject.Spec.ClusterServicePlanRef.Name != "cspguid" {
		t.Fatalf("ClusterServicePlanRef was not resolved correctly during reconcile")
	}
	assertNumberOfActions(t, fakeKubeClient.Actions(), 0)
	assertNumEvents(t, getRecordedEvents(testController), 0)
}
func TestReconcileServiceInstanceAppliesDefaultProvisioningParams(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.ServicePlanDefaults))
	if err != nil {
		t.Fatalf("Could not enable ServicePlanDefaults feature flag.")
	}
	_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sc := getTestClusterServiceClass()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(sc)
	sp := getTestClusterServicePlan()
	classParams := `{"secure": false, "class-default": 1}`
	sc.Spec.DefaultProvisionParameters = &runtime.RawExtension{Raw: []byte(classParams)}
	planParams := `{"secure": true, "plan-default": 2}`
	sp.Spec.DefaultProvisionParameters = &runtime.RawExtension{Raw: []byte(planParams)}
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(sp)
	instance := getTestServiceInstanceWithClusterRefs()
	var scItems []v1beta1.ClusterServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: scItems}, nil
	})
	var spItems []v1beta1.ClusterServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: spItems}, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	updatedServiceInstance := assertUpdate(t, actions[0], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	wantParams := `{"class-default":1,"plan-default":2,"secure":true}`
	gotParams := string(updateObject.Spec.Parameters.Raw)
	if gotParams != wantParams {
		t.Fatalf("DefaultProvisioningParameters was not applied to the service instance during reconcile.\n\nWANT: %v\nGOT: %v", wantParams, gotParams)
	}
	updatedServiceInstance = assertUpdateStatus(t, actions[1], instance)
	updateObject, ok = updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	gotParams = string(updateObject.Status.DefaultProvisionParameters.Raw)
	if gotParams != wantParams {
		t.Fatalf("DefaultProvisioningParameters was not persisted to the service instance status during reconcile.\n\nWANT: %v\nGOT: %v", wantParams, gotParams)
	}
}
func TestReconcileServiceInstanceRespectsServicePlanDefaultsFeatureGate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.ServicePlanDefaults))
	if err != nil {
		t.Fatalf("Could not disable ServicePlanDefaults feature flag.")
	}
	_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sc := getTestClusterServiceClass()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(sc)
	sp := getTestClusterServicePlan()
	defaultProvParams := `{"secure": true}`
	sp.Spec.DefaultProvisionParameters = &runtime.RawExtension{Raw: []byte(defaultProvParams)}
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(sp)
	instance := getTestServiceInstanceWithClusterRefs()
	var scItems []v1beta1.ClusterServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: scItems}, nil
	})
	var spItems []v1beta1.ClusterServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: spItems}, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Status.DefaultProvisionParameters != nil {
		t.Fatal("DefaultProvisioningParameters should not be set on the status because the feature is disabled")
	}
}
func TestReconcileServiceInstanceResolvesReferencesClusterServiceClassRefAlreadySet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sc := getTestClusterServiceClass()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(sc)
	sp := getTestClusterServicePlan()
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(sp)
	instance := getTestServiceInstance()
	instance.Spec.ClusterServiceClassRef = &v1beta1.ClusterObjectReference{Name: testClusterServiceClassGUID}
	var scItems []v1beta1.ClusterServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: scItems}, nil
	})
	var spItems []v1beta1.ClusterServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: spItems}, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=test-clusterserviceplan,spec.clusterServiceBrokerName=test-clusterservicebroker,spec.clusterServiceClassRef.name=cscguid")}
	assertList(t, actions[0], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateReference(t, actions[1], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ClusterServiceClassRef == nil || updateObject.Spec.ClusterServiceClassRef.Name != "cscguid" {
		t.Fatalf("ClusterServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ClusterServicePlanRef == nil || updateObject.Spec.ClusterServicePlanRef.Name != "cspguid" {
		t.Fatalf("ClusterServicePlanRef was not resolved correctly during reconcile")
	}
	assertNumberOfActions(t, fakeKubeClient.Actions(), 0)
	assertNumEvents(t, getRecordedEvents(testController), 0)
}
func TestReconcileServiceInstanceWithProvisionCallFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Error: errors.New("fake creation failure")}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("Should not be able to make the ServiceInstance")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestRetriableError(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, errorErrorCallingProvisionReason, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorErrorCallingProvisionReason).msg("The provision call failed and will be retried:").msgf("Error communicating with broker for provisioning:").msg("fake creation failure")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceWithTemporaryProvisionFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Error: osb.HTTPStatusCodeError{StatusCode: http.StatusInternalServerError, ErrorMessage: strPtr("InternalServerError"), Description: strPtr("Something went wrong!")}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("Reconcile not expected to fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	expectedKubeActions := []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, expectedKubeActions); err != nil {
		t.Fatal(err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	events := getRecordedEvents(testController)
	updatedServiceInstance = assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	instance = updatedServiceInstance.(*v1beta1.ServiceInstance)
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("Should not be able to make the ServiceInstance")
	}
	brokerActions = fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	kubeActions = fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions = fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance = assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestFailingErrorStartOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, startingInstanceOrphanMitigationReason, "", errorProvisionCallFailedReason, instance)
	events = getRecordedEvents(testController)
	message := fmt.Sprintf("Error provisioning ServiceInstance of ClusterServiceClass (K8S: %q ExternalName: %q) at ClusterServiceBroker %q: Status: %v; ErrorMessage: %s", "cscguid", "test-clusterserviceclass", "test-clusterservicebroker", 500, "InternalServerError; Description: Something went wrong!; ResponseError: <nil>")
	expectedProvisionCallEvent := warningEventBuilder(errorProvisionCallFailedReason).msg(message)
	expectedOrphanMitigationEvent := warningEventBuilder(startingInstanceOrphanMitigationReason).msg("The instance provision call failed with an ambiguous error; attempting to deprovision the instance in order to mitigate an orphaned resource")
	expectedEvents := []string{expectedProvisionCallEvent.String(), expectedOrphanMitigationEvent.String()}
	if err := checkEvents(events, expectedEvents); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceWithTerminalProvisionFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Error: osb.HTTPStatusCodeError{StatusCode: http.StatusBadRequest, ErrorMessage: strPtr("BadRequest"), Description: strPtr("Your parameters are incorrect!")}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceProvisionRequestFailingErrorNoOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, errorProvisionCallFailedReason, "ClusterServiceBrokerReturnedFailure", instance)
	events := getRecordedEvents(testController)
	message := fmt.Sprintf("Error provisioning ServiceInstance of ClusterServiceClass (K8S: %q ExternalName: %q) at ClusterServiceBroker %q: Status: %v; ErrorMessage: %s", "cscguid", "test-clusterserviceclass", "test-clusterservicebroker", 400, "BadRequest; Description: Your parameters are incorrect!; ResponseError: <nil>")
	expectedEvents := []string{warningEventBuilder(errorProvisionCallFailedReason).msg(message).String(), warningEventBuilder("ClusterServiceBrokerReturnedFailure").msg(message).String()}
	if err := checkEvents(events, expectedEvents); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstance(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{DashboardURL: &testDashboardURL}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	assertNumberOfBrokerActions(t, fakeClusterServiceBrokerClient.Actions(), 0)
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successProvisionReason).msg(successProvisionMessage)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceFailsWithDeletedPlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sp := getTestClusterServicePlan()
	sp.Status.RemovedFromBrokerCatalog = true
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(sp)
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("This should fail")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceReadyFalse(t, updatedServiceInstance, errorDeletedClusterServicePlanReason)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorDeletedClusterServicePlanReason).msgf("ClusterServicePlan (K8S: %q ExternalName: %q) has been deleted; cannot provision.", "cspguid", "test-clusterserviceplan")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceFailsWithDeletedClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sc := getTestClusterServiceClass()
	sc.Status.RemovedFromBrokerCatalog = true
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(sc)
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("This should have failed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceReadyFalse(t, updatedServiceInstance, errorDeletedClusterServiceClassReason)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorDeletedClusterServiceClassReason).msgf("ClusterServiceClass (K8S: %q ExternalName: %q) has been deleted; cannot provision.", "cscguid", "test-clusterserviceclass")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceSuccessWithK8SNames(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{DashboardURL: &testDashboardURL}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceK8SNames()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateReference(t, actions[0], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ClusterServiceClassRef == nil || updateObject.Spec.ClusterServiceClassRef.Name != "cscguid" {
		t.Fatalf("ClusterServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ClusterServicePlanRef == nil || updateObject.Spec.ClusterServicePlanRef.Name != "cspguid" {
		t.Fatalf("ClusterServicePlanRef was not resolved correctly during reconcile")
	}
	instance = updateObject
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions = fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance = assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successProvisionReason).msg(successProvisionMessage)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceAsynchronous(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := osb.OperationKey(testOperation)
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{Async: true, DashboardURL: &testDashboardURL, OperationKey: &key}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testOperation, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
	kubeActions := fakeKubeClient.Actions()
	if e, a := 1, len(kubeActions); e != a {
		t.Fatalf("Unexpected number of actions: expected %v, got %v", e, a)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have a record of seeing test instance once")
	}
}
func TestReconcileServiceInstanceAsynchronousNoOperation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{Async: true, DashboardURL: &testDashboardURL}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, "", testClusterServicePlanName, testClusterServicePlanGUID, instance)
	assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
	kubeActions := fakeKubeClient.Actions()
	if e, a := 1, len(kubeActions); e != a {
		t.Fatalf("Unexpected number of actions: expected %v, got %v", e, a)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have a record of seeing test instance once")
	}
}
func TestReconcileServiceInstanceNamespaceError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	fakeKubeClient.PrependReactor("get", "namespaces", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &corev1.Namespace{}, errors.New("No namespace")
	})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("There should not be a namespace for the ServiceInstance to be created in")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceErrorBeforeRequest(t, updatedServiceInstance, errorFindingNamespaceServiceInstanceReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorFindingNamespaceServiceInstanceReason).msgf("Failed to get namespace %q:", "test-ns").msg("No namespace")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceDeleteBlockedByCredentials(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	credentials := getTestServiceBinding()
	sharedInformers.ServiceBindings().Informer().GetStore().Add(credentials)
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("expected reconcileServiceInstance to return an error, but there was none")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updateObject := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceErrorBeforeRequest(t, updateObject, errorDeprovisionBlockedByCredentialsReason, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorDeprovisionBlockedByCredentialsReason).msg("All associated ServiceBindings must be removed before this ServiceInstance can be deleted")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
	sharedInformers.ServiceBindings().Informer().GetStore().Delete(credentials)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	instance = updateObject.(*v1beta1.ServiceInstance)
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions = fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
	kubeActions = fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions = fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updateObject = assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updateObject, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events = getRecordedEvents(testController)
	expectedEvent = normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceDeleteAsynchronous(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := osb.OperationKey(testOperation)
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{Async: true, OperationKey: &key}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have a record of seeing test instance once")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testOperation, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(asyncDeprovisioningReason).msg("The instance is being deprovisioned asynchronously")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceDeleteFailedProvisionWithRequest(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name			string
		currentOperation	v1beta1.ServiceInstanceOperation
		inProgressProperties	*v1beta1.ServiceInstancePropertiesState
	}{{name: "With failed provisioning operation in progress", currentOperation: v1beta1.ServiceInstanceOperationProvision, inProgressProperties: &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}}, {name: "With terminally failed provisioning", currentOperation: "", inProgressProperties: nil}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceWithFailedStatus()
			instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
			instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
			instance.Status.CurrentOperation = tc.currentOperation
			instance.Status.InProgressProperties = tc.inProgressProperties
			instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
			instance.Generation = 2
			instance.Status.ReconciledGeneration = 1
			instance.Status.ObservedGeneration = 1
			instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
			fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
				return true, instance, nil
			})
			if err := reconcileServiceInstance(t, testController, instance); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
			fakeCatalogClient.ClearActions()
			fakeKubeClient.ClearActions()
			err := reconcileServiceInstance(t, testController, instance)
			if err != nil {
				t.Fatalf("Unexpected error from reconcileServiceInstance: %v", err)
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
			kubeActions := fakeKubeClient.Actions()
			assertNumberOfActions(t, kubeActions, 0)
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
			events := getRecordedEvents(testController)
			expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
			if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestReconsileServiceInstanceDeleteWithParameters(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name			string
		externalProperties	*v1beta1.ServiceInstancePropertiesState
		deprovisionStatus	v1beta1.ServiceInstanceDeprovisionStatus
		serviceBinding		*v1beta1.ServiceBinding
		generation		int64
		reconciledGeneration	int64
	}{{name: "With a failed to provision instance and without making a provision request", externalProperties: &v1beta1.ServiceInstancePropertiesState{}, deprovisionStatus: v1beta1.ServiceInstanceDeprovisionStatusNotRequired, serviceBinding: nil, generation: 1, reconciledGeneration: 0}, {name: "With a failed to provision instance, with inactive binding, and without making a provision request", externalProperties: &v1beta1.ServiceInstancePropertiesState{}, deprovisionStatus: v1beta1.ServiceInstanceDeprovisionStatusNotRequired, serviceBinding: getTestServiceInactiveBinding(), generation: 1, reconciledGeneration: 0}, {name: "With a deprovisioned instance and without making a deprovision request", externalProperties: nil, deprovisionStatus: v1beta1.ServiceInstanceDeprovisionStatusSucceeded, serviceBinding: nil, generation: 2, reconciledGeneration: 1}, {name: "With a deprovisioned instance, with inactive binding, and without making a deprovision request", externalProperties: nil, deprovisionStatus: v1beta1.ServiceInstanceDeprovisionStatusSucceeded, serviceBinding: getTestServiceInactiveBinding(), generation: 2, reconciledGeneration: 1}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			if tc.serviceBinding != nil {
				sharedInformers.ClusterServicePlans().Informer().GetStore().Add(tc.serviceBinding)
			}
			instance := getTestServiceInstanceWithFailedStatus()
			instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
			instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
			instance.Status.ExternalProperties = tc.externalProperties
			instance.Status.DeprovisionStatus = tc.deprovisionStatus
			instance.Generation = tc.generation
			instance.Status.ReconciledGeneration = tc.reconciledGeneration
			instance.Status.ObservedGeneration = 1
			instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
			fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
				return true, instance, nil
			})
			err := reconcileServiceInstance(t, testController, instance)
			if err != nil {
				t.Fatalf("Unexpected error from reconcileServiceInstance: %v", err)
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 0)
			kubeActions := fakeKubeClient.Actions()
			assertNumberOfActions(t, kubeActions, 0)
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			assertEmptyFinalizers(t, updatedServiceInstance)
			events := getRecordedEvents(testController)
			assertNumEvents(t, events, 0)
		})
	}
}
func TestReconcileServiceInstanceDeleteWhenAlreadyDeprovisionedUnsuccessfully(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithFailedStatus()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{}
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusFailed
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("Unexpected error from reconcileServiceInstance: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 0)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestReconcileServiceInstanceDeleteFailedUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 2
	instance.Status.ObservedGeneration = 2
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("Unexpected error from reconcileServiceInstance: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
	assertEmptyFinalizers(t, updatedServiceInstance)
}
func TestReconcileServiceInstanceDeleteDoesNotInvokeClusterServiceBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 1
	instance.Status.ReconciledGeneration = 0
	instance.Status.ObservedGeneration = 0
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusNotRequired
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertEmptyFinalizers(t, updatedServiceInstance)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestFinalizerClearedWhen409ConflictEncounteredOnStatusUpdate(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ResourceVersion = "1"
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 1
	instance.Status.ReconciledGeneration = 0
	instance.Status.ObservedGeneration = 0
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusNotRequired
	newerInstance := instance.DeepCopy()
	newerInstance.ResourceVersion = "2"
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, newerInstance, nil
	})
	fakeCatalogClient.AddReactor("update", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		updateAction := action.(clientgotesting.UpdateAction)
		object := updateAction.GetObject()
		instance := object.(*v1beta1.ServiceInstance)
		if instance.ResourceVersion == "1" {
			return true, nil, apierrors.NewConflict(action.GetResource().GroupResource(), instance.Name, errors.New("object has changed"))
		}
		return false, nil, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 3)
	assertUpdateStatus(t, actions[0], instance)
	assertGet(t, actions[1], instance)
	updatedServiceInstance := assertUpdateStatus(t, actions[2], instance)
	assertEmptyFinalizers(t, updatedServiceInstance)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestReconcileServiceInstanceWithFailedCondition(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithFailedStatus()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 1)
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 1)
	expectedEvent := normalEventBuilder(successProvisionReason).msg("The instance was provisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceInProgressProvisioningWithOperation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncProvisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have record of seeing test instance once")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testOperation, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	assertServiceInstanceConditionHasLastOperationDescription(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, lastOperationDescription)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestPollServiceInstanceSuccessProvisioningWithOperation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncProvisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have requeues of test instance after polling have completed with a 'success' state")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
}
func TestPollServiceInstanceFailureProvisioningWithOperation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncProvisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) == 0 {
		t.Fatalf("Expected polling queue to have a record of test instance to process orphan mitigation")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestFailingErrorStartOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, startingInstanceOrphanMitigationReason, errorProvisionCallFailedReason, errorProvisionCallFailedReason, instance)
}
func TestPollServiceInstanceInProgressDeprovisioningWithOperationNoFinalizer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name	string
		setup	func(instance *v1beta1.ServiceInstance)
	}{{name: "nil plan", setup: func(instance *v1beta1.ServiceInstance) {
		instance.Spec.ClusterServicePlanExternalName = "plan-that-does-not-exist"
		instance.Spec.ClusterServicePlanRef = nil
	}}, {name: "With plan", setup: func(instance *v1beta1.ServiceInstance) {
	}}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
			tc.setup(instance)
			instanceKey := testNamespace + "/" + testServiceInstanceName
			if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
				t.Fatalf("Expected polling queue to not have any record of test instance")
			}
			err := testController.pollServiceInstance(instance)
			if err != nil {
				t.Fatalf("pollServiceInstance failed: %s", err)
			}
			if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
				t.Fatalf("Expected polling queue to have record of seeing test instance once")
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			operationKey := osb.OperationKey(testOperation)
			assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			assertServiceInstanceAsyncStillInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testOperation, testClusterServicePlanName, testClusterServicePlanGUID, instance)
			assertServiceInstanceConditionHasLastOperationDescription(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, lastOperationDescription)
			kubeActions := fakeKubeClient.Actions()
			assertNumberOfActions(t, kubeActions, 0)
		})
	}
}
func TestPollServiceInstanceSuccessDeprovisioningWithOperationNoFinalizer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceFailureDeprovisioning(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err == nil {
		t.Fatalf("Expected pollServiceInstance to return an error but there was none")
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestRetriableError(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, errorDeprovisionCallFailedReason, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorDeprovisionCallFailedReason).msg("Deprovision call failed: (no description provided)")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceFailureDeprovisioningWithReconciliationTimeout(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceUpdateRequestFailingErrorNoOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, errorDeprovisionCallFailedReason, errorReconciliationRetryTimeoutReason, instance)
	events := getRecordedEvents(testController)
	expectedEvents := []string{warningEventBuilder(errorDeprovisionCallFailedReason).msg("Deprovision call failed: (no description provided)").String(), warningEventBuilder(errorReconciliationRetryTimeoutReason).msg("Stopping reconciliation retries because too much time has elapsed").String()}
	if err := checkEvents(events, expectedEvents); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceStatusGoneDeprovisioningWithOperationNoFinalizer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Error: osb.HTTPStatusCodeError{StatusCode: http.StatusGone}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceClusterServiceBrokerTemporaryError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Error: osb.HTTPStatusCodeError{StatusCode: http.StatusForbidden}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err == nil {
		t.Fatal("Expected pollServiceInstance to return error")
	}
	expectedErr := "Error polling last operation: Status: 403; ErrorMessage: <nil>; Description: <nil>; ResponseError: <nil>"
	if e, a := expectedErr, err.Error(); e != a {
		t.Fatalf("unexpected error returned: expected %q, got %q", e, a)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	assertUpdateStatus(t, actions[0], instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorPollingLastOperationReason).msg("Error polling last operation:").msg("Status: 403; ErrorMessage: <nil>; Description: <nil>; ResponseError: <nil>")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceClusterServiceBrokerTerminalError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Error: osb.HTTPStatusCodeError{StatusCode: http.StatusBadRequest}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %v", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have requeues")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	assertUpdateStatus(t, actions[0], instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorPollingLastOperationReason).msg("Error polling last operation:").msg("Status: 400; ErrorMessage: <nil>; Description: <nil>; ResponseError: <nil>")
	if err := checkEvents(events, []string{expectedEvent.String(), expectedEvent.String()}); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceSuccessDeprovisioningWithOperationWithFinalizer(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioningWithFinalizer(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceSuccessOnFinalRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Status.CurrentOperation = v1beta1.ServiceInstanceOperationProvision
	instance.Status.InProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	instance.Status.ObservedGeneration = instance.Generation
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successProvisionReason).msg("The instance was provisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceUpdateInProgressPropertiesOnRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Status.CurrentOperation = v1beta1.ServiceInstanceOperationProvision
	instance.Status.InProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID, Parameters: &runtime.RawExtension{Raw: []byte(`{ "staleParameter": "value" }`)}, ParameterChecksum: "staleChecksum"}
	instance.Status.ObservedGeneration = instance.Generation
	instance.Status.Conditions = []v1beta1.ServiceInstanceCondition{{Type: v1beta1.ServiceInstanceConditionReady, Status: v1beta1.ConditionFalse, Reason: provisioningInFlightReason}}
	parameters := instanceParameters{Name: "test-param", Args: make(map[string]string)}
	parameters.Args["first"] = "first-arg"
	parameters.Args["second"] = "new-second-arg"
	b, err := json.Marshal(parameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters %v : %v", parameters, err)
	}
	instance.Spec.Parameters = &runtime.RawExtension{Raw: b}
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	if err := testController.reconcileServiceInstance(instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	expectedParameters := map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "new-second-arg"}, "name": "test-param"}
	expectedParametersChecksum := generateChecksumOfParametersOrFail(t, expectedParameters)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance).(*v1beta1.ServiceInstance)
	assertServiceInstanceOperationInProgressWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, expectedParameters, expectedParametersChecksum, instance)
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	checkEventCounts(events, []string{})
}
func TestReconcileServiceInstanceFailureOnFinalRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Error: errors.New("fake creation failure")}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Status.CurrentOperation = v1beta1.ServiceInstanceOperationProvision
	instance.Status.InProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalID: testClusterServicePlanGUID, ClusterServicePlanExternalName: testClusterServicePlanName}
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	instance.Status.ObservedGeneration = instance.Generation
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("Should have returned no error because the retry duration has elapsed: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceProvisionRequestFailingErrorNoOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, errorErrorCallingProvisionReason, errorReconciliationRetryTimeoutReason, instance)
	events := getRecordedEvents(testController)
	expectedEventPrefixes := []string{corev1.EventTypeWarning + " " + errorErrorCallingProvisionReason, corev1.EventTypeWarning + " " + errorReconciliationRetryTimeoutReason}
	if err := checkEventPrefixes(events, expectedEventPrefixes); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceSuccessOnFinalRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncProvisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	if err := testController.pollServiceInstance(instance); err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
}
func TestPollServiceInstanceFailureOnFinalRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncProvisioning(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	if err := testController.pollServiceInstance(instance); err == nil {
		t.Fatalf("Expected error to be returned in order to requeue instance for orphan mitigation")
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestFailingErrorStartOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, startingInstanceOrphanMitigationReason, errorReconciliationRetryTimeoutReason, asyncProvisioningReason, instance)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestReconcileServiceInstanceWithStatusUpdateError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, noFakeActions())
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	fakeCatalogClient.AddReactor("update", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, nil, errors.New("update error")
	})
	err := reconcileServiceInstance(t, testController, instance)
	if err == nil {
		t.Fatalf("expected error from but got none")
	}
	if e, a := "update error", err.Error(); e != a {
		t.Fatalf("unexpected error returned: expected %q, got %q", e, a)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestSetServiceInstanceCondition(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceWithCondition := func(condition *v1beta1.ServiceInstanceCondition) *v1beta1.ServiceInstance {
		instance := getTestServiceInstance()
		instance.Status.Conditions = []v1beta1.ServiceInstanceCondition{*condition}
		return instance
	}
	newTs := metav1.Now()
	oldTs := metav1.NewTime(newTs.Add(-5 * time.Minute))
	condition := func(cType v1beta1.ServiceInstanceConditionType, status v1beta1.ConditionStatus, s ...string) *v1beta1.ServiceInstanceCondition {
		c := &v1beta1.ServiceInstanceCondition{Type: cType, Status: status}
		if len(s) > 0 {
			c.Reason = s[0]
		}
		if len(s) > 1 {
			c.Message = s[1]
		}
		c.LastTransitionTime = oldTs
		return c
	}
	readyFalse := func() *v1beta1.ServiceInstanceCondition {
		return condition(v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, "Reason", "Message")
	}
	readyFalsef := func(reason, message string) *v1beta1.ServiceInstanceCondition {
		return condition(v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, reason, message)
	}
	readyTrue := func() *v1beta1.ServiceInstanceCondition {
		return condition(v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionTrue, "Reason", "Message")
	}
	failedTrue := func() *v1beta1.ServiceInstanceCondition {
		return condition(v1beta1.ServiceInstanceConditionFailed, v1beta1.ConditionTrue, "Reason", "Message")
	}
	withNewTs := func(c *v1beta1.ServiceInstanceCondition) *v1beta1.ServiceInstanceCondition {
		c.LastTransitionTime = newTs
		return c
	}
	cases := []struct {
		name		string
		input		*v1beta1.ServiceInstance
		condition	*v1beta1.ServiceInstanceCondition
		result		*v1beta1.ServiceInstance
	}{{name: "new ready condition", input: getTestServiceInstance(), condition: readyFalse(), result: instanceWithCondition(withNewTs(readyFalse()))}, {name: "not ready -> not ready; no ts update", input: instanceWithCondition(readyFalse()), condition: readyFalse(), result: instanceWithCondition(readyFalse())}, {name: "not ready -> not ready, reason and message change; no ts update", input: instanceWithCondition(readyFalse()), condition: readyFalsef("DifferentReason", "DifferentMessage"), result: instanceWithCondition(readyFalsef("DifferentReason", "DifferentMessage"))}, {name: "not ready -> ready", input: instanceWithCondition(readyFalse()), condition: readyTrue(), result: instanceWithCondition(withNewTs(readyTrue()))}, {name: "ready -> ready; no ts update", input: instanceWithCondition(readyTrue()), condition: readyTrue(), result: instanceWithCondition(readyTrue())}, {name: "ready -> not ready", input: instanceWithCondition(readyTrue()), condition: readyFalse(), result: instanceWithCondition(withNewTs(readyFalse()))}, {name: "not ready -> not ready + failed", input: instanceWithCondition(readyFalse()), condition: failedTrue(), result: func() *v1beta1.ServiceInstance {
		i := instanceWithCondition(readyFalse())
		i.Status.Conditions = append(i.Status.Conditions, *withNewTs(failedTrue()))
		return i
	}()}}
	for _, tc := range cases {
		setServiceInstanceConditionInternal(tc.input, tc.condition.Type, tc.condition.Status, tc.condition.Reason, tc.condition.Message, newTs)
		if !reflect.DeepEqual(tc.input, tc.result) {
			t.Errorf("%v: unexpected diff: %v", tc.name, diff.ObjectReflectDiff(tc.input, tc.result))
		}
	}
}
func TestUpdateServiceInstanceCondition(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	getTestServiceInstanceWithStatus := func(status v1beta1.ConditionStatus) *v1beta1.ServiceInstance {
		instance := getTestServiceInstance()
		instance.Status = v1beta1.ServiceInstanceStatus{Conditions: []v1beta1.ServiceInstanceCondition{{Type: v1beta1.ServiceInstanceConditionReady, Status: status, Message: "message", LastTransitionTime: metav1.NewTime(time.Now().Add(-5 * time.Minute))}}}
		return instance
	}
	cases := []struct {
		name			string
		input			*v1beta1.ServiceInstance
		status			v1beta1.ConditionStatus
		reason			string
		message			string
		transitionTimeChanged	bool
	}{{name: "initially unset", input: getTestServiceInstance(), status: v1beta1.ConditionFalse, message: "message", transitionTimeChanged: true}, {name: "not ready -> not ready", input: getTestServiceInstanceWithStatus(v1beta1.ConditionFalse), status: v1beta1.ConditionFalse, transitionTimeChanged: false}, {name: "not ready -> not ready, reason and message change", input: getTestServiceInstanceWithStatus(v1beta1.ConditionFalse), status: v1beta1.ConditionFalse, reason: "foo", message: "bar", transitionTimeChanged: false}, {name: "not ready -> ready", input: getTestServiceInstanceWithStatus(v1beta1.ConditionFalse), status: v1beta1.ConditionTrue, message: "message", transitionTimeChanged: true}, {name: "ready -> ready", input: getTestServiceInstanceWithStatus(v1beta1.ConditionTrue), status: v1beta1.ConditionTrue, message: "message", transitionTimeChanged: false}, {name: "ready -> not ready", input: getTestServiceInstanceWithStatus(v1beta1.ConditionTrue), status: v1beta1.ConditionFalse, message: "message", transitionTimeChanged: true}, {name: "message -> message2", input: getTestServiceInstanceWithStatus(v1beta1.ConditionTrue), status: v1beta1.ConditionFalse, message: "message2", transitionTimeChanged: true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, noFakeActions())
			inputClone := tc.input.DeepCopy()
			_, err := testController.updateServiceInstanceCondition(tc.input, v1beta1.ServiceInstanceConditionReady, tc.status, tc.reason, tc.message)
			if err != nil {
				t.Fatalf("%v: error updating instance condition: %v", tc.name, err)
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 0)
			if !reflect.DeepEqual(tc.input, inputClone) {
				t.Fatalf("%v: updating broker condition mutated input: %s", tc.name, expectedGot(inputClone, tc.input))
			}
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], tc.input)
			updateActionObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
			if !ok {
				t.Fatalf("%v: couldn't convert to instance", tc.name)
			}
			var initialTs metav1.Time
			if len(inputClone.Status.Conditions) != 0 {
				initialTs = inputClone.Status.Conditions[0].LastTransitionTime
			}
			if e, a := 1, len(updateActionObject.Status.Conditions); e != a {
				t.Fatalf("%v: condition(s) %s", tc.name, expectedGot(e, a))
			}
			outputCondition := updateActionObject.Status.Conditions[0]
			newTs := outputCondition.LastTransitionTime
			if tc.transitionTimeChanged && initialTs == newTs {
				t.Fatalf("%v: transition time didn't change when it should have", tc.name)
			} else if !tc.transitionTimeChanged && initialTs != newTs {
				t.Fatalf("%v: transition time changed when it shouldn't have", tc.name)
			}
			if e, a := tc.reason, outputCondition.Reason; e != "" && e != a {
				t.Fatalf("%v: condition reasons didn't match; %s", tc.name, expectedGot(e, a))
			}
			if e, a := tc.message, outputCondition.Message; e != "" && e != a {
				t.Fatalf("%v: condition reasons didn't match; %s", tc.name, expectedGot(e, a))
			}
		})
	}
}
func TestReconcileInstanceUsingOriginatingIdentity(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, tc := range originatingIdentityTestCases {
		func() {
			prevOrigIDEnablement := sctestutil.EnableOriginatingIdentity(t, tc.enableOriginatingIdentity)
			defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=%v", scfeatures.OriginatingIdentity, prevOrigIDEnablement))
			fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{DashboardURL: &testDashboardURL}}})
			addGetNamespaceReaction(fakeKubeClient)
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceWithClusterRefs()
			if tc.includeUserInfo {
				instance.Spec.UserInfo = testUserInfo
			}
			if err := reconcileServiceInstance(t, testController, instance); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			instance = updatedServiceInstance.(*v1beta1.ServiceInstance)
			if err := reconcileServiceInstance(t, testController, instance); err != nil {
				t.Fatalf("This should not fail : %v", err)
			}
			brokerActions := fakeBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			actualRequest, ok := brokerActions[0].Request.(*osb.ProvisionRequest)
			if !ok {
				t.Errorf("%v: unexpected request type; expected %T, got %T", tc.name, &osb.ProvisionRequest{}, actualRequest)
				return
			}
			var expectedOriginatingIdentity *osb.OriginatingIdentity
			if tc.expectedOriginatingIdentity {
				expectedOriginatingIdentity = testOriginatingIdentity
			}
			assertOriginatingIdentity(t, expectedOriginatingIdentity, actualRequest.OriginatingIdentity)
		}()
	}
}
func TestReconcileInstanceDeleteUsingOriginatingIdentity(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, tc := range originatingIdentityTestCases {
		func() {
			prevOrigIDEnablement := sctestutil.EnableOriginatingIdentity(t, tc.enableOriginatingIdentity)
			defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=%v", scfeatures.OriginatingIdentity, prevOrigIDEnablement))
			_, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceWithClusterRefs()
			instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
			instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
			instance.Generation = 2
			instance.Status.ReconciledGeneration = 1
			instance.Status.ObservedGeneration = 1
			instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
			instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
			instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
			if tc.includeUserInfo {
				instance.Spec.UserInfo = testUserInfo
			}
			fakeCatalogClient.AddReactor("get", "instances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
				return true, instance, nil
			})
			if err := reconcileServiceInstance(t, testController, instance); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
			fakeCatalogClient.ClearActions()
			err := reconcileServiceInstance(t, testController, instance)
			if err != nil {
				t.Fatalf("This should not fail")
			}
			brokerActions := fakeBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			actualRequest, ok := brokerActions[0].Request.(*osb.DeprovisionRequest)
			if !ok {
				t.Errorf("%v: unexpected request type; expected %T, got %T", tc.name, &osb.DeprovisionRequest{}, actualRequest)
				return
			}
			var expectedOriginatingIdentity *osb.OriginatingIdentity
			if tc.expectedOriginatingIdentity {
				expectedOriginatingIdentity = testOriginatingIdentity
			}
			assertOriginatingIdentity(t, expectedOriginatingIdentity, actualRequest.OriginatingIdentity)
		}()
	}
}
func TestPollInstanceUsingOriginatingIdentity(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, tc := range originatingIdentityTestCases {
		func() {
			prevOrigIDEnablement := sctestutil.EnableOriginatingIdentity(t, tc.enableOriginatingIdentity)
			defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=%v", scfeatures.OriginatingIdentity, prevOrigIDEnablement))
			_, _, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceAsyncProvisioning(testOperation)
			if tc.includeUserInfo {
				instance.Spec.UserInfo = testUserInfo
			}
			err := testController.pollServiceInstance(instance)
			if err != nil {
				t.Fatalf("Expected pollServiceInstance to not fail while in progress")
			}
			brokerActions := fakeBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			actualRequest, ok := brokerActions[0].Request.(*osb.LastOperationRequest)
			if !ok {
				t.Errorf("%v: unexpected request type; expected %T, got %T", tc.name, &osb.LastOperationRequest{}, actualRequest)
				return
			}
			var expectedOriginatingIdentity *osb.OriginatingIdentity
			if tc.expectedOriginatingIdentity {
				expectedOriginatingIdentity = testOriginatingIdentity
			}
			assertOriginatingIdentity(t, expectedOriginatingIdentity, actualRequest.OriginatingIdentity)
		}()
	}
}
func TestReconcileServiceInstanceWithHTTPStatusCodeErrorOrphanMitigation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name				string
		statusCode			int
		triggersOrphanMitigation	bool
		terminalFailure			bool
	}{{name: "Status OK", statusCode: 200, triggersOrphanMitigation: false}, {name: "other 2XX", statusCode: 201, triggersOrphanMitigation: true}, {name: "3XX", statusCode: 300, triggersOrphanMitigation: false}, {name: "400", statusCode: 400, triggersOrphanMitigation: false, terminalFailure: true}, {name: "408", statusCode: 408, triggersOrphanMitigation: false}, {name: "other 4XX", statusCode: 400, triggersOrphanMitigation: false}, {name: "5XX", statusCode: 500, triggersOrphanMitigation: true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Error: osb.HTTPStatusCodeError{StatusCode: tc.statusCode}}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceWithClusterRefs()
			if err := reconcileServiceInstance(t, testController, instance); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
			fakeCatalogClient.ClearActions()
			err := reconcileServiceInstance(t, testController, instance)
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedObject := assertUpdateStatus(t, actions[0], instance)
			updatedServiceInstance, _ := updatedObject.(*v1beta1.ServiceInstance)
			assertServiceInstanceOrphanMitigationInProgress(t, updatedServiceInstance, tc.triggersOrphanMitigation)
			if tc.triggersOrphanMitigation {
				assertServiceInstanceStartingOrphanMitigation(t, updatedServiceInstance, instance)
				if err == nil {
					t.Fatalf("%v: Reconciler should return error so that instance is orphan mitigated", tc.name)
				}
			} else {
				if err != nil {
					if tc.terminalFailure {
						t.Fatalf("%v: Reconciler should treat as terminal condition and not requeue", tc.name)
					}
				}
			}
		})
	}
}
func TestReconcileServiceInstanceTimeoutTriggersOrphanMitigation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Error: &url.Error{Err: getTestTimeoutError()}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatal("Reconciler should return error for timeout so that instance is orphan mitigated")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedObject := assertUpdateStatus(t, actions[0], instance)
	updatedServiceInstance, ok := updatedObject.(*v1beta1.ServiceInstance)
	if !ok {
		fatalf(t, "Couldn't convert object %+v into a *v1beta1.ServiceInstance", updatedObject)
	}
	assertServiceInstanceReadyCondition(t, updatedServiceInstance, v1beta1.ConditionFalse, startingInstanceOrphanMitigationReason)
	assertServiceInstanceOrphanMitigationTrue(t, updatedServiceInstance, errorErrorCallingProvisionReason)
	assertServiceInstanceOrphanMitigationInProgressTrue(t, updatedServiceInstance)
}
func TestReconcileServiceInstanceOrphanMitigation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := osb.OperationKey(testOperation)
	description := "description"
	cases := []struct {
		name				string
		deprovReaction			*fakeosb.DeprovisionReaction
		pollReaction			*fakeosb.PollLastOperationReaction
		async				bool
		finishedOrphanMitigation	bool
		shouldError			bool
		retryDurationExceeded		bool
		expectedReadyConditionStatus	v1beta1.ConditionStatus
		expectedReadyConditionReason	string
	}{{name: "sync - success", deprovReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}, finishedOrphanMitigation: true, expectedReadyConditionStatus: v1beta1.ConditionFalse, expectedReadyConditionReason: successOrphanMitigationReason}, {name: "sync - 202 accepted", deprovReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{Async: true, OperationKey: &key}}, finishedOrphanMitigation: false, expectedReadyConditionStatus: v1beta1.ConditionFalse, expectedReadyConditionReason: asyncDeprovisioningReason}, {name: "sync - http error", deprovReaction: &fakeosb.DeprovisionReaction{Error: fakeosb.AsyncRequiredError()}, finishedOrphanMitigation: false, shouldError: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorDeprovisionCallFailedReason}, {name: "sync - http error - retry duration exceeded", deprovReaction: &fakeosb.DeprovisionReaction{Error: fakeosb.AsyncRequiredError()}, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}, {name: "sync - other error", deprovReaction: &fakeosb.DeprovisionReaction{Error: fmt.Errorf("other error")}, finishedOrphanMitigation: false, shouldError: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorDeprovisionCallFailedReason}, {name: "sync - other error - retry duration exceeded", deprovReaction: &fakeosb.DeprovisionReaction{Error: fmt.Errorf("other error")}, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}, {name: "poll - success", pollReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded}}, async: true, finishedOrphanMitigation: true, expectedReadyConditionStatus: v1beta1.ConditionFalse, expectedReadyConditionReason: successOrphanMitigationReason}, {name: "poll - gone", pollReaction: &fakeosb.PollLastOperationReaction{Error: osb.HTTPStatusCodeError{StatusCode: http.StatusGone}}, async: true, finishedOrphanMitigation: true, expectedReadyConditionStatus: v1beta1.ConditionFalse, expectedReadyConditionReason: successOrphanMitigationReason}, {name: "poll - in progress", pollReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: &description}}, async: true, finishedOrphanMitigation: false, expectedReadyConditionStatus: v1beta1.ConditionFalse, expectedReadyConditionReason: asyncDeprovisioningReason}, {name: "poll - failed", pollReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}, async: true, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}, {name: "poll - failed - retry duration exceeded", pollReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}, async: true, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}, {name: "poll - error - retry duration exceeded", pollReaction: &fakeosb.PollLastOperationReaction{Error: fmt.Errorf("other error")}, async: true, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}, {name: "poll - in progress - retry duration exceeded", pollReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress}}, async: true, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}, {name: "poll - invalid state - retry duration exceeded", pollReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: "invalid state"}}, async: true, finishedOrphanMitigation: false, retryDurationExceeded: true, expectedReadyConditionStatus: v1beta1.ConditionUnknown, expectedReadyConditionReason: errorOrphanMitigationFailedReason}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: tc.deprovReaction, PollLastOperationReaction: tc.pollReaction})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceWithClusterRefs()
			instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
			instance.Status.CurrentOperation = v1beta1.ServiceInstanceOperationProvision
			instance.Status.OrphanMitigationInProgress = true
			setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionOrphanMitigation, v1beta1.ConditionTrue, startingInstanceOrphanMitigationReason, startingInstanceOrphanMitigationMessage)
			instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
			instance.Status.InProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
			if tc.async {
				instance.Status.AsyncOpInProgress = true
			}
			var startTime metav1.Time
			if tc.retryDurationExceeded {
				startTime = metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
			} else {
				startTime = metav1.NewTime(time.Now())
			}
			instance.Status.OperationStartTime = &startTime
			fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
				return true, instance, nil
			})
			err := reconcileServiceInstance(t, testController, instance)
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedObject := assertUpdateStatus(t, actions[0], instance)
			updatedServiceInstance, _ := updatedObject.(*v1beta1.ServiceInstance)
			assertServiceInstanceOrphanMitigationInProgress(t, updatedServiceInstance, !tc.finishedOrphanMitigation)
			if tc.finishedOrphanMitigation {
				assertServiceInstanceOrphanMitigationMissing(t, updatedServiceInstance)
			} else {
				assertServiceInstanceOrphanMitigationTrue(t, updatedServiceInstance, startingInstanceOrphanMitigationReason)
			}
			assertServiceInstanceReadyCondition(t, updatedServiceInstance, tc.expectedReadyConditionStatus, tc.expectedReadyConditionReason)
			if tc.shouldError {
				if err == nil {
					t.Fatalf("%v: Expected error; this should not be a terminal state", tc.name)
				}
			} else {
				if err != nil {
					t.Fatalf("%v: Unexpected error; this should be a terminal state", tc.name)
				}
			}
			assertCatalogFinalizerExists(t, updatedServiceInstance)
		})
	}
}
func TestReconcileServiceInstanceWithSecretParameters(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{}}})
	paramSecret := &corev1.Secret{Data: map[string][]byte{"param-secret-key": []byte("{\"b\":\"2\"}")}}
	addGetSecretReaction(fakeKubeClient, paramSecret)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	parameters := map[string]interface{}{"a": "1"}
	b, err := json.Marshal(parameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters %v : %v", parameters, err)
	}
	instance.Spec.Parameters = &runtime.RawExtension{Raw: b}
	instance.Spec.ParametersFrom = []v1beta1.ParametersFromSource{{SecretKeyRef: &v1beta1.SecretKeyReference{Name: "param-secret-name", Key: "param-secret-key"}}}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedParameters := map[string]interface{}{"a": "1", "b": "<redacted>"}
	expectedParametersChecksum := generateChecksumOfParametersOrFail(t, map[string]interface{}{"a": "1", "b": "2"})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationInProgressWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, expectedParameters, expectedParametersChecksum, instance)
	instance = updatedServiceInstance.(*v1beta1.ServiceInstance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err = reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext, Parameters: map[string]interface{}{"a": "1", "b": "2"}})
	actions = fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance = assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccessWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testClusterServicePlanName, testClusterServicePlanGUID, expectedParameters, expectedParametersChecksum, instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if len(updateObject.Spec.Parameters.Raw) == 0 {
		t.Fatalf("Parameters was unexpectedly empty")
	}
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}, {verb: "get", resourceName: "secrets", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successProvisionReason).msg("The instance was provisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestResolveReferencesReferencesAlreadySet(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, _ := newTestController(t, noFakeActions())
	instance := getTestServiceInstanceWithClusterRefs()
	modified, err := testController.resolveReferences(instance)
	if err != nil {
		t.Fatalf("resolveReferences failed unexpectedly: %q", err)
	}
	if modified {
		t.Fatalf("Should have returned false")
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 0)
}
func TestResolveReferencesNoClusterServiceClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, _ := newTestController(t, noFakeActions())
	instance := getTestServiceInstance()
	modified, err := testController.resolveReferences(instance)
	if err == nil {
		t.Fatalf("Should have failed with no service class")
	}
	if e, a := "a non-existent ClusterServiceClass", err.Error(); !strings.Contains(a, e) {
		t.Fatalf("Did not get the expected error message %q got %q", e, a)
	}
	if !modified {
		t.Fatalf("Should have returned true")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.externalName", instance.Spec.ClusterServiceClassExternalName)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	updatedServiceInstance := assertUpdateStatus(t, actions[1], instance)
	updatedObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updatedObject.Spec.ClusterServiceClassRef != nil {
		t.Fatalf("ClusterServiceClassRef was unexpectedly set: %+v", updatedObject)
	}
	if updatedObject.Spec.ClusterServicePlanRef != nil {
		t.Fatalf("ClusterServicePlanRef was unexpectedly set: %+v", updatedObject)
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServiceClassReason).msg(fmt.Sprintf(`References a non-existent ClusterServiceClass %c or there is more than one (found: 0)`, instance.Spec.PlanReference))
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceUpdateParameters(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	oldParameters := map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}, "name": "test-param"}
	oldParametersMarshaled, err := MarshalRawParameters(oldParameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters: %v", err)
	}
	oldParametersRaw := &runtime.RawExtension{Raw: oldParametersMarshaled}
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID, Parameters: oldParametersRaw, ParameterChecksum: generateChecksumOfParametersOrFail(t, oldParameters)}
	parameters := instanceParameters{Name: "test-param", Args: make(map[string]string)}
	parameters.Args["first"] = "first-arg"
	parameters.Args["second"] = "new-second-arg"
	b, err := json.Marshal(parameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters %v : %v", parameters, err)
	}
	instance.Spec.Parameters = &runtime.RawExtension{Raw: b}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectedParameters := map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "new-second-arg"}, "name": "test-param"}
	expectedParametersChecksum := generateChecksumOfParametersOrFail(t, expectedParameters)
	instance = assertServiceInstanceOperationInProgressWithParametersIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, expectedParameters, expectedParametersChecksum)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err = reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: nil, Context: testContext, Parameters: map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "new-second-arg"}, "name": "test-param"}})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccessWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, expectedParameters, expectedParametersChecksum, instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if len(updateObject.Spec.Parameters.Raw) == 0 {
		t.Fatalf("Parameters was unexpectedly empty")
	}
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successUpdateInstanceReason).msg("The instance was updated successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceDeleteParameters(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	oldParameters := map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}, "name": "test-param"}
	oldParametersMarshaled, err := MarshalRawParameters(oldParameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters: %v", err)
	}
	oldParametersRaw := &runtime.RawExtension{Raw: oldParametersMarshaled}
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID, Parameters: oldParametersRaw, ParameterChecksum: generateChecksumOfParametersOrFail(t, oldParameters)}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceUpdateInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err = reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: nil, Context: testContext, Parameters: make(map[string]interface{})})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.Parameters != nil {
		t.Fatalf("Parameters was unexpectedly not empty")
	}
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successUpdateInstanceReason).msg("The instance was updated successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestResolveReferencesNoClusterServicePlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, _ := newTestController(t, noFakeActions())
	instance := getTestServiceInstance()
	sc := getTestClusterServiceClass()
	var scItems []v1beta1.ClusterServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: scItems}, nil
	})
	modified, err := testController.resolveReferences(instance)
	if err == nil {
		t.Fatalf("Should have failed with no service plan")
	}
	if e, a := "a non-existent ClusterServicePlan", err.Error(); !strings.Contains(a, e) {
		t.Fatalf("Did not get the expected error message %q got %q", e, a)
	}
	if !modified {
		t.Fatalf("Should have returned true")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 3)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.externalName", instance.Spec.ClusterServiceClassExternalName)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	listRestrictions = clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=test-clusterserviceplan,spec.clusterServiceBrokerName=test-clusterservicebroker,spec.clusterServiceClassRef.name=cscguid")}
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateStatus(t, actions[2], instance)
	updatedObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updatedObject.Spec.ClusterServiceClassRef == nil || updatedObject.Spec.ClusterServiceClassRef.Name != testClusterServiceClassGUID {
		t.Fatalf("ClusterServiceClassRef.Name was not set correctly, expected %q got: %+v", testClusterServiceClassGUID, updatedObject.Spec.ClusterServiceClassRef.Name)
	}
	if updatedObject.Spec.ClusterServicePlanRef != nil {
		t.Fatalf("ClusterServicePlanRef was unexpectedly set: %+v", updatedObject)
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorNonexistentClusterServicePlanReason).msgf(`References a non-existent ClusterServicePlan %b on ClusterServiceClass %s %c or there is more than one (found: 0)`, instance.Spec.PlanReference, instance.Spec.ClusterServiceClassRef.Name, instance.Spec.PlanReference)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceUpdateDashboardURLResponse(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name				string
		enableUpdateDashboardURL	bool
		newDashboardURL			string
	}{{name: "new dashboard url returned and alpha feature enabled", enableUpdateDashboardURL: true, newDashboardURL: "http://foobar.com"}, {name: "dashboard url blank not returned and alpha feature enabled", enableUpdateDashboardURL: true, newDashboardURL: ""}, {name: "new dashboard url returned and alpha feature disabled", enableUpdateDashboardURL: false, newDashboardURL: "http://banana.com"}}
	for _, tc := range cases {
		fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{DashboardURL: &tc.newDashboardURL}}})
		if tc.enableUpdateDashboardURL {
			err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.UpdateDashboardURL))
			if err != nil {
				t.Fatalf("Failed to enable updatable dashboard URL feature: %v", err)
			}
		} else {
			err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.UpdateDashboardURL))
			if err != nil {
				t.Fatalf("Failed to enable updatable dashboard URL feature: %v", err)
			}
		}
		defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.UpdateDashboardURL))
		sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
		sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
		sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
		instance := getTestServiceInstanceWithClusterRefs()
		instance.Generation = 2
		instance.Status.ReconciledGeneration = 1
		instance.Status.ObservedGeneration = 1
		instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
		instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
		instance.Status.DashboardURL = &testDashboardURL
		oldParameters := map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}, "name": "test-param"}
		oldParametersMarshaled, err := MarshalRawParameters(oldParameters)
		if err != nil {
			t.Fatalf("Failed to marshal parameters: %v", err)
		}
		oldParametersRaw := &runtime.RawExtension{Raw: oldParametersMarshaled}
		oldParametersChecksum := generateChecksumOfParametersOrFail(t, oldParameters)
		instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: "old-plan-name", ClusterServicePlanExternalID: "old-plan-id", Parameters: oldParametersRaw, ParameterChecksum: oldParametersChecksum}
		parameters := instanceParameters{Name: "test-param", Args: make(map[string]string)}
		parameters.Args["first"] = "first-arg"
		parameters.Args["second"] = "second-arg"
		b, err := json.Marshal(parameters)
		if err != nil {
			t.Fatalf("Failed to marshal parameters %v : %v", parameters, err)
		}
		instance.Spec.Parameters = &runtime.RawExtension{Raw: b}
		if err := testController.reconcileServiceInstance(instance); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		instance = assertServiceInstanceOperationInProgressWithParametersIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, oldParameters, oldParametersChecksum)
		fakeCatalogClient.ClearActions()
		fakeKubeClient.ClearActions()
		if err = testController.reconcileServiceInstance(instance); err != nil {
			t.Fatalf("This should not fail : %v", err)
		}
		brokerActions := fakeClusterServiceBrokerClient.Actions()
		assertNumberOfBrokerActions(t, brokerActions, 1)
		expectedPlanID := testClusterServicePlanGUID
		assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: &expectedPlanID, Context: testContext, Parameters: nil})
		actions := fakeCatalogClient.Actions()
		assertNumberOfActions(t, actions, 1)
		updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
		if tc.enableUpdateDashboardURL {
			if tc.newDashboardURL != "" {
				assertServiceInstanceDashboardURL(t, updatedServiceInstance, tc.newDashboardURL)
			} else {
				assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
			}
		} else {
			assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
		}
		kubeActions := fakeKubeClient.Actions()
		if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
			t.Fatal(err)
		}
		events := getRecordedEvents(testController)
		expectedEvent := normalEventBuilder(successUpdateInstanceReason).msg("The instance was updated successfully")
		if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
			t.Fatal(err)
		}
	}
}
func TestReconcileServiceInstanceUpdatePlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	oldParameters := map[string]interface{}{"args": map[string]interface{}{"first": "first-arg", "second": "second-arg"}, "name": "test-param"}
	oldParametersMarshaled, err := MarshalRawParameters(oldParameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters: %v", err)
	}
	oldParametersRaw := &runtime.RawExtension{Raw: oldParametersMarshaled}
	oldParametersChecksum := generateChecksumOfParametersOrFail(t, oldParameters)
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: "old-plan-name", ClusterServicePlanExternalID: "old-plan-id", Parameters: oldParametersRaw, ParameterChecksum: oldParametersChecksum}
	parameters := instanceParameters{Name: "test-param", Args: make(map[string]string)}
	parameters.Args["first"] = "first-arg"
	parameters.Args["second"] = "second-arg"
	b, err := json.Marshal(parameters)
	if err != nil {
		t.Fatalf("Failed to marshal parameters %v : %v", parameters, err)
	}
	instance.Spec.Parameters = &runtime.RawExtension{Raw: b}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceOperationInProgressWithParametersIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, oldParameters, oldParametersChecksum)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err = reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	expectedPlanID := testClusterServicePlanGUID
	assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: &expectedPlanID, Context: testContext, Parameters: nil})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccessWithParameters(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, oldParameters, oldParametersChecksum, instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if len(updateObject.Spec.Parameters.Raw) == 0 {
		t.Fatalf("Parameters was unexpectedly empty")
	}
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successUpdateInstanceReason).msg("The instance was updated successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceWithUpdateCallFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Error: errors.New("fake update failure")}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceUpdatingPlan()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceUpdateInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err == nil {
		t.Fatalf("Should not be able to make the ServiceInstance.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	expectedPlanID := testClusterServicePlanGUID
	assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: &expectedPlanID, Context: testContext})
	kubeActions := fakeKubeClient.Actions()
	if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
		t.Fatal(err)
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestRetriableError(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, errorErrorCallingUpdateInstanceReason, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorErrorCallingUpdateInstanceReason).msg("The update call failed and will be retried:").msg("Error communicating with broker for updating:").msg("fake update failure")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceWithUpdateFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name			string
		brokerHTTPError		osb.HTTPStatusCodeError
		errorExpected		bool
		expectedFailureReason	string
		expectedEventMessage	string
	}{{name: "retriable failure", brokerHTTPError: osb.HTTPStatusCodeError{StatusCode: http.StatusConflict, ErrorMessage: strPtr("OutOfQuota"), Description: strPtr("You're out of quota!")}, errorExpected: true, expectedFailureReason: "", expectedEventMessage: "ServiceBroker returned a failure for update call; update will be retried: " + "Status: 409; ErrorMessage: OutOfQuota; Description: You're out of quota!; ResponseError: <nil>"}, {name: "terminal failure", brokerHTTPError: osb.HTTPStatusCodeError{StatusCode: http.StatusBadRequest, ErrorMessage: strPtr("BadRequest"), Description: strPtr("Something's wrong with the request")}, errorExpected: false, expectedFailureReason: errorUpdateInstanceCallFailedReason, expectedEventMessage: "ServiceBroker returned a failure for update call; update will not be retried: " + "Status: 400; ErrorMessage: BadRequest; Description: Something's wrong with the request; ResponseError: <nil>"}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Error: tc.brokerHTTPError}})
			sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
			sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
			sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
			instance := getTestServiceInstanceUpdatingPlan()
			if err := reconcileServiceInstance(t, testController, instance); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			instance = assertServiceInstanceUpdateInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
			fakeCatalogClient.ClearActions()
			fakeKubeClient.ClearActions()
			err := reconcileServiceInstance(t, testController, instance)
			if tc.errorExpected && err == nil {
				t.Fatal("expected error to be returned")
			} else if !tc.errorExpected && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			expectedPlanID := testClusterServicePlanGUID
			assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: &expectedPlanID, Context: testContext})
			kubeActions := fakeKubeClient.Actions()
			if err := checkKubeClientActions(kubeActions, []kubeClientAction{{verb: "get", resourceName: "namespaces", checkType: checkGetActionType}}); err != nil {
				t.Fatal(err)
			}
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			assertServiceInstanceUpdateRequestFailingErrorNoOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, errorUpdateInstanceCallFailedReason, tc.expectedFailureReason, instance)
			events := getRecordedEvents(testController)
			expectedEvent := warningEventBuilder(errorUpdateInstanceCallFailedReason).msg(tc.expectedEventMessage)
			if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestResolveReferencesWorks(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, _ := newTestController(t, noFakeActions())
	instance := getTestServiceInstance()
	sc := getTestClusterServiceClass()
	var scItems []v1beta1.ClusterServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: scItems}, nil
	})
	sp := getTestClusterServicePlan()
	var spItems []v1beta1.ClusterServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: spItems}, nil
	})
	modified, err := testController.resolveReferences(instance)
	if err != nil {
		t.Fatalf("Should not have failed, but failed with: %q", err)
	}
	if !modified {
		t.Fatalf("Should have returned true")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 3)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.externalName", instance.Spec.ClusterServiceClassExternalName)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	listRestrictions = clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=test-clusterserviceplan,spec.clusterServiceBrokerName=test-clusterservicebroker,spec.clusterServiceClassRef.name=cscguid")}
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateReference(t, actions[2], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ClusterServiceClassRef == nil || updateObject.Spec.ClusterServiceClassRef.Name != testClusterServiceClassGUID {
		t.Fatalf("ClusterServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ClusterServicePlanRef == nil || updateObject.Spec.ClusterServicePlanRef.Name != testClusterServicePlanGUID {
		t.Fatalf("ClusterServicePlanRef was not resolved correctly during reconcile")
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestResolveReferencesForPlanChange(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	instance := getTestServiceInstanceWithClusterRefs()
	newPlanID := "new-plan-id"
	newPlanName := "new-plan-name"
	sp := &v1beta1.ClusterServicePlan{ObjectMeta: metav1.ObjectMeta{Name: newPlanID}, Spec: v1beta1.ClusterServicePlanSpec{CommonServicePlanSpec: v1beta1.CommonServicePlanSpec{ExternalID: newPlanID, ExternalName: newPlanName, Bindable: truePtr()}}}
	var spItems []v1beta1.ClusterServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: spItems}, nil
	})
	instance.Spec.ClusterServicePlanExternalName = newPlanName
	instance.Spec.ClusterServicePlanRef = nil
	modified, err := testController.resolveReferences(instance)
	if err != nil {
		t.Fatalf("Should not have failed, but failed with: %q", err)
	}
	if !modified {
		t.Fatalf("Should have returned true")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=new-plan-name,spec.clusterServiceBrokerName=test-clusterservicebroker,spec.clusterServiceClassRef.name=cscguid")}
	assertList(t, actions[0], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateReference(t, actions[1], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ClusterServiceClassRef == nil || updateObject.Spec.ClusterServiceClassRef.Name != testClusterServiceClassGUID {
		t.Fatalf("ClusterServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ClusterServicePlanRef == nil || updateObject.Spec.ClusterServicePlanRef.Name != newPlanID {
		t.Fatalf("ClusterServicePlanRef was not resolved correctly during reconcile")
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestResolveReferencesWorksK8SNames(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, noFakeActions())
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceK8SNames()
	modified, err := testController.resolveReferences(instance)
	if err != nil {
		t.Fatalf("Should not have failed, but failed with: %q", err)
	}
	if !modified {
		t.Fatalf("Should have returned true")
	}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateReference(t, actions[0], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ClusterServiceClassRef == nil || updateObject.Spec.ClusterServiceClassRef.Name != testClusterServiceClassGUID {
		t.Fatalf("ClusterServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ClusterServicePlanRef == nil || updateObject.Spec.ClusterServicePlanRef.Name != testClusterServicePlanGUID {
		t.Fatalf("ClusterServicePlanRef was not resolved correctly during reconcile")
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
func TestReconcileServiceInstanceUpdateAsynchronous(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key := osb.OperationKey(testOperation)
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{Async: true, OperationKey: &key}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: "old-plan-name", ClusterServicePlanExternalID: "old-plan-id"}
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceUpdateInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	expectedPlanID := testClusterServicePlanGUID
	assertUpdateInstance(t, brokerActions[0], &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: &expectedPlanID, Context: testContext})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, testOperation, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	kubeActions := fakeKubeClient.Actions()
	if e, a := 1, len(kubeActions); e != a {
		t.Fatalf("Unexpected number of actions: expected %v, got %v", e, a)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have a record of seeing test instance once")
	}
}
func TestPollServiceInstanceAsyncInProgressUpdating(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncUpdating(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have record of seeing test instance once")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStillInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, testOperation, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	assertServiceInstanceConditionHasLastOperationDescription(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, lastOperationDescription)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestPollServiceInstanceAsyncSuccessUpdating(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncUpdating(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, testClusterServicePlanName, testClusterServicePlanGUID, instance)
}
func TestPollServiceInstanceAsyncFailureUpdating(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceAsyncUpdating(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err := testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue not to have a record of test instance as update should not have retried")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testClusterServiceClassGUID), PlanID: strPtr(testClusterServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceUpdateRequestFailingErrorNoOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationUpdate, errorUpdateInstanceCallFailedReason, errorUpdateInstanceCallFailedReason, instance)
}
func TestCheckClassAndPlanForDeletion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		instance	*v1beta1.ServiceInstance
		class		*v1beta1.ClusterServiceClass
		plan		*v1beta1.ClusterServicePlan
		success		bool
		expectedReason	string
		expectedErrors	[]string
	}{{name: "non-deleted plan and class works", instance: getTestServiceInstance(), class: getTestClusterServiceClass(), plan: getTestClusterServicePlan(), success: true}, {name: "deleted plan fails", instance: getTestServiceInstance(), class: getTestClusterServiceClass(), plan: getTestMarkedAsRemovedClusterServicePlan(), success: false, expectedReason: errorDeletedClusterServicePlanReason, expectedErrors: []string{"ClusterServicePlan", "has been deleted"}}, {name: "deleted class fails", instance: getTestServiceInstance(), class: getTestMarkedAsRemovedClusterServiceClass(), plan: getTestClusterServicePlan(), success: false, expectedReason: errorDeletedClusterServiceClassReason, expectedErrors: []string{"ClusterServiceClass", "has been deleted"}}, {name: "deleted plan and class fails", instance: getTestServiceInstance(), class: getTestClusterServiceClass(), plan: getTestMarkedAsRemovedClusterServicePlan(), success: false, expectedReason: errorDeletedClusterServicePlanReason, expectedErrors: []string{"ClusterServicePlan", "has been deleted"}}, {name: "Updating plan fails", instance: getTestServiceInstanceUpdatingPlan(), class: getTestClusterServiceClass(), plan: getTestMarkedAsRemovedClusterServicePlan(), success: false, expectedReason: errorDeletedClusterServicePlanReason, expectedErrors: []string{"ClusterServicePlan", "has been deleted"}}, {name: "Updating parameters works", instance: getTestServiceInstanceUpdatingParametersOfDeletedPlan(), class: getTestClusterServiceClass(), plan: getTestMarkedAsRemovedClusterServicePlan(), success: true}}
	for _, tc := range cases {
		fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, noFakeActions())
		err := testController.checkForRemovedClusterClassAndPlan(tc.instance, tc.class, tc.plan)
		if err != nil {
			if tc.success {
				t.Errorf("%q: Unexpected error %v", tc.name, err)
			}
			for _, exp := range tc.expectedErrors {
				if e, a := exp, err.Error(); !strings.Contains(a, e) {
					t.Errorf("%q: Did not find expected error %q : got %q", tc.name, e, a)
				}
			}
		} else if !tc.success {
			t.Errorf("%q: Did not get a failure when expected one", tc.name)
		}
		assertNumberOfActions(t, fakeKubeClient.Actions(), 0)
		brokerActions := fakeClusterServiceBrokerClient.Actions()
		assertNumberOfBrokerActions(t, brokerActions, 0)
		actions := fakeCatalogClient.Actions()
		assertNumberOfActions(t, actions, 0)
	}
}
func TestReconcileServiceInstanceDeleteDuringOngoingOperation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Status.CurrentOperation = v1beta1.ServiceInstanceOperationProvision
	startTime := metav1.NewTime(time.Now().Add(-1 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	instance.Status.InProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	timeOfReconciliation := metav1.Now()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	if instance.Status.OperationStartTime.Before(&timeOfReconciliation) {
		t.Fatalf("OperationStartTime should not be before the time that the reconciliation started. OperationStartTime=%v. timeOfReconciliation=%v", instance.Status.OperationStartTime, timeOfReconciliation)
	}
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance).(*v1beta1.ServiceInstance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceDeleteWithOngoingOperation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Status.CurrentOperation = v1beta1.ServiceInstanceOperationProvision
	startTime := metav1.NewTime(time.Now().Add(-1 * time.Hour))
	instance.Status.OperationStartTime = &startTime
	instance.Status.OrphanMitigationInProgress = true
	instance.Status.InProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionOrphanMitigation, v1beta1.ConditionTrue, startingInstanceOrphanMitigationReason, startingInstanceOrphanMitigationMessage)
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	timeOfReconciliation := metav1.Now()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	if instance.Status.OperationStartTime.Before(&timeOfReconciliation) {
		t.Fatalf("OperationStartTime should not be before the time that the reconciliation started. OperationStartTime=%v. timeOfReconciliation=%v", instance.Status.OperationStartTime, timeOfReconciliation)
	}
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: testClusterServicePlanGUID})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance).(*v1beta1.ServiceInstance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testClusterServicePlanName, testClusterServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 1)
	expectedEvent := corev1.EventTypeNormal + " " + successDeprovisionReason + " " + "The instance was deprovisioned successfully"
	if e, a := expectedEvent, events[0]; e != a {
		t.Fatalf("Received unexpected event: %v\nExpected: %v", a, e)
	}
}
func TestReconcileServiceInstanceDeleteWithNonExistentPlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: "old-plan-name", ClusterServicePlanExternalID: "old-plan-id"}
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	instance.Spec.ClusterServicePlanRef = nil
	fakeCatalogClient.AddReactor("get", "serviceinstances", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, instance, nil
	})
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceOperationInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationDeprovision, "old-plan-name", "old-plan-id")
	fakeCatalogClient.ClearActions()
	fakeKubeClient.ClearActions()
	err := reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testClusterServiceClassGUID, PlanID: "old-plan-id"})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, "old-plan-name", "old-plan-id", instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceUpdateMissingObservedGeneration(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 0
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance).(*v1beta1.ServiceInstance)
	if updatedServiceInstance.Status.ObservedGeneration == 0 || updatedServiceInstance.Status.ObservedGeneration != instance.Status.ReconciledGeneration {
		t.Fatalf("Unexpected ObservedGeneration value: %d", updatedServiceInstance.Status.ObservedGeneration)
	}
	if updatedServiceInstance.Status.ProvisionStatus != v1beta1.ServiceInstanceProvisionStatusProvisioned {
		t.Fatalf("The instance was expected to be marked as Provisioned")
	}
}
func TestReconcileServiceInstanceUpdateMissingOrphanMitigation(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{UpdateInstanceReaction: &fakeosb.UpdateInstanceReaction{Response: &osb.UpdateInstanceResponse{}}})
	sharedInformers.ClusterServiceBrokers().Informer().GetStore().Add(getTestClusterServiceBroker())
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(getTestClusterServiceClass())
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(getTestClusterServicePlan())
	instance := getTestServiceInstanceWithClusterRefs()
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	instance.Status.OrphanMitigationInProgress = true
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: testClusterServicePlanName, ClusterServicePlanExternalID: testClusterServicePlanGUID}
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance).(*v1beta1.ServiceInstance)
	if !isServiceInstanceOrphanMitigation(updatedServiceInstance) {
		t.Fatal("Expected instance status to have an OrphanMitigation condition set to True")
	}
}
func generateChecksumOfParametersOrFail(t *testing.T, params map[string]interface{}) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	expectedParametersChecksum, err := generateChecksumOfParameters(params)
	if err != nil {
		t.Fatalf("Failed to generate parameters checksum: %v", err)
	}
	return expectedParametersChecksum
}
func assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t *testing.T, fakeCatalogClient *fake.Clientset, instance *v1beta1.ServiceInstance) *v1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var planName, planGUID string
	if instance.Spec.ClusterServiceClassSpecified() {
		planName = testClusterServicePlanName
		planGUID = testClusterServicePlanGUID
	} else {
		planName = testServicePlanName
		planGUID = testServicePlanGUID
	}
	return assertServiceInstanceOperationInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationProvision, planName, planGUID)
}
func assertServiceInstanceUpdateInProgressIsTheOnlyCatalogClientAction(t *testing.T, fakeCatalogClient *fake.Clientset, instance *v1beta1.ServiceInstance) *v1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var planName, planGUID string
	if instance.Spec.ClusterServiceClassSpecified() {
		planName = testClusterServicePlanName
		planGUID = testClusterServicePlanGUID
	} else {
		planName = testServicePlanName
		planGUID = testServicePlanGUID
	}
	return assertServiceInstanceOperationInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationUpdate, planName, planGUID)
}
func assertServiceInstanceDeprovisionInProgressIsTheOnlyCatalogClientAction(t *testing.T, fakeCatalogClient *fake.Clientset, instance *v1beta1.ServiceInstance) *v1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var planName, planGUID string
	if instance.Spec.ClusterServiceClassSpecified() {
		planName = testClusterServicePlanName
		planGUID = testClusterServicePlanGUID
	} else {
		planName = testServicePlanName
		planGUID = testServicePlanGUID
	}
	return assertServiceInstanceOperationInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, v1beta1.ServiceInstanceOperationDeprovision, planName, planGUID)
}
func assertServiceInstanceOperationInProgressIsTheOnlyCatalogClientAction(t *testing.T, fakeCatalogClient *fake.Clientset, instance *v1beta1.ServiceInstance, operation v1beta1.ServiceInstanceOperation, planName string, planGUID string) *v1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return assertServiceInstanceOperationInProgressWithParametersIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance, operation, planName, planGUID, nil, "")
}
func assertServiceInstanceOperationInProgressWithParametersIsTheOnlyCatalogClientAction(t *testing.T, fakeCatalogClient *fake.Clientset, instance *v1beta1.ServiceInstance, operation v1beta1.ServiceInstanceOperation, planName string, planGUID string, parameters map[string]interface{}, parametersChecksum string) *v1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updateObject := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationInProgressWithParameters(t, updateObject, operation, planName, planGUID, parameters, parametersChecksum, instance)
	return updateObject.(*v1beta1.ServiceInstance)
}
func reconcileServiceInstance(t *testing.T, testController *controller, instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clone := instance.DeepCopy()
	err := testController.reconcileServiceInstance(instance)
	if !reflect.DeepEqual(instance, clone) {
		t.Errorf("reconcileServiceInstance shouldn't mutate input, but it does: %s", expectedGot(clone, instance))
	}
	return err
}
