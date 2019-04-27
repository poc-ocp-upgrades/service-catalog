package pretty

type Kind int

const (
	Unknown	Kind	= iota
	ClusterServiceBroker
	ServiceBroker
	ClusterServiceClass
	ServiceClass
	ClusterServicePlan
	ServicePlan
	ServiceBinding
	ServiceInstance
)

func (k Kind) String() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	switch k {
	case ClusterServiceBroker:
		return "ClusterServiceBroker"
	case ServiceBroker:
		return "ServiceBroker"
	case ClusterServiceClass:
		return "ClusterServiceClass"
	case ServiceClass:
		return "ServiceClass"
	case ClusterServicePlan:
		return "ClusterServicePlan"
	case ServicePlan:
		return "ServicePlan"
	case ServiceBinding:
		return "ServiceBinding"
	case ServiceInstance:
		return "ServiceInstance"
	default:
		return ""
	}
}
