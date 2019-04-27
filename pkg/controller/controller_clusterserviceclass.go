package controller

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func (c *controller) clusterServiceClassAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.clusterServiceClassQueue.Add(key)
}
func (c *controller) clusterServiceClassUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.clusterServiceClassAdd(newObj)
}
func (c *controller) clusterServiceClassDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	serviceClass, ok := obj.(*v1beta1.ClusterServiceClass)
	if serviceClass == nil || !ok {
		return
	}
	klog.V(4).Infof("Received delete event for ServiceClass %v; no further processing will occur", serviceClass.Name)
}
func (c *controller) reconcileClusterServiceClassKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	class, err := c.clusterServiceClassLister.Get(key)
	if errors.IsNotFound(err) {
		klog.Infof("ClusterServiceClass %q: Not doing work because it has been deleted", key)
		return nil
	}
	if err != nil {
		klog.Infof("ClusterServiceClass %q: Unable to retrieve object from store: %v", key, err)
		return err
	}
	return c.reconcileClusterServiceClass(class)
}
func (c *controller) reconcileClusterServiceClass(serviceClass *v1beta1.ClusterServiceClass) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("ClusterServiceClass %q (ExternalName: %q): processing", serviceClass.Name, serviceClass.Spec.ExternalName)
	if !serviceClass.Status.RemovedFromBrokerCatalog {
		return nil
	}
	klog.Infof("ClusterServiceClass %q (ExternalName: %q): has been removed from broker catalog; determining whether there are instances remaining", serviceClass.Name, serviceClass.Spec.ExternalName)
	serviceInstances, err := c.findServiceInstancesOnClusterServiceClass(serviceClass)
	if err != nil {
		return err
	}
	if len(serviceInstances.Items) != 0 {
		return nil
	}
	klog.Infof("ClusterServiceClass %q (ExternalName: %q): has been removed from broker catalog and has zero instances remaining; deleting", serviceClass.Name, serviceClass.Spec.ExternalName)
	return c.serviceCatalogClient.ClusterServiceClasses().Delete(serviceClass.Name, &metav1.DeleteOptions{})
}
func (c *controller) findServiceInstancesOnClusterServiceClass(serviceClass *v1beta1.ClusterServiceClass) (*v1beta1.ServiceInstanceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{"spec.clusterServiceClassRef.name": serviceClass.Name}
	fieldSelector := fields.SelectorFromSet(fieldSet).String()
	listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
	return c.serviceCatalogClient.ServiceInstances(metav1.NamespaceAll).List(listOpts)
}
