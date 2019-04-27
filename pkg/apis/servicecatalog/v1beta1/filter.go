package v1beta1

import (
	"strconv"
	"github.com/kubernetes-incubator/service-catalog/pkg/filter"
	"k8s.io/apimachinery/pkg/labels"
)

func ConvertServiceClassToProperties(serviceClass *ServiceClass) filter.Properties {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if serviceClass == nil {
		return labels.Set{}
	}
	return labels.Set{FilterName: serviceClass.Name, FilterSpecExternalName: serviceClass.Spec.ExternalName, FilterSpecExternalID: serviceClass.Spec.ExternalID}
}
func IsValidServiceClassProperty(p string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID
}
func ConvertServicePlanToProperties(servicePlan *ServicePlan) filter.Properties {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if servicePlan == nil {
		return labels.Set{}
	}
	return labels.Set{FilterName: servicePlan.Name, FilterSpecExternalName: servicePlan.Spec.ExternalName, FilterSpecExternalID: servicePlan.Spec.ExternalID, FilterSpecServiceClassName: servicePlan.Spec.ServiceClassRef.Name, FilterSpecFree: strconv.FormatBool(servicePlan.Spec.Free)}
}
func IsValidServicePlanProperty(p string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID || p == FilterSpecServiceClassName || p == FilterSpecFree
}
func ConvertClusterServiceClassToProperties(serviceClass *ClusterServiceClass) filter.Properties {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if serviceClass == nil {
		return labels.Set{}
	}
	return labels.Set{FilterName: serviceClass.Name, FilterSpecExternalName: serviceClass.Spec.ExternalName, FilterSpecExternalID: serviceClass.Spec.ExternalID}
}
func IsValidClusterServiceClassProperty(p string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID
}
func ConvertClusterServicePlanToProperties(servicePlan *ClusterServicePlan) filter.Properties {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if servicePlan == nil {
		return labels.Set{}
	}
	return labels.Set{FilterName: servicePlan.Name, FilterSpecExternalName: servicePlan.Spec.ExternalName, FilterSpecExternalID: servicePlan.Spec.ExternalID, FilterSpecClusterServiceClassName: servicePlan.Spec.ClusterServiceClassRef.Name, FilterSpecFree: strconv.FormatBool(servicePlan.Spec.Free)}
}
func IsValidClusterServicePlanProperty(p string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p == FilterName || p == FilterSpecExternalName || p == FilterSpecExternalID || p == FilterSpecClusterServiceClassName || p == FilterSpecFree
}
