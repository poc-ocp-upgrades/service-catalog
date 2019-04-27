package validation

import (
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"regexp"
)

var hexademicalStringRegexp = regexp.MustCompile("^[[:xdigit:]]*$")

func stringIsHexadecimal(s string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return hexademicalStringRegexp.MatchString(s)
}
func validateParametersFromSource(parametersFrom []sc.ParametersFromSource, fldPath *field.Path) field.ErrorList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	allErrs := field.ErrorList{}
	for _, paramsFrom := range parametersFrom {
		if paramsFrom.SecretKeyRef != nil {
			if paramsFrom.SecretKeyRef.Name == "" {
				allErrs = append(allErrs, field.Required(fldPath.Child("parametersFrom.secretKeyRef.name"), "name is required"))
			}
			if paramsFrom.SecretKeyRef.Key == "" {
				allErrs = append(allErrs, field.Required(fldPath.Child("parametersFrom.secretKeyRef.key"), "key is required"))
			}
		} else {
			allErrs = append(allErrs, field.Required(fldPath.Child("parametersFrom"), "source must not be empty if present"))
		}
	}
	return allErrs
}
