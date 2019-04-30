package controller

import (
	stderrors "errors"
	"fmt"
	"net/url"
	"sync"
	"time"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	"github.com/kubernetes-incubator/service-catalog/pkg/pretty"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	successDeprovisionReason			string		= "DeprovisionedSuccessfully"
	successDeprovisionMessage			string		= "The instance was deprovisioned successfully"
	successUpdateInstanceReason			string		= "InstanceUpdatedSuccessfully"
	successUpdateInstanceMessage			string		= "The instance was updated successfully"
	successProvisionReason				string		= "ProvisionedSuccessfully"
	successProvisionMessage				string		= "The instance was provisioned successfully"
	successOrphanMitigationReason			string		= "OrphanMitigationSuccessful"
	successOrphanMitigationMessage			string		= "Orphan mitigation was completed successfully"
	errorWithParametersReason			string		= "ErrorWithParameters"
	errorProvisionCallFailedReason			string		= "ProvisionCallFailed"
	errorErrorCallingProvisionReason		string		= "ErrorCallingProvision"
	errorUpdateInstanceCallFailedReason		string		= "UpdateInstanceCallFailed"
	errorErrorCallingUpdateInstanceReason		string		= "ErrorCallingUpdateInstance"
	errorDeprovisionCallFailedReason		string		= "DeprovisionCallFailed"
	errorDeprovisionBlockedByCredentialsReason	string		= "DeprovisionBlockedByExistingCredentials"
	errorPollingLastOperationReason			string		= "ErrorPollingLastOperation"
	errorWithOriginatingIdentityReason		string		= "ErrorWithOriginatingIdentity"
	errorWithOngoingAsyncOperationReason		string		= "ErrorAsyncOperationInProgress"
	errorNonexistentClusterServiceClassReason	string		= "ReferencesNonexistentServiceClass"
	errorNonexistentClusterServiceClassMessage	string		= "ReferencesNonexistentServiceClass"
	errorNonexistentClusterServicePlanReason	string		= "ReferencesNonexistentServicePlan"
	errorNonexistentClusterServiceBrokerReason	string		= "ReferencesNonexistentBroker"
	errorNonexistentServiceClassReason		string		= "ReferencesNonexistentServiceClass"
	errorNonexistentServicePlanReason		string		= "ReferencesNonexistentServicePlan"
	errorNonexistentServiceBrokerReason		string		= "ReferencesNonexistentBroker"
	errorDeletedClusterServiceClassReason		string		= "ReferencesDeletedServiceClass"
	errorDeletedClusterServicePlanReason		string		= "ReferencesDeletedServicePlan"
	errorDeletedServiceClassReason			string		= "ReferencesDeletedServiceClass"
	errorDeletedServicePlanReason			string		= "ReferencesDeletedServicePlan"
	errorFindingNamespaceServiceInstanceReason	string		= "ErrorFindingNamespaceForInstance"
	errorOrphanMitigationFailedReason		string		= "OrphanMitigationFailed"
	errorInvalidDeprovisionStatusReason		string		= "InvalidDeprovisionStatus"
	errorAmbiguousPlanReferenceScope		string		= "couldn't determine if the instance refers to a Cluster or Namespaced ServiceClass/Plan"
	asyncProvisioningReason				string		= "Provisioning"
	asyncProvisioningMessage			string		= "The instance is being provisioned asynchronously"
	asyncUpdatingInstanceReason			string		= "UpdatingInstance"
	asyncUpdatingInstanceMessage			string		= "The instance is being updated asynchronously"
	asyncDeprovisioningReason			string		= "Deprovisioning"
	asyncDeprovisioningMessage			string		= "The instance is being deprovisioned asynchronously"
	provisioningInFlightReason			string		= "ProvisionRequestInFlight"
	provisioningInFlightMessage			string		= "Provision request for ServiceInstance in-flight to Broker"
	instanceUpdatingInFlightReason			string		= "UpdateInstanceRequestInFlight"
	instanceUpdatingInFlightMessage			string		= "Update request for ServiceInstance in-flight to Broker"
	deprovisioningInFlightReason			string		= "DeprovisionRequestInFlight"
	deprovisioningInFlightMessage			string		= "Deprovision request for ServiceInstance in-flight to Broker"
	startingInstanceOrphanMitigationReason		string		= "StartingInstanceOrphanMitigation"
	startingInstanceOrphanMitigationMessage		string		= "The instance provision call failed with an ambiguous error; attempting to deprovision the instance in order to mitigate an orphaned resource"
	clusterIdentifierKey				string		= "clusterid"
	minBrokerOperationRetryDelay			time.Duration	= time.Second * 1
	maxBrokerOperationRetryDelay			time.Duration	= time.Minute * 20
	eventHandlerLogLevel						= 4
)

type backoffEntry struct {
	generation		int64
	calculatedRetryTime	time.Time
	dirty			bool
}
type instanceOperationBackoff struct {
	mutex		sync.RWMutex
	instances	map[string]backoffEntry
	rateLimiter	workqueue.RateLimiter
}

func (c *controller) enqueueInstance(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.instanceQueue.Add(key)
}
func (c *controller) enqueueInstanceAfter(obj interface{}, d time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.instanceQueue.AddAfter(key, d)
}
func (c *controller) instanceAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if klog.V(eventHandlerLogLevel) {
		instance := obj.(*v1beta1.ServiceInstance)
		pcb := pretty.NewInstanceContextBuilder(instance)
		klog.Info(pcb.Messagef("Received ADD event: %v", toJSON(instance)))
	}
	c.enqueueInstance(obj)
}
func (c *controller) instanceUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance := newObj.(*v1beta1.ServiceInstance)
	pcb := pretty.NewInstanceContextBuilder(instance)
	if klog.V(eventHandlerLogLevel) {
		pcb := pretty.NewInstanceContextBuilder(instance)
		klog.Info(pcb.Messagef("Received UPDATE event: %v", toJSON(instance)))
	}
	if instance.Status.AsyncOpInProgress {
		klog.V(eventHandlerLogLevel).Info(pcb.Message("NOT enqueueing instance because an async operation is in progress"))
		return
	}
	klog.V(eventHandlerLogLevel).Info(pcb.Message("Enqueueing instance"))
	c.enqueueInstance(newObj)
}
func (c *controller) instanceDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, ok := obj.(*v1beta1.ServiceInstance)
	if instance == nil || !ok {
		return
	}
	if klog.V(eventHandlerLogLevel) {
		pcb := pretty.NewInstanceContextBuilder(instance)
		klog.Info(pcb.Messagef("Received DELETE event: %v", toJSON(instance)))
		klog.Info(pcb.Message("no further processing will occur"))
	}
}
func (c *controller) requeueServiceInstanceForPoll(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.instanceQueue.Add(key)
	return nil
}
func (c *controller) beginPollingServiceInstance(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(instance)
	if err != nil {
		pcb := pretty.NewInstanceContextBuilder(instance)
		s := fmt.Sprintf("Couldn't create a key for object %+v: %v", instance, err)
		klog.Errorf(pcb.Message(s))
		return fmt.Errorf(s)
	}
	c.instancePollingQueue.AddRateLimited(key)
	return nil
}
func (c *controller) continuePollingServiceInstance(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.beginPollingServiceInstance(instance)
}
func (c *controller) finishPollingServiceInstance(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(instance)
	if err != nil {
		pcb := pretty.NewInstanceContextBuilder(instance)
		s := fmt.Sprintf("Couldn't create a key for object %+v: %v", instance, err)
		klog.Errorf(pcb.Message(s))
		return fmt.Errorf(s)
	}
	c.instancePollingQueue.Forget(key)
	return nil
}
func (c *controller) resetPollingRateLimiterForServiceInstance(instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(instance)
	if err != nil {
		pcb := pretty.NewInstanceContextBuilder(instance)
		s := fmt.Sprintf("Couldn't create a key for object %+v: %v", instance, err)
		klog.Errorf(pcb.Message(s))
		return
	}
	c.instancePollingQueue.Forget(key)
}
func getReconciliationActionForServiceInstance(instance *v1beta1.ServiceInstance) ReconciliationAction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case instance.Status.AsyncOpInProgress:
		return reconcilePoll
	case instance.ObjectMeta.DeletionTimestamp != nil || instance.Status.OrphanMitigationInProgress:
		return reconcileDelete
	case instance.Status.ProvisionStatus == v1beta1.ServiceInstanceProvisionStatusProvisioned:
		return reconcileUpdate
	default:
		return reconcileAdd
	}
}
func (c *controller) reconcileServiceInstanceKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	pcb := pretty.NewContextBuilder(pretty.ServiceInstance, namespace, name, "")
	instance, err := c.instanceLister.ServiceInstances(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.Info(pcb.Messagef("Not doing work for %v because it has been deleted", key))
		return nil
	}
	if err != nil {
		klog.Errorf(pcb.Messagef("Unable to retrieve %v from store: %v", key, err))
		return err
	}
	return c.reconcileServiceInstance(instance)
}
func (c *controller) reconcileServiceInstance(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	updated, err := c.initObservedGeneration(instance)
	if err != nil {
		return err
	}
	if updated {
		return nil
	}
	updated, err = c.initOrphanMitigationCondition(instance)
	if err != nil {
		return err
	}
	if updated {
		return nil
	}
	reconciliationAction := getReconciliationActionForServiceInstance(instance)
	switch reconciliationAction {
	case reconcileAdd:
		return c.reconcileServiceInstanceAdd(instance)
	case reconcileUpdate:
		return c.reconcileServiceInstanceUpdate(instance)
	case reconcileDelete:
		return c.reconcileServiceInstanceDelete(instance)
	case reconcilePoll:
		return c.pollServiceInstance(instance)
	default:
		pcb := pretty.NewInstanceContextBuilder(instance)
		return fmt.Errorf(pcb.Messagef("Unknown reconciliation action %v", reconciliationAction))
	}
}
func (c *controller) initObservedGeneration(instance *v1beta1.ServiceInstance) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Status.ObservedGeneration == 0 && instance.Status.ReconciledGeneration != 0 {
		instance = instance.DeepCopy()
		instance.Status.ObservedGeneration = instance.Status.ReconciledGeneration
		provisioned := !isServiceInstanceFailed(instance)
		if provisioned {
			instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
		} else {
			instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
		}
		updatedInstance, err := c.updateServiceInstanceStatus(instance)
		if err != nil {
			return false, err
		}
		return updatedInstance.ResourceVersion != instance.ResourceVersion, nil
	}
	return false, nil
}
func (c *controller) initOrphanMitigationCondition(instance *v1beta1.ServiceInstance) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !isServiceInstanceOrphanMitigation(instance) && instance.Status.OrphanMitigationInProgress {
		instance := instance.DeepCopy()
		reason := startingInstanceOrphanMitigationReason
		message := startingInstanceOrphanMitigationMessage
		c.recorder.Event(instance, corev1.EventTypeWarning, reason, message)
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionOrphanMitigation, v1beta1.ConditionTrue, reason, message)
		updatedInstance, err := c.updateServiceInstanceStatus(instance)
		if err != nil {
			return false, err
		}
		return updatedInstance.ResourceVersion != instance.ResourceVersion, nil
	}
	return false, nil
}
func (c *controller) setRetryBackoffRequired(instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	c.instanceOperationRetryQueue.mutex.Lock()
	defer c.instanceOperationRetryQueue.mutex.Unlock()
	key := string(instance.GetUID())
	retryEntry, found := c.instanceOperationRetryQueue.instances[key]
	if !found || retryEntry.generation != instance.Generation {
		retryEntry.generation = instance.Generation
		if found {
			c.instanceOperationRetryQueue.rateLimiter.Forget(key)
		}
	}
	retryEntry.dirty = true
	c.instanceOperationRetryQueue.instances[key] = retryEntry
	klog.V(4).Info(pcb.Messagef("BrokerOpRetry: added %v (%v/%v) generation %v to backoffBeforeRetrying map", key, instance.GetNamespace(), instance.GetName(), instance.Generation))
}
func (c *controller) backoffAndRequeueIfRetrying(instance *v1beta1.ServiceInstance, operation string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	key := string(instance.GetUID())
	delay := time.Millisecond * 0
	c.instanceOperationRetryQueue.mutex.Lock()
	defer c.instanceOperationRetryQueue.mutex.Unlock()
	retryEntry, exists := c.instanceOperationRetryQueue.instances[key]
	if exists {
		if retryEntry.generation != instance.Generation {
			delete(c.instanceOperationRetryQueue.instances, key)
			c.instanceOperationRetryQueue.rateLimiter.Forget(key)
			return false
		}
		if retryEntry.dirty {
			retryEntry.calculatedRetryTime = time.Now().Add(c.instanceOperationRetryQueue.rateLimiter.When(key))
			retryEntry.dirty = false
			c.instanceOperationRetryQueue.instances[key] = retryEntry
			klog.V(4).Infof(pcb.Messagef("BrokerOpRetry: generation %v retryTime calculated as %v", instance.Generation, retryEntry.calculatedRetryTime))
		}
		now := time.Now()
		delay = retryEntry.calculatedRetryTime.Sub(now)
		if delay > 0 {
			msg := fmt.Sprintf("Delaying %s retry, next attempt will be after %s", operation, retryEntry.calculatedRetryTime)
			c.recorder.Event(instance, corev1.EventTypeWarning, "RetryBackoff", msg)
			klog.V(2).Info(pcb.Messagef("BrokerOpRetry: %s", msg))
			c.enqueueInstanceAfter(instance, delay)
			return true
		}
	}
	return false
}
func (c *controller) purgeExpiredRetryEntries() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	now := time.Now()
	c.instanceOperationRetryQueue.mutex.Lock()
	defer c.instanceOperationRetryQueue.mutex.Unlock()
	overDue := now.Add(-maxBrokerOperationRetryDelay)
	purgedEntries := 0
	for k, v := range c.instanceOperationRetryQueue.instances {
		if v.calculatedRetryTime.Before(overDue) {
			klog.V(5).Infof("BrokerOpRetry: removing %s from instanceOperationRetryQueue which had retry time of %v", k, v.calculatedRetryTime)
			delete(c.instanceOperationRetryQueue.instances, k)
			c.instanceOperationRetryQueue.rateLimiter.Forget(k)
			purgedEntries++
		}
	}
	klog.V(5).Infof("BrokerOpRetry: purged %v expired entries from instanceOperationRetryQueue.instances, number of entries remaining: %v", purgedEntries, len(c.instanceOperationRetryQueue.instances))
}
func (c *controller) removeInstanceFromRetryMap(instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	key := string(instance.GetUID())
	c.instanceOperationRetryQueue.mutex.Lock()
	defer c.instanceOperationRetryQueue.mutex.Unlock()
	delete(c.instanceOperationRetryQueue.instances, key)
	c.instanceOperationRetryQueue.rateLimiter.Forget(key)
	klog.V(4).Infof(pcb.Message("BrokerOpRetry: removed %v from instanceOperationRetryQueue"), key)
}
func (c *controller) reconcileServiceInstanceAdd(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	if isServiceInstanceProcessedAlready(instance) {
		klog.V(4).Info(pcb.Message("Not processing event because status showed there is no work to do"))
		return nil
	}
	if c.backoffAndRequeueIfRetrying(instance, "provision") {
		return nil
	}
	instance = instance.DeepCopy()
	if instance.Status.ObservedGeneration != instance.Generation {
		c.prepareObservedGeneration(instance)
	}
	modified, err := c.resolveReferences(instance)
	if err != nil {
		return err
	}
	if modified {
		return nil
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.ServicePlanDefaults) {
		modified, err = c.applyDefaultProvisioningParameters(instance)
		if err != nil {
			return err
		}
		if modified {
			return nil
		}
	}
	klog.V(4).Info(pcb.Message("Processing adding event"))
	request, inProgressProperties, err := c.prepareProvisionRequest(instance)
	if err != nil {
		return c.handleServiceInstanceReconciliationError(instance, err)
	}
	if instance.Status.CurrentOperation == "" || !isServiceInstancePropertiesStateEqual(instance.Status.InProgressProperties, inProgressProperties) {
		updatedInstance, err := c.recordStartOfServiceInstanceOperation(instance, v1beta1.ServiceInstanceOperationProvision, inProgressProperties)
		if err != nil {
			return err
		}
		if updatedInstance.ResourceVersion != instance.ResourceVersion {
			return nil
		}
		instance = updatedInstance
	} else if instance.Status.DeprovisionStatus != v1beta1.ServiceInstanceDeprovisionStatusRequired {
		instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
		updatedInstance, err := c.updateServiceInstanceStatus(instance)
		if err != nil {
			return err
		}
		if updatedInstance.ResourceVersion != instance.ResourceVersion {
			return nil
		}
		instance = updatedInstance
	}
	var prettyClass string
	var brokerName string
	var brokerClient osb.Client
	if instance.Spec.ClusterServiceClassSpecified() {
		var serviceClass *v1beta1.ClusterServiceClass
		serviceClass, _, brokerName, brokerClient, _ = c.getClusterServiceClassPlanAndClusterServiceBroker(instance)
		prettyClass = pretty.ClusterServiceClassName(serviceClass)
	} else {
		var serviceClass *v1beta1.ServiceClass
		serviceClass, _, brokerName, brokerClient, _ = c.getServiceClassPlanAndServiceBroker(instance)
		prettyClass = pretty.ServiceClassName(serviceClass)
	}
	klog.V(4).Info(pcb.Messagef("Provisioning a new ServiceInstance of %s at Broker %q", prettyClass, brokerName))
	c.setRetryBackoffRequired(instance)
	response, err := brokerClient.ProvisionInstance(request)
	if err != nil {
		if httpErr, ok := osb.IsHTTPError(err); ok {
			msg := fmt.Sprintf("Error provisioning ServiceInstance of %s at ClusterServiceBroker %q: %s", prettyClass, brokerName, httpErr)
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, errorProvisionCallFailedReason, msg)
			shouldMitigateOrphan := shouldStartOrphanMitigation(httpErr.StatusCode)
			if isRetriableHTTPStatus(httpErr.StatusCode) {
				return c.processTemporaryProvisionFailure(instance, readyCond, shouldMitigateOrphan)
			}
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, "ClusterServiceBrokerReturnedFailure", msg)
			return c.processTerminalProvisionFailure(instance, readyCond, failedCond, shouldMitigateOrphan)
		}
		reason := errorErrorCallingProvisionReason
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			msg := fmt.Sprintf("Communication with the ClusterServiceBroker timed out; operation will be retried: %v", urlErr)
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, msg)
			return c.processTemporaryProvisionFailure(instance, readyCond, true)
		}
		msg := fmt.Sprintf("The provision call failed and will be retried: Error communicating with broker for provisioning: %v", err)
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, msg)
		if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
			msg := "Stopping reconciliation retries because too much time has elapsed"
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
			return c.processTerminalProvisionFailure(instance, readyCond, failedCond, false)
		}
		return c.processServiceInstanceOperationError(instance, readyCond)
	}
	if response.Async {
		return c.processProvisionAsyncResponse(instance, response)
	}
	return c.processProvisionSuccess(instance, response.DashboardURL)
}
func (c *controller) reconcileServiceInstanceUpdate(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	if isServiceInstanceProcessedAlready(instance) {
		klog.V(4).Info(pcb.Message("Not processing event because status showed there is no work to do"))
		return nil
	}
	if c.backoffAndRequeueIfRetrying(instance, "update") {
		return nil
	}
	instance = instance.DeepCopy()
	if instance.Status.ObservedGeneration != instance.Generation {
		c.prepareObservedGeneration(instance)
	}
	modified, err := c.resolveReferences(instance)
	if err != nil {
		return err
	}
	if modified {
		return nil
	}
	klog.V(4).Info(pcb.Message("Processing updating event"))
	var brokerClient osb.Client
	var request *osb.UpdateInstanceRequest
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, servicePlan, brokerName, bClient, err := c.getClusterServiceClassPlanAndClusterServiceBroker(instance)
		if err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		brokerClient = bClient
		if err := c.checkForRemovedClusterClassAndPlan(instance, serviceClass, servicePlan); err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		req, inProgressProperties, err := c.prepareUpdateInstanceRequest(instance)
		if err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		request = req
		if instance.Status.CurrentOperation == "" || !isServiceInstancePropertiesStateEqual(instance.Status.InProgressProperties, inProgressProperties) {
			updatedInstance, err := c.recordStartOfServiceInstanceOperation(instance, v1beta1.ServiceInstanceOperationUpdate, inProgressProperties)
			if err != nil {
				return err
			}
			if updatedInstance.ResourceVersion != instance.ResourceVersion {
				return nil
			}
			instance = updatedInstance
		}
		klog.V(4).Info(pcb.Messagef("Updating ServiceInstance of %s at ClusterServiceBroker %q", pretty.ClusterServiceClassName(serviceClass), brokerName))
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, servicePlan, brokerName, bClient, err := c.getServiceClassPlanAndServiceBroker(instance)
		if err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		brokerClient = bClient
		if err := c.checkForRemovedClassAndPlan(instance, serviceClass, servicePlan); err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		req, inProgressProperties, err := c.prepareUpdateInstanceRequest(instance)
		if err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		request = req
		if instance.Status.CurrentOperation == "" || !isServiceInstancePropertiesStateEqual(instance.Status.InProgressProperties, inProgressProperties) {
			updatedInstance, err := c.recordStartOfServiceInstanceOperation(instance, v1beta1.ServiceInstanceOperationUpdate, inProgressProperties)
			if err != nil {
				return err
			}
			if updatedInstance.ResourceVersion != instance.ResourceVersion {
				return nil
			}
			instance = updatedInstance
		}
		klog.V(4).Info(pcb.Messagef("Updating ServiceInstance of %s at ServiceBroker %q", pretty.ServiceClassName(serviceClass), brokerName))
	}
	c.setRetryBackoffRequired(instance)
	response, err := brokerClient.UpdateInstance(request)
	if err != nil {
		if httpErr, ok := osb.IsHTTPError(err); ok {
			if isRetriableHTTPStatus(httpErr.StatusCode) {
				msg := fmt.Sprintf("ServiceBroker returned a failure for update call; update will be retried: %v", httpErr)
				readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, errorUpdateInstanceCallFailedReason, msg)
				return c.processTemporaryUpdateServiceInstanceFailure(instance, readyCond)
			}
			msg := fmt.Sprintf("ServiceBroker returned a failure for update call; update will not be retried: %v", httpErr)
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, errorUpdateInstanceCallFailedReason, msg)
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, errorUpdateInstanceCallFailedReason, msg)
			return c.processTerminalUpdateServiceInstanceFailure(instance, readyCond, failedCond)
		}
		reason := errorErrorCallingUpdateInstanceReason
		if urlErr, ok := err.(*url.Error); ok && urlErr.Timeout() {
			msg := fmt.Sprintf("Communication with the ServiceBroker timed out; update will be retried: %v", urlErr)
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, msg)
			return c.processTemporaryUpdateServiceInstanceFailure(instance, readyCond)
		}
		msg := fmt.Sprintf("The update call failed and will be retried: Error communicating with broker for updating: %s", err)
		if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
			klog.Info(pcb.Message(msg))
			c.recorder.Event(instance, corev1.EventTypeWarning, reason, msg)
			msg = "Stopping reconciliation retries because too much time has elapsed"
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, errorReconciliationRetryTimeoutReason, msg)
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
			return c.processTerminalUpdateServiceInstanceFailure(instance, readyCond, failedCond)
		}
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, msg)
		return c.processServiceInstanceOperationError(instance, readyCond)
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.UpdateDashboardURL) {
		if *response.DashboardURL != "" {
			instance.Status.DashboardURL = response.DashboardURL
		}
	}
	if response.Async {
		return c.processUpdateServiceInstanceAsyncResponse(instance, response)
	}
	return c.processUpdateServiceInstanceSuccess(instance)
}
func (c *controller) reconcileServiceInstanceDelete(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if finalizers := sets.NewString(instance.Finalizers...); !finalizers.Has(v1beta1.FinalizerServiceCatalog) {
		return nil
	}
	pcb := pretty.NewInstanceContextBuilder(instance)
	if instance.Status.DeprovisionStatus == v1beta1.ServiceInstanceDeprovisionStatusFailed {
		klog.V(4).Info(pcb.Message("Not processing deleting event because deprovisioning has failed"))
		return nil
	}
	if instance.Status.OrphanMitigationInProgress {
		klog.V(4).Info(pcb.Message("Performing orphan mitigation"))
	} else {
		klog.V(4).Info(pcb.Message("Processing deleting event"))
	}
	instance = instance.DeepCopy()
	if !instance.Status.OrphanMitigationInProgress && instance.Status.ObservedGeneration != instance.Generation {
		c.prepareObservedGeneration(instance)
	}
	if instance.Status.DeprovisionStatus == v1beta1.ServiceInstanceDeprovisionStatusNotRequired || instance.Status.DeprovisionStatus == v1beta1.ServiceInstanceDeprovisionStatusSucceeded {
		return c.processServiceInstanceGracefulDeletionSuccess(instance)
	}
	if instance.Status.DeprovisionStatus != v1beta1.ServiceInstanceDeprovisionStatusRequired {
		msg := fmt.Sprintf("ServiceInstance has invalid DeprovisionStatus field: %v", instance.Status.DeprovisionStatus)
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionUnknown, errorInvalidDeprovisionStatusReason, msg)
		failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, errorInvalidDeprovisionStatusReason, msg)
		return c.processDeprovisionFailure(instance, readyCond, failedCond)
	}
	if err := c.checkServiceInstanceHasExistingBindings(instance); err != nil {
		return c.handleServiceInstanceReconciliationError(instance, err)
	}
	var prettyName string
	var brokerName string
	var brokerClient osb.Client
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, name, bClient, err := c.getClusterServiceClassAndClusterServiceBroker(instance)
		if err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		brokerName = name
		brokerClient = bClient
		prettyName = pretty.ClusterServiceClassName(serviceClass)
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, name, bClient, err := c.getServiceClassAndServiceBroker(instance)
		if err != nil {
			return c.handleServiceInstanceReconciliationError(instance, err)
		}
		brokerName = name
		brokerClient = bClient
		prettyName = pretty.ServiceClassName(serviceClass)
	}
	request, inProgressProperties, err := c.prepareDeprovisionRequest(instance)
	if err != nil {
		return c.handleServiceInstanceReconciliationError(instance, err)
	}
	if instance.DeletionTimestamp == nil {
		if instance.Status.OperationStartTime == nil {
			now := metav1.Now()
			instance.Status.OperationStartTime = &now
		}
	} else {
		if instance.Status.CurrentOperation != v1beta1.ServiceInstanceOperationDeprovision {
			if instance.Status.OrphanMitigationInProgress {
				removeServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionOrphanMitigation)
				instance.Status.OrphanMitigationInProgress = false
			}
			updatedInstance, err := c.recordStartOfServiceInstanceOperation(instance, v1beta1.ServiceInstanceOperationDeprovision, inProgressProperties)
			if err != nil {
				return err
			}
			if updatedInstance.ResourceVersion != instance.ResourceVersion {
				return nil
			}
			instance = updatedInstance
		}
	}
	klog.V(4).Info(pcb.Message("Sending deprovision request to broker"))
	response, err := brokerClient.DeprovisionInstance(request)
	if err != nil {
		msg := fmt.Sprintf(`Error deprovisioning, %s at ClusterServiceBroker %q: %v`, prettyName, brokerName, err)
		if httpErr, ok := osb.IsHTTPError(err); ok {
			msg = fmt.Sprintf("Deprovision call failed; received error response from broker: %v", httpErr)
		}
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionUnknown, errorDeprovisionCallFailedReason, msg)
		if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
			msg := "Stopping reconciliation retries because too much time has elapsed"
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
			return c.processDeprovisionFailure(instance, readyCond, failedCond)
		}
		return c.processServiceInstanceOperationError(instance, readyCond)
	}
	if response.Async {
		return c.processDeprovisionAsyncResponse(instance, response)
	}
	return c.processDeprovisionSuccess(instance)
}
func (c *controller) pollServiceInstance(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	klog.V(4).Info(pcb.Message("Processing poll event"))
	instance = instance.DeepCopy()
	var brokerClient osb.Client
	var err error
	if instance.Spec.ClusterServiceClassSpecified() {
		_, _, _, brokerClient, err = c.getClusterServiceClassPlanAndClusterServiceBroker(instance)
	} else {
		_, _, _, brokerClient, err = c.getServiceClassPlanAndServiceBroker(instance)
	}
	if err != nil {
		return c.handleServiceInstanceReconciliationError(instance, err)
	}
	mitigatingOrphan := instance.Status.OrphanMitigationInProgress
	provisioning := instance.Status.CurrentOperation == v1beta1.ServiceInstanceOperationProvision && !mitigatingOrphan
	deleting := instance.Status.CurrentOperation == v1beta1.ServiceInstanceOperationDeprovision || mitigatingOrphan
	request, err := c.prepareServiceInstanceLastOperationRequest(instance)
	if err != nil {
		return c.handleServiceInstanceReconciliationError(instance, err)
	}
	klog.V(5).Info(pcb.Message("Polling last operation"))
	response, err := brokerClient.PollLastOperation(request)
	if err != nil {
		if osb.IsGoneError(err) && deleting {
			if err := c.processDeprovisionSuccess(instance); err != nil {
				return c.handleServiceInstancePollingError(instance, err)
			}
			return c.finishPollingServiceInstance(instance)
		}
		reason := errorPollingLastOperationReason
		message := fmt.Sprintf("Error polling last operation: %v", err)
		klog.V(4).Info(pcb.Message(message))
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, message)
		if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
			return c.processServiceInstancePollingFailureRetryTimeout(instance, readyCond)
		}
		if httpErr, ok := osb.IsHTTPError(err); ok {
			if isRetriableHTTPStatus(httpErr.StatusCode) {
				return c.processServiceInstancePollingTemporaryFailure(instance, readyCond)
			}
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, reason, message)
			return c.processServiceInstancePollingTerminalFailure(instance, readyCond, failedCond)
		}
		return c.processServiceInstancePollingTemporaryFailure(instance, readyCond)
	}
	description := "(no description provided)"
	if response.Description != nil {
		description = *response.Description
	}
	klog.V(4).Info(pcb.Messagef("Poll returned %q : %q", response.State, description))
	switch response.State {
	case osb.StateInProgress:
		var message string
		var reason string
		switch {
		case deleting:
			reason = asyncDeprovisioningReason
			message = asyncDeprovisioningMessage
		case provisioning:
			reason = asyncProvisioningReason
			message = asyncProvisioningMessage
		default:
			reason = asyncUpdatingInstanceReason
			message = asyncUpdatingInstanceMessage
		}
		if response.Description != nil {
			message = fmt.Sprintf("%s (%s)", message, *response.Description)
		}
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, message)
		if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
			return c.processServiceInstancePollingFailureRetryTimeout(instance, readyCond)
		}
		if response.Description != nil {
			c.recorder.Event(instance, corev1.EventTypeNormal, readyCond.Reason, readyCond.Message)
			setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
			if _, err := c.updateServiceInstanceStatus(instance); err != nil {
				return c.handleServiceInstancePollingError(instance, err)
			}
		}
		klog.V(4).Info(pcb.Message("Last operation not completed (still in progress)"))
		return c.continuePollingServiceInstance(instance)
	case osb.StateSucceeded:
		var err error
		switch {
		case deleting:
			err = c.processDeprovisionSuccess(instance)
		case provisioning:
			err = c.processProvisionSuccess(instance, nil)
		default:
			err = c.processUpdateServiceInstanceSuccess(instance)
		}
		if err != nil {
			return c.handleServiceInstancePollingError(instance, err)
		}
		return c.finishPollingServiceInstance(instance)
	case osb.StateFailed:
		var err error
		switch {
		case deleting:
			msg := "Deprovision call failed: " + description
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionUnknown, errorDeprovisionCallFailedReason, msg)
			if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
				return c.processServiceInstancePollingFailureRetryTimeout(instance, readyCond)
			}
			clearServiceInstanceAsyncOsbOperation(instance)
			c.finishPollingServiceInstance(instance)
			return c.processServiceInstanceOperationError(instance, readyCond)
		case provisioning:
			reason := errorProvisionCallFailedReason
			message := "Provision call failed: " + description
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, message)
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, reason, message)
			err = c.processTerminalProvisionFailure(instance, readyCond, failedCond, true)
		default:
			reason := errorUpdateInstanceCallFailedReason
			message := "Update call failed: " + description
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, reason, message)
			failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, reason, message)
			err = c.processTerminalUpdateServiceInstanceFailure(instance, readyCond, failedCond)
		}
		if err != nil {
			return c.handleServiceInstancePollingError(instance, err)
		}
		return c.finishPollingServiceInstance(instance)
	default:
		message := pcb.Messagef("Got invalid state in LastOperationResponse: %q", response.State)
		klog.Warning(message)
		if c.reconciliationRetryDurationExceeded(instance.Status.OperationStartTime) {
			readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionUnknown, errorPollingLastOperationReason, message)
			return c.processServiceInstancePollingFailureRetryTimeout(instance, readyCond)
		}
		err := fmt.Errorf(`Got invalid state in LastOperationResponse: %q`, response.State)
		return c.handleServiceInstancePollingError(instance, err)
	}
}
func clearServiceInstanceAsyncOsbOperation(instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance.Status.AsyncOpInProgress = false
	instance.Status.LastOperation = nil
}
func isServiceInstanceProcessedAlready(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return instance.Status.ObservedGeneration >= instance.Generation && (isServiceInstanceReady(instance) || isServiceInstanceFailed(instance)) && !instance.Status.OrphanMitigationInProgress
}
func (c *controller) processServiceInstancePollingFailureRetryTimeout(instance *v1beta1.ServiceInstance, readyCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	msg := "Stopping reconciliation retries because too much time has elapsed"
	failedCond := newServiceInstanceFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
	return c.processServiceInstancePollingTerminalFailure(instance, readyCond, failedCond)
}
func (c *controller) processServiceInstancePollingTerminalFailure(instance *v1beta1.ServiceInstance, readyCond, failedCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mitigatingOrphan := instance.Status.OrphanMitigationInProgress
	provisioning := instance.Status.CurrentOperation == v1beta1.ServiceInstanceOperationProvision && !mitigatingOrphan
	deleting := instance.Status.CurrentOperation == v1beta1.ServiceInstanceOperationDeprovision || mitigatingOrphan
	var err error
	switch {
	case deleting:
		err = c.processDeprovisionFailure(instance, readyCond, failedCond)
	case provisioning:
		c.finishPollingServiceInstance(instance)
		return c.processTerminalProvisionFailure(instance, readyCond, failedCond, true)
	default:
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionFalse, failedCond.Reason, failedCond.Message)
		err = c.processTerminalUpdateServiceInstanceFailure(instance, readyCond, failedCond)
	}
	if err != nil {
		c.recorder.Event(instance, corev1.EventTypeWarning, failedCond.Reason, failedCond.Message)
		return c.handleServiceInstancePollingError(instance, err)
	}
	return c.finishPollingServiceInstance(instance)
}
func (c *controller) processServiceInstancePollingTemporaryFailure(instance *v1beta1.ServiceInstance, readyCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.recorder.Event(instance, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return c.handleServiceInstancePollingError(instance, err)
	}
	return fmt.Errorf(readyCond.Message)
}
func (c *controller) resolveReferences(instance *v1beta1.ServiceInstance) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Spec.ClusterServiceClassSpecified() {
		return c.resolveClusterReferences(instance)
	} else if instance.Spec.ServiceClassSpecified() {
		return c.resolveNamespacedReferences(instance)
	}
	return false, stderrors.New(errorAmbiguousPlanReferenceScope)
}
func (c *controller) resolveClusterReferences(instance *v1beta1.ServiceInstance) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Spec.ClusterServiceClassRef != nil && instance.Spec.ClusterServicePlanRef != nil {
		return false, nil
	}
	var sc *v1beta1.ClusterServiceClass
	var err error
	if instance.Spec.ClusterServiceClassRef == nil {
		sc, err = c.resolveClusterServiceClassRef(instance)
		if err != nil {
			pcb := pretty.NewInstanceContextBuilder(instance)
			klog.Warning(pcb.Message(err.Error()))
			updatedInstance, _ := c.updateServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServiceClassReason, "The instance references a ClusterServiceClass that does not exist. "+err.Error())
			c.recorder.Event(instance, corev1.EventTypeWarning, errorNonexistentClusterServiceClassReason, err.Error())
			return updatedInstance.ResourceVersion != instance.ResourceVersion, err
		}
	}
	if instance.Spec.ClusterServicePlanRef == nil {
		if sc == nil {
			sc, err = c.clusterServiceClassLister.Get(instance.Spec.ClusterServiceClassRef.Name)
			if err != nil {
				return false, fmt.Errorf(`couldn't find ClusterServiceClass "(K8S: %s)": %v`, instance.Spec.ClusterServiceClassRef.Name, err.Error())
			}
		}
		err = c.resolveClusterServicePlanRef(instance, sc.Spec.ClusterServiceBrokerName)
		if err != nil {
			pcb := pretty.NewInstanceContextBuilder(instance)
			klog.Warning(pcb.Message(err.Error()))
			updatedInstance, _ := c.updateServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServicePlanReason, "The instance references a ClusterServicePlan that does not exist. "+err.Error())
			c.recorder.Event(instance, corev1.EventTypeWarning, errorNonexistentClusterServicePlanReason, err.Error())
			return updatedInstance.ResourceVersion != instance.ResourceVersion, err
		}
	}
	updatedInstance, err := c.updateServiceInstanceReferences(instance)
	return updatedInstance.ResourceVersion != instance.ResourceVersion, err
}
func (c *controller) resolveNamespacedReferences(instance *v1beta1.ServiceInstance) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Spec.ServiceClassRef != nil && instance.Spec.ServicePlanRef != nil {
		return false, nil
	}
	var sc *v1beta1.ServiceClass
	var err error
	if instance.Spec.ServiceClassRef == nil {
		sc, err = c.resolveServiceClassRef(instance)
		if err != nil {
			pcb := pretty.NewInstanceContextBuilder(instance)
			klog.Warning(pcb.Message(err.Error()))
			updatedInstance, _ := c.updateServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, errorNonexistentServiceClassReason, "The instance references a ServiceClass that does not exist. "+err.Error())
			c.recorder.Event(instance, corev1.EventTypeWarning, errorNonexistentServiceClassReason, err.Error())
			return updatedInstance.ResourceVersion != instance.ResourceVersion, err
		}
	}
	if instance.Spec.ServicePlanRef == nil {
		if sc == nil {
			sc, err = c.serviceClassLister.ServiceClasses(instance.Namespace).Get(instance.Spec.ServiceClassRef.Name)
			if err != nil {
				return false, fmt.Errorf(`couldn't find ServiceClass "(K8S: %s)": %v`, instance.Spec.ServiceClassRef.Name, err.Error())
			}
		}
		err = c.resolveServicePlanRef(instance, sc.Spec.ServiceBrokerName)
		if err != nil {
			pcb := pretty.NewInstanceContextBuilder(instance)
			klog.Warning(pcb.Message(err.Error()))
			updatedInstance, _ := c.updateServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, errorNonexistentServicePlanReason, "The instance references a ServicePlan that does not exist. "+err.Error())
			c.recorder.Event(instance, corev1.EventTypeWarning, errorNonexistentServicePlanReason, err.Error())
			return updatedInstance.ResourceVersion != instance.ResourceVersion, err
		}
	}
	updatedInstance, err := c.updateServiceInstanceReferences(instance)
	return updatedInstance.ResourceVersion != instance.ResourceVersion, err
}
func (c *controller) resolveClusterServiceClassRef(instance *v1beta1.ServiceInstance) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !instance.Spec.ClusterServiceClassSpecified() {
		return nil, fmt.Errorf("ServiceInstance %s/%s is in invalid state, neither ClusterServiceClassExternalName, ClusterServiceClassExternalID, nor ClusterServiceClassName is set", instance.Namespace, instance.Name)
	}
	pcb := pretty.NewInstanceContextBuilder(instance)
	var sc *v1beta1.ClusterServiceClass
	if instance.Spec.ClusterServiceClassName != "" {
		klog.V(4).Info(pcb.Messagef("looking up a ClusterServiceClass from K8S Name: %q", instance.Spec.ClusterServiceClassName))
		var err error
		sc, err = c.clusterServiceClassLister.Get(instance.Spec.ClusterServiceClassName)
		if err == nil {
			instance.Spec.ClusterServiceClassRef = &v1beta1.ClusterObjectReference{Name: sc.Name}
			klog.V(4).Info(pcb.Messagef("resolved ClusterServiceClass %c to ClusterServiceClass with external Name %q", instance.Spec.PlanReference, sc.Spec.ExternalName))
		} else {
			return nil, fmt.Errorf("References a non-existent ClusterServiceClass %c", instance.Spec.PlanReference)
		}
	} else {
		filterField := instance.Spec.GetClusterServiceClassFilterFieldName()
		filterValue := instance.Spec.GetSpecifiedClusterServiceClass()
		klog.V(4).Info(pcb.Messagef("looking up a ClusterServiceClass from %s: %q", filterField, filterValue))
		listOpts := metav1.ListOptions{FieldSelector: fields.OneTermEqualSelector(filterField, filterValue).String()}
		serviceClasses, err := c.serviceCatalogClient.ClusterServiceClasses().List(listOpts)
		if err == nil && len(serviceClasses.Items) == 1 {
			sc = &serviceClasses.Items[0]
			instance.Spec.ClusterServiceClassRef = &v1beta1.ClusterObjectReference{Name: sc.Name}
			klog.V(4).Info(pcb.Messagef("resolved %c to K8S ClusterServiceClass %q", instance.Spec.PlanReference, sc.Name))
		} else {
			return nil, fmt.Errorf("References a non-existent ClusterServiceClass %c or there is more than one (found: %d)", instance.Spec.PlanReference, len(serviceClasses.Items))
		}
	}
	return sc, nil
}
func (c *controller) resolveServiceClassRef(instance *v1beta1.ServiceInstance) (*v1beta1.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !instance.Spec.ServiceClassSpecified() {
		return nil, fmt.Errorf("ServiceInstance %s/%s is in invalid state, neither ServiceClassExternalName, ServiceClassExternalID, nor ServiceClassName is set", instance.Namespace, instance.Name)
	}
	pcb := pretty.NewInstanceContextBuilder(instance)
	var sc *v1beta1.ServiceClass
	if instance.Spec.ServiceClassName != "" {
		klog.V(4).Info(pcb.Messagef("looking up a ServiceClass from K8S Name: %q", instance.Spec.ServiceClassName))
		var err error
		sc, err = c.serviceClassLister.ServiceClasses(instance.Namespace).Get(instance.Spec.ServiceClassName)
		if err == nil {
			instance.Spec.ServiceClassRef = &v1beta1.LocalObjectReference{Name: sc.Name}
			klog.V(4).Info(pcb.Messagef("resolved ServiceClass %c to ServiceClass with external Name %q", instance.Spec.PlanReference, sc.Spec.ExternalName))
		} else {
			return nil, fmt.Errorf("References a non-existent ServiceClass %c", instance.Spec.PlanReference)
		}
	} else {
		filterField := instance.Spec.GetServiceClassFilterFieldName()
		filterValue := instance.Spec.GetSpecifiedServiceClass()
		klog.V(4).Info(pcb.Messagef("looking up a ServiceClass from %s: %q", filterField, filterValue))
		listOpts := metav1.ListOptions{FieldSelector: fields.OneTermEqualSelector(filterField, filterValue).String()}
		serviceClasses, err := c.serviceCatalogClient.ServiceClasses(instance.Namespace).List(listOpts)
		if err == nil && len(serviceClasses.Items) == 1 {
			sc = &serviceClasses.Items[0]
			instance.Spec.ServiceClassRef = &v1beta1.LocalObjectReference{Name: sc.Name}
			klog.V(4).Info(pcb.Messagef("resolved %c to K8S ServiceClass %q", instance.Spec.PlanReference, sc.Name))
		} else {
			return nil, fmt.Errorf("References a non-existent ServiceClass %c or there is more than one (found: %d)", instance.Spec.PlanReference, len(serviceClasses.Items))
		}
	}
	return sc, nil
}
func (c *controller) resolveClusterServicePlanRef(instance *v1beta1.ServiceInstance, brokerName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !instance.Spec.ClusterServicePlanSpecified() {
		return fmt.Errorf("ServiceInstance %s/%s is in invalid state, neither ClusterServicePlanExternalName, ClusterServicePlanExternalID, nor ClusterServicePlanName is set", instance.Namespace, instance.Name)
	}
	pcb := pretty.NewInstanceContextBuilder(instance)
	if instance.Spec.ClusterServicePlanName != "" {
		sp, err := c.clusterServicePlanLister.Get(instance.Spec.ClusterServicePlanName)
		if err == nil {
			instance.Spec.ClusterServicePlanRef = &v1beta1.ClusterObjectReference{Name: sp.Name}
			klog.V(4).Info(pcb.Messagef("resolved ClusterServicePlan with K8S name %q to ClusterServicePlan with external name %q", instance.Spec.ClusterServicePlanName, sp.Spec.ExternalName))
		} else {
			return fmt.Errorf("References a non-existent ClusterServicePlan %v", instance.Spec.PlanReference)
		}
	} else {
		fieldSet := fields.Set{instance.Spec.GetClusterServicePlanFilterFieldName(): instance.Spec.GetSpecifiedClusterServicePlan(), "spec.clusterServiceClassRef.name": instance.Spec.ClusterServiceClassRef.Name, "spec.clusterServiceBrokerName": brokerName}
		fieldSelector := fields.SelectorFromSet(fieldSet).String()
		listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
		servicePlans, err := c.serviceCatalogClient.ClusterServicePlans().List(listOpts)
		if err == nil && len(servicePlans.Items) == 1 {
			sp := &servicePlans.Items[0]
			instance.Spec.ClusterServicePlanRef = &v1beta1.ClusterObjectReference{Name: sp.Name}
			klog.V(4).Info(pcb.Messagef("resolved %v to ClusterServicePlan (K8S: %q)", instance.Spec.PlanReference, sp.Name))
		} else {
			return fmt.Errorf("References a non-existent ClusterServicePlan %b on ClusterServiceClass %s %c or there is more than one (found: %d)", instance.Spec.PlanReference, instance.Spec.ClusterServiceClassRef.Name, instance.Spec.PlanReference, len(servicePlans.Items))
		}
	}
	return nil
}
func (c *controller) resolveServicePlanRef(instance *v1beta1.ServiceInstance, brokerName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !instance.Spec.ServicePlanSpecified() {
		return fmt.Errorf("ServiceInstance %s/%s is in invalid state, neither ServicePlanExternalName, ServicePlanExternalID, nor ServicePlanName is set", instance.Namespace, instance.Name)
	}
	pcb := pretty.NewInstanceContextBuilder(instance)
	if instance.Spec.ServicePlanName != "" {
		sp, err := c.servicePlanLister.ServicePlans(instance.Namespace).Get(instance.Spec.ServicePlanName)
		if err == nil {
			instance.Spec.ServicePlanRef = &v1beta1.LocalObjectReference{Name: sp.Name}
			klog.V(4).Info(pcb.Messagef("resolved ServicePlan with K8S name %q to ServicePlan with external name %q", instance.Spec.ServicePlanName, sp.Spec.ExternalName))
		} else {
			return fmt.Errorf("References a non-existent ServicePlan %v", instance.Spec.PlanReference)
		}
	} else {
		fieldSet := fields.Set{instance.Spec.GetServicePlanFilterFieldName(): instance.Spec.GetSpecifiedServicePlan(), "spec.serviceClassRef.name": instance.Spec.ServiceClassRef.Name, "spec.serviceBrokerName": brokerName}
		fieldSelector := fields.SelectorFromSet(fieldSet).String()
		listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
		servicePlans, err := c.serviceCatalogClient.ServicePlans(instance.Namespace).List(listOpts)
		if err == nil && len(servicePlans.Items) == 1 {
			sp := &servicePlans.Items[0]
			instance.Spec.ServicePlanRef = &v1beta1.LocalObjectReference{Name: sp.Name}
			klog.V(4).Info(pcb.Messagef("resolved %v to ServicePlan (K8S: %q)", instance.Spec.PlanReference, sp.Name))
		} else {
			return fmt.Errorf("References a non-existent ServicePlan %b on ServiceClass %s %c or there is more than one (found: %d)", instance.Spec.PlanReference, instance.Spec.ServiceClassRef.Name, instance.Spec.PlanReference, len(servicePlans.Items))
		}
	}
	return nil
}
func (c *controller) applyDefaultProvisioningParameters(instance *v1beta1.ServiceInstance) (bool, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Status.DefaultProvisionParameters != nil {
		return false, nil
	}
	defaultParams, err := c.getDefaultProvisioningParameters(instance)
	if err != nil {
		return false, err
	}
	finalParams, err := mergeParameters(instance.Spec.Parameters, defaultParams)
	if err != nil {
		return false, err
	}
	if instance.Spec.Parameters == finalParams {
		return false, nil
	}
	pcb := pretty.NewContextBuilder(pretty.ServiceInstance, instance.Namespace, instance.Name, "")
	klog.V(4).Info(pcb.Message("Applying default provisioning parameters"))
	instance.Spec.Parameters = finalParams
	_, err = c.updateServiceInstanceWithRetries(instance, func(conflictedInstance *v1beta1.ServiceInstance) {
		conflictedInstance.Spec.Parameters = finalParams
	})
	if err != nil {
		s := fmt.Sprintf("error updating service instance to apply default parameters: %s", err)
		klog.Warning(pcb.Message(s))
		c.recorder.Event(instance, corev1.EventTypeWarning, errorWithParametersReason, s)
		return false, fmt.Errorf(s)
	}
	instance.Status.DefaultProvisionParameters = defaultParams
	updatedInstance, err := c.updateServiceInstanceStatus(instance)
	return updatedInstance.ResourceVersion != instance.ResourceVersion, err
}
func (c *controller) getDefaultProvisioningParameters(instance *v1beta1.ServiceInstance) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var classDefaults, planDefaults *runtime.RawExtension
	if instance.Spec.ClusterServiceClassSpecified() {
		class, err := c.clusterServiceClassLister.Get(instance.Spec.ClusterServiceClassRef.Name)
		if err != nil {
			return nil, err
		}
		classDefaults = class.Spec.DefaultProvisionParameters
	} else if instance.Spec.ServiceClassSpecified() {
		class, err := c.serviceClassLister.ServiceClasses(instance.Namespace).Get(instance.Spec.ServiceClassRef.Name)
		if err != nil {
			return nil, err
		}
		classDefaults = class.Spec.DefaultProvisionParameters
	} else {
		return nil, fmt.Errorf("invalid class reference %v", instance.Spec.PlanReference)
	}
	if instance.Spec.ClusterServicePlanSpecified() {
		plan, err := c.clusterServicePlanLister.Get(instance.Spec.ClusterServicePlanRef.Name)
		if err != nil {
			return nil, err
		}
		planDefaults = plan.Spec.DefaultProvisionParameters
	} else if instance.Spec.ServicePlanSpecified() {
		plan, err := c.servicePlanLister.ServicePlans(instance.Namespace).Get(instance.Spec.ServicePlanRef.Name)
		if err != nil {
			return nil, err
		}
		planDefaults = plan.Spec.DefaultProvisionParameters
	} else {
		return nil, fmt.Errorf("invalid plan reference %v", instance.Spec.PlanReference)
	}
	return mergeParameters(planDefaults, classDefaults)
}
func (c *controller) prepareProvisionRequest(instance *v1beta1.ServiceInstance) (*osb.ProvisionRequest, *v1beta1.ServiceInstancePropertiesState, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, servicePlan, _, _, err := c.getClusterServiceClassPlanAndClusterServiceBroker(instance)
		if err != nil {
			return nil, nil, err
		}
		if err = c.checkForRemovedClusterClassAndPlan(instance, serviceClass, servicePlan); err != nil {
			return nil, nil, err
		}
		request, inProgressProperties, err := c.innerPrepareProvisionRequest(instance, serviceClass.Spec.CommonServiceClassSpec, servicePlan.Spec.CommonServicePlanSpec)
		if err != nil {
			return nil, nil, err
		}
		return request, inProgressProperties, nil
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, servicePlan, _, _, err := c.getServiceClassPlanAndServiceBroker(instance)
		if err != nil {
			return nil, nil, err
		}
		if err = c.checkForRemovedClassAndPlan(instance, serviceClass, servicePlan); err != nil {
			return nil, nil, err
		}
		request, inProgressProperties, err := c.innerPrepareProvisionRequest(instance, serviceClass.Spec.CommonServiceClassSpec, servicePlan.Spec.CommonServicePlanSpec)
		if err != nil {
			return nil, nil, err
		}
		return request, inProgressProperties, nil
	}
	return nil, nil, stderrors.New(errorAmbiguousPlanReferenceScope)
}
func newServiceInstanceCondition(status v1beta1.ConditionStatus, condType v1beta1.ServiceInstanceConditionType, reason, message string) *v1beta1.ServiceInstanceCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v1beta1.ServiceInstanceCondition{Type: condType, Status: status, Reason: reason, Message: message, LastTransitionTime: metav1.Now()}
}
func newServiceInstanceReadyCondition(status v1beta1.ConditionStatus, reason, message string) *v1beta1.ServiceInstanceCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceInstanceCondition(status, v1beta1.ServiceInstanceConditionReady, reason, message)
}
func newServiceInstanceFailedCondition(status v1beta1.ConditionStatus, reason, message string) *v1beta1.ServiceInstanceCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newServiceInstanceCondition(status, v1beta1.ServiceInstanceConditionFailed, reason, message)
}
func removeServiceInstanceCondition(toUpdate *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(toUpdate)
	klog.V(5).Info(pcb.Messagef("Removing condition %q", conditionType))
	newStatusConditions := make([]v1beta1.ServiceInstanceCondition, 0, len(toUpdate.Status.Conditions))
	for _, cond := range toUpdate.Status.Conditions {
		if cond.Type == conditionType {
			klog.V(5).Info(pcb.Messagef("Found existing condition %q: %q; removing it", conditionType, cond.Status))
			continue
		}
		newStatusConditions = append(newStatusConditions, cond)
	}
	toUpdate.Status.Conditions = newStatusConditions
}
func setServiceInstanceCondition(toUpdate *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType, status v1beta1.ConditionStatus, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceConditionInternal(toUpdate, conditionType, status, reason, message, metav1.Now())
}
func setServiceInstanceConditionInternal(toUpdate *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType, status v1beta1.ConditionStatus, reason, message string, t metav1.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(toUpdate)
	klog.Info(pcb.Message(message))
	klog.V(5).Info(pcb.Messagef("Setting condition %q to %v", conditionType, status))
	newCondition := v1beta1.ServiceInstanceCondition{Type: conditionType, Status: status, Reason: reason, Message: message}
	if len(toUpdate.Status.Conditions) == 0 {
		klog.V(3).Info(pcb.Messagef("Setting lastTransitionTime, condition %q to %v", conditionType, t))
		newCondition.LastTransitionTime = t
		toUpdate.Status.Conditions = []v1beta1.ServiceInstanceCondition{newCondition}
		return
	}
	for i, cond := range toUpdate.Status.Conditions {
		if cond.Type == conditionType {
			if cond.Status != newCondition.Status {
				klog.V(3).Info(pcb.Messagef("Found status change, condition %q: %q -> %q; setting lastTransitionTime to %v", conditionType, cond.Status, status, t))
				newCondition.LastTransitionTime = t
			} else {
				newCondition.LastTransitionTime = cond.LastTransitionTime
			}
			toUpdate.Status.Conditions[i] = newCondition
			return
		}
	}
	klog.V(3).Info(pcb.Messagef("Setting lastTransitionTime, condition %q to %v", conditionType, t))
	newCondition.LastTransitionTime = t
	toUpdate.Status.Conditions = append(toUpdate.Status.Conditions, newCondition)
}
func (c *controller) updateServiceInstanceReferences(toUpdate *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(toUpdate)
	klog.V(4).Info(pcb.Message("Updating references"))
	status := toUpdate.Status
	updatedInstance, err := c.serviceCatalogClient.ServiceInstances(toUpdate.Namespace).UpdateReferences(toUpdate)
	if err != nil {
		klog.Errorf(pcb.Messagef("Failed to update references: %v", err))
	}
	updatedInstance.Status = status
	return updatedInstance, err
}
func (c *controller) updateServiceInstanceWithRetries(instance *v1beta1.ServiceInstance, conflictResolutionFunc func(*v1beta1.ServiceInstance)) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	const interval = 100 * time.Millisecond
	const timeout = 10 * time.Second
	var updatedInstance *v1beta1.ServiceInstance
	instanceToUpdate := instance
	err := wait.PollImmediate(interval, timeout, func() (bool, error) {
		klog.V(4).Info(pcb.Message("Updating instance"))
		upd, err := c.serviceCatalogClient.ServiceInstances(instanceToUpdate.Namespace).Update(instanceToUpdate)
		if err != nil {
			if !errors.IsConflict(err) {
				return false, err
			}
			klog.V(4).Info(pcb.Message("Couldn't update instance because the resource was stale"))
			instanceToUpdate, err = c.serviceCatalogClient.ServiceInstances(instance.Namespace).Get(instance.Name, metav1.GetOptions{})
			if err != nil {
				return false, err
			}
			conflictResolutionFunc(instanceToUpdate)
			return false, nil
		}
		updatedInstance = upd
		return true, nil
	})
	if err != nil {
		klog.Errorf(pcb.Messagef("Failed to update instance: %v", err))
	}
	return updatedInstance, err
}
func (c *controller) updateServiceInstanceStatus(instance *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.updateServiceInstanceStatusWithRetries(instance, nil)
}
func (c *controller) updateServiceInstanceStatusWithRetries(instance *v1beta1.ServiceInstance, postConflictUpdateFunc func(*v1beta1.ServiceInstance)) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	const interval = 100 * time.Millisecond
	const timeout = 10 * time.Second
	var updatedInstance *v1beta1.ServiceInstance
	instanceToUpdate := instance
	err := wait.PollImmediate(interval, timeout, func() (bool, error) {
		klog.V(4).Info(pcb.Message("Updating status"))
		upd, err := c.serviceCatalogClient.ServiceInstances(instanceToUpdate.Namespace).UpdateStatus(instanceToUpdate)
		if err != nil {
			if !errors.IsConflict(err) {
				return false, err
			}
			klog.V(4).Info(pcb.Message("Couldn't update status because the resource was stale"))
			instanceToUpdate, err = c.serviceCatalogClient.ServiceInstances(instance.Namespace).Get(instance.Name, metav1.GetOptions{})
			if err != nil {
				return false, err
			}
			instanceToUpdate.Status = instance.Status
			if postConflictUpdateFunc != nil {
				postConflictUpdateFunc(instanceToUpdate)
			}
			return false, nil
		}
		updatedInstance = upd
		return true, nil
	})
	if err != nil {
		klog.Errorf(pcb.Messagef("Failed to update status: %v", err))
	}
	return updatedInstance, err
}
func (c *controller) updateServiceInstanceCondition(instance *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType, status v1beta1.ConditionStatus, reason, message string) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	toUpdate := instance.DeepCopy()
	setServiceInstanceCondition(toUpdate, conditionType, status, reason, message)
	klog.V(4).Info(pcb.Messagef("Updating %v condition to %v", conditionType, status))
	updatedInstance, err := c.serviceCatalogClient.ServiceInstances(instance.Namespace).UpdateStatus(toUpdate)
	if err != nil {
		klog.Errorf(pcb.Messagef("Failed to update condition %v to true: %v", conditionType, err))
	}
	return updatedInstance, err
}
func (c *controller) prepareObservedGeneration(toUpdate *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	toUpdate.Status.ObservedGeneration = toUpdate.Generation
	removeServiceInstanceCondition(toUpdate, v1beta1.ServiceInstanceConditionFailed)
}
func isServiceInstancePropertiesStateEqual(s1 *v1beta1.ServiceInstancePropertiesState, s2 *v1beta1.ServiceInstancePropertiesState) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s1 == nil && s2 == nil {
		return true
	}
	if (s1 == nil && s2 != nil) || (s1 != nil && s2 == nil) {
		return false
	}
	if s1.ClusterServicePlanExternalID != s2.ClusterServicePlanExternalID {
		return false
	}
	if s1.ClusterServicePlanExternalName != s2.ClusterServicePlanExternalName {
		return false
	}
	if s1.ParameterChecksum != s2.ParameterChecksum {
		return false
	}
	if s1.UserInfo != nil || s2.UserInfo != nil {
		u1 := s1.UserInfo
		u2 := s2.UserInfo
		if (u1 == nil && u2 != nil) || (u1 != nil && u2 == nil) {
			return false
		}
		if u1.UID != u2.UID {
			return false
		}
	}
	return true
}
func (c *controller) recordStartOfServiceInstanceOperation(toUpdate *v1beta1.ServiceInstance, operation v1beta1.ServiceInstanceOperation, inProgressProperties *v1beta1.ServiceInstancePropertiesState) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clearServiceInstanceCurrentOperation(toUpdate)
	toUpdate.Status.CurrentOperation = operation
	now := metav1.Now()
	toUpdate.Status.OperationStartTime = &now
	toUpdate.Status.InProgressProperties = inProgressProperties
	reason := ""
	message := ""
	switch operation {
	case v1beta1.ServiceInstanceOperationProvision:
		reason = provisioningInFlightReason
		message = provisioningInFlightMessage
		toUpdate.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusRequired
	case v1beta1.ServiceInstanceOperationUpdate:
		reason = instanceUpdatingInFlightReason
		message = instanceUpdatingInFlightMessage
	case v1beta1.ServiceInstanceOperationDeprovision:
		reason = deprovisioningInFlightReason
		message = deprovisioningInFlightMessage
	}
	setServiceInstanceCondition(toUpdate, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, reason, message)
	c.resetPollingRateLimiterForServiceInstance(toUpdate)
	return c.updateServiceInstanceStatus(toUpdate)
}
func (c *controller) checkForRemovedClusterClassAndPlan(instance *v1beta1.ServiceInstance, serviceClass *v1beta1.ClusterServiceClass, servicePlan *v1beta1.ClusterServicePlan) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	classDeleted := serviceClass.Status.RemovedFromBrokerCatalog
	planDeleted := servicePlan.Status.RemovedFromBrokerCatalog
	if !classDeleted && !planDeleted {
		return nil
	}
	isProvisioning := instance.Status.ProvisionStatus != v1beta1.ServiceInstanceProvisionStatusProvisioned
	if !isProvisioning && instance.Status.ExternalProperties != nil && servicePlan.Spec.ExternalID == instance.Status.ExternalProperties.ClusterServicePlanExternalID {
		return nil
	}
	if planDeleted {
		return &operationError{reason: errorDeletedClusterServicePlanReason, message: fmt.Sprintf("%s has been deleted; cannot provision.", pretty.ClusterServicePlanName(servicePlan))}
	}
	return &operationError{reason: errorDeletedClusterServiceClassReason, message: fmt.Sprintf("%s has been deleted; cannot provision.", pretty.ClusterServiceClassName(serviceClass))}
}
func (c *controller) checkForRemovedClassAndPlan(instance *v1beta1.ServiceInstance, serviceClass *v1beta1.ServiceClass, servicePlan *v1beta1.ServicePlan) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	classDeleted := serviceClass.Status.RemovedFromBrokerCatalog
	planDeleted := servicePlan.Status.RemovedFromBrokerCatalog
	if !classDeleted && !planDeleted {
		return nil
	}
	isProvisioning := instance.Status.ProvisionStatus != v1beta1.ServiceInstanceProvisionStatusProvisioned
	if !isProvisioning && instance.Status.ExternalProperties != nil && servicePlan.Spec.ExternalID == instance.Status.ExternalProperties.ServicePlanExternalID {
		return nil
	}
	if planDeleted {
		return &operationError{reason: errorDeletedServicePlanReason, message: fmt.Sprintf("%s has been deleted; cannot provision.", pretty.ServicePlanName(servicePlan))}
	}
	return &operationError{reason: errorDeletedServiceClassReason, message: fmt.Sprintf("%s has been deleted; cannot provision.", pretty.ServiceClassName(serviceClass))}
}
func clearServiceInstanceCurrentOperation(toUpdate *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	toUpdate.Status.CurrentOperation = ""
	toUpdate.Status.OperationStartTime = nil
	toUpdate.Status.AsyncOpInProgress = false
	toUpdate.Status.LastOperation = nil
	toUpdate.Status.InProgressProperties = nil
}
func (c *controller) checkServiceInstanceHasExistingBindings(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bindingLister := c.bindingLister.ServiceBindings(instance.Namespace)
	selector := labels.NewSelector()
	bindingList, err := bindingLister.List(selector)
	if err != nil {
		return err
	}
	for _, binding := range bindingList {
		if instance.Name == binding.Spec.InstanceRef.Name {
			return &operationError{reason: errorDeprovisionBlockedByCredentialsReason, message: "All associated ServiceBindings must be removed before this ServiceInstance can be deleted"}
		}
	}
	return nil
}

type requestHelper struct {
	ns			*corev1.Namespace
	parameters		map[string]interface{}
	inProgressProperties	*v1beta1.ServiceInstancePropertiesState
	originatingIdentity	*osb.OriginatingIdentity
	requestContext		map[string]interface{}
}

func (c *controller) prepareRequestHelper(instance *v1beta1.ServiceInstance, planName string, planID string, setInProgressProperties bool) (*requestHelper, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rh := &requestHelper{}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		originatingIdentity, err := buildOriginatingIdentity(instance.Spec.UserInfo)
		if err != nil {
			return nil, &operationError{reason: errorWithOriginatingIdentityReason, message: fmt.Sprintf("Error building originating identity headers: %v", err)}
		}
		rh.originatingIdentity = originatingIdentity
	}
	reconciliationAction := getReconciliationActionForServiceInstance(instance)
	if reconciliationAction == reconcileDelete || reconciliationAction == reconcilePoll {
		return rh, nil
	}
	ns, err := c.kubeClient.CoreV1().Namespaces().Get(instance.Namespace, metav1.GetOptions{})
	if err != nil {
		return nil, &operationError{reason: errorFindingNamespaceServiceInstanceReason, message: fmt.Sprintf("Failed to get namespace %q: %s", instance.Namespace, err)}
	}
	rh.ns = ns
	if setInProgressProperties {
		parameters, parametersChecksum, rawParametersWithRedaction, err := prepareInProgressPropertyParameters(c.kubeClient, instance.Namespace, instance.Spec.Parameters, instance.Spec.ParametersFrom)
		if err != nil {
			return nil, &operationError{reason: errorWithParametersReason, message: err.Error()}
		}
		rh.parameters = parameters
		rh.inProgressProperties = &v1beta1.ServiceInstancePropertiesState{Parameters: rawParametersWithRedaction, ParameterChecksum: parametersChecksum, UserInfo: instance.Spec.UserInfo}
		if instance.Spec.ClusterServiceClassSpecified() {
			rh.inProgressProperties.ClusterServicePlanExternalName = planName
			rh.inProgressProperties.ClusterServicePlanExternalID = planID
		} else {
			rh.inProgressProperties.ServicePlanExternalName = planName
			rh.inProgressProperties.ServicePlanExternalID = planID
		}
	}
	id := c.getClusterID()
	rh.requestContext = map[string]interface{}{"platform": ContextProfilePlatformKubernetes, "namespace": instance.Namespace, clusterIdentifierKey: id}
	return rh, nil
}
func (c *controller) innerPrepareProvisionRequest(instance *v1beta1.ServiceInstance, classCommon v1beta1.CommonServiceClassSpec, planCommon v1beta1.CommonServicePlanSpec) (*osb.ProvisionRequest, *v1beta1.ServiceInstancePropertiesState, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rh, err := c.prepareRequestHelper(instance, planCommon.ExternalName, planCommon.ExternalID, true)
	if err != nil {
		return nil, nil, err
	}
	request := &osb.ProvisionRequest{AcceptsIncomplete: true, InstanceID: instance.Spec.ExternalID, ServiceID: classCommon.ExternalID, PlanID: planCommon.ExternalID, Parameters: rh.parameters, OrganizationGUID: c.getClusterID(), SpaceGUID: string(rh.ns.UID), Context: rh.requestContext, OriginatingIdentity: rh.originatingIdentity}
	return request, rh.inProgressProperties, nil
}
func (c *controller) prepareUpdateInstanceRequest(instance *v1beta1.ServiceInstance) (*osb.UpdateInstanceRequest, *v1beta1.ServiceInstancePropertiesState, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var rh *requestHelper
	var request *osb.UpdateInstanceRequest
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, servicePlan, _, _, err := c.getClusterServiceClassPlanAndClusterServiceBroker(instance)
		if err != nil {
			return nil, nil, c.handleServiceInstanceReconciliationError(instance, err)
		}
		rh, err = c.prepareRequestHelper(instance, servicePlan.Spec.ExternalName, servicePlan.Spec.ExternalID, true)
		if err != nil {
			return nil, nil, err
		}
		request = &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: instance.Spec.ExternalID, ServiceID: serviceClass.Spec.ExternalID, Context: rh.requestContext, OriginatingIdentity: rh.originatingIdentity}
		if instance.Status.ExternalProperties == nil || servicePlan.Spec.ExternalID != instance.Status.ExternalProperties.ClusterServicePlanExternalID {
			planID := servicePlan.Spec.ExternalID
			request.PlanID = &planID
		}
		if instance.Status.ExternalProperties == nil || rh.inProgressProperties.ParameterChecksum != instance.Status.ExternalProperties.ParameterChecksum {
			if rh.parameters != nil {
				request.Parameters = rh.parameters
			} else {
				request.Parameters = make(map[string]interface{})
			}
		}
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, servicePlan, _, _, err := c.getServiceClassPlanAndServiceBroker(instance)
		if err != nil {
			return nil, nil, c.handleServiceInstanceReconciliationError(instance, err)
		}
		rh, err = c.prepareRequestHelper(instance, servicePlan.Spec.ExternalName, servicePlan.Spec.ExternalID, true)
		if err != nil {
			return nil, nil, err
		}
		request = &osb.UpdateInstanceRequest{AcceptsIncomplete: true, InstanceID: instance.Spec.ExternalID, ServiceID: serviceClass.Spec.ExternalID, Context: rh.requestContext, OriginatingIdentity: rh.originatingIdentity}
		if instance.Status.ExternalProperties == nil || servicePlan.Spec.ExternalID != instance.Status.ExternalProperties.ServicePlanExternalID {
			planID := servicePlan.Spec.ExternalID
			request.PlanID = &planID
		}
		if instance.Status.ExternalProperties == nil || rh.inProgressProperties.ParameterChecksum != instance.Status.ExternalProperties.ParameterChecksum {
			if rh.parameters != nil {
				request.Parameters = rh.parameters
			} else {
				request.Parameters = make(map[string]interface{})
			}
		}
	}
	return request, rh.inProgressProperties, nil
}
func (c *controller) prepareDeprovisionRequest(instance *v1beta1.ServiceInstance) (*osb.DeprovisionRequest, *v1beta1.ServiceInstancePropertiesState, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rh, err := c.prepareRequestHelper(instance, "", "", true)
	if err != nil {
		return nil, nil, err
	}
	var scExternalID string
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, _, _, err := c.getClusterServiceClassAndClusterServiceBroker(instance)
		if err != nil {
			return nil, nil, c.handleServiceInstanceReconciliationError(instance, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, _, _, err := c.getServiceClassAndServiceBroker(instance)
		if err != nil {
			return nil, nil, c.handleServiceInstanceReconciliationError(instance, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
	}
	if instance.Status.CurrentOperation != "" || instance.Status.OrphanMitigationInProgress {
		if instance.Status.InProgressProperties == nil {
			return nil, nil, stderrors.New("InProgressProperties must be set when there is an operation or orphan mitigation in progress")
		}
		rh.inProgressProperties = instance.Status.InProgressProperties
	} else if instance.Status.ProvisionStatus != v1beta1.ServiceInstanceProvisionStatusProvisioned {
		if instance.Spec.ClusterServiceClassSpecified() {
			servicePlan, err := c.clusterServicePlanLister.Get(instance.Spec.ClusterServicePlanRef.Name)
			if err != nil {
				return nil, nil, &operationError{reason: errorNonexistentClusterServicePlanReason, message: fmt.Sprintf("The instance references a non-existent ClusterServicePlan %q - %v", instance.Spec.ClusterServicePlanRef.Name, instance.Spec.PlanReference)}
			}
			rh.inProgressProperties = &v1beta1.ServiceInstancePropertiesState{ClusterServicePlanExternalName: servicePlan.Spec.ExternalName, ClusterServicePlanExternalID: servicePlan.Spec.ExternalID}
		} else {
			servicePlan, err := c.servicePlanLister.ServicePlans(instance.Namespace).Get(instance.Spec.ServicePlanRef.Name)
			if err != nil {
				return nil, nil, &operationError{reason: errorNonexistentServicePlanReason, message: fmt.Sprintf("The instance references a non-existent ServicePlan %q - %v", instance.Spec.ServicePlanRef.Name, instance.Spec.PlanReference)}
			}
			rh.inProgressProperties = &v1beta1.ServiceInstancePropertiesState{ServicePlanExternalName: servicePlan.Spec.ExternalName, ServicePlanExternalID: servicePlan.Spec.ExternalID}
		}
	} else {
		if instance.Status.ExternalProperties == nil {
			return nil, nil, stderrors.New("ExternalProperties must be set before deprovisioning")
		}
		rh.inProgressProperties = instance.Status.ExternalProperties
	}
	var planExternalID string
	if instance.Spec.ClusterServiceClassSpecified() {
		planExternalID = rh.inProgressProperties.ClusterServicePlanExternalID
	} else if instance.Spec.ServiceClassSpecified() {
		planExternalID = rh.inProgressProperties.ServicePlanExternalID
	}
	request := &osb.DeprovisionRequest{InstanceID: instance.Spec.ExternalID, ServiceID: scExternalID, PlanID: planExternalID, OriginatingIdentity: rh.originatingIdentity, AcceptsIncomplete: true}
	return request, rh.inProgressProperties, nil
}
func (c *controller) prepareServiceInstanceLastOperationRequest(instance *v1beta1.ServiceInstance) (*osb.LastOperationRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if instance.Status.InProgressProperties == nil {
		pcb := pretty.NewInstanceContextBuilder(instance)
		err := stderrors.New("Instance.Status.InProgressProperties can not be nil")
		klog.Error(pcb.Message(err.Error()))
		return nil, err
	}
	var rh *requestHelper
	var scExternalID string
	var spExternalID string
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, servicePlan, _, _, err := c.getClusterServiceClassPlanAndClusterServiceBroker(instance)
		if err != nil {
			return nil, c.handleServiceInstanceReconciliationError(instance, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
		var spExternalName string
		if servicePlan != nil {
			spExternalName = servicePlan.Spec.ExternalName
			spExternalID = servicePlan.Spec.ExternalID
		} else {
			spExternalID = instance.Status.InProgressProperties.ClusterServicePlanExternalID
		}
		rh, err = c.prepareRequestHelper(instance, spExternalName, spExternalID, false)
		if err != nil {
			return nil, err
		}
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, servicePlan, _, _, err := c.getServiceClassPlanAndServiceBroker(instance)
		if err != nil {
			return nil, c.handleServiceInstanceReconciliationError(instance, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
		var spExternalName string
		if servicePlan != nil {
			spExternalName = servicePlan.Spec.ExternalName
			spExternalID = servicePlan.Spec.ExternalID
		} else {
			spExternalID = instance.Status.InProgressProperties.ServicePlanExternalID
		}
		rh, err = c.prepareRequestHelper(instance, spExternalName, spExternalID, false)
		if err != nil {
			return nil, err
		}
	}
	request := &osb.LastOperationRequest{InstanceID: instance.Spec.ExternalID, ServiceID: &scExternalID, PlanID: &spExternalID, OriginatingIdentity: rh.originatingIdentity}
	if instance.Status.LastOperation != nil && *instance.Status.LastOperation != "" {
		key := osb.OperationKey(*instance.Status.LastOperation)
		request.OperationKey = &key
	}
	return request, nil
}
func (c *controller) processServiceInstanceGracefulDeletionSuccess(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.removeFinalizer(instance)
	if _, err := c.updateServiceInstanceStatusWithRetries(instance, c.removeFinalizer); err != nil {
		return err
	}
	pcb := pretty.NewInstanceContextBuilder(instance)
	klog.Info(pcb.Message("Cleared finalizer"))
	c.removeInstanceFromRetryMap(instance)
	return nil
}
func (c *controller) removeFinalizer(instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	finalizers := sets.NewString(instance.Finalizers...)
	finalizers.Delete(v1beta1.FinalizerServiceCatalog)
	instance.Finalizers = finalizers.List()
}
func (c *controller) handleServiceInstanceReconciliationError(instance *v1beta1.ServiceInstance, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resourceErr, ok := err.(*operationError); ok {
		status := v1beta1.ConditionFalse
		if instance.Status.CurrentOperation == v1beta1.ServiceInstanceOperationDeprovision {
			status = v1beta1.ConditionUnknown
		}
		readyCond := newServiceInstanceReadyCondition(status, resourceErr.reason, resourceErr.message)
		return c.processServiceInstanceOperationError(instance, readyCond)
	}
	return err
}
func (c *controller) processServiceInstanceOperationError(instance *v1beta1.ServiceInstance, readyCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	c.recorder.Event(instance, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	return fmt.Errorf(readyCond.Message)
}
func (c *controller) processProvisionSuccess(instance *v1beta1.ServiceInstance, dashboardURL *string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceDashboardURL(instance, dashboardURL)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionTrue, successProvisionReason, successProvisionMessage)
	instance.Status.ExternalProperties = instance.Status.InProgressProperties
	clearServiceInstanceCurrentOperation(instance)
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusProvisioned
	instance.Status.ReconciledGeneration = instance.Status.ObservedGeneration
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	c.removeInstanceFromRetryMap(instance)
	c.recorder.Eventf(instance, corev1.EventTypeNormal, successProvisionReason, successProvisionMessage)
	return nil
}
func (c *controller) processTerminalProvisionFailure(instance *v1beta1.ServiceInstance, readyCond, failedCond *v1beta1.ServiceInstanceCondition, shouldMitigateOrphan bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if failedCond == nil {
		return fmt.Errorf("failedCond must not be nil")
	}
	c.removeInstanceFromRetryMap(instance)
	return c.processProvisionFailure(instance, readyCond, failedCond, shouldMitigateOrphan)
}
func (c *controller) processTemporaryProvisionFailure(instance *v1beta1.ServiceInstance, readyCond *v1beta1.ServiceInstanceCondition, shouldMitigateOrphan bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.processProvisionFailure(instance, readyCond, nil, shouldMitigateOrphan)
}
func (c *controller) processProvisionFailure(instance *v1beta1.ServiceInstance, readyCond, failedCond *v1beta1.ServiceInstanceCondition, shouldMitigateOrphan bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.recorder.Event(instance, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
	var errorMessage error
	if failedCond != nil {
		c.recorder.Event(instance, corev1.EventTypeWarning, failedCond.Reason, failedCond.Message)
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionFailed, failedCond.Status, failedCond.Reason, failedCond.Message)
		errorMessage = fmt.Errorf(failedCond.Message)
	} else {
		errorMessage = fmt.Errorf(readyCond.Message)
	}
	if shouldMitigateOrphan {
		c.recorder.Event(instance, corev1.EventTypeWarning, startingInstanceOrphanMitigationReason, startingInstanceOrphanMitigationMessage)
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionOrphanMitigation, v1beta1.ConditionTrue, readyCond.Reason, readyCond.Message)
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, startingInstanceOrphanMitigationReason, startingInstanceOrphanMitigationMessage)
		instance.Status.OrphanMitigationInProgress = true
	} else {
		instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusNotRequired
	}
	if failedCond == nil || shouldMitigateOrphan {
		clearServiceInstanceAsyncOsbOperation(instance)
	} else {
		clearServiceInstanceCurrentOperation(instance)
	}
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	if failedCond == nil || shouldMitigateOrphan {
		return errorMessage
	}
	return nil
}
func (c *controller) processProvisionAsyncResponse(instance *v1beta1.ServiceInstance, response *osb.ProvisionResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceDashboardURL(instance, response.DashboardURL)
	setServiceInstanceLastOperation(instance, response.OperationKey)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, asyncProvisioningReason, asyncProvisioningMessage)
	instance.Status.AsyncOpInProgress = true
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	c.recorder.Event(instance, corev1.EventTypeNormal, asyncProvisioningReason, asyncProvisioningMessage)
	return c.beginPollingServiceInstance(instance)
}
func (c *controller) processUpdateServiceInstanceSuccess(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionTrue, successUpdateInstanceReason, successUpdateInstanceMessage)
	instance.Status.ExternalProperties = instance.Status.InProgressProperties
	clearServiceInstanceCurrentOperation(instance)
	instance.Status.ReconciledGeneration = instance.Status.ObservedGeneration
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	c.removeInstanceFromRetryMap(instance)
	c.recorder.Eventf(instance, corev1.EventTypeNormal, successUpdateInstanceReason, successUpdateInstanceMessage)
	return nil
}
func (c *controller) processTerminalUpdateServiceInstanceFailure(instance *v1beta1.ServiceInstance, readyCond, failedCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if failedCond == nil {
		return fmt.Errorf("failedCond must not be nil")
	}
	c.removeInstanceFromRetryMap(instance)
	return c.processUpdateServiceInstanceFailure(instance, readyCond, failedCond)
}
func (c *controller) processTemporaryUpdateServiceInstanceFailure(instance *v1beta1.ServiceInstance, readyCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.processUpdateServiceInstanceFailure(instance, readyCond, nil)
}
func (c *controller) processUpdateServiceInstanceFailure(instance *v1beta1.ServiceInstance, readyCond, failedCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.recorder.Event(instance, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
	if failedCond != nil {
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionFailed, failedCond.Status, failedCond.Reason, failedCond.Message)
		clearServiceInstanceCurrentOperation(instance)
	} else {
		clearServiceInstanceAsyncOsbOperation(instance)
	}
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	if failedCond == nil {
		return fmt.Errorf(readyCond.Message)
	}
	return nil
}
func (c *controller) processUpdateServiceInstanceAsyncResponse(instance *v1beta1.ServiceInstance, response *osb.UpdateInstanceResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceLastOperation(instance, response.OperationKey)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, asyncUpdatingInstanceReason, asyncUpdatingInstanceMessage)
	instance.Status.AsyncOpInProgress = true
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	c.recorder.Event(instance, corev1.EventTypeNormal, asyncUpdatingInstanceReason, asyncUpdatingInstanceMessage)
	return c.beginPollingServiceInstance(instance)
}
func (c *controller) processDeprovisionSuccess(instance *v1beta1.ServiceInstance) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mitigatingOrphan := instance.Status.OrphanMitigationInProgress
	reason := successDeprovisionReason
	msg := successDeprovisionMessage
	if mitigatingOrphan {
		removeServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionOrphanMitigation)
		instance.Status.OrphanMitigationInProgress = false
		reason = successOrphanMitigationReason
		msg = successOrphanMitigationMessage
	}
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, reason, msg)
	clearServiceInstanceCurrentOperation(instance)
	instance.Status.ExternalProperties = nil
	instance.Status.ProvisionStatus = v1beta1.ServiceInstanceProvisionStatusNotProvisioned
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusSucceeded
	if mitigatingOrphan {
		if _, err := c.updateServiceInstanceStatus(instance); err != nil {
			return err
		}
	} else {
		if err := c.processServiceInstanceGracefulDeletionSuccess(instance); err != nil {
			return err
		}
	}
	c.recorder.Event(instance, corev1.EventTypeNormal, reason, msg)
	return nil
}
func (c *controller) processDeprovisionFailure(instance *v1beta1.ServiceInstance, readyCond, failedCond *v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if failedCond == nil {
		return fmt.Errorf("failedCond must not be nil")
	}
	if instance.Status.OrphanMitigationInProgress {
		msg := "Orphan mitigation failed: " + failedCond.Message
		readyCond := newServiceInstanceReadyCondition(v1beta1.ConditionUnknown, errorOrphanMitigationFailedReason, msg)
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
		c.recorder.Event(instance, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	} else {
		if readyCond != nil {
			setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionUnknown, readyCond.Reason, readyCond.Message)
			c.recorder.Event(instance, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
		}
		setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionFailed, failedCond.Status, failedCond.Reason, failedCond.Message)
		c.recorder.Event(instance, corev1.EventTypeWarning, failedCond.Reason, failedCond.Message)
	}
	clearServiceInstanceCurrentOperation(instance)
	instance.Status.DeprovisionStatus = v1beta1.ServiceInstanceDeprovisionStatusFailed
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	return nil
}
func (c *controller) processDeprovisionAsyncResponse(instance *v1beta1.ServiceInstance, response *osb.DeprovisionResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceInstanceLastOperation(instance, response.OperationKey)
	setServiceInstanceCondition(instance, v1beta1.ServiceInstanceConditionReady, v1beta1.ConditionFalse, asyncDeprovisioningReason, asyncDeprovisioningMessage)
	instance.Status.AsyncOpInProgress = true
	if _, err := c.updateServiceInstanceStatus(instance); err != nil {
		return err
	}
	c.recorder.Event(instance, corev1.EventTypeNormal, asyncDeprovisioningReason, asyncDeprovisioningMessage)
	return c.beginPollingServiceInstance(instance)
}
func (c *controller) handleServiceInstancePollingError(instance *v1beta1.ServiceInstance, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	klog.V(4).Info(pcb.Messagef("Error during polling: %v", err))
	return c.continuePollingServiceInstance(instance)
}
func setServiceInstanceDashboardURL(instance *v1beta1.ServiceInstance, dashboardURL *string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if dashboardURL != nil && *dashboardURL != "" {
		url := *dashboardURL
		instance.Status.DashboardURL = &url
	}
}
func setServiceInstanceLastOperation(instance *v1beta1.ServiceInstance, operationKey *osb.OperationKey) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if operationKey != nil && *operationKey != "" {
		key := string(*operationKey)
		instance.Status.LastOperation = &key
	}
}
