package servicecatalogfakes

import (
	"sync"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"time"
	apiv1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	apicorev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/version"
)

type FakeSvcatClient struct {
	BindStub	func(string, string, string, string, string, interface{}, map[string]string) (*apiv1beta1.ServiceBinding, error)
	bindMutex	sync.RWMutex
	bindArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	string
		arg4	string
		arg5	string
		arg6	interface{}
		arg7	map[string]string
	}
	bindReturns	struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}
	bindReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}
	BindingParentHierarchyStub		func(*apiv1beta1.ServiceBinding) (*apiv1beta1.ServiceInstance, *apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, *apiv1beta1.ClusterServiceBroker, error)
	bindingParentHierarchyMutex		sync.RWMutex
	bindingParentHierarchyArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceBinding }
	bindingParentHierarchyReturns		struct {
		result1	*apiv1beta1.ServiceInstance
		result2	*apiv1beta1.ClusterServiceClass
		result3	*apiv1beta1.ClusterServicePlan
		result4	*apiv1beta1.ClusterServiceBroker
		result5	error
	}
	bindingParentHierarchyReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstance
		result2	*apiv1beta1.ClusterServiceClass
		result3	*apiv1beta1.ClusterServicePlan
		result4	*apiv1beta1.ClusterServiceBroker
		result5	error
	}
	DeleteBindingStub		func(string, string) error
	deleteBindingMutex		sync.RWMutex
	deleteBindingArgsForCall	[]struct {
		arg1	string
		arg2	string
	}
	deleteBindingReturns		struct{ result1 error }
	deleteBindingReturnsOnCall	map[int]struct{ result1 error }
	DeleteBindingsStub		func([]types.NamespacedName) ([]types.NamespacedName, error)
	deleteBindingsMutex		sync.RWMutex
	deleteBindingsArgsForCall	[]struct{ arg1 []types.NamespacedName }
	deleteBindingsReturns		struct {
		result1	[]types.NamespacedName
		result2	error
	}
	deleteBindingsReturnsOnCall	map[int]struct {
		result1	[]types.NamespacedName
		result2	error
	}
	IsBindingFailedStub		func(*apiv1beta1.ServiceBinding) bool
	isBindingFailedMutex		sync.RWMutex
	isBindingFailedArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceBinding }
	isBindingFailedReturns		struct{ result1 bool }
	isBindingFailedReturnsOnCall	map[int]struct{ result1 bool }
	IsBindingReadyStub		func(*apiv1beta1.ServiceBinding) bool
	isBindingReadyMutex		sync.RWMutex
	isBindingReadyArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceBinding }
	isBindingReadyReturns		struct{ result1 bool }
	isBindingReadyReturnsOnCall	map[int]struct{ result1 bool }
	RetrieveBindingStub		func(string, string) (*apiv1beta1.ServiceBinding, error)
	retrieveBindingMutex		sync.RWMutex
	retrieveBindingArgsForCall	[]struct {
		arg1	string
		arg2	string
	}
	retrieveBindingReturns	struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}
	retrieveBindingReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}
	RetrieveBindingsStub		func(string) (*apiv1beta1.ServiceBindingList, error)
	retrieveBindingsMutex		sync.RWMutex
	retrieveBindingsArgsForCall	[]struct{ arg1 string }
	retrieveBindingsReturns		struct {
		result1	*apiv1beta1.ServiceBindingList
		result2	error
	}
	retrieveBindingsReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceBindingList
		result2	error
	}
	RetrieveBindingsByInstanceStub		func(*apiv1beta1.ServiceInstance) ([]apiv1beta1.ServiceBinding, error)
	retrieveBindingsByInstanceMutex		sync.RWMutex
	retrieveBindingsByInstanceArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceInstance }
	retrieveBindingsByInstanceReturns	struct {
		result1	[]apiv1beta1.ServiceBinding
		result2	error
	}
	retrieveBindingsByInstanceReturnsOnCall	map[int]struct {
		result1	[]apiv1beta1.ServiceBinding
		result2	error
	}
	UnbindStub		func(string, string) ([]types.NamespacedName, error)
	unbindMutex		sync.RWMutex
	unbindArgsForCall	[]struct {
		arg1	string
		arg2	string
	}
	unbindReturns	struct {
		result1	[]types.NamespacedName
		result2	error
	}
	unbindReturnsOnCall	map[int]struct {
		result1	[]types.NamespacedName
		result2	error
	}
	WaitForBindingStub		func(string, string, time.Duration, *time.Duration) (*apiv1beta1.ServiceBinding, error)
	waitForBindingMutex		sync.RWMutex
	waitForBindingArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	time.Duration
		arg4	*time.Duration
	}
	waitForBindingReturns	struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}
	waitForBindingReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}
	DeregisterStub		func(string, *servicecatalog.ScopeOptions) error
	deregisterMutex		sync.RWMutex
	deregisterArgsForCall	[]struct {
		arg1	string
		arg2	*servicecatalog.ScopeOptions
	}
	deregisterReturns		struct{ result1 error }
	deregisterReturnsOnCall		map[int]struct{ result1 error }
	RetrieveBrokersStub		func(opts servicecatalog.ScopeOptions) ([]servicecatalog.Broker, error)
	retrieveBrokersMutex		sync.RWMutex
	retrieveBrokersArgsForCall	[]struct{ opts servicecatalog.ScopeOptions }
	retrieveBrokersReturns		struct {
		result1	[]servicecatalog.Broker
		result2	error
	}
	retrieveBrokersReturnsOnCall	map[int]struct {
		result1	[]servicecatalog.Broker
		result2	error
	}
	RetrieveBrokerStub		func(string) (*apiv1beta1.ClusterServiceBroker, error)
	retrieveBrokerMutex		sync.RWMutex
	retrieveBrokerArgsForCall	[]struct{ arg1 string }
	retrieveBrokerReturns		struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}
	retrieveBrokerReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}
	RetrieveBrokerByClassStub		func(*apiv1beta1.ClusterServiceClass) (*apiv1beta1.ClusterServiceBroker, error)
	retrieveBrokerByClassMutex		sync.RWMutex
	retrieveBrokerByClassArgsForCall	[]struct {
		arg1 *apiv1beta1.ClusterServiceClass
	}
	retrieveBrokerByClassReturns	struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}
	retrieveBrokerByClassReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}
	RegisterStub		func(string, string, *servicecatalog.RegisterOptions, *servicecatalog.ScopeOptions) (servicecatalog.Broker, error)
	registerMutex		sync.RWMutex
	registerArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	*servicecatalog.RegisterOptions
		arg4	*servicecatalog.ScopeOptions
	}
	registerReturns	struct {
		result1	servicecatalog.Broker
		result2	error
	}
	registerReturnsOnCall	map[int]struct {
		result1	servicecatalog.Broker
		result2	error
	}
	SyncStub	func(string, servicecatalog.ScopeOptions, int) error
	syncMutex	sync.RWMutex
	syncArgsForCall	[]struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
		arg3	int
	}
	syncReturns			struct{ result1 error }
	syncReturnsOnCall		map[int]struct{ result1 error }
	WaitForBrokerStub		func(string, time.Duration, *time.Duration) (servicecatalog.Broker, error)
	waitForBrokerMutex		sync.RWMutex
	waitForBrokerArgsForCall	[]struct {
		arg1	string
		arg2	time.Duration
		arg3	*time.Duration
	}
	waitForBrokerReturns	struct {
		result1	servicecatalog.Broker
		result2	error
	}
	waitForBrokerReturnsOnCall	map[int]struct {
		result1	servicecatalog.Broker
		result2	error
	}
	RetrieveClassesStub		func(servicecatalog.ScopeOptions) ([]servicecatalog.Class, error)
	retrieveClassesMutex		sync.RWMutex
	retrieveClassesArgsForCall	[]struct{ arg1 servicecatalog.ScopeOptions }
	retrieveClassesReturns		struct {
		result1	[]servicecatalog.Class
		result2	error
	}
	retrieveClassesReturnsOnCall	map[int]struct {
		result1	[]servicecatalog.Class
		result2	error
	}
	RetrieveClassByNameStub		func(string, servicecatalog.ScopeOptions) (servicecatalog.Class, error)
	retrieveClassByNameMutex	sync.RWMutex
	retrieveClassByNameArgsForCall	[]struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}
	retrieveClassByNameReturns	struct {
		result1	servicecatalog.Class
		result2	error
	}
	retrieveClassByNameReturnsOnCall	map[int]struct {
		result1	servicecatalog.Class
		result2	error
	}
	RetrieveClassByIDStub		func(string) (*apiv1beta1.ClusterServiceClass, error)
	retrieveClassByIDMutex		sync.RWMutex
	retrieveClassByIDArgsForCall	[]struct{ arg1 string }
	retrieveClassByIDReturns	struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}
	retrieveClassByIDReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}
	RetrieveClassByPlanStub		func(servicecatalog.Plan) (*apiv1beta1.ClusterServiceClass, error)
	retrieveClassByPlanMutex	sync.RWMutex
	retrieveClassByPlanArgsForCall	[]struct{ arg1 servicecatalog.Plan }
	retrieveClassByPlanReturns	struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}
	retrieveClassByPlanReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}
	CreateClassFromStub		func(servicecatalog.CreateClassFromOptions) (servicecatalog.Class, error)
	createClassFromMutex		sync.RWMutex
	createClassFromArgsForCall	[]struct {
		arg1 servicecatalog.CreateClassFromOptions
	}
	createClassFromReturns	struct {
		result1	servicecatalog.Class
		result2	error
	}
	createClassFromReturnsOnCall	map[int]struct {
		result1	servicecatalog.Class
		result2	error
	}
	DeprovisionStub		func(string, string) error
	deprovisionMutex	sync.RWMutex
	deprovisionArgsForCall	[]struct {
		arg1	string
		arg2	string
	}
	deprovisionReturns			struct{ result1 error }
	deprovisionReturnsOnCall		map[int]struct{ result1 error }
	InstanceParentHierarchyStub		func(*apiv1beta1.ServiceInstance) (*apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, *apiv1beta1.ClusterServiceBroker, error)
	instanceParentHierarchyMutex		sync.RWMutex
	instanceParentHierarchyArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceInstance }
	instanceParentHierarchyReturns		struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	*apiv1beta1.ClusterServiceBroker
		result4	error
	}
	instanceParentHierarchyReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	*apiv1beta1.ClusterServiceBroker
		result4	error
	}
	InstanceToServiceClassAndPlanStub		func(*apiv1beta1.ServiceInstance) (*apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, error)
	instanceToServiceClassAndPlanMutex		sync.RWMutex
	instanceToServiceClassAndPlanArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceInstance }
	instanceToServiceClassAndPlanReturns		struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	error
	}
	instanceToServiceClassAndPlanReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	error
	}
	IsInstanceFailedStub		func(*apiv1beta1.ServiceInstance) bool
	isInstanceFailedMutex		sync.RWMutex
	isInstanceFailedArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceInstance }
	isInstanceFailedReturns		struct{ result1 bool }
	isInstanceFailedReturnsOnCall	map[int]struct{ result1 bool }
	IsInstanceReadyStub		func(*apiv1beta1.ServiceInstance) bool
	isInstanceReadyMutex		sync.RWMutex
	isInstanceReadyArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceInstance }
	isInstanceReadyReturns		struct{ result1 bool }
	isInstanceReadyReturnsOnCall	map[int]struct{ result1 bool }
	ProvisionStub			func(string, string, string, *servicecatalog.ProvisionOptions) (*apiv1beta1.ServiceInstance, error)
	provisionMutex			sync.RWMutex
	provisionArgsForCall		[]struct {
		arg1	string
		arg2	string
		arg3	string
		arg4	*servicecatalog.ProvisionOptions
	}
	provisionReturns	struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	provisionReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	RetrieveInstanceStub		func(string, string) (*apiv1beta1.ServiceInstance, error)
	retrieveInstanceMutex		sync.RWMutex
	retrieveInstanceArgsForCall	[]struct {
		arg1	string
		arg2	string
	}
	retrieveInstanceReturns	struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	retrieveInstanceReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	RetrieveInstanceByBindingStub		func(*apiv1beta1.ServiceBinding) (*apiv1beta1.ServiceInstance, error)
	retrieveInstanceByBindingMutex		sync.RWMutex
	retrieveInstanceByBindingArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceBinding }
	retrieveInstanceByBindingReturns	struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	retrieveInstanceByBindingReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	RetrieveInstancesStub		func(string, string, string) (*apiv1beta1.ServiceInstanceList, error)
	retrieveInstancesMutex		sync.RWMutex
	retrieveInstancesArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	string
	}
	retrieveInstancesReturns	struct {
		result1	*apiv1beta1.ServiceInstanceList
		result2	error
	}
	retrieveInstancesReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstanceList
		result2	error
	}
	RetrieveInstancesByPlanStub		func(servicecatalog.Plan) ([]apiv1beta1.ServiceInstance, error)
	retrieveInstancesByPlanMutex		sync.RWMutex
	retrieveInstancesByPlanArgsForCall	[]struct{ arg1 servicecatalog.Plan }
	retrieveInstancesByPlanReturns		struct {
		result1	[]apiv1beta1.ServiceInstance
		result2	error
	}
	retrieveInstancesByPlanReturnsOnCall	map[int]struct {
		result1	[]apiv1beta1.ServiceInstance
		result2	error
	}
	TouchInstanceStub		func(string, string, int) error
	touchInstanceMutex		sync.RWMutex
	touchInstanceArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	int
	}
	touchInstanceReturns		struct{ result1 error }
	touchInstanceReturnsOnCall	map[int]struct{ result1 error }
	WaitForInstanceStub		func(string, string, time.Duration, *time.Duration) (*apiv1beta1.ServiceInstance, error)
	waitForInstanceMutex		sync.RWMutex
	waitForInstanceArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	time.Duration
		arg4	*time.Duration
	}
	waitForInstanceReturns	struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	waitForInstanceReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	WaitForInstanceToNotExistStub		func(string, string, time.Duration, *time.Duration) (*apiv1beta1.ServiceInstance, error)
	waitForInstanceToNotExistMutex		sync.RWMutex
	waitForInstanceToNotExistArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	time.Duration
		arg4	*time.Duration
	}
	waitForInstanceToNotExistReturns	struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	waitForInstanceToNotExistReturnsOnCall	map[int]struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}
	RetrievePlansStub		func(string, servicecatalog.ScopeOptions) ([]servicecatalog.Plan, error)
	retrievePlansMutex		sync.RWMutex
	retrievePlansArgsForCall	[]struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}
	retrievePlansReturns	struct {
		result1	[]servicecatalog.Plan
		result2	error
	}
	retrievePlansReturnsOnCall	map[int]struct {
		result1	[]servicecatalog.Plan
		result2	error
	}
	RetrievePlanByNameStub		func(string, servicecatalog.ScopeOptions) (servicecatalog.Plan, error)
	retrievePlanByNameMutex		sync.RWMutex
	retrievePlanByNameArgsForCall	[]struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}
	retrievePlanByNameReturns	struct {
		result1	servicecatalog.Plan
		result2	error
	}
	retrievePlanByNameReturnsOnCall	map[int]struct {
		result1	servicecatalog.Plan
		result2	error
	}
	RetrievePlanByClassAndNameStub		func(string, string, servicecatalog.ScopeOptions) (servicecatalog.Plan, error)
	retrievePlanByClassAndNameMutex		sync.RWMutex
	retrievePlanByClassAndNameArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	servicecatalog.ScopeOptions
	}
	retrievePlanByClassAndNameReturns	struct {
		result1	servicecatalog.Plan
		result2	error
	}
	retrievePlanByClassAndNameReturnsOnCall	map[int]struct {
		result1	servicecatalog.Plan
		result2	error
	}
	RetrievePlanByClassIDAndNameStub	func(string, string, servicecatalog.ScopeOptions) (servicecatalog.Plan, error)
	retrievePlanByClassIDAndNameMutex	sync.RWMutex
	retrievePlanByClassIDAndNameArgsForCall	[]struct {
		arg1	string
		arg2	string
		arg3	servicecatalog.ScopeOptions
	}
	retrievePlanByClassIDAndNameReturns	struct {
		result1	servicecatalog.Plan
		result2	error
	}
	retrievePlanByClassIDAndNameReturnsOnCall	map[int]struct {
		result1	servicecatalog.Plan
		result2	error
	}
	RetrievePlanByIDStub		func(string, servicecatalog.ScopeOptions) (servicecatalog.Plan, error)
	retrievePlanByIDMutex		sync.RWMutex
	retrievePlanByIDArgsForCall	[]struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}
	retrievePlanByIDReturns	struct {
		result1	servicecatalog.Plan
		result2	error
	}
	retrievePlanByIDReturnsOnCall	map[int]struct {
		result1	servicecatalog.Plan
		result2	error
	}
	RetrieveSecretByBindingStub		func(*apiv1beta1.ServiceBinding) (*apicorev1.Secret, error)
	retrieveSecretByBindingMutex		sync.RWMutex
	retrieveSecretByBindingArgsForCall	[]struct{ arg1 *apiv1beta1.ServiceBinding }
	retrieveSecretByBindingReturns		struct {
		result1	*apicorev1.Secret
		result2	error
	}
	retrieveSecretByBindingReturnsOnCall	map[int]struct {
		result1	*apicorev1.Secret
		result2	error
	}
	ServerVersionStub		func() (*version.Info, error)
	serverVersionMutex		sync.RWMutex
	serverVersionArgsForCall	[]struct{}
	serverVersionReturns		struct {
		result1	*version.Info
		result2	error
	}
	serverVersionReturnsOnCall	map[int]struct {
		result1	*version.Info
		result2	error
	}
	invocations		map[string][][]interface{}
	invocationsMutex	sync.RWMutex
}

func (fake *FakeSvcatClient) Bind(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 interface{}, arg7 map[string]string) (*apiv1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	string
		arg4	string
		arg5	string
		arg6	interface{}
		arg7	map[string]string
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("Bind", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindReturns.result1, fake.bindReturns.result2
}
func (fake *FakeSvcatClient) BindCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}
func (fake *FakeSvcatClient) BindArgsForCall(i int) (string, string, string, string, string, interface{}, map[string]string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].arg1, fake.bindArgsForCall[i].arg2, fake.bindArgsForCall[i].arg3, fake.bindArgsForCall[i].arg4, fake.bindArgsForCall[i].arg5, fake.bindArgsForCall[i].arg6, fake.bindArgsForCall[i].arg7
}
func (fake *FakeSvcatClient) BindReturns(result1 *apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) BindReturnsOnCall(i int, result1 *apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceBinding
			result2	error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) BindingParentHierarchy(arg1 *apiv1beta1.ServiceBinding) (*apiv1beta1.ServiceInstance, *apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, *apiv1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.bindingParentHierarchyMutex.Lock()
	ret, specificReturn := fake.bindingParentHierarchyReturnsOnCall[len(fake.bindingParentHierarchyArgsForCall)]
	fake.bindingParentHierarchyArgsForCall = append(fake.bindingParentHierarchyArgsForCall, struct{ arg1 *apiv1beta1.ServiceBinding }{arg1})
	fake.recordInvocation("BindingParentHierarchy", []interface{}{arg1})
	fake.bindingParentHierarchyMutex.Unlock()
	if fake.BindingParentHierarchyStub != nil {
		return fake.BindingParentHierarchyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4, ret.result5
	}
	return fake.bindingParentHierarchyReturns.result1, fake.bindingParentHierarchyReturns.result2, fake.bindingParentHierarchyReturns.result3, fake.bindingParentHierarchyReturns.result4, fake.bindingParentHierarchyReturns.result5
}
func (fake *FakeSvcatClient) BindingParentHierarchyCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.bindingParentHierarchyMutex.RLock()
	defer fake.bindingParentHierarchyMutex.RUnlock()
	return len(fake.bindingParentHierarchyArgsForCall)
}
func (fake *FakeSvcatClient) BindingParentHierarchyArgsForCall(i int) *apiv1beta1.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.bindingParentHierarchyMutex.RLock()
	defer fake.bindingParentHierarchyMutex.RUnlock()
	return fake.bindingParentHierarchyArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) BindingParentHierarchyReturns(result1 *apiv1beta1.ServiceInstance, result2 *apiv1beta1.ClusterServiceClass, result3 *apiv1beta1.ClusterServicePlan, result4 *apiv1beta1.ClusterServiceBroker, result5 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.BindingParentHierarchyStub = nil
	fake.bindingParentHierarchyReturns = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	*apiv1beta1.ClusterServiceClass
		result3	*apiv1beta1.ClusterServicePlan
		result4	*apiv1beta1.ClusterServiceBroker
		result5	error
	}{result1, result2, result3, result4, result5}
}
func (fake *FakeSvcatClient) BindingParentHierarchyReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstance, result2 *apiv1beta1.ClusterServiceClass, result3 *apiv1beta1.ClusterServicePlan, result4 *apiv1beta1.ClusterServiceBroker, result5 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.BindingParentHierarchyStub = nil
	if fake.bindingParentHierarchyReturnsOnCall == nil {
		fake.bindingParentHierarchyReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstance
			result2	*apiv1beta1.ClusterServiceClass
			result3	*apiv1beta1.ClusterServicePlan
			result4	*apiv1beta1.ClusterServiceBroker
			result5	error
		})
	}
	fake.bindingParentHierarchyReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	*apiv1beta1.ClusterServiceClass
		result3	*apiv1beta1.ClusterServicePlan
		result4	*apiv1beta1.ClusterServiceBroker
		result5	error
	}{result1, result2, result3, result4, result5}
}
func (fake *FakeSvcatClient) DeleteBinding(arg1 string, arg2 string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deleteBindingMutex.Lock()
	ret, specificReturn := fake.deleteBindingReturnsOnCall[len(fake.deleteBindingArgsForCall)]
	fake.deleteBindingArgsForCall = append(fake.deleteBindingArgsForCall, struct {
		arg1	string
		arg2	string
	}{arg1, arg2})
	fake.recordInvocation("DeleteBinding", []interface{}{arg1, arg2})
	fake.deleteBindingMutex.Unlock()
	if fake.DeleteBindingStub != nil {
		return fake.DeleteBindingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteBindingReturns.result1
}
func (fake *FakeSvcatClient) DeleteBindingCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	return len(fake.deleteBindingArgsForCall)
}
func (fake *FakeSvcatClient) DeleteBindingArgsForCall(i int) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	return fake.deleteBindingArgsForCall[i].arg1, fake.deleteBindingArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) DeleteBindingReturns(result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeleteBindingStub = nil
	fake.deleteBindingReturns = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) DeleteBindingReturnsOnCall(i int, result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeleteBindingStub = nil
	if fake.deleteBindingReturnsOnCall == nil {
		fake.deleteBindingReturnsOnCall = make(map[int]struct{ result1 error })
	}
	fake.deleteBindingReturnsOnCall[i] = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) DeleteBindings(arg1 []types.NamespacedName) ([]types.NamespacedName, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var arg1Copy []types.NamespacedName
	if arg1 != nil {
		arg1Copy = make([]types.NamespacedName, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteBindingsMutex.Lock()
	ret, specificReturn := fake.deleteBindingsReturnsOnCall[len(fake.deleteBindingsArgsForCall)]
	fake.deleteBindingsArgsForCall = append(fake.deleteBindingsArgsForCall, struct{ arg1 []types.NamespacedName }{arg1Copy})
	fake.recordInvocation("DeleteBindings", []interface{}{arg1Copy})
	fake.deleteBindingsMutex.Unlock()
	if fake.DeleteBindingsStub != nil {
		return fake.DeleteBindingsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteBindingsReturns.result1, fake.deleteBindingsReturns.result2
}
func (fake *FakeSvcatClient) DeleteBindingsCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deleteBindingsMutex.RLock()
	defer fake.deleteBindingsMutex.RUnlock()
	return len(fake.deleteBindingsArgsForCall)
}
func (fake *FakeSvcatClient) DeleteBindingsArgsForCall(i int) []types.NamespacedName {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deleteBindingsMutex.RLock()
	defer fake.deleteBindingsMutex.RUnlock()
	return fake.deleteBindingsArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) DeleteBindingsReturns(result1 []types.NamespacedName, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeleteBindingsStub = nil
	fake.deleteBindingsReturns = struct {
		result1	[]types.NamespacedName
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) DeleteBindingsReturnsOnCall(i int, result1 []types.NamespacedName, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeleteBindingsStub = nil
	if fake.deleteBindingsReturnsOnCall == nil {
		fake.deleteBindingsReturnsOnCall = make(map[int]struct {
			result1	[]types.NamespacedName
			result2	error
		})
	}
	fake.deleteBindingsReturnsOnCall[i] = struct {
		result1	[]types.NamespacedName
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) IsBindingFailed(arg1 *apiv1beta1.ServiceBinding) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isBindingFailedMutex.Lock()
	ret, specificReturn := fake.isBindingFailedReturnsOnCall[len(fake.isBindingFailedArgsForCall)]
	fake.isBindingFailedArgsForCall = append(fake.isBindingFailedArgsForCall, struct{ arg1 *apiv1beta1.ServiceBinding }{arg1})
	fake.recordInvocation("IsBindingFailed", []interface{}{arg1})
	fake.isBindingFailedMutex.Unlock()
	if fake.IsBindingFailedStub != nil {
		return fake.IsBindingFailedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isBindingFailedReturns.result1
}
func (fake *FakeSvcatClient) IsBindingFailedCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isBindingFailedMutex.RLock()
	defer fake.isBindingFailedMutex.RUnlock()
	return len(fake.isBindingFailedArgsForCall)
}
func (fake *FakeSvcatClient) IsBindingFailedArgsForCall(i int) *apiv1beta1.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isBindingFailedMutex.RLock()
	defer fake.isBindingFailedMutex.RUnlock()
	return fake.isBindingFailedArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) IsBindingFailedReturns(result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsBindingFailedStub = nil
	fake.isBindingFailedReturns = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) IsBindingFailedReturnsOnCall(i int, result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsBindingFailedStub = nil
	if fake.isBindingFailedReturnsOnCall == nil {
		fake.isBindingFailedReturnsOnCall = make(map[int]struct{ result1 bool })
	}
	fake.isBindingFailedReturnsOnCall[i] = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) IsBindingReady(arg1 *apiv1beta1.ServiceBinding) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isBindingReadyMutex.Lock()
	ret, specificReturn := fake.isBindingReadyReturnsOnCall[len(fake.isBindingReadyArgsForCall)]
	fake.isBindingReadyArgsForCall = append(fake.isBindingReadyArgsForCall, struct{ arg1 *apiv1beta1.ServiceBinding }{arg1})
	fake.recordInvocation("IsBindingReady", []interface{}{arg1})
	fake.isBindingReadyMutex.Unlock()
	if fake.IsBindingReadyStub != nil {
		return fake.IsBindingReadyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isBindingReadyReturns.result1
}
func (fake *FakeSvcatClient) IsBindingReadyCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isBindingReadyMutex.RLock()
	defer fake.isBindingReadyMutex.RUnlock()
	return len(fake.isBindingReadyArgsForCall)
}
func (fake *FakeSvcatClient) IsBindingReadyArgsForCall(i int) *apiv1beta1.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isBindingReadyMutex.RLock()
	defer fake.isBindingReadyMutex.RUnlock()
	return fake.isBindingReadyArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) IsBindingReadyReturns(result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsBindingReadyStub = nil
	fake.isBindingReadyReturns = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) IsBindingReadyReturnsOnCall(i int, result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsBindingReadyStub = nil
	if fake.isBindingReadyReturnsOnCall == nil {
		fake.isBindingReadyReturnsOnCall = make(map[int]struct{ result1 bool })
	}
	fake.isBindingReadyReturnsOnCall[i] = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) RetrieveBinding(arg1 string, arg2 string) (*apiv1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingMutex.Lock()
	ret, specificReturn := fake.retrieveBindingReturnsOnCall[len(fake.retrieveBindingArgsForCall)]
	fake.retrieveBindingArgsForCall = append(fake.retrieveBindingArgsForCall, struct {
		arg1	string
		arg2	string
	}{arg1, arg2})
	fake.recordInvocation("RetrieveBinding", []interface{}{arg1, arg2})
	fake.retrieveBindingMutex.Unlock()
	if fake.RetrieveBindingStub != nil {
		return fake.RetrieveBindingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveBindingReturns.result1, fake.retrieveBindingReturns.result2
}
func (fake *FakeSvcatClient) RetrieveBindingCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingMutex.RLock()
	defer fake.retrieveBindingMutex.RUnlock()
	return len(fake.retrieveBindingArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveBindingArgsForCall(i int) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingMutex.RLock()
	defer fake.retrieveBindingMutex.RUnlock()
	return fake.retrieveBindingArgsForCall[i].arg1, fake.retrieveBindingArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) RetrieveBindingReturns(result1 *apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBindingStub = nil
	fake.retrieveBindingReturns = struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBindingReturnsOnCall(i int, result1 *apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBindingStub = nil
	if fake.retrieveBindingReturnsOnCall == nil {
		fake.retrieveBindingReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceBinding
			result2	error
		})
	}
	fake.retrieveBindingReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBindings(arg1 string) (*apiv1beta1.ServiceBindingList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingsMutex.Lock()
	ret, specificReturn := fake.retrieveBindingsReturnsOnCall[len(fake.retrieveBindingsArgsForCall)]
	fake.retrieveBindingsArgsForCall = append(fake.retrieveBindingsArgsForCall, struct{ arg1 string }{arg1})
	fake.recordInvocation("RetrieveBindings", []interface{}{arg1})
	fake.retrieveBindingsMutex.Unlock()
	if fake.RetrieveBindingsStub != nil {
		return fake.RetrieveBindingsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveBindingsReturns.result1, fake.retrieveBindingsReturns.result2
}
func (fake *FakeSvcatClient) RetrieveBindingsCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingsMutex.RLock()
	defer fake.retrieveBindingsMutex.RUnlock()
	return len(fake.retrieveBindingsArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveBindingsArgsForCall(i int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingsMutex.RLock()
	defer fake.retrieveBindingsMutex.RUnlock()
	return fake.retrieveBindingsArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveBindingsReturns(result1 *apiv1beta1.ServiceBindingList, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBindingsStub = nil
	fake.retrieveBindingsReturns = struct {
		result1	*apiv1beta1.ServiceBindingList
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBindingsReturnsOnCall(i int, result1 *apiv1beta1.ServiceBindingList, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBindingsStub = nil
	if fake.retrieveBindingsReturnsOnCall == nil {
		fake.retrieveBindingsReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceBindingList
			result2	error
		})
	}
	fake.retrieveBindingsReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceBindingList
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBindingsByInstance(arg1 *apiv1beta1.ServiceInstance) ([]apiv1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingsByInstanceMutex.Lock()
	ret, specificReturn := fake.retrieveBindingsByInstanceReturnsOnCall[len(fake.retrieveBindingsByInstanceArgsForCall)]
	fake.retrieveBindingsByInstanceArgsForCall = append(fake.retrieveBindingsByInstanceArgsForCall, struct{ arg1 *apiv1beta1.ServiceInstance }{arg1})
	fake.recordInvocation("RetrieveBindingsByInstance", []interface{}{arg1})
	fake.retrieveBindingsByInstanceMutex.Unlock()
	if fake.RetrieveBindingsByInstanceStub != nil {
		return fake.RetrieveBindingsByInstanceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveBindingsByInstanceReturns.result1, fake.retrieveBindingsByInstanceReturns.result2
}
func (fake *FakeSvcatClient) RetrieveBindingsByInstanceCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingsByInstanceMutex.RLock()
	defer fake.retrieveBindingsByInstanceMutex.RUnlock()
	return len(fake.retrieveBindingsByInstanceArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveBindingsByInstanceArgsForCall(i int) *apiv1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBindingsByInstanceMutex.RLock()
	defer fake.retrieveBindingsByInstanceMutex.RUnlock()
	return fake.retrieveBindingsByInstanceArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveBindingsByInstanceReturns(result1 []apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBindingsByInstanceStub = nil
	fake.retrieveBindingsByInstanceReturns = struct {
		result1	[]apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBindingsByInstanceReturnsOnCall(i int, result1 []apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBindingsByInstanceStub = nil
	if fake.retrieveBindingsByInstanceReturnsOnCall == nil {
		fake.retrieveBindingsByInstanceReturnsOnCall = make(map[int]struct {
			result1	[]apiv1beta1.ServiceBinding
			result2	error
		})
	}
	fake.retrieveBindingsByInstanceReturnsOnCall[i] = struct {
		result1	[]apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) Unbind(arg1 string, arg2 string) ([]types.NamespacedName, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		arg1	string
		arg2	string
	}{arg1, arg2})
	fake.recordInvocation("Unbind", []interface{}{arg1, arg2})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unbindReturns.result1, fake.unbindReturns.result2
}
func (fake *FakeSvcatClient) UnbindCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}
func (fake *FakeSvcatClient) UnbindArgsForCall(i int) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].arg1, fake.unbindArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) UnbindReturns(result1 []types.NamespacedName, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1	[]types.NamespacedName
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) UnbindReturnsOnCall(i int, result1 []types.NamespacedName, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1	[]types.NamespacedName
			result2	error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1	[]types.NamespacedName
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) WaitForBinding(arg1 string, arg2 string, arg3 time.Duration, arg4 *time.Duration) (*apiv1beta1.ServiceBinding, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForBindingMutex.Lock()
	ret, specificReturn := fake.waitForBindingReturnsOnCall[len(fake.waitForBindingArgsForCall)]
	fake.waitForBindingArgsForCall = append(fake.waitForBindingArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	time.Duration
		arg4	*time.Duration
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("WaitForBinding", []interface{}{arg1, arg2, arg3, arg4})
	fake.waitForBindingMutex.Unlock()
	if fake.WaitForBindingStub != nil {
		return fake.WaitForBindingStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.waitForBindingReturns.result1, fake.waitForBindingReturns.result2
}
func (fake *FakeSvcatClient) WaitForBindingCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForBindingMutex.RLock()
	defer fake.waitForBindingMutex.RUnlock()
	return len(fake.waitForBindingArgsForCall)
}
func (fake *FakeSvcatClient) WaitForBindingArgsForCall(i int) (string, string, time.Duration, *time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForBindingMutex.RLock()
	defer fake.waitForBindingMutex.RUnlock()
	return fake.waitForBindingArgsForCall[i].arg1, fake.waitForBindingArgsForCall[i].arg2, fake.waitForBindingArgsForCall[i].arg3, fake.waitForBindingArgsForCall[i].arg4
}
func (fake *FakeSvcatClient) WaitForBindingReturns(result1 *apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForBindingStub = nil
	fake.waitForBindingReturns = struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) WaitForBindingReturnsOnCall(i int, result1 *apiv1beta1.ServiceBinding, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForBindingStub = nil
	if fake.waitForBindingReturnsOnCall == nil {
		fake.waitForBindingReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceBinding
			result2	error
		})
	}
	fake.waitForBindingReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceBinding
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) Deregister(arg1 string, arg2 *servicecatalog.ScopeOptions) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deregisterMutex.Lock()
	ret, specificReturn := fake.deregisterReturnsOnCall[len(fake.deregisterArgsForCall)]
	fake.deregisterArgsForCall = append(fake.deregisterArgsForCall, struct {
		arg1	string
		arg2	*servicecatalog.ScopeOptions
	}{arg1, arg2})
	fake.recordInvocation("Deregister", []interface{}{arg1, arg2})
	fake.deregisterMutex.Unlock()
	if fake.DeregisterStub != nil {
		return fake.DeregisterStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deregisterReturns.result1
}
func (fake *FakeSvcatClient) DeregisterCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	return len(fake.deregisterArgsForCall)
}
func (fake *FakeSvcatClient) DeregisterArgsForCall(i int) (string, *servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	return fake.deregisterArgsForCall[i].arg1, fake.deregisterArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) DeregisterReturns(result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeregisterStub = nil
	fake.deregisterReturns = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) DeregisterReturnsOnCall(i int, result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeregisterStub = nil
	if fake.deregisterReturnsOnCall == nil {
		fake.deregisterReturnsOnCall = make(map[int]struct{ result1 error })
	}
	fake.deregisterReturnsOnCall[i] = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) RetrieveBrokers(opts servicecatalog.ScopeOptions) ([]servicecatalog.Broker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokersMutex.Lock()
	ret, specificReturn := fake.retrieveBrokersReturnsOnCall[len(fake.retrieveBrokersArgsForCall)]
	fake.retrieveBrokersArgsForCall = append(fake.retrieveBrokersArgsForCall, struct{ opts servicecatalog.ScopeOptions }{opts})
	fake.recordInvocation("RetrieveBrokers", []interface{}{opts})
	fake.retrieveBrokersMutex.Unlock()
	if fake.RetrieveBrokersStub != nil {
		return fake.RetrieveBrokersStub(opts)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveBrokersReturns.result1, fake.retrieveBrokersReturns.result2
}
func (fake *FakeSvcatClient) RetrieveBrokersCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokersMutex.RLock()
	defer fake.retrieveBrokersMutex.RUnlock()
	return len(fake.retrieveBrokersArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveBrokersArgsForCall(i int) servicecatalog.ScopeOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokersMutex.RLock()
	defer fake.retrieveBrokersMutex.RUnlock()
	return fake.retrieveBrokersArgsForCall[i].opts
}
func (fake *FakeSvcatClient) RetrieveBrokersReturns(result1 []servicecatalog.Broker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBrokersStub = nil
	fake.retrieveBrokersReturns = struct {
		result1	[]servicecatalog.Broker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBrokersReturnsOnCall(i int, result1 []servicecatalog.Broker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBrokersStub = nil
	if fake.retrieveBrokersReturnsOnCall == nil {
		fake.retrieveBrokersReturnsOnCall = make(map[int]struct {
			result1	[]servicecatalog.Broker
			result2	error
		})
	}
	fake.retrieveBrokersReturnsOnCall[i] = struct {
		result1	[]servicecatalog.Broker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBroker(arg1 string) (*apiv1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokerMutex.Lock()
	ret, specificReturn := fake.retrieveBrokerReturnsOnCall[len(fake.retrieveBrokerArgsForCall)]
	fake.retrieveBrokerArgsForCall = append(fake.retrieveBrokerArgsForCall, struct{ arg1 string }{arg1})
	fake.recordInvocation("RetrieveBroker", []interface{}{arg1})
	fake.retrieveBrokerMutex.Unlock()
	if fake.RetrieveBrokerStub != nil {
		return fake.RetrieveBrokerStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveBrokerReturns.result1, fake.retrieveBrokerReturns.result2
}
func (fake *FakeSvcatClient) RetrieveBrokerCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokerMutex.RLock()
	defer fake.retrieveBrokerMutex.RUnlock()
	return len(fake.retrieveBrokerArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveBrokerArgsForCall(i int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokerMutex.RLock()
	defer fake.retrieveBrokerMutex.RUnlock()
	return fake.retrieveBrokerArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveBrokerReturns(result1 *apiv1beta1.ClusterServiceBroker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBrokerStub = nil
	fake.retrieveBrokerReturns = struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBrokerReturnsOnCall(i int, result1 *apiv1beta1.ClusterServiceBroker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBrokerStub = nil
	if fake.retrieveBrokerReturnsOnCall == nil {
		fake.retrieveBrokerReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ClusterServiceBroker
			result2	error
		})
	}
	fake.retrieveBrokerReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBrokerByClass(arg1 *apiv1beta1.ClusterServiceClass) (*apiv1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokerByClassMutex.Lock()
	ret, specificReturn := fake.retrieveBrokerByClassReturnsOnCall[len(fake.retrieveBrokerByClassArgsForCall)]
	fake.retrieveBrokerByClassArgsForCall = append(fake.retrieveBrokerByClassArgsForCall, struct {
		arg1 *apiv1beta1.ClusterServiceClass
	}{arg1})
	fake.recordInvocation("RetrieveBrokerByClass", []interface{}{arg1})
	fake.retrieveBrokerByClassMutex.Unlock()
	if fake.RetrieveBrokerByClassStub != nil {
		return fake.RetrieveBrokerByClassStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveBrokerByClassReturns.result1, fake.retrieveBrokerByClassReturns.result2
}
func (fake *FakeSvcatClient) RetrieveBrokerByClassCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokerByClassMutex.RLock()
	defer fake.retrieveBrokerByClassMutex.RUnlock()
	return len(fake.retrieveBrokerByClassArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveBrokerByClassArgsForCall(i int) *apiv1beta1.ClusterServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveBrokerByClassMutex.RLock()
	defer fake.retrieveBrokerByClassMutex.RUnlock()
	return fake.retrieveBrokerByClassArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveBrokerByClassReturns(result1 *apiv1beta1.ClusterServiceBroker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBrokerByClassStub = nil
	fake.retrieveBrokerByClassReturns = struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveBrokerByClassReturnsOnCall(i int, result1 *apiv1beta1.ClusterServiceBroker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveBrokerByClassStub = nil
	if fake.retrieveBrokerByClassReturnsOnCall == nil {
		fake.retrieveBrokerByClassReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ClusterServiceBroker
			result2	error
		})
	}
	fake.retrieveBrokerByClassReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ClusterServiceBroker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) Register(arg1 string, arg2 string, arg3 *servicecatalog.RegisterOptions, arg4 *servicecatalog.ScopeOptions) (servicecatalog.Broker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.registerMutex.Lock()
	ret, specificReturn := fake.registerReturnsOnCall[len(fake.registerArgsForCall)]
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	*servicecatalog.RegisterOptions
		arg4	*servicecatalog.ScopeOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Register", []interface{}{arg1, arg2, arg3, arg4})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.registerReturns.result1, fake.registerReturns.result2
}
func (fake *FakeSvcatClient) RegisterCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}
func (fake *FakeSvcatClient) RegisterArgsForCall(i int) (string, string, *servicecatalog.RegisterOptions, *servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].arg1, fake.registerArgsForCall[i].arg2, fake.registerArgsForCall[i].arg3, fake.registerArgsForCall[i].arg4
}
func (fake *FakeSvcatClient) RegisterReturns(result1 servicecatalog.Broker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1	servicecatalog.Broker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RegisterReturnsOnCall(i int, result1 servicecatalog.Broker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RegisterStub = nil
	if fake.registerReturnsOnCall == nil {
		fake.registerReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Broker
			result2	error
		})
	}
	fake.registerReturnsOnCall[i] = struct {
		result1	servicecatalog.Broker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) Sync(arg1 string, arg2 servicecatalog.ScopeOptions, arg3 int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.syncMutex.Lock()
	ret, specificReturn := fake.syncReturnsOnCall[len(fake.syncArgsForCall)]
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
		arg3	int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Sync", []interface{}{arg1, arg2, arg3})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		return fake.SyncStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.syncReturns.result1
}
func (fake *FakeSvcatClient) SyncCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}
func (fake *FakeSvcatClient) SyncArgsForCall(i int) (string, servicecatalog.ScopeOptions, int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return fake.syncArgsForCall[i].arg1, fake.syncArgsForCall[i].arg2, fake.syncArgsForCall[i].arg3
}
func (fake *FakeSvcatClient) SyncReturns(result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.SyncStub = nil
	fake.syncReturns = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) SyncReturnsOnCall(i int, result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.SyncStub = nil
	if fake.syncReturnsOnCall == nil {
		fake.syncReturnsOnCall = make(map[int]struct{ result1 error })
	}
	fake.syncReturnsOnCall[i] = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) WaitForBroker(arg1 string, arg2 time.Duration, arg3 *time.Duration) (servicecatalog.Broker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForBrokerMutex.Lock()
	ret, specificReturn := fake.waitForBrokerReturnsOnCall[len(fake.waitForBrokerArgsForCall)]
	fake.waitForBrokerArgsForCall = append(fake.waitForBrokerArgsForCall, struct {
		arg1	string
		arg2	time.Duration
		arg3	*time.Duration
	}{arg1, arg2, arg3})
	fake.recordInvocation("WaitForBroker", []interface{}{arg1, arg2, arg3})
	fake.waitForBrokerMutex.Unlock()
	if fake.WaitForBrokerStub != nil {
		return fake.WaitForBrokerStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.waitForBrokerReturns.result1, fake.waitForBrokerReturns.result2
}
func (fake *FakeSvcatClient) WaitForBrokerCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForBrokerMutex.RLock()
	defer fake.waitForBrokerMutex.RUnlock()
	return len(fake.waitForBrokerArgsForCall)
}
func (fake *FakeSvcatClient) WaitForBrokerArgsForCall(i int) (string, time.Duration, *time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForBrokerMutex.RLock()
	defer fake.waitForBrokerMutex.RUnlock()
	return fake.waitForBrokerArgsForCall[i].arg1, fake.waitForBrokerArgsForCall[i].arg2, fake.waitForBrokerArgsForCall[i].arg3
}
func (fake *FakeSvcatClient) WaitForBrokerReturns(result1 servicecatalog.Broker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForBrokerStub = nil
	fake.waitForBrokerReturns = struct {
		result1	servicecatalog.Broker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) WaitForBrokerReturnsOnCall(i int, result1 servicecatalog.Broker, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForBrokerStub = nil
	if fake.waitForBrokerReturnsOnCall == nil {
		fake.waitForBrokerReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Broker
			result2	error
		})
	}
	fake.waitForBrokerReturnsOnCall[i] = struct {
		result1	servicecatalog.Broker
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClasses(arg1 servicecatalog.ScopeOptions) ([]servicecatalog.Class, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassesMutex.Lock()
	ret, specificReturn := fake.retrieveClassesReturnsOnCall[len(fake.retrieveClassesArgsForCall)]
	fake.retrieveClassesArgsForCall = append(fake.retrieveClassesArgsForCall, struct{ arg1 servicecatalog.ScopeOptions }{arg1})
	fake.recordInvocation("RetrieveClasses", []interface{}{arg1})
	fake.retrieveClassesMutex.Unlock()
	if fake.RetrieveClassesStub != nil {
		return fake.RetrieveClassesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveClassesReturns.result1, fake.retrieveClassesReturns.result2
}
func (fake *FakeSvcatClient) RetrieveClassesCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassesMutex.RLock()
	defer fake.retrieveClassesMutex.RUnlock()
	return len(fake.retrieveClassesArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveClassesArgsForCall(i int) servicecatalog.ScopeOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassesMutex.RLock()
	defer fake.retrieveClassesMutex.RUnlock()
	return fake.retrieveClassesArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveClassesReturns(result1 []servicecatalog.Class, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassesStub = nil
	fake.retrieveClassesReturns = struct {
		result1	[]servicecatalog.Class
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassesReturnsOnCall(i int, result1 []servicecatalog.Class, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassesStub = nil
	if fake.retrieveClassesReturnsOnCall == nil {
		fake.retrieveClassesReturnsOnCall = make(map[int]struct {
			result1	[]servicecatalog.Class
			result2	error
		})
	}
	fake.retrieveClassesReturnsOnCall[i] = struct {
		result1	[]servicecatalog.Class
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassByName(arg1 string, arg2 servicecatalog.ScopeOptions) (servicecatalog.Class, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByNameMutex.Lock()
	ret, specificReturn := fake.retrieveClassByNameReturnsOnCall[len(fake.retrieveClassByNameArgsForCall)]
	fake.retrieveClassByNameArgsForCall = append(fake.retrieveClassByNameArgsForCall, struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}{arg1, arg2})
	fake.recordInvocation("RetrieveClassByName", []interface{}{arg1, arg2})
	fake.retrieveClassByNameMutex.Unlock()
	if fake.RetrieveClassByNameStub != nil {
		return fake.RetrieveClassByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveClassByNameReturns.result1, fake.retrieveClassByNameReturns.result2
}
func (fake *FakeSvcatClient) RetrieveClassByNameCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByNameMutex.RLock()
	defer fake.retrieveClassByNameMutex.RUnlock()
	return len(fake.retrieveClassByNameArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveClassByNameArgsForCall(i int) (string, servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByNameMutex.RLock()
	defer fake.retrieveClassByNameMutex.RUnlock()
	return fake.retrieveClassByNameArgsForCall[i].arg1, fake.retrieveClassByNameArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) RetrieveClassByNameReturns(result1 servicecatalog.Class, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassByNameStub = nil
	fake.retrieveClassByNameReturns = struct {
		result1	servicecatalog.Class
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassByNameReturnsOnCall(i int, result1 servicecatalog.Class, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassByNameStub = nil
	if fake.retrieveClassByNameReturnsOnCall == nil {
		fake.retrieveClassByNameReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Class
			result2	error
		})
	}
	fake.retrieveClassByNameReturnsOnCall[i] = struct {
		result1	servicecatalog.Class
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassByID(arg1 string) (*apiv1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByIDMutex.Lock()
	ret, specificReturn := fake.retrieveClassByIDReturnsOnCall[len(fake.retrieveClassByIDArgsForCall)]
	fake.retrieveClassByIDArgsForCall = append(fake.retrieveClassByIDArgsForCall, struct{ arg1 string }{arg1})
	fake.recordInvocation("RetrieveClassByID", []interface{}{arg1})
	fake.retrieveClassByIDMutex.Unlock()
	if fake.RetrieveClassByIDStub != nil {
		return fake.RetrieveClassByIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveClassByIDReturns.result1, fake.retrieveClassByIDReturns.result2
}
func (fake *FakeSvcatClient) RetrieveClassByIDCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByIDMutex.RLock()
	defer fake.retrieveClassByIDMutex.RUnlock()
	return len(fake.retrieveClassByIDArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveClassByIDArgsForCall(i int) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByIDMutex.RLock()
	defer fake.retrieveClassByIDMutex.RUnlock()
	return fake.retrieveClassByIDArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveClassByIDReturns(result1 *apiv1beta1.ClusterServiceClass, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassByIDStub = nil
	fake.retrieveClassByIDReturns = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassByIDReturnsOnCall(i int, result1 *apiv1beta1.ClusterServiceClass, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassByIDStub = nil
	if fake.retrieveClassByIDReturnsOnCall == nil {
		fake.retrieveClassByIDReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ClusterServiceClass
			result2	error
		})
	}
	fake.retrieveClassByIDReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassByPlan(arg1 servicecatalog.Plan) (*apiv1beta1.ClusterServiceClass, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByPlanMutex.Lock()
	ret, specificReturn := fake.retrieveClassByPlanReturnsOnCall[len(fake.retrieveClassByPlanArgsForCall)]
	fake.retrieveClassByPlanArgsForCall = append(fake.retrieveClassByPlanArgsForCall, struct{ arg1 servicecatalog.Plan }{arg1})
	fake.recordInvocation("RetrieveClassByPlan", []interface{}{arg1})
	fake.retrieveClassByPlanMutex.Unlock()
	if fake.RetrieveClassByPlanStub != nil {
		return fake.RetrieveClassByPlanStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveClassByPlanReturns.result1, fake.retrieveClassByPlanReturns.result2
}
func (fake *FakeSvcatClient) RetrieveClassByPlanCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByPlanMutex.RLock()
	defer fake.retrieveClassByPlanMutex.RUnlock()
	return len(fake.retrieveClassByPlanArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveClassByPlanArgsForCall(i int) servicecatalog.Plan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveClassByPlanMutex.RLock()
	defer fake.retrieveClassByPlanMutex.RUnlock()
	return fake.retrieveClassByPlanArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveClassByPlanReturns(result1 *apiv1beta1.ClusterServiceClass, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassByPlanStub = nil
	fake.retrieveClassByPlanReturns = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveClassByPlanReturnsOnCall(i int, result1 *apiv1beta1.ClusterServiceClass, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveClassByPlanStub = nil
	if fake.retrieveClassByPlanReturnsOnCall == nil {
		fake.retrieveClassByPlanReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ClusterServiceClass
			result2	error
		})
	}
	fake.retrieveClassByPlanReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) CreateClassFrom(arg1 servicecatalog.CreateClassFromOptions) (servicecatalog.Class, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.createClassFromMutex.Lock()
	ret, specificReturn := fake.createClassFromReturnsOnCall[len(fake.createClassFromArgsForCall)]
	fake.createClassFromArgsForCall = append(fake.createClassFromArgsForCall, struct {
		arg1 servicecatalog.CreateClassFromOptions
	}{arg1})
	fake.recordInvocation("CreateClassFrom", []interface{}{arg1})
	fake.createClassFromMutex.Unlock()
	if fake.CreateClassFromStub != nil {
		return fake.CreateClassFromStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createClassFromReturns.result1, fake.createClassFromReturns.result2
}
func (fake *FakeSvcatClient) CreateClassFromCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.createClassFromMutex.RLock()
	defer fake.createClassFromMutex.RUnlock()
	return len(fake.createClassFromArgsForCall)
}
func (fake *FakeSvcatClient) CreateClassFromArgsForCall(i int) servicecatalog.CreateClassFromOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.createClassFromMutex.RLock()
	defer fake.createClassFromMutex.RUnlock()
	return fake.createClassFromArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) CreateClassFromReturns(result1 servicecatalog.Class, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.CreateClassFromStub = nil
	fake.createClassFromReturns = struct {
		result1	servicecatalog.Class
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) CreateClassFromReturnsOnCall(i int, result1 servicecatalog.Class, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.CreateClassFromStub = nil
	if fake.createClassFromReturnsOnCall == nil {
		fake.createClassFromReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Class
			result2	error
		})
	}
	fake.createClassFromReturnsOnCall[i] = struct {
		result1	servicecatalog.Class
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) Deprovision(arg1 string, arg2 string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		arg1	string
		arg2	string
	}{arg1, arg2})
	fake.recordInvocation("Deprovision", []interface{}{arg1, arg2})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deprovisionReturns.result1
}
func (fake *FakeSvcatClient) DeprovisionCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}
func (fake *FakeSvcatClient) DeprovisionArgsForCall(i int) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].arg1, fake.deprovisionArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) DeprovisionReturns(result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) DeprovisionReturnsOnCall(i int, result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct{ result1 error })
	}
	fake.deprovisionReturnsOnCall[i] = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) InstanceParentHierarchy(arg1 *apiv1beta1.ServiceInstance) (*apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, *apiv1beta1.ClusterServiceBroker, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.instanceParentHierarchyMutex.Lock()
	ret, specificReturn := fake.instanceParentHierarchyReturnsOnCall[len(fake.instanceParentHierarchyArgsForCall)]
	fake.instanceParentHierarchyArgsForCall = append(fake.instanceParentHierarchyArgsForCall, struct{ arg1 *apiv1beta1.ServiceInstance }{arg1})
	fake.recordInvocation("InstanceParentHierarchy", []interface{}{arg1})
	fake.instanceParentHierarchyMutex.Unlock()
	if fake.InstanceParentHierarchyStub != nil {
		return fake.InstanceParentHierarchyStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.instanceParentHierarchyReturns.result1, fake.instanceParentHierarchyReturns.result2, fake.instanceParentHierarchyReturns.result3, fake.instanceParentHierarchyReturns.result4
}
func (fake *FakeSvcatClient) InstanceParentHierarchyCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.instanceParentHierarchyMutex.RLock()
	defer fake.instanceParentHierarchyMutex.RUnlock()
	return len(fake.instanceParentHierarchyArgsForCall)
}
func (fake *FakeSvcatClient) InstanceParentHierarchyArgsForCall(i int) *apiv1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.instanceParentHierarchyMutex.RLock()
	defer fake.instanceParentHierarchyMutex.RUnlock()
	return fake.instanceParentHierarchyArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) InstanceParentHierarchyReturns(result1 *apiv1beta1.ClusterServiceClass, result2 *apiv1beta1.ClusterServicePlan, result3 *apiv1beta1.ClusterServiceBroker, result4 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.InstanceParentHierarchyStub = nil
	fake.instanceParentHierarchyReturns = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	*apiv1beta1.ClusterServiceBroker
		result4	error
	}{result1, result2, result3, result4}
}
func (fake *FakeSvcatClient) InstanceParentHierarchyReturnsOnCall(i int, result1 *apiv1beta1.ClusterServiceClass, result2 *apiv1beta1.ClusterServicePlan, result3 *apiv1beta1.ClusterServiceBroker, result4 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.InstanceParentHierarchyStub = nil
	if fake.instanceParentHierarchyReturnsOnCall == nil {
		fake.instanceParentHierarchyReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ClusterServiceClass
			result2	*apiv1beta1.ClusterServicePlan
			result3	*apiv1beta1.ClusterServiceBroker
			result4	error
		})
	}
	fake.instanceParentHierarchyReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	*apiv1beta1.ClusterServiceBroker
		result4	error
	}{result1, result2, result3, result4}
}
func (fake *FakeSvcatClient) InstanceToServiceClassAndPlan(arg1 *apiv1beta1.ServiceInstance) (*apiv1beta1.ClusterServiceClass, *apiv1beta1.ClusterServicePlan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.instanceToServiceClassAndPlanMutex.Lock()
	ret, specificReturn := fake.instanceToServiceClassAndPlanReturnsOnCall[len(fake.instanceToServiceClassAndPlanArgsForCall)]
	fake.instanceToServiceClassAndPlanArgsForCall = append(fake.instanceToServiceClassAndPlanArgsForCall, struct{ arg1 *apiv1beta1.ServiceInstance }{arg1})
	fake.recordInvocation("InstanceToServiceClassAndPlan", []interface{}{arg1})
	fake.instanceToServiceClassAndPlanMutex.Unlock()
	if fake.InstanceToServiceClassAndPlanStub != nil {
		return fake.InstanceToServiceClassAndPlanStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.instanceToServiceClassAndPlanReturns.result1, fake.instanceToServiceClassAndPlanReturns.result2, fake.instanceToServiceClassAndPlanReturns.result3
}
func (fake *FakeSvcatClient) InstanceToServiceClassAndPlanCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.instanceToServiceClassAndPlanMutex.RLock()
	defer fake.instanceToServiceClassAndPlanMutex.RUnlock()
	return len(fake.instanceToServiceClassAndPlanArgsForCall)
}
func (fake *FakeSvcatClient) InstanceToServiceClassAndPlanArgsForCall(i int) *apiv1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.instanceToServiceClassAndPlanMutex.RLock()
	defer fake.instanceToServiceClassAndPlanMutex.RUnlock()
	return fake.instanceToServiceClassAndPlanArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) InstanceToServiceClassAndPlanReturns(result1 *apiv1beta1.ClusterServiceClass, result2 *apiv1beta1.ClusterServicePlan, result3 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.InstanceToServiceClassAndPlanStub = nil
	fake.instanceToServiceClassAndPlanReturns = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	error
	}{result1, result2, result3}
}
func (fake *FakeSvcatClient) InstanceToServiceClassAndPlanReturnsOnCall(i int, result1 *apiv1beta1.ClusterServiceClass, result2 *apiv1beta1.ClusterServicePlan, result3 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.InstanceToServiceClassAndPlanStub = nil
	if fake.instanceToServiceClassAndPlanReturnsOnCall == nil {
		fake.instanceToServiceClassAndPlanReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ClusterServiceClass
			result2	*apiv1beta1.ClusterServicePlan
			result3	error
		})
	}
	fake.instanceToServiceClassAndPlanReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ClusterServiceClass
		result2	*apiv1beta1.ClusterServicePlan
		result3	error
	}{result1, result2, result3}
}
func (fake *FakeSvcatClient) IsInstanceFailed(arg1 *apiv1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isInstanceFailedMutex.Lock()
	ret, specificReturn := fake.isInstanceFailedReturnsOnCall[len(fake.isInstanceFailedArgsForCall)]
	fake.isInstanceFailedArgsForCall = append(fake.isInstanceFailedArgsForCall, struct{ arg1 *apiv1beta1.ServiceInstance }{arg1})
	fake.recordInvocation("IsInstanceFailed", []interface{}{arg1})
	fake.isInstanceFailedMutex.Unlock()
	if fake.IsInstanceFailedStub != nil {
		return fake.IsInstanceFailedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isInstanceFailedReturns.result1
}
func (fake *FakeSvcatClient) IsInstanceFailedCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isInstanceFailedMutex.RLock()
	defer fake.isInstanceFailedMutex.RUnlock()
	return len(fake.isInstanceFailedArgsForCall)
}
func (fake *FakeSvcatClient) IsInstanceFailedArgsForCall(i int) *apiv1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isInstanceFailedMutex.RLock()
	defer fake.isInstanceFailedMutex.RUnlock()
	return fake.isInstanceFailedArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) IsInstanceFailedReturns(result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsInstanceFailedStub = nil
	fake.isInstanceFailedReturns = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) IsInstanceFailedReturnsOnCall(i int, result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsInstanceFailedStub = nil
	if fake.isInstanceFailedReturnsOnCall == nil {
		fake.isInstanceFailedReturnsOnCall = make(map[int]struct{ result1 bool })
	}
	fake.isInstanceFailedReturnsOnCall[i] = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) IsInstanceReady(arg1 *apiv1beta1.ServiceInstance) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isInstanceReadyMutex.Lock()
	ret, specificReturn := fake.isInstanceReadyReturnsOnCall[len(fake.isInstanceReadyArgsForCall)]
	fake.isInstanceReadyArgsForCall = append(fake.isInstanceReadyArgsForCall, struct{ arg1 *apiv1beta1.ServiceInstance }{arg1})
	fake.recordInvocation("IsInstanceReady", []interface{}{arg1})
	fake.isInstanceReadyMutex.Unlock()
	if fake.IsInstanceReadyStub != nil {
		return fake.IsInstanceReadyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isInstanceReadyReturns.result1
}
func (fake *FakeSvcatClient) IsInstanceReadyCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isInstanceReadyMutex.RLock()
	defer fake.isInstanceReadyMutex.RUnlock()
	return len(fake.isInstanceReadyArgsForCall)
}
func (fake *FakeSvcatClient) IsInstanceReadyArgsForCall(i int) *apiv1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.isInstanceReadyMutex.RLock()
	defer fake.isInstanceReadyMutex.RUnlock()
	return fake.isInstanceReadyArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) IsInstanceReadyReturns(result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsInstanceReadyStub = nil
	fake.isInstanceReadyReturns = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) IsInstanceReadyReturnsOnCall(i int, result1 bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.IsInstanceReadyStub = nil
	if fake.isInstanceReadyReturnsOnCall == nil {
		fake.isInstanceReadyReturnsOnCall = make(map[int]struct{ result1 bool })
	}
	fake.isInstanceReadyReturnsOnCall[i] = struct{ result1 bool }{result1}
}
func (fake *FakeSvcatClient) Provision(arg1 string, arg2 string, arg3 string, arg4 *servicecatalog.ProvisionOptions) (*apiv1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	string
		arg4	*servicecatalog.ProvisionOptions
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Provision", []interface{}{arg1, arg2, arg3, arg4})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.provisionReturns.result1, fake.provisionReturns.result2
}
func (fake *FakeSvcatClient) ProvisionCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}
func (fake *FakeSvcatClient) ProvisionArgsForCall(i int) (string, string, string, *servicecatalog.ProvisionOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].arg1, fake.provisionArgsForCall[i].arg2, fake.provisionArgsForCall[i].arg3, fake.provisionArgsForCall[i].arg4
}
func (fake *FakeSvcatClient) ProvisionReturns(result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) ProvisionReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstance
			result2	error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstance(arg1 string, arg2 string) (*apiv1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstanceMutex.Lock()
	ret, specificReturn := fake.retrieveInstanceReturnsOnCall[len(fake.retrieveInstanceArgsForCall)]
	fake.retrieveInstanceArgsForCall = append(fake.retrieveInstanceArgsForCall, struct {
		arg1	string
		arg2	string
	}{arg1, arg2})
	fake.recordInvocation("RetrieveInstance", []interface{}{arg1, arg2})
	fake.retrieveInstanceMutex.Unlock()
	if fake.RetrieveInstanceStub != nil {
		return fake.RetrieveInstanceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveInstanceReturns.result1, fake.retrieveInstanceReturns.result2
}
func (fake *FakeSvcatClient) RetrieveInstanceCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstanceMutex.RLock()
	defer fake.retrieveInstanceMutex.RUnlock()
	return len(fake.retrieveInstanceArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveInstanceArgsForCall(i int) (string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstanceMutex.RLock()
	defer fake.retrieveInstanceMutex.RUnlock()
	return fake.retrieveInstanceArgsForCall[i].arg1, fake.retrieveInstanceArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) RetrieveInstanceReturns(result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstanceStub = nil
	fake.retrieveInstanceReturns = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstanceReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstanceStub = nil
	if fake.retrieveInstanceReturnsOnCall == nil {
		fake.retrieveInstanceReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstance
			result2	error
		})
	}
	fake.retrieveInstanceReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstanceByBinding(arg1 *apiv1beta1.ServiceBinding) (*apiv1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstanceByBindingMutex.Lock()
	ret, specificReturn := fake.retrieveInstanceByBindingReturnsOnCall[len(fake.retrieveInstanceByBindingArgsForCall)]
	fake.retrieveInstanceByBindingArgsForCall = append(fake.retrieveInstanceByBindingArgsForCall, struct{ arg1 *apiv1beta1.ServiceBinding }{arg1})
	fake.recordInvocation("RetrieveInstanceByBinding", []interface{}{arg1})
	fake.retrieveInstanceByBindingMutex.Unlock()
	if fake.RetrieveInstanceByBindingStub != nil {
		return fake.RetrieveInstanceByBindingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveInstanceByBindingReturns.result1, fake.retrieveInstanceByBindingReturns.result2
}
func (fake *FakeSvcatClient) RetrieveInstanceByBindingCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstanceByBindingMutex.RLock()
	defer fake.retrieveInstanceByBindingMutex.RUnlock()
	return len(fake.retrieveInstanceByBindingArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveInstanceByBindingArgsForCall(i int) *apiv1beta1.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstanceByBindingMutex.RLock()
	defer fake.retrieveInstanceByBindingMutex.RUnlock()
	return fake.retrieveInstanceByBindingArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveInstanceByBindingReturns(result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstanceByBindingStub = nil
	fake.retrieveInstanceByBindingReturns = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstanceByBindingReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstanceByBindingStub = nil
	if fake.retrieveInstanceByBindingReturnsOnCall == nil {
		fake.retrieveInstanceByBindingReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstance
			result2	error
		})
	}
	fake.retrieveInstanceByBindingReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstances(arg1 string, arg2 string, arg3 string) (*apiv1beta1.ServiceInstanceList, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstancesMutex.Lock()
	ret, specificReturn := fake.retrieveInstancesReturnsOnCall[len(fake.retrieveInstancesArgsForCall)]
	fake.retrieveInstancesArgsForCall = append(fake.retrieveInstancesArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	string
	}{arg1, arg2, arg3})
	fake.recordInvocation("RetrieveInstances", []interface{}{arg1, arg2, arg3})
	fake.retrieveInstancesMutex.Unlock()
	if fake.RetrieveInstancesStub != nil {
		return fake.RetrieveInstancesStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveInstancesReturns.result1, fake.retrieveInstancesReturns.result2
}
func (fake *FakeSvcatClient) RetrieveInstancesCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstancesMutex.RLock()
	defer fake.retrieveInstancesMutex.RUnlock()
	return len(fake.retrieveInstancesArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveInstancesArgsForCall(i int) (string, string, string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstancesMutex.RLock()
	defer fake.retrieveInstancesMutex.RUnlock()
	return fake.retrieveInstancesArgsForCall[i].arg1, fake.retrieveInstancesArgsForCall[i].arg2, fake.retrieveInstancesArgsForCall[i].arg3
}
func (fake *FakeSvcatClient) RetrieveInstancesReturns(result1 *apiv1beta1.ServiceInstanceList, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstancesStub = nil
	fake.retrieveInstancesReturns = struct {
		result1	*apiv1beta1.ServiceInstanceList
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstancesReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstanceList, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstancesStub = nil
	if fake.retrieveInstancesReturnsOnCall == nil {
		fake.retrieveInstancesReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstanceList
			result2	error
		})
	}
	fake.retrieveInstancesReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstanceList
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstancesByPlan(arg1 servicecatalog.Plan) ([]apiv1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstancesByPlanMutex.Lock()
	ret, specificReturn := fake.retrieveInstancesByPlanReturnsOnCall[len(fake.retrieveInstancesByPlanArgsForCall)]
	fake.retrieveInstancesByPlanArgsForCall = append(fake.retrieveInstancesByPlanArgsForCall, struct{ arg1 servicecatalog.Plan }{arg1})
	fake.recordInvocation("RetrieveInstancesByPlan", []interface{}{arg1})
	fake.retrieveInstancesByPlanMutex.Unlock()
	if fake.RetrieveInstancesByPlanStub != nil {
		return fake.RetrieveInstancesByPlanStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveInstancesByPlanReturns.result1, fake.retrieveInstancesByPlanReturns.result2
}
func (fake *FakeSvcatClient) RetrieveInstancesByPlanCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstancesByPlanMutex.RLock()
	defer fake.retrieveInstancesByPlanMutex.RUnlock()
	return len(fake.retrieveInstancesByPlanArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveInstancesByPlanArgsForCall(i int) servicecatalog.Plan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveInstancesByPlanMutex.RLock()
	defer fake.retrieveInstancesByPlanMutex.RUnlock()
	return fake.retrieveInstancesByPlanArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveInstancesByPlanReturns(result1 []apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstancesByPlanStub = nil
	fake.retrieveInstancesByPlanReturns = struct {
		result1	[]apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveInstancesByPlanReturnsOnCall(i int, result1 []apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveInstancesByPlanStub = nil
	if fake.retrieveInstancesByPlanReturnsOnCall == nil {
		fake.retrieveInstancesByPlanReturnsOnCall = make(map[int]struct {
			result1	[]apiv1beta1.ServiceInstance
			result2	error
		})
	}
	fake.retrieveInstancesByPlanReturnsOnCall[i] = struct {
		result1	[]apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) TouchInstance(arg1 string, arg2 string, arg3 int) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.touchInstanceMutex.Lock()
	ret, specificReturn := fake.touchInstanceReturnsOnCall[len(fake.touchInstanceArgsForCall)]
	fake.touchInstanceArgsForCall = append(fake.touchInstanceArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	int
	}{arg1, arg2, arg3})
	fake.recordInvocation("TouchInstance", []interface{}{arg1, arg2, arg3})
	fake.touchInstanceMutex.Unlock()
	if fake.TouchInstanceStub != nil {
		return fake.TouchInstanceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.touchInstanceReturns.result1
}
func (fake *FakeSvcatClient) TouchInstanceCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.touchInstanceMutex.RLock()
	defer fake.touchInstanceMutex.RUnlock()
	return len(fake.touchInstanceArgsForCall)
}
func (fake *FakeSvcatClient) TouchInstanceArgsForCall(i int) (string, string, int) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.touchInstanceMutex.RLock()
	defer fake.touchInstanceMutex.RUnlock()
	return fake.touchInstanceArgsForCall[i].arg1, fake.touchInstanceArgsForCall[i].arg2, fake.touchInstanceArgsForCall[i].arg3
}
func (fake *FakeSvcatClient) TouchInstanceReturns(result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.TouchInstanceStub = nil
	fake.touchInstanceReturns = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) TouchInstanceReturnsOnCall(i int, result1 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.TouchInstanceStub = nil
	if fake.touchInstanceReturnsOnCall == nil {
		fake.touchInstanceReturnsOnCall = make(map[int]struct{ result1 error })
	}
	fake.touchInstanceReturnsOnCall[i] = struct{ result1 error }{result1}
}
func (fake *FakeSvcatClient) WaitForInstance(arg1 string, arg2 string, arg3 time.Duration, arg4 *time.Duration) (*apiv1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForInstanceMutex.Lock()
	ret, specificReturn := fake.waitForInstanceReturnsOnCall[len(fake.waitForInstanceArgsForCall)]
	fake.waitForInstanceArgsForCall = append(fake.waitForInstanceArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	time.Duration
		arg4	*time.Duration
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("WaitForInstance", []interface{}{arg1, arg2, arg3, arg4})
	fake.waitForInstanceMutex.Unlock()
	if fake.WaitForInstanceStub != nil {
		return fake.WaitForInstanceStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.waitForInstanceReturns.result1, fake.waitForInstanceReturns.result2
}
func (fake *FakeSvcatClient) WaitForInstanceCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForInstanceMutex.RLock()
	defer fake.waitForInstanceMutex.RUnlock()
	return len(fake.waitForInstanceArgsForCall)
}
func (fake *FakeSvcatClient) WaitForInstanceArgsForCall(i int) (string, string, time.Duration, *time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForInstanceMutex.RLock()
	defer fake.waitForInstanceMutex.RUnlock()
	return fake.waitForInstanceArgsForCall[i].arg1, fake.waitForInstanceArgsForCall[i].arg2, fake.waitForInstanceArgsForCall[i].arg3, fake.waitForInstanceArgsForCall[i].arg4
}
func (fake *FakeSvcatClient) WaitForInstanceReturns(result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForInstanceStub = nil
	fake.waitForInstanceReturns = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) WaitForInstanceReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForInstanceStub = nil
	if fake.waitForInstanceReturnsOnCall == nil {
		fake.waitForInstanceReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstance
			result2	error
		})
	}
	fake.waitForInstanceReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) WaitForInstanceToNotExist(arg1 string, arg2 string, arg3 time.Duration, arg4 *time.Duration) (*apiv1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForInstanceToNotExistMutex.Lock()
	ret, specificReturn := fake.waitForInstanceToNotExistReturnsOnCall[len(fake.waitForInstanceToNotExistArgsForCall)]
	fake.waitForInstanceToNotExistArgsForCall = append(fake.waitForInstanceToNotExistArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	time.Duration
		arg4	*time.Duration
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("WaitForInstanceToNotExist", []interface{}{arg1, arg2, arg3, arg4})
	fake.waitForInstanceToNotExistMutex.Unlock()
	if fake.WaitForInstanceToNotExistStub != nil {
		return fake.WaitForInstanceToNotExistStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.waitForInstanceToNotExistReturns.result1, fake.waitForInstanceToNotExistReturns.result2
}
func (fake *FakeSvcatClient) WaitForInstanceToNotExistCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForInstanceToNotExistMutex.RLock()
	defer fake.waitForInstanceToNotExistMutex.RUnlock()
	return len(fake.waitForInstanceToNotExistArgsForCall)
}
func (fake *FakeSvcatClient) WaitForInstanceToNotExistArgsForCall(i int) (string, string, time.Duration, *time.Duration) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.waitForInstanceToNotExistMutex.RLock()
	defer fake.waitForInstanceToNotExistMutex.RUnlock()
	return fake.waitForInstanceToNotExistArgsForCall[i].arg1, fake.waitForInstanceToNotExistArgsForCall[i].arg2, fake.waitForInstanceToNotExistArgsForCall[i].arg3, fake.waitForInstanceToNotExistArgsForCall[i].arg4
}
func (fake *FakeSvcatClient) WaitForInstanceToNotExistReturns(result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForInstanceToNotExistStub = nil
	fake.waitForInstanceToNotExistReturns = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) WaitForInstanceToNotExistReturnsOnCall(i int, result1 *apiv1beta1.ServiceInstance, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.WaitForInstanceToNotExistStub = nil
	if fake.waitForInstanceToNotExistReturnsOnCall == nil {
		fake.waitForInstanceToNotExistReturnsOnCall = make(map[int]struct {
			result1	*apiv1beta1.ServiceInstance
			result2	error
		})
	}
	fake.waitForInstanceToNotExistReturnsOnCall[i] = struct {
		result1	*apiv1beta1.ServiceInstance
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlans(arg1 string, arg2 servicecatalog.ScopeOptions) ([]servicecatalog.Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlansMutex.Lock()
	ret, specificReturn := fake.retrievePlansReturnsOnCall[len(fake.retrievePlansArgsForCall)]
	fake.retrievePlansArgsForCall = append(fake.retrievePlansArgsForCall, struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}{arg1, arg2})
	fake.recordInvocation("RetrievePlans", []interface{}{arg1, arg2})
	fake.retrievePlansMutex.Unlock()
	if fake.RetrievePlansStub != nil {
		return fake.RetrievePlansStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrievePlansReturns.result1, fake.retrievePlansReturns.result2
}
func (fake *FakeSvcatClient) RetrievePlansCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlansMutex.RLock()
	defer fake.retrievePlansMutex.RUnlock()
	return len(fake.retrievePlansArgsForCall)
}
func (fake *FakeSvcatClient) RetrievePlansArgsForCall(i int) (string, servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlansMutex.RLock()
	defer fake.retrievePlansMutex.RUnlock()
	return fake.retrievePlansArgsForCall[i].arg1, fake.retrievePlansArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) RetrievePlansReturns(result1 []servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlansStub = nil
	fake.retrievePlansReturns = struct {
		result1	[]servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlansReturnsOnCall(i int, result1 []servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlansStub = nil
	if fake.retrievePlansReturnsOnCall == nil {
		fake.retrievePlansReturnsOnCall = make(map[int]struct {
			result1	[]servicecatalog.Plan
			result2	error
		})
	}
	fake.retrievePlansReturnsOnCall[i] = struct {
		result1	[]servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByName(arg1 string, arg2 servicecatalog.ScopeOptions) (servicecatalog.Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByNameMutex.Lock()
	ret, specificReturn := fake.retrievePlanByNameReturnsOnCall[len(fake.retrievePlanByNameArgsForCall)]
	fake.retrievePlanByNameArgsForCall = append(fake.retrievePlanByNameArgsForCall, struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}{arg1, arg2})
	fake.recordInvocation("RetrievePlanByName", []interface{}{arg1, arg2})
	fake.retrievePlanByNameMutex.Unlock()
	if fake.RetrievePlanByNameStub != nil {
		return fake.RetrievePlanByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrievePlanByNameReturns.result1, fake.retrievePlanByNameReturns.result2
}
func (fake *FakeSvcatClient) RetrievePlanByNameCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByNameMutex.RLock()
	defer fake.retrievePlanByNameMutex.RUnlock()
	return len(fake.retrievePlanByNameArgsForCall)
}
func (fake *FakeSvcatClient) RetrievePlanByNameArgsForCall(i int) (string, servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByNameMutex.RLock()
	defer fake.retrievePlanByNameMutex.RUnlock()
	return fake.retrievePlanByNameArgsForCall[i].arg1, fake.retrievePlanByNameArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) RetrievePlanByNameReturns(result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByNameStub = nil
	fake.retrievePlanByNameReturns = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByNameReturnsOnCall(i int, result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByNameStub = nil
	if fake.retrievePlanByNameReturnsOnCall == nil {
		fake.retrievePlanByNameReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Plan
			result2	error
		})
	}
	fake.retrievePlanByNameReturnsOnCall[i] = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByClassAndName(arg1 string, arg2 string, arg3 servicecatalog.ScopeOptions) (servicecatalog.Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByClassAndNameMutex.Lock()
	ret, specificReturn := fake.retrievePlanByClassAndNameReturnsOnCall[len(fake.retrievePlanByClassAndNameArgsForCall)]
	fake.retrievePlanByClassAndNameArgsForCall = append(fake.retrievePlanByClassAndNameArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	servicecatalog.ScopeOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("RetrievePlanByClassAndName", []interface{}{arg1, arg2, arg3})
	fake.retrievePlanByClassAndNameMutex.Unlock()
	if fake.RetrievePlanByClassAndNameStub != nil {
		return fake.RetrievePlanByClassAndNameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrievePlanByClassAndNameReturns.result1, fake.retrievePlanByClassAndNameReturns.result2
}
func (fake *FakeSvcatClient) RetrievePlanByClassAndNameCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByClassAndNameMutex.RLock()
	defer fake.retrievePlanByClassAndNameMutex.RUnlock()
	return len(fake.retrievePlanByClassAndNameArgsForCall)
}
func (fake *FakeSvcatClient) RetrievePlanByClassAndNameArgsForCall(i int) (string, string, servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByClassAndNameMutex.RLock()
	defer fake.retrievePlanByClassAndNameMutex.RUnlock()
	return fake.retrievePlanByClassAndNameArgsForCall[i].arg1, fake.retrievePlanByClassAndNameArgsForCall[i].arg2, fake.retrievePlanByClassAndNameArgsForCall[i].arg3
}
func (fake *FakeSvcatClient) RetrievePlanByClassAndNameReturns(result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByClassAndNameStub = nil
	fake.retrievePlanByClassAndNameReturns = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByClassAndNameReturnsOnCall(i int, result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByClassAndNameStub = nil
	if fake.retrievePlanByClassAndNameReturnsOnCall == nil {
		fake.retrievePlanByClassAndNameReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Plan
			result2	error
		})
	}
	fake.retrievePlanByClassAndNameReturnsOnCall[i] = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByClassIDAndName(arg1 string, arg2 string, arg3 servicecatalog.ScopeOptions) (servicecatalog.Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByClassIDAndNameMutex.Lock()
	ret, specificReturn := fake.retrievePlanByClassIDAndNameReturnsOnCall[len(fake.retrievePlanByClassIDAndNameArgsForCall)]
	fake.retrievePlanByClassIDAndNameArgsForCall = append(fake.retrievePlanByClassIDAndNameArgsForCall, struct {
		arg1	string
		arg2	string
		arg3	servicecatalog.ScopeOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("RetrievePlanByClassIDAndName", []interface{}{arg1, arg2, arg3})
	fake.retrievePlanByClassIDAndNameMutex.Unlock()
	if fake.RetrievePlanByClassIDAndNameStub != nil {
		return fake.RetrievePlanByClassIDAndNameStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrievePlanByClassIDAndNameReturns.result1, fake.retrievePlanByClassIDAndNameReturns.result2
}
func (fake *FakeSvcatClient) RetrievePlanByClassIDAndNameCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByClassIDAndNameMutex.RLock()
	defer fake.retrievePlanByClassIDAndNameMutex.RUnlock()
	return len(fake.retrievePlanByClassIDAndNameArgsForCall)
}
func (fake *FakeSvcatClient) RetrievePlanByClassIDAndNameArgsForCall(i int) (string, string, servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByClassIDAndNameMutex.RLock()
	defer fake.retrievePlanByClassIDAndNameMutex.RUnlock()
	return fake.retrievePlanByClassIDAndNameArgsForCall[i].arg1, fake.retrievePlanByClassIDAndNameArgsForCall[i].arg2, fake.retrievePlanByClassIDAndNameArgsForCall[i].arg3
}
func (fake *FakeSvcatClient) RetrievePlanByClassIDAndNameReturns(result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByClassIDAndNameStub = nil
	fake.retrievePlanByClassIDAndNameReturns = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByClassIDAndNameReturnsOnCall(i int, result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByClassIDAndNameStub = nil
	if fake.retrievePlanByClassIDAndNameReturnsOnCall == nil {
		fake.retrievePlanByClassIDAndNameReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Plan
			result2	error
		})
	}
	fake.retrievePlanByClassIDAndNameReturnsOnCall[i] = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByID(arg1 string, arg2 servicecatalog.ScopeOptions) (servicecatalog.Plan, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByIDMutex.Lock()
	ret, specificReturn := fake.retrievePlanByIDReturnsOnCall[len(fake.retrievePlanByIDArgsForCall)]
	fake.retrievePlanByIDArgsForCall = append(fake.retrievePlanByIDArgsForCall, struct {
		arg1	string
		arg2	servicecatalog.ScopeOptions
	}{arg1, arg2})
	fake.recordInvocation("RetrievePlanByID", []interface{}{arg1, arg2})
	fake.retrievePlanByIDMutex.Unlock()
	if fake.RetrievePlanByIDStub != nil {
		return fake.RetrievePlanByIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrievePlanByIDReturns.result1, fake.retrievePlanByIDReturns.result2
}
func (fake *FakeSvcatClient) RetrievePlanByIDCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByIDMutex.RLock()
	defer fake.retrievePlanByIDMutex.RUnlock()
	return len(fake.retrievePlanByIDArgsForCall)
}
func (fake *FakeSvcatClient) RetrievePlanByIDArgsForCall(i int) (string, servicecatalog.ScopeOptions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrievePlanByIDMutex.RLock()
	defer fake.retrievePlanByIDMutex.RUnlock()
	return fake.retrievePlanByIDArgsForCall[i].arg1, fake.retrievePlanByIDArgsForCall[i].arg2
}
func (fake *FakeSvcatClient) RetrievePlanByIDReturns(result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByIDStub = nil
	fake.retrievePlanByIDReturns = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrievePlanByIDReturnsOnCall(i int, result1 servicecatalog.Plan, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrievePlanByIDStub = nil
	if fake.retrievePlanByIDReturnsOnCall == nil {
		fake.retrievePlanByIDReturnsOnCall = make(map[int]struct {
			result1	servicecatalog.Plan
			result2	error
		})
	}
	fake.retrievePlanByIDReturnsOnCall[i] = struct {
		result1	servicecatalog.Plan
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveSecretByBinding(arg1 *apiv1beta1.ServiceBinding) (*apicorev1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveSecretByBindingMutex.Lock()
	ret, specificReturn := fake.retrieveSecretByBindingReturnsOnCall[len(fake.retrieveSecretByBindingArgsForCall)]
	fake.retrieveSecretByBindingArgsForCall = append(fake.retrieveSecretByBindingArgsForCall, struct{ arg1 *apiv1beta1.ServiceBinding }{arg1})
	fake.recordInvocation("RetrieveSecretByBinding", []interface{}{arg1})
	fake.retrieveSecretByBindingMutex.Unlock()
	if fake.RetrieveSecretByBindingStub != nil {
		return fake.RetrieveSecretByBindingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrieveSecretByBindingReturns.result1, fake.retrieveSecretByBindingReturns.result2
}
func (fake *FakeSvcatClient) RetrieveSecretByBindingCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveSecretByBindingMutex.RLock()
	defer fake.retrieveSecretByBindingMutex.RUnlock()
	return len(fake.retrieveSecretByBindingArgsForCall)
}
func (fake *FakeSvcatClient) RetrieveSecretByBindingArgsForCall(i int) *apiv1beta1.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.retrieveSecretByBindingMutex.RLock()
	defer fake.retrieveSecretByBindingMutex.RUnlock()
	return fake.retrieveSecretByBindingArgsForCall[i].arg1
}
func (fake *FakeSvcatClient) RetrieveSecretByBindingReturns(result1 *apicorev1.Secret, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveSecretByBindingStub = nil
	fake.retrieveSecretByBindingReturns = struct {
		result1	*apicorev1.Secret
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) RetrieveSecretByBindingReturnsOnCall(i int, result1 *apicorev1.Secret, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.RetrieveSecretByBindingStub = nil
	if fake.retrieveSecretByBindingReturnsOnCall == nil {
		fake.retrieveSecretByBindingReturnsOnCall = make(map[int]struct {
			result1	*apicorev1.Secret
			result2	error
		})
	}
	fake.retrieveSecretByBindingReturnsOnCall[i] = struct {
		result1	*apicorev1.Secret
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) ServerVersion() (*version.Info, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.serverVersionMutex.Lock()
	ret, specificReturn := fake.serverVersionReturnsOnCall[len(fake.serverVersionArgsForCall)]
	fake.serverVersionArgsForCall = append(fake.serverVersionArgsForCall, struct{}{})
	fake.recordInvocation("ServerVersion", []interface{}{})
	fake.serverVersionMutex.Unlock()
	if fake.ServerVersionStub != nil {
		return fake.ServerVersionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.serverVersionReturns.result1, fake.serverVersionReturns.result2
}
func (fake *FakeSvcatClient) ServerVersionCallCount() int {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	return len(fake.serverVersionArgsForCall)
}
func (fake *FakeSvcatClient) ServerVersionReturns(result1 *version.Info, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.ServerVersionStub = nil
	fake.serverVersionReturns = struct {
		result1	*version.Info
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) ServerVersionReturnsOnCall(i int, result1 *version.Info, result2 error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.ServerVersionStub = nil
	if fake.serverVersionReturnsOnCall == nil {
		fake.serverVersionReturnsOnCall = make(map[int]struct {
			result1	*version.Info
			result2	error
		})
	}
	fake.serverVersionReturnsOnCall[i] = struct {
		result1	*version.Info
		result2	error
	}{result1, result2}
}
func (fake *FakeSvcatClient) Invocations() map[string][][]interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.bindingParentHierarchyMutex.RLock()
	defer fake.bindingParentHierarchyMutex.RUnlock()
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	fake.deleteBindingsMutex.RLock()
	defer fake.deleteBindingsMutex.RUnlock()
	fake.isBindingFailedMutex.RLock()
	defer fake.isBindingFailedMutex.RUnlock()
	fake.isBindingReadyMutex.RLock()
	defer fake.isBindingReadyMutex.RUnlock()
	fake.retrieveBindingMutex.RLock()
	defer fake.retrieveBindingMutex.RUnlock()
	fake.retrieveBindingsMutex.RLock()
	defer fake.retrieveBindingsMutex.RUnlock()
	fake.retrieveBindingsByInstanceMutex.RLock()
	defer fake.retrieveBindingsByInstanceMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.waitForBindingMutex.RLock()
	defer fake.waitForBindingMutex.RUnlock()
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	fake.retrieveBrokersMutex.RLock()
	defer fake.retrieveBrokersMutex.RUnlock()
	fake.retrieveBrokerMutex.RLock()
	defer fake.retrieveBrokerMutex.RUnlock()
	fake.retrieveBrokerByClassMutex.RLock()
	defer fake.retrieveBrokerByClassMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	fake.waitForBrokerMutex.RLock()
	defer fake.waitForBrokerMutex.RUnlock()
	fake.retrieveClassesMutex.RLock()
	defer fake.retrieveClassesMutex.RUnlock()
	fake.retrieveClassByNameMutex.RLock()
	defer fake.retrieveClassByNameMutex.RUnlock()
	fake.retrieveClassByIDMutex.RLock()
	defer fake.retrieveClassByIDMutex.RUnlock()
	fake.retrieveClassByPlanMutex.RLock()
	defer fake.retrieveClassByPlanMutex.RUnlock()
	fake.createClassFromMutex.RLock()
	defer fake.createClassFromMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.instanceParentHierarchyMutex.RLock()
	defer fake.instanceParentHierarchyMutex.RUnlock()
	fake.instanceToServiceClassAndPlanMutex.RLock()
	defer fake.instanceToServiceClassAndPlanMutex.RUnlock()
	fake.isInstanceFailedMutex.RLock()
	defer fake.isInstanceFailedMutex.RUnlock()
	fake.isInstanceReadyMutex.RLock()
	defer fake.isInstanceReadyMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.retrieveInstanceMutex.RLock()
	defer fake.retrieveInstanceMutex.RUnlock()
	fake.retrieveInstanceByBindingMutex.RLock()
	defer fake.retrieveInstanceByBindingMutex.RUnlock()
	fake.retrieveInstancesMutex.RLock()
	defer fake.retrieveInstancesMutex.RUnlock()
	fake.retrieveInstancesByPlanMutex.RLock()
	defer fake.retrieveInstancesByPlanMutex.RUnlock()
	fake.touchInstanceMutex.RLock()
	defer fake.touchInstanceMutex.RUnlock()
	fake.waitForInstanceMutex.RLock()
	defer fake.waitForInstanceMutex.RUnlock()
	fake.waitForInstanceToNotExistMutex.RLock()
	defer fake.waitForInstanceToNotExistMutex.RUnlock()
	fake.retrievePlansMutex.RLock()
	defer fake.retrievePlansMutex.RUnlock()
	fake.retrievePlanByNameMutex.RLock()
	defer fake.retrievePlanByNameMutex.RUnlock()
	fake.retrievePlanByClassAndNameMutex.RLock()
	defer fake.retrievePlanByClassAndNameMutex.RUnlock()
	fake.retrievePlanByClassIDAndNameMutex.RLock()
	defer fake.retrievePlanByClassIDAndNameMutex.RUnlock()
	fake.retrievePlanByIDMutex.RLock()
	defer fake.retrievePlanByIDMutex.RUnlock()
	fake.retrieveSecretByBindingMutex.RLock()
	defer fake.retrieveSecretByBindingMutex.RUnlock()
	fake.serverVersionMutex.RLock()
	defer fake.serverVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}
func (fake *FakeSvcatClient) recordInvocation(key string, args []interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ servicecatalog.SvcatClient = new(FakeSvcatClient)

func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
