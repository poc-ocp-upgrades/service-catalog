package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterServiceBroker struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ClusterServiceBrokerSpec	`json:"spec,omitempty"`
	Status			ClusterServiceBrokerStatus	`json:"status,omitempty"`
}
type ClusterServiceBrokerList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ClusterServiceBroker	`json:"items"`
}
type ServiceBroker struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ServiceBrokerSpec	`json:"spec,omitempty"`
	Status			ServiceBrokerStatus	`json:"status,omitempty"`
}
type ServiceBrokerList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ServiceBroker	`json:"items"`
}
type CommonServiceBrokerSpec struct {
	URL			string				`json:"url"`
	InsecureSkipTLSVerify	bool				`json:"insecureSkipTLSVerify,omitempty"`
	CABundle		[]byte				`json:"caBundle,omitempty"`
	RelistBehavior		ServiceBrokerRelistBehavior	`json:"relistBehavior"`
	RelistDuration		*metav1.Duration		`json:"relistDuration,omitempty"`
	RelistRequests		int64				`json:"relistRequests"`
	CatalogRestrictions	*CatalogRestrictions		`json:"catalogRestrictions,omitempty"`
}
type CatalogRestrictions struct {
	ServiceClass	[]string	`json:"serviceClass,omitempty"`
	ServicePlan	[]string	`json:"servicePlan,omitempty"`
}
type ClusterServiceBrokerSpec struct {
	CommonServiceBrokerSpec	`json:",inline"`
	AuthInfo		*ClusterServiceBrokerAuthInfo	`json:"authInfo,omitempty"`
}
type ServiceBrokerSpec struct {
	CommonServiceBrokerSpec	`json:",inline"`
	AuthInfo		*ServiceBrokerAuthInfo	`json:"authInfo,omitempty"`
}
type ServiceBrokerRelistBehavior string

const (
	ServiceBrokerRelistBehaviorDuration	ServiceBrokerRelistBehavior	= "Duration"
	ServiceBrokerRelistBehaviorManual	ServiceBrokerRelistBehavior	= "Manual"
)

type ClusterServiceBrokerAuthInfo struct {
	Basic	*ClusterBasicAuthConfig		`json:"basic,omitempty"`
	Bearer	*ClusterBearerTokenAuthConfig	`json:"bearer,omitempty"`
}
type ClusterBasicAuthConfig struct {
	SecretRef *ObjectReference `json:"secretRef,omitempty"`
}
type ClusterBearerTokenAuthConfig struct {
	SecretRef *ObjectReference `json:"secretRef,omitempty"`
}
type ServiceBrokerAuthInfo struct {
	Basic	*BasicAuthConfig	`json:"basic,omitempty"`
	Bearer	*BearerTokenAuthConfig	`json:"bearer,omitempty"`
}
type BasicAuthConfig struct {
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`
}
type BearerTokenAuthConfig struct {
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`
}

const (
	BasicAuthUsernameKey	= "username"
	BasicAuthPasswordKey	= "password"
	BearerTokenKey		= "token"
)

type CommonServiceBrokerStatus struct {
	Conditions			[]ServiceBrokerCondition	`json:"conditions"`
	ReconciledGeneration		int64				`json:"reconciledGeneration"`
	OperationStartTime		*metav1.Time			`json:"operationStartTime,omitempty"`
	LastCatalogRetrievalTime	*metav1.Time			`json:"lastCatalogRetrievalTime,omitempty"`
}
type ClusterServiceBrokerStatus struct {
	CommonServiceBrokerStatus `json:",inline"`
}
type ServiceBrokerStatus struct {
	CommonServiceBrokerStatus `json:",inline"`
}
type ServiceBrokerCondition struct {
	Type			ServiceBrokerConditionType	`json:"type"`
	Status			ConditionStatus			`json:"status"`
	LastTransitionTime	metav1.Time			`json:"lastTransitionTime"`
	Reason			string				`json:"reason"`
	Message			string				`json:"message"`
}
type ServiceBrokerConditionType string

const (
	ServiceBrokerConditionReady	ServiceBrokerConditionType	= "Ready"
	ServiceBrokerConditionFailed	ServiceBrokerConditionType	= "Failed"
)

type ConditionStatus string

const (
	ConditionTrue		ConditionStatus	= "True"
	ConditionFalse		ConditionStatus	= "False"
	ConditionUnknown	ConditionStatus	= "Unknown"
)

type ClusterServiceClassList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ClusterServiceClass	`json:"items"`
}
type ClusterServiceClass struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ClusterServiceClassSpec		`json:"spec,omitempty"`
	Status			ClusterServiceClassStatus	`json:"status,omitempty"`
}
type ServiceClassList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ServiceClass	`json:"items"`
}
type ServiceClass struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ServiceClassSpec	`json:"spec,omitempty"`
	Status			ServiceClassStatus	`json:"status,omitempty"`
}
type ServiceClassStatus struct {
	CommonServiceClassStatus `json:",inline"`
}
type ClusterServiceClassStatus struct {
	CommonServiceClassStatus `json:",inline"`
}
type CommonServiceClassStatus struct {
	RemovedFromBrokerCatalog bool `json:"removedFromBrokerCatalog"`
}
type CommonServiceClassSpec struct {
	ExternalName			string			`json:"externalName"`
	ExternalID			string			`json:"externalID"`
	Description			string			`json:"description"`
	Bindable			bool			`json:"bindable"`
	BindingRetrievable		bool			`json:"bindingRetrievable"`
	PlanUpdatable			bool			`json:"planUpdatable"`
	ExternalMetadata		*runtime.RawExtension	`json:"externalMetadata,omitempty"`
	Tags				[]string		`json:"tags,omitempty"`
	Requires			[]string		`json:"requires,omitempty"`
	DefaultProvisionParameters	*runtime.RawExtension	`json:"defaultProvisionParameters,omitempty"`
}
type ClusterServiceClassSpec struct {
	CommonServiceClassSpec		`json:",inline"`
	ClusterServiceBrokerName	string	`json:"clusterServiceBrokerName"`
}
type ServiceClassSpec struct {
	CommonServiceClassSpec	`json:",inline"`
	ServiceBrokerName	string	`json:"serviceBrokerName"`
}
type ClusterServicePlanList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ClusterServicePlan	`json:"items"`
}
type ClusterServicePlan struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ClusterServicePlanSpec		`json:"spec,omitempty"`
	Status			ClusterServicePlanStatus	`json:"status,omitempty"`
}
type CommonServicePlanSpec struct {
	ExternalName				string			`json:"externalName"`
	ExternalID				string			`json:"externalID"`
	Description				string			`json:"description"`
	Bindable				*bool			`json:"bindable,omitempty"`
	Free					bool			`json:"free"`
	ExternalMetadata			*runtime.RawExtension	`json:"externalMetadata,omitempty"`
	InstanceCreateParameterSchema		*runtime.RawExtension	`json:"instanceCreateParameterSchema,omitempty"`
	InstanceUpdateParameterSchema		*runtime.RawExtension	`json:"instanceUpdateParameterSchema,omitempty"`
	ServiceBindingCreateParameterSchema	*runtime.RawExtension	`json:"serviceBindingCreateParameterSchema,omitempty"`
	ServiceBindingCreateResponseSchema	*runtime.RawExtension	`json:"serviceBindingCreateResponseSchema,omitempty"`
	DefaultProvisionParameters		*runtime.RawExtension	`json:"defaultProvisionParameters,omitempty"`
}
type ClusterServicePlanSpec struct {
	CommonServicePlanSpec		`json:",inline"`
	ClusterServiceBrokerName	string			`json:"clusterServiceBrokerName"`
	ClusterServiceClassRef		ClusterObjectReference	`json:"clusterServiceClassRef"`
}
type ClusterServicePlanStatus struct {
	CommonServicePlanStatus `json:",inline"`
}
type CommonServicePlanStatus struct {
	RemovedFromBrokerCatalog bool `json:"removedFromBrokerCatalog"`
}
type ServicePlanList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ServicePlan	`json:"items"`
}
type ServicePlan struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ServicePlanSpec		`json:"spec,omitempty"`
	Status			ServicePlanStatus	`json:"status,omitempty"`
}
type ServicePlanSpec struct {
	CommonServicePlanSpec	`json:",inline"`
	ServiceBrokerName	string			`json:"serviceBrokerName"`
	ServiceClassRef		LocalObjectReference	`json:"serviceClassRef"`
}
type ServicePlanStatus struct {
	CommonServicePlanStatus `json:",inline"`
}
type ServiceInstanceList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ServiceInstance	`json:"items"`
}
type UserInfo struct {
	Username	string			`json:"username"`
	UID		string			`json:"uid"`
	Groups		[]string		`json:"groups,omitempty"`
	Extra		map[string]ExtraValue	`json:"extra,omitempty"`
}
type ExtraValue []string
type ServiceInstance struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ServiceInstanceSpec	`json:"spec,omitempty"`
	Status			ServiceInstanceStatus	`json:"status,omitempty"`
}
type PlanReference struct {
	ClusterServiceClassExternalName	string	`json:"clusterServiceClassExternalName,omitempty"`
	ClusterServicePlanExternalName	string	`json:"clusterServicePlanExternalName,omitempty"`
	ClusterServiceClassExternalID	string	`json:"clusterServiceClassExternalID,omitempty"`
	ClusterServicePlanExternalID	string	`json:"clusterServicePlanExternalID,omitempty"`
	ClusterServiceClassName		string	`json:"clusterServiceClassName,omitempty"`
	ClusterServicePlanName		string	`json:"clusterServicePlanName,omitempty"`
	ServiceClassExternalName	string	`json:"serviceClassExternalName,omitempty"`
	ServicePlanExternalName		string	`json:"servicePlanExternalName,omitempty"`
	ServiceClassExternalID		string	`json:"serviceClassExternalID,omitempty"`
	ServicePlanExternalID		string	`json:"servicePlanExternalID,omitempty"`
	ServiceClassName		string	`json:"serviceClassName,omitempty"`
	ServicePlanName			string	`json:"servicePlanName,omitempty"`
}
type ServiceInstanceSpec struct {
	PlanReference		`json:",inline"`
	ClusterServiceClassRef	*ClusterObjectReference	`json:"clusterServiceClassRef,omitempty"`
	ClusterServicePlanRef	*ClusterObjectReference	`json:"clusterServicePlanRef,omitempty"`
	ServiceClassRef		*LocalObjectReference	`json:"serviceClassRef,omitempty"`
	ServicePlanRef		*LocalObjectReference	`json:"servicePlanRef,omitempty"`
	Parameters		*runtime.RawExtension	`json:"parameters,omitempty"`
	ParametersFrom		[]ParametersFromSource	`json:"parametersFrom,omitempty"`
	ExternalID		string			`json:"externalID"`
	UserInfo		*UserInfo		`json:"userInfo,omitempty"`
	UpdateRequests		int64			`json:"updateRequests"`
}
type ServiceInstanceStatus struct {
	Conditions			[]ServiceInstanceCondition		`json:"conditions"`
	AsyncOpInProgress		bool					`json:"asyncOpInProgress"`
	OrphanMitigationInProgress	bool					`json:"orphanMitigationInProgress"`
	LastOperation			*string					`json:"lastOperation,omitempty"`
	DashboardURL			*string					`json:"dashboardURL,omitempty"`
	CurrentOperation		ServiceInstanceOperation		`json:"currentOperation,omitempty"`
	ReconciledGeneration		int64					`json:"reconciledGeneration"`
	ObservedGeneration		int64					`json:"observedGeneration"`
	OperationStartTime		*metav1.Time				`json:"operationStartTime,omitempty"`
	InProgressProperties		*ServiceInstancePropertiesState		`json:"inProgressProperties,omitempty"`
	ExternalProperties		*ServiceInstancePropertiesState		`json:"externalProperties,omitempty"`
	ProvisionStatus			ServiceInstanceProvisionStatus		`json:"provisionStatus"`
	DeprovisionStatus		ServiceInstanceDeprovisionStatus	`json:"deprovisionStatus"`
	DefaultProvisionParameters	*runtime.RawExtension			`json:"defaultProvisionParameters,omitempty"`
}
type ServiceInstanceCondition struct {
	Type			ServiceInstanceConditionType	`json:"type"`
	Status			ConditionStatus			`json:"status"`
	LastTransitionTime	metav1.Time			`json:"lastTransitionTime"`
	Reason			string				`json:"reason"`
	Message			string				`json:"message"`
}
type ServiceInstanceConditionType string

const (
	ServiceInstanceConditionReady			ServiceInstanceConditionType	= "Ready"
	ServiceInstanceConditionFailed			ServiceInstanceConditionType	= "Failed"
	ServiceInstanceConditionOrphanMitigation	ServiceInstanceConditionType	= "OrphanMitigation"
)

type ServiceInstanceOperation string

const (
	ServiceInstanceOperationProvision	ServiceInstanceOperation	= "Provision"
	ServiceInstanceOperationUpdate		ServiceInstanceOperation	= "Update"
	ServiceInstanceOperationDeprovision	ServiceInstanceOperation	= "Deprovision"
)

type ServiceInstancePropertiesState struct {
	ClusterServicePlanExternalName	string			`json:"clusterServicePlanExternalName"`
	ClusterServicePlanExternalID	string			`json:"clusterServicePlanExternalID"`
	ServicePlanExternalName		string			`json:"servicePlanExternalName,omitempty"`
	ServicePlanExternalID		string			`json:"servicePlanExternalID,omitempty"`
	Parameters			*runtime.RawExtension	`json:"parameters,omitempty"`
	ParameterChecksum		string			`json:"parameterChecksum,omitempty"`
	UserInfo			*UserInfo		`json:"userInfo,omitempty"`
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
	ServiceInstanceProvisionStatusProvisioned	ServiceInstanceProvisionStatus	= "Provisioned"
	ServiceInstanceProvisionStatusNotProvisioned	ServiceInstanceProvisionStatus	= "NotProvisioned"
)

type ServiceBindingList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]ServiceBinding	`json:"items"`
}
type ServiceBinding struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			ServiceBindingSpec	`json:"spec,omitempty"`
	Status			ServiceBindingStatus	`json:"status,omitempty"`
}
type ServiceBindingSpec struct {
	InstanceRef		LocalObjectReference	`json:"instanceRef"`
	Parameters		*runtime.RawExtension	`json:"parameters,omitempty"`
	ParametersFrom		[]ParametersFromSource	`json:"parametersFrom,omitempty"`
	SecretName		string			`json:"secretName,omitempty"`
	SecretTransforms	[]SecretTransform	`json:"secretTransforms,omitempty"`
	ExternalID		string			`json:"externalID"`
	UserInfo		*UserInfo		`json:"userInfo,omitempty"`
}
type ServiceBindingStatus struct {
	Conditions			[]ServiceBindingCondition	`json:"conditions"`
	AsyncOpInProgress		bool				`json:"asyncOpInProgress"`
	LastOperation			*string				`json:"lastOperation,omitempty"`
	CurrentOperation		ServiceBindingOperation		`json:"currentOperation,omitempty"`
	ReconciledGeneration		int64				`json:"reconciledGeneration"`
	OperationStartTime		*metav1.Time			`json:"operationStartTime,omitempty"`
	InProgressProperties		*ServiceBindingPropertiesState	`json:"inProgressProperties,omitempty"`
	ExternalProperties		*ServiceBindingPropertiesState	`json:"externalProperties,omitempty"`
	OrphanMitigationInProgress	bool				`json:"orphanMitigationInProgress"`
	UnbindStatus			ServiceBindingUnbindStatus	`json:"unbindStatus"`
}
type ServiceBindingCondition struct {
	Type			ServiceBindingConditionType	`json:"type"`
	Status			ConditionStatus			`json:"status"`
	LastTransitionTime	metav1.Time			`json:"lastTransitionTime"`
	Reason			string				`json:"reason"`
	Message			string				`json:"message"`
}
type ServiceBindingConditionType string

const (
	ServiceBindingConditionReady	ServiceBindingConditionType	= "Ready"
	ServiceBindingConditionFailed	ServiceBindingConditionType	= "Failed"
)

type ServiceBindingOperation string

const (
	ServiceBindingOperationBind	ServiceBindingOperation	= "Bind"
	ServiceBindingOperationUnbind	ServiceBindingOperation	= "Unbind"
)

type ServiceBindingUnbindStatus string

const (
	ServiceBindingUnbindStatusNotRequired	ServiceBindingUnbindStatus	= "NotRequired"
	ServiceBindingUnbindStatusRequired	ServiceBindingUnbindStatus	= "Required"
	ServiceBindingUnbindStatusSucceeded	ServiceBindingUnbindStatus	= "Succeeded"
	ServiceBindingUnbindStatusFailed	ServiceBindingUnbindStatus	= "Failed"
)
const (
	FinalizerServiceCatalog string = "kubernetes-incubator/service-catalog"
)

type ServiceBindingPropertiesState struct {
	Parameters		*runtime.RawExtension	`json:"parameters,omitempty"`
	ParameterChecksum	string			`json:"parameterChecksum,omitempty"`
	UserInfo		*UserInfo		`json:"userInfo,omitempty"`
}
type ParametersFromSource struct {
	SecretKeyRef *SecretKeyReference `json:"secretKeyRef,omitempty"`
}
type SecretKeyReference struct {
	Name	string	`json:"name"`
	Key	string	`json:"key"`
}
type ObjectReference struct {
	Namespace	string	`json:"namespace,omitempty"`
	Name		string	`json:"name,omitempty"`
}
type LocalObjectReference struct {
	Name string `json:"name,omitempty"`
}
type ClusterObjectReference struct {
	Name string `json:"name,omitempty"`
}

const (
	FilterName				= "name"
	FilterSpecExternalName			= "spec.externalName"
	FilterSpecExternalID			= "spec.externalID"
	FilterSpecServiceBrokerName		= "spec.serviceBrokerName"
	FilterSpecClusterServiceClassName	= "spec.clusterServiceClass.name"
	FilterSpecServiceClassName		= "spec.serviceClass.name"
	FilterSpecFree				= "spec.free"
)

type SecretTransform struct {
	RenameKey	*RenameKeyTransform	`json:"renameKey,omitempty"`
	AddKey		*AddKeyTransform	`json:"addKey,omitempty"`
	AddKeysFrom	*AddKeysFromTransform	`json:"addKeysFrom,omitempty"`
	RemoveKey	*RemoveKeyTransform	`json:"removeKey,omitempty"`
}
type RenameKeyTransform struct {
	From	string	`json:"from"`
	To	string	`json:"to"`
}
type AddKeyTransform struct {
	Key			string	`json:"key"`
	Value			[]byte	`json:"value"`
	StringValue		*string	`json:"stringValue"`
	JSONPathExpression	*string	`json:"jsonPathExpression"`
}
type AddKeysFromTransform struct {
	SecretRef *ObjectReference `json:"secretRef,omitempty"`
}
type RemoveKeyTransform struct {
	Key string `json:"key"`
}
