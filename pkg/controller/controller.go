package controller

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
	osb "github.com/pmorie/go-open-service-broker-client/v2"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeutil "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/wait"
	corev1 "k8s.io/api/core/v1"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	servicecatalogclientset "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
	informers "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/externalversions/servicecatalog/v1beta1"
	listers "github.com/kubernetes-incubator/service-catalog/pkg/client/listers_generated/servicecatalog/v1beta1"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	"github.com/kubernetes-incubator/service-catalog/pkg/filter"
	"github.com/kubernetes-incubator/service-catalog/pkg/pretty"
)

const (
	maxRetries									= 15
	pollingStartInterval						= 1 * time.Second
	ContextProfilePlatformKubernetes	string	= "kubernetes"
	DefaultClusterIDConfigMapName		string	= "cluster-info"
	DefaultClusterIDConfigMapNamespace	string	= "default"
)

func NewController(kubeClient kubernetes.Interface, serviceCatalogClient servicecatalogclientset.ServicecatalogV1beta1Interface, clusterServiceBrokerInformer informers.ClusterServiceBrokerInformer, serviceBrokerInformer informers.ServiceBrokerInformer, clusterServiceClassInformer informers.ClusterServiceClassInformer, serviceClassInformer informers.ServiceClassInformer, instanceInformer informers.ServiceInstanceInformer, bindingInformer informers.ServiceBindingInformer, clusterServicePlanInformer informers.ClusterServicePlanInformer, servicePlanInformer informers.ServicePlanInformer, brokerClientCreateFunc osb.CreateFunc, brokerRelistInterval time.Duration, osbAPIPreferredVersion string, recorder record.EventRecorder, reconciliationRetryDuration time.Duration, operationPollingMaximumBackoffDuration time.Duration, clusterIDConfigMapName string, clusterIDConfigMapNamespace string) (Controller, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	controller := &controller{kubeClient: kubeClient, serviceCatalogClient: serviceCatalogClient, brokerRelistInterval: brokerRelistInterval, OSBAPIPreferredVersion: osbAPIPreferredVersion, recorder: recorder, reconciliationRetryDuration: reconciliationRetryDuration, clusterServiceBrokerQueue: workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(pollingStartInterval, operationPollingMaximumBackoffDuration), "cluster-service-broker"), serviceBrokerQueue: workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(pollingStartInterval, operationPollingMaximumBackoffDuration), "service-broker"), clusterServiceClassQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "cluster-service-class"), serviceClassQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "service-class"), clusterServicePlanQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "cluster-service-plan"), servicePlanQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "service-plan"), instanceQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "service-instance"), bindingQueue: workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "service-binding"), instancePollingQueue: workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(pollingStartInterval, operationPollingMaximumBackoffDuration), "instance-poller"), bindingPollingQueue: workqueue.NewNamedRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(pollingStartInterval, operationPollingMaximumBackoffDuration), "binding-poller"), clusterIDConfigMapName: clusterIDConfigMapName, clusterIDConfigMapNamespace: clusterIDConfigMapNamespace, brokerClientManager: NewBrokerClientManager(brokerClientCreateFunc)}
	controller.clusterServiceBrokerLister = clusterServiceBrokerInformer.Lister()
	clusterServiceBrokerInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.clusterServiceBrokerAdd, UpdateFunc: controller.clusterServiceBrokerUpdate, DeleteFunc: controller.clusterServiceBrokerDelete})
	controller.clusterServiceClassLister = clusterServiceClassInformer.Lister()
	clusterServiceClassInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.clusterServiceClassAdd, UpdateFunc: controller.clusterServiceClassUpdate, DeleteFunc: controller.clusterServiceClassDelete})
	controller.clusterServicePlanLister = clusterServicePlanInformer.Lister()
	clusterServicePlanInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.clusterServicePlanAdd, UpdateFunc: controller.clusterServicePlanUpdate, DeleteFunc: controller.clusterServicePlanDelete})
	controller.instanceLister = instanceInformer.Lister()
	instanceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.instanceAdd, UpdateFunc: controller.instanceUpdate, DeleteFunc: controller.instanceDelete})
	controller.bindingLister = bindingInformer.Lister()
	bindingInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.bindingAdd, UpdateFunc: controller.bindingUpdate, DeleteFunc: controller.bindingDelete})
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.NamespacedServiceBroker) {
		controller.serviceBrokerLister = serviceBrokerInformer.Lister()
		serviceBrokerInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.serviceBrokerAdd, UpdateFunc: controller.serviceBrokerUpdate, DeleteFunc: controller.serviceBrokerDelete})
		controller.serviceClassLister = serviceClassInformer.Lister()
		serviceClassInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.serviceClassAdd, UpdateFunc: controller.serviceClassUpdate, DeleteFunc: controller.serviceClassDelete})
		controller.servicePlanLister = servicePlanInformer.Lister()
		servicePlanInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{AddFunc: controller.servicePlanAdd, UpdateFunc: controller.servicePlanUpdate, DeleteFunc: controller.servicePlanDelete})
	}
	controller.instanceOperationRetryQueue.instances = make(map[string]backoffEntry)
	controller.instanceOperationRetryQueue.rateLimiter = workqueue.NewItemExponentialFailureRateLimiter(minBrokerOperationRetryDelay, maxBrokerOperationRetryDelay)
	return controller, nil
}

type Controller interface {
	Run(workers int, stopCh <-chan struct{})
}
type controller struct {
	kubeClient					kubernetes.Interface
	serviceCatalogClient		servicecatalogclientset.ServicecatalogV1beta1Interface
	clusterServiceBrokerLister	listers.ClusterServiceBrokerLister
	serviceBrokerLister			listers.ServiceBrokerLister
	clusterServiceClassLister	listers.ClusterServiceClassLister
	serviceClassLister			listers.ServiceClassLister
	instanceLister				listers.ServiceInstanceLister
	bindingLister				listers.ServiceBindingLister
	clusterServicePlanLister	listers.ClusterServicePlanLister
	servicePlanLister			listers.ServicePlanLister
	brokerRelistInterval		time.Duration
	OSBAPIPreferredVersion		string
	recorder					record.EventRecorder
	reconciliationRetryDuration	time.Duration
	clusterServiceBrokerQueue	workqueue.RateLimitingInterface
	serviceBrokerQueue			workqueue.RateLimitingInterface
	clusterServiceClassQueue	workqueue.RateLimitingInterface
	serviceClassQueue			workqueue.RateLimitingInterface
	clusterServicePlanQueue		workqueue.RateLimitingInterface
	servicePlanQueue			workqueue.RateLimitingInterface
	instanceQueue				workqueue.RateLimitingInterface
	bindingQueue				workqueue.RateLimitingInterface
	instancePollingQueue		workqueue.RateLimitingInterface
	bindingPollingQueue			workqueue.RateLimitingInterface
	clusterIDConfigMapName		string
	clusterIDConfigMapNamespace	string
	clusterID					string
	clusterIDLock				sync.RWMutex
	instanceOperationRetryQueue	instanceOperationBackoff
	brokerClientManager			*BrokerClientManager
}

func (c *controller) Run(workers int, stopCh <-chan struct{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	defer runtimeutil.HandleCrash()
	klog.Info("Starting service-catalog controller")
	var waitGroup sync.WaitGroup
	for i := 0; i < workers; i++ {
		createWorker(c.clusterServiceBrokerQueue, "ClusterServiceBroker", maxRetries, true, c.reconcileClusterServiceBrokerKey, stopCh, &waitGroup)
		createWorker(c.clusterServiceClassQueue, "ClusterServiceClass", maxRetries, true, c.reconcileClusterServiceClassKey, stopCh, &waitGroup)
		createWorker(c.clusterServicePlanQueue, "ClusterServicePlan", maxRetries, true, c.reconcileClusterServicePlanKey, stopCh, &waitGroup)
		createWorker(c.instanceQueue, "ServiceInstance", maxRetries, true, c.reconcileServiceInstanceKey, stopCh, &waitGroup)
		createWorker(c.bindingQueue, "ServiceBinding", maxRetries, true, c.reconcileServiceBindingKey, stopCh, &waitGroup)
		createWorker(c.instancePollingQueue, "InstancePoller", maxRetries, false, c.requeueServiceInstanceForPoll, stopCh, &waitGroup)
		if utilfeature.DefaultFeatureGate.Enabled(scfeatures.NamespacedServiceBroker) {
			createWorker(c.serviceBrokerQueue, "ServiceBroker", maxRetries, true, c.reconcileServiceBrokerKey, stopCh, &waitGroup)
			createWorker(c.serviceClassQueue, "ServiceClass", maxRetries, true, c.reconcileServiceClassKey, stopCh, &waitGroup)
			createWorker(c.servicePlanQueue, "ServicePlan", maxRetries, true, c.reconcileServicePlanKey, stopCh, &waitGroup)
		}
		if utilfeature.DefaultFeatureGate.Enabled(scfeatures.AsyncBindingOperations) {
			createWorker(c.bindingPollingQueue, "BindingPoller", maxRetries, false, c.requeueServiceBindingForPoll, stopCh, &waitGroup)
		}
	}
	c.createConfigMapMonitorWorker(stopCh, &waitGroup)
	c.createPurgeExpiredRetryEntriesWorker(stopCh, &waitGroup)
	<-stopCh
	klog.Info("Shutting down service-catalog controller")
	c.clusterServiceBrokerQueue.ShutDown()
	c.clusterServiceClassQueue.ShutDown()
	c.clusterServicePlanQueue.ShutDown()
	c.instanceQueue.ShutDown()
	c.bindingQueue.ShutDown()
	c.instancePollingQueue.ShutDown()
	c.bindingPollingQueue.ShutDown()
	if utilfeature.DefaultFeatureGate.Enabled(scfeatures.NamespacedServiceBroker) {
		c.serviceBrokerQueue.ShutDown()
		c.serviceClassQueue.ShutDown()
		c.servicePlanQueue.ShutDown()
	}
	waitGroup.Wait()
	klog.Info("Shutdown service-catalog controller")
}
func createWorker(queue workqueue.RateLimitingInterface, resourceType string, maxRetries int, forgetAfterSuccess bool, reconciler func(key string) error, stopCh <-chan struct{}, waitGroup *sync.WaitGroup) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	waitGroup.Add(1)
	go func() {
		wait.Until(worker(queue, resourceType, maxRetries, forgetAfterSuccess, reconciler), time.Second, stopCh)
		waitGroup.Done()
	}()
}
func (c *controller) createConfigMapMonitorWorker(stopCh <-chan struct{}, waitGroup *sync.WaitGroup) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	waitGroup.Add(1)
	go func() {
		wait.Until(c.monitorConfigMap, 15*time.Second, stopCh)
		waitGroup.Done()
	}()
}
func (c *controller) createPurgeExpiredRetryEntriesWorker(stopCh <-chan struct{}, waitGroup *sync.WaitGroup) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	waitGroup.Add(1)
	go func() {
		wait.Until(c.purgeExpiredRetryEntries, 2*maxBrokerOperationRetryDelay, stopCh)
		waitGroup.Done()
	}()
}
func (c *controller) monitorConfigMap() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.V(9).Info("cluster ID monitor loop enter")
	cm, err := c.kubeClient.CoreV1().ConfigMaps(c.clusterIDConfigMapNamespace).Get(c.clusterIDConfigMapName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		m := make(map[string]string)
		m["id"] = c.getClusterID()
		cm := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: c.clusterIDConfigMapName}, Data: m}
		if _, err := c.kubeClient.CoreV1().ConfigMaps(c.clusterIDConfigMapNamespace).Create(cm); err != nil {
			klog.Warningf("due to error %q, could not set clusterid configmap to %#v ", err, cm)
		}
	} else if err == nil {
		if id := cm.Data["id"]; "" != id {
			c.setClusterID(id)
		} else {
			m := cm.Data
			if m == nil {
				m = make(map[string]string)
				cm.Data = m
			}
			m["id"] = c.getClusterID()
			c.kubeClient.CoreV1().ConfigMaps(c.clusterIDConfigMapNamespace).Update(cm)
		}
	} else {
		klog.V(4).Infof("error getting the cluster info configmap: %q", err)
	}
	klog.V(9).Info("cluster ID monitor loop exit")
}
func worker(queue workqueue.RateLimitingInterface, resourceType string, maxRetries int, forgetAfterSuccess bool, reconciler func(key string) error) func() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() {
		exit := false
		for !exit {
			exit = func() bool {
				key, quit := queue.Get()
				if quit {
					return true
				}
				defer queue.Done(key)
				err := reconciler(key.(string))
				if err == nil {
					if forgetAfterSuccess {
						queue.Forget(key)
					}
					return false
				}
				numRequeues := queue.NumRequeues(key)
				if numRequeues < maxRetries {
					klog.V(4).Infof("Error syncing %s %v (retry: %d/%d): %v", resourceType, key, numRequeues, maxRetries, err)
					queue.AddRateLimited(key)
					return false
				}
				klog.V(4).Infof("Dropping %s %q out of the queue: %v", resourceType, key, err)
				queue.Forget(key)
				return false
			}()
		}
	}
}

type operationError struct {
	reason	string
	message	string
}

func (e *operationError) Error() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return e.message
}
func (c *controller) getClusterServiceClassPlanAndClusterServiceBroker(instance *v1beta1.ServiceInstance) (*v1beta1.ClusterServiceClass, *v1beta1.ClusterServicePlan, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, brokerName, brokerClient, err := c.getClusterServiceClassAndClusterServiceBroker(instance)
	if err != nil {
		return nil, nil, "", nil, err
	}
	var servicePlan *v1beta1.ClusterServicePlan
	if instance.Spec.ClusterServicePlanRef != nil {
		var err error
		servicePlan, err = c.clusterServicePlanLister.Get(instance.Spec.ClusterServicePlanRef.Name)
		if nil != err {
			return nil, nil, "", nil, &operationError{reason: errorNonexistentClusterServicePlanReason, message: fmt.Sprintf("The instance references a non-existent ClusterServicePlan %q - %v", instance.Spec.ClusterServicePlanRef.Name, instance.Spec.PlanReference)}
		}
	}
	return serviceClass, servicePlan, brokerName, brokerClient, nil
}
func (c *controller) getServiceClassPlanAndServiceBroker(instance *v1beta1.ServiceInstance) (*v1beta1.ServiceClass, *v1beta1.ServicePlan, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, brokerName, brokerClient, err := c.getServiceClassAndServiceBroker(instance)
	if err != nil {
		return nil, nil, "", nil, err
	}
	var servicePlan *v1beta1.ServicePlan
	if instance.Spec.ServicePlanRef != nil {
		var err error
		servicePlan, err = c.servicePlanLister.ServicePlans(instance.Namespace).Get(instance.Spec.ServicePlanRef.Name)
		if nil != err {
			return nil, nil, "", nil, &operationError{reason: errorNonexistentServicePlanReason, message: fmt.Sprintf("The instance references a non-existent ServicePlan %q - %v", instance.Spec.ServicePlanRef.Name, instance.Spec.PlanReference)}
		}
	}
	return serviceClass, servicePlan, brokerName, brokerClient, nil
}
func (c *controller) getClusterServiceClassAndClusterServiceBroker(instance *v1beta1.ServiceInstance) (*v1beta1.ClusterServiceClass, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, err := c.clusterServiceClassLister.Get(instance.Spec.ClusterServiceClassRef.Name)
	if err != nil {
		return nil, "", nil, &operationError{reason: errorNonexistentClusterServiceClassReason, message: fmt.Sprintf("The instance references a non-existent ClusterServiceClass (K8S: %q ExternalName: %q)", instance.Spec.ClusterServiceClassRef.Name, instance.Spec.ClusterServiceClassExternalName)}
	}
	broker, err := c.clusterServiceBrokerLister.Get(serviceClass.Spec.ClusterServiceBrokerName)
	if err != nil {
		return nil, "", nil, &operationError{reason: errorNonexistentClusterServiceBrokerReason, message: fmt.Sprintf("The instance references a non-existent broker %q", serviceClass.Spec.ClusterServiceBrokerName)}
	}
	brokerClient, found := c.brokerClientManager.BrokerClient(NewClusterServiceBrokerKey(serviceClass.Spec.ClusterServiceBrokerName))
	if !found {
		return nil, "", nil, &operationError{reason: errorNonexistentClusterServiceBrokerReason, message: fmt.Sprintf("The instance references a broker %q which has no OSB client created", serviceClass.Spec.ClusterServiceBrokerName)}
	}
	return serviceClass, broker.Name, brokerClient, nil
}
func (c *controller) getServiceClassAndServiceBroker(instance *v1beta1.ServiceInstance) (*v1beta1.ServiceClass, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, err := c.serviceClassLister.ServiceClasses(instance.Namespace).Get(instance.Spec.ServiceClassRef.Name)
	if err != nil {
		return nil, "", nil, &operationError{reason: errorNonexistentServiceClassReason, message: fmt.Sprintf("The instance references a non-existent ServiceClass (K8S: %q ExternalName: %q)", instance.Spec.ServiceClassRef.Name, instance.Spec.ServiceClassExternalName)}
	}
	broker, err := c.serviceBrokerLister.ServiceBrokers(instance.Namespace).Get(serviceClass.Spec.ServiceBrokerName)
	if err != nil {
		return nil, "", nil, &operationError{reason: errorNonexistentServiceBrokerReason, message: fmt.Sprintf("The instance references a non-existent broker %q", serviceClass.Spec.ServiceBrokerName)}
	}
	brokerClient, found := c.brokerClientManager.BrokerClient(NewServiceBrokerKey(instance.Namespace, serviceClass.Spec.ServiceBrokerName))
	if !found {
		return nil, "", nil, &operationError{reason: errorNonexistentClusterServiceBrokerReason, message: fmt.Sprintf("The instance references a broker %q which has no OSB client created", serviceClass.Spec.ServiceBrokerName)}
	}
	return serviceClass, broker.Name, brokerClient, nil
}
func (c *controller) getClusterServiceClassPlanAndClusterServiceBrokerForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (*v1beta1.ClusterServiceClass, *v1beta1.ClusterServicePlan, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, serviceBrokerName, osbClient, err := c.getClusterServiceClassAndClusterServiceBrokerForServiceBinding(instance, binding)
	if err != nil {
		return nil, nil, "", nil, err
	}
	servicePlan, err := c.getClusterServicePlanForServiceBinding(instance, binding, serviceClass)
	if err != nil {
		return nil, nil, "", nil, err
	}
	return serviceClass, servicePlan, serviceBrokerName, osbClient, nil
}
func (c *controller) getClusterServiceClassAndClusterServiceBrokerForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (*v1beta1.ClusterServiceClass, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, err := c.getClusterServiceClassForServiceBinding(instance, binding)
	if err != nil {
		return nil, "", nil, err
	}
	serviceBroker, err := c.getClusterServiceBrokerForServiceBinding(instance, binding, serviceClass)
	if err != nil {
		return nil, "", nil, err
	}
	osbClient, err := c.getBrokerClientForServiceBinding(instance, binding)
	if err != nil {
		return nil, "", nil, err
	}
	return serviceClass, serviceBroker.Name, osbClient, nil
}
func (c *controller) getClusterServiceClassForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (*v1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	serviceClass, err := c.clusterServiceClassLister.Get(instance.Spec.ClusterServiceClassRef.Name)
	if err != nil {
		s := fmt.Sprintf("References a non-existent ClusterServiceClass %q - %c", instance.Spec.ClusterServiceClassRef.Name, instance.Spec.PlanReference)
		klog.Warning(pcb.Message(s))
		c.updateServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServiceClassReason, "The binding references a ClusterServiceClass that does not exist. "+s)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorNonexistentClusterServiceClassMessage, s)
		return nil, err
	}
	return serviceClass, nil
}
func (c *controller) getClusterServicePlanForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding, serviceClass *v1beta1.ClusterServiceClass) (*v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	servicePlan, err := c.clusterServicePlanLister.Get(instance.Spec.ClusterServicePlanRef.Name)
	if nil != err {
		s := fmt.Sprintf("References a non-existent ClusterServicePlan %q - %v", instance.Spec.ClusterServicePlanRef.Name, instance.Spec.PlanReference)
		klog.Warning(pcb.Message(s))
		c.updateServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServicePlanReason, "The ServiceBinding references an ServiceInstance which references ClusterServicePlan that does not exist. "+s)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorNonexistentClusterServicePlanReason, s)
		return nil, fmt.Errorf(s)
	}
	return servicePlan, nil
}
func (c *controller) getClusterServiceBrokerForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding, serviceClass *v1beta1.ClusterServiceClass) (*v1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	broker, err := c.clusterServiceBrokerLister.Get(serviceClass.Spec.ClusterServiceBrokerName)
	if err != nil {
		s := fmt.Sprintf("References a non-existent ClusterServiceBroker %q", serviceClass.Spec.ClusterServiceBrokerName)
		klog.Warning(pcb.Message(s))
		c.updateServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServiceBrokerReason, "The binding references a ClusterServiceBroker that does not exist. "+s)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorNonexistentClusterServiceBrokerReason, s)
		return nil, err
	}
	return broker, nil
}
func (c *controller) getBrokerClientForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var brokerClient osb.Client
	if instance.Spec.ClusterServiceClassSpecified() {
		serviceClass, err := c.getClusterServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, err
		}
		broker, err := c.getClusterServiceBrokerForServiceBinding(instance, binding, serviceClass)
		if err != nil {
			return nil, err
		}
		var found bool
		brokerClient, found = c.brokerClientManager.BrokerClient(NewClusterServiceBrokerKey(broker.Name))
		if !found {
			return nil, fmt.Errorf("OSB client not found for the broker %s", broker.Name)
		}
	} else if instance.Spec.ServiceClassSpecified() {
		serviceClass, err := c.getServiceClassForServiceBinding(instance, binding)
		if err != nil {
			return nil, err
		}
		broker, err := c.getServiceBrokerForServiceBinding(instance, binding, serviceClass)
		if err != nil {
			return nil, err
		}
		var found bool
		brokerClient, found = c.brokerClientManager.BrokerClient(NewServiceBrokerKey(broker.Namespace, broker.Name))
		if !found {
			return nil, fmt.Errorf("OSB client not found for the broker %s", broker.Name)
		}
	}
	return brokerClient, nil
}
func getAuthCredentialsFromClusterServiceBroker(client kubernetes.Interface, broker *v1beta1.ClusterServiceBroker) (*osb.AuthConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if broker.Spec.AuthInfo == nil {
		return nil, nil
	}
	authInfo := broker.Spec.AuthInfo
	if authInfo.Basic != nil {
		secretRef := authInfo.Basic.SecretRef
		secret, err := client.CoreV1().Secrets(secretRef.Namespace).Get(secretRef.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		basicAuthConfig, err := getBasicAuthConfig(secret)
		if err != nil {
			return nil, err
		}
		return &osb.AuthConfig{BasicAuthConfig: basicAuthConfig}, nil
	} else if authInfo.Bearer != nil {
		secretRef := authInfo.Bearer.SecretRef
		secret, err := client.CoreV1().Secrets(secretRef.Namespace).Get(secretRef.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		bearerConfig, err := getBearerConfig(secret)
		if err != nil {
			return nil, err
		}
		return &osb.AuthConfig{BearerConfig: bearerConfig}, nil
	}
	return nil, fmt.Errorf("empty auth info or unsupported auth mode: %s", authInfo)
}
func getAuthCredentialsFromServiceBroker(client kubernetes.Interface, broker *v1beta1.ServiceBroker) (*osb.AuthConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if broker.Spec.AuthInfo == nil {
		return nil, nil
	}
	authInfo := broker.Spec.AuthInfo
	if authInfo.Basic != nil {
		secretRef := authInfo.Basic.SecretRef
		secret, err := client.CoreV1().Secrets(broker.Namespace).Get(secretRef.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		basicAuthConfig, err := getBasicAuthConfig(secret)
		if err != nil {
			return nil, err
		}
		return &osb.AuthConfig{BasicAuthConfig: basicAuthConfig}, nil
	} else if authInfo.Bearer != nil {
		secretRef := authInfo.Bearer.SecretRef
		secret, err := client.CoreV1().Secrets(broker.Namespace).Get(secretRef.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		bearerConfig, err := getBearerConfig(secret)
		if err != nil {
			return nil, err
		}
		return &osb.AuthConfig{BearerConfig: bearerConfig}, nil
	}
	return nil, fmt.Errorf("empty auth info or unsupported auth mode: %s", authInfo)
}
func getBasicAuthConfig(secret *corev1.Secret) (*osb.BasicAuthConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	usernameBytes, ok := secret.Data["username"]
	if !ok {
		return nil, fmt.Errorf("auth secret didn't contain username")
	}
	passwordBytes, ok := secret.Data["password"]
	if !ok {
		return nil, fmt.Errorf("auth secret didn't contain password")
	}
	return &osb.BasicAuthConfig{Username: string(usernameBytes), Password: string(passwordBytes)}, nil
}
func getBearerConfig(secret *corev1.Secret) (*osb.BearerConfig, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tokenBytes, ok := secret.Data["token"]
	if !ok {
		return nil, fmt.Errorf("auth secret didn't contain token")
	}
	return &osb.BearerConfig{Token: string(tokenBytes)}, nil
}
func convertAndFilterCatalogToNamespacedTypes(namespace string, in *osb.CatalogResponse, restrictions *v1beta1.CatalogRestrictions, existingServiceClasses map[string]*v1beta1.ServiceClass, existingServicePlans map[string]*v1beta1.ServicePlan) ([]*v1beta1.ServiceClass, []*v1beta1.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicate filter.Predicate
	var err error
	if restrictions != nil && len(restrictions.ServiceClass) > 0 {
		predicate, err = filter.CreatePredicate(restrictions.ServiceClass)
		if err != nil {
			return nil, nil, err
		}
	} else {
		predicate = filter.NewPredicate()
	}
	serviceClasses := []*v1beta1.ServiceClass(nil)
	servicePlans := []*v1beta1.ServicePlan(nil)
	for _, svc := range in.Services {
		serviceClass := &v1beta1.ServiceClass{Spec: v1beta1.ServiceClassSpec{CommonServiceClassSpec: v1beta1.CommonServiceClassSpec{Bindable: svc.Bindable, PlanUpdatable: svc.PlanUpdatable != nil && *svc.PlanUpdatable, ExternalID: svc.ID, ExternalName: svc.Name, Tags: svc.Tags, Description: svc.Description, Requires: svc.Requires}}}
		if utilfeature.DefaultFeatureGate.Enabled(scfeatures.AsyncBindingOperations) {
			serviceClass.Spec.BindingRetrievable = svc.BindingsRetrievable
		}
		if svc.Metadata != nil {
			metadata, err := json.Marshal(svc.Metadata)
			if err != nil {
				err = fmt.Errorf("Failed to marshal metadata\n%+v\n %v", svc.Metadata, err)
				klog.Error(err)
				return nil, nil, err
			}
			serviceClass.Spec.ExternalMetadata = &runtime.RawExtension{Raw: metadata}
		}
		if existingServiceClasses[svc.ID] != nil {
			serviceClass.SetName(existingServiceClasses[svc.ID].Name)
		} else {
			serviceClass.SetName(GenerateEscapedName(svc.ID))
		}
		serviceClass.SetNamespace(namespace)
		if fields := v1beta1.ConvertServiceClassToProperties(serviceClass); predicate.Accepts(fields) {
			plans, err := convertServicePlans(namespace, svc.Plans, serviceClass.Name, existingServicePlans)
			if err != nil {
				return nil, nil, err
			}
			acceptedPlans, _, err := filterNamespacedServicePlans(restrictions, plans)
			if err != nil {
				return nil, nil, err
			}
			if len(acceptedPlans) > 0 {
				serviceClasses = append(serviceClasses, serviceClass)
				servicePlans = append(servicePlans, acceptedPlans...)
			}
		}
	}
	return serviceClasses, servicePlans, nil
}
func GenerateEscapedName(externalID string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	buffer := bytes.Buffer{}
	lenOrigin := len(externalID)
	prevDot := false
	prevDash := false
	for i, ch := range externalID {
		if (ch >= 'a' && ch <= 'y') || (ch >= '0' && ch <= '9') {
			buffer.WriteString(string(ch))
			prevDash = false
			prevDot = false
		} else if ch == '.' && i != 0 && i != lenOrigin-1 && !prevDot && !prevDash {
			buffer.WriteString(string(ch))
			prevDash = false
			prevDot = true
		} else if ch == '-' && i != 0 && i != lenOrigin-1 && !prevDot {
			buffer.WriteString(string(ch))
			prevDash = true
			prevDot = false
		} else {
			start, end := "z", "z"
			buffer.WriteString(fmt.Sprintf("%s%x%s", start, ch, end))
			prevDash = false
			prevDot = false
		}
	}
	escapedName := buffer.String()
	if len(escapedName) > validation.DNS1123LabelMaxLength {
		escapedName = escapedName[0:30]
		escapedName = strings.TrimSuffix(escapedName, ".")
		escapedName = escapedName + "-" + fmt.Sprintf("%x", md5.Sum([]byte(externalID)))
	}
	return escapedName
}
func convertAndFilterCatalog(in *osb.CatalogResponse, restrictions *v1beta1.CatalogRestrictions, existingServiceClasses map[string]*v1beta1.ClusterServiceClass, existingServicePlans map[string]*v1beta1.ClusterServicePlan) ([]*v1beta1.ClusterServiceClass, []*v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicate filter.Predicate
	var err error
	if restrictions != nil && len(restrictions.ServiceClass) > 0 {
		predicate, err = filter.CreatePredicate(restrictions.ServiceClass)
		if err != nil {
			return nil, nil, err
		}
	} else {
		predicate = filter.NewPredicate()
	}
	serviceClasses := []*v1beta1.ClusterServiceClass(nil)
	servicePlans := []*v1beta1.ClusterServicePlan(nil)
	for _, svc := range in.Services {
		serviceClass := &v1beta1.ClusterServiceClass{Spec: v1beta1.ClusterServiceClassSpec{CommonServiceClassSpec: v1beta1.CommonServiceClassSpec{Bindable: svc.Bindable, PlanUpdatable: svc.PlanUpdatable != nil && *svc.PlanUpdatable, ExternalID: svc.ID, ExternalName: svc.Name, Tags: svc.Tags, Description: svc.Description, Requires: svc.Requires}}}
		if utilfeature.DefaultFeatureGate.Enabled(scfeatures.AsyncBindingOperations) {
			serviceClass.Spec.BindingRetrievable = svc.BindingsRetrievable
		}
		if svc.Metadata != nil {
			metadata, err := json.Marshal(svc.Metadata)
			if err != nil {
				err = fmt.Errorf("Failed to marshal metadata\n%+v\n %v", svc.Metadata, err)
				klog.Error(err)
				return nil, nil, err
			}
			serviceClass.Spec.ExternalMetadata = &runtime.RawExtension{Raw: metadata}
		}
		if existingServiceClasses[svc.ID] != nil {
			serviceClass.SetName(existingServiceClasses[svc.ID].Name)
		} else {
			serviceClass.SetName(GenerateEscapedName(svc.ID))
		}
		if fields := v1beta1.ConvertClusterServiceClassToProperties(serviceClass); predicate.Accepts(fields) {
			plans, err := convertClusterServicePlans(svc.Plans, serviceClass.Name, existingServicePlans)
			if err != nil {
				return nil, nil, err
			}
			acceptedPlans, _, err := filterServicePlans(restrictions, plans)
			if err != nil {
				return nil, nil, err
			}
			if len(acceptedPlans) > 0 {
				serviceClasses = append(serviceClasses, serviceClass)
				servicePlans = append(servicePlans, acceptedPlans...)
			}
		}
	}
	return serviceClasses, servicePlans, nil
}
func filterNamespacedServicePlans(restrictions *v1beta1.CatalogRestrictions, servicePlans []*v1beta1.ServicePlan) ([]*v1beta1.ServicePlan, []*v1beta1.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicate filter.Predicate
	var err error
	if restrictions != nil && len(restrictions.ServicePlan) > 0 {
		predicate, err = filter.CreatePredicate(restrictions.ServicePlan)
		if err != nil {
			return nil, nil, err
		}
	} else {
		predicate = filter.NewPredicate()
	}
	if predicate.Empty() {
		return servicePlans, []*v1beta1.ServicePlan(nil), nil
	}
	accepted := []*v1beta1.ServicePlan(nil)
	rejected := []*v1beta1.ServicePlan(nil)
	for _, sp := range servicePlans {
		fields := v1beta1.ConvertServicePlanToProperties(sp)
		if predicate.Accepts(fields) {
			accepted = append(accepted, sp)
		} else {
			rejected = append(rejected, sp)
		}
	}
	return accepted, rejected, nil
}
func filterServicePlans(restrictions *v1beta1.CatalogRestrictions, servicePlans []*v1beta1.ClusterServicePlan) ([]*v1beta1.ClusterServicePlan, []*v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var predicate filter.Predicate
	var err error
	if restrictions != nil && len(restrictions.ServicePlan) > 0 {
		predicate, err = filter.CreatePredicate(restrictions.ServicePlan)
		if err != nil {
			return nil, nil, err
		}
	} else {
		predicate = filter.NewPredicate()
	}
	if predicate.Empty() {
		return servicePlans, []*v1beta1.ClusterServicePlan(nil), nil
	}
	accepted := []*v1beta1.ClusterServicePlan(nil)
	rejected := []*v1beta1.ClusterServicePlan(nil)
	for _, sp := range servicePlans {
		fields := v1beta1.ConvertClusterServicePlanToProperties(sp)
		if predicate.Accepts(fields) {
			accepted = append(accepted, sp)
		} else {
			rejected = append(rejected, sp)
		}
	}
	return accepted, rejected, nil
}
func convertServicePlans(namespace string, plans []osb.Plan, serviceClassID string, existingServicePlans map[string]*v1beta1.ServicePlan) ([]*v1beta1.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if 0 == len(plans) {
		return nil, fmt.Errorf("ServiceClass (K8S: %q) must have at least one plan", serviceClassID)
	}
	servicePlans := make([]*v1beta1.ServicePlan, len(plans))
	for i, plan := range plans {
		servicePlan := &v1beta1.ServicePlan{Spec: v1beta1.ServicePlanSpec{CommonServicePlanSpec: v1beta1.CommonServicePlanSpec{ExternalName: plan.Name, ExternalID: plan.ID, Free: plan.Free != nil && *plan.Free, Description: plan.Description}, ServiceClassRef: v1beta1.LocalObjectReference{Name: serviceClassID}}}
		servicePlans[i] = servicePlan
		if existingServicePlans[plan.ID] != nil {
			servicePlans[i].SetName(existingServicePlans[plan.ID].Name)
		} else {
			servicePlans[i].SetName(GenerateEscapedName(plan.ID))
		}
		servicePlan.SetNamespace(namespace)
		err := convertCommonServicePlan(plan, &servicePlan.Spec.CommonServicePlanSpec)
		if err != nil {
			return nil, err
		}
	}
	return servicePlans, nil
}
func convertCommonServicePlan(plan osb.Plan, commonServicePlanSpec *v1beta1.CommonServicePlanSpec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if plan.Bindable != nil {
		b := plan.Bindable
		commonServicePlanSpec.Bindable = b
	}
	if plan.Metadata != nil {
		metadata, err := json.Marshal(plan.Metadata)
		if err != nil {
			err = fmt.Errorf("Failed to marshal metadata\n%+v\n %v", plan.Metadata, err)
			klog.Error(err)
			return err
		}
		commonServicePlanSpec.ExternalMetadata = &runtime.RawExtension{Raw: metadata}
	}
	if schemas := plan.Schemas; schemas != nil {
		if instanceSchemas := schemas.ServiceInstance; instanceSchemas != nil {
			if instanceCreateSchema := instanceSchemas.Create; instanceCreateSchema != nil && instanceCreateSchema.Parameters != nil {
				schema, err := json.Marshal(instanceCreateSchema.Parameters)
				if err != nil {
					err = fmt.Errorf("Failed to marshal instance create schema \n%+v\n %v", instanceCreateSchema.Parameters, err)
					klog.Error(err)
					return err
				}
				commonServicePlanSpec.InstanceCreateParameterSchema = &runtime.RawExtension{Raw: schema}
			}
			if instanceUpdateSchema := instanceSchemas.Update; instanceUpdateSchema != nil && instanceUpdateSchema.Parameters != nil {
				schema, err := json.Marshal(instanceUpdateSchema.Parameters)
				if err != nil {
					err = fmt.Errorf("Failed to marshal instance update schema \n%+v\n %v", instanceUpdateSchema.Parameters, err)
					klog.Error(err)
					return err
				}
				commonServicePlanSpec.InstanceUpdateParameterSchema = &runtime.RawExtension{Raw: schema}
			}
		}
		if bindingSchemas := schemas.ServiceBinding; bindingSchemas != nil {
			if bindingCreateSchema := bindingSchemas.Create; bindingCreateSchema != nil {
				if bindingCreateSchema.Parameters != nil {
					schema, err := json.Marshal(bindingCreateSchema.Parameters)
					if err != nil {
						err = fmt.Errorf("Failed to marshal binding create schema \n%+v\n %v", bindingCreateSchema.Parameters, err)
						klog.Error(err)
						return err
					}
					commonServicePlanSpec.ServiceBindingCreateParameterSchema = &runtime.RawExtension{Raw: schema}
				}
				if utilfeature.DefaultFeatureGate.Enabled(scfeatures.ResponseSchema) && bindingCreateSchema.Response != nil {
					schema, err := json.Marshal(bindingCreateSchema.Response)
					if err != nil {
						err = fmt.Errorf("Failed to marshal binding create response schema \n%+v\n %v", bindingCreateSchema.Response, err)
						klog.Error(err)
						return err
					}
					commonServicePlanSpec.ServiceBindingCreateResponseSchema = &runtime.RawExtension{Raw: schema}
				}
			}
		}
	}
	return nil
}
func convertClusterServicePlans(plans []osb.Plan, serviceClassID string, existingServicePlans map[string]*v1beta1.ClusterServicePlan) ([]*v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if 0 == len(plans) {
		return nil, fmt.Errorf("ClusterServiceClass (K8S: %q) must have at least one plan", serviceClassID)
	}
	servicePlans := make([]*v1beta1.ClusterServicePlan, len(plans))
	for i, plan := range plans {
		servicePlans[i] = &v1beta1.ClusterServicePlan{Spec: v1beta1.ClusterServicePlanSpec{CommonServicePlanSpec: v1beta1.CommonServicePlanSpec{ExternalName: plan.Name, ExternalID: plan.ID, Free: plan.Free != nil && *plan.Free, Description: plan.Description}, ClusterServiceClassRef: v1beta1.ClusterObjectReference{Name: serviceClassID}}}
		if existingServicePlans[plan.ID] != nil {
			servicePlans[i].SetName(existingServicePlans[plan.ID].Name)
		} else {
			servicePlans[i].SetName(GenerateEscapedName(plan.ID))
		}
		if plan.Bindable != nil {
			b := *plan.Bindable
			servicePlans[i].Spec.Bindable = &b
		}
		if plan.Metadata != nil {
			metadata, err := json.Marshal(plan.Metadata)
			if err != nil {
				err = fmt.Errorf("Failed to marshal metadata\n%+v\n %v", plan.Metadata, err)
				klog.Error(err)
				return nil, err
			}
			servicePlans[i].Spec.ExternalMetadata = &runtime.RawExtension{Raw: metadata}
		}
		if schemas := plan.Schemas; schemas != nil {
			if instanceSchemas := schemas.ServiceInstance; instanceSchemas != nil {
				if instanceCreateSchema := instanceSchemas.Create; instanceCreateSchema != nil && instanceCreateSchema.Parameters != nil {
					schema, err := json.Marshal(instanceCreateSchema.Parameters)
					if err != nil {
						err = fmt.Errorf("Failed to marshal instance create schema \n%+v\n %v", instanceCreateSchema.Parameters, err)
						klog.Error(err)
						return nil, err
					}
					servicePlans[i].Spec.InstanceCreateParameterSchema = &runtime.RawExtension{Raw: schema}
				}
				if instanceUpdateSchema := instanceSchemas.Update; instanceUpdateSchema != nil && instanceUpdateSchema.Parameters != nil {
					schema, err := json.Marshal(instanceUpdateSchema.Parameters)
					if err != nil {
						err = fmt.Errorf("Failed to marshal instance update schema \n%+v\n %v", instanceUpdateSchema.Parameters, err)
						klog.Error(err)
						return nil, err
					}
					servicePlans[i].Spec.InstanceUpdateParameterSchema = &runtime.RawExtension{Raw: schema}
				}
			}
			if bindingSchemas := schemas.ServiceBinding; bindingSchemas != nil {
				if bindingCreateSchema := bindingSchemas.Create; bindingCreateSchema != nil {
					if bindingCreateSchema.Parameters != nil {
						schema, err := json.Marshal(bindingCreateSchema.Parameters)
						if err != nil {
							err = fmt.Errorf("Failed to marshal binding create schema \n%+v\n %v", bindingCreateSchema.Parameters, err)
							klog.Error(err)
							return nil, err
						}
						servicePlans[i].Spec.ServiceBindingCreateParameterSchema = &runtime.RawExtension{Raw: schema}
					}
					if utilfeature.DefaultFeatureGate.Enabled(scfeatures.ResponseSchema) && bindingCreateSchema.Response != nil {
						schema, err := json.Marshal(bindingCreateSchema.Response)
						if err != nil {
							err = fmt.Errorf("Failed to marshal binding create response schema \n%+v\n %v", bindingCreateSchema.Response, err)
							klog.Error(err)
							return nil, err
						}
						servicePlans[i].Spec.ServiceBindingCreateResponseSchema = &runtime.RawExtension{Raw: schema}
					}
				}
			}
		}
	}
	return servicePlans, nil
}
func isServiceInstanceConditionTrue(instance *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, cond := range instance.Status.Conditions {
		if cond.Type == conditionType {
			return cond.Status == v1beta1.ConditionTrue
		}
	}
	return false
}
func isServiceInstanceReady(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isServiceInstanceConditionTrue(instance, v1beta1.ServiceInstanceConditionReady)
}
func isServiceInstanceFailed(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isServiceInstanceConditionTrue(instance, v1beta1.ServiceInstanceConditionFailed)
}
func isServiceInstanceOrphanMitigation(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isServiceInstanceConditionTrue(instance, v1beta1.ServiceInstanceConditionOrphanMitigation)
}
func NewClientConfigurationForBroker(meta metav1.ObjectMeta, commonSpec *v1beta1.CommonServiceBrokerSpec, authConfig *osb.AuthConfig) *osb.ClientConfiguration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clientConfig := osb.DefaultClientConfiguration()
	clientConfig.Name = meta.Name
	clientConfig.URL = commonSpec.URL
	clientConfig.AuthConfig = authConfig
	clientConfig.EnableAlphaFeatures = true
	clientConfig.Insecure = commonSpec.InsecureSkipTLSVerify
	clientConfig.CAData = commonSpec.CABundle
	return clientConfig
}
func (c *controller) reconciliationRetryDurationExceeded(operationStartTime *metav1.Time) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if operationStartTime == nil || time.Now().Before(operationStartTime.Time.Add(c.reconciliationRetryDuration)) {
		return false
	}
	return true
}
func shouldStartOrphanMitigation(statusCode int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	is2XX := statusCode >= 200 && statusCode < 300
	is5XX := statusCode >= 500 && statusCode < 600
	return (is2XX && statusCode != http.StatusOK) || is5XX
}
func isRetriableHTTPStatus(statusCode int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return statusCode != http.StatusBadRequest
}

type ReconciliationAction string

const (
	reconcileAdd	ReconciliationAction	= "Add"
	reconcileUpdate	ReconciliationAction	= "Update"
	reconcileDelete	ReconciliationAction	= "Delete"
	reconcilePoll	ReconciliationAction	= "Poll"
)

func (c *controller) getClusterID() (id string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.clusterIDLock.RLock()
	id = c.clusterID
	c.clusterIDLock.RUnlock()
	if id != "" {
		return
	}
	c.clusterIDLock.Lock()
	if id = c.clusterID; id == "" {
		id = string(uuid.NewUUID())
		c.clusterID = id
	}
	c.clusterIDLock.Unlock()
	return
}
func (c *controller) setClusterID(id string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.clusterIDLock.Lock()
	c.clusterID = id
	c.clusterIDLock.Unlock()
}
func (c *controller) getServiceClassPlanAndServiceBrokerForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (*v1beta1.ServiceClass, *v1beta1.ServicePlan, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, serviceBrokerName, osbClient, err := c.getServiceClassAndServiceBrokerForServiceBinding(instance, binding)
	if err != nil {
		return nil, nil, "", nil, err
	}
	servicePlan, err := c.getServicePlanForServiceBinding(instance, binding, serviceClass)
	if err != nil {
		return nil, nil, "", nil, err
	}
	return serviceClass, servicePlan, serviceBrokerName, osbClient, nil
}
func (c *controller) getServiceClassAndServiceBrokerForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (*v1beta1.ServiceClass, string, osb.Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, err := c.getServiceClassForServiceBinding(instance, binding)
	if err != nil {
		return nil, "", nil, err
	}
	serviceBroker, err := c.getServiceBrokerForServiceBinding(instance, binding, serviceClass)
	if err != nil {
		return nil, "", nil, err
	}
	osbClient, err := c.getBrokerClientForServiceBinding(instance, binding)
	if err != nil {
		return nil, "", nil, err
	}
	return serviceClass, serviceBroker.Name, osbClient, nil
}
func (c *controller) getServiceClassForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding) (*v1beta1.ServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	serviceClass, err := c.serviceClassLister.ServiceClasses(instance.Namespace).Get(instance.Spec.ServiceClassRef.Name)
	if err != nil {
		s := fmt.Sprintf("References a non-existent ServiceClass %q - %c", instance.Spec.ServiceClassRef.Name, instance.Spec.PlanReference)
		klog.Warning(pcb.Message(s))
		c.updateServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServiceClassReason, "The binding references a ServiceClass that does not exist. "+s)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorNonexistentClusterServiceClassMessage, s)
		return nil, err
	}
	return serviceClass, nil
}
func (c *controller) getServicePlanForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding, serviceClass *v1beta1.ServiceClass) (*v1beta1.ServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	servicePlan, err := c.servicePlanLister.ServicePlans(instance.Namespace).Get(instance.Spec.ServicePlanRef.Name)
	if nil != err {
		s := fmt.Sprintf("References a non-existent ServicePlan %q - %v", instance.Spec.ServicePlanRef.Name, instance.Spec.PlanReference)
		klog.Warning(pcb.Message(s))
		c.updateServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServicePlanReason, "The ServiceBinding references an ServiceInstance which references ServicePlan that does not exist. "+s)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorNonexistentClusterServicePlanReason, s)
		return nil, fmt.Errorf(s)
	}
	return servicePlan, nil
}
func (c *controller) getServiceBrokerForServiceBinding(instance *v1beta1.ServiceInstance, binding *v1beta1.ServiceBinding, serviceClass *v1beta1.ServiceClass) (*v1beta1.ServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewInstanceContextBuilder(instance)
	broker, err := c.serviceBrokerLister.ServiceBrokers(instance.Namespace).Get(serviceClass.Spec.ServiceBrokerName)
	if err != nil {
		s := fmt.Sprintf("References a non-existent ServiceBroker %q", serviceClass.Spec.ServiceBrokerName)
		klog.Warning(pcb.Message(s))
		c.updateServiceBindingCondition(binding, v1beta1.ServiceBindingConditionReady, v1beta1.ConditionFalse, errorNonexistentClusterServiceBrokerReason, "The binding references a ServiceBroker that does not exist. "+s)
		c.recorder.Event(binding, corev1.EventTypeWarning, errorNonexistentClusterServiceBrokerReason, s)
		return nil, err
	}
	return broker, nil
}
func shouldReconcileServiceBrokerCommon(pcb *pretty.ContextBuilder, brokerMeta *metav1.ObjectMeta, brokerSpec *v1beta1.CommonServiceBrokerSpec, brokerStatus *v1beta1.CommonServiceBrokerStatus, now time.Time, defaultRelistInterval time.Duration) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if brokerStatus.ReconciledGeneration != brokerMeta.Generation {
		return true
	}
	if brokerMeta.DeletionTimestamp != nil || len(brokerStatus.Conditions) == 0 {
		return true
	}
	for _, condition := range brokerStatus.Conditions {
		if condition.Type == v1beta1.ServiceBrokerConditionReady {
			if condition.Status == v1beta1.ConditionTrue {
				if brokerSpec.RelistBehavior == v1beta1.ServiceBrokerRelistBehaviorManual {
					klog.V(10).Info(pcb.Message("Not processing because RelistBehavior is set to Manual"))
					return false
				}
				duration := defaultRelistInterval
				if brokerSpec.RelistDuration != nil {
					duration = brokerSpec.RelistDuration.Duration
				}
				intervalPassed := true
				if brokerStatus.LastCatalogRetrievalTime != nil {
					intervalPassed = now.After(brokerStatus.LastCatalogRetrievalTime.Time.Add(duration))
				}
				if intervalPassed == false {
					klog.V(10).Info(pcb.Message("Not processing because RelistDuration has not elapsed since the last relist"))
				}
				return intervalPassed
			}
			return true
		}
	}
	return true
}
func toJSON(obj interface{}) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, _ := json.Marshal(obj)
	return string(bytes)
}
