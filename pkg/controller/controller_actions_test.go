package controller

import (
	"fmt"
	"reflect"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	fakeosb "github.com/pmorie/go-open-service-broker-client/v2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/testing"
)

type kubeClientAction struct {
	verb		string
	resourceName	string
	checkType	func(testing.Action) error
}

func checkGetActionType(a testing.Action) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := a.(testing.GetAction); !ok {
		return fmt.Errorf("expected a GetAction, got %s", reflect.TypeOf(a))
	}
	return nil
}
func checkUpdateActionType(a testing.Action) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if _, ok := a.(testing.UpdateAction); !ok {
		return fmt.Errorf("expected an UpdateAction, got %s", reflect.TypeOf(a))
	}
	return nil
}

type catalogClientAction struct {
	verb			string
	getRuntimeObject	func(testing.Action) (runtime.Object, error)
	checkObject		func(runtime.Object) error
}

func getRuntimeObjectFromUpdateAction(t testing.Action) (runtime.Object, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	up, ok := t.(testing.UpdateAction)
	if !ok {
		return nil, fmt.Errorf("action was not a testing.UpdateAction")
	}
	return up.GetObject(), nil
}
func checkServiceInstance(descr instanceDescription) func(runtime.Object) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func(obj runtime.Object) error {
		inst, ok := obj.(*v1beta1.ServiceInstance)
		if !ok {
			return fmt.Errorf("expected an instance, got a %s", reflect.TypeOf(obj))
		}
		if inst.Name != descr.name {
			return fmt.Errorf("expected instance name %s, got %s", descr.name, inst.Name)
		}
		if len(descr.conditionReasons) != len(inst.Status.Conditions) {
			return fmt.Errorf("expected %d conditions, got %d", len(descr.conditionReasons), len(inst.Status.Conditions))
		}
		for i, expectedConditionReason := range descr.conditionReasons {
			actualCondition := inst.Status.Conditions[i]
			if expectedConditionReason != actualCondition.Reason {
				return fmt.Errorf("condition %d: expected condition reason %s, got %s", i, expectedConditionReason, actualCondition.Reason)
			}
		}
		return nil
	}
}

type instanceDescription struct {
	name			string
	conditionReasons	[]string
}

func checkKubeClientActions(actual []testing.Action, expected []kubeClientAction) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(actual) != len(expected) {
		return fmt.Errorf("expected %d kube client actions, got %d; full action list: %v", len(expected), len(actual), actual)
	}
	for i, actualAction := range actual {
		expectedAction := expected[i]
		if actualAction.GetVerb() != expectedAction.verb {
			return fmt.Errorf("action %d: expected verb '%s', got '%s'", i, expectedAction.verb, actualAction.GetVerb())
		}
		getAction, ok := actualAction.(testing.GetAction)
		if !ok {
			return fmt.Errorf("action %d: expected a GetAction, got %s", i, reflect.TypeOf(actualAction))
		}
		if expectedAction.resourceName != getAction.GetResource().Resource {
			return fmt.Errorf("expected resource name '%s', got '%s'", expectedAction.resourceName, getAction.GetResource().Resource)
		}
	}
	return nil
}
func checkCatalogClientActions(actual []testing.Action, expected []catalogClientAction) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(actual) != len(expected) {
		return fmt.Errorf("expected %d actions, got %d", len(expected), len(actual))
	}
	for i, actualAction := range actual {
		expectedAction := expected[i]
		if actualAction.GetVerb() != expectedAction.verb {
			return fmt.Errorf("action %d: expected verb %s, got %s", i, expectedAction.verb, actualAction.GetVerb())
		}
		obj, err := expectedAction.getRuntimeObject(actualAction)
		if err != nil {
			return fmt.Errorf("action %d: %s", i, err)
		}
		if err := expectedAction.checkObject(obj); err != nil {
			return fmt.Errorf("action %d: %s", i, err)
		}
	}
	return nil
}

type brokerClientAction struct{ actionType fakeosb.ActionType }
