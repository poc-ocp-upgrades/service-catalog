package controller

import (
	"fmt"
	"testing"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	fakeosb "github.com/pmorie/go-open-service-broker-client/v2/fake"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	clientgotesting "k8s.io/client-go/testing"
)

func TestReconcileServiceInstanceNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{DashboardURL: &testDashboardURL}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceWithNamespacedRefs()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	instance = assertServiceInstanceProvisionInProgressIsTheOnlyCatalogClientAction(t, fakeCatalogClient, instance)
	fakeCatalogClient.ClearActions()
	assertNumberOfBrokerActions(t, fakeBrokerClient.Actions(), 0)
	fakeKubeClient.ClearActions()
	if err := reconcileServiceInstance(t, testController, instance); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testServiceClassGUID, PlanID: testServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
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
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testServicePlanName, testServicePlanGUID, instance)
	assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successProvisionReason).msg(successProvisionMessage)
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileServiceInstanceAsynchronousNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	key := osb.OperationKey(testOperation)
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{ProvisionReaction: &fakeosb.ProvisionReaction{Response: &osb.ProvisionResponse{Async: true, DashboardURL: &testDashboardURL, OperationKey: &key}}})
	addGetNamespaceReaction(fakeKubeClient)
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceWithNamespacedRefs()
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
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertProvision(t, brokerActions[0], &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testServiceClassGUID, PlanID: testServicePlanGUID, OrganizationGUID: testClusterID, SpaceGUID: testNamespaceGUID, Context: testContext})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testOperation, testServicePlanName, testServicePlanGUID, instance)
	assertServiceInstanceDashboardURL(t, updatedServiceInstance, testDashboardURL)
	kubeActions := fakeKubeClient.Actions()
	if e, a := 1, len(kubeActions); e != a {
		t.Fatalf("Unexpected number of actions: expected %v, got %v", e, a)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have a record of seeing test instance once")
	}
}
func TestPollServiceInstanceInProgressProvisioningWithOperationNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceAsyncProvisioningWithNamespacedRefs(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err = testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have record of seeing test instance once")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testServiceClassGUID), PlanID: strPtr(testServicePlanGUID), OperationKey: &operationKey})
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testOperation, testServicePlanName, testServicePlanGUID, instance)
	assertServiceInstanceConditionHasLastOperationDescription(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, lastOperationDescription)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestPollServiceInstanceSuccessProvisioningWithOperationNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded, Description: strPtr(lastOperationDescription)}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceAsyncProvisioningWithNamespacedRefs(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err = testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have requeues of test instance after polling have completed with a 'success' state")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testServiceClassGUID), PlanID: strPtr(testServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, testServicePlanName, testServicePlanGUID, instance)
}
func TestPollServiceInstanceFailureProvisioningWithOperationNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceAsyncProvisioningWithNamespacedRefs(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err = testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) == 0 {
		t.Fatalf("Expected polling queue to have a record of test instance to process orphan mitigation")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testServiceClassGUID), PlanID: strPtr(testServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestFailingErrorStartOrphanMitigation(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationProvision, startingInstanceOrphanMitigationReason, errorProvisionCallFailedReason, errorProvisionCallFailedReason, instance)
}
func TestReconcileServiceInstanceDeleteWithNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceWithNamespacedRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ServicePlanExternalName: testServicePlanName, ServicePlanExternalID: testServicePlanGUID}
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
	err = reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testServiceClassGUID, PlanID: testServicePlanGUID})
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
func TestReconcileServiceInstanceDeleteAsynchronousWithNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	key := osb.OperationKey(testOperation)
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{DeprovisionReaction: &fakeosb.DeprovisionReaction{Response: &osb.DeprovisionResponse{Async: true, OperationKey: &key}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceWithNamespacedRefs()
	instance.ObjectMeta.DeletionTimestamp = &metav1.Time{}
	instance.ObjectMeta.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
	instance.Generation = 2
	instance.Status.ReconciledGeneration = 1
	instance.Status.ObservedGeneration = 1
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ExternalProperties = &v1beta1.ServiceInstancePropertiesState{ServicePlanExternalName: testServicePlanName, ServicePlanExternalID: testServicePlanGUID}
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
	err = reconcileServiceInstance(t, testController, instance)
	if err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
		t.Fatalf("Expected polling queue to have a record of seeing test instance once")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertDeprovision(t, brokerActions[0], &osb.DeprovisionRequest{AcceptsIncomplete: true, InstanceID: testServiceInstanceGUID, ServiceID: testServiceClassGUID, PlanID: testServicePlanGUID})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceAsyncStartInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testOperation, testServicePlanName, testServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(asyncDeprovisioningReason).msg("The instance is being deprovisioned asynchronously")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceInProgressDeprovisioningWithOperationNoFinalizerNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	cases := []struct {
		name	string
		setup	func(instance *v1beta1.ServiceInstance)
	}{{name: "nil plan", setup: func(instance *v1beta1.ServiceInstance) {
		instance.Spec.ServicePlanExternalName = "plan-that-does-not-exist"
		instance.Spec.ServicePlanRef = nil
	}}, {name: "With plan", setup: func(instance *v1beta1.ServiceInstance) {
	}}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateInProgress, Description: strPtr(lastOperationDescription)}}})
			sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
			sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
			sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
			instance := getTestServiceInstanceAsyncDeprovisioningWithNamespacedRefs(testOperation)
			tc.setup(instance)
			instanceKey := testNamespace + "/" + testServiceInstanceName
			if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
				t.Fatalf("Expected polling queue to not have any record of test instance")
			}
			err = testController.pollServiceInstance(instance)
			if err != nil {
				t.Fatalf("pollServiceInstance failed: %s", err)
			}
			if testController.instancePollingQueue.NumRequeues(instanceKey) != 1 {
				t.Fatalf("Expected polling queue to have record of seeing test instance once")
			}
			brokerActions := fakeBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 1)
			operationKey := osb.OperationKey(testOperation)
			assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testServiceClassGUID), PlanID: strPtr(testServicePlanGUID), OperationKey: &operationKey})
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
			assertServiceInstanceAsyncStillInProgress(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testOperation, testServicePlanName, testServicePlanGUID, instance)
			assertServiceInstanceConditionHasLastOperationDescription(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, lastOperationDescription)
			kubeActions := fakeKubeClient.Actions()
			assertNumberOfActions(t, kubeActions, 0)
		})
	}
}
func TestPollServiceInstanceSuccessDeprovisioningWithOperationNoFinalizerNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateSucceeded}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioningWithNamespacedRefs(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err = testController.pollServiceInstance(instance)
	if err != nil {
		t.Fatalf("pollServiceInstance failed: %s", err)
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testServiceClassGUID), PlanID: strPtr(testServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceOperationSuccess(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, testServicePlanName, testServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := normalEventBuilder(successDeprovisionReason).msg("The instance was deprovisioned successfully")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestPollServiceInstanceFailureDeprovisioningNamespacedRefs(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", scfeatures.NamespacedServiceBroker))
	if err != nil {
		t.Fatalf("Could not enable NamespacedServiceBroker feature flag.")
	}
	defer utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=false", scfeatures.NamespacedServiceBroker))
	fakeKubeClient, fakeCatalogClient, fakeBrokerClient, testController, sharedInformers := newTestController(t, fakeosb.FakeClientConfiguration{PollLastOperationReaction: &fakeosb.PollLastOperationReaction{Response: &osb.LastOperationResponse{State: osb.StateFailed}}})
	sharedInformers.ServiceBrokers().Informer().GetStore().Add(getTestServiceBroker())
	sharedInformers.ServiceClasses().Informer().GetStore().Add(getTestServiceClass())
	sharedInformers.ServicePlans().Informer().GetStore().Add(getTestServicePlan())
	instance := getTestServiceInstanceAsyncDeprovisioningWithNamespacedRefs(testOperation)
	instanceKey := testNamespace + "/" + testServiceInstanceName
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance")
	}
	err = testController.pollServiceInstance(instance)
	if err == nil {
		t.Fatalf("Expected pollServiceInstance to return an error but there was none")
	}
	if testController.instancePollingQueue.NumRequeues(instanceKey) != 0 {
		t.Fatalf("Expected polling queue to not have any record of test instance as polling should have completed")
	}
	brokerActions := fakeBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	operationKey := osb.OperationKey(testOperation)
	assertPollLastOperation(t, brokerActions[0], &osb.LastOperationRequest{InstanceID: testServiceInstanceGUID, ServiceID: strPtr(testServiceClassGUID), PlanID: strPtr(testServicePlanGUID), OperationKey: &operationKey})
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 1)
	updatedServiceInstance := assertUpdateStatus(t, actions[0], instance)
	assertServiceInstanceRequestRetriableError(t, updatedServiceInstance, v1beta1.ServiceInstanceOperationDeprovision, errorDeprovisionCallFailedReason, testServicePlanName, testServicePlanGUID, instance)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorDeprovisionCallFailedReason).msg("Deprovision call failed: (no description provided)")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestResolveNamespacedReferencesWorks(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, _, testController, _ := newTestController(t, noFakeActions())
	instance := getTestServiceInstanceWithNamespacedPlanReference()
	sc := getTestServiceClass()
	var scItems []v1beta1.ServiceClass
	scItems = append(scItems, *sc)
	fakeCatalogClient.AddReactor("list", "serviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ServiceClassList{Items: scItems}, nil
	})
	sp := getTestServicePlan()
	var spItems []v1beta1.ServicePlan
	spItems = append(spItems, *sp)
	fakeCatalogClient.AddReactor("list", "serviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ServicePlanList{Items: spItems}, nil
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
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.externalName", instance.Spec.ServiceClassExternalName)}
	assertList(t, actions[0], &v1beta1.ServiceClass{}, listRestrictions)
	listRestrictions = clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.ParseSelectorOrDie("spec.externalName=test-serviceplan,spec.serviceBrokerName=test-servicebroker,spec.serviceClassRef.name=scguid")}
	assertList(t, actions[1], &v1beta1.ServicePlan{}, listRestrictions)
	updatedServiceInstance := assertUpdateReference(t, actions[2], instance)
	updateObject, ok := updatedServiceInstance.(*v1beta1.ServiceInstance)
	if !ok {
		t.Fatalf("couldn't convert to *v1beta1.ServiceInstance")
	}
	if updateObject.Spec.ServiceClassRef == nil || updateObject.Spec.ServiceClassRef.Name != testServiceClassGUID {
		t.Fatalf("ServiceClassRef was not resolved correctly during reconcile")
	}
	if updateObject.Spec.ServicePlanRef == nil || updateObject.Spec.ServicePlanRef.Name != testServicePlanGUID {
		t.Fatalf("ServicePlanRef was not resolved correctly during reconcile")
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 0)
}
