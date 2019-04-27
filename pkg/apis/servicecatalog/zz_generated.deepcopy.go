package servicecatalog

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *AddKeyTransform) DeepCopyInto(out *AddKeyTransform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.StringValue != nil {
		in, out := &in.StringValue, &out.StringValue
		*out = new(string)
		**out = **in
	}
	if in.JSONPathExpression != nil {
		in, out := &in.JSONPathExpression, &out.JSONPathExpression
		*out = new(string)
		**out = **in
	}
	return
}
func (in *AddKeyTransform) DeepCopy() *AddKeyTransform {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AddKeyTransform)
	in.DeepCopyInto(out)
	return out
}
func (in *AddKeysFromTransform) DeepCopyInto(out *AddKeysFromTransform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectReference)
		**out = **in
	}
	return
}
func (in *AddKeysFromTransform) DeepCopy() *AddKeysFromTransform {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(AddKeysFromTransform)
	in.DeepCopyInto(out)
	return out
}
func (in *BasicAuthConfig) DeepCopyInto(out *BasicAuthConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	return
}
func (in *BasicAuthConfig) DeepCopy() *BasicAuthConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BasicAuthConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *BearerTokenAuthConfig) DeepCopyInto(out *BearerTokenAuthConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	return
}
func (in *BearerTokenAuthConfig) DeepCopy() *BearerTokenAuthConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(BearerTokenAuthConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *CatalogRestrictions) DeepCopyInto(out *CatalogRestrictions) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ServiceClass != nil {
		in, out := &in.ServiceClass, &out.ServiceClass
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ServicePlan != nil {
		in, out := &in.ServicePlan, &out.ServicePlan
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *CatalogRestrictions) DeepCopy() *CatalogRestrictions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CatalogRestrictions)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterBasicAuthConfig) DeepCopyInto(out *ClusterBasicAuthConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectReference)
		**out = **in
	}
	return
}
func (in *ClusterBasicAuthConfig) DeepCopy() *ClusterBasicAuthConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterBasicAuthConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterBearerTokenAuthConfig) DeepCopyInto(out *ClusterBearerTokenAuthConfig) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(ObjectReference)
		**out = **in
	}
	return
}
func (in *ClusterBearerTokenAuthConfig) DeepCopy() *ClusterBearerTokenAuthConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterBearerTokenAuthConfig)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterObjectReference) DeepCopyInto(out *ClusterObjectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ClusterObjectReference) DeepCopy() *ClusterObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterObjectReference)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceBroker) DeepCopyInto(out *ClusterServiceBroker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ClusterServiceBroker) DeepCopy() *ClusterServiceBroker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBroker)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceBroker) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterServiceBrokerAuthInfo) DeepCopyInto(out *ClusterServiceBrokerAuthInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		*out = new(ClusterBasicAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Bearer != nil {
		in, out := &in.Bearer, &out.Bearer
		*out = new(ClusterBearerTokenAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ClusterServiceBrokerAuthInfo) DeepCopy() *ClusterServiceBrokerAuthInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBrokerAuthInfo)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceBrokerList) DeepCopyInto(out *ClusterServiceBrokerList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServiceBroker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ClusterServiceBrokerList) DeepCopy() *ClusterServiceBrokerList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBrokerList)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceBrokerList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterServiceBrokerSpec) DeepCopyInto(out *ClusterServiceBrokerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServiceBrokerSpec.DeepCopyInto(&out.CommonServiceBrokerSpec)
	if in.AuthInfo != nil {
		in, out := &in.AuthInfo, &out.AuthInfo
		*out = new(ClusterServiceBrokerAuthInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ClusterServiceBrokerSpec) DeepCopy() *ClusterServiceBrokerSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBrokerSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceBrokerStatus) DeepCopyInto(out *ClusterServiceBrokerStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServiceBrokerStatus.DeepCopyInto(&out.CommonServiceBrokerStatus)
	return
}
func (in *ClusterServiceBrokerStatus) DeepCopy() *ClusterServiceBrokerStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceBrokerStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceClass) DeepCopyInto(out *ClusterServiceClass) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}
func (in *ClusterServiceClass) DeepCopy() *ClusterServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClass)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceClass) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterServiceClassList) DeepCopyInto(out *ClusterServiceClassList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServiceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ClusterServiceClassList) DeepCopy() *ClusterServiceClassList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClassList)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceClassList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterServiceClassSpec) DeepCopyInto(out *ClusterServiceClassSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServiceClassSpec.DeepCopyInto(&out.CommonServiceClassSpec)
	return
}
func (in *ClusterServiceClassSpec) DeepCopy() *ClusterServiceClassSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClassSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServiceClassStatus) DeepCopyInto(out *ClusterServiceClassStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.CommonServiceClassStatus = in.CommonServiceClassStatus
	return
}
func (in *ClusterServiceClassStatus) DeepCopy() *ClusterServiceClassStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServiceClassStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServicePlan) DeepCopyInto(out *ClusterServicePlan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}
func (in *ClusterServicePlan) DeepCopy() *ClusterServicePlan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlan)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServicePlan) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterServicePlanList) DeepCopyInto(out *ClusterServicePlanList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterServicePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ClusterServicePlanList) DeepCopy() *ClusterServicePlanList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlanList)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServicePlanList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ClusterServicePlanSpec) DeepCopyInto(out *ClusterServicePlanSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServicePlanSpec.DeepCopyInto(&out.CommonServicePlanSpec)
	out.ClusterServiceClassRef = in.ClusterServiceClassRef
	return
}
func (in *ClusterServicePlanSpec) DeepCopy() *ClusterServicePlanSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlanSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ClusterServicePlanStatus) DeepCopyInto(out *ClusterServicePlanStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.CommonServicePlanStatus = in.CommonServicePlanStatus
	return
}
func (in *ClusterServicePlanStatus) DeepCopy() *ClusterServicePlanStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ClusterServicePlanStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonServiceBrokerSpec) DeepCopyInto(out *CommonServiceBrokerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.RelistDuration != nil {
		in, out := &in.RelistDuration, &out.RelistDuration
		*out = new(v1.Duration)
		**out = **in
	}
	if in.CatalogRestrictions != nil {
		in, out := &in.CatalogRestrictions, &out.CatalogRestrictions
		*out = new(CatalogRestrictions)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *CommonServiceBrokerSpec) DeepCopy() *CommonServiceBrokerSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonServiceBrokerSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonServiceBrokerStatus) DeepCopyInto(out *CommonServiceBrokerStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceBrokerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		*out = (*in).DeepCopy()
	}
	if in.LastCatalogRetrievalTime != nil {
		in, out := &in.LastCatalogRetrievalTime, &out.LastCatalogRetrievalTime
		*out = (*in).DeepCopy()
	}
	return
}
func (in *CommonServiceBrokerStatus) DeepCopy() *CommonServiceBrokerStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonServiceBrokerStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonServiceClassSpec) DeepCopyInto(out *CommonServiceClassSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.ExternalMetadata != nil {
		in, out := &in.ExternalMetadata, &out.ExternalMetadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultProvisionParameters != nil {
		in, out := &in.DefaultProvisionParameters, &out.DefaultProvisionParameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *CommonServiceClassSpec) DeepCopy() *CommonServiceClassSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonServiceClassSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonServiceClassStatus) DeepCopyInto(out *CommonServiceClassStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *CommonServiceClassStatus) DeepCopy() *CommonServiceClassStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonServiceClassStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonServicePlanSpec) DeepCopyInto(out *CommonServicePlanSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Bindable != nil {
		in, out := &in.Bindable, &out.Bindable
		*out = new(bool)
		**out = **in
	}
	if in.ExternalMetadata != nil {
		in, out := &in.ExternalMetadata, &out.ExternalMetadata
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceCreateParameterSchema != nil {
		in, out := &in.InstanceCreateParameterSchema, &out.InstanceCreateParameterSchema
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceUpdateParameterSchema != nil {
		in, out := &in.InstanceUpdateParameterSchema, &out.InstanceUpdateParameterSchema
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceBindingCreateParameterSchema != nil {
		in, out := &in.ServiceBindingCreateParameterSchema, &out.ServiceBindingCreateParameterSchema
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceBindingCreateResponseSchema != nil {
		in, out := &in.ServiceBindingCreateResponseSchema, &out.ServiceBindingCreateResponseSchema
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultProvisionParameters != nil {
		in, out := &in.DefaultProvisionParameters, &out.DefaultProvisionParameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *CommonServicePlanSpec) DeepCopy() *CommonServicePlanSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonServicePlanSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *CommonServicePlanStatus) DeepCopyInto(out *CommonServicePlanStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *CommonServicePlanStatus) DeepCopy() *CommonServicePlanStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(CommonServicePlanStatus)
	in.DeepCopyInto(out)
	return out
}
func (in ExtraValue) DeepCopyInto(out *ExtraValue) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	{
		in := &in
		*out = make(ExtraValue, len(*in))
		copy(*out, *in)
		return
	}
}
func (in ExtraValue) DeepCopy() ExtraValue {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ExtraValue)
	in.DeepCopyInto(out)
	return *out
}
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *ObjectReference) DeepCopy() *ObjectReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}
func (in *ParametersFromSource) DeepCopyInto(out *ParametersFromSource) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(SecretKeyReference)
		**out = **in
	}
	return
}
func (in *ParametersFromSource) DeepCopy() *ParametersFromSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ParametersFromSource)
	in.DeepCopyInto(out)
	return out
}
func (in *PlanReference) DeepCopyInto(out *PlanReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *PlanReference) DeepCopy() *PlanReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(PlanReference)
	in.DeepCopyInto(out)
	return out
}
func (in *RemoveKeyTransform) DeepCopyInto(out *RemoveKeyTransform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *RemoveKeyTransform) DeepCopy() *RemoveKeyTransform {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RemoveKeyTransform)
	in.DeepCopyInto(out)
	return out
}
func (in *RenameKeyTransform) DeepCopyInto(out *RenameKeyTransform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *RenameKeyTransform) DeepCopy() *RenameKeyTransform {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RenameKeyTransform)
	in.DeepCopyInto(out)
	return out
}
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}
func (in *SecretTransform) DeepCopyInto(out *SecretTransform) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.RenameKey != nil {
		in, out := &in.RenameKey, &out.RenameKey
		*out = new(RenameKeyTransform)
		**out = **in
	}
	if in.AddKey != nil {
		in, out := &in.AddKey, &out.AddKey
		*out = new(AddKeyTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.AddKeysFrom != nil {
		in, out := &in.AddKeysFrom, &out.AddKeysFrom
		*out = new(AddKeysFromTransform)
		(*in).DeepCopyInto(*out)
	}
	if in.RemoveKey != nil {
		in, out := &in.RemoveKey, &out.RemoveKey
		*out = new(RemoveKeyTransform)
		**out = **in
	}
	return
}
func (in *SecretTransform) DeepCopy() *SecretTransform {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecretTransform)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBinding) DeepCopyInto(out *ServiceBinding) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ServiceBinding) DeepCopy() *ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBinding)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBinding) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceBindingCondition) DeepCopyInto(out *ServiceBindingCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *ServiceBindingCondition) DeepCopy() *ServiceBindingCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBindingCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBindingList) DeepCopyInto(out *ServiceBindingList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ServiceBindingList) DeepCopy() *ServiceBindingList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBindingList)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBindingList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceBindingPropertiesState) DeepCopyInto(out *ServiceBindingPropertiesState) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(UserInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceBindingPropertiesState) DeepCopy() *ServiceBindingPropertiesState {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBindingPropertiesState)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBindingSpec) DeepCopyInto(out *ServiceBindingSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.InstanceRef = in.InstanceRef
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.ParametersFrom != nil {
		in, out := &in.ParametersFrom, &out.ParametersFrom
		*out = make([]ParametersFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretTransforms != nil {
		in, out := &in.SecretTransforms, &out.SecretTransforms
		*out = make([]SecretTransform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(UserInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceBindingSpec) DeepCopy() *ServiceBindingSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBindingSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBindingStatus) DeepCopyInto(out *ServiceBindingStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceBindingCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(string)
		**out = **in
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		*out = (*in).DeepCopy()
	}
	if in.InProgressProperties != nil {
		in, out := &in.InProgressProperties, &out.InProgressProperties
		*out = new(ServiceBindingPropertiesState)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalProperties != nil {
		in, out := &in.ExternalProperties, &out.ExternalProperties
		*out = new(ServiceBindingPropertiesState)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceBindingStatus) DeepCopy() *ServiceBindingStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBindingStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBroker) DeepCopyInto(out *ServiceBroker) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ServiceBroker) DeepCopy() *ServiceBroker {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBroker)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBroker) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceBrokerAuthInfo) DeepCopyInto(out *ServiceBrokerAuthInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		*out = new(BasicAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Bearer != nil {
		in, out := &in.Bearer, &out.Bearer
		*out = new(BearerTokenAuthConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceBrokerAuthInfo) DeepCopy() *ServiceBrokerAuthInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerAuthInfo)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBrokerCondition) DeepCopyInto(out *ServiceBrokerCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *ServiceBrokerCondition) DeepCopy() *ServiceBrokerCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBrokerList) DeepCopyInto(out *ServiceBrokerList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceBroker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ServiceBrokerList) DeepCopy() *ServiceBrokerList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerList)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBrokerList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceBrokerSpec) DeepCopyInto(out *ServiceBrokerSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServiceBrokerSpec.DeepCopyInto(&out.CommonServiceBrokerSpec)
	if in.AuthInfo != nil {
		in, out := &in.AuthInfo, &out.AuthInfo
		*out = new(ServiceBrokerAuthInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceBrokerSpec) DeepCopy() *ServiceBrokerSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceBrokerStatus) DeepCopyInto(out *ServiceBrokerStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServiceBrokerStatus.DeepCopyInto(&out.CommonServiceBrokerStatus)
	return
}
func (in *ServiceBrokerStatus) DeepCopy() *ServiceBrokerStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceBrokerStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceClass) DeepCopyInto(out *ServiceClass) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}
func (in *ServiceClass) DeepCopy() *ServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceClass)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceClass) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceClassList) DeepCopyInto(out *ServiceClassList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ServiceClassList) DeepCopy() *ServiceClassList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceClassList)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceClassList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceClassSpec) DeepCopyInto(out *ServiceClassSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServiceClassSpec.DeepCopyInto(&out.CommonServiceClassSpec)
	return
}
func (in *ServiceClassSpec) DeepCopy() *ServiceClassSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceClassSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceClassStatus) DeepCopyInto(out *ServiceClassStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.CommonServiceClassStatus = in.CommonServiceClassStatus
	return
}
func (in *ServiceClassStatus) DeepCopy() *ServiceClassStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceClassStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceInstance) DeepCopyInto(out *ServiceInstance) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}
func (in *ServiceInstance) DeepCopy() *ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceInstance)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceInstance) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceInstanceCondition) DeepCopyInto(out *ServiceInstanceCondition) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}
func (in *ServiceInstanceCondition) DeepCopy() *ServiceInstanceCondition {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceCondition)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceInstanceList) DeepCopyInto(out *ServiceInstanceList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ServiceInstanceList) DeepCopy() *ServiceInstanceList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceList)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceInstanceList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServiceInstancePropertiesState) DeepCopyInto(out *ServiceInstancePropertiesState) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(UserInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceInstancePropertiesState) DeepCopy() *ServiceInstancePropertiesState {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceInstancePropertiesState)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceInstanceSpec) DeepCopyInto(out *ServiceInstanceSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.PlanReference = in.PlanReference
	if in.ClusterServiceClassRef != nil {
		in, out := &in.ClusterServiceClassRef, &out.ClusterServiceClassRef
		*out = new(ClusterObjectReference)
		**out = **in
	}
	if in.ClusterServicePlanRef != nil {
		in, out := &in.ClusterServicePlanRef, &out.ClusterServicePlanRef
		*out = new(ClusterObjectReference)
		**out = **in
	}
	if in.ServiceClassRef != nil {
		in, out := &in.ServiceClassRef, &out.ServiceClassRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.ServicePlanRef != nil {
		in, out := &in.ServicePlanRef, &out.ServicePlanRef
		*out = new(LocalObjectReference)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.ParametersFrom != nil {
		in, out := &in.ParametersFrom, &out.ParametersFrom
		*out = make([]ParametersFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserInfo != nil {
		in, out := &in.UserInfo, &out.UserInfo
		*out = new(UserInfo)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceInstanceSpec) DeepCopy() *ServiceInstanceSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ServiceInstanceStatus) DeepCopyInto(out *ServiceInstanceStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ServiceInstanceCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(string)
		**out = **in
	}
	if in.DashboardURL != nil {
		in, out := &in.DashboardURL, &out.DashboardURL
		*out = new(string)
		**out = **in
	}
	if in.OperationStartTime != nil {
		in, out := &in.OperationStartTime, &out.OperationStartTime
		*out = (*in).DeepCopy()
	}
	if in.InProgressProperties != nil {
		in, out := &in.InProgressProperties, &out.InProgressProperties
		*out = new(ServiceInstancePropertiesState)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalProperties != nil {
		in, out := &in.ExternalProperties, &out.ExternalProperties
		*out = new(ServiceInstancePropertiesState)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultProvisionParameters != nil {
		in, out := &in.DefaultProvisionParameters, &out.DefaultProvisionParameters
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *ServiceInstanceStatus) DeepCopy() *ServiceInstanceStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServiceInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *ServicePlan) DeepCopyInto(out *ServicePlan) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}
func (in *ServicePlan) DeepCopy() *ServicePlan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServicePlan)
	in.DeepCopyInto(out)
	return out
}
func (in *ServicePlan) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServicePlanList) DeepCopyInto(out *ServicePlanList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServicePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *ServicePlanList) DeepCopy() *ServicePlanList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServicePlanList)
	in.DeepCopyInto(out)
	return out
}
func (in *ServicePlanList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *ServicePlanSpec) DeepCopyInto(out *ServicePlanSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.CommonServicePlanSpec.DeepCopyInto(&out.CommonServicePlanSpec)
	out.ServiceClassRef = in.ServiceClassRef
	return
}
func (in *ServicePlanSpec) DeepCopy() *ServicePlanSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServicePlanSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *ServicePlanStatus) DeepCopyInto(out *ServicePlanStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.CommonServicePlanStatus = in.CommonServicePlanStatus
	return
}
func (in *ServicePlanStatus) DeepCopy() *ServicePlanStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(ServicePlanStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *UserInfo) DeepCopyInto(out *UserInfo) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ExtraValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	return
}
func (in *UserInfo) DeepCopy() *UserInfo {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(UserInfo)
	in.DeepCopyInto(out)
	return out
}
