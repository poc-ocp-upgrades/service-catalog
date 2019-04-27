package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (p *ClusterServicePlan) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Name
}
func (p *ServicePlan) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Name
}
func (p *ClusterServicePlan) GetNamespace() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ""
}
func (p *ServicePlan) GetNamespace() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Namespace
}
func (p *ClusterServicePlan) GetShortStatus() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.Status.RemovedFromBrokerCatalog {
		return "Deprecated"
	}
	return "Active"
}
func (p *ServicePlan) GetShortStatus() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if p.Status.RemovedFromBrokerCatalog {
		return "Deprecated"
	}
	return "Active"
}
func (p *ClusterServicePlan) GetExternalName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.ExternalName
}
func (p *ServicePlan) GetExternalName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.ExternalName
}
func (p *ClusterServicePlan) GetDescription() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.Description
}
func (p *ServicePlan) GetDescription() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.Description
}
func (p *ClusterServicePlan) GetFree() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.Free
}
func (p *ServicePlan) GetFree() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.Free
}
func (p *ClusterServicePlan) GetClassID() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.ClusterServiceClassRef.Name
}
func (p *ServicePlan) GetClassID() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.ServiceClassRef.Name
}
func (p *ClusterServicePlan) GetDefaultProvisionParameters() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.DefaultProvisionParameters
}
func (p *ServicePlan) GetDefaultProvisionParameters() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.DefaultProvisionParameters
}
func (p *ClusterServicePlan) GetInstanceCreateSchema() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.InstanceCreateParameterSchema
}
func (p *ServicePlan) GetInstanceCreateSchema() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.InstanceCreateParameterSchema
}
func (p *ClusterServicePlan) GetInstanceUpdateSchema() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.InstanceUpdateParameterSchema
}
func (p *ServicePlan) GetInstanceUpdateSchema() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.InstanceUpdateParameterSchema
}
func (p *ClusterServicePlan) GetBindingCreateSchema() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.ServiceBindingCreateParameterSchema
}
func (p *ServicePlan) GetBindingCreateSchema() *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return p.Spec.ServiceBindingCreateParameterSchema
}
