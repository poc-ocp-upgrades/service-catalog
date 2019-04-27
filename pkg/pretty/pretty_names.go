package pretty

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
)

func Name(kind Kind, k8sName, externalName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := fmt.Sprintf("%s", kind)
	if k8sName != "" && externalName != "" {
		s += fmt.Sprintf(" (K8S: %q ExternalName: %q)", k8sName, externalName)
	} else if k8sName != "" {
		s += fmt.Sprintf(" (K8S: %q)", k8sName)
	} else if externalName != "" {
		s += fmt.Sprintf(" (ExternalName: %q)", externalName)
	}
	return s
}
func ServiceInstanceName(instance *v1beta1.ServiceInstance) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf(`%s "%s/%s"`, ServiceInstance, instance.Namespace, instance.Name)
}
func ClusterServiceBrokerName(clusterServiceBrokerName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf(`%s %q`, ClusterServiceBroker, clusterServiceBrokerName)
}
func ServiceBrokerName(serviceBrokerName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf(`%s %q`, ServiceBroker, serviceBrokerName)
}
func ClusterServiceClassName(serviceClass *v1beta1.ClusterServiceClass) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if serviceClass != nil {
		return Name(ClusterServiceClass, serviceClass.Name, serviceClass.Spec.ExternalName)
	}
	return Name(ClusterServiceClass, "", "")
}
func ServiceClassName(serviceClass *v1beta1.ServiceClass) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if serviceClass != nil {
		return Name(ServiceClass, fmt.Sprintf("%s/%s", serviceClass.Namespace, serviceClass.Name), serviceClass.Spec.ExternalName)
	}
	return Name(ServiceClass, "", "")
}
func ClusterServicePlanName(servicePlan *v1beta1.ClusterServicePlan) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if servicePlan != nil {
		return Name(ClusterServicePlan, servicePlan.Name, servicePlan.Spec.ExternalName)
	}
	return Name(ClusterServicePlan, "", "")
}
func ServicePlanName(servicePlan *v1beta1.ServicePlan) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if servicePlan != nil {
		return Name(ServicePlan, fmt.Sprintf("%s/%s", servicePlan.Namespace, servicePlan.Name), servicePlan.Spec.ExternalName)
	}
	return Name(ServicePlan, "", "")
}
func FromServiceInstanceOfClusterServiceClassAtBrokerName(instance *v1beta1.ServiceInstance, serviceClass *v1beta1.ClusterServiceClass, brokerName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s of %s at %s", ServiceInstanceName(instance), ClusterServiceClassName(serviceClass), ClusterServiceBrokerName(brokerName))
}
func FromServiceInstanceOfServiceClassAtBrokerName(instance *v1beta1.ServiceInstance, serviceClass *v1beta1.ServiceClass, brokerName string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s of %s at %s", ServiceInstanceName(instance), ServiceClassName(serviceClass), ServiceBrokerName(brokerName))
}
