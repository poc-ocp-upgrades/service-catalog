package controller

import (
	"fmt"
	"strings"
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
	errorListingClusterServiceClassesReason		string	= "ErrorListingClusterServiceClasses"
	errorListingClusterServiceClassesMessage	string	= "Error listing cluster service classes."
	errorListingClusterServicePlansReason		string	= "ErrorListingClusterServicePlans"
	errorListingClusterServicePlansMessage		string	= "Error listing cluster service plans."
	errorDeletingClusterServiceClassReason		string	= "ErrorDeletingClusterServiceClass"
	errorDeletingClusterServiceClassMessage		string	= "Error deleting cluster service class."
	errorDeletingClusterServicePlanReason		string	= "ErrorDeletingClusterServicePlan"
	errorDeletingClusterServicePlanMessage		string	= "Error deleting cluster service plan."
	errorAuthCredentialsReason			string	= "ErrorGettingAuthCredentials"
	successClusterServiceBrokerDeletedReason	string	= "DeletedClusterServiceBrokerSuccessfully"
	successClusterServiceBrokerDeletedMessage	string	= "The broker %v was deleted successfully."
	errorFetchingCatalogReason			string	= "ErrorFetchingCatalog"
	errorFetchingCatalogMessage			string	= "Error fetching catalog."
	errorSyncingCatalogReason			string	= "ErrorSyncingCatalog"
	errorSyncingCatalogMessage			string	= "Error syncing catalog from ClusterServiceBroker."
	successFetchedCatalogReason			string	= "FetchedCatalog"
	successFetchedCatalogMessage			string	= "Successfully fetched catalog entries from broker."
	errorReconciliationRetryTimeoutReason		string	= "ErrorReconciliationRetryTimeout"
)

func (c *controller) clusterServiceBrokerAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.clusterServiceBrokerQueue.Add(key)
}
func (c *controller) clusterServiceBrokerUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.clusterServiceBrokerAdd(newObj)
}
func (c *controller) clusterServiceBrokerDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, ok := obj.(*v1beta1.ClusterServiceBroker)
	if broker == nil || !ok {
		return
	}
	klog.V(4).Infof("Received delete event for ClusterServiceBroker %v; no further processing will occur", broker.Name)
}
func shouldReconcileClusterServiceBroker(broker *v1beta1.ClusterServiceBroker, now time.Time, defaultRelistInterval time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return shouldReconcileServiceBrokerCommon(pretty.NewClusterServiceBrokerContextBuilder(broker), &broker.ObjectMeta, &broker.Spec.CommonServiceBrokerSpec, &broker.Status.CommonServiceBrokerStatus, now, defaultRelistInterval)
}
func (c *controller) reconcileClusterServiceBrokerKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	broker, err := c.clusterServiceBrokerLister.Get(key)
	pcb := pretty.NewContextBuilder(pretty.ClusterServiceBroker, "", key, "")
	klog.V(4).Info(pcb.Message("Processing service broker"))
	if errors.IsNotFound(err) {
		klog.Info(pcb.Message("Not doing work because it has been deleted"))
		c.brokerClientManager.RemoveBrokerClient(NewClusterServiceBrokerKey(key))
		return nil
	}
	if err != nil {
		klog.Info(pcb.Messagef("Unable to retrieve object from store: %v", err))
		return err
	}
	return c.reconcileClusterServiceBroker(broker)
}
func (c *controller) updateClusterServiceBrokerClient(broker *v1beta1.ClusterServiceBroker) (osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewClusterServiceBrokerContextBuilder(broker)
	klog.V(4).Info(pcb.Message("Updating broker client"))
	authConfig, err := getAuthCredentialsFromClusterServiceBroker(c.kubeClient, broker)
	if err != nil {
		s := fmt.Sprintf("Error getting broker auth credentials: %s", err)
		klog.Info(pcb.Message(s))
		c.recorder.Event(broker, corev1.EventTypeWarning, errorAuthCredentialsReason, s)
		if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorFetchingCatalogReason, errorFetchingCatalogMessage+s); err != nil {
			return nil, err
		}
		return nil, err
	}
	clientConfig := NewClientConfigurationForBroker(broker.ObjectMeta, &broker.Spec.CommonServiceBrokerSpec, authConfig)
	brokerClient, err := c.brokerClientManager.UpdateBrokerClient(NewClusterServiceBrokerKey(broker.Name), clientConfig)
	if err != nil {
		s := fmt.Sprintf("Error creating client for broker %q: %s", broker.Name, err)
		klog.Info(pcb.Message(s))
		c.recorder.Event(broker, corev1.EventTypeWarning, errorAuthCredentialsReason, s)
		if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorFetchingCatalogReason, errorFetchingCatalogMessage+s); err != nil {
			return nil, err
		}
		return nil, err
	}
	return brokerClient, nil
}
func (c *controller) reconcileClusterServiceBroker(broker *v1beta1.ClusterServiceBroker) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewClusterServiceBrokerContextBuilder(broker)
	klog.V(4).Infof(pcb.Message("Processing"))
	if !shouldReconcileClusterServiceBroker(broker, time.Now(), c.brokerRelistInterval) {
		return nil
	}
	if broker.DeletionTimestamp == nil {
		klog.V(4).Info(pcb.Message("Processing adding/update event"))
		brokerClient, err := c.updateClusterServiceBrokerClient(broker)
		if err != nil {
			return err
		}
		now := metav1.Now()
		brokerCatalog, err := brokerClient.GetCatalog()
		if err != nil {
			s := fmt.Sprintf("Error getting broker catalog: %s", err)
			klog.Warning(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorFetchingCatalogReason, s)
			if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorFetchingCatalogReason, errorFetchingCatalogMessage+s); err != nil {
				return err
			}
			if broker.Status.OperationStartTime == nil {
				toUpdate := broker.DeepCopy()
				toUpdate.Status.OperationStartTime = &now
				if _, err := c.serviceCatalogClient.ClusterServiceBrokers().UpdateStatus(toUpdate); err != nil {
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
				return c.updateClusterServiceBrokerCondition(toUpdate, v1beta1.ServiceBrokerConditionFailed, v1beta1.ConditionTrue, errorReconciliationRetryTimeoutReason, s)
			}
			return err
		}
		klog.V(5).Info(pcb.Messagef("Successfully fetched %v catalog entries", len(brokerCatalog.Services)))
		if broker.Status.OperationStartTime != nil {
			toUpdate := broker.DeepCopy()
			toUpdate.Status.OperationStartTime = nil
			if _, err := c.serviceCatalogClient.ClusterServiceBrokers().UpdateStatus(toUpdate); err != nil {
				klog.Error(pcb.Messagef("Error updating operation start time: %v", err))
				return err
			}
		}
		existingServiceClasses, existingServicePlans, err := c.getCurrentServiceClassesAndPlansForBroker(broker)
		if err != nil {
			return err
		}
		existingServiceClassMap := convertClusterServiceClassListToMap(existingServiceClasses)
		existingServicePlanMap := convertClusterServicePlanListToMap(existingServicePlans)
		klog.V(4).Info(pcb.Message("Converting catalog response into service-catalog API"))
		payloadServiceClasses, payloadServicePlans, err := convertAndFilterCatalog(brokerCatalog, broker.Spec.CatalogRestrictions, existingServiceClassMap, existingServicePlanMap)
		if err != nil {
			s := fmt.Sprintf("Error converting catalog payload for broker %q to service-catalog API: %s", broker.Name, err)
			klog.Warning(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
			if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
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
			klog.V(4).Info(pcb.Messagef("Reconciling %s", pretty.ClusterServiceClassName(payloadServiceClass)))
			if err := c.reconcileClusterServiceClassFromClusterServiceBrokerCatalog(broker, payloadServiceClass, existingServiceClass); err != nil {
				s := fmt.Sprintf("Error reconciling %s (broker %q): %s", pretty.ClusterServiceClassName(payloadServiceClass), broker.Name, err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
					return err
				}
				return err
			}
			klog.V(5).Info(pcb.Messagef("Reconciled %s", pretty.ClusterServiceClassName(payloadServiceClass)))
		}
		for _, existingServiceClass := range existingServiceClassMap {
			if existingServiceClass.Status.RemovedFromBrokerCatalog {
				continue
			}
			if !isServiceCatalogManagedResource(existingServiceClass) {
				continue
			}
			klog.V(4).Info(pcb.Messagef("%s has been removed from broker's catalog; marking", pretty.ClusterServiceClassName(existingServiceClass)))
			existingServiceClass.Status.RemovedFromBrokerCatalog = true
			_, err := c.serviceCatalogClient.ClusterServiceClasses().UpdateStatus(existingServiceClass)
			if err != nil {
				s := fmt.Sprintf("Error updating status of %s: %v", pretty.ClusterServiceClassName(existingServiceClass), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
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
			klog.V(4).Infof("ClusterServiceBroker %q: reconciling %s", broker.Name, pretty.ClusterServicePlanName(payloadServicePlan))
			if err := c.reconcileClusterServicePlanFromClusterServiceBrokerCatalog(broker, payloadServicePlan, existingServicePlan); err != nil {
				s := fmt.Sprintf("Error reconciling %s: %s", pretty.ClusterServicePlanName(payloadServicePlan), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s)
				return err
			}
			klog.V(5).Info(pcb.Messagef("Reconciled %s", pretty.ClusterServicePlanName(payloadServicePlan)))
		}
		for _, existingServicePlan := range existingServicePlanMap {
			if existingServicePlan.Status.RemovedFromBrokerCatalog {
				continue
			}
			if !isServiceCatalogManagedResource(existingServicePlan) {
				continue
			}
			klog.V(4).Info(pcb.Messagef("%s has been removed from broker's catalog; marking", pretty.ClusterServicePlanName(existingServicePlan)))
			existingServicePlan.Status.RemovedFromBrokerCatalog = true
			_, err := c.serviceCatalogClient.ClusterServicePlans().UpdateStatus(existingServicePlan)
			if err != nil {
				s := fmt.Sprintf("Error updating status of %s: %v", pretty.ClusterServicePlanName(existingServicePlan), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
				if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
					return err
				}
				return err
			}
		}
		if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionTrue, successFetchedCatalogReason, successFetchedCatalogMessage); err != nil {
			return err
		}
		c.recorder.Event(broker, corev1.EventTypeNormal, successFetchedCatalogReason, successFetchedCatalogMessage)
		metrics.BrokerServiceClassCount.WithLabelValues(broker.Name).Set(float64(len(payloadServiceClasses)))
		metrics.BrokerServicePlanCount.WithLabelValues(broker.Name).Set(float64(len(payloadServicePlans)))
		return nil
	}
	if finalizers := sets.NewString(broker.Finalizers...); finalizers.Has(v1beta1.FinalizerServiceCatalog) {
		klog.V(4).Info(pcb.Message("Finalizing"))
		existingServiceClasses, existingServicePlans, err := c.getCurrentServiceClassesAndPlansForBroker(broker)
		if err != nil {
			return err
		}
		klog.V(4).Info(pcb.Messagef("Found %d ClusterServiceClasses and %d ClusterServicePlans to delete", len(existingServiceClasses), len(existingServicePlans)))
		for _, plan := range existingServicePlans {
			klog.V(4).Info(pcb.Messagef("Deleting %s", pretty.ClusterServicePlanName(&plan)))
			err := c.serviceCatalogClient.ClusterServicePlans().Delete(plan.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				s := fmt.Sprintf("Error deleting %s: %s", pretty.ClusterServicePlanName(&plan), err)
				klog.Warning(pcb.Message(s))
				c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorDeletingClusterServicePlanMessage, errorDeletingClusterServicePlanReason+s)
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorDeletingClusterServicePlanReason, "%v %v", errorDeletingClusterServicePlanMessage, s)
				return err
			}
		}
		for _, svcClass := range existingServiceClasses {
			klog.V(4).Info(pcb.Messagef("Deleting %s", pretty.ClusterServiceClassName(&svcClass)))
			err = c.serviceCatalogClient.ClusterServiceClasses().Delete(svcClass.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				s := fmt.Sprintf("Error deleting %s: %s", pretty.ClusterServiceClassName(&svcClass), err)
				klog.Warning(pcb.Message(s))
				c.recorder.Eventf(broker, corev1.EventTypeWarning, errorDeletingClusterServiceClassReason, "%v %v", errorDeletingClusterServiceClassMessage, s)
				if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorDeletingClusterServiceClassMessage, errorDeletingClusterServiceClassReason+s); err != nil {
					return err
				}
				return err
			}
		}
		if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, successClusterServiceBrokerDeletedReason, "The broker was deleted successfully"); err != nil {
			return err
		}
		finalizers.Delete(v1beta1.FinalizerServiceCatalog)
		c.updateClusterServiceBrokerFinalizers(broker, finalizers.List())
		c.recorder.Eventf(broker, corev1.EventTypeNormal, successClusterServiceBrokerDeletedReason, successClusterServiceBrokerDeletedMessage, broker.Name)
		klog.V(5).Info(pcb.Message("Successfully deleted"))
		metrics.BrokerServiceClassCount.DeleteLabelValues(broker.Name)
		metrics.BrokerServicePlanCount.DeleteLabelValues(broker.Name)
		return nil
	}
	return nil
}
func (c *controller) reconcileClusterServiceClassFromClusterServiceBrokerCatalog(broker *v1beta1.ClusterServiceBroker, serviceClass, existingServiceClass *v1beta1.ClusterServiceClass) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewClusterServiceBrokerContextBuilder(broker)
	serviceClass.Spec.ClusterServiceBrokerName = broker.Name
	if existingServiceClass == nil {
		otherServiceClass, err := c.clusterServiceClassLister.Get(serviceClass.Name)
		if err != nil {
			if !errors.IsNotFound(err) {
				return err
			}
		} else {
			if otherServiceClass.Spec.ClusterServiceBrokerName != broker.Name {
				errMsg := fmt.Sprintf("%s already exists for Broker %q", pretty.ClusterServiceClassName(serviceClass), otherServiceClass.Spec.ClusterServiceBrokerName)
				klog.Error(pcb.Message(errMsg))
				return fmt.Errorf(errMsg)
			}
		}
		markAsServiceCatalogManagedResource(serviceClass, broker)
		klog.V(5).Info(pcb.Messagef("Fresh %s; creating", pretty.ClusterServiceClassName(serviceClass)))
		if _, err := c.serviceCatalogClient.ClusterServiceClasses().Create(serviceClass); err != nil {
			klog.Error(pcb.Messagef("Error creating %s: %v", pretty.ClusterServiceClassName(serviceClass), err))
			return err
		}
		return nil
	}
	if existingServiceClass.Spec.ExternalID != serviceClass.Spec.ExternalID {
		errMsg := fmt.Sprintf("%s already exists with OSB guid %q, received different guid %q", pretty.ClusterServiceClassName(serviceClass), existingServiceClass.Name, serviceClass.Name)
		klog.Error(pcb.Message(errMsg))
		return fmt.Errorf(errMsg)
	}
	klog.V(5).Info(pcb.Messagef("Found existing %s; updating", pretty.ClusterServiceClassName(serviceClass)))
	toUpdate := existingServiceClass.DeepCopy()
	toUpdate.Spec.BindingRetrievable = serviceClass.Spec.BindingRetrievable
	toUpdate.Spec.Bindable = serviceClass.Spec.Bindable
	toUpdate.Spec.PlanUpdatable = serviceClass.Spec.PlanUpdatable
	toUpdate.Spec.Tags = serviceClass.Spec.Tags
	toUpdate.Spec.Description = serviceClass.Spec.Description
	toUpdate.Spec.Requires = serviceClass.Spec.Requires
	toUpdate.Spec.ExternalName = serviceClass.Spec.ExternalName
	toUpdate.Spec.ExternalMetadata = serviceClass.Spec.ExternalMetadata
	markAsServiceCatalogManagedResource(toUpdate, broker)
	updatedServiceClass, err := c.serviceCatalogClient.ClusterServiceClasses().Update(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating %s: %v", pretty.ClusterServiceClassName(serviceClass), err))
		return err
	}
	if updatedServiceClass.Status.RemovedFromBrokerCatalog {
		klog.V(4).Info(pcb.Messagef("Resetting RemovedFromBrokerCatalog status on %s", pretty.ClusterServiceClassName(serviceClass)))
		updatedServiceClass.Status.RemovedFromBrokerCatalog = false
		_, err := c.serviceCatalogClient.ClusterServiceClasses().UpdateStatus(updatedServiceClass)
		if err != nil {
			s := fmt.Sprintf("Error updating status of %s: %v", pretty.ClusterServiceClassName(updatedServiceClass), err)
			klog.Warning(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
			if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
				return err
			}
			return err
		}
	}
	return nil
}
func (c *controller) reconcileClusterServicePlanFromClusterServiceBrokerCatalog(broker *v1beta1.ClusterServiceBroker, servicePlan, existingServicePlan *v1beta1.ClusterServicePlan) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewClusterServiceBrokerContextBuilder(broker)
	servicePlan.Spec.ClusterServiceBrokerName = broker.Name
	if existingServicePlan == nil {
		otherServicePlan, err := c.clusterServicePlanLister.Get(servicePlan.Name)
		if err != nil {
			if !errors.IsNotFound(err) {
				return err
			}
		} else {
			if otherServicePlan.Spec.ClusterServiceBrokerName != broker.Name {
				errMsg := fmt.Sprintf("%s already exists for Broker %q", pretty.ClusterServicePlanName(servicePlan), otherServicePlan.Spec.ClusterServiceBrokerName)
				klog.Error(pcb.Message(errMsg))
				return fmt.Errorf(errMsg)
			}
		}
		markAsServiceCatalogManagedResource(servicePlan, broker)
		if _, err := c.serviceCatalogClient.ClusterServicePlans().Create(servicePlan); err != nil {
			klog.Error(pcb.Messagef("Error creating %s: %v", pretty.ClusterServicePlanName(servicePlan), err))
			return err
		}
		return nil
	}
	if existingServicePlan.Spec.ExternalID != servicePlan.Spec.ExternalID {
		errMsg := fmt.Sprintf("%s already exists with OSB guid %q, received different guid %q", pretty.ClusterServicePlanName(servicePlan), existingServicePlan.Spec.ExternalID, servicePlan.Spec.ExternalID)
		klog.Error(pcb.Message(errMsg))
		return fmt.Errorf(errMsg)
	}
	klog.V(5).Info(pcb.Messagef("Found existing %s; updating", pretty.ClusterServicePlanName(servicePlan)))
	toUpdate := existingServicePlan.DeepCopy()
	toUpdate.Spec.Description = servicePlan.Spec.Description
	toUpdate.Spec.Bindable = servicePlan.Spec.Bindable
	toUpdate.Spec.Free = servicePlan.Spec.Free
	toUpdate.Spec.ExternalName = servicePlan.Spec.ExternalName
	toUpdate.Spec.ExternalMetadata = servicePlan.Spec.ExternalMetadata
	toUpdate.Spec.InstanceCreateParameterSchema = servicePlan.Spec.InstanceCreateParameterSchema
	toUpdate.Spec.InstanceUpdateParameterSchema = servicePlan.Spec.InstanceUpdateParameterSchema
	toUpdate.Spec.ServiceBindingCreateParameterSchema = servicePlan.Spec.ServiceBindingCreateParameterSchema
	markAsServiceCatalogManagedResource(toUpdate, broker)
	updatedPlan, err := c.serviceCatalogClient.ClusterServicePlans().Update(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating %s: %v", pretty.ClusterServicePlanName(servicePlan), err))
		return err
	}
	if updatedPlan.Status.RemovedFromBrokerCatalog {
		updatedPlan.Status.RemovedFromBrokerCatalog = false
		klog.V(4).Info(pcb.Messagef("Resetting RemovedFromBrokerCatalog status on %s", pretty.ClusterServicePlanName(updatedPlan)))
		_, err := c.serviceCatalogClient.ClusterServicePlans().UpdateStatus(updatedPlan)
		if err != nil {
			s := fmt.Sprintf("Error updating status of %s: %v", pretty.ClusterServicePlanName(updatedPlan), err)
			klog.Error(pcb.Message(s))
			c.recorder.Eventf(broker, corev1.EventTypeWarning, errorSyncingCatalogReason, s)
			if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionFalse, errorSyncingCatalogReason, errorSyncingCatalogMessage+s); err != nil {
				return err
			}
			return err
		}
	}
	return nil
}
func (c *controller) updateClusterServiceBrokerCondition(broker *v1beta1.ClusterServiceBroker, conditionType v1beta1.ServiceBrokerConditionType, status v1beta1.ConditionStatus, reason, message string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewClusterServiceBrokerContextBuilder(broker)
	toUpdate := broker.DeepCopy()
	newCondition := v1beta1.ServiceBrokerCondition{Type: conditionType, Status: status, Reason: reason, Message: message}
	t := time.Now()
	if len(broker.Status.Conditions) == 0 {
		klog.Info(pcb.Messagef("Setting lastTransitionTime for condition %q to %v", conditionType, t))
		newCondition.LastTransitionTime = metav1.NewTime(t)
		toUpdate.Status.Conditions = []v1beta1.ServiceBrokerCondition{newCondition}
	} else {
		for i, cond := range broker.Status.Conditions {
			if cond.Type == conditionType {
				if cond.Status != newCondition.Status {
					klog.Info(pcb.Messagef("Found status change for condition %q: %q -> %q; setting lastTransitionTime to %v", conditionType, cond.Status, status, t))
					newCondition.LastTransitionTime = metav1.NewTime(t)
				} else {
					newCondition.LastTransitionTime = cond.LastTransitionTime
				}
				toUpdate.Status.Conditions[i] = newCondition
				break
			}
		}
	}
	if conditionType == v1beta1.ServiceBrokerConditionReady && status == v1beta1.ConditionTrue {
		toUpdate.Status.ReconciledGeneration = toUpdate.Generation
		now := metav1.NewTime(t)
		toUpdate.Status.LastCatalogRetrievalTime = &now
	}
	klog.V(4).Info(pcb.Messagef("Updating ready condition to %v", status))
	_, err := c.serviceCatalogClient.ClusterServiceBrokers().UpdateStatus(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating ready condition: %v", err))
	} else {
		klog.V(5).Info(pcb.Messagef("Updated ready condition to %v", status))
	}
	return err
}
func (c *controller) updateClusterServiceBrokerFinalizers(broker *v1beta1.ClusterServiceBroker, finalizers []string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewClusterServiceBrokerContextBuilder(broker)
	broker, err := c.serviceCatalogClient.ClusterServiceBrokers().Get(broker.Name, metav1.GetOptions{})
	if err != nil {
		klog.Error(pcb.Messagef("Error finalizing: %v", err))
	}
	toUpdate := broker.DeepCopy()
	toUpdate.Finalizers = finalizers
	logContext := fmt.Sprint(pcb.Messagef("Updating finalizers to %v", finalizers))
	klog.V(4).Info(pcb.Messagef("Updating %v", logContext))
	_, err = c.serviceCatalogClient.ClusterServiceBrokers().UpdateStatus(toUpdate)
	if err != nil {
		klog.Error(pcb.Messagef("Error updating %v: %v", logContext, err))
	}
	return err
}
func (c *controller) getCurrentServiceClassesAndPlansForBroker(broker *v1beta1.ClusterServiceBroker) ([]v1beta1.ClusterServiceClass, []v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{"spec.clusterServiceBrokerName": broker.Name}
	fieldSelector := fields.SelectorFromSet(fieldSet).String()
	listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
	existingServiceClasses, err := c.serviceCatalogClient.ClusterServiceClasses().List(listOpts)
	if err != nil {
		c.recorder.Eventf(broker, corev1.EventTypeWarning, errorListingClusterServiceClassesReason, "%v %v", errorListingClusterServiceClassesMessage, err)
		if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorListingClusterServiceClassesReason, errorListingClusterServiceClassesMessage); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}
	existingServicePlans, err := c.serviceCatalogClient.ClusterServicePlans().List(listOpts)
	if err != nil {
		c.recorder.Eventf(broker, corev1.EventTypeWarning, errorListingClusterServicePlansReason, "%v %v", errorListingClusterServicePlansMessage, err)
		if err := c.updateClusterServiceBrokerCondition(broker, v1beta1.ServiceBrokerConditionReady, v1beta1.ConditionUnknown, errorListingClusterServicePlansReason, errorListingClusterServicePlansMessage); err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}
	return existingServiceClasses.Items, existingServicePlans.Items, nil
}
func convertClusterServiceClassListToMap(list []v1beta1.ClusterServiceClass) map[string]*v1beta1.ClusterServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := make(map[string]*v1beta1.ClusterServiceClass, len(list))
	for i := range list {
		ret[list[i].Name] = &list[i]
	}
	return ret
}
func convertClusterServicePlanListToMap(list []v1beta1.ClusterServicePlan) map[string]*v1beta1.ClusterServicePlan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ret := make(map[string]*v1beta1.ClusterServicePlan, len(list))
	for i := range list {
		ret[list[i].Name] = &list[i]
	}
	return ret
}
func markAsServiceCatalogManagedResource(obj metav1.Object, broker *v1beta1.ClusterServiceBroker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if isServiceCatalogManagedResource(obj) {
		return
	}
	var blockOwnerDeletion = false
	controllerRef := *metav1.NewControllerRef(broker, v1beta1.SchemeGroupVersion.WithKind("ClusterServiceBroker"))
	controllerRef.BlockOwnerDeletion = &blockOwnerDeletion
	obj.SetOwnerReferences(append(obj.GetOwnerReferences(), controllerRef))
}
func isServiceCatalogManagedResource(resource metav1.Object) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := metav1.GetControllerOf(resource)
	if c == nil {
		return false
	}
	return strings.HasPrefix(c.APIVersion, v1beta1.GroupName)
}
