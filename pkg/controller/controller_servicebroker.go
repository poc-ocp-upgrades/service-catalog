package controller

import (
	"fmt"
	"time"
	"k8s.io/klog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/cache"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/metrics"
	"github.com/kubernetes-incubator/service-catalog/pkg/pretty"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
)

const (
	errorListingServiceClassesReason	string	= "ErrorListingServiceClasses"
	errorListingServiceClassesMessage	string	= "Error listing service classes."
	errorListingServicePlansReason		string	= "ErrorListingServicePlans"
	errorListingServicePlansMessage		string	= "Error listing service plans."
	errorDeletingServiceClassReason		string	= "ErrorDeletingServiceClass"
	errorDeletingServiceClassMessage	string	= "Error deleting service class."
	errorDeletingServicePlanReason		string	= "ErrorDeletingServicePlan"
	errorDeletingServicePlanMessage		string	= "Error deleting service plan."
	successServiceBrokerDeletedReason	string	= "DeletedSuccessfully"
	successServiceBrokerDeletedMessage	string	= "The servicebroker %v was deleted successfully."
)

func (c *controller) serviceBrokerAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.serviceBrokerQueue.Add(key)
}
func (c *controller) serviceBrokerUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.serviceBrokerAdd(newObj)
}
func (c *controller) serviceBrokerDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, ok := obj.(*v1beta1.ServiceBroker)
	if broker == nil || !ok {
		return
	}
	klog.V(4).Infof("Received delete event for ServiceBroker %v; no further processing will occur", broker.Name)
}
func shouldReconcileServiceBroker(broker *v1beta1.ServiceBroker, now time.Time, defaultRelistInterval time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return shouldReconcileServiceBrokerCommon(pretty.NewServiceBrokerContextBuilder(broker), &broker.ObjectMeta, &broker.Spec.CommonServiceBrokerSpec, &broker.Status.CommonServiceBrokerStatus, now, defaultRelistInterval)
}
func (c *controller) reconcileServiceBrokerKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	pcb := pretty.NewContextBuilder(pretty.ServiceBroker, namespace, name, "")
	broker, err := c.serviceBrokerLister.ServiceBrokers(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.Info(pcb.Message("Not doing work because the ServiceBroker has been deleted"))
		c.brokerClientManager.RemoveBrokerClient(NewServiceBrokerKey(namespace, name))
		return nil
	}
	if err != nil {
		klog.Info(pcb.Messagef("Unable to retrieve ServiceBroker: %v", err))
		return err
	}
	return c.reconcileServiceBroker(broker)
}
func (c *controller) updateServiceBrokerClient(broker *v1beta1.ServiceBroker) (osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewServiceBrokerContextBuilder(broker)
	authConfig, err := getAuthCredentialsFromServiceBroker(c.kubeClient, broker)
	if err != nil {
		s := fmt.Sprintf("Error getting broker auth credentials: %s", err)
		klog.Info(pcb.Message(s))
		c.recorder.Event(broker, corev1.EventTypeWarning, errorAuthCredentialsReason, s)
		if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorFetchingCatalogReason, errorFetchingCatalogMessage+s); err != nil {
			return nil, err
		}
		return nil, err
	}
	clientConfig := NewClientConfigurationForBroker(broker.ObjectMeta, &broker.Spec.CommonServiceBrokerSpec, authConfig)
	brokerClient, err := c.brokerClientManager.UpdateBrokerClient(NewServiceBrokerKey(broker.Namespace, broker.Name), clientConfig)
	if err != nil {
		s := fmt.Sprintf("Error creating client for broker %q: %s", broker.Name, err)
		klog.Info(pcb.Message(s))
		c.recorder.Event(broker, corev1.EventTypeWarning, errorAuthCredentialsReason, s)
		if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorFetchingCatalogReason, errorFetchingCatalogMessage+s); err != nil {
			return nil, err
		}
		return nil, err
	}
	return brokerClient, nil
}
func (c *controller) reconcileServiceBroker(broker *v1beta1.ServiceBroker) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewServiceBrokerContextBuilder(broker)
	klog.V(4).Infof(pcb.Message("Processing"))
	if !shouldReconcileServiceBroker(broker, time.Now(), c.brokerRelistInterval) {
		return nil
	}
	if broker.DeletionTimestamp == nil {
		klog.V(4).Info(pcb.Message("Processing adding/update event"))
		brokerClient, err := c.updateServiceBrokerClient(broker)
		if err != nil {
			return err
		}
		now := metav1.Now()
		brokerCatalog, err := brokerClient.GetCatalog()
		if err != nil {
			s := fmt.Sprintf("Error getting broker catalog: %s", err)
			klog.Warning(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorFetchingCatalogReason, s)
			if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorFetchingCatalogReason, errorFetchingCatalogMessage+s); err != nil {
				return err
			}
			if broker.Status.OperationStartTime == nil {
				toUpdate := broker.DeepCopy()
				toUpdate.Status.OperationStartTime = &now
				if _, err := c.serviceCatalogClient.ServiceBrokers(broker.Namespace).UpdateStatus(toUpdate); err != nil {
					klog.Error(pcb.Messagef("Error updating operation start time: %v", err))
					return err
				}
			} else if !time.Now().Before(broker.Status.OperationStartTime.Time.Add(c.reconciliationRetryDuration)) {
				s := "Stopping reconciliation retries because too much time has elapsed"
				klog.Info(pcb.Message(s))
				c.recorder.Event(broker, corev1.EventTypeWarning, errorReconciliationRetryTimeoutReason, s)
				toUpdate := broker.DeepCopy()
				toUpdate.Status.OperationStartTime = nil
				toUpdate.Status.ReconciledGeneration = toUpdate.Generation
				return c.updateServiceBrokerCondition(toUpdate, v1beta1.ServiceBrokerConditionFailed, v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, s)
			}
			return err
		}
		klog.V(5).Info(pcb.Messagef("Successfully fetched %v catalog entries", len(brokerCatalog.Services)))
		if broker.Status.OperationStartTime != nil {
			toUpdate := broker.DeepCopy()
			toUpdate.Status.OperationStartTime = nil
			if _, err := c.serviceCatalogClient.ServiceBrokers(broker.Namespace).UpdateStatus(toUpdate); err != nil {
				klog.Error(pcb.Messagef("Error updating operation start time: %v", err))
				return err
			}
		}
		existingServiceClasses, existingServicePlans, err := c.getCurrentServiceClassesAndPlansForNamespacedBroker(broker)
		if err != nil {
			return err
		}
		existingServiceClassMap := convertServiceClassListToMap(existingServiceClasses)
		existingServicePlanMap := convertServicePlanListToMap(existingServicePlans)
		klog.V(4).Info(pcb.Message("Converting catalog response into service-catalog API"))
		payloadServiceClasses, payloadServicePlans, err := convertAndFilterCatalogToNamespacedTypes(broker.Namespace, brokerCatalog, broker.Spec.CatalogRestrictions, existingServiceClassMap, existingServicePlanMap)
		if err != nil {
			s := fmt.Sprintf("Error converting catalog payload for broker %q to service-catalog API: %s", broker.Name, err)
			klog.Warning(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
			if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
				return err
			}
			return err
		}
		klog.V(5).Info(pcb.Message("Successfully converted catalog payload from to service-catalog API"))
		for _, payloadServiceClass := range payloadServiceClasses {
			existingServiceClass, _ := existingServiceClassMap[payloadServiceClass.Name]
			delete(existingServiceClassMap, payloadServiceClass.Name)
			if existingServiceClass == nil {
				existingServiceClass, _ = existingServiceClassMap[payloadServiceClass.Spec.ExternalID]
				delete(existingServiceClassMap, payloadServiceClass.Spec.ExternalID)
			}
			klog.V(4).Info(pcb.Messagef("Reconciling %s", pretty.ServiceClassName(payloadServiceClass)))
			if err := c.reconcileServiceClassFromServiceBrokerCatalog(broker, payloadServiceClass, existingServiceClass); err != nil {
				s := fmt.Sprintf("Error reconciling %s (broker %q): %s", pretty.ServiceClassName(payloadServiceClass), broker.Name, err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
					return err
				}
				return err
			}
			klog.V(5).Info(pcb.Messagef("Reconciled %s", pretty.ServiceClassName(payloadServiceClass)))
		}
		for _, existingServiceClass := range existingServiceClassMap {
			if existingServiceClass.Status.RemovedFromBrokerCatalog {
				continue
			}
			klog.V(4).Info(pcb.Messagef("%s has been removed from broker's catalog; marking", pretty.ServiceClassName(existingServiceClass)))
			existingServiceClass.Status.RemovedFromBrokerCatalog = true
			_, err := c.serviceCatalogClient.ServiceClasses(broker.Namespace).UpdateStatus(existingServiceClass)
			if err != nil {
				s := fmt.Sprintf("Error updating status of %s: %v", pretty.ServiceClassName(existingServiceClass), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
					return err
				}
				return err
			}
		}
		for _, payloadServicePlan := range payloadServicePlans {
			existingServicePlan, _ := existingServicePlanMap[payloadServicePlan.Name]
			delete(existingServicePlanMap, payloadServicePlan.Name)
			if existingServicePlan == nil {
				existingServicePlan, _ = existingServicePlanMap[payloadServicePlan.Spec.ExternalID]
				delete(existingServicePlanMap, payloadServicePlan.Spec.ExternalID)
			}
			klog.V(4).Infof("ServiceBroker %q: reconciling %s", broker.Name, pretty.ServicePlanName(payloadServicePlan))
			if err := c.reconcileServicePlanFromServiceBrokerCatalog(broker, payloadServicePlan, existingServicePlan); err != nil {
				s := fmt.Sprintf("Error reconciling %s: %s", pretty.ServicePlanName(payloadServicePlan), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s)
				return err
			}
			klog.V(5).Info(pcb.Messagef("Reconciled %s", pretty.ServicePlanName(payloadServicePlan)))
		}
		for _, existingServicePlan := range existingServicePlanMap {
			if existingServicePlan.Status.RemovedFromBrokerCatalog {
				continue
			}
			klog.V(4).Info(pcb.Messagef("%s has been removed from broker's catalog; marking", pretty.ServicePlanName(existingServicePlan)))
			existingServicePlan.Status.RemovedFromBrokerCatalog = true
			_, err := c.serviceCatalogClient.ServicePlans(broker.Namespace).UpdateStatus(existingServicePlan)
			if err != nil {
				s := fmt.Sprintf("Error updating status of %s: %v", pretty.ServicePlanName(existingServicePlan), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
					return err
				}
				return err
			}
		}
		if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionTrue, successFetchedCatalogReason, successFetchedCatalogMessage); err != nil {
			return err
		}
		c.recorder.Event(broker, corev1.EventTypeNormal, successFetchedCatalogReason, successFetchedCatalogMessage)
		metrics.BrokerServiceClassCount.WithLabelValues(broker.Name).Set(float64(len(payloadServiceClasses)))
		metrics.BrokerServicePlanCount.WithLabelValues(broker.Name).Set(float64(len(payloadServicePlans)))
		return nil
	}
	if finalizers := sets.NewString(broker.Finalizers...); finalizers.Has(v1beta1.FinalizerServiceCatalog) {
		klog.V(4).Info(pcb.Message("Finalizing"))
		existingServiceClasses, existingServicePlans, err := c.getCurrentServiceClassesAndPlansForNamespacedBroker(broker)
		if err != nil {
			return err
		}
		klog.V(4).Info(pcb.Messagef("Found %d ServiceClasses and %d ServicePlans to delete", len(existingServiceClasses), len(existingServicePlans)))
		for _, plan := range existingServicePlans {
			klog.V(4).Info(pcb.Messagef("Deleting %s", pretty.ServicePlanName(&plan)))
			err := c.serviceCatalogClient.ServicePlans(broker.Namespace).Delete(plan.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				s := fmt.Sprintf("Error deleting %s: %s", pretty.ServicePlanName(&plan), err)
				klog.Warning(pcb.Message(s))
				c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorDeletingServicePlanMessage, errorDeletingServicePlanReason+s)
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorDeletingServicePlanReason, "%v %v", errorDeletingServicePlanMessage, s)
				return err
			}
		}
		for _, svcClass := range existingServiceClasses {
			klog.V(4).Info(pcb.Messagef("Deleting %s", pretty.ServiceClassName(&svcClass)))
			err = c.serviceCatalogClient.ServiceClasses(broker.Namespace).Delete(svcClass.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				s := fmt.Sprintf("Error deleting %s: %s", pretty.ServiceClassName(&svcClass), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorDeletingServiceClassReason, "%v %v", errorDeletingServiceClassMessage, s)
				if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorDeletingServiceClassMessage, errorDeletingServiceClassReason+s); err != nil {
					return err
				}
				return err
			}
		}
		if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, successServiceBrokerDeletedReason, "The broker was deleted successfully"); err != nil {
			return err
		}
		finalizers.Delete(v1beta1.FinalizerServiceCatalog)
		c.updateServiceBrokerFinalizers(broker, finalizers.List())
		c.recorder.Eventf(broker, corev1.EventTypeNormal, successServiceBrokerDeletedReason, successServiceBrokerDeletedMessage, broker.Name)
		klog.V(5).Info(pcb.Message("Successfully deleted"))
		metrics.BrokerServiceClassCount.DeleteLabelValues(broker.Name)
		metrics.BrokerServicePlanCount.DeleteLabelValues(broker.Name)
		return nil
	}
	return nil
}
func (c *controller) reconcileServiceClassFromServiceBrokerCatalog(broker *v1beta1.ServiceBroker, serviceClass, existingServiceClass *v1beta1.ServiceClass) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewServiceBrokerContextBuilder(broker)
	serviceClass.Spec.ServiceBrokerName = broker.Name
	if existingServiceClass == nil {
		otherServiceClass, err := c.serviceClassLister.ServiceClasses(broker.Namespace).Get(serviceClass.Name)
		if err != nil {
			if !errors.IsNotFound(err) {
				return err
			}
		} else {
			if otherServiceClass.Spec.ServiceBrokerName != broker.Name {
				errMsg := fmt.Sprintf("%s already exists for Broker %q", pretty.ServiceClassName(serviceClass), otherServiceClass.Spec.ServiceBrokerName)
				klog.Error(pcb.Message(errMsg))
				return fmt.Errorf(errMsg)
			}
		}
		klog.V(5).Info(pcb.Messagef("Fresh %s; creating", pretty.ServiceClassName(serviceClass)))
		if _, err := c.serviceCatalogClient.ServiceClasses(broker.Namespace).Create(serviceClass); err != nil {
			klog.Error(pcb.Messagef("Error creating %s: %v", pretty.ServiceClassName(serviceClass), err))
			return err
		}
		return nil
	}
	if existingServiceClass.Spec.ExternalID != serviceClass.Spec.ExternalID {
		errMsg := fmt.Sprintf("%s already exists with OSB guid %q, received different guid %q", pretty.ServiceClassName(serviceClass), existingServiceClass.Name, serviceClass.Name)
		klog.Error(pcb.Message(errMsg))
		return fmt.Errorf(errMsg)
	}
	klog.V(5).Info(pcb.Messagef("Found existing %s; updating", pretty.ServiceClassName(serviceClass)))
	toUpdate := existingServiceClass.DeepCopy()
	toUpdate.Spec.BindingRetrievable = serviceClass.Spec.BindingRetrievable
	toUpdate.Spec.Bindable = serviceClass.Spec.Bindable
	toUpdate.Spec.PlanUpdatable = serviceClass.Spec.PlanUpdatable
	toUpdate.Spec.Tags = serviceClass.Spec.Tags
	toUpdate.Spec.Description = serviceClass.Spec.Description
	toUpdate.Spec.Requires = serviceClass.Spec.Requires
	toUpdate.Spec.ExternalName = serviceClass.Spec.ExternalName
	toUpdate.Spec.ExternalMetadata = serviceClass.Spec.ExternalMetadata
	updatedServiceClass, err := c.serviceCatalogClient.ServiceClasses(broker.Namespace).Update(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating %s: %v", pretty.ServiceClassName(serviceClass), err))
		return err
	}
	if updatedServiceClass.Status.RemovedFromBrokerCatalog {
		klog.V(4).Info(pcb.Messagef("Resetting RemovedFromBrokerCatalog status on %s", pretty.ServiceClassName(serviceClass)))
		updatedServiceClass.Status.RemovedFromBrokerCatalog = false
		_, err := c.serviceCatalogClient.ServiceClasses(broker.Namespace).UpdateStatus(updatedServiceClass)
		if err != nil {
			s := fmt.Sprintf("Error updating status of %s: %v", pretty.ServiceClassName(updatedServiceClass), err)
			klog.Warning(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
			if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
				return err
			}
			return err
		}
	}
	return nil
}
func (c *controller) reconcileServicePlanFromServiceBrokerCatalog(broker *v1beta1.ServiceBroker, servicePlan, existingServicePlan *v1beta1.ServicePlan) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewServiceBrokerContextBuilder(broker)
	servicePlan.Spec.ServiceBrokerName = broker.Name
	if existingServicePlan == nil {
		otherServicePlan, err := c.servicePlanLister.ServicePlans(broker.Namespace).Get(servicePlan.Name)
		if err != nil {
			if !errors.IsNotFound(err) {
				return err
			}
		} else {
			if otherServicePlan.Spec.ServiceBrokerName != broker.Name {
				errMsg := fmt.Sprintf("%s already exists for Broker %q", pretty.ServicePlanName(servicePlan), otherServicePlan.Spec.ServiceBrokerName)
				klog.Error(pcb.Message(errMsg))
				return fmt.Errorf(errMsg)
			}
		}
		if _, err := c.serviceCatalogClient.ServicePlans(broker.Namespace).Create(servicePlan); err != nil {
			klog.Error(pcb.Messagef("Error creating %s: %v", pretty.ServicePlanName(servicePlan), err))
			return err
		}
		return nil
	}
	if existingServicePlan.Spec.ExternalID != servicePlan.Spec.ExternalID {
		errMsg := fmt.Sprintf("%s already exists with OSB guid %q, received different guid %q", pretty.ServicePlanName(servicePlan), existingServicePlan.Spec.ExternalID, servicePlan.Spec.ExternalID)
		klog.Error(pcb.Message(errMsg))
		return fmt.Errorf(errMsg)
	}
	klog.V(5).Info(pcb.Messagef("Found existing %s; updating", pretty.ServicePlanName(servicePlan)))
	toUpdate := existingServicePlan.DeepCopy()
	toUpdate.Spec.Description = servicePlan.Spec.Description
	toUpdate.Spec.Bindable = servicePlan.Spec.Bindable
	toUpdate.Spec.Free = servicePlan.Spec.Free
	toUpdate.Spec.ExternalName = servicePlan.Spec.ExternalName
	toUpdate.Spec.ExternalMetadata = servicePlan.Spec.ExternalMetadata
	toUpdate.Spec.InstanceCreateParameterSchema = servicePlan.Spec.InstanceCreateParameterSchema
	toUpdate.Spec.InstanceUpdateParameterSchema = servicePlan.Spec.InstanceUpdateParameterSchema
	toUpdate.Spec.ServiceBindingCreateParameterSchema = servicePlan.Spec.ServiceBindingCreateParameterSchema
	updatedPlan, err := c.serviceCatalogClient.ServicePlans(broker.Namespace).Update(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating %s: %v", pretty.ServicePlanName(servicePlan), err))
		return err
	}
	if updatedPlan.Status.RemovedFromBrokerCatalog {
		updatedPlan.Status.RemovedFromBrokerCatalog = false
		klog.V(4).Info(pcb.Messagef("Resetting RemovedFromBrokerCatalog status on %s", pretty.ServicePlanName(updatedPlan)))
		_, err := c.serviceCatalogClient.ServicePlans(broker.Namespace).UpdateStatus(updatedPlan)
		if err != nil {
			s := fmt.Sprintf("Error updating status of %s: %v", pretty.ServicePlanName(updatedPlan), err)
			klog.Error(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
			if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
				return err
			}
			return err
		}
	}
	return nil
}
func updateCommonStatusCondition(pcb *pretty.ContextBuilder, meta metav1.ObjectMeta, commonStatus *v1beta1.CommonServiceBrokerStatus, conditionType v1beta1.ServiceBrokerConditionType, status v1beta1.ConditionStatus, reason, message string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	newCondition := v1beta1.ServiceBrokerCondition{Type: conditionType, Status: status, Reason: reason, Message: message}
	t := time.Now()
	if len(commonStatus.Conditions) == 0 {
		klog.Info(pcb.Messagef("Setting lastTransitionTime for condition %q to %v", conditionType, t))
		newCondition.LastTransitionTime = metav1.NewTime(t)
		commonStatus.Conditions = []v1beta1.ServiceBrokerCondition{newCondition}
	} else {
		for i, cond := range commonStatus.Conditions {
			if cond.Type == conditionType {
				if cond.Status != newCondition.Status {
					klog.Info(pcb.Messagef("Found status change for condition %q: %q -> %q; setting lastTransitionTime to %v", conditionType, cond.Status, status, t))
					newCondition.LastTransitionTime = metav1.NewTime(t)
				} else {
					newCondition.LastTransitionTime = cond.LastTransitionTime
				}
				commonStatus.Conditions[i] = newCondition
				break
			}
		}
	}
	if conditionType == v1beta1.ServiceBrokerConditionReady && status == v1beta1.ConditionTrue {
		commonStatus.ReconciledGeneration = meta.Generation
		now := metav1.NewTime(t)
		commonStatus.LastCatalogRetrievalTime = &now
	}
}
func (c *controller) updateServiceBrokerCondition(broker *v1beta1.ServiceBroker, conditionType v1beta1.ServiceBrokerConditionType, status v1beta1.ConditionStatus, reason, message string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	toUpdate := broker.DeepCopy()
	pcb := pretty.NewServiceBrokerContextBuilder(toUpdate)
	updateCommonStatusCondition(pcb, toUpdate.ObjectMeta, &toUpdate.Status.CommonServiceBrokerStatus, conditionType, status, reason, message)
	klog.V(4).Info(pcb.Messagef("Updating ready condition to %v", status))
	_, err := c.serviceCatalogClient.ServiceBrokers(broker.Namespace).UpdateStatus(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating ready condition: %v", err))
	} else {
		klog.V(5).Info(pcb.Messagef("Updated ready condition to %v", status))
	}
	return err
}
func (c *controller) updateServiceBrokerFinalizers(broker *v1beta1.ServiceBroker, finalizers []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewServiceBrokerContextBuilder(broker)
	broker, err := c.serviceCatalogClient.ServiceBrokers(broker.Namespace).Get(broker.Name, metav1.GetOptions{})
	if err != nil {
		klog.Error(pcb.Messagef("Error finalizing: %v", err))
	}
	toUpdate := broker.DeepCopy()
	toUpdate.Finalizers = finalizers
	logContext := fmt.Sprint(pcb.Messagef("Updating finalizers to %v", finalizers))
	klog.V(4).Info(pcb.Messagef("Updating %v", logContext))
	_, err = c.serviceCatalogClient.ServiceBrokers(broker.Namespace).UpdateStatus(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating %v: %v", logContext, err))
	}
	return err
}
func (c *controller) getCurrentServiceClassesAndPlansForNamespacedBroker(broker *v1beta1.ServiceBroker) ([]v1beta1.ServiceClass, []v1beta1.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{v1beta1.FilterSpecServiceBrokerName: broker.Name}
	fieldSelector := fields.SelectorFromSet(fieldSet).String()
	listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
	existingServiceClasses, err := c.serviceCatalogClient.ServiceClasses(broker.Namespace).List(listOpts)
	if err != nil {
		c.recorder.Eventf(broker, corev1.EventTypeWarning, errorListingServiceClassesReason, "%v %v", errorListingServiceClassesMessage, err)
		if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorListingServiceClassesReason, errorListingServiceClassesMessage); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}
	existingServicePlans, err := c.serviceCatalogClient.ServicePlans(broker.Namespace).List(listOpts)
	if err != nil {
		c.recorder.Eventf(broker, corev1.EventTypeWarning, errorListingServicePlansReason, "%v %v", errorListingServicePlansMessage, err)
		if err := c.updateServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorListingServicePlansReason, errorListingServicePlansMessage); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}
	return existingServiceClasses.Items, existingServicePlans.Items, nil
}
func convertServiceClassListToMap(list []v1beta1.ServiceClass) map[string]*v1beta1.ServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := make(map[string]*v1beta1.ServiceClass, len(list))
	for i := range list {
		ret[list[i].Name] = &list[i]
	}
	return ret
}
func convertServicePlanListToMap(list []v1beta1.ServicePlan) map[string]*v1beta1.ServicePlan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := make(map[string]*v1beta1.ServicePlan, len(list))
	for i := range list {
		ret[list[i].Name] = &list[i]
	}
	return ret
}
