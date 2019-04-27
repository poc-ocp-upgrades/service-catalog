package validation

import (
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	utilvalidation "k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
)

const commonServiceClassNameMaxLength int = 63
const guidMaxLength int = 63

func validateCommonServiceClassName(value string, prefix bool) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs []string
	if len(value) > commonServiceClassNameMaxLength {
		errs = append(errs, utilvalidation.MaxLenError(commonServiceClassNameMaxLength))
	}
	if len(value) == 0 {
		errs = append(errs, utilvalidation.EmptyError())
	}
	return errs
}
func validateExternalID(value string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var errs []string
	if len(value) > guidMaxLength {
		errs = append(errs, utilvalidation.MaxLenError(guidMaxLength))
	}
	if len(value) == 0 {
		errs = append(errs, utilvalidation.EmptyError())
	}
	return errs
}
func ValidateClusterServiceClass(serviceclass *sc.ClusterServiceClass) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return internalValidateClusterServiceClass(serviceclass)
}
func ValidateClusterServiceClassUpdate(new *sc.ClusterServiceClass, old *sc.ClusterServiceClass) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, internalValidateClusterServiceClass(new)...)
	return allErrs
}
func internalValidateClusterServiceClass(clusterserviceclass *sc.ClusterServiceClass) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMeta(&clusterserviceclass.ObjectMeta, false, validateCommonServiceClassName, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateClusterServiceClassSpec(&clusterserviceclass.Spec, field.NewPath("spec"), true)...)
	return allErrs
}
func validateClusterServiceClassSpec(spec *sc.ClusterServiceClassSpec, fldPath *field.Path, create bool) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if "" == spec.ClusterServiceBrokerName {
		allErrs = append(allErrs, field.Required(fldPath.Child("clusterServiceBrokerName"), "clusterServiceBrokerName is required"))
	}
	commonErrs := validateCommonServiceClassSpec(&spec.CommonServiceClassSpec, fldPath, create)
	if len(commonErrs) != 0 {
		allErrs = append(allErrs, commonErrs...)
	}
	return allErrs
}
func ValidateServiceClass(serviceclass *sc.ServiceClass) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return internalValidateServiceClass(serviceclass)
}
func ValidateServiceClassUpdate(new *sc.ServiceClass, old *sc.ServiceClass) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, internalValidateServiceClass(new)...)
	return allErrs
}
func internalValidateServiceClass(clusterserviceclass *sc.ServiceClass) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMeta(&clusterserviceclass.ObjectMeta, true, validateCommonServiceClassName, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateServiceClassSpec(&clusterserviceclass.Spec, field.NewPath("spec"), true)...)
	return allErrs
}
func validateServiceClassSpec(spec *sc.ServiceClassSpec, fldPath *field.Path, create bool) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	if "" == spec.ServiceBrokerName {
		allErrs = append(allErrs, field.Required(fldPath.Child("serviceBrokerName"), "serviceBrokerName is required"))
	}
	commonErrs := validateCommonServiceClassSpec(&spec.CommonServiceClassSpec, fldPath, create)
	if len(commonErrs) != 0 {
		allErrs = append(commonErrs)
	}
	return allErrs
}
func validateCommonServiceClassSpec(spec *sc.CommonServiceClassSpec, fldPath *field.Path, create bool) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	commonErrs := field.ErrorList{}
	if "" == spec.ExternalID {
		commonErrs = append(commonErrs, field.Required(fldPath.Child("externalID"), "externalID is required"))
	}
	if "" == spec.Description {
		commonErrs = append(commonErrs, field.Required(fldPath.Child("description"), "description is required"))
	}
	for _, msg := range validateCommonServiceClassName(spec.ExternalName, false) {
		commonErrs = append(commonErrs, field.Invalid(fldPath.Child("externalName"), spec.ExternalName, msg))
	}
	for _, msg := range validateExternalID(spec.ExternalID) {
		commonErrs = append(commonErrs, field.Invalid(fldPath.Child("externalID"), spec.ExternalID, msg))
	}
	return commonErrs
}
