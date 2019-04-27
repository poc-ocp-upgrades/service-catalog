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

func (c *controller) servicePlanAdd(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("ServicePlan: Couldn't get key for object %+v: %v", obj, err)
		return
	}
	c.servicePlanQueue.Add(key)
}
func (c *controller) servicePlanUpdate(oldObj, newObj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c.servicePlanAdd(newObj)
}
func (c *controller) servicePlanDelete(obj interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	servicePlan, ok := obj.(*v1beta1.ServicePlan)
	if servicePlan == nil || !ok {
		return
	}
	klog.V(4).Infof("ServicePlan: Received delete event for %v; no further processing will occur", servicePlan.Name)
}
func (c *controller) reconcileServicePlanKey(key string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	pcb := pretty.NewContextBuilder(pretty.ServicePlan, namespace, name, "")
	plan, err := c.servicePlanLister.ServicePlans(namespace).Get(key)
	if errors.IsNotFound(err) {
		klog.Infof(pcb.Message("not doing work because plan has been deleted"))
		return nil
	}
	if err != nil {
		klog.Infof(pcb.Message("unable to retrieve object from store: %v"))
		return err
	}
	return c.reconcileServicePlan(plan)
}
func (c *controller) reconcileServicePlan(servicePlan *v1beta1.ServicePlan) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := pretty.NewContextBuilder(pretty.ServicePlan, servicePlan.Namespace, servicePlan.Name, "")
	klog.Infof("ServicePlan %q (ExternalName: %q): processing", servicePlan.Name, servicePlan.Spec.ExternalName)
	if !servicePlan.Status.RemovedFromBrokerCatalog {
		return nil
	}
	klog.Infof(pcb.Message("removed from broker catalog; determining whether there are instances remaining"))
	serviceInstances, err := c.findServiceInstancesOnServicePlan(servicePlan)
	if err != nil {
		return err
	}
	if len(serviceInstances.Items) != 0 {
		return nil
	}
	klog.Infof(pcb.Message("removed from broker catalog and has zero instances remaining; deleting"))
	return c.serviceCatalogClient.ServicePlans(servicePlan.Namespace).Delete(servicePlan.Name, &metav1.DeleteOptions{})
}
func (c *controller) findServiceInstancesOnServicePlan(servicePlan *v1beta1.ServicePlan) (*v1beta1.ServiceInstanceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fieldSet := fields.Set{"spec.servicePlanRef.name": servicePlan.Name}
	fieldSelector := fields.SelectorFromSet(fieldSet).String()
	listOpts := metav1.ListOptions{FieldSelector: fieldSelector}
	return c.serviceCatalogClient.ServiceInstances(metav1.NamespaceAll).List(listOpts)
}
