package controller

import (
	"bytes"
	"fmt"
	"net"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/klog"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	"github.com/kubernetes-incubator/service-catalog/pkg/pretty"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/jsonpath"
)

const (
	errorNonexistentServiceInstanceReason		string	= "ReferencesNonexistentInstance"
	errorBindCallReason				string	= "BindCallFailed"
	errorInjectingBindResultReason			string	= "ErrorInjectingBindResult"
	errorEjectingBindReason				string	= "ErrorEjectingServiceBinding"
	errorUnbindCallReason				string	= "UnbindCallFailed"
	errorNonbindableClusterServiceClassReason	string	= "ErrorNonbindableServiceClass"
	errorServiceInstanceRefsUnresolved		string	= "ErrorInstanceRefsUnresolved"
	errorServiceInstanceNotReadyReason		string	= "ErrorInstanceNotReady"
	errorServiceBindingOrphanMitigation		string	= "ServiceBindingNeedsOrphanMitigation"
	errorFetchingBindingFailedReason		string	= "FetchingBindingFailed"
	errorAsyncOpTimeoutReason			string	= "AsyncOperationTimeout"
	successInjectedBindResultReason			string	= "InjectedBindResult"
	successInjectedBindResultMessage		string	= "Injected bind result"
	successUnboundReason				string	= "UnboundSuccessfully"
	asyncBindingReason				string	= "Binding"
	asyncBindingMessage				string	= "The binding is being created asynchronously"
	asyncUnbindingReason				string	= "Unbinding"
	asyncUnbindingMessage				string	= "The binding is being deleted asynchronously"
	bindingInFlightReason				string	= "BindingRequestInFlight"
	bindingInFlightMessage				string	= "Binding request for ServiceBinding in-flight to Broker"
	unbindingInFlightReason				string	= "UnbindingRequestInFlight"
	unbindingInFlightMessage			string	= "Unbind request for ServiceBinding in-flight to Broker"
)

var bindingControllerKind = v1beta1.SchemeGroupVersion.WithKind("ServiceBinding")

func (c *controller) bindingAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		pcb := pretty.NewContextBuilder(pretty.ServiceBinding, "", "", "")
		klog.Errorf(pcb.Messagef("Couldn't get key for object %+v: %v", obj, err))
		return
	}
	pcb := pretty.NewContextBuilder(pretty.ServiceBinding, "", key, "")
	acc, err := meta.Accessor(obj)
	if err != nil {
		klog.Errorf(pcb.Messagef("error creating meta accessor: %v", err))
		return
	}
	klog.V(6).Info(pcb.Messagef("received ADD/UPDATE event for: resourceVersion: %v", acc.GetResourceVersion()))
	c.bindingQueue.Add(key)
}
func (c *controller) bindingUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	binding := newObj.(*v1beta1.ServiceBinding)
	if !binding.Status.AsyncOpInProgress {
		c.bindingAdd(newObj)
	}
}
func (c *controller) bindingDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	binding, ok := obj.(*v1beta1.ServiceBinding)
	if binding == nil || !ok {
		return
	}
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.V(4).Info(pcb.Messagef("Received DELETE event; no further processing will occur; resourceVersion %v", binding.ResourceVersion))
}
func (c *controller) reconcileServiceBindingKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	pcb := pretty.NewContextBuilder(pretty.ServiceBinding, namespace, name, "")
	binding, err := c.bindingLister.ServiceBindings(namespace).Get(name)
	if apierrors.IsNotFound(err) {
		klog.Info(pcb.Message("Not doing work because the ServiceBinding has been deleted"))
		return nil
	}
	if err != nil {
		klog.Info(pcb.Messagef("Unable to retrieve store: %v", err))
		return err
	}
	return c.reconcileServiceBinding(binding)
}
func isServiceBindingFailed(binding *v1beta1.ServiceBinding) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, condition := range binding.Status.Conditions {
		if condition.Type == v1beta1.ServiceBindingConditionFailed && condition.Status == v1beta1.ConditionTrue {
			return true
		}
	}
	return false
}
func getReconciliationActionForServiceBinding(binding *v1beta1.ServiceBinding) ReconciliationAction {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch {
	case binding.Status.AsyncOpInProgress:
		return reconcilePoll
	case binding.ObjectMeta.DeletionTimestamp != nil || binding.Status.OrphanMitigationInProgress:
		return reconcileDelete
	default:
		return reconcileAdd
	}
}
func (c *controller) reconcileServiceBinding(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.V(6).Info(pcb.Messagef(`beginning to process resourceVersion: %v`, binding.ResourceVersion))
	reconciliationAction := getReconciliationActionForServiceBinding(binding)
	switch reconciliationAction {
	case reconcileAdd:
		return c.reconcileServiceBindingAdd(binding)
	case reconcileDelete:
		return c.reconcileServiceBindingDelete(binding)
	case reconcilePoll:
		return c.pollServiceBinding(binding)
	default:
		return fmt.Errorf(pcb.Messagef("Unknown reconciliation action %v", reconciliationAction))
	}
}
func (c *controller) reconcileServiceBindingAdd(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(binding)
	if isServiceBindingFailed(binding) {
		klog.V(4).Info(pcb.Message("not processing event; status showed that it has failed"))
		return nil
	}
	if binding.Status.ReconciledGeneration == binding.Generation {
		klog.V(4).Info(pcb.Message("Not processing event; reconciled generation showed there is no work to do"))
		return nil
	}
	klog.V(4).Info(pcb.Message("Processing"))
	binding = binding.DeepCopy()
	instance, err := c.instanceLister.ServiceInstances(binding.Namespace).Get(binding.Spec.InstanceRef.Name)
	if err != nil {
		msg := fmt.Sprintf(`References a non-existent %s "%s/%s"`, pretty.ServiceInstance, binding.Namespace, binding.Spec.InstanceRef.Name)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorNonexistentServiceInstanceReason, msg)
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	var prettyName string
	var brokerClient osb.Client
	var request *osb.BindRequest
	var inProgressProperties *v1beta1.ServiceBindingPropertiesState
	if instance.Spec.ClusterServiceClassSpecified() {
		if instance.Spec.ClusterServiceClassRef == nil || instance.Spec.ClusterServicePlanRef == nil {
			msg := fmt.Sprintf(`Binding cannot begin because ClusterServiceClass and ClusterServicePlan references for %s have not been resolved yet`, pretty.ServiceInstanceName(instance))
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorServiceInstanceRefsUnresolved, msg)
			return c.processServiceBindingOperationError(binding, readyCond)
		}
		serviceClass, servicePlan, brokerName, bClient, err := c.getClusterServiceClassPlanAndClusterServiceBrokerForServiceBinding(instance, binding)
		if err != nil {
			return c.handleServiceBindingReconciliationError(binding, err)
		}
		brokerClient = bClient
		if !isClusterServicePlanBindable(serviceClass, servicePlan) {
			msg := fmt.Sprintf(`References a non-bindable %s and Plan (%q) combination`, pretty.ClusterServiceClassName(serviceClass), instance.Spec.ClusterServicePlanExternalName)
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorNonbindableClusterServiceClassReason, msg)
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, errorNonbindableClusterServiceClassReason, msg)
			return c.processBindFailure(binding, readyCond, failedCond, false)
		}
		if !isServiceInstanceReady(instance) {
			msg := fmt.Sprintf(`Binding cannot begin because referenced %s is not ready`, pretty.ServiceInstanceName(instance))
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorServiceInstanceNotReadyReason, msg)
			return c.processServiceBindingOperationError(binding, readyCond)
		}
		klog.V(4).Info(pcb.Message("Adding/Updating"))
		request, inProgressProperties, err = c.prepareBindRequest(binding, instance)
		if err != nil {
			return c.handleServiceBindingReconciliationError(binding, err)
		}
		prettyName = pretty.FromServiceInstanceOfClusterServiceClassAtBrokerName(instance, serviceClass, brokerName)
	} else if instance.Spec.ServiceClassSpecified() {
		if instance.Spec.ServiceClassRef == nil || instance.Spec.ServicePlanRef == nil {
			msg := fmt.Sprintf(`Binding cannot begin because ServiceClass and ServicePlan references for %s have not been resolved yet`, pretty.ServiceInstanceName(instance))
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorServiceInstanceRefsUnresolved, msg)
			return c.processServiceBindingOperationError(binding, readyCond)
		}
		serviceClass, servicePlan, brokerName, bClient, err := c.getServiceClassPlanAndServiceBrokerForServiceBinding(instance, binding)
		if err != nil {
			return c.handleServiceBindingReconciliationError(binding, err)
		}
		brokerClient = bClient
		if !isServicePlanBindable(serviceClass, servicePlan) {
			msg := fmt.Sprintf(`References a non-bindable %s and Plan (%q) combination`, pretty.ServiceClassName(serviceClass), instance.Spec.ClusterServicePlanExternalName)
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorNonbindableClusterServiceClassReason, msg)
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, errorNonbindableClusterServiceClassReason, msg)
			return c.processBindFailure(binding, readyCond, failedCond, false)
		}
		if !isServiceInstanceReady(instance) {
			msg := fmt.Sprintf(`Binding cannot begin because referenced %s is not ready`, pretty.ServiceInstanceName(instance))
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorServiceInstanceNotReadyReason, msg)
			return c.processServiceBindingOperationError(binding, readyCond)
		}
		klog.V(4).Info(pcb.Message("Adding/Updating"))
		request, inProgressProperties, err = c.prepareBindRequest(binding, instance)
		if err != nil {
			return c.handleServiceBindingReconciliationError(binding, err)
		}
		prettyName = pretty.FromServiceInstanceOfServiceClassAtBrokerName(instance, serviceClass, brokerName)
	}
	if binding.Status.CurrentOperation == "" {
		binding, err = c.recordStartOfServiceBindingOperation(binding, v1beta1.ServiceBindingOperationBind, inProgressProperties)
		if err != nil {
			return err
		}
		return nil
	}
	response, err := brokerClient.Bind(request)
	if err != nil {
		if httpErr, ok := osb.IsHTTPError(err); ok {
			msg := fmt.Sprintf("ServiceBroker returned failure; bind operation will not be retried: %v", err.Error())
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorBindCallReason, msg)
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, "ServiceBindingReturnedFailure", msg)
			return c.processBindFailure(binding, readyCond, failedCond, shouldStartOrphanMitigation(httpErr.StatusCode))
		}
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			msg := "Communication with the ServiceBroker timed out; Bind operation will not be retried: " + err.Error()
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, errorBindCallReason, msg)
			return c.processBindFailure(binding, nil, failedCond, true)
		}
		msg := fmt.Sprintf(`Error creating ServiceBinding for %s: %s`, prettyName, err)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorBindCallReason, msg)
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			msg := "Stopping reconciliation retries, too much time has elapsed"
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
			return c.processBindFailure(binding, readyCond, failedCond, false)
		}
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	if response.Async {
		return c.processBindAsyncResponse(binding, response)
	}
	binding.Status.ExternalProperties = binding.Status.InProgressProperties
	err = c.injectServiceBinding(binding, response.Credentials)
	if err != nil {
		msg := fmt.Sprintf(`Error injecting bind result: %s`, err)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorInjectingBindResultReason, msg)
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			msg := "Stopping reconciliation retries, too much time has elapsed"
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
			return c.processBindFailure(binding, readyCond, failedCond, true)
		}
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	return c.processBindSuccess(binding)
}
func (c *controller) reconcileServiceBindingDelete(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	pcb := pretty.NewBindingContextBuilder(binding)
	if binding.DeletionTimestamp == nil && !binding.Status.OrphanMitigationInProgress {
		return nil
	}
	if finalizers := sets.NewString(binding.Finalizers...); !finalizers.Has(v1beta1.FinalizerServiceCatalog) {
		return nil
	}
	if binding.Status.UnbindStatus == v1beta1.ServiceBindingUnbindStatusFailed {
		klog.V(4).Info(pcb.Message("Not processing delete event because unbinding has failed"))
		return nil
	}
	klog.V(4).Info(pcb.Message("Processing Delete"))
	binding = binding.DeepCopy()
	if binding.Status.UnbindStatus == v1beta1.ServiceBindingUnbindStatusNotRequired || binding.Status.UnbindStatus == v1beta1.ServiceBindingUnbindStatusSucceeded {
		return c.processServiceBindingGracefulDeletionSuccess(binding)
	}
	if err := c.ejectServiceBinding(binding); err != nil {
		msg := fmt.Sprintf(`Error ejecting binding. Error deleting secret: %s`, err)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorEjectingBindReason, msg)
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	if binding.DeletionTimestamp == nil {
		if binding.Status.OperationStartTime == nil {
			now := metav1.Now()
			binding.Status.OperationStartTime = &now
		}
	} else {
		if binding.Status.CurrentOperation != v1beta1.ServiceBindingOperationUnbind {
			binding, err = c.recordStartOfServiceBindingOperation(binding, v1beta1.ServiceBindingOperationUnbind, nil)
			if err != nil {
				return err
			}
			return nil
		}
	}
	instance, err := c.instanceLister.ServiceInstances(binding.Namespace).Get(binding.Spec.InstanceRef.Name)
	if err != nil {
		msg := fmt.Sprintf(`References a non-existent %s "%s/%s"`, pretty.ServiceInstance, binding.Namespace, binding.Spec.InstanceRef.Name)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorNonexistentServiceInstanceReason, msg)
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	if instance.Status.AsyncOpInProgress {
		msg := fmt.Sprintf(`trying to unbind to %s "%s/%s" that has ongoing asynchronous operation`, pretty.ServiceInstance, binding.Namespace, binding.Spec.InstanceRef.Name)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorWithOngoingAsyncOperationReason, msg)
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	var brokerClient osb.Client
	var prettyBrokerName string
	if instance.Spec.ClusterServiceClassSpecified() {
		if instance.Spec.ClusterServiceClassRef == nil {
			return fmt.Errorf("ClusterServiceClass reference for Instance has not been resolved yet")
		}
		if instance.Status.ExternalProperties == nil || instance.Status.ExternalProperties.ClusterServicePlanExternalID == "" {
			return fmt.Errorf("ClusterServicePlanExternalID for Instance has not been set yet")
		}
		serviceClass, brokerName, bClient, err := c.getClusterServiceClassAndClusterServiceBrokerForServiceBinding(instance, binding)
		if err != nil {
			return c.handleServiceBindingReconciliationError(binding, err)
		}
		brokerClient = bClient
		prettyBrokerName = pretty.FromServiceInstanceOfClusterServiceClassAtBrokerName(instance, serviceClass, brokerName)
	} else if instance.Spec.ServiceClassSpecified() {
		if instance.Spec.ServiceClassRef == nil {
			return fmt.Errorf("ServiceClass reference for Instance has not been resolved yet")
		}
		if instance.Status.ExternalProperties == nil || instance.Status.ExternalProperties.ServicePlanExternalID == "" {
			return fmt.Errorf("ServicePlanExternalID for Instance has not been set yet")
		}
		serviceClass, brokerName, bClient, err := c.getServiceClassAndServiceBrokerForServiceBinding(instance, binding)
		if err != nil {
			return c.handleServiceBindingReconciliationError(binding, err)
		}
		brokerClient = bClient
		prettyBrokerName = pretty.FromServiceInstanceOfServiceClassAtBrokerName(instance, serviceClass, brokerName)
	}
	request, err := c.prepareUnbindRequest(binding, instance)
	if err != nil {
		return c.handleServiceBindingReconciliationError(binding, err)
	}
	response, err := brokerClient.Unbind(request)
	if err != nil {
		msg := fmt.Sprintf(`Error unbinding from %s: %s`, prettyBrokerName, err)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionUnknown, errorUnbindCallReason, msg)
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			msg := "Stopping reconciliation retries, too much time has elapsed"
			failedCond := newServiceBindingReadyCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
			return c.processUnbindFailure(binding, readyCond, failedCond)
		}
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	if response.Async {
		return c.processUnbindAsyncResponse(binding, response)
	}
	return c.processUnbindSuccess(binding)
}
func isClusterServicePlanBindable(serviceClass *v1beta1.ClusterServiceClass, plan *v1beta1.ClusterServicePlan) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if plan.Spec.Bindable != nil {
		return *plan.Spec.Bindable
	}
	return serviceClass.Spec.Bindable
}
func isServicePlanBindable(serviceClass *v1beta1.ServiceClass, plan *v1beta1.ServicePlan) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if plan.Spec.Bindable != nil {
		return *plan.Spec.Bindable
	}
	return serviceClass.Spec.Bindable
}
func (c *controller) injectServiceBinding(binding *v1beta1.ServiceBinding, credentials map[string]interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.V(5).Info(pcb.Messagef(`Creating/updating Secret "%s/%s" with %d keys`, binding.Namespace, binding.Spec.SecretName, len(credentials)))
	if err := c.transformCredentials(binding.Spec.SecretTransforms, credentials); err != nil {
		return fmt.Errorf(`Unexpected error while transforming credentials for ServiceBinding "%s/%s": %v`, binding.Namespace, binding.Name, err)
	}
	secretData := make(map[string][]byte)
	for k, v := range credentials {
		var err error
		if secretData[k], err = serialize(v); err != nil {
			return fmt.Errorf("Unable to serialize value for credential key %q (value is intentionally not logged): %s", k, err)
		}
	}
	secretClient := c.kubeClient.CoreV1().Secrets(binding.Namespace)
	existingSecret, err := secretClient.Get(binding.Spec.SecretName, metav1.GetOptions{})
	if err == nil {
		if !metav1.IsControlledBy(existingSecret, binding) {
			controllerRef := metav1.GetControllerOf(existingSecret)
			return fmt.Errorf(`Secret "%s/%s" is not owned by ServiceBinding, controllerRef: %v`, binding.Namespace, existingSecret.Name, controllerRef)
		}
		existingSecret.Data = secretData
		if _, err = secretClient.Update(existingSecret); err != nil {
			if apierrors.IsConflict(err) {
				return fmt.Errorf(`Conflicting Secret "%s/%s" update detected`, binding.Namespace, existingSecret.Name)
			}
			return fmt.Errorf(`Unexpected error updating Secret "%s/%s": %v`, binding.Namespace, existingSecret.Name, err)
		}
	} else {
		if !apierrors.IsNotFound(err) {
			return fmt.Errorf(`Unexpected error getting Secret "%s/%s": %v`, binding.Namespace, existingSecret.Name, err)
		}
		err = nil
		secret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: binding.Spec.SecretName, Namespace: binding.Namespace, OwnerReferences: []metav1.OwnerReference{*metav1.NewControllerRef(binding, bindingControllerKind)}}, Data: secretData}
		if _, err = secretClient.Create(secret); err != nil {
			if apierrors.IsAlreadyExists(err) {
				return fmt.Errorf(`Conflicting Secret "%s/%s" creation detected`, binding.Namespace, secret.Name)
			}
			return fmt.Errorf(`Unexpected error creating Secret "%s/%s": %v`, binding.Namespace, secret.Name, err)
		}
	}
	return err
}
func (c *controller) transformCredentials(transforms []v1beta1.SecretTransform, credentials map[string]interface{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, t := range transforms {
		switch {
		case t.AddKey != nil:
			var value interface{}
			if t.AddKey.JSONPathExpression != nil {
				result, err := evaluateJSONPath(*t.AddKey.JSONPathExpression, credentials)
				if err != nil {
					return err
				}
				value = result
			} else if t.AddKey.StringValue != nil {
				value = *t.AddKey.StringValue
			} else {
				value = t.AddKey.Value
			}
			credentials[t.AddKey.Key] = value
		case t.RenameKey != nil:
			value, ok := credentials[t.RenameKey.From]
			if ok {
				credentials[t.RenameKey.To] = value
				delete(credentials, t.RenameKey.From)
			}
		case t.AddKeysFrom != nil:
			secret, err := c.kubeClient.CoreV1().Secrets(t.AddKeysFrom.SecretRef.Namespace).Get(t.AddKeysFrom.SecretRef.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
			for k, v := range secret.Data {
				credentials[k] = v
			}
		case t.RemoveKey != nil:
			delete(credentials, t.RemoveKey.Key)
		}
	}
	return nil
}
func evaluateJSONPath(jsonPath string, credentials map[string]interface{}) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	j := jsonpath.New("expression")
	buf := new(bytes.Buffer)
	if err := j.Parse(jsonPath); err != nil {
		return "", err
	}
	if err := j.Execute(buf, credentials); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func (c *controller) ejectServiceBinding(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.V(5).Info(pcb.Messagef(`Deleting Secret "%s/%s"`, binding.Namespace, binding.Spec.SecretName))
	if err = c.kubeClient.CoreV1().Secrets(binding.Namespace).Delete(binding.Spec.SecretName, &metav1.DeleteOptions{}); err != nil && !apierrors.IsNotFound(err) {
		return err
	}
	return nil
}
func setServiceBindingCondition(toUpdate *v1beta1.ServiceBinding, conditionType v1beta1.ServiceBindingConditionType, status v1beta1.ConditionStatus, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceBindingConditionInternal(toUpdate, conditionType, status, reason, message, metav1.Now())
}
func setServiceBindingConditionInternal(toUpdate *v1beta1.ServiceBinding, conditionType v1beta1.ServiceBindingConditionType, status v1beta1.ConditionStatus, reason, message string, t metav1.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(toUpdate)
	klog.Info(pcb.Message(message))
	klog.V(5).Info(pcb.Messagef("Setting condition %q to %v", conditionType, status))
	newCondition := v1beta1.ServiceBindingCondition{Type: conditionType, Status: status, Reason: reason, Message: message}
	if len(toUpdate.Status.Conditions) == 0 {
		klog.Info(pcb.Messagef("Setting lastTransitionTime for condition %q to %v", conditionType, t))
		newCondition.LastTransitionTime = t
		toUpdate.Status.Conditions = []v1beta1.ServiceBindingCondition{newCondition}
		return
	}
	for i, cond := range toUpdate.Status.Conditions {
		if cond.Type == conditionType {
			if cond.Status != newCondition.Status {
				klog.V(3).Info(pcb.Messagef("Found status change for condition %q: %q -> %q; setting lastTransitionTime to %v", conditionType, cond.Status, status, t))
				newCondition.LastTransitionTime = t
			} else {
				newCondition.LastTransitionTime = cond.LastTransitionTime
			}
			toUpdate.Status.Conditions[i] = newCondition
			return
		}
	}
	klog.V(3).Info(pcb.Messagef("Setting lastTransitionTime for condition %q to %v", conditionType, t))
	newCondition.LastTransitionTime = t
	toUpdate.Status.Conditions = append(toUpdate.Status.Conditions, newCondition)
}
func (c *controller) updateServiceBindingStatus(toUpdate *v1beta1.ServiceBinding) (*v1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(toUpdate)
	klog.V(4).Info(pcb.Message("Updating status"))
	updatedBinding, err := c.serviceCatalogClient.ServiceBindings(toUpdate.Namespace).UpdateStatus(toUpdate)
	if err != nil {
		klog.Errorf(pcb.Messagef("Error updating status: %v", err))
	} else {
		klog.V(6).Info(pcb.Messagef(`Updated status of resourceVersion: %v; got resourceVersion: %v`, toUpdate.ResourceVersion, updatedBinding.ResourceVersion))
	}
	return updatedBinding, err
}
func (c *controller) updateServiceBindingCondition(binding *v1beta1.ServiceBinding, conditionType v1beta1.ServiceBindingConditionType, status v1beta1.ConditionStatus, reason, message string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(binding)
	toUpdate := binding.DeepCopy()
	setServiceBindingCondition(toUpdate, conditionType, status, reason, message)
	klog.V(4).Info(pcb.Messagef("Updating %v condition to %v (Reason: %q, Message: %q)", conditionType, status, reason, message))
	_, err := c.serviceCatalogClient.ServiceBindings(binding.Namespace).UpdateStatus(toUpdate)
	if err != nil {
		klog.Errorf(pcb.Messagef("Error updating %v condition to %v: %v", conditionType, status, err))
	}
	return err
}
func (c *controller) recordStartOfServiceBindingOperation(toUpdate *v1beta1.ServiceBinding, operation v1beta1.ServiceBindingOperation, inProgressProperties *v1beta1.ServiceBindingPropertiesState) (*v1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentReconciledGeneration := toUpdate.Status.ReconciledGeneration
	clearServiceBindingCurrentOperation(toUpdate)
	toUpdate.Status.ReconciledGeneration = currentReconciledGeneration
	toUpdate.Status.CurrentOperation = operation
	now := metav1.Now()
	toUpdate.Status.OperationStartTime = &now
	toUpdate.Status.InProgressProperties = inProgressProperties
	reason := ""
	message := ""
	switch operation {
	case v1beta1.ServiceBindingOperationBind:
		reason = bindingInFlightReason
		message = bindingInFlightMessage
		toUpdate.Status.UnbindStatus = v1beta1.ServiceBindingUnbindStatusRequired
	case v1beta1.ServiceBindingOperationUnbind:
		reason = unbindingInFlightReason
		message = unbindingInFlightMessage
	}
	setServiceBindingCondition(toUpdate, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, reason, message)
	return c.updateServiceBindingStatus(toUpdate)
}
func clearServiceBindingCurrentOperation(toUpdate *v1beta1.ServiceBinding) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	toUpdate.Status.CurrentOperation = ""
	toUpdate.Status.OperationStartTime = nil
	toUpdate.Status.AsyncOpInProgress = false
	toUpdate.Status.LastOperation = nil
	toUpdate.Status.ReconciledGeneration = toUpdate.Generation
	toUpdate.Status.InProgressProperties = nil
	toUpdate.Status.OrphanMitigationInProgress = false
}
func rollbackBindingReconciledGenerationOnDeletion(binding *v1beta1.ServiceBinding, currentReconciledGeneration int64) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if binding.DeletionTimestamp != nil {
		klog.V(4).Infof("Not updating ReconciledGeneration after async operation because there is a deletion pending.")
		binding.Status.ReconciledGeneration = currentReconciledGeneration
	}
}
func (c *controller) requeueServiceBindingForPoll(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.bindingQueue.Add(key)
	return nil
}
func (c *controller) beginPollingServiceBinding(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(binding)
	if err != nil {
		klog.Errorf("Couldn't create a key for object %+v: %v", binding, err)
		return fmt.Errorf("Couldn't create a key for object %+v: %v", binding, err)
	}
	c.bindingPollingQueue.AddRateLimited(key)
	return nil
}
func (c *controller) continuePollingServiceBinding(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.beginPollingServiceBinding(binding)
}
func (c *controller) finishPollingServiceBinding(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(binding)
	if err != nil {
		klog.Errorf("Couldn't create a key for object %+v: %v", binding, err)
		return fmt.Errorf("Couldn't create a key for object %+v: %v", binding, err)
	}
	c.bindingPollingQueue.Forget(key)
	return nil
}
func (c *controller) pollServiceBinding(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.V(4).Infof(pcb.Message("Processing"))
	binding = binding.DeepCopy()
	instance, err := c.instanceLister.ServiceInstances(binding.Namespace).Get(binding.Spec.InstanceRef.Name)
	if err != nil {
		msg := fmt.Sprintf(`References a non-existent %s "%s/%s"`, pretty.ServiceInstance, binding.Namespace, binding.Spec.InstanceRef.Name)
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorNonexistentServiceInstanceReason, msg)
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	brokerClient, err := c.getBrokerClientForServiceBinding(instance, binding)
	if err != nil {
		return c.handleServiceBindingReconciliationError(binding, err)
	}
	mitigatingOrphan := binding.Status.OrphanMitigationInProgress
	deleting := binding.Status.CurrentOperation == v1beta1.ServiceBindingOperationUnbind || mitigatingOrphan
	request, err := c.prepareServiceBindingLastOperationRequest(binding, instance)
	if err != nil {
		return c.handleServiceBindingReconciliationError(binding, err)
	}
	klog.V(5).Info(pcb.Message("Polling last operation"))
	response, err := brokerClient.PollBindingLastOperation(request)
	if err != nil {
		if osb.IsGoneError(err) && deleting {
			if err := c.processUnbindSuccess(binding); err != nil {
				return c.handleServiceBindingPollingError(binding, err)
			}
			return c.finishPollingServiceBinding(binding)
		}
		s := fmt.Sprintf("Error polling last operation: %v", err)
		klog.V(4).Info(pcb.Message(s))
		c.recorder.Event(binding, corev1.EventTypeWarning, errorPollingLastOperationReason, s)
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			return c.processServiceBindingPollingFailureRetryTimeout(binding, nil)
		}
		return c.continuePollingServiceBinding(binding)
	}
	description := "(no description provided)"
	if response.Description != nil {
		description = *response.Description
	}
	klog.V(4).Info(pcb.Messagef("Poll returned %q : %q", response.State, description))
	switch response.State {
	case osb.StateInProgress:
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			return c.processServiceBindingPollingFailureRetryTimeout(binding, nil)
		}
		if response.Description != nil {
			reason := asyncBindingReason
			message := asyncBindingMessage
			if deleting {
				reason = asyncUnbindingReason
				message = asyncUnbindingMessage
			}
			message = fmt.Sprintf("%s (%s)", message, *response.Description)
			setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, reason, message)
			c.recorder.Event(binding, corev1.EventTypeNormal, reason, message)
			if _, err := c.updateServiceBindingStatus(binding); err != nil {
				return err
			}
		}
		klog.V(4).Info(pcb.Message("Last operation not completed (still in progress)"))
		return c.continuePollingServiceBinding(binding)
	case osb.StateSucceeded:
		if deleting {
			if err := c.processUnbindSuccess(binding); err != nil {
				return err
			}
			return c.finishPollingServiceBinding(binding)
		}
		binding.Status.ExternalProperties = binding.Status.InProgressProperties
		getBindingRequest := &osb.GetBindingRequest{InstanceID: instance.Spec.ExternalID, BindingID: binding.Spec.ExternalID}
		getBindingResponse, err := brokerClient.GetBinding(getBindingRequest)
		if err != nil {
			reason := errorFetchingBindingFailedReason
			msg := fmt.Sprintf("Could not do a GET on binding resource: %v", err)
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, reason, msg)
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, reason, msg)
			if err := c.processBindFailure(binding, readyCond, failedCond, true); err != nil {
				return err
			}
			return c.finishPollingServiceBinding(binding)
		}
		if err := c.injectServiceBinding(binding, getBindingResponse.Credentials); err != nil {
			reason := errorInjectingBindResultReason
			msg := fmt.Sprintf("Error injecting bind results: %v", err)
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, reason, msg)
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, reason, msg)
			if err := c.processBindFailure(binding, readyCond, failedCond, true); err != nil {
				return err
			}
			return c.finishPollingServiceBinding(binding)
		}
		if err := c.processBindSuccess(binding); err != nil {
			return err
		}
		return c.finishPollingServiceBinding(binding)
	case osb.StateFailed:
		if !deleting {
			reason := errorBindCallReason
			message := "Bind call failed: " + description
			readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, reason, message)
			failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, reason, message)
			if err := c.processBindFailure(binding, readyCond, failedCond, false); err != nil {
				return c.handleServiceBindingPollingError(binding, err)
			}
			return c.finishPollingServiceBinding(binding)
		}
		msg := "Unbind call failed: " + description
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionUnknown, errorUnbindCallReason, msg)
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			return c.processServiceBindingPollingFailureRetryTimeout(binding, readyCond)
		}
		setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorUnbindCallReason, msg)
		binding.Status.AsyncOpInProgress = false
		binding.Status.LastOperation = nil
		if _, err := c.updateServiceBindingStatus(binding); err != nil {
			return err
		}
		c.finishPollingServiceBinding(binding)
		return fmt.Errorf(readyCond.Message)
	default:
		klog.Warning(pcb.Messagef("Got invalid state in LastOperationResponse: %q", response.State))
		if c.reconciliationRetryDurationExceeded(binding.Status.OperationStartTime) {
			return c.processServiceBindingPollingFailureRetryTimeout(binding, nil)
		}
		return c.continuePollingServiceBinding(binding)
	}
}
func (c *controller) processServiceBindingPollingFailureRetryTimeout(binding *v1beta1.ServiceBinding, readyCond *v1beta1.ServiceBindingCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mitigatingOrphan := binding.Status.OrphanMitigationInProgress
	deleting := binding.Status.CurrentOperation == v1beta1.ServiceBindingOperationUnbind || mitigatingOrphan
	if readyCond == nil {
		operation := "Bind"
		status := v1beta1.ConditionFalse
		if deleting {
			operation = "Unbind"
			status = v1beta1.ConditionUnknown
		}
		msg := fmt.Sprintf("The asynchronous %v operation timed out and will not be retried", operation)
		readyCond = newServiceBindingReadyCondition(status, errorAsyncOpTimeoutReason, msg)
	}
	msg := "Stopping reconciliation retries because too much time has elapsed"
	failedCond := newServiceBindingFailedCondition(v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, msg)
	var err error
	if deleting {
		err = c.processUnbindFailure(binding, readyCond, failedCond)
	} else {
		c.finishPollingServiceBinding(binding)
		return c.processBindFailure(binding, readyCond, failedCond, true)
	}
	if err != nil {
		return c.handleServiceBindingPollingError(binding, err)
	}
	return c.finishPollingServiceBinding(binding)
}
func newServiceBindingReadyCondition(status v1beta1.ConditionStatus, reason, message string) *v1beta1.ServiceBindingCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v1beta1.ServiceBindingCondition{Type: v1beta1.ServiceBindingConditionReady, Status: status, Reason: reason, Message: message, LastTransitionTime: metav1.Now()}
}
func newServiceBindingFailedCondition(status v1beta1.ConditionStatus, reason, message string) *v1beta1.ServiceBindingCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v1beta1.ServiceBindingCondition{Type: v1beta1.ServiceBindingConditionFailed, Status: status, Reason: reason, Message: message, LastTransitionTime: metav1.Now()}
}
func setServiceBindingLastOperation(binding *v1beta1.ServiceBinding, operationKey *osb.OperationKey) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if operationKey != nil && *operationKey != "" {
		key := string(*operationKey)
		binding.Status.LastOperation = &key
	}
}
func (c *controller) prepareBindRequest(binding *v1beta1.ServiceBinding, instance *v1beta1.ServiceInstance) (*osb.BindRequest, *v1beta1.ServiceBindingPropertiesState, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var scExternalID string
	var spExternalID string
	var scBindingRetrievable bool
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, err := c.getClusterServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, nil, &operationError{reason: errorNonexistentClusterServiceClassReason, message: err.Error()}
		}
		servicePlan, err := c.getClusterServicePlanForServiceBinding(instance, binding, serviceClass)
		if err != nil {
			return nil, nil, &operationError{reason: errorNonexistentClusterServicePlanReason, message: err.Error()}
		}
		scExternalID = serviceClass.Spec.ExternalID
		spExternalID = servicePlan.Spec.ExternalID
		scBindingRetrievable = serviceClass.Spec.BindingRetrievable
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, err := c.getServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, nil, &operationError{reason: errorNonexistentServiceClassReason, message: err.Error()}
		}
		servicePlan, err := c.getServicePlanForServiceBinding(instance, binding, serviceClass)
		if err != nil {
			return nil, nil, &operationError{reason: errorNonexistentServicePlanReason, message: err.Error()}
		}
		scExternalID = serviceClass.Spec.ExternalID
		spExternalID = servicePlan.Spec.ExternalID
		scBindingRetrievable = serviceClass.Spec.BindingRetrievable
	}
	ns, err := c.kubeClient.CoreV1().Namespaces().Get(instance.Namespace, metav1.GetOptions{})
	if err != nil {
		return nil, nil, &operationError{reason: errorFindingNamespaceServiceInstanceReason, message: fmt.Sprintf(`Failed to get namespace %q during binding: %s`, instance.Namespace, err)}
	}
	parameters, parametersChecksum, rawParametersWithRedaction, err := prepareInProgressPropertyParameters(c.kubeClient, binding.Namespace, binding.Spec.Parameters, binding.Spec.ParametersFrom)
	if err != nil {
		return nil, nil, &operationError{reason: errorWithParametersReason, message: err.Error()}
	}
	inProgressProperties := &v1beta1.ServiceBindingPropertiesState{Parameters: rawParametersWithRedaction, ParameterChecksum: parametersChecksum, UserInfo: binding.Spec.UserInfo}
	appGUID := string(ns.UID)
	clusterID := c.getClusterID()
	requestContext := map[string]interface{}{"platform": ContextProfilePlatformKubernetes, "namespace": instance.Namespace, clusterIdentifierKey: clusterID}
	request := &osb.BindRequest{BindingID: binding.Spec.ExternalID, InstanceID: instance.Spec.ExternalID, ServiceID: scExternalID, PlanID: spExternalID, AppGUID: &appGUID, Parameters: parameters, BindResource: &osb.BindResource{AppGUID: &appGUID}, Context: requestContext}
	if scBindingRetrievable && utilfeature.DefaultFeatureGate.Enabled(scfeatures.AsyncBindingOperations) {
		request.AcceptsIncomplete = true
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		originatingIdentity, err := buildOriginatingIdentity(binding.Spec.UserInfo)
		if err != nil {
			return nil, nil, &operationError{reason: errorWithOriginatingIdentityReason, message: fmt.Sprintf(`Error building originating identity headers for binding: %v`, err)}
		}
		request.OriginatingIdentity = originatingIdentity
	}
	return request, inProgressProperties, nil
}
func (c *controller) prepareUnbindRequest(binding *v1beta1.ServiceBinding, instance *v1beta1.ServiceInstance) (*osb.UnbindRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var scExternalID string
	var scBindingRetrievable bool
	var planExternalID string
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, err := c.getClusterServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, c.handleServiceBindingReconciliationError(binding, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
		scBindingRetrievable = serviceClass.Spec.BindingRetrievable
		planExternalID = instance.Status.ExternalProperties.ClusterServicePlanExternalID
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, err := c.getServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, c.handleServiceBindingReconciliationError(binding, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
		scBindingRetrievable = serviceClass.Spec.BindingRetrievable
		planExternalID = instance.Status.ExternalProperties.ServicePlanExternalID
	}
	request := &osb.UnbindRequest{BindingID: binding.Spec.ExternalID, InstanceID: instance.Spec.ExternalID, ServiceID: scExternalID, PlanID: planExternalID}
	if scBindingRetrievable && utilfeature.DefaultFeatureGate.Enabled(scfeatures.AsyncBindingOperations) {
		request.AcceptsIncomplete = true
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		originatingIdentity, err := buildOriginatingIdentity(binding.Spec.UserInfo)
		if err != nil {
			return nil, &operationError{reason: errorWithOriginatingIdentityReason, message: fmt.Sprintf(`Error building originating identity headers for binding: %v`, err)}
		}
		request.OriginatingIdentity = originatingIdentity
	}
	return request, nil
}
func (c *controller) prepareServiceBindingLastOperationRequest(binding *v1beta1.ServiceBinding, instance *v1beta1.ServiceInstance) (*osb.BindingLastOperationRequest, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var scExternalID string
	var spExternalID string
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, err := c.getClusterServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, c.handleServiceBindingReconciliationError(binding, err)
		}
		servicePlan, err := c.getClusterServicePlanForServiceBinding(instance, binding, serviceClass)
		if err != nil {
			return nil, c.handleServiceBindingReconciliationError(binding, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
		spExternalID = servicePlan.Spec.ExternalID
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, err := c.getServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, c.handleServiceBindingReconciliationError(binding, err)
		}
		servicePlan, err := c.getServicePlanForServiceBinding(instance, binding, serviceClass)
		if err != nil {
			return nil, c.handleServiceBindingReconciliationError(binding, err)
		}
		scExternalID = serviceClass.Spec.ExternalID
		spExternalID = servicePlan.Spec.ExternalID
	}
	request := &osb.BindingLastOperationRequest{InstanceID: instance.Spec.ExternalID, BindingID: binding.Spec.ExternalID, ServiceID: &scExternalID, PlanID: &spExternalID}
	if binding.Status.LastOperation != nil && *binding.Status.LastOperation != "" {
		key := osb.OperationKey(*binding.Status.LastOperation)
		request.OperationKey = &key
	}
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity) {
		originatingIdentity, err := buildOriginatingIdentity(binding.Spec.UserInfo)
		if err != nil {
			return nil, &operationError{reason: errorWithOriginatingIdentityReason, message: fmt.Sprintf(`Error building originating identity headers for polling binding last operation: %v`, err)}
		}
		request.OriginatingIdentity = originatingIdentity
	}
	return request, nil
}
func (c *controller) processServiceBindingOperationError(binding *v1beta1.ServiceBinding, readyCond *v1beta1.ServiceBindingCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.recorder.Event(binding, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	setServiceBindingCondition(binding, readyCond.Type, readyCond.Status, readyCond.Reason, readyCond.Message)
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	return fmt.Errorf(readyCond.Message)
}
func (c *controller) processBindSuccess(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionTrue, successInjectedBindResultReason, successInjectedBindResultMessage)
	currentReconciledGeneration := binding.Status.ReconciledGeneration
	clearServiceBindingCurrentOperation(binding)
	rollbackBindingReconciledGenerationOnDeletion(binding, currentReconciledGeneration)
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	c.recorder.Event(binding, corev1.EventTypeNormal, successInjectedBindResultReason, successInjectedBindResultMessage)
	return nil
}
func (c *controller) processBindFailure(binding *v1beta1.ServiceBinding, readyCond, failedCond *v1beta1.ServiceBindingCondition, shouldMitigateOrphan bool) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentReconciledGeneration := binding.Status.ReconciledGeneration
	if readyCond != nil {
		c.recorder.Event(binding, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
		setServiceBindingCondition(binding, readyCond.Type, readyCond.Status, readyCond.Reason, readyCond.Message)
	}
	c.recorder.Event(binding, corev1.EventTypeWarning, failedCond.Reason, failedCond.Message)
	setServiceBindingCondition(binding, failedCond.Type, failedCond.Status, failedCond.Reason, failedCond.Message)
	if shouldMitigateOrphan {
		msg := "Starting orphan mitigation"
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, errorServiceBindingOrphanMitigation, msg)
		setServiceBindingCondition(binding, readyCond.Type, readyCond.Status, readyCond.Reason, readyCond.Message)
		c.recorder.Event(binding, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
		binding.Status.OrphanMitigationInProgress = true
		binding.Status.AsyncOpInProgress = false
		binding.Status.OperationStartTime = nil
	} else {
		clearServiceBindingCurrentOperation(binding)
		rollbackBindingReconciledGenerationOnDeletion(binding, currentReconciledGeneration)
	}
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	return nil
}
func (c *controller) processBindAsyncResponse(binding *v1beta1.ServiceBinding, response *osb.BindResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceBindingLastOperation(binding, response.OperationKey)
	setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, asyncBindingReason, asyncBindingMessage)
	binding.Status.AsyncOpInProgress = true
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	c.recorder.Event(binding, corev1.EventTypeNormal, asyncBindingReason, asyncBindingMessage)
	return c.beginPollingServiceBinding(binding)
}
func (c *controller) handleServiceBindingReconciliationError(binding *v1beta1.ServiceBinding, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if resourceErr, ok := err.(*operationError); ok {
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionFalse, resourceErr.reason, resourceErr.message)
		return c.processServiceBindingOperationError(binding, readyCond)
	}
	return err
}
func (c *controller) processServiceBindingGracefulDeletionSuccess(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	finalizers := sets.NewString(binding.Finalizers...)
	finalizers.Delete(v1beta1.FinalizerServiceCatalog)
	binding.Finalizers = finalizers.List()
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.Info(pcb.Message("Cleared finalizer"))
	return nil
}
func (c *controller) processUnbindSuccess(binding *v1beta1.ServiceBinding) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	mitigatingOrphan := binding.Status.OrphanMitigationInProgress
	reason := successUnboundReason
	msg := "The binding was deleted successfully"
	if mitigatingOrphan {
		reason = successOrphanMitigationReason
		msg = successOrphanMitigationMessage
	}
	setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, reason, msg)
	clearServiceBindingCurrentOperation(binding)
	binding.Status.ExternalProperties = nil
	binding.Status.UnbindStatus = v1beta1.ServiceBindingUnbindStatusSucceeded
	if mitigatingOrphan {
		if _, err := c.updateServiceBindingStatus(binding); err != nil {
			return err
		}
	} else {
		if err := c.processServiceBindingGracefulDeletionSuccess(binding); err != nil {
			return err
		}
	}
	c.recorder.Event(binding, corev1.EventTypeNormal, reason, msg)
	return nil
}
func (c *controller) processUnbindFailure(binding *v1beta1.ServiceBinding, readyCond, failedCond *v1beta1.ServiceBindingCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if failedCond == nil {
		return fmt.Errorf("failedCond must not be nil")
	}
	if readyCond != nil {
		setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionUnknown, readyCond.Reason, readyCond.Message)
		c.recorder.Event(binding, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	}
	if binding.Status.OrphanMitigationInProgress {
		msg := "Orphan mitigation failed: " + failedCond.Message
		readyCond := newServiceBindingReadyCondition(v1beta1.ConditionUnknown, errorOrphanMitigationFailedReason, msg)
		setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, readyCond.Status, readyCond.Reason, readyCond.Message)
		c.recorder.Event(binding, corev1.EventTypeWarning, readyCond.Reason, readyCond.Message)
	} else {
		setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionFailed, failedCond.Status, failedCond.Reason, failedCond.Message)
		c.recorder.Event(binding, corev1.EventTypeWarning, failedCond.Reason, failedCond.Message)
	}
	clearServiceBindingCurrentOperation(binding)
	binding.Status.UnbindStatus = v1beta1.ServiceBindingUnbindStatusFailed
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	return nil
}
func (c *controller) processUnbindAsyncResponse(binding *v1beta1.ServiceBinding, response *osb.UnbindResponse) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	setServiceBindingLastOperation(binding, response.OperationKey)
	setServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, asyncUnbindingReason, asyncUnbindingMessage)
	binding.Status.AsyncOpInProgress = true
	if _, err := c.updateServiceBindingStatus(binding); err != nil {
		return err
	}
	c.recorder.Event(binding, corev1.EventTypeNormal, asyncUnbindingReason, asyncUnbindingMessage)
	return c.beginPollingServiceBinding(binding)
}
func (c *controller) handleServiceBindingPollingError(binding *v1beta1.ServiceBinding, err error) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewBindingContextBuilder(binding)
	klog.V(4).Info(pcb.Messagef("Error during polling: %v", err))
	return c.continuePollingServiceBinding(binding)
}
