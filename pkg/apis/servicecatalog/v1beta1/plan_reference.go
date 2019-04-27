package v1beta1

import (
	"fmt"
	"strings"
)

func (pr PlanReference) ClusterServiceClassSpecified() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pr.ClusterServiceClassExternalName != "" || pr.ClusterServiceClassExternalID != "" || pr.ClusterServiceClassName != ""
}
func (pr PlanReference) ClusterServicePlanSpecified() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pr.ClusterServicePlanExternalName != "" || pr.ClusterServicePlanExternalID != "" || pr.ClusterServicePlanName != ""
}
func (pr PlanReference) ServiceClassSpecified() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pr.ServiceClassExternalName != "" || pr.ServiceClassExternalID != "" || pr.ServiceClassName != ""
}
func (pr PlanReference) ServicePlanSpecified() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return pr.ServicePlanExternalName != "" || pr.ServicePlanExternalID != "" || pr.ServicePlanName != ""
}
func (pr PlanReference) GetSpecifiedClusterServiceClass() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ClusterServiceClassExternalName != "" {
		return pr.ClusterServiceClassExternalName
	}
	if pr.ClusterServiceClassExternalID != "" {
		return pr.ClusterServiceClassExternalID
	}
	if pr.ClusterServiceClassName != "" {
		return pr.ClusterServiceClassName
	}
	return ""
}
func (pr PlanReference) GetSpecifiedServiceClass() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ServiceClassExternalName != "" {
		return pr.ServiceClassExternalName
	}
	if pr.ServiceClassExternalID != "" {
		return pr.ServiceClassExternalID
	}
	if pr.ServiceClassName != "" {
		return pr.ServiceClassName
	}
	return ""
}
func (pr PlanReference) GetSpecifiedClusterServicePlan() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ClusterServicePlanExternalName != "" {
		return pr.ClusterServicePlanExternalName
	}
	if pr.ClusterServicePlanExternalID != "" {
		return pr.ClusterServicePlanExternalID
	}
	if pr.ClusterServicePlanName != "" {
		return pr.ClusterServicePlanName
	}
	return ""
}
func (pr PlanReference) GetSpecifiedServicePlan() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ServicePlanExternalName != "" {
		return pr.ServicePlanExternalName
	}
	if pr.ServicePlanExternalID != "" {
		return pr.ServicePlanExternalID
	}
	if pr.ServicePlanName != "" {
		return pr.ServicePlanName
	}
	return ""
}
func (pr PlanReference) GetClusterServiceClassFilterFieldName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ClusterServiceClassExternalName != "" {
		return "spec.externalName"
	}
	if pr.ClusterServiceClassExternalID != "" {
		return "spec.externalID"
	}
	return ""
}
func (pr PlanReference) GetClusterServicePlanFilterFieldName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ClusterServicePlanExternalName != "" {
		return "spec.externalName"
	}
	if pr.ClusterServicePlanExternalID != "" {
		return "spec.externalID"
	}
	return ""
}
func (pr PlanReference) GetServiceClassFilterFieldName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ServiceClassExternalName != "" {
		return "spec.externalName"
	}
	if pr.ServiceClassExternalID != "" {
		return "spec.externalID"
	}
	return ""
}
func (pr PlanReference) GetServicePlanFilterFieldName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pr.ServicePlanExternalName != "" {
		return "spec.externalName"
	}
	if pr.ServicePlanExternalID != "" {
		return "spec.externalID"
	}
	return ""
}
func (pr PlanReference) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var rep string
	if pr.ClusterServiceClassSpecified() && pr.ClusterServicePlanSpecified() {
		rep = fmt.Sprintf("%s/%s", pr.GetSpecifiedClusterServiceClass(), pr.GetSpecifiedClusterServicePlan())
	} else {
		rep = fmt.Sprintf("%s/%s", pr.GetSpecifiedServiceClass(), pr.GetSpecifiedServicePlan())
	}
	return rep
}
func (pr PlanReference) Format(s fmt.State, verb rune) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var classFields []string
	var planFields []string
	if pr.ClusterServiceClassExternalName != "" {
		classFields = append(classFields, fmt.Sprintf("ClusterServiceClassExternalName:%q", pr.ClusterServiceClassExternalName))
	}
	if pr.ClusterServiceClassExternalID != "" {
		classFields = append(classFields, fmt.Sprintf("ClusterServiceClassExternalID:%q", pr.ClusterServiceClassExternalID))
	}
	if pr.ClusterServiceClassName != "" {
		classFields = append(classFields, fmt.Sprintf("ClusterServiceClassName:%q", pr.ClusterServiceClassName))
	}
	if pr.ClusterServicePlanExternalName != "" {
		planFields = append(planFields, fmt.Sprintf("ClusterServicePlanExternalName:%q", pr.ClusterServicePlanExternalName))
	}
	if pr.ClusterServicePlanExternalID != "" {
		planFields = append(planFields, fmt.Sprintf("ClusterServicePlanExternalID:%q", pr.ClusterServicePlanExternalID))
	}
	if pr.ClusterServicePlanName != "" {
		planFields = append(planFields, fmt.Sprintf("ClusterServicePlanName:%q", pr.ClusterServicePlanName))
	}
	if pr.ServiceClassExternalName != "" {
		classFields = append(classFields, fmt.Sprintf("ServiceClassExternalName:%q", pr.ServiceClassExternalName))
	}
	if pr.ServiceClassExternalID != "" {
		classFields = append(classFields, fmt.Sprintf("ServiceClassExternalID:%q", pr.ServiceClassExternalID))
	}
	if pr.ServiceClassName != "" {
		classFields = append(classFields, fmt.Sprintf("ServiceClassName:%q", pr.ServiceClassName))
	}
	if pr.ServicePlanExternalName != "" {
		planFields = append(planFields, fmt.Sprintf("ServicePlanExternalName:%q", pr.ServicePlanExternalName))
	}
	if pr.ServicePlanExternalID != "" {
		planFields = append(planFields, fmt.Sprintf("ServicePlanExternalID:%q", pr.ServicePlanExternalID))
	}
	if pr.ServicePlanName != "" {
		planFields = append(planFields, fmt.Sprintf("ServicePlanName:%q", pr.ServicePlanName))
	}
	switch verb {
	case 'c':
		fmt.Fprintf(s, "{%s}", strings.Join(classFields, ", "))
	case 'b':
		fmt.Fprintf(s, "{%s}", strings.Join(planFields, ", "))
	case 'v':
		fmt.Fprintf(s, "{%s}", strings.Join(append(classFields, planFields...), ", "))
	}
}
