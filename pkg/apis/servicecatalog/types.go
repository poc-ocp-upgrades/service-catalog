package servicecatalog

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterServiceBroker struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ClusterServiceBrokerSpec
	Status	ClusterServiceBrokerStatus
}
type ClusterServiceBrokerList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ClusterServiceBroker
}
type ServiceBroker struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ServiceBrokerSpec
	Status	ServiceBrokerStatus
}
type ServiceBrokerList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ServiceBroker
}
type CommonServiceBrokerSpec struct {
	URL						string
	InsecureSkipTLSVerify	bool
	CABundle				[]byte
	RelistBehavior			ServiceBrokerRelistBehavior
	RelistDuration			*metav1.Duration
	RelistRequests			int64
	CatalogRestrictions		*CatalogRestrictions
}
type CatalogRestrictions struct {
	ServiceClass	[]string
	ServicePlan		[]string
}
type ClusterServiceBrokerSpec struct {
	CommonServiceBrokerSpec
	AuthInfo	*ClusterServiceBrokerAuthInfo
}
type ServiceBrokerSpec struct {
	CommonServiceBrokerSpec
	AuthInfo	*ServiceBrokerAuthInfo
}
type ServiceBrokerRelistBehavior string

const (
	ServiceBrokerRelistBehaviorDuration	ServiceBrokerRelistBehavior	= "Duration"
	ServiceBrokerRelistBehaviorManual	ServiceBrokerRelistBehavior	= "Manual"
)

type ClusterServiceBrokerAuthInfo struct {
	Basic	*ClusterBasicAuthConfig
	Bearer	*ClusterBearerTokenAuthConfig
}
type ClusterBasicAuthConfig struct{ SecretRef *ObjectReference }
type ClusterBearerTokenAuthConfig struct{ SecretRef *ObjectReference }
type ServiceBrokerAuthInfo struct {
	Basic	*BasicAuthConfig
	Bearer	*BearerTokenAuthConfig
}
type BasicAuthConfig struct{ SecretRef *LocalObjectReference }
type BearerTokenAuthConfig struct{ SecretRef *LocalObjectReference }

const (
	BasicAuthUsernameKey	= "username"
	BasicAuthPasswordKey	= "password"
	BearerTokenKey			= "token"
)

type CommonServiceBrokerStatus struct {
	Conditions					[]ServiceBrokerCondition
	ReconciledGeneration		int64
	OperationStartTime			*metav1.Time
	LastCatalogRetrievalTime	*metav1.Time
}
type ClusterServiceBrokerStatus struct{ CommonServiceBrokerStatus }
type ServiceBrokerStatus struct{ CommonServiceBrokerStatus }
type ServiceBrokerCondition struct {
	Type				ServiceBrokerConditionType
	Status				ConditionStatus
	LastTransitionTime	metav1.Time
	Reason				string
	Message				string
}
type ServiceBrokerConditionType string

const (
	ServiceBrokerConditionReady		ServiceBrokerConditionType	= "Ready"
	ServiceBrokerConditionFailed	ServiceBrokerConditionType	= "Failed"
)

type ConditionStatus string

const (
	ConditionTrue		ConditionStatus	= "True"
	ConditionFalse		ConditionStatus	= "False"
	ConditionUnknown	ConditionStatus	= "Unknown"
)

type ClusterServiceClassList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ClusterServiceClass
}
type ClusterServiceClass struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ClusterServiceClassSpec
	Status	ClusterServiceClassStatus
}
type ServiceClassList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ServiceClass
}
type ServiceClass struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ServiceClassSpec
	Status	ServiceClassStatus
}
type ServiceClassStatus struct{ CommonServiceClassStatus }
type ClusterServiceClassStatus struct{ CommonServiceClassStatus }
type CommonServiceClassStatus struct{ RemovedFromBrokerCatalog bool }
type CommonServiceClassSpec struct {
	ExternalName				string
	ExternalID					string
	Description					string
	Bindable					bool
	BindingRetrievable			bool
	PlanUpdatable				bool
	ExternalMetadata			*runtime.RawExtension
	Tags						[]string
	Requires					[]string
	DefaultProvisionParameters	*runtime.RawExtension
}
type ClusterServiceClassSpec struct {
	CommonServiceClassSpec
	ClusterServiceBrokerName	string
}
type ServiceClassSpec struct {
	CommonServiceClassSpec
	ServiceBrokerName	string
}
type ClusterServicePlanList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ClusterServicePlan
}
type ClusterServicePlan struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ClusterServicePlanSpec
	Status	ClusterServicePlanStatus
}
type CommonServicePlanSpec struct {
	ExternalName						string
	ExternalID							string
	Description							string
	Bindable							*bool
	Free								bool
	ExternalMetadata					*runtime.RawExtension
	InstanceCreateParameterSchema		*runtime.RawExtension
	InstanceUpdateParameterSchema		*runtime.RawExtension
	ServiceBindingCreateParameterSchema	*runtime.RawExtension
	ServiceBindingCreateResponseSchema	*runtime.RawExtension
	DefaultProvisionParameters			*runtime.RawExtension
}
type ClusterServicePlanSpec struct {
	CommonServicePlanSpec
	ClusterServiceBrokerName	string
	ClusterServiceClassRef		ClusterObjectReference
}
type ClusterServicePlanStatus struct{ CommonServicePlanStatus }
type CommonServicePlanStatus struct{ RemovedFromBrokerCatalog bool }
type ServicePlanList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ServicePlan
}
type ServicePlan struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ServicePlanSpec
	Status	ServicePlanStatus
}
type ServicePlanSpec struct {
	CommonServicePlanSpec
	ServiceBrokerName	string
	ServiceClassRef		LocalObjectReference
}
type ServicePlanStatus struct{ CommonServicePlanStatus }
type ServiceInstanceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ServiceInstance
}
type UserInfo struct {
	Username	string
	UID			string
	Groups		[]string
	Extra		map[string]ExtraValue
}
type ExtraValue []string
type ServiceInstance struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ServiceInstanceSpec
	Status	ServiceInstanceStatus
}
type PlanReference struct {
	ClusterServiceClassExternalName	string
	ClusterServicePlanExternalName	string
	ClusterServiceClassExternalID	string
	ClusterServicePlanExternalID	string
	ClusterServiceClassName			string
	ClusterServicePlanName			string
	ServiceClassExternalName		string
	ServicePlanExternalName			string
	ServiceClassExternalID			string
	ServicePlanExternalID			string
	ServiceClassName				string
	ServicePlanName					string
}
type ServiceInstanceSpec struct {
	PlanReference
	ClusterServiceClassRef	*ClusterObjectReference
	ClusterServicePlanRef	*ClusterObjectReference
	ServiceClassRef			*LocalObjectReference
	ServicePlanRef			*LocalObjectReference
	Parameters				*runtime.RawExtension
	ParametersFrom			[]ParametersFromSource
	ExternalID				string
	UserInfo				*UserInfo
	UpdateRequests			int64
}
type ServiceInstanceStatus struct {
	Conditions					[]ServiceInstanceCondition
	AsyncOpInProgress			bool
	OrphanMitigationInProgress	bool
	LastOperation				*string
	DashboardURL				*string
	CurrentOperation			ServiceInstanceOperation
	ReconciledGeneration		int64
	ObservedGeneration			int64
	OperationStartTime			*metav1.Time
	InProgressProperties		*ServiceInstancePropertiesState
	ExternalProperties			*ServiceInstancePropertiesState
	ProvisionStatus				ServiceInstanceProvisionStatus
	DeprovisionStatus			ServiceInstanceDeprovisionStatus
	DefaultProvisionParameters	*runtime.RawExtension
}
type ServiceInstanceCondition struct {
	Type				ServiceInstanceConditionType
	Status				ConditionStatus
	LastTransitionTime	metav1.Time
	Reason				string
	Message				string
}
type ServiceInstanceConditionType string

const (
	ServiceInstanceConditionReady				ServiceInstanceConditionType	= "Ready"
	ServiceInstanceConditionFailed				ServiceInstanceConditionType	= "Failed"
	ServiceInstanceConditionOrphanMitigation	ServiceInstanceConditionType	= "OrphanMitigation"
)

type ServiceInstanceOperation string

const (
	ServiceInstanceOperationProvision	ServiceInstanceOperation	= "Provision"
	ServiceInstanceOperationUpdate		ServiceInstanceOperation	= "Update"
	ServiceInstanceOperationDeprovision	ServiceInstanceOperation	= "Deprovision"
)

type ServiceInstancePropertiesState struct {
	ClusterServicePlanExternalName	string
	ClusterServicePlanExternalID	string
	ServicePlanExternalName			string
	ServicePlanExternalID			string
	Parameters						*runtime.RawExtension
	ParameterChecksum				string
	UserInfo						*UserInfo
}
type ServiceInstanceDeprovisionStatus string

const (
	ServiceInstanceDeprovisionStatusNotRequired	ServiceInstanceDeprovisionStatus	= "NotRequired"
	ServiceInstanceDeprovisionStatusRequired	ServiceInstanceDeprovisionStatus	= "Required"
	ServiceInstanceDeprovisionStatusSucceeded	ServiceInstanceDeprovisionStatus	= "Succeeded"
	ServiceInstanceDeprovisionStatusFailed		ServiceInstanceDeprovisionStatus	= "Failed"
)

type ServiceInstanceProvisionStatus string

const (
	ServiceInstanceProvisionStatusProvisioned		ServiceInstanceProvisionStatus	= "Provisioned"
	ServiceInstanceProvisionStatusNotProvisioned	ServiceInstanceProvisionStatus	= "NotProvisioned"
)

type ServiceBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]ServiceBinding
}
type ServiceBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	ServiceBindingSpec
	Status	ServiceBindingStatus
}
type ServiceBindingSpec struct {
	InstanceRef			LocalObjectReference
	Parameters			*runtime.RawExtension
	ParametersFrom		[]ParametersFromSource
	SecretName			string
	SecretTransforms	[]SecretTransform
	ExternalID			string
	UserInfo			*UserInfo
}
type ServiceBindingStatus struct {
	Conditions					[]ServiceBindingCondition
	AsyncOpInProgress			bool
	LastOperation				*string
	CurrentOperation			ServiceBindingOperation
	ReconciledGeneration		int64
	OperationStartTime			*metav1.Time
	InProgressProperties		*ServiceBindingPropertiesState
	ExternalProperties			*ServiceBindingPropertiesState
	OrphanMitigationInProgress	bool
	UnbindStatus				ServiceBindingUnbindStatus
}
type ServiceBindingCondition struct {
	Type				ServiceBindingConditionType
	Status				ConditionStatus
	LastTransitionTime	metav1.Time
	Reason				string
	Message				string
}
type ServiceBindingConditionType string

const (
	ServiceBindingConditionReady	ServiceBindingConditionType	= "Ready"
	ServiceBindingConditionFailed	ServiceBindingConditionType	= "Failed"
)

type ServiceBindingOperation string

const (
	ServiceBindingOperationBind		ServiceBindingOperation	= "Bind"
	ServiceBindingOperationUnbind	ServiceBindingOperation	= "Unbind"
)
const (
	FinalizerServiceCatalog string = "kubernetes-incubator/service-catalog"
)

type ServiceBindingPropertiesState struct {
	Parameters			*runtime.RawExtension
	ParameterChecksum	string
	UserInfo			*UserInfo
}
type ServiceBindingUnbindStatus string

const (
	ServiceBindingUnbindStatusNotRequired	ServiceBindingUnbindStatus	= "NotRequired"
	ServiceBindingUnbindStatusRequired		ServiceBindingUnbindStatus	= "Required"
	ServiceBindingUnbindStatusSucceeded		ServiceBindingUnbindStatus	= "Succeeded"
	ServiceBindingUnbindStatusFailed		ServiceBindingUnbindStatus	= "Failed"
)

type ParametersFromSource struct{ SecretKeyRef *SecretKeyReference }
type SecretKeyReference struct {
	Name	string
	Key		string
}
type ObjectReference struct {
	Namespace	string
	Name		string
}
type LocalObjectReference struct{ Name string }
type ClusterObjectReference struct{ Name string }
type SecretTransform struct {
	RenameKey	*RenameKeyTransform
	AddKey		*AddKeyTransform
	AddKeysFrom	*AddKeysFromTransform
	RemoveKey	*RemoveKeyTransform
}
type RenameKeyTransform struct {
	From	string
	To		string
}
type AddKeyTransform struct {
	Key					string
	Value				[]byte
	StringValue			*string
	JSONPathExpression	*string
}
type AddKeysFromTransform struct{ SecretRef *ObjectReference }
type RemoveKeyTransform struct{ Key string }
