package output

import (
	"fmt"
	"io"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/olekukonko/tablewriter"
)

func getInstanceStatusCondition(status v1beta1.ServiceInstanceStatus) v1beta1.ServiceInstanceCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(status.Conditions) > 0 {
		return status.Conditions[len(status.Conditions)-1]
	}
	return v1beta1.ServiceInstanceCondition{}
}
func getInstanceStatusFull(status v1beta1.ServiceInstanceStatus) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastCond := getInstanceStatusCondition(status)
	return formatStatusFull(string(lastCond.Type), lastCond.Status, lastCond.Reason, lastCond.Message, lastCond.LastTransitionTime)
}
func getInstanceStatusShort(status v1beta1.ServiceInstanceStatus) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastCond := getInstanceStatusCondition(status)
	return formatStatusShort(string(lastCond.Type), lastCond.Status, lastCond.Reason)
}
func appendInstanceDashboardURL(status v1beta1.ServiceInstanceStatus, table *tablewriter.Table) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if status.DashboardURL != nil {
		dashboardURL := *status.DashboardURL
		table.AppendBulk([][]string{{"DashboardURL:", dashboardURL}})
	}
}
func writeInstanceListTable(w io.Writer, instanceList *v1beta1.ServiceInstanceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewListTable(w)
	t.SetHeader([]string{"Name", "Namespace", "Class", "Plan", "Status"})
	for _, instance := range instanceList.Items {
		t.Append([]string{instance.Name, instance.Namespace, instance.Spec.GetSpecifiedClusterServiceClass(), instance.Spec.GetSpecifiedClusterServicePlan(), getInstanceStatusShort(instance.Status)})
	}
	t.Render()
}
func WriteInstanceList(w io.Writer, outputFormat string, instanceList *v1beta1.ServiceInstanceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, instanceList)
	case FormatYAML:
		writeYAML(w, instanceList, 0)
	case FormatTable:
		writeInstanceListTable(w, instanceList)
	}
}
func WriteInstance(w io.Writer, outputFormat string, instance v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, instance)
	case FormatYAML:
		writeYAML(w, instance, 0)
	case FormatTable:
		p := v1beta1.ServiceInstanceList{Items: []v1beta1.ServiceInstance{instance}}
		writeInstanceListTable(w, &p)
	}
}
func WriteParentInstance(w io.Writer, instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(w, "\nInstance:")
	t := NewDetailsTable(w)
	t.AppendBulk([][]string{{"Name:", instance.Name}, {"Namespace:", instance.Namespace}, {"Status:", getInstanceStatusShort(instance.Status)}})
	t.Render()
}
func WriteAssociatedInstances(w io.Writer, instances []v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(w, "\nInstances:")
	if len(instances) == 0 {
		fmt.Fprintln(w, "No instances defined")
		return
	}
	t := NewListTable(w)
	t.SetHeader([]string{"Name", "Namespace", "Status"})
	for _, instance := range instances {
		t.Append([]string{instance.Name, instance.Namespace, getInstanceStatusShort(instance.Status)})
	}
	t.Render()
}
func WriteInstanceDetails(w io.Writer, instance *v1beta1.ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewDetailsTable(w)
	t.AppendBulk([][]string{{"Name:", instance.Name}, {"Namespace:", instance.Namespace}, {"Status:", getInstanceStatusFull(instance.Status)}})
	appendInstanceDashboardURL(instance.Status, t)
	t.AppendBulk([][]string{{"Class:", instance.Spec.GetSpecifiedClusterServiceClass()}, {"Plan:", instance.Spec.GetSpecifiedClusterServicePlan()}})
	t.Render()
	writeParameters(w, instance.Spec.Parameters)
	writeParametersFrom(w, instance.Spec.ParametersFrom)
}
