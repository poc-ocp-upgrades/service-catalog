package servicecatalog

import (
	"fmt"
	"math"
	"time"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	FieldServicePlanRef = "spec.clusterServicePlanRef.name"
)

func (sdk *SDK) RetrieveInstances(ns, classFilter, planFilter string) (*v1beta1.ServiceInstanceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instances, err := sdk.ServiceCatalog().ServiceInstances(ns).List(v1.ListOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to list instances in %s", ns)
	}
	if classFilter == "" && planFilter == "" {
		return instances, nil
	}
	filtered := v1beta1.ServiceInstanceList{Items: []v1beta1.ServiceInstance{}}
	for _, instance := range instances.Items {
		if classFilter != "" && instance.Spec.GetSpecifiedClusterServiceClass() != classFilter {
			continue
		}
		if planFilter != "" && instance.Spec.GetSpecifiedClusterServicePlan() != planFilter {
			continue
		}
		filtered.Items = append(filtered.Items, instance)
	}
	return &filtered, nil
}
func (sdk *SDK) RetrieveInstance(ns, name string) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance, err := sdk.ServiceCatalog().ServiceInstances(ns).Get(name, v1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to get instance '%s.%s' (%s)", ns, name, err)
	}
	return instance, nil
}
func (sdk *SDK) RetrieveInstanceByBinding(b *v1beta1.ServiceBinding) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := b.Namespace
	instName := b.Spec.InstanceRef.Name
	inst, err := sdk.ServiceCatalog().ServiceInstances(ns).Get(instName, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return inst, nil
}
func (sdk *SDK) RetrieveInstancesByPlan(plan Plan) ([]v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	planOpts := v1.ListOptions{FieldSelector: fields.OneTermEqualSelector(FieldServicePlanRef, plan.GetName()).String()}
	instances, err := sdk.ServiceCatalog().ServiceInstances("").List(planOpts)
	if err != nil {
		return nil, fmt.Errorf("unable to list instances (%s)", err)
	}
	return instances.Items, nil
}
func (sdk *SDK) InstanceParentHierarchy(instance *v1beta1.ServiceInstance) (*v1beta1.ClusterServiceClass, *v1beta1.ClusterServicePlan, *v1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	class, plan, err := sdk.InstanceToServiceClassAndPlan(instance)
	if err != nil {
		return nil, nil, nil, err
	}
	broker, err := sdk.RetrieveBrokerByClass(class)
	if err != nil {
		return nil, nil, nil, err
	}
	return class, plan, broker, nil
}
func (sdk *SDK) InstanceToServiceClassAndPlan(instance *v1beta1.ServiceInstance) (*v1beta1.ClusterServiceClass, *v1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	classID := instance.Spec.ClusterServiceClassRef.Name
	classCh := make(chan *v1beta1.ClusterServiceClass)
	classErrCh := make(chan error)
	go func() {
		class, err := sdk.ServiceCatalog().ClusterServiceClasses().Get(classID, v1.GetOptions{})
		if err != nil {
			classErrCh <- err
			return
		}
		classCh <- class
	}()
	planID := instance.Spec.ClusterServicePlanRef.Name
	planCh := make(chan *v1beta1.ClusterServicePlan)
	planErrCh := make(chan error)
	go func() {
		plan, err := sdk.ServiceCatalog().ClusterServicePlans().Get(planID, v1.GetOptions{})
		if err != nil {
			planErrCh <- err
			return
		}
		planCh <- plan
	}()
	var class *v1beta1.ClusterServiceClass
	var plan *v1beta1.ClusterServicePlan
	for {
		select {
		case cl := <-classCh:
			class = cl
			if class != nil && plan != nil {
				return class, plan, nil
			}
		case err := <-classErrCh:
			return nil, nil, err
		case pl := <-planCh:
			plan = pl
			if class != nil && plan != nil {
				return class, plan, nil
			}
		case err := <-planErrCh:
			return nil, nil, err
		}
	}
}
func (sdk *SDK) Provision(instanceName, className, planName string, opts *ProvisionOptions) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	request := &v1beta1.ServiceInstance{ObjectMeta: v1.ObjectMeta{Name: instanceName, Namespace: opts.Namespace}, Spec: v1beta1.ServiceInstanceSpec{ExternalID: opts.ExternalID, PlanReference: v1beta1.PlanReference{ClusterServiceClassExternalName: className, ClusterServicePlanExternalName: planName}, Parameters: BuildParameters(opts.Params), ParametersFrom: BuildParametersFrom(opts.Secrets)}}
	result, err := sdk.ServiceCatalog().ServiceInstances(opts.Namespace).Create(request)
	if err != nil {
		return nil, fmt.Errorf("provision request failed (%s)", err)
	}
	return result, nil
}
func (sdk *SDK) Deprovision(namespace, instanceName string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := sdk.ServiceCatalog().ServiceInstances(namespace).Delete(instanceName, &v1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("deprovision request failed (%s)", err)
	}
	return nil
}
func (sdk *SDK) TouchInstance(ns, name string, retries int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for j := 0; j < retries; j++ {
		inst, err := sdk.RetrieveInstance(ns, name)
		if err != nil {
			return err
		}
		inst.Spec.UpdateRequests = inst.Spec.UpdateRequests + 1
		_, err = sdk.ServiceCatalog().ServiceInstances(ns).Update(inst)
		if err == nil {
			return nil
		}
		if !apierrors.IsConflict(err) {
			return fmt.Errorf("could not touch instance (%s)", err)
		}
	}
	return fmt.Errorf("could not sync service broker after %d tries", retries)
}
func (sdk *SDK) WaitForInstanceToNotExist(ns, name string, interval time.Duration, timeout *time.Duration) (instance *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if timeout == nil {
		notimeout := time.Duration(math.MaxInt64)
		timeout = &notimeout
	}
	err = wait.PollImmediate(interval, *timeout, func() (bool, error) {
		instance, err = sdk.ServiceCatalog().ServiceInstances(ns).Get(name, v1.GetOptions{})
		if err != nil {
			if apierrors.IsNotFound(err) {
				err = nil
			}
			return true, err
		}
		return false, err
	})
	return instance, err
}
func (sdk *SDK) WaitForInstance(ns, name string, interval time.Duration, timeout *time.Duration) (instance *v1beta1.ServiceInstance, err error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if timeout == nil {
		notimeout := time.Duration(math.MaxInt64)
		timeout = &notimeout
	}
	err = wait.PollImmediate(interval, *timeout, func() (bool, error) {
		instance, err = sdk.RetrieveInstance(ns, name)
		if nil != err {
			return false, err
		}
		if len(instance.Status.Conditions) == 0 {
			return false, nil
		}
		isDone := (sdk.IsInstanceReady(instance) || sdk.IsInstanceFailed(instance)) && !instance.Status.AsyncOpInProgress
		return isDone, nil
	})
	return instance, err
}
func (sdk *SDK) IsInstanceReady(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdk.InstanceHasStatus(instance, v1beta1.ServiceInstanceConditionReady)
}
func (sdk *SDK) IsInstanceFailed(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return sdk.InstanceHasStatus(instance, v1beta1.ServiceInstanceConditionFailed)
}
func (sdk *SDK) InstanceHasStatus(instance *v1beta1.ServiceInstance, status v1beta1.ServiceInstanceConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, cond := range instance.Status.Conditions {
		if cond.Type == status && cond.Status == v1beta1.ConditionTrue {
			return true
		}
	}
	return false
}
