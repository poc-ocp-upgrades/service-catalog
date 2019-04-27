package output

import (
	"fmt"
	"io"
	"strings"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	statusActive		= "Active"
	statusDeprecated	= "Deprecated"
)
const (
	FormatJSON	= "json"
	FormatTable	= "table"
	FormatYAML	= "yaml"
)

func formatStatusShort(condition string, conditionStatus v1beta1.ConditionStatus, reason string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if conditionStatus == v1beta1.ConditionTrue {
		return condition
	}
	return reason
}
func formatStatusFull(condition string, conditionStatus v1beta1.ConditionStatus, reason string, message string, timestamp v1.Time) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status := formatStatusShort(condition, conditionStatus, reason)
	if status == "" {
		return ""
	}
	message = strings.TrimRight(message, ".")
	return fmt.Sprintf("%s - %s @ %s", status, message, timestamp.UTC())
}
func WriteDeletedResourceName(w io.Writer, resourceName string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(w, "deleted %s\n", resourceName)
}
