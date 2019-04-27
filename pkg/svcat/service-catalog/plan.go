package servicecatalog

import (
	"errors"
	"fmt"
	"strings"
	"github.com/hashicorp/go-multierror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
)

const (
	FieldExternalPlanName		= "spec.externalName"
	FieldClusterServiceClassRef	= "spec.clusterServiceClassRef.name"
	FieldServiceClassRef		= "spec.serviceClassRef.name"
)

type Plan interface {
	GetName() string
	GetShortStatus() string
	GetNamespace() string
	GetExternalName() string
	GetDescription() string
	GetFree() bool
	GetClassID() string
	GetInstanceCreateSchema() *runtime.RawExtension
	GetInstanceUpdateSchema() *runtime.RawExtension
	GetBindingCreateSchema() *runtime.RawExtension
	GetDefaultProvisionParameters() *runtime.RawExtension
}

func (sdk *SDK) RetrievePlans(classID string, opts ScopeOptions) ([]Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	plans, err := sdk.retrievePlansByListOptions(opts, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	if classID == "" {
		return plans, nil
	}
	var filtered []Plan
	for _, p := range plans {
		if p.GetClassID() == classID {
			filtered = append(filtered, p)
		}
	}
	return filtered, nil
}
func (sdk *SDK) retrievePlansByListOptions(scopeOpts ScopeOptions, listOpts metav1.ListOptions) ([]Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var plans []Plan
	if scopeOpts.Scope.Matches(ClusterScope) {
		csp, err := sdk.ServiceCatalog().ClusterServicePlans().List(listOpts)
		if err != nil {
			return nil, fmt.Errorf("unable to list cluster-scoped plans (%s)", err)
		}
		for _, p := range csp.Items {
			plan := p
			plans = append(plans, &plan)
		}
	}
	if scopeOpts.Scope.Matches(NamespaceScope) {
		sp, err := sdk.ServiceCatalog().ServicePlans(scopeOpts.Namespace).List(listOpts)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return plans, nil
			}
			return nil, fmt.Errorf("unable to list plans in %q (%s)", scopeOpts.Namespace, err)
		}
		for _, p := range sp.Items {
			plan := p
			plans = append(plans, &plan)
		}
	}
	return plans, nil
}
func (sdk *SDK) RetrievePlanByName(name string, opts ScopeOptions) (Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}
	listOpts := metav1.ListOptions{FieldSelector: fields.OneTermEqualSelector(FieldExternalPlanName, name).String()}
	return sdk.retrieveSinglePlanByListOptions(name, opts, listOpts)
}
func (sdk *SDK) RetrievePlanByClassAndName(className, planName string, opts ScopeOptions) (Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}
	class, err := sdk.RetrieveClassByName(className, opts)
	if err != nil {
		return nil, err
	}
	var classRefSelector fields.Selector
	if opts.Scope.Matches(ClusterScope) {
		classRefSelector = fields.OneTermEqualSelector(FieldClusterServiceClassRef, class.GetName())
	} else {
		classRefSelector = fields.OneTermEqualSelector(FieldServiceClassRef, class.GetName())
	}
	listOpts := metav1.ListOptions{FieldSelector: fields.AndSelectors(classRefSelector, fields.OneTermEqualSelector(FieldExternalPlanName, planName)).String()}
	ss := []string{class.GetName(), planName}
	return sdk.retrieveSinglePlanByListOptions(strings.Join(ss, "/"), opts, listOpts)
}
func (sdk *SDK) RetrievePlanByClassIDAndName(classKubeName, planName string, scopeOpts ScopeOptions) (Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var classRefSelector fields.Selector
	findError := &multierror.Error{ErrorFormat: func(errors []error) string {
		return joinErrors("error:", errors, "\n  ")
	}}
	if scopeOpts.Scope.Matches(ClusterScope) {
		classRefSelector = fields.OneTermEqualSelector(FieldClusterServiceClassRef, classKubeName)
		listOpts := metav1.ListOptions{FieldSelector: fields.AndSelectors(classRefSelector, fields.OneTermEqualSelector(FieldExternalPlanName, planName)).String()}
		ss := []string{classKubeName, planName}
		plan, err := sdk.retrieveSinglePlanByListOptions(strings.Join(ss, "/"), scopeOpts, listOpts)
		if err != nil {
			findError = multierror.Append(findError, err)
		} else if plan != nil {
			return plan, nil
		}
	}
	if scopeOpts.Scope.Matches(NamespaceScope) {
		classRefSelector = fields.OneTermEqualSelector(FieldServiceClassRef, classKubeName)
		listOpts := metav1.ListOptions{FieldSelector: fields.AndSelectors(classRefSelector, fields.OneTermEqualSelector(FieldExternalPlanName, planName)).String()}
		ss := []string{classKubeName, planName}
		plan, err := sdk.retrieveSinglePlanByListOptions(strings.Join(ss, "/"), scopeOpts, listOpts)
		if err != nil {
			findError = multierror.Append(findError, err)
		} else if plan != nil {
			return plan, nil
		}
	}
	return nil, fmt.Errorf("plan '%s' not found:%s", planName, findError.Error())
}
func (sdk *SDK) retrieveSinglePlanByListOptions(name string, scopeOpts ScopeOptions, listOpts metav1.ListOptions) (Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	plans, err := sdk.retrievePlansByListOptions(scopeOpts, listOpts)
	if err != nil {
		return nil, err
	}
	if len(plans) == 0 {
		return nil, fmt.Errorf("plan not found '%s'", name)
	}
	if len(plans) > 1 {
		return nil, fmt.Errorf("more than one matching plan found for '%s'", name)
	}
	return plans[0], nil
}
func (sdk *SDK) RetrievePlanByID(kubeName string, opts ScopeOptions) (Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if opts.Scope == AllScope {
		return nil, errors.New("invalid scope: all")
	}
	if opts.Scope.Matches(ClusterScope) {
		p, err := sdk.ServiceCatalog().ClusterServicePlans().Get(kubeName, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get cluster-scoped plan by Kubernetes name'%s' (%s)", kubeName, err)
		}
		return p, nil
	}
	if opts.Scope.Matches(NamespaceScope) {
		p, err := sdk.ServiceCatalog().ServicePlans(opts.Namespace).Get(kubeName, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("unable to get plan by Kubernetes name'%s' (%s)", kubeName, err)
		}
		return p, nil
	}
	return nil, fmt.Errorf("unable to get plan by Kubernetes name'%s'", kubeName)
}
