package output

import (
	"strings"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/olekukonko/tablewriter"
)

func Test_appendInstanceDashboardURL(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	dashboardURL := "grafana.example.com"
	table := &tablewriter.Table{}
	tests := []struct {
		name		string
		status		v1beta1.ServiceInstanceStatus
		table		*tablewriter.Table
		expectedString	string
	}{{"dashboardURLOK", v1beta1.ServiceInstanceStatus{DashboardURL: &dashboardURL}, table, "DashboardURL:   grafana.example.com"}, {"dashboardURLEmpty", v1beta1.ServiceInstanceStatus{}, table, ""}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stringBuilder strings.Builder
			tt.table = NewDetailsTable(&stringBuilder)
			appendInstanceDashboardURL(tt.status, tt.table)
			tt.table.Render()
			actualString := strings.Trim(stringBuilder.String(), " \n")
			if actualString != tt.expectedString {
				t.Fatalf("%v failed; expected %v; got %v", tt.name, tt.expectedString, actualString)
			}
		})
	}
}
