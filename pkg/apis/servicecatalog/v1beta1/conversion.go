package v1beta1

import (
	"fmt"
)

func ClusterServicePlanFieldLabelConversionFunc(label, value string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "metadata.name", "spec.externalID", "spec.externalName", "spec.clusterServiceBrokerName", "spec.clusterServiceClassRef.name":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
func ServicePlanFieldLabelConversionFunc(label, value string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "metadata.name", "metadata.namespace", "spec.externalID", "spec.externalName", "spec.serviceBrokerName", "spec.serviceClassRef.name":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
func ServiceClassFieldLabelConversionFunc(label, value string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "metadata.name", "metadata.namespace", "spec.externalID", "spec.externalName", "spec.serviceBrokerName":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
func ClusterServiceClassFieldLabelConversionFunc(label, value string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "metadata.name", "spec.externalID", "spec.externalName", "spec.clusterServiceBrokerName":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
func ServiceInstanceFieldLabelConversionFunc(label, value string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "metadata.name", "metadata.namespace", "spec.externalID", "spec.clusterServiceClassRef.name", "spec.clusterServicePlanRef.name", "spec.serviceClassRef.name", "spec.servicePlanRef.name":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
func ServiceBindingFieldLabelConversionFunc(label, value string) (string, string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch label {
	case "metadata.name", "metadata.namespace", "spec.externalID":
		return label, value, nil
	default:
		return "", "", fmt.Errorf("field label not supported: %s", label)
	}
}
