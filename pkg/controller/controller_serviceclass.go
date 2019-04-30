package controller

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/klog"
	"github.com/kubernetes-incubator/service-catalog/pkg/pretty"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func (c *controller) serviceClassAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.serviceClassQueue.Add(key)
}
func (c *controller) serviceClassUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.serviceClassAdd(newObj)
}
func (c *controller) serviceClassDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, ok := obj.(*v1beta1.ServiceClass)
	if serviceClass == nil || !ok {
		return
	}
	klog.V(4).Infof("Received delete event for ServiceClass %v; no further processing will occur", serviceClass.Name)
}
func (c *controller) reconcileServiceClassKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	pcb := pretty.NewContextBuilder(pretty.ServiceClass, namespace, name, "")
	class, err := c.serviceClassLister.ServiceClasses(namespace).Get(name)
	if errors.IsNotFound(err) {
		klog.Info(pcb.Message("Not doing work because the ServiceClass has been deleted"))
		return nil
	}
	if err != nil {
		klog.Infof(pcb.Message("Unable to retrieve"))
		return err
	}
	return c.reconcileServiceClass(class)
}
func (c *controller) reconcileServiceClass(serviceClass *v1beta1.ServiceClass) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewContextBuilder(pretty.ServiceClass, serviceClass.Namespace, serviceClass.Name, "")
	klog.Info(pcb.Message("Processing"))
	if !serviceClass.Status.RemovedFromBrokerCatalog {
		return nil
	}
	klog.Info(pcb.Message("Removed from broker catalog; determining whether there are instances remaining"))
	serviceInstances, err := c.findServiceInstancesOnServiceClass(serviceClass)
	if err != nil {
		return err
	}
	if len(serviceInstances.Items) != 0 {
		return nil
	}
	klog.Info(pcb.Message("Removed from broker catalog and has zero instances remaining; deleting"))
	return c.serviceCatalogClient.ServiceClasses(serviceClass.Namespace).Delete(serviceClass.Name, &metav1.DeleteOptions{})
}
func (c *controller) findServiceInstancesOnServiceClass(serviceClass *v1beta1.ServiceClass) (*v1beta1.ServiceInstanceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{"spec.serviceClassRef.name": serviceClass.Name}
	fieldSelector := fields.SelectorFromSet(fieldSet).String()
	listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
	return c.serviceCatalogClient.ServiceInstances(serviceClass.Namespace).List(listOpts)
}
