package v1beta1

const (
	statusActive		= "Active"
	statusDeprecated	= "Deprecated"
)

func (c *ClusterServiceClass) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Name
}
func (c *ServiceClass) GetName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Name
}
func (c *ClusterServiceClass) GetNamespace() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return ""
}
func (c *ServiceClass) GetNamespace() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Namespace
}
func (c *ClusterServiceClass) GetExternalName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.ExternalName
}
func (c *ServiceClass) GetExternalName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.ExternalName
}
func (c *ClusterServiceClass) GetDescription() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.Description
}
func (c *ServiceClass) GetDescription() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.Description
}
func (c *ServiceClass) GetSpec() CommonServiceClassSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.CommonServiceClassSpec
}
func (c *ClusterServiceClass) GetSpec() CommonServiceClassSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.CommonServiceClassSpec
}
func (c *ServiceClass) GetServiceBrokerName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.ServiceBrokerName
}
func (c *ClusterServiceClass) GetServiceBrokerName() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Spec.ClusterServiceBrokerName
}
func (c *ServiceClass) GetStatusText() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Status.GetStatusText()
}
func (c *ClusterServiceClass) GetStatusText() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.Status.GetStatusText()
}
func (c *CommonServiceClassStatus) GetStatusText() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c.RemovedFromBrokerCatalog {
		return statusDeprecated
	}
	return statusActive
}
