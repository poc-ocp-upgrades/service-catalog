package output

import (
	"io"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
)

func getBrokerStatusCondition(status v1beta1.CommonServiceBrokerStatus) v1beta1.ServiceBrokerCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(status.Conditions) > 0 {
		return status.Conditions[len(status.Conditions)-1]
	}
	return v1beta1.ServiceBrokerCondition{}
}
func getBrokerStatusShort(status v1beta1.CommonServiceBrokerStatus) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastCond := getBrokerStatusCondition(status)
	return formatStatusShort(string(lastCond.Type), lastCond.Status, lastCond.Reason)
}
func getBrokerStatusFull(status v1beta1.CommonServiceBrokerStatus) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	lastCond := getBrokerStatusCondition(status)
	return formatStatusFull(string(lastCond.Type), lastCond.Status, lastCond.Reason, lastCond.Message, lastCond.LastTransitionTime)
}
func writeBrokerListTable(w io.Writer, brokers []servicecatalog.Broker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewListTable(w)
	t.SetHeader([]string{"Name", "Namespace", "URL", "Status"})
	for _, broker := range brokers {
		t.Append([]string{broker.GetName(), broker.GetNamespace(), broker.GetURL(), getBrokerStatusShort(broker.GetStatus())})
	}
	t.Render()
}
func WriteBrokerList(w io.Writer, outputFormat string, brokers ...servicecatalog.Broker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, brokers)
	case FormatYAML:
		writeYAML(w, brokers, 0)
	case FormatTable:
		writeBrokerListTable(w, brokers)
	}
}
func WriteBroker(w io.Writer, outputFormat string, broker v1beta1.ClusterServiceBroker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch outputFormat {
	case FormatJSON:
		writeJSON(w, broker)
	case FormatYAML:
		writeYAML(w, broker, 0)
	case FormatTable:
		writeBrokerListTable(w, []servicecatalog.Broker{&broker})
	}
}
func WriteBrokerDetails(w io.Writer, broker servicecatalog.Broker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	t := NewDetailsTable(w)
	t.AppendBulk([][]string{{"Name:", broker.GetName()}, {"URL:", broker.GetURL()}, {"Status:", getBrokerStatusFull(broker.GetStatus())}})
	t.Render()
}
