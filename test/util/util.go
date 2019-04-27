package util

import (
	"context"
	"fmt"
	"testing"
	"time"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/authentication/user"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1beta1servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/servicecatalog/v1beta1"
	scfeatures "github.com/kubernetes-incubator/service-catalog/pkg/features"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
)

func WaitForBrokerCondition(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, name string, condition v1beta1.ServiceBrokerCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, 3*time.Minute, func() (bool, error) {
		klog.V(5).Infof("Waiting for broker %v condition %#v", name, condition)
		broker, err := client.ClusterServiceBrokers().Get(name, metav1.GetOptions{})
		if nil != err {
			return false, fmt.Errorf("error getting Broker %v: %v", name, err)
		}
		if len(broker.Status.Conditions) == 0 {
			return false, nil
		}
		klog.V(5).Infof("Conditions = %#v", broker.Status.Conditions)
		for _, cond := range broker.Status.Conditions {
			if condition.Type == cond.Type && condition.Status == cond.Status {
				if condition.Reason == "" || condition.Reason == cond.Reason {
					return true, nil
				}
			}
		}
		return false, nil
	})
}
func WaitForBrokerToNotExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for broker %v to not exist", name)
		_, err := client.ClusterServiceBrokers().Get(name, metav1.GetOptions{})
		if nil == err {
			return false, nil
		}
		if errors.IsNotFound(err) {
			return true, nil
		}
		return false, nil
	})
}
func WaitForClusterServiceClassToExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for serviceClass %v to exist", name)
		_, err := client.ClusterServiceClasses().Get(name, metav1.GetOptions{})
		if nil == err {
			return true, nil
		}
		return false, nil
	})
}
func WaitForClusterServicePlanToExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for ClusterServicePlan %v to exist", name)
		_, err := client.ClusterServicePlans().Get(name, metav1.GetOptions{})
		if nil == err {
			return true, nil
		}
		return false, nil
	})
}
func WaitForClusterServicePlanToNotExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for ClusterServicePlan %q to not exist", name)
		_, err := client.ClusterServicePlans().Get(name, metav1.GetOptions{})
		if nil == err {
			return false, nil
		}
		if errors.IsNotFound(err) {
			return true, nil
		}
		return false, nil
	})
}
func WaitForClusterServiceClassToNotExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for serviceClass %v to not exist", name)
		_, err := client.ClusterServiceClasses().Get(name, metav1.GetOptions{})
		if nil == err {
			return false, nil
		}
		if errors.IsNotFound(err) {
			return true, nil
		}
		return false, nil
	})
}
func WaitForInstanceCondition(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, namespace, name string, condition v1beta1.ServiceInstanceCondition) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for instance %v/%v condition %#v", namespace, name, condition)
		instance, err := client.ServiceInstances(namespace).Get(name, metav1.GetOptions{})
		if nil != err {
			return false, fmt.Errorf("error getting Instance %v/%v: %v", namespace, name, err)
		}
		if len(instance.Status.Conditions) == 0 {
			return false, nil
		}
		klog.V(5).Infof("Conditions = %#v", instance.Status.Conditions)
		for _, cond := range instance.Status.Conditions {
			if condition.Type == cond.Type && condition.Status == cond.Status {
				if condition.Reason == "" || condition.Reason == cond.Reason {
					return true, nil
				}
			}
		}
		return false, nil
	})
}
func WaitForInstanceToNotExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for instance %v/%v to not exist", namespace, name)
		_, err := client.ServiceInstances(namespace).Get(name, metav1.GetOptions{})
		if nil == err {
			return false, nil
		}
		if errors.IsNotFound(err) {
			return true, nil
		}
		return false, nil
	})
}
func WaitForInstanceProcessedGeneration(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, namespace, name string, processedGeneration int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for instance %v/%v to have processed generation of %v", namespace, name, processedGeneration)
		instance, err := client.ServiceInstances(namespace).Get(name, metav1.GetOptions{})
		if nil != err {
			return false, fmt.Errorf("error getting Instance %v/%v: %v", namespace, name, err)
		}
		if instance.Status.ObservedGeneration >= processedGeneration && (isServiceInstanceReady(instance) || isServiceInstanceFailed(instance)) && !instance.Status.OrphanMitigationInProgress {
			return true, nil
		}
		return false, nil
	})
}
func isServiceInstanceConditionTrue(instance *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isServiceInstanceConditionTrue(instance, v1beta1.ServiceInstanceConditionReady)
}
func isServiceInstanceFailed(instance *v1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return isServiceInstanceConditionTrue(instance, v1beta1.ServiceInstanceConditionFailed)
}
func WaitForBindingCondition(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, namespace, name string, condition v1beta1.ServiceBindingCondition) (*v1beta1.ServiceBindingCondition, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var lastSeenCondition *v1beta1.ServiceBindingCondition
	return lastSeenCondition, wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for binding %v/%v condition %#v", namespace, name, condition)
		binding, err := client.ServiceBindings(namespace).Get(name, metav1.GetOptions{})
		if nil != err {
			return false, fmt.Errorf("error getting Binding %v/%v: %v", namespace, name, err)
		}
		if len(binding.Status.Conditions) == 0 {
			return false, nil
		}
		klog.V(5).Infof("Conditions = %#v", binding.Status.Conditions)
		for _, cond := range binding.Status.Conditions {
			if condition.Type == cond.Type {
				lastSeenCondition = &cond
			}
			if condition.Type == cond.Type && condition.Status == cond.Status {
				if condition.Reason == "" || condition.Reason == cond.Reason {
					return true, nil
				}
			}
		}
		return false, nil
	})
}
func WaitForBindingToNotExist(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for binding %v/%v to not exist", namespace, name)
		_, err := client.ServiceBindings(namespace).Get(name, metav1.GetOptions{})
		if nil == err {
			return false, nil
		}
		if errors.IsNotFound(err) {
			return true, nil
		}
		return false, nil
	})
}
func WaitForBindingReconciledGeneration(client v1beta1servicecatalog.ServicecatalogV1beta1Interface, namespace, name string, reconciledGeneration int64) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(500*time.Millisecond, wait.ForeverTestTimeout, func() (bool, error) {
		klog.V(5).Infof("Waiting for binding %v/%v to have reconciled generation of %v", namespace, name, reconciledGeneration)
		binding, err := client.ServiceBindings(namespace).Get(name, metav1.GetOptions{})
		if nil != err {
			return false, fmt.Errorf("error getting ServiceBinding %v/%v: %v", namespace, name, err)
		}
		if binding.Status.ReconciledGeneration == reconciledGeneration {
			return true, nil
		}
		return false, nil
	})
}
func AssertServiceInstanceCondition(t *testing.T, instance *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType, status v1beta1.ConditionStatus, reason ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	foundCondition := false
	for _, condition := range instance.Status.Conditions {
		if condition.Type == conditionType {
			foundCondition = true
			if condition.Status != status {
				t.Fatalf("%v condition had unexpected status; expected %v, got %v", conditionType, status, condition.Status)
			}
			if len(reason) == 1 && condition.Reason != reason[0] {
				t.Fatalf("unexpected reason; expected %v, got %v", reason[0], condition.Reason)
			}
		}
	}
	if !foundCondition {
		t.Fatalf("%v condition not found", conditionType)
	}
}
func AssertServiceBindingCondition(t *testing.T, binding *v1beta1.ServiceBinding, conditionType v1beta1.ServiceBindingConditionType, status v1beta1.ConditionStatus, reason ...string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	foundCondition := false
	for _, condition := range binding.Status.Conditions {
		if condition.Type == conditionType {
			foundCondition = true
			if condition.Status != status {
				t.Fatalf("%v condition had unexpected status; expected %v, got %v", conditionType, status, condition.Status)
			}
			if len(reason) == 1 && condition.Reason != reason[0] {
				t.Fatalf("unexpected reason; expected %v, got %v", reason[0], condition.Reason)
			}
		}
	}
	if !foundCondition {
		t.Fatalf("%v condition not found", conditionType)
	}
}
func AssertServiceInstanceConditionFalseOrAbsent(t *testing.T, instance *v1beta1.ServiceInstance, conditionType v1beta1.ServiceInstanceConditionType) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, condition := range instance.Status.Conditions {
		if condition.Type == conditionType {
			if e, a := v1beta1.ConditionFalse, condition.Status; e != a {
				t.Fatalf("%v condition had unexpected status; expected %v, got %v", conditionType, e, a)
			}
		}
	}
}
func EnableOriginatingIdentity(t *testing.T, enabled bool) (previousState bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	prevOrigIdentEnablement := utilfeature.DefaultFeatureGate.Enabled(scfeatures.OriginatingIdentity)
	if prevOrigIdentEnablement != enabled {
		err := utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=%v", scfeatures.OriginatingIdentity, enabled))
		if err != nil {
			t.Fatalf("Failed to enable originating identity feature: %v", err)
		}
	}
	return prevOrigIdentEnablement
}
func ContextWithUserName(userName string) context.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx := genericapirequest.NewContext()
	userInfo := &user.DefaultInfo{Name: userName}
	return genericapirequest.WithUser(ctx, userInfo)
}
