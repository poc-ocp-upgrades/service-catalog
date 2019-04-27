package validation

import (
	"fmt"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	utilvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

const commonServicePlanNameMaxLength int = 63

func validateCommonServicePlanName(value string, prefix bool) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs []string
	if len(value) > commonServicePlanNameMaxLength {
		errs = append(errs, utilvalidation.MaxLenError(commonServicePlanNameMaxLength))
	}
	if len(value) == 0 {
		errs = append(errs, utilvalidation.EmptyError())
	}
	return errs
}
func ValidateClusterServicePlan(clusterServicePlan *sc.ClusterServicePlan) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validateClusterServicePlan(clusterServicePlan)
}
func validateClusterServicePlan(clusterServicePlan *sc.ClusterServicePlan) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMeta(&clusterServicePlan.ObjectMeta, false, validateCommonServicePlanName, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateClusterServicePlanSpec(&clusterServicePlan.Spec, field.NewPath("spec"))...)
	return allErrs
}
func ValidateServicePlan(servicePlan *sc.ServicePlan) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return validateServicePlan(servicePlan)
}
func validateServicePlan(servicePlan *sc.ServicePlan) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMeta(&servicePlan.ObjectMeta, true, validateCommonServicePlanName, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateServicePlanSpec(&servicePlan.Spec, field.NewPath("spec"))...)
	return allErrs
}
func validateCommonServicePlanSpec(spec sc.CommonServicePlanSpec, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if "" == spec.ExternalID {
		allErrs = append(allErrs, field.Required(fldPath.Child("externalID"), "externalID is required"))
	}
	if "" == spec.Description {
		allErrs = append(allErrs, field.Required(fldPath.Child("description"), "description is required"))
	}
	for _, msg := range validateExternalID(spec.ExternalID) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("externalID"), spec.ExternalID, msg))
	}
	for _, msg := range validateCommonServicePlanName(spec.ExternalName, false) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("externalName"), spec.ExternalName, msg))
	}
	return allErrs
}
func validateClusterServicePlanSpec(spec *sc.ClusterServicePlanSpec, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := validateCommonServicePlanSpec(spec.CommonServicePlanSpec, fldPath)
	if "" == spec.ClusterServiceBrokerName {
		allErrs = append(allErrs, field.Required(fldPath.Child("clusterServiceBrokerName"), "clusterServiceBrokerName is required"))
	}
	if "" == spec.ClusterServiceClassRef.Name {
		allErrs = append(allErrs, field.Required(fldPath.Child("clusterServiceClassRef"), "an owning serviceclass is required"))
	}
	for _, msg := range validateCommonServiceClassName(spec.ClusterServiceClassRef.Name, false) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("clusterServiceClassRef", "name"), spec.ClusterServiceClassRef.Name, msg))
	}
	return allErrs
}
func validateServicePlanSpec(spec *sc.ServicePlanSpec, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := validateCommonServicePlanSpec(spec.CommonServicePlanSpec, fldPath)
	if "" == spec.ServiceBrokerName {
		allErrs = append(allErrs, field.Required(fldPath.Child("serviceBrokerName"), "serviceBrokerName is required"))
	}
	if "" == spec.ServiceClassRef.Name {
		allErrs = append(allErrs, field.Required(fldPath.Child("serviceClassRef"), "an owning serviceclass is required"))
	}
	for _, msg := range validateCommonServiceClassName(spec.ServiceClassRef.Name, false) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("serviceClassRef", "name"), spec.ServiceClassRef.Name, msg))
	}
	return allErrs
}
func ValidateClusterServicePlanUpdate(new *sc.ClusterServicePlan, old *sc.ClusterServicePlan) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validateClusterServicePlan(new)...)
	allErrs = append(allErrs, validateCommonServicePlanUpdate(new.Spec.CommonServicePlanSpec, old.Spec.CommonServicePlanSpec, "ClusterServicePlan")...)
	return allErrs
}
func validateCommonServicePlanUpdate(new sc.CommonServicePlanSpec, old sc.CommonServicePlanSpec, resourceType string) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if new.ExternalID != old.ExternalID {
		allErrs = append(allErrs, field.Invalid(field.NewPath("spec").Child("externalID"), new.ExternalID, fmt.Sprintf("externalID cannot change when updating a %s", resourceType)))
	}
	return allErrs
}
func ValidateServicePlanUpdate(new *sc.ServicePlan, old *sc.ServicePlan) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, validateServicePlan(new)...)
	allErrs = append(allErrs, validateCommonServicePlanUpdate(new.Spec.CommonServicePlanSpec, old.Spec.CommonServicePlanSpec, "ServicePlan")...)
	return allErrs
}
