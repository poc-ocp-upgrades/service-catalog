package output

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
)

func getPlanStatusShort(status v1beta1.ClusterServicePlanStatus) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if status.RemovedFromBrokerCatalog {
		return statusDeprecated
	}
	return statusActive
}

type byClass []servicecatalog.Plan

func (a byClass) Len() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return len(a)
}
func (a byClass) Swap(i, j int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a[i], a[j] = a[j], a[i]
}
func (a byClass) Less(i, j int) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return a[i].GetClassID() < a[j].GetClassID()
}
func writePlanListTable(w io.Writer, plans []servicecatalog.Plan, classNames map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sort.Sort(byClass(plans))
	t := NewListTable(w)
	t.SetHeader([]string{"Name", "Namespace", "Class", "Description"})
	for _, plan := range plans {
		t.Append([]string{plan.GetExternalName(), plan.GetNamespace(), classNames[plan.GetClassID()], plan.GetDescription()})
	}
	t.SetVariableColumn(4)
	t.Render()
}
func WritePlanList(w io.Writer, outputFormat string, plans []servicecatalog.Plan, classes []servicecatalog.Class) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	classNames := map[string]string{}
	for _, class := range classes {
		classNames[class.GetName()] = class.GetExternalName()
	}
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, plans)
	case FormatYAML:
		writeYAML(w, plans, 0)
	case FormatTable:
		writePlanListTable(w, plans, classNames)
	}
}
func WritePlan(w io.Writer, outputFormat string, plan servicecatalog.Plan, class v1beta1.ClusterServiceClass) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, plan)
	case FormatYAML:
		writeYAML(w, plan, 0)
	case FormatTable:
		classNames := map[string]string{}
		classNames[class.Name] = class.Spec.ExternalName
		writePlanListTable(w, []servicecatalog.Plan{plan}, classNames)
	}
}
func WriteAssociatedPlans(w io.Writer, plans []servicecatalog.Plan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(w, "\nPlans:")
	if len(plans) == 0 {
		fmt.Fprintln(w, "No plans defined")
		return
	}
	t := NewListTable(w)
	t.SetHeader([]string{"Name", "Description"})
	for _, plan := range plans {
		t.Append([]string{plan.GetExternalName(), plan.GetDescription()})
	}
	t.Render()
}
func WriteParentPlan(w io.Writer, plan *v1beta1.ClusterServicePlan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(w, "\nPlan:")
	t := NewDetailsTable(w)
	t.AppendBulk([][]string{{"Name:", plan.Spec.ExternalName}, {"Kubernetes Name:", string(plan.Name)}, {"Status:", getPlanStatusShort(plan.Status)}})
	t.Render()
}
func WritePlanDetails(w io.Writer, plan servicecatalog.Plan, class *v1beta1.ClusterServiceClass) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewDetailsTable(w)
	t.AppendBulk([][]string{{"Name:", plan.GetExternalName()}, {"Description:", plan.GetDescription()}, {"Kubernetes Name:", string(plan.GetName())}, {"Status:", plan.GetShortStatus()}, {"Free:", strconv.FormatBool(plan.GetFree())}, {"Class:", class.Spec.ExternalName}})
	t.Render()
}
func WriteDefaultProvisionParameters(w io.Writer, plan servicecatalog.Plan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	defaultProvisionParameters := plan.GetDefaultProvisionParameters()
	if defaultProvisionParameters != nil {
		fmt.Fprintln(w, "\nDefault Provision Parameters:")
		writeYAML(w, defaultProvisionParameters, 2)
	}
}
func WritePlanSchemas(w io.Writer, plan servicecatalog.Plan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceCreateSchema := plan.GetInstanceCreateSchema()
	instanceUpdateSchema := plan.GetInstanceUpdateSchema()
	bindingCreateSchema := plan.GetBindingCreateSchema()
	if instanceCreateSchema != nil {
		fmt.Fprintln(w, "\nInstance Create Parameter Schema:")
		writeYAML(w, instanceCreateSchema, 2)
	}
	if instanceUpdateSchema != nil {
		fmt.Fprintln(w, "\nInstance Update Parameter Schema:")
		writeYAML(w, instanceUpdateSchema, 2)
	}
	if bindingCreateSchema != nil {
		fmt.Fprintln(w, "\nBinding Create Parameter Schema:")
		writeYAML(w, bindingCreateSchema, 2)
	}
}
