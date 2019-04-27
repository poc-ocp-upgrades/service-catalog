package controller

import (
	"errors"
	"reflect"
	"testing"
	"time"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	fakeosb "github.com/pmorie/go-open-service-broker-client/v2/fake"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/test/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	"strings"
	corev1 "k8s.io/api/core/v1"
	clientgotesting "k8s.io/client-go/testing"
)

func TestShouldReconcileClusterServiceBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		broker		*v1beta1.ClusterServiceBroker
		now		time.Time
		reconcile	bool
		err		error
	}{{name: "no status", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBroker()
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Minute}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "deletionTimestamp set", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue)
		broker.DeletionTimestamp = &metav1.Time{}
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Hour}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "no ready condition", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBroker()
		broker.Status = v1beta1.ClusterServiceBrokerStatus{CommonServiceBrokerStatus: v1beta1.CommonServiceBrokerStatus{Conditions: []v1beta1.ServiceBrokerCondition{{Type: v1beta1.ServiceBrokerConditionType("NotARealCondition"), Status: v1beta1.ConditionTrue}}}}
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Minute}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "not ready", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBrokerWithStatus(v1beta1.ConditionFalse)
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Minute}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "ready, interval elapsed", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue)
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Minute}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "good steady state - ready, interval not elapsed, but last state change was a long time ago", broker: func() *v1beta1.ClusterServiceBroker {
		lastTransitionTime := metav1.NewTime(time.Now().Add(-30 * time.Minute))
		lastRelistTime := metav1.NewTime(time.Now().Add(-2 * time.Minute))
		broker := getTestClusterServiceBrokerWithStatusAndTime(v1beta1.ConditionTrue, lastTransitionTime, lastRelistTime)
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Minute}
		return broker
	}(), now: time.Now(), reconcile: false}, {name: "good steady state - ready, interval has elapsed, last state change was a long time ago", broker: func() *v1beta1.ClusterServiceBroker {
		lastTransitionTime := metav1.NewTime(time.Now().Add(-30 * time.Minute))
		lastRelistTime := metav1.NewTime(time.Now().Add(-4 * time.Minute))
		broker := getTestClusterServiceBrokerWithStatusAndTime(v1beta1.ConditionTrue, lastTransitionTime, lastRelistTime)
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Minute}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "ready, interval not elapsed", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue)
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Hour}
		return broker
	}(), now: time.Now(), reconcile: false}, {name: "ready, interval not elapsed, spec changed", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue)
		broker.Generation = 2
		broker.Status.ReconciledGeneration = 1
		broker.Spec.RelistDuration = &metav1.Duration{Duration: 3 * time.Hour}
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "ready, duration behavior, nil duration, interval not elapsed", broker: func() *v1beta1.ClusterServiceBroker {
		t := metav1.NewTime(time.Now().Add(-23 * time.Hour))
		broker := getTestClusterServiceBrokerWithStatusAndTime(v1beta1.ConditionTrue, t, t)
		broker.Spec.RelistBehavior = v1beta1.ServiceBrokerRelistBehaviorDuration
		broker.Spec.RelistDuration = nil
		return broker
	}(), now: time.Now(), reconcile: false}, {name: "ready, duration behavior, nil duration, interval elapsed", broker: func() *v1beta1.ClusterServiceBroker {
		t := metav1.NewTime(time.Now().Add(-25 * time.Hour))
		broker := getTestClusterServiceBrokerWithStatusAndTime(v1beta1.ConditionTrue, t, t)
		broker.Spec.RelistBehavior = v1beta1.ServiceBrokerRelistBehaviorDuration
		broker.Spec.RelistDuration = nil
		return broker
	}(), now: time.Now(), reconcile: true}, {name: "ready, manual behavior", broker: func() *v1beta1.ClusterServiceBroker {
		broker := getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue)
		broker.Spec.RelistBehavior = v1beta1.ServiceBrokerRelistBehaviorManual
		return broker
	}(), now: time.Now(), reconcile: false}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var ltt *time.Time
			if len(tc.broker.Status.Conditions) != 0 {
				ltt = &tc.broker.Status.Conditions[0].LastTransitionTime.Time
			}
			if tc.broker.Spec.RelistDuration != nil {
				interval := tc.broker.Spec.RelistDuration.Duration
				lastRelistTime := tc.broker.Status.LastCatalogRetrievalTime
				t.Logf("now: %v, interval: %v, last transition time: %v, last relist time: %v", tc.now, interval, ltt, lastRelistTime)
			} else {
				t.Logf("broker.Spec.RelistDuration set to nil")
			}
			actual := shouldReconcileClusterServiceBroker(tc.broker, tc.now, 24*time.Hour)
			if e, a := tc.reconcile, actual; e != a {
				t.Errorf("unexpected result: %s", expectedGot(e, a))
			}
		})
	}
}
func TestReconcileClusterServiceBrokerExistingServiceClassAndServicePlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServicePlanNonbindable := getTestClusterServicePlanNonbindable()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass}}, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", "test-clusterservicebroker")}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 6)
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertUpdate(t, actions[2], testClusterServiceClass)
	assertCreate(t, actions[3], testClusterServicePlan)
	assertCreate(t, actions[4], testClusterServicePlanNonbindable)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[5], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestReconcileClusterServiceBrokerRemovedClusterServiceClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testRemovedClusterServiceClass := getTestRemovedClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServicePlanNonbindable := getTestClusterServicePlanNonbindable()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testRemovedClusterServiceClass)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass, *testRemovedClusterServiceClass}}, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", "test-clusterservicebroker")}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 7)
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertUpdate(t, actions[2], testClusterServiceClass)
	assertUpdateStatus(t, actions[3], testRemovedClusterServiceClass)
	assertCreate(t, actions[4], testClusterServicePlan)
	assertCreate(t, actions[5], testClusterServicePlanNonbindable)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[6], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestReconcileClusterServiceBrokerRemovedAndRestoredClusterServiceClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServicePlan.Status.RemovedFromBrokerCatalog = true
	testClusterServicePlanNonbindable := getTestClusterServicePlanNonbindable()
	testClusterServiceClass.Status.RemovedFromBrokerCatalog = true
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(testClusterServicePlan)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass}}, nil
	})
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: []v1beta1.ClusterServicePlan{*testClusterServicePlan}}, nil
	})
	fakeCatalogClient.AddReactor("update", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, testClusterServiceClass, nil
	})
	fakeCatalogClient.AddReactor("update", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, testClusterServicePlan, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", "test-clusterservicebroker")}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 8)
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertUpdate(t, actions[2], testClusterServiceClass)
	class := assertUpdateStatus(t, actions[3], testClusterServiceClass)
	assertClassRemovedFromBrokerCatalogFalse(t, class)
	assertUpdate(t, actions[4], testClusterServicePlan)
	plan := assertUpdateStatus(t, actions[5], testClusterServicePlan)
	assertPlanRemovedFromBrokerCatalogFalse(t, plan)
	assertCreate(t, actions[6], testClusterServicePlanNonbindable)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[7], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestReconcileClusterServiceBrokerRemovedClusterServicePlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServicePlanNonbindable := getTestClusterServicePlanNonbindable()
	testRemovedClusterServicePlan := getTestRemovedClusterServicePlan()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(testRemovedClusterServicePlan)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass}}, nil
	})
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: []v1beta1.ClusterServicePlan{*testRemovedClusterServicePlan}}, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", "test-clusterservicebroker")}
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 7)
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertUpdate(t, actions[2], testClusterServiceClass)
	assertCreate(t, actions[3], testClusterServicePlan)
	assertCreate(t, actions[4], testClusterServicePlanNonbindable)
	assertUpdateStatus(t, actions[5], testRemovedClusterServicePlan)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[6], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestReconcileClusterServiceBrokerExistingClusterServiceClassDifferentBroker(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServiceClass.Spec.ClusterServiceBrokerName = "notTheSame"
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err == nil {
		t.Fatal("The same service class should not belong to two different brokers.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 3)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", "test-clusterservicebroker")}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[2], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorSyncingCatalogReason).msgf("Error reconciling ClusterServiceClass (K8S: %q ExternalName: %q) (broker %q):", testClusterServiceClassGUID, testClusterServiceClassName, testClusterServiceBrokerName).msgf("ClusterServiceClass (K8S: %q ExternalName: %q) already exists for Broker %q", testClusterServiceClassGUID, testClusterServiceClassName, "notTheSame")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileClusterServiceBrokerExistingClusterServicePlanDifferentClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServicePlan.Spec.ClusterServiceBrokerName = "notTheSame"
	testClusterServicePlan.Spec.ClusterServiceClassRef = v1beta1.ClusterObjectReference{Name: "notTheSameClass"}
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(testClusterServicePlan)
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err == nil {
		t.Fatal("The same service class should not belong to two different brokers.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 4)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", "test-clusterservicebroker")}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertCreate(t, actions[2], getTestClusterServiceClass())
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[3], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorSyncingCatalogReason).msgf("Error reconciling ClusterServicePlan (K8S: %q ExternalName: %q):", testClusterServicePlanGUID, testClusterServicePlanName).msgf("ClusterServicePlan (K8S: %q ExternalName: %q) already exists for Broker %q", testClusterServicePlanGUID, testClusterServicePlanName, "notTheSame")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func getClusterServiceBrokerReactor(broker *v1beta1.ClusterServiceBroker) (string, string, clientgotesting.ReactionFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "get", "clusterservicebrokers", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, broker, nil
	}
}
func listClusterServiceClassesReactor(classes []v1beta1.ClusterServiceClass) (string, string, clientgotesting.ReactionFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: classes}, nil
	}
}
func listClusterServicePlansReactor(plans []v1beta1.ClusterServicePlan) (string, string, clientgotesting.ReactionFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: plans}, nil
	}
}
func TestReconcileClusterServiceBrokerDelete(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		authInfo	*v1beta1.ClusterServiceBrokerAuthInfo
		secret		*corev1.Secret
	}{{name: "no auth", authInfo: nil, secret: nil}, {name: "basic auth", authInfo: getTestClusterBrokerBasicAuthInfo(), secret: getTestBasicAuthSecret()}, {name: "bearer auth", authInfo: getTestClusterBrokerBearerAuthInfo(), secret: getTestBearerAuthSecret()}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, getTestCatalogConfig())
			testClusterServiceClass := getTestClusterServiceClass()
			testClusterServicePlan := getTestClusterServicePlan()
			addGetSecretReaction(fakeKubeClient, tc.secret)
			broker := getTestClusterServiceBrokerWithAuth(tc.authInfo)
			broker.DeletionTimestamp = &metav1.Time{}
			broker.Finalizers = []string{v1beta1.FinalizerServiceCatalog}
			updateBrokerClientCalled := false
			testController.brokerClientManager = NewBrokerClientManager(func(_ *osb.ClientConfiguration) (osb.Client, error) {
				updateBrokerClientCalled = true
				return nil, nil
			})
			fakeCatalogClient.AddReactor(getClusterServiceBrokerReactor(broker))
			fakeCatalogClient.AddReactor(listClusterServiceClassesReactor([]v1beta1.ClusterServiceClass{*testClusterServiceClass}))
			fakeCatalogClient.AddReactor(listClusterServicePlansReactor([]v1beta1.ClusterServicePlan{*testClusterServicePlan}))
			err := reconcileClusterServiceBroker(t, testController, broker)
			if err != nil {
				t.Fatalf("This should not fail : %v", err)
			}
			if updateBrokerClientCalled {
				t.Errorf("Unexpected broker client update action")
			}
			brokerActions := fakeClusterServiceBrokerClient.Actions()
			assertNumberOfBrokerActions(t, brokerActions, 0)
			kubeActions := fakeKubeClient.Actions()
			assertNumberOfActions(t, kubeActions, 0)
			catalogActions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, catalogActions, 7)
			listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", broker.Name)}
			assertList(t, catalogActions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
			assertList(t, catalogActions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
			assertDelete(t, catalogActions[2], testClusterServicePlan)
			assertDelete(t, catalogActions[3], testClusterServiceClass)
			updatedClusterServiceBroker := assertUpdateStatus(t, catalogActions[4], broker)
			assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
			assertGet(t, catalogActions[5], broker)
			updatedClusterServiceBroker = assertUpdateStatus(t, catalogActions[6], broker)
			assertEmptyFinalizers(t, updatedClusterServiceBroker)
			events := getRecordedEvents(testController)
			expectedEvent := normalEventBuilder(successClusterServiceBrokerDeletedReason).msg("The broker test-clusterservicebroker was deleted successfully.")
			if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
				t.Fatal(err)
			}
		})
	}
}
func TestReconcileClusterServiceBrokerErrorFetchingCatalog(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, fakeosb.FakeClientConfiguration{CatalogReaction: &fakeosb.CatalogReaction{Error: errors.New("ooops")}})
	broker := getTestClusterServiceBroker()
	if err := reconcileClusterServiceBroker(t, testController, broker); err == nil {
		t.Fatal("Should have failed to get the catalog.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[0], broker)
	assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
	updatedClusterServiceBroker = assertUpdateStatus(t, actions[1], broker)
	assertClusterServiceBrokerOperationStartTimeSet(t, updatedClusterServiceBroker, true)
	assertNumberOfActions(t, fakeKubeClient.Actions(), 0)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorFetchingCatalogReason).msg("Error getting broker catalog:").msg("ooops")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileClusterServiceBrokerZeroServices(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, fakeosb.FakeClientConfiguration{CatalogReaction: &fakeosb.CatalogReaction{Response: &osb.CatalogResponse{}}})
	broker := getTestClusterServiceBroker()
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{}}, nil
	})
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: []v1beta1.ClusterServicePlan{}}, nil
	})
	err := reconcileClusterServiceBroker(t, testController, broker)
	if err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 3)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", broker.Name)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[2], broker)
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	events := getRecordedEvents(testController)
	expectedEvent := corev1.EventTypeNormal + " " + successFetchedCatalogReason + " " + successFetchedCatalogMessage
	if e, a := expectedEvent, events[0]; !strings.HasPrefix(a, e) {
		t.Fatalf("Received unexpected event, %s", expectedGot(e, a))
	}
}
func TestReconcileClusterServiceBrokerWithAuth(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		authInfo	*v1beta1.ClusterServiceBrokerAuthInfo
		secret		*corev1.Secret
		shouldSucceed	bool
	}{{name: "basic auth - normal", authInfo: getTestClusterBrokerBasicAuthInfo(), secret: getTestBasicAuthSecret(), shouldSucceed: true}, {name: "basic auth - invalid secret", authInfo: getTestClusterBrokerBasicAuthInfo(), secret: getTestBearerAuthSecret(), shouldSucceed: false}, {name: "basic auth - secret not found", authInfo: getTestClusterBrokerBasicAuthInfo(), secret: nil, shouldSucceed: false}, {name: "bearer auth - normal", authInfo: getTestClusterBrokerBearerAuthInfo(), secret: getTestBearerAuthSecret(), shouldSucceed: true}, {name: "bearer auth - invalid secret", authInfo: getTestClusterBrokerBearerAuthInfo(), secret: getTestBasicAuthSecret(), shouldSucceed: false}, {name: "bearer auth - secret not found", authInfo: getTestClusterBrokerBearerAuthInfo(), secret: nil, shouldSucceed: false}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testReconcileClusterServiceBrokerWithAuth(t, tc.authInfo, tc.secret, tc.shouldSucceed)
		})
	}
}
func testReconcileClusterServiceBrokerWithAuth(t *testing.T, authInfo *v1beta1.ClusterServiceBrokerAuthInfo, secret *corev1.Secret, shouldSucceed bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, fakeosb.FakeClientConfiguration{})
	broker := getTestClusterServiceBrokerWithAuth(authInfo)
	if secret != nil {
		addGetSecretReaction(fakeKubeClient, secret)
	} else {
		addGetSecretNotFoundReaction(fakeKubeClient)
	}
	testClusterServiceClass := getTestClusterServiceClass()
	fakeClusterServiceBrokerClient.CatalogReaction = &fakeosb.CatalogReaction{Response: &osb.CatalogResponse{Services: []osb.Service{{ID: testClusterServiceClass.Spec.ExternalID, Name: testClusterServiceClass.Name}}}}
	err := reconcileClusterServiceBroker(t, testController, broker)
	if shouldSucceed && err != nil {
		t.Fatal("Should have succeeded to get the catalog for the broker. got error: ", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	if shouldSucceed {
		assertNumberOfBrokerActions(t, brokerActions, 1)
		assertGetCatalog(t, brokerActions[0])
	} else {
		assertNumberOfBrokerActions(t, brokerActions, 0)
	}
	actions := fakeCatalogClient.Actions()
	if shouldSucceed {
		assertNumberOfActions(t, actions, 2)
		assertCreate(t, actions[0], testClusterServiceClass)
		updatedClusterServiceBroker := assertUpdateStatus(t, actions[1], broker)
		assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	} else {
		assertNumberOfActions(t, actions, 1)
		updatedClusterServiceBroker := assertUpdateStatus(t, actions[0], broker)
		assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
	}
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 1)
	getAction := kubeActions[0].(clientgotesting.GetAction)
	if e, a := "get", getAction.GetVerb(); e != a {
		t.Fatalf("Unexpected verb on action; %s", expectedGot(e, a))
	}
	if e, a := "secrets", getAction.GetResource().Resource; e != a {
		t.Fatalf("Unexpected resource on action; %s", expectedGot(e, a))
	}
	events := getRecordedEvents(testController)
	assertNumEvents(t, events, 1)
	var expectedEvent string
	if shouldSucceed {
		expectedEvent = corev1.EventTypeNormal + " " + successFetchedCatalogReason + " " + successFetchedCatalogMessage
	} else {
		expectedEvent = corev1.EventTypeWarning + " " + errorAuthCredentialsReason + " " + `Error getting broker auth credentials`
	}
	if e, a := expectedEvent, events[0]; !strings.HasPrefix(a, e) {
		t.Fatalf("Received unexpected event, %s", expectedGot(e, a))
	}
}
func TestReconcileClusterServiceBrokerWithReconcileError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, getTestCatalogConfig())
	broker := getTestClusterServiceBroker()
	fakeCatalogClient.AddReactor("create", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, nil, errors.New("error creating serviceclass")
	})
	if err := reconcileClusterServiceBroker(t, testController, broker); err == nil {
		t.Fatal("There should have been an error.")
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 4)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", broker.Name)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	createSCAction := actions[2].(clientgotesting.CreateAction)
	createdSC, ok := createSCAction.GetObject().(*v1beta1.ClusterServiceClass)
	if !ok {
		t.Fatalf("couldn't convert to a ClusterServiceClass: %+v", createSCAction.GetObject())
	}
	if e, a := getTestClusterServiceClass(), createdSC; !reflect.DeepEqual(e, a) {
		t.Fatalf("unexpected diff for created ClusterServiceClass: %v,\n\nEXPECTED: %+v\n\nACTUAL:  %+v", diff.ObjectReflectDiff(e, a), e, a)
	}
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[3], broker)
	assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
	events := getRecordedEvents(testController)
	expectedEvent := warningEventBuilder(errorSyncingCatalogReason).msgf("Error reconciling ClusterServiceClass (K8S: %q ExternalName: %q) (broker %q):", testClusterServiceClassGUID, testClusterServiceClassName, testClusterServiceBrokerName).msg("error creating serviceclass")
	if err := checkEvents(events, expectedEvent.stringArr()); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileClusterServiceBrokerSuccessOnFinalRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	broker := getTestClusterServiceBroker()
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	broker.Status.OperationStartTime = &startTime
	if err := reconcileClusterServiceBroker(t, testController, broker); err != nil {
		t.Fatalf("This should not fail : %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 7)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", broker.Name)}
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[0], getTestClusterServiceBroker())
	assertClusterServiceBrokerOperationStartTimeSet(t, updatedClusterServiceBroker, false)
	assertList(t, actions[1], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[2], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertCreate(t, actions[3], testClusterServiceClass)
	assertCreate(t, actions[4], getTestClusterServicePlan())
	assertCreate(t, actions[5], getTestClusterServicePlanNonbindable())
	updatedClusterServiceBroker = assertUpdateStatus(t, actions[6], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestReconcileClusterServiceBrokerFailureOnFinalRetry(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, fakeosb.FakeClientConfiguration{CatalogReaction: &fakeosb.CatalogReaction{Error: errors.New("ooops")}})
	broker := getTestClusterServiceBroker()
	startTime := metav1.NewTime(time.Now().Add(-7 * 24 * time.Hour))
	broker.Status.OperationStartTime = &startTime
	if err := reconcileClusterServiceBroker(t, testController, broker); err != nil {
		t.Fatalf("Should have return no error because the retry duration has elapsed: %v", err)
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 2)
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[0], broker)
	assertClusterServiceBrokerReadyFalse(t, updatedClusterServiceBroker)
	updatedClusterServiceBroker = assertUpdateStatus(t, actions[1], broker)
	assertClusterServiceBrokerCondition(t, updatedClusterServiceBroker, v1beta1.ServiceBrokerConditionFailed, v1beta1.ConditionTrue)
	assertClusterServiceBrokerOperationStartTimeSet(t, updatedClusterServiceBroker, false)
	assertNumberOfActions(t, fakeKubeClient.Actions(), 0)
	events := getRecordedEvents(testController)
	expectedEventPrefixes := []string{warningEventBuilder(errorFetchingCatalogReason).String(), warningEventBuilder(errorReconciliationRetryTimeoutReason).String()}
	if err := checkEventPrefixes(events, expectedEventPrefixes); err != nil {
		t.Fatal(err)
	}
}
func TestReconcileClusterServiceBrokerWithStatusUpdateError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeKubeClient, fakeCatalogClient, fakeClusterServiceBrokerClient, testController, _ := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	broker := getTestClusterServiceBroker()
	fakeCatalogClient.AddReactor("update", "clusterservicebrokers", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, nil, errors.New("update error")
	})
	err := reconcileClusterServiceBroker(t, testController, broker)
	if err == nil {
		t.Fatalf("expected error from but got none")
	}
	if e, a := "update error", err.Error(); e != a {
		t.Fatalf("unexpected error returned: %s", expectedGot(e, a))
	}
	brokerActions := fakeClusterServiceBrokerClient.Actions()
	assertNumberOfBrokerActions(t, brokerActions, 1)
	assertGetCatalog(t, brokerActions[0])
	actions := fakeCatalogClient.Actions()
	assertNumberOfActions(t, actions, 6)
	listRestrictions := clientgotesting.ListRestrictions{Labels: labels.Everything(), Fields: fields.OneTermEqualSelector("spec.clusterServiceBrokerName", broker.Name)}
	assertList(t, actions[0], &v1beta1.ClusterServiceClass{}, listRestrictions)
	assertList(t, actions[1], &v1beta1.ClusterServicePlan{}, listRestrictions)
	assertCreate(t, actions[2], testClusterServiceClass)
	assertCreate(t, actions[3], getTestClusterServicePlan())
	assertCreate(t, actions[4], getTestClusterServicePlanNonbindable())
	updatedClusterServiceBroker := assertUpdateStatus(t, actions[5], getTestClusterServiceBroker())
	assertClusterServiceBrokerReadyTrue(t, updatedClusterServiceBroker)
	kubeActions := fakeKubeClient.Actions()
	assertNumberOfActions(t, kubeActions, 0)
}
func TestUpdateServiceBrokerCondition(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name			string
		input			*v1beta1.ClusterServiceBroker
		status			v1beta1.ConditionStatus
		reason			string
		message			string
		transitionTimeChanged	bool
	}{{name: "initially unset", input: getTestClusterServiceBroker(), status: v1beta1.ConditionFalse, transitionTimeChanged: true}, {name: "not ready -> not ready", input: getTestClusterServiceBrokerWithStatus(v1beta1.ConditionFalse), status: v1beta1.ConditionFalse, transitionTimeChanged: false}, {name: "not ready -> not ready with reason and message change", input: getTestClusterServiceBrokerWithStatus(v1beta1.ConditionFalse), status: v1beta1.ConditionFalse, reason: "foo", message: "bar", transitionTimeChanged: false}, {name: "not ready -> ready", input: getTestClusterServiceBrokerWithStatus(v1beta1.ConditionFalse), status: v1beta1.ConditionTrue, transitionTimeChanged: true}, {name: "ready -> ready", input: getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue), status: v1beta1.ConditionTrue, transitionTimeChanged: false}, {name: "ready -> not ready", input: getTestClusterServiceBrokerWithStatus(v1beta1.ConditionTrue), status: v1beta1.ConditionFalse, transitionTimeChanged: true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, fakeCatalogClient, _, testController, _ := newTestController(t, getTestCatalogConfig())
			inputClone := tc.input.DeepCopy()
			err := testController.updateClusterServiceBrokerCondition(tc.input, v1beta1.ServiceBrokerConditionReady, tc.status, tc.reason, tc.message)
			if err != nil {
				t.Fatalf("%v: error updating broker condition: %v", tc.name, err)
			}
			if !reflect.DeepEqual(tc.input, inputClone) {
				t.Fatalf("%v: updating broker condition mutated input: %s", tc.name, expectedGot(inputClone, tc.input))
			}
			actions := fakeCatalogClient.Actions()
			assertNumberOfActions(t, actions, 1)
			updatedClusterServiceBroker := assertUpdateStatus(t, actions[0], tc.input)
			updateActionObject, ok := updatedClusterServiceBroker.(*v1beta1.ClusterServiceBroker)
			if !ok {
				t.Fatalf("%v: couldn't convert to broker", tc.name)
			}
			var initialTs metav1.Time
			if len(inputClone.Status.Conditions) != 0 {
				initialTs = inputClone.Status.Conditions[0].LastTransitionTime
			}
			if e, a := 1, len(updateActionObject.Status.Conditions); e != a {
				t.Fatalf("%v: %s", tc.name, expectedGot(e, a))
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
				t.Fatalf("%v: condition message didn't match; %s", tc.name, expectedGot(e, a))
			}
		})
	}
}
func TestReconcileClusterServicePlanFromClusterServiceBrokerCatalog(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	updatedPlan := func() *v1beta1.ClusterServicePlan {
		p := getTestClusterServicePlan()
		p.Spec.Description = "new-description"
		p.Spec.ExternalName = "new-value"
		p.Spec.Free = false
		p.Spec.ExternalMetadata = &runtime.RawExtension{Raw: []byte(`{"field1": "value1"}`)}
		p.Spec.InstanceCreateParameterSchema = &runtime.RawExtension{Raw: []byte(`{"field1": "value1"}`)}
		p.Spec.InstanceUpdateParameterSchema = &runtime.RawExtension{Raw: []byte(`{"field1": "value1"}`)}
		p.Spec.ServiceBindingCreateParameterSchema = &runtime.RawExtension{Raw: []byte(`{"field1": "value1"}`)}
		return p
	}
	cases := []struct {
		name			string
		newServicePlan		*v1beta1.ClusterServicePlan
		existingServicePlan	*v1beta1.ClusterServicePlan
		listerServicePlan	*v1beta1.ClusterServicePlan
		shouldError		bool
		errText			*string
		catalogClientPrepFunc	func(*fake.Clientset)
		catalogActionsCheckFunc	func(t *testing.T, name string, actions []clientgotesting.Action)
	}{{name: "new plan", newServicePlan: getTestClusterServicePlan(), shouldError: false, catalogActionsCheckFunc: func(t *testing.T, name string, actions []clientgotesting.Action) {
		assertNumberOfActions(t, actions, 1)
		assertCreate(t, actions[0], getTestClusterServicePlan())
	}}, {name: "exists, but for a different broker", newServicePlan: getTestClusterServicePlan(), existingServicePlan: getTestClusterServicePlan(), listerServicePlan: func() *v1beta1.ClusterServicePlan {
		p := getTestClusterServicePlan()
		p.Spec.ClusterServiceBrokerName = "something-else"
		return p
	}(), shouldError: true, errText: strPtr(`ClusterServiceBroker "test-clusterservicebroker": ClusterServicePlan "test-clusterserviceplan" already exists for Broker "something-else"`)}, {name: "plan update", newServicePlan: updatedPlan(), existingServicePlan: getTestClusterServicePlan(), shouldError: false, catalogActionsCheckFunc: func(t *testing.T, name string, actions []clientgotesting.Action) {
		assertNumberOfActions(t, actions, 1)
		assertUpdate(t, actions[0], updatedPlan())
	}}, {name: "plan update - failure", newServicePlan: updatedPlan(), existingServicePlan: getTestClusterServicePlan(), catalogClientPrepFunc: func(client *fake.Clientset) {
		client.AddReactor("update", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
			return true, nil, errors.New("oops")
		})
	}, shouldError: true, errText: strPtr("oops")}}
	broker := getTestClusterServiceBroker()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, noFakeActions())
			if tc.catalogClientPrepFunc != nil {
				tc.catalogClientPrepFunc(fakeCatalogClient)
			}
			if tc.listerServicePlan != nil {
				sharedInformers.ClusterServicePlans().Informer().GetStore().Add(tc.listerServicePlan)
			}
			err := testController.reconcileClusterServicePlanFromClusterServiceBrokerCatalog(broker, tc.newServicePlan, tc.existingServicePlan)
			if err != nil {
				if !tc.shouldError {
					t.Fatalf("%v: unexpected error from method under test: %v", tc.name, err)
				} else if tc.errText != nil && *tc.errText != err.Error() {
					t.Fatalf("%v: unexpected error text from method under test; %s", tc.name, expectedGot(tc.errText, err.Error()))
				}
			}
			if tc.catalogActionsCheckFunc != nil {
				actions := fakeCatalogClient.Actions()
				tc.catalogActionsCheckFunc(t, tc.name, actions)
			}
		})
	}
}
func reconcileClusterServiceBroker(t *testing.T, testController *controller, broker *v1beta1.ClusterServiceBroker) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clone := broker.DeepCopy()
	err := testController.reconcileClusterServiceBroker(broker)
	if !reflect.DeepEqual(broker, clone) {
		t.Errorf("reconcileClusterServiceBroker shouldn't mutate input, but it does: %s", expectedGot(clone, broker))
	}
	return err
}
func TestReconcileUpdatesManagedClassesAndPlans(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(testClusterServicePlan)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass}}, nil
	})
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: []v1beta1.ClusterServicePlan{*testClusterServicePlan}}, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	actions := fakeCatalogClient.Actions()
	c := assertUpdate(t, actions[2], testClusterServiceClass)
	updatedClass, ok := c.(metav1.Object)
	if !ok {
		t.Fatalf("could not cast %T to metav1.Object", c)
	}
	if !isServiceCatalogManagedResource(updatedClass) {
		t.Error("expected the class to have a service catalog controller reference")
	}
	p := assertUpdate(t, actions[3], testClusterServicePlan)
	updatedPlan, ok := p.(metav1.Object)
	if !ok {
		t.Fatalf("could not cast %T to metav1.Object", p)
	}
	if !isServiceCatalogManagedResource(updatedPlan) {
		t.Error("expected the plan to have a service catalog controller reference")
	}
}
func TestReconcileCreatesManagedClassesAndPlans(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, _, testController, _ := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	actions := fakeCatalogClient.Actions()
	c := assertCreate(t, actions[2], testClusterServiceClass)
	createdClass, ok := c.(metav1.Object)
	if !ok {
		t.Fatalf("could not cast %T to metav1.Object", c)
	}
	if !isServiceCatalogManagedResource(createdClass) {
		t.Error("expected the class to have a service catalog controller reference")
	}
	p := assertCreate(t, actions[3], testClusterServicePlan)
	createdPlan, ok := p.(metav1.Object)
	if !ok {
		t.Fatalf("could not cast %T to metav1.Object", p)
	}
	if !isServiceCatalogManagedResource(createdPlan) {
		t.Error("expected the plan to have a service catalog controller reference")
	}
}
func TestReconcileMarksExistingClassesAndPlansAsManaged(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServiceClass.ObjectMeta.OwnerReferences = nil
	testClusterServicePlan.ObjectMeta.OwnerReferences = nil
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(testClusterServicePlan)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass}}, nil
	})
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: []v1beta1.ClusterServicePlan{*testClusterServicePlan}}, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	actions := fakeCatalogClient.Actions()
	c := assertUpdate(t, actions[2], testClusterServiceClass)
	updatedClass, ok := c.(metav1.Object)
	if !ok {
		t.Fatalf("could not cast %T to metav1.Object", c)
	}
	if !isServiceCatalogManagedResource(updatedClass) {
		t.Error("expected the class to have a service catalog controller reference")
	}
	p := assertUpdate(t, actions[3], testClusterServicePlan)
	updatedPlan, ok := p.(metav1.Object)
	if !ok {
		t.Fatalf("could not cast %T to metav1.Object", p)
	}
	if !isServiceCatalogManagedResource(updatedPlan) {
		t.Error("expected the plan to have a service catalog controller reference")
	}
}
func TestReconcileDoesNotUpdateUserDefinedClassesAndPlans(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, fakeCatalogClient, _, testController, sharedInformers := newTestController(t, getTestCatalogConfig())
	testClusterServiceClass := getTestClusterServiceClass()
	testClusterServicePlan := getTestClusterServicePlan()
	testClusterServiceClass.OwnerReferences = nil
	testClusterServiceClass.Name = "user-defined-class"
	testClusterServicePlan.OwnerReferences = nil
	testClusterServicePlan.Name = "user-defined-plan"
	sharedInformers.ClusterServiceClasses().Informer().GetStore().Add(testClusterServiceClass)
	sharedInformers.ClusterServicePlans().Informer().GetStore().Add(testClusterServicePlan)
	fakeCatalogClient.AddReactor("list", "clusterserviceclasses", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServiceClassList{Items: []v1beta1.ClusterServiceClass{*testClusterServiceClass}}, nil
	})
	fakeCatalogClient.AddReactor("list", "clusterserviceplans", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, &v1beta1.ClusterServicePlanList{Items: []v1beta1.ClusterServicePlan{*testClusterServicePlan}}, nil
	})
	if err := reconcileClusterServiceBroker(t, testController, getTestClusterServiceBroker()); err != nil {
		t.Fatalf("This should not fail: %v", err)
	}
	actions := fakeCatalogClient.Actions()
	for _, a := range actions {
		r := a.GetResource().Resource
		if a.GetVerb() == "update" && (r == "clusterserviceclasses" || r == "clusterserviceplans") {
			t.Errorf("expected user-defined classes and plans to be ignored but found action %+v", a)
		}
	}
}
func TestIsServiceCatalogManagedResource(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testcases := []struct {
		name		string
		resource	metav1.Object
		want		bool
	}{{"unmanaged service class", &v1beta1.ServiceClass{}, false}, {"unmanaged service plan", &v1beta1.ServicePlan{}, false}, {"managed service class", &v1beta1.ServiceClass{ObjectMeta: metav1.ObjectMeta{OwnerReferences: []metav1.OwnerReference{{Controller: truePtr(), APIVersion: v1beta1.SchemeGroupVersion.String()}}}}, true}, {"managed service plan", &v1beta1.ServicePlan{ObjectMeta: metav1.ObjectMeta{OwnerReferences: []metav1.OwnerReference{{Controller: truePtr(), APIVersion: v1beta1.SchemeGroupVersion.String()}}}}, true}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got := isServiceCatalogManagedResource(tc.resource)
			if tc.want != got {
				t.Fatalf("WANT: %v, GOT: %v", tc.want, got)
			}
		})
	}
}
func TestMarkAsServiceCatalogManagedResource(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	testcases := []struct {
		name		string
		resource	metav1.Object
	}{{"service class", &v1beta1.ServiceClass{}}, {"service plan", &v1beta1.ServicePlan{}}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			broker := getTestClusterServiceBroker()
			markAsServiceCatalogManagedResource(tc.resource, broker)
			numOwners := len(tc.resource.GetOwnerReferences())
			if numOwners != 1 {
				t.Fatalf("Expected 1 owner reference, got %v", numOwners)
			}
			gotOwner := tc.resource.GetOwnerReferences()[0]
			gotIsController := gotOwner.Controller != nil && *gotOwner.Controller == true
			if !gotIsController {
				t.Errorf("Expected a controller reference, but Controller is false")
			}
			gotBlockOwnerDeletion := gotOwner.BlockOwnerDeletion != nil && *gotOwner.BlockOwnerDeletion == true
			if gotBlockOwnerDeletion {
				t.Errorf("Expected the controller reference to not modify deletion semantics, but BlockOwnerDeletion is true")
			}
			wantAPIVersion := v1beta1.SchemeGroupVersion.String()
			gotAPIVersion := gotOwner.APIVersion
			if wantAPIVersion != gotAPIVersion {
				t.Errorf("unexpected APIVersion. WANT: %q, GOT: %q", wantAPIVersion, gotAPIVersion)
			}
			if !isServiceCatalogManagedResource(tc.resource) {
				t.Fatal("expected isServiceCatalogManagedResource to return true")
			}
		})
	}
}
