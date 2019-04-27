package controller

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func (c *controller) clusterServicePlanAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("ClusterServicePlan: Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.clusterServicePlanQueue.Add(key)
}
func (c *controller) clusterServicePlanUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.clusterServicePlanAdd(newObj)
}
func (c *controller) clusterServicePlanDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	clusterServicePlan, ok := obj.(*v1beta1.ClusterServicePlan)
	if clusterServicePlan == nil || !ok {
		return
	}
	klog.V(4).Infof("ClusterServicePlan: Received delete event for %v; no further processing will occur", clusterServicePlan.Name)
}
func (c *controller) reconcileClusterServicePlanKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	plan, err := c.clusterServicePlanLister.Get(key)
	if errors.IsNotFound(err) {
		klog.Infof("ClusterServicePlan %q: Not doing work because it has been deleted", key)
		return nil
	}
	if err != nil {
		klog.Infof("ClusterServicePlan %q: Unable to retrieve object from store: %v", key, err)
		return err
	}
	return c.reconcileClusterServicePlan(plan)
}
func (c *controller) reconcileClusterServicePlan(clusterServicePlan *v1beta1.ClusterServicePlan) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Infof("ClusterServicePlan %q (ExternalName: %q): processing", clusterServicePlan.Name, clusterServicePlan.Spec.ExternalName)
	if !clusterServicePlan.Status.RemovedFromBrokerCatalog {
		return nil
	}
	klog.Infof("ClusterServicePlan %q (ExternalName: %q): has been removed from broker catalog; determining whether there are instances remaining", clusterServicePlan.Name, clusterServicePlan.Spec.ExternalName)
	serviceInstances, err := c.findServiceInstancesOnClusterServicePlan(clusterServicePlan)
	if err != nil {
		return err
	}
	if len(serviceInstances.Items) != 0 {
		return nil
	}
	klog.Infof("ClusterServicePlan %q (ExternalName: %q): has been removed from broker catalog and has zero instances remaining; deleting", clusterServicePlan.Name, clusterServicePlan.Spec.ExternalName)
	return c.serviceCatalogClient.ClusterServicePlans().Delete(clusterServicePlan.Name, &metav1.DeleteOptions{})
}
func (c *controller) findServiceInstancesOnClusterServicePlan(clusterServicePlan *v1beta1.ClusterServicePlan) (*v1beta1.ServiceInstanceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{"spec.clusterServicePlanRef.name": clusterServicePlan.Name}
	fieldSelector := fields.SelectorFromSet(fieldSet).String()
	listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
	return c.serviceCatalogClient.ServiceInstances(metav1.NamespaceAll).List(listOpts)
}
