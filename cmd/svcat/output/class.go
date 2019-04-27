package output

import (
	"io"
	"strings"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
)

func getScope(class servicecatalog.Class) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if class.GetNamespace() != "" {
		return servicecatalog.NamespaceScope
	}
	return servicecatalog.ClusterScope
}
func writeClassListTable(w io.Writer, classes []servicecatalog.Class) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewListTable(w)
	t.SetHeader([]string{"Name", "Namespace", "Description"})
	t.SetVariableColumn(3)
	for _, class := range classes {
		t.Append([]string{class.GetExternalName(), class.GetNamespace(), class.GetDescription()})
	}
	t.Render()
}
func WriteClassList(w io.Writer, outputFormat string, classes ...servicecatalog.Class) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, classes)
	case FormatYAML:
		writeYAML(w, classes, 0)
	case FormatTable:
		writeClassListTable(w, classes)
	}
}
func WriteClass(w io.Writer, outputFormat string, class servicecatalog.Class) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, class)
	case FormatYAML:
		writeYAML(w, class, 0)
	case FormatTable:
		writeClassListTable(w, []servicecatalog.Class{class})
	}
}
func WriteClassDetails(w io.Writer, class servicecatalog.Class) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	scope := getScope(class)
	spec := class.GetSpec()
	t := NewDetailsTable(w)
	t.Append([]string{"Name:", spec.ExternalName})
	if class.GetNamespace() != "" {
		t.Append([]string{"Namespace:", class.GetNamespace()})
	}
	t.AppendBulk([][]string{{"Scope:", scope}, {"Description:", spec.Description}, {"Kubernetes Name:", class.GetName()}, {"Status:", class.GetStatusText()}, {"Tags:", strings.Join(spec.Tags, ", ")}, {"Broker:", class.GetServiceBrokerName()}})
	t.Render()
}
func WriteClassAndPlanDetails(w io.Writer, classes []servicecatalog.Class, plans [][]servicecatalog.Plan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewListTable(w)
	t.SetHeader([]string{"Class", "Plans", "Description"})
	for i, class := range classes {
		for i, plan := range plans[i] {
			if i == 0 {
				t.Append([]string{class.GetExternalName(), plan.GetExternalName(), class.GetSpec().Description})
			} else {
				t.Append([]string{"", plan.GetExternalName(), ""})
			}
		}
	}
	t.table.SetAutoWrapText(true)
	t.SetVariableColumn(3)
	t.Render()
}
