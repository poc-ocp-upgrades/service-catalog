package v1beta1

import (
	unsafe "unsafe"
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*AddKeyTransform)(nil), (*servicecatalog.AddKeyTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddKeyTransform_To_servicecatalog_AddKeyTransform(a.(*AddKeyTransform), b.(*servicecatalog.AddKeyTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.AddKeyTransform)(nil), (*AddKeyTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_AddKeyTransform_To_v1beta1_AddKeyTransform(a.(*servicecatalog.AddKeyTransform), b.(*AddKeyTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AddKeysFromTransform)(nil), (*servicecatalog.AddKeysFromTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_AddKeysFromTransform_To_servicecatalog_AddKeysFromTransform(a.(*AddKeysFromTransform), b.(*servicecatalog.AddKeysFromTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.AddKeysFromTransform)(nil), (*AddKeysFromTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_AddKeysFromTransform_To_v1beta1_AddKeysFromTransform(a.(*servicecatalog.AddKeysFromTransform), b.(*AddKeysFromTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BasicAuthConfig)(nil), (*servicecatalog.BasicAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_BasicAuthConfig_To_servicecatalog_BasicAuthConfig(a.(*BasicAuthConfig), b.(*servicecatalog.BasicAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.BasicAuthConfig)(nil), (*BasicAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_BasicAuthConfig_To_v1beta1_BasicAuthConfig(a.(*servicecatalog.BasicAuthConfig), b.(*BasicAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*BearerTokenAuthConfig)(nil), (*servicecatalog.BearerTokenAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_BearerTokenAuthConfig_To_servicecatalog_BearerTokenAuthConfig(a.(*BearerTokenAuthConfig), b.(*servicecatalog.BearerTokenAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.BearerTokenAuthConfig)(nil), (*BearerTokenAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_BearerTokenAuthConfig_To_v1beta1_BearerTokenAuthConfig(a.(*servicecatalog.BearerTokenAuthConfig), b.(*BearerTokenAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CatalogRestrictions)(nil), (*servicecatalog.CatalogRestrictions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CatalogRestrictions_To_servicecatalog_CatalogRestrictions(a.(*CatalogRestrictions), b.(*servicecatalog.CatalogRestrictions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CatalogRestrictions)(nil), (*CatalogRestrictions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CatalogRestrictions_To_v1beta1_CatalogRestrictions(a.(*servicecatalog.CatalogRestrictions), b.(*CatalogRestrictions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterBasicAuthConfig)(nil), (*servicecatalog.ClusterBasicAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterBasicAuthConfig_To_servicecatalog_ClusterBasicAuthConfig(a.(*ClusterBasicAuthConfig), b.(*servicecatalog.ClusterBasicAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterBasicAuthConfig)(nil), (*ClusterBasicAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterBasicAuthConfig_To_v1beta1_ClusterBasicAuthConfig(a.(*servicecatalog.ClusterBasicAuthConfig), b.(*ClusterBasicAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterBearerTokenAuthConfig)(nil), (*servicecatalog.ClusterBearerTokenAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterBearerTokenAuthConfig_To_servicecatalog_ClusterBearerTokenAuthConfig(a.(*ClusterBearerTokenAuthConfig), b.(*servicecatalog.ClusterBearerTokenAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterBearerTokenAuthConfig)(nil), (*ClusterBearerTokenAuthConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterBearerTokenAuthConfig_To_v1beta1_ClusterBearerTokenAuthConfig(a.(*servicecatalog.ClusterBearerTokenAuthConfig), b.(*ClusterBearerTokenAuthConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterObjectReference)(nil), (*servicecatalog.ClusterObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterObjectReference_To_servicecatalog_ClusterObjectReference(a.(*ClusterObjectReference), b.(*servicecatalog.ClusterObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterObjectReference)(nil), (*ClusterObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterObjectReference_To_v1beta1_ClusterObjectReference(a.(*servicecatalog.ClusterObjectReference), b.(*ClusterObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceBroker)(nil), (*servicecatalog.ClusterServiceBroker)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceBroker_To_servicecatalog_ClusterServiceBroker(a.(*ClusterServiceBroker), b.(*servicecatalog.ClusterServiceBroker), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceBroker)(nil), (*ClusterServiceBroker)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceBroker_To_v1beta1_ClusterServiceBroker(a.(*servicecatalog.ClusterServiceBroker), b.(*ClusterServiceBroker), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceBrokerAuthInfo)(nil), (*servicecatalog.ClusterServiceBrokerAuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceBrokerAuthInfo_To_servicecatalog_ClusterServiceBrokerAuthInfo(a.(*ClusterServiceBrokerAuthInfo), b.(*servicecatalog.ClusterServiceBrokerAuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceBrokerAuthInfo)(nil), (*ClusterServiceBrokerAuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceBrokerAuthInfo_To_v1beta1_ClusterServiceBrokerAuthInfo(a.(*servicecatalog.ClusterServiceBrokerAuthInfo), b.(*ClusterServiceBrokerAuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceBrokerList)(nil), (*servicecatalog.ClusterServiceBrokerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceBrokerList_To_servicecatalog_ClusterServiceBrokerList(a.(*ClusterServiceBrokerList), b.(*servicecatalog.ClusterServiceBrokerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceBrokerList)(nil), (*ClusterServiceBrokerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceBrokerList_To_v1beta1_ClusterServiceBrokerList(a.(*servicecatalog.ClusterServiceBrokerList), b.(*ClusterServiceBrokerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceBrokerSpec)(nil), (*servicecatalog.ClusterServiceBrokerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceBrokerSpec_To_servicecatalog_ClusterServiceBrokerSpec(a.(*ClusterServiceBrokerSpec), b.(*servicecatalog.ClusterServiceBrokerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceBrokerSpec)(nil), (*ClusterServiceBrokerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceBrokerSpec_To_v1beta1_ClusterServiceBrokerSpec(a.(*servicecatalog.ClusterServiceBrokerSpec), b.(*ClusterServiceBrokerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceBrokerStatus)(nil), (*servicecatalog.ClusterServiceBrokerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceBrokerStatus_To_servicecatalog_ClusterServiceBrokerStatus(a.(*ClusterServiceBrokerStatus), b.(*servicecatalog.ClusterServiceBrokerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceBrokerStatus)(nil), (*ClusterServiceBrokerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceBrokerStatus_To_v1beta1_ClusterServiceBrokerStatus(a.(*servicecatalog.ClusterServiceBrokerStatus), b.(*ClusterServiceBrokerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceClass)(nil), (*servicecatalog.ClusterServiceClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceClass_To_servicecatalog_ClusterServiceClass(a.(*ClusterServiceClass), b.(*servicecatalog.ClusterServiceClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceClass)(nil), (*ClusterServiceClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceClass_To_v1beta1_ClusterServiceClass(a.(*servicecatalog.ClusterServiceClass), b.(*ClusterServiceClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceClassList)(nil), (*servicecatalog.ClusterServiceClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceClassList_To_servicecatalog_ClusterServiceClassList(a.(*ClusterServiceClassList), b.(*servicecatalog.ClusterServiceClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceClassList)(nil), (*ClusterServiceClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceClassList_To_v1beta1_ClusterServiceClassList(a.(*servicecatalog.ClusterServiceClassList), b.(*ClusterServiceClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceClassSpec)(nil), (*servicecatalog.ClusterServiceClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceClassSpec_To_servicecatalog_ClusterServiceClassSpec(a.(*ClusterServiceClassSpec), b.(*servicecatalog.ClusterServiceClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceClassSpec)(nil), (*ClusterServiceClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceClassSpec_To_v1beta1_ClusterServiceClassSpec(a.(*servicecatalog.ClusterServiceClassSpec), b.(*ClusterServiceClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServiceClassStatus)(nil), (*servicecatalog.ClusterServiceClassStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServiceClassStatus_To_servicecatalog_ClusterServiceClassStatus(a.(*ClusterServiceClassStatus), b.(*servicecatalog.ClusterServiceClassStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServiceClassStatus)(nil), (*ClusterServiceClassStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServiceClassStatus_To_v1beta1_ClusterServiceClassStatus(a.(*servicecatalog.ClusterServiceClassStatus), b.(*ClusterServiceClassStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServicePlan)(nil), (*servicecatalog.ClusterServicePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServicePlan_To_servicecatalog_ClusterServicePlan(a.(*ClusterServicePlan), b.(*servicecatalog.ClusterServicePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServicePlan)(nil), (*ClusterServicePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServicePlan_To_v1beta1_ClusterServicePlan(a.(*servicecatalog.ClusterServicePlan), b.(*ClusterServicePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServicePlanList)(nil), (*servicecatalog.ClusterServicePlanList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServicePlanList_To_servicecatalog_ClusterServicePlanList(a.(*ClusterServicePlanList), b.(*servicecatalog.ClusterServicePlanList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServicePlanList)(nil), (*ClusterServicePlanList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServicePlanList_To_v1beta1_ClusterServicePlanList(a.(*servicecatalog.ClusterServicePlanList), b.(*ClusterServicePlanList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServicePlanSpec)(nil), (*servicecatalog.ClusterServicePlanSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServicePlanSpec_To_servicecatalog_ClusterServicePlanSpec(a.(*ClusterServicePlanSpec), b.(*servicecatalog.ClusterServicePlanSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServicePlanSpec)(nil), (*ClusterServicePlanSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServicePlanSpec_To_v1beta1_ClusterServicePlanSpec(a.(*servicecatalog.ClusterServicePlanSpec), b.(*ClusterServicePlanSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterServicePlanStatus)(nil), (*servicecatalog.ClusterServicePlanStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ClusterServicePlanStatus_To_servicecatalog_ClusterServicePlanStatus(a.(*ClusterServicePlanStatus), b.(*servicecatalog.ClusterServicePlanStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ClusterServicePlanStatus)(nil), (*ClusterServicePlanStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ClusterServicePlanStatus_To_v1beta1_ClusterServicePlanStatus(a.(*servicecatalog.ClusterServicePlanStatus), b.(*ClusterServicePlanStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CommonServiceBrokerSpec)(nil), (*servicecatalog.CommonServiceBrokerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CommonServiceBrokerSpec_To_servicecatalog_CommonServiceBrokerSpec(a.(*CommonServiceBrokerSpec), b.(*servicecatalog.CommonServiceBrokerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CommonServiceBrokerSpec)(nil), (*CommonServiceBrokerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CommonServiceBrokerSpec_To_v1beta1_CommonServiceBrokerSpec(a.(*servicecatalog.CommonServiceBrokerSpec), b.(*CommonServiceBrokerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CommonServiceBrokerStatus)(nil), (*servicecatalog.CommonServiceBrokerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CommonServiceBrokerStatus_To_servicecatalog_CommonServiceBrokerStatus(a.(*CommonServiceBrokerStatus), b.(*servicecatalog.CommonServiceBrokerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CommonServiceBrokerStatus)(nil), (*CommonServiceBrokerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CommonServiceBrokerStatus_To_v1beta1_CommonServiceBrokerStatus(a.(*servicecatalog.CommonServiceBrokerStatus), b.(*CommonServiceBrokerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CommonServiceClassSpec)(nil), (*servicecatalog.CommonServiceClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CommonServiceClassSpec_To_servicecatalog_CommonServiceClassSpec(a.(*CommonServiceClassSpec), b.(*servicecatalog.CommonServiceClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CommonServiceClassSpec)(nil), (*CommonServiceClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CommonServiceClassSpec_To_v1beta1_CommonServiceClassSpec(a.(*servicecatalog.CommonServiceClassSpec), b.(*CommonServiceClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CommonServiceClassStatus)(nil), (*servicecatalog.CommonServiceClassStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CommonServiceClassStatus_To_servicecatalog_CommonServiceClassStatus(a.(*CommonServiceClassStatus), b.(*servicecatalog.CommonServiceClassStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CommonServiceClassStatus)(nil), (*CommonServiceClassStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CommonServiceClassStatus_To_v1beta1_CommonServiceClassStatus(a.(*servicecatalog.CommonServiceClassStatus), b.(*CommonServiceClassStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CommonServicePlanSpec)(nil), (*servicecatalog.CommonServicePlanSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CommonServicePlanSpec_To_servicecatalog_CommonServicePlanSpec(a.(*CommonServicePlanSpec), b.(*servicecatalog.CommonServicePlanSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CommonServicePlanSpec)(nil), (*CommonServicePlanSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CommonServicePlanSpec_To_v1beta1_CommonServicePlanSpec(a.(*servicecatalog.CommonServicePlanSpec), b.(*CommonServicePlanSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CommonServicePlanStatus)(nil), (*servicecatalog.CommonServicePlanStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CommonServicePlanStatus_To_servicecatalog_CommonServicePlanStatus(a.(*CommonServicePlanStatus), b.(*servicecatalog.CommonServicePlanStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.CommonServicePlanStatus)(nil), (*CommonServicePlanStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_CommonServicePlanStatus_To_v1beta1_CommonServicePlanStatus(a.(*servicecatalog.CommonServicePlanStatus), b.(*CommonServicePlanStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LocalObjectReference)(nil), (*servicecatalog.LocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_LocalObjectReference_To_servicecatalog_LocalObjectReference(a.(*LocalObjectReference), b.(*servicecatalog.LocalObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.LocalObjectReference)(nil), (*LocalObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_LocalObjectReference_To_v1beta1_LocalObjectReference(a.(*servicecatalog.LocalObjectReference), b.(*LocalObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ObjectReference)(nil), (*servicecatalog.ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ObjectReference_To_servicecatalog_ObjectReference(a.(*ObjectReference), b.(*servicecatalog.ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ObjectReference)(nil), (*ObjectReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ObjectReference_To_v1beta1_ObjectReference(a.(*servicecatalog.ObjectReference), b.(*ObjectReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ParametersFromSource)(nil), (*servicecatalog.ParametersFromSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ParametersFromSource_To_servicecatalog_ParametersFromSource(a.(*ParametersFromSource), b.(*servicecatalog.ParametersFromSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ParametersFromSource)(nil), (*ParametersFromSource)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ParametersFromSource_To_v1beta1_ParametersFromSource(a.(*servicecatalog.ParametersFromSource), b.(*ParametersFromSource), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PlanReference)(nil), (*servicecatalog.PlanReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_PlanReference_To_servicecatalog_PlanReference(a.(*PlanReference), b.(*servicecatalog.PlanReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.PlanReference)(nil), (*PlanReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_PlanReference_To_v1beta1_PlanReference(a.(*servicecatalog.PlanReference), b.(*PlanReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RemoveKeyTransform)(nil), (*servicecatalog.RemoveKeyTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RemoveKeyTransform_To_servicecatalog_RemoveKeyTransform(a.(*RemoveKeyTransform), b.(*servicecatalog.RemoveKeyTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.RemoveKeyTransform)(nil), (*RemoveKeyTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_RemoveKeyTransform_To_v1beta1_RemoveKeyTransform(a.(*servicecatalog.RemoveKeyTransform), b.(*RemoveKeyTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RenameKeyTransform)(nil), (*servicecatalog.RenameKeyTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_RenameKeyTransform_To_servicecatalog_RenameKeyTransform(a.(*RenameKeyTransform), b.(*servicecatalog.RenameKeyTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.RenameKeyTransform)(nil), (*RenameKeyTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_RenameKeyTransform_To_v1beta1_RenameKeyTransform(a.(*servicecatalog.RenameKeyTransform), b.(*RenameKeyTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecretKeyReference)(nil), (*servicecatalog.SecretKeyReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SecretKeyReference_To_servicecatalog_SecretKeyReference(a.(*SecretKeyReference), b.(*servicecatalog.SecretKeyReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.SecretKeyReference)(nil), (*SecretKeyReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_SecretKeyReference_To_v1beta1_SecretKeyReference(a.(*servicecatalog.SecretKeyReference), b.(*SecretKeyReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecretTransform)(nil), (*servicecatalog.SecretTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SecretTransform_To_servicecatalog_SecretTransform(a.(*SecretTransform), b.(*servicecatalog.SecretTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.SecretTransform)(nil), (*SecretTransform)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_SecretTransform_To_v1beta1_SecretTransform(a.(*servicecatalog.SecretTransform), b.(*SecretTransform), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBinding)(nil), (*servicecatalog.ServiceBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBinding_To_servicecatalog_ServiceBinding(a.(*ServiceBinding), b.(*servicecatalog.ServiceBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBinding)(nil), (*ServiceBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBinding_To_v1beta1_ServiceBinding(a.(*servicecatalog.ServiceBinding), b.(*ServiceBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBindingCondition)(nil), (*servicecatalog.ServiceBindingCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBindingCondition_To_servicecatalog_ServiceBindingCondition(a.(*ServiceBindingCondition), b.(*servicecatalog.ServiceBindingCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBindingCondition)(nil), (*ServiceBindingCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBindingCondition_To_v1beta1_ServiceBindingCondition(a.(*servicecatalog.ServiceBindingCondition), b.(*ServiceBindingCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBindingList)(nil), (*servicecatalog.ServiceBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBindingList_To_servicecatalog_ServiceBindingList(a.(*ServiceBindingList), b.(*servicecatalog.ServiceBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBindingList)(nil), (*ServiceBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBindingList_To_v1beta1_ServiceBindingList(a.(*servicecatalog.ServiceBindingList), b.(*ServiceBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBindingPropertiesState)(nil), (*servicecatalog.ServiceBindingPropertiesState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBindingPropertiesState_To_servicecatalog_ServiceBindingPropertiesState(a.(*ServiceBindingPropertiesState), b.(*servicecatalog.ServiceBindingPropertiesState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBindingPropertiesState)(nil), (*ServiceBindingPropertiesState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBindingPropertiesState_To_v1beta1_ServiceBindingPropertiesState(a.(*servicecatalog.ServiceBindingPropertiesState), b.(*ServiceBindingPropertiesState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBindingSpec)(nil), (*servicecatalog.ServiceBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBindingSpec_To_servicecatalog_ServiceBindingSpec(a.(*ServiceBindingSpec), b.(*servicecatalog.ServiceBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBindingSpec)(nil), (*ServiceBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBindingSpec_To_v1beta1_ServiceBindingSpec(a.(*servicecatalog.ServiceBindingSpec), b.(*ServiceBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBindingStatus)(nil), (*servicecatalog.ServiceBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBindingStatus_To_servicecatalog_ServiceBindingStatus(a.(*ServiceBindingStatus), b.(*servicecatalog.ServiceBindingStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBindingStatus)(nil), (*ServiceBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBindingStatus_To_v1beta1_ServiceBindingStatus(a.(*servicecatalog.ServiceBindingStatus), b.(*ServiceBindingStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBroker)(nil), (*servicecatalog.ServiceBroker)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBroker_To_servicecatalog_ServiceBroker(a.(*ServiceBroker), b.(*servicecatalog.ServiceBroker), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBroker)(nil), (*ServiceBroker)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBroker_To_v1beta1_ServiceBroker(a.(*servicecatalog.ServiceBroker), b.(*ServiceBroker), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBrokerAuthInfo)(nil), (*servicecatalog.ServiceBrokerAuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBrokerAuthInfo_To_servicecatalog_ServiceBrokerAuthInfo(a.(*ServiceBrokerAuthInfo), b.(*servicecatalog.ServiceBrokerAuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBrokerAuthInfo)(nil), (*ServiceBrokerAuthInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBrokerAuthInfo_To_v1beta1_ServiceBrokerAuthInfo(a.(*servicecatalog.ServiceBrokerAuthInfo), b.(*ServiceBrokerAuthInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBrokerCondition)(nil), (*servicecatalog.ServiceBrokerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBrokerCondition_To_servicecatalog_ServiceBrokerCondition(a.(*ServiceBrokerCondition), b.(*servicecatalog.ServiceBrokerCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBrokerCondition)(nil), (*ServiceBrokerCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBrokerCondition_To_v1beta1_ServiceBrokerCondition(a.(*servicecatalog.ServiceBrokerCondition), b.(*ServiceBrokerCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBrokerList)(nil), (*servicecatalog.ServiceBrokerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBrokerList_To_servicecatalog_ServiceBrokerList(a.(*ServiceBrokerList), b.(*servicecatalog.ServiceBrokerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBrokerList)(nil), (*ServiceBrokerList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBrokerList_To_v1beta1_ServiceBrokerList(a.(*servicecatalog.ServiceBrokerList), b.(*ServiceBrokerList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBrokerSpec)(nil), (*servicecatalog.ServiceBrokerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBrokerSpec_To_servicecatalog_ServiceBrokerSpec(a.(*ServiceBrokerSpec), b.(*servicecatalog.ServiceBrokerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBrokerSpec)(nil), (*ServiceBrokerSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBrokerSpec_To_v1beta1_ServiceBrokerSpec(a.(*servicecatalog.ServiceBrokerSpec), b.(*ServiceBrokerSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceBrokerStatus)(nil), (*servicecatalog.ServiceBrokerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceBrokerStatus_To_servicecatalog_ServiceBrokerStatus(a.(*ServiceBrokerStatus), b.(*servicecatalog.ServiceBrokerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceBrokerStatus)(nil), (*ServiceBrokerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceBrokerStatus_To_v1beta1_ServiceBrokerStatus(a.(*servicecatalog.ServiceBrokerStatus), b.(*ServiceBrokerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceClass)(nil), (*servicecatalog.ServiceClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceClass_To_servicecatalog_ServiceClass(a.(*ServiceClass), b.(*servicecatalog.ServiceClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceClass)(nil), (*ServiceClass)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceClass_To_v1beta1_ServiceClass(a.(*servicecatalog.ServiceClass), b.(*ServiceClass), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceClassList)(nil), (*servicecatalog.ServiceClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceClassList_To_servicecatalog_ServiceClassList(a.(*ServiceClassList), b.(*servicecatalog.ServiceClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceClassList)(nil), (*ServiceClassList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceClassList_To_v1beta1_ServiceClassList(a.(*servicecatalog.ServiceClassList), b.(*ServiceClassList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceClassSpec)(nil), (*servicecatalog.ServiceClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceClassSpec_To_servicecatalog_ServiceClassSpec(a.(*ServiceClassSpec), b.(*servicecatalog.ServiceClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceClassSpec)(nil), (*ServiceClassSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceClassSpec_To_v1beta1_ServiceClassSpec(a.(*servicecatalog.ServiceClassSpec), b.(*ServiceClassSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceClassStatus)(nil), (*servicecatalog.ServiceClassStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceClassStatus_To_servicecatalog_ServiceClassStatus(a.(*ServiceClassStatus), b.(*servicecatalog.ServiceClassStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceClassStatus)(nil), (*ServiceClassStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceClassStatus_To_v1beta1_ServiceClassStatus(a.(*servicecatalog.ServiceClassStatus), b.(*ServiceClassStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceInstance)(nil), (*servicecatalog.ServiceInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceInstance_To_servicecatalog_ServiceInstance(a.(*ServiceInstance), b.(*servicecatalog.ServiceInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceInstance)(nil), (*ServiceInstance)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceInstance_To_v1beta1_ServiceInstance(a.(*servicecatalog.ServiceInstance), b.(*ServiceInstance), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceInstanceCondition)(nil), (*servicecatalog.ServiceInstanceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceInstanceCondition_To_servicecatalog_ServiceInstanceCondition(a.(*ServiceInstanceCondition), b.(*servicecatalog.ServiceInstanceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceInstanceCondition)(nil), (*ServiceInstanceCondition)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceInstanceCondition_To_v1beta1_ServiceInstanceCondition(a.(*servicecatalog.ServiceInstanceCondition), b.(*ServiceInstanceCondition), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceInstanceList)(nil), (*servicecatalog.ServiceInstanceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceInstanceList_To_servicecatalog_ServiceInstanceList(a.(*ServiceInstanceList), b.(*servicecatalog.ServiceInstanceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceInstanceList)(nil), (*ServiceInstanceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceInstanceList_To_v1beta1_ServiceInstanceList(a.(*servicecatalog.ServiceInstanceList), b.(*ServiceInstanceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceInstancePropertiesState)(nil), (*servicecatalog.ServiceInstancePropertiesState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceInstancePropertiesState_To_servicecatalog_ServiceInstancePropertiesState(a.(*ServiceInstancePropertiesState), b.(*servicecatalog.ServiceInstancePropertiesState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceInstancePropertiesState)(nil), (*ServiceInstancePropertiesState)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceInstancePropertiesState_To_v1beta1_ServiceInstancePropertiesState(a.(*servicecatalog.ServiceInstancePropertiesState), b.(*ServiceInstancePropertiesState), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceInstanceSpec)(nil), (*servicecatalog.ServiceInstanceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceInstanceSpec_To_servicecatalog_ServiceInstanceSpec(a.(*ServiceInstanceSpec), b.(*servicecatalog.ServiceInstanceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceInstanceSpec)(nil), (*ServiceInstanceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceInstanceSpec_To_v1beta1_ServiceInstanceSpec(a.(*servicecatalog.ServiceInstanceSpec), b.(*ServiceInstanceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServiceInstanceStatus)(nil), (*servicecatalog.ServiceInstanceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServiceInstanceStatus_To_servicecatalog_ServiceInstanceStatus(a.(*ServiceInstanceStatus), b.(*servicecatalog.ServiceInstanceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServiceInstanceStatus)(nil), (*ServiceInstanceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServiceInstanceStatus_To_v1beta1_ServiceInstanceStatus(a.(*servicecatalog.ServiceInstanceStatus), b.(*ServiceInstanceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServicePlan)(nil), (*servicecatalog.ServicePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServicePlan_To_servicecatalog_ServicePlan(a.(*ServicePlan), b.(*servicecatalog.ServicePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServicePlan)(nil), (*ServicePlan)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServicePlan_To_v1beta1_ServicePlan(a.(*servicecatalog.ServicePlan), b.(*ServicePlan), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServicePlanList)(nil), (*servicecatalog.ServicePlanList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServicePlanList_To_servicecatalog_ServicePlanList(a.(*ServicePlanList), b.(*servicecatalog.ServicePlanList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServicePlanList)(nil), (*ServicePlanList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServicePlanList_To_v1beta1_ServicePlanList(a.(*servicecatalog.ServicePlanList), b.(*ServicePlanList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServicePlanSpec)(nil), (*servicecatalog.ServicePlanSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServicePlanSpec_To_servicecatalog_ServicePlanSpec(a.(*ServicePlanSpec), b.(*servicecatalog.ServicePlanSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServicePlanSpec)(nil), (*ServicePlanSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServicePlanSpec_To_v1beta1_ServicePlanSpec(a.(*servicecatalog.ServicePlanSpec), b.(*ServicePlanSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ServicePlanStatus)(nil), (*servicecatalog.ServicePlanStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ServicePlanStatus_To_servicecatalog_ServicePlanStatus(a.(*ServicePlanStatus), b.(*servicecatalog.ServicePlanStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.ServicePlanStatus)(nil), (*ServicePlanStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_ServicePlanStatus_To_v1beta1_ServicePlanStatus(a.(*servicecatalog.ServicePlanStatus), b.(*ServicePlanStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UserInfo)(nil), (*servicecatalog.UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_UserInfo_To_servicecatalog_UserInfo(a.(*UserInfo), b.(*servicecatalog.UserInfo), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*servicecatalog.UserInfo)(nil), (*UserInfo)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_servicecatalog_UserInfo_To_v1beta1_UserInfo(a.(*servicecatalog.UserInfo), b.(*UserInfo), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1beta1_AddKeyTransform_To_servicecatalog_AddKeyTransform(in *AddKeyTransform, out *servicecatalog.AddKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Key = in.Key
	out.Value = *(*[]byte)(unsafe.Pointer(&in.Value))
	out.StringValue = (*string)(unsafe.Pointer(in.StringValue))
	out.JSONPathExpression = (*string)(unsafe.Pointer(in.JSONPathExpression))
	return nil
}
func Convert_v1beta1_AddKeyTransform_To_servicecatalog_AddKeyTransform(in *AddKeyTransform, out *servicecatalog.AddKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_AddKeyTransform_To_servicecatalog_AddKeyTransform(in, out, s)
}
func autoConvert_servicecatalog_AddKeyTransform_To_v1beta1_AddKeyTransform(in *servicecatalog.AddKeyTransform, out *AddKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Key = in.Key
	out.Value = *(*[]byte)(unsafe.Pointer(&in.Value))
	out.StringValue = (*string)(unsafe.Pointer(in.StringValue))
	out.JSONPathExpression = (*string)(unsafe.Pointer(in.JSONPathExpression))
	return nil
}
func Convert_servicecatalog_AddKeyTransform_To_v1beta1_AddKeyTransform(in *servicecatalog.AddKeyTransform, out *AddKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_AddKeyTransform_To_v1beta1_AddKeyTransform(in, out, s)
}
func autoConvert_v1beta1_AddKeysFromTransform_To_servicecatalog_AddKeysFromTransform(in *AddKeysFromTransform, out *servicecatalog.AddKeysFromTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*servicecatalog.ObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_v1beta1_AddKeysFromTransform_To_servicecatalog_AddKeysFromTransform(in *AddKeysFromTransform, out *servicecatalog.AddKeysFromTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_AddKeysFromTransform_To_servicecatalog_AddKeysFromTransform(in, out, s)
}
func autoConvert_servicecatalog_AddKeysFromTransform_To_v1beta1_AddKeysFromTransform(in *servicecatalog.AddKeysFromTransform, out *AddKeysFromTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*ObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_servicecatalog_AddKeysFromTransform_To_v1beta1_AddKeysFromTransform(in *servicecatalog.AddKeysFromTransform, out *AddKeysFromTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_AddKeysFromTransform_To_v1beta1_AddKeysFromTransform(in, out, s)
}
func autoConvert_v1beta1_BasicAuthConfig_To_servicecatalog_BasicAuthConfig(in *BasicAuthConfig, out *servicecatalog.BasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*servicecatalog.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_v1beta1_BasicAuthConfig_To_servicecatalog_BasicAuthConfig(in *BasicAuthConfig, out *servicecatalog.BasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_BasicAuthConfig_To_servicecatalog_BasicAuthConfig(in, out, s)
}
func autoConvert_servicecatalog_BasicAuthConfig_To_v1beta1_BasicAuthConfig(in *servicecatalog.BasicAuthConfig, out *BasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_servicecatalog_BasicAuthConfig_To_v1beta1_BasicAuthConfig(in *servicecatalog.BasicAuthConfig, out *BasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_BasicAuthConfig_To_v1beta1_BasicAuthConfig(in, out, s)
}
func autoConvert_v1beta1_BearerTokenAuthConfig_To_servicecatalog_BearerTokenAuthConfig(in *BearerTokenAuthConfig, out *servicecatalog.BearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*servicecatalog.LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_v1beta1_BearerTokenAuthConfig_To_servicecatalog_BearerTokenAuthConfig(in *BearerTokenAuthConfig, out *servicecatalog.BearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_BearerTokenAuthConfig_To_servicecatalog_BearerTokenAuthConfig(in, out, s)
}
func autoConvert_servicecatalog_BearerTokenAuthConfig_To_v1beta1_BearerTokenAuthConfig(in *servicecatalog.BearerTokenAuthConfig, out *BearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*LocalObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_servicecatalog_BearerTokenAuthConfig_To_v1beta1_BearerTokenAuthConfig(in *servicecatalog.BearerTokenAuthConfig, out *BearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_BearerTokenAuthConfig_To_v1beta1_BearerTokenAuthConfig(in, out, s)
}
func autoConvert_v1beta1_CatalogRestrictions_To_servicecatalog_CatalogRestrictions(in *CatalogRestrictions, out *servicecatalog.CatalogRestrictions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ServiceClass = *(*[]string)(unsafe.Pointer(&in.ServiceClass))
	out.ServicePlan = *(*[]string)(unsafe.Pointer(&in.ServicePlan))
	return nil
}
func Convert_v1beta1_CatalogRestrictions_To_servicecatalog_CatalogRestrictions(in *CatalogRestrictions, out *servicecatalog.CatalogRestrictions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CatalogRestrictions_To_servicecatalog_CatalogRestrictions(in, out, s)
}
func autoConvert_servicecatalog_CatalogRestrictions_To_v1beta1_CatalogRestrictions(in *servicecatalog.CatalogRestrictions, out *CatalogRestrictions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ServiceClass = *(*[]string)(unsafe.Pointer(&in.ServiceClass))
	out.ServicePlan = *(*[]string)(unsafe.Pointer(&in.ServicePlan))
	return nil
}
func Convert_servicecatalog_CatalogRestrictions_To_v1beta1_CatalogRestrictions(in *servicecatalog.CatalogRestrictions, out *CatalogRestrictions, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CatalogRestrictions_To_v1beta1_CatalogRestrictions(in, out, s)
}
func autoConvert_v1beta1_ClusterBasicAuthConfig_To_servicecatalog_ClusterBasicAuthConfig(in *ClusterBasicAuthConfig, out *servicecatalog.ClusterBasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*servicecatalog.ObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_v1beta1_ClusterBasicAuthConfig_To_servicecatalog_ClusterBasicAuthConfig(in *ClusterBasicAuthConfig, out *servicecatalog.ClusterBasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterBasicAuthConfig_To_servicecatalog_ClusterBasicAuthConfig(in, out, s)
}
func autoConvert_servicecatalog_ClusterBasicAuthConfig_To_v1beta1_ClusterBasicAuthConfig(in *servicecatalog.ClusterBasicAuthConfig, out *ClusterBasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*ObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_servicecatalog_ClusterBasicAuthConfig_To_v1beta1_ClusterBasicAuthConfig(in *servicecatalog.ClusterBasicAuthConfig, out *ClusterBasicAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterBasicAuthConfig_To_v1beta1_ClusterBasicAuthConfig(in, out, s)
}
func autoConvert_v1beta1_ClusterBearerTokenAuthConfig_To_servicecatalog_ClusterBearerTokenAuthConfig(in *ClusterBearerTokenAuthConfig, out *servicecatalog.ClusterBearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*servicecatalog.ObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_v1beta1_ClusterBearerTokenAuthConfig_To_servicecatalog_ClusterBearerTokenAuthConfig(in *ClusterBearerTokenAuthConfig, out *servicecatalog.ClusterBearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterBearerTokenAuthConfig_To_servicecatalog_ClusterBearerTokenAuthConfig(in, out, s)
}
func autoConvert_servicecatalog_ClusterBearerTokenAuthConfig_To_v1beta1_ClusterBearerTokenAuthConfig(in *servicecatalog.ClusterBearerTokenAuthConfig, out *ClusterBearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretRef = (*ObjectReference)(unsafe.Pointer(in.SecretRef))
	return nil
}
func Convert_servicecatalog_ClusterBearerTokenAuthConfig_To_v1beta1_ClusterBearerTokenAuthConfig(in *servicecatalog.ClusterBearerTokenAuthConfig, out *ClusterBearerTokenAuthConfig, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterBearerTokenAuthConfig_To_v1beta1_ClusterBearerTokenAuthConfig(in, out, s)
}
func autoConvert_v1beta1_ClusterObjectReference_To_servicecatalog_ClusterObjectReference(in *ClusterObjectReference, out *servicecatalog.ClusterObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	return nil
}
func Convert_v1beta1_ClusterObjectReference_To_servicecatalog_ClusterObjectReference(in *ClusterObjectReference, out *servicecatalog.ClusterObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterObjectReference_To_servicecatalog_ClusterObjectReference(in, out, s)
}
func autoConvert_servicecatalog_ClusterObjectReference_To_v1beta1_ClusterObjectReference(in *servicecatalog.ClusterObjectReference, out *ClusterObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	return nil
}
func Convert_servicecatalog_ClusterObjectReference_To_v1beta1_ClusterObjectReference(in *servicecatalog.ClusterObjectReference, out *ClusterObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterObjectReference_To_v1beta1_ClusterObjectReference(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceBroker_To_servicecatalog_ClusterServiceBroker(in *ClusterServiceBroker, out *servicecatalog.ClusterServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterServiceBrokerSpec_To_servicecatalog_ClusterServiceBrokerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterServiceBrokerStatus_To_servicecatalog_ClusterServiceBrokerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServiceBroker_To_servicecatalog_ClusterServiceBroker(in *ClusterServiceBroker, out *servicecatalog.ClusterServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceBroker_To_servicecatalog_ClusterServiceBroker(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceBroker_To_v1beta1_ClusterServiceBroker(in *servicecatalog.ClusterServiceBroker, out *ClusterServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ClusterServiceBrokerSpec_To_v1beta1_ClusterServiceBrokerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ClusterServiceBrokerStatus_To_v1beta1_ClusterServiceBrokerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServiceBroker_To_v1beta1_ClusterServiceBroker(in *servicecatalog.ClusterServiceBroker, out *ClusterServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceBroker_To_v1beta1_ClusterServiceBroker(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceBrokerAuthInfo_To_servicecatalog_ClusterServiceBrokerAuthInfo(in *ClusterServiceBrokerAuthInfo, out *servicecatalog.ClusterServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Basic = (*servicecatalog.ClusterBasicAuthConfig)(unsafe.Pointer(in.Basic))
	out.Bearer = (*servicecatalog.ClusterBearerTokenAuthConfig)(unsafe.Pointer(in.Bearer))
	return nil
}
func Convert_v1beta1_ClusterServiceBrokerAuthInfo_To_servicecatalog_ClusterServiceBrokerAuthInfo(in *ClusterServiceBrokerAuthInfo, out *servicecatalog.ClusterServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceBrokerAuthInfo_To_servicecatalog_ClusterServiceBrokerAuthInfo(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceBrokerAuthInfo_To_v1beta1_ClusterServiceBrokerAuthInfo(in *servicecatalog.ClusterServiceBrokerAuthInfo, out *ClusterServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Basic = (*ClusterBasicAuthConfig)(unsafe.Pointer(in.Basic))
	out.Bearer = (*ClusterBearerTokenAuthConfig)(unsafe.Pointer(in.Bearer))
	return nil
}
func Convert_servicecatalog_ClusterServiceBrokerAuthInfo_To_v1beta1_ClusterServiceBrokerAuthInfo(in *servicecatalog.ClusterServiceBrokerAuthInfo, out *ClusterServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceBrokerAuthInfo_To_v1beta1_ClusterServiceBrokerAuthInfo(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceBrokerList_To_servicecatalog_ClusterServiceBrokerList(in *ClusterServiceBrokerList, out *servicecatalog.ClusterServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ClusterServiceBroker)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ClusterServiceBrokerList_To_servicecatalog_ClusterServiceBrokerList(in *ClusterServiceBrokerList, out *servicecatalog.ClusterServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceBrokerList_To_servicecatalog_ClusterServiceBrokerList(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceBrokerList_To_v1beta1_ClusterServiceBrokerList(in *servicecatalog.ClusterServiceBrokerList, out *ClusterServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterServiceBroker)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ClusterServiceBrokerList_To_v1beta1_ClusterServiceBrokerList(in *servicecatalog.ClusterServiceBrokerList, out *ClusterServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceBrokerList_To_v1beta1_ClusterServiceBrokerList(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceBrokerSpec_To_servicecatalog_ClusterServiceBrokerSpec(in *ClusterServiceBrokerSpec, out *servicecatalog.ClusterServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceBrokerSpec_To_servicecatalog_CommonServiceBrokerSpec(&in.CommonServiceBrokerSpec, &out.CommonServiceBrokerSpec, s); err != nil {
		return err
	}
	out.AuthInfo = (*servicecatalog.ClusterServiceBrokerAuthInfo)(unsafe.Pointer(in.AuthInfo))
	return nil
}
func Convert_v1beta1_ClusterServiceBrokerSpec_To_servicecatalog_ClusterServiceBrokerSpec(in *ClusterServiceBrokerSpec, out *servicecatalog.ClusterServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceBrokerSpec_To_servicecatalog_ClusterServiceBrokerSpec(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceBrokerSpec_To_v1beta1_ClusterServiceBrokerSpec(in *servicecatalog.ClusterServiceBrokerSpec, out *ClusterServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceBrokerSpec_To_v1beta1_CommonServiceBrokerSpec(&in.CommonServiceBrokerSpec, &out.CommonServiceBrokerSpec, s); err != nil {
		return err
	}
	out.AuthInfo = (*ClusterServiceBrokerAuthInfo)(unsafe.Pointer(in.AuthInfo))
	return nil
}
func Convert_servicecatalog_ClusterServiceBrokerSpec_To_v1beta1_ClusterServiceBrokerSpec(in *servicecatalog.ClusterServiceBrokerSpec, out *ClusterServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceBrokerSpec_To_v1beta1_ClusterServiceBrokerSpec(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceBrokerStatus_To_servicecatalog_ClusterServiceBrokerStatus(in *ClusterServiceBrokerStatus, out *servicecatalog.ClusterServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceBrokerStatus_To_servicecatalog_CommonServiceBrokerStatus(&in.CommonServiceBrokerStatus, &out.CommonServiceBrokerStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServiceBrokerStatus_To_servicecatalog_ClusterServiceBrokerStatus(in *ClusterServiceBrokerStatus, out *servicecatalog.ClusterServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceBrokerStatus_To_servicecatalog_ClusterServiceBrokerStatus(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceBrokerStatus_To_v1beta1_ClusterServiceBrokerStatus(in *servicecatalog.ClusterServiceBrokerStatus, out *ClusterServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceBrokerStatus_To_v1beta1_CommonServiceBrokerStatus(&in.CommonServiceBrokerStatus, &out.CommonServiceBrokerStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServiceBrokerStatus_To_v1beta1_ClusterServiceBrokerStatus(in *servicecatalog.ClusterServiceBrokerStatus, out *ClusterServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceBrokerStatus_To_v1beta1_ClusterServiceBrokerStatus(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceClass_To_servicecatalog_ClusterServiceClass(in *ClusterServiceClass, out *servicecatalog.ClusterServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterServiceClassSpec_To_servicecatalog_ClusterServiceClassSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterServiceClassStatus_To_servicecatalog_ClusterServiceClassStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServiceClass_To_servicecatalog_ClusterServiceClass(in *ClusterServiceClass, out *servicecatalog.ClusterServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceClass_To_servicecatalog_ClusterServiceClass(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceClass_To_v1beta1_ClusterServiceClass(in *servicecatalog.ClusterServiceClass, out *ClusterServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ClusterServiceClassSpec_To_v1beta1_ClusterServiceClassSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ClusterServiceClassStatus_To_v1beta1_ClusterServiceClassStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServiceClass_To_v1beta1_ClusterServiceClass(in *servicecatalog.ClusterServiceClass, out *ClusterServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceClass_To_v1beta1_ClusterServiceClass(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceClassList_To_servicecatalog_ClusterServiceClassList(in *ClusterServiceClassList, out *servicecatalog.ClusterServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ClusterServiceClass)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ClusterServiceClassList_To_servicecatalog_ClusterServiceClassList(in *ClusterServiceClassList, out *servicecatalog.ClusterServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceClassList_To_servicecatalog_ClusterServiceClassList(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceClassList_To_v1beta1_ClusterServiceClassList(in *servicecatalog.ClusterServiceClassList, out *ClusterServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterServiceClass)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ClusterServiceClassList_To_v1beta1_ClusterServiceClassList(in *servicecatalog.ClusterServiceClassList, out *ClusterServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceClassList_To_v1beta1_ClusterServiceClassList(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceClassSpec_To_servicecatalog_ClusterServiceClassSpec(in *ClusterServiceClassSpec, out *servicecatalog.ClusterServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceClassSpec_To_servicecatalog_CommonServiceClassSpec(&in.CommonServiceClassSpec, &out.CommonServiceClassSpec, s); err != nil {
		return err
	}
	out.ClusterServiceBrokerName = in.ClusterServiceBrokerName
	return nil
}
func Convert_v1beta1_ClusterServiceClassSpec_To_servicecatalog_ClusterServiceClassSpec(in *ClusterServiceClassSpec, out *servicecatalog.ClusterServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceClassSpec_To_servicecatalog_ClusterServiceClassSpec(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceClassSpec_To_v1beta1_ClusterServiceClassSpec(in *servicecatalog.ClusterServiceClassSpec, out *ClusterServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceClassSpec_To_v1beta1_CommonServiceClassSpec(&in.CommonServiceClassSpec, &out.CommonServiceClassSpec, s); err != nil {
		return err
	}
	out.ClusterServiceBrokerName = in.ClusterServiceBrokerName
	return nil
}
func Convert_servicecatalog_ClusterServiceClassSpec_To_v1beta1_ClusterServiceClassSpec(in *servicecatalog.ClusterServiceClassSpec, out *ClusterServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceClassSpec_To_v1beta1_ClusterServiceClassSpec(in, out, s)
}
func autoConvert_v1beta1_ClusterServiceClassStatus_To_servicecatalog_ClusterServiceClassStatus(in *ClusterServiceClassStatus, out *servicecatalog.ClusterServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceClassStatus_To_servicecatalog_CommonServiceClassStatus(&in.CommonServiceClassStatus, &out.CommonServiceClassStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServiceClassStatus_To_servicecatalog_ClusterServiceClassStatus(in *ClusterServiceClassStatus, out *servicecatalog.ClusterServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServiceClassStatus_To_servicecatalog_ClusterServiceClassStatus(in, out, s)
}
func autoConvert_servicecatalog_ClusterServiceClassStatus_To_v1beta1_ClusterServiceClassStatus(in *servicecatalog.ClusterServiceClassStatus, out *ClusterServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceClassStatus_To_v1beta1_CommonServiceClassStatus(&in.CommonServiceClassStatus, &out.CommonServiceClassStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServiceClassStatus_To_v1beta1_ClusterServiceClassStatus(in *servicecatalog.ClusterServiceClassStatus, out *ClusterServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServiceClassStatus_To_v1beta1_ClusterServiceClassStatus(in, out, s)
}
func autoConvert_v1beta1_ClusterServicePlan_To_servicecatalog_ClusterServicePlan(in *ClusterServicePlan, out *servicecatalog.ClusterServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ClusterServicePlanSpec_To_servicecatalog_ClusterServicePlanSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ClusterServicePlanStatus_To_servicecatalog_ClusterServicePlanStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServicePlan_To_servicecatalog_ClusterServicePlan(in *ClusterServicePlan, out *servicecatalog.ClusterServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServicePlan_To_servicecatalog_ClusterServicePlan(in, out, s)
}
func autoConvert_servicecatalog_ClusterServicePlan_To_v1beta1_ClusterServicePlan(in *servicecatalog.ClusterServicePlan, out *ClusterServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ClusterServicePlanSpec_To_v1beta1_ClusterServicePlanSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ClusterServicePlanStatus_To_v1beta1_ClusterServicePlanStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServicePlan_To_v1beta1_ClusterServicePlan(in *servicecatalog.ClusterServicePlan, out *ClusterServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServicePlan_To_v1beta1_ClusterServicePlan(in, out, s)
}
func autoConvert_v1beta1_ClusterServicePlanList_To_servicecatalog_ClusterServicePlanList(in *ClusterServicePlanList, out *servicecatalog.ClusterServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ClusterServicePlan)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ClusterServicePlanList_To_servicecatalog_ClusterServicePlanList(in *ClusterServicePlanList, out *servicecatalog.ClusterServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServicePlanList_To_servicecatalog_ClusterServicePlanList(in, out, s)
}
func autoConvert_servicecatalog_ClusterServicePlanList_To_v1beta1_ClusterServicePlanList(in *servicecatalog.ClusterServicePlanList, out *ClusterServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterServicePlan)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ClusterServicePlanList_To_v1beta1_ClusterServicePlanList(in *servicecatalog.ClusterServicePlanList, out *ClusterServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServicePlanList_To_v1beta1_ClusterServicePlanList(in, out, s)
}
func autoConvert_v1beta1_ClusterServicePlanSpec_To_servicecatalog_ClusterServicePlanSpec(in *ClusterServicePlanSpec, out *servicecatalog.ClusterServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServicePlanSpec_To_servicecatalog_CommonServicePlanSpec(&in.CommonServicePlanSpec, &out.CommonServicePlanSpec, s); err != nil {
		return err
	}
	out.ClusterServiceBrokerName = in.ClusterServiceBrokerName
	if err := Convert_v1beta1_ClusterObjectReference_To_servicecatalog_ClusterObjectReference(&in.ClusterServiceClassRef, &out.ClusterServiceClassRef, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServicePlanSpec_To_servicecatalog_ClusterServicePlanSpec(in *ClusterServicePlanSpec, out *servicecatalog.ClusterServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServicePlanSpec_To_servicecatalog_ClusterServicePlanSpec(in, out, s)
}
func autoConvert_servicecatalog_ClusterServicePlanSpec_To_v1beta1_ClusterServicePlanSpec(in *servicecatalog.ClusterServicePlanSpec, out *ClusterServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServicePlanSpec_To_v1beta1_CommonServicePlanSpec(&in.CommonServicePlanSpec, &out.CommonServicePlanSpec, s); err != nil {
		return err
	}
	out.ClusterServiceBrokerName = in.ClusterServiceBrokerName
	if err := Convert_servicecatalog_ClusterObjectReference_To_v1beta1_ClusterObjectReference(&in.ClusterServiceClassRef, &out.ClusterServiceClassRef, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServicePlanSpec_To_v1beta1_ClusterServicePlanSpec(in *servicecatalog.ClusterServicePlanSpec, out *ClusterServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServicePlanSpec_To_v1beta1_ClusterServicePlanSpec(in, out, s)
}
func autoConvert_v1beta1_ClusterServicePlanStatus_To_servicecatalog_ClusterServicePlanStatus(in *ClusterServicePlanStatus, out *servicecatalog.ClusterServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServicePlanStatus_To_servicecatalog_CommonServicePlanStatus(&in.CommonServicePlanStatus, &out.CommonServicePlanStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ClusterServicePlanStatus_To_servicecatalog_ClusterServicePlanStatus(in *ClusterServicePlanStatus, out *servicecatalog.ClusterServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ClusterServicePlanStatus_To_servicecatalog_ClusterServicePlanStatus(in, out, s)
}
func autoConvert_servicecatalog_ClusterServicePlanStatus_To_v1beta1_ClusterServicePlanStatus(in *servicecatalog.ClusterServicePlanStatus, out *ClusterServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServicePlanStatus_To_v1beta1_CommonServicePlanStatus(&in.CommonServicePlanStatus, &out.CommonServicePlanStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ClusterServicePlanStatus_To_v1beta1_ClusterServicePlanStatus(in *servicecatalog.ClusterServicePlanStatus, out *ClusterServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ClusterServicePlanStatus_To_v1beta1_ClusterServicePlanStatus(in, out, s)
}
func autoConvert_v1beta1_CommonServiceBrokerSpec_To_servicecatalog_CommonServiceBrokerSpec(in *CommonServiceBrokerSpec, out *servicecatalog.CommonServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.RelistBehavior = servicecatalog.ServiceBrokerRelistBehavior(in.RelistBehavior)
	out.RelistDuration = (*v1.Duration)(unsafe.Pointer(in.RelistDuration))
	out.RelistRequests = in.RelistRequests
	out.CatalogRestrictions = (*servicecatalog.CatalogRestrictions)(unsafe.Pointer(in.CatalogRestrictions))
	return nil
}
func Convert_v1beta1_CommonServiceBrokerSpec_To_servicecatalog_CommonServiceBrokerSpec(in *CommonServiceBrokerSpec, out *servicecatalog.CommonServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CommonServiceBrokerSpec_To_servicecatalog_CommonServiceBrokerSpec(in, out, s)
}
func autoConvert_servicecatalog_CommonServiceBrokerSpec_To_v1beta1_CommonServiceBrokerSpec(in *servicecatalog.CommonServiceBrokerSpec, out *CommonServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.URL = in.URL
	out.InsecureSkipTLSVerify = in.InsecureSkipTLSVerify
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	out.RelistBehavior = ServiceBrokerRelistBehavior(in.RelistBehavior)
	out.RelistDuration = (*v1.Duration)(unsafe.Pointer(in.RelistDuration))
	out.RelistRequests = in.RelistRequests
	out.CatalogRestrictions = (*CatalogRestrictions)(unsafe.Pointer(in.CatalogRestrictions))
	return nil
}
func Convert_servicecatalog_CommonServiceBrokerSpec_To_v1beta1_CommonServiceBrokerSpec(in *servicecatalog.CommonServiceBrokerSpec, out *CommonServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CommonServiceBrokerSpec_To_v1beta1_CommonServiceBrokerSpec(in, out, s)
}
func autoConvert_v1beta1_CommonServiceBrokerStatus_To_servicecatalog_CommonServiceBrokerStatus(in *CommonServiceBrokerStatus, out *servicecatalog.CommonServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]servicecatalog.ServiceBrokerCondition)(unsafe.Pointer(&in.Conditions))
	out.ReconciledGeneration = in.ReconciledGeneration
	out.OperationStartTime = (*v1.Time)(unsafe.Pointer(in.OperationStartTime))
	out.LastCatalogRetrievalTime = (*v1.Time)(unsafe.Pointer(in.LastCatalogRetrievalTime))
	return nil
}
func Convert_v1beta1_CommonServiceBrokerStatus_To_servicecatalog_CommonServiceBrokerStatus(in *CommonServiceBrokerStatus, out *servicecatalog.CommonServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CommonServiceBrokerStatus_To_servicecatalog_CommonServiceBrokerStatus(in, out, s)
}
func autoConvert_servicecatalog_CommonServiceBrokerStatus_To_v1beta1_CommonServiceBrokerStatus(in *servicecatalog.CommonServiceBrokerStatus, out *CommonServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]ServiceBrokerCondition)(unsafe.Pointer(&in.Conditions))
	out.ReconciledGeneration = in.ReconciledGeneration
	out.OperationStartTime = (*v1.Time)(unsafe.Pointer(in.OperationStartTime))
	out.LastCatalogRetrievalTime = (*v1.Time)(unsafe.Pointer(in.LastCatalogRetrievalTime))
	return nil
}
func Convert_servicecatalog_CommonServiceBrokerStatus_To_v1beta1_CommonServiceBrokerStatus(in *servicecatalog.CommonServiceBrokerStatus, out *CommonServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CommonServiceBrokerStatus_To_v1beta1_CommonServiceBrokerStatus(in, out, s)
}
func autoConvert_v1beta1_CommonServiceClassSpec_To_servicecatalog_CommonServiceClassSpec(in *CommonServiceClassSpec, out *servicecatalog.CommonServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExternalName = in.ExternalName
	out.ExternalID = in.ExternalID
	out.Description = in.Description
	out.Bindable = in.Bindable
	out.BindingRetrievable = in.BindingRetrievable
	out.PlanUpdatable = in.PlanUpdatable
	out.ExternalMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.ExternalMetadata))
	out.Tags = *(*[]string)(unsafe.Pointer(&in.Tags))
	out.Requires = *(*[]string)(unsafe.Pointer(&in.Requires))
	out.DefaultProvisionParameters = (*runtime.RawExtension)(unsafe.Pointer(in.DefaultProvisionParameters))
	return nil
}
func Convert_v1beta1_CommonServiceClassSpec_To_servicecatalog_CommonServiceClassSpec(in *CommonServiceClassSpec, out *servicecatalog.CommonServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CommonServiceClassSpec_To_servicecatalog_CommonServiceClassSpec(in, out, s)
}
func autoConvert_servicecatalog_CommonServiceClassSpec_To_v1beta1_CommonServiceClassSpec(in *servicecatalog.CommonServiceClassSpec, out *CommonServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExternalName = in.ExternalName
	out.ExternalID = in.ExternalID
	out.Description = in.Description
	out.Bindable = in.Bindable
	out.BindingRetrievable = in.BindingRetrievable
	out.PlanUpdatable = in.PlanUpdatable
	out.ExternalMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.ExternalMetadata))
	out.Tags = *(*[]string)(unsafe.Pointer(&in.Tags))
	out.Requires = *(*[]string)(unsafe.Pointer(&in.Requires))
	out.DefaultProvisionParameters = (*runtime.RawExtension)(unsafe.Pointer(in.DefaultProvisionParameters))
	return nil
}
func Convert_servicecatalog_CommonServiceClassSpec_To_v1beta1_CommonServiceClassSpec(in *servicecatalog.CommonServiceClassSpec, out *CommonServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CommonServiceClassSpec_To_v1beta1_CommonServiceClassSpec(in, out, s)
}
func autoConvert_v1beta1_CommonServiceClassStatus_To_servicecatalog_CommonServiceClassStatus(in *CommonServiceClassStatus, out *servicecatalog.CommonServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RemovedFromBrokerCatalog = in.RemovedFromBrokerCatalog
	return nil
}
func Convert_v1beta1_CommonServiceClassStatus_To_servicecatalog_CommonServiceClassStatus(in *CommonServiceClassStatus, out *servicecatalog.CommonServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CommonServiceClassStatus_To_servicecatalog_CommonServiceClassStatus(in, out, s)
}
func autoConvert_servicecatalog_CommonServiceClassStatus_To_v1beta1_CommonServiceClassStatus(in *servicecatalog.CommonServiceClassStatus, out *CommonServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RemovedFromBrokerCatalog = in.RemovedFromBrokerCatalog
	return nil
}
func Convert_servicecatalog_CommonServiceClassStatus_To_v1beta1_CommonServiceClassStatus(in *servicecatalog.CommonServiceClassStatus, out *CommonServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CommonServiceClassStatus_To_v1beta1_CommonServiceClassStatus(in, out, s)
}
func autoConvert_v1beta1_CommonServicePlanSpec_To_servicecatalog_CommonServicePlanSpec(in *CommonServicePlanSpec, out *servicecatalog.CommonServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExternalName = in.ExternalName
	out.ExternalID = in.ExternalID
	out.Description = in.Description
	out.Bindable = (*bool)(unsafe.Pointer(in.Bindable))
	out.Free = in.Free
	out.ExternalMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.ExternalMetadata))
	out.InstanceCreateParameterSchema = (*runtime.RawExtension)(unsafe.Pointer(in.InstanceCreateParameterSchema))
	out.InstanceUpdateParameterSchema = (*runtime.RawExtension)(unsafe.Pointer(in.InstanceUpdateParameterSchema))
	out.ServiceBindingCreateParameterSchema = (*runtime.RawExtension)(unsafe.Pointer(in.ServiceBindingCreateParameterSchema))
	out.ServiceBindingCreateResponseSchema = (*runtime.RawExtension)(unsafe.Pointer(in.ServiceBindingCreateResponseSchema))
	out.DefaultProvisionParameters = (*runtime.RawExtension)(unsafe.Pointer(in.DefaultProvisionParameters))
	return nil
}
func Convert_v1beta1_CommonServicePlanSpec_To_servicecatalog_CommonServicePlanSpec(in *CommonServicePlanSpec, out *servicecatalog.CommonServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CommonServicePlanSpec_To_servicecatalog_CommonServicePlanSpec(in, out, s)
}
func autoConvert_servicecatalog_CommonServicePlanSpec_To_v1beta1_CommonServicePlanSpec(in *servicecatalog.CommonServicePlanSpec, out *CommonServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ExternalName = in.ExternalName
	out.ExternalID = in.ExternalID
	out.Description = in.Description
	out.Bindable = (*bool)(unsafe.Pointer(in.Bindable))
	out.Free = in.Free
	out.ExternalMetadata = (*runtime.RawExtension)(unsafe.Pointer(in.ExternalMetadata))
	out.InstanceCreateParameterSchema = (*runtime.RawExtension)(unsafe.Pointer(in.InstanceCreateParameterSchema))
	out.InstanceUpdateParameterSchema = (*runtime.RawExtension)(unsafe.Pointer(in.InstanceUpdateParameterSchema))
	out.ServiceBindingCreateParameterSchema = (*runtime.RawExtension)(unsafe.Pointer(in.ServiceBindingCreateParameterSchema))
	out.ServiceBindingCreateResponseSchema = (*runtime.RawExtension)(unsafe.Pointer(in.ServiceBindingCreateResponseSchema))
	out.DefaultProvisionParameters = (*runtime.RawExtension)(unsafe.Pointer(in.DefaultProvisionParameters))
	return nil
}
func Convert_servicecatalog_CommonServicePlanSpec_To_v1beta1_CommonServicePlanSpec(in *servicecatalog.CommonServicePlanSpec, out *CommonServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CommonServicePlanSpec_To_v1beta1_CommonServicePlanSpec(in, out, s)
}
func autoConvert_v1beta1_CommonServicePlanStatus_To_servicecatalog_CommonServicePlanStatus(in *CommonServicePlanStatus, out *servicecatalog.CommonServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RemovedFromBrokerCatalog = in.RemovedFromBrokerCatalog
	return nil
}
func Convert_v1beta1_CommonServicePlanStatus_To_servicecatalog_CommonServicePlanStatus(in *CommonServicePlanStatus, out *servicecatalog.CommonServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_CommonServicePlanStatus_To_servicecatalog_CommonServicePlanStatus(in, out, s)
}
func autoConvert_servicecatalog_CommonServicePlanStatus_To_v1beta1_CommonServicePlanStatus(in *servicecatalog.CommonServicePlanStatus, out *CommonServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RemovedFromBrokerCatalog = in.RemovedFromBrokerCatalog
	return nil
}
func Convert_servicecatalog_CommonServicePlanStatus_To_v1beta1_CommonServicePlanStatus(in *servicecatalog.CommonServicePlanStatus, out *CommonServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_CommonServicePlanStatus_To_v1beta1_CommonServicePlanStatus(in, out, s)
}
func autoConvert_v1beta1_LocalObjectReference_To_servicecatalog_LocalObjectReference(in *LocalObjectReference, out *servicecatalog.LocalObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	return nil
}
func Convert_v1beta1_LocalObjectReference_To_servicecatalog_LocalObjectReference(in *LocalObjectReference, out *servicecatalog.LocalObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_LocalObjectReference_To_servicecatalog_LocalObjectReference(in, out, s)
}
func autoConvert_servicecatalog_LocalObjectReference_To_v1beta1_LocalObjectReference(in *servicecatalog.LocalObjectReference, out *LocalObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	return nil
}
func Convert_servicecatalog_LocalObjectReference_To_v1beta1_LocalObjectReference(in *servicecatalog.LocalObjectReference, out *LocalObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_LocalObjectReference_To_v1beta1_LocalObjectReference(in, out, s)
}
func autoConvert_v1beta1_ObjectReference_To_servicecatalog_ObjectReference(in *ObjectReference, out *servicecatalog.ObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}
func Convert_v1beta1_ObjectReference_To_servicecatalog_ObjectReference(in *ObjectReference, out *servicecatalog.ObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ObjectReference_To_servicecatalog_ObjectReference(in, out, s)
}
func autoConvert_servicecatalog_ObjectReference_To_v1beta1_ObjectReference(in *servicecatalog.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}
func Convert_servicecatalog_ObjectReference_To_v1beta1_ObjectReference(in *servicecatalog.ObjectReference, out *ObjectReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ObjectReference_To_v1beta1_ObjectReference(in, out, s)
}
func autoConvert_v1beta1_ParametersFromSource_To_servicecatalog_ParametersFromSource(in *ParametersFromSource, out *servicecatalog.ParametersFromSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretKeyRef = (*servicecatalog.SecretKeyReference)(unsafe.Pointer(in.SecretKeyRef))
	return nil
}
func Convert_v1beta1_ParametersFromSource_To_servicecatalog_ParametersFromSource(in *ParametersFromSource, out *servicecatalog.ParametersFromSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ParametersFromSource_To_servicecatalog_ParametersFromSource(in, out, s)
}
func autoConvert_servicecatalog_ParametersFromSource_To_v1beta1_ParametersFromSource(in *servicecatalog.ParametersFromSource, out *ParametersFromSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.SecretKeyRef = (*SecretKeyReference)(unsafe.Pointer(in.SecretKeyRef))
	return nil
}
func Convert_servicecatalog_ParametersFromSource_To_v1beta1_ParametersFromSource(in *servicecatalog.ParametersFromSource, out *ParametersFromSource, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ParametersFromSource_To_v1beta1_ParametersFromSource(in, out, s)
}
func autoConvert_v1beta1_PlanReference_To_servicecatalog_PlanReference(in *PlanReference, out *servicecatalog.PlanReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClusterServiceClassExternalName = in.ClusterServiceClassExternalName
	out.ClusterServicePlanExternalName = in.ClusterServicePlanExternalName
	out.ClusterServiceClassExternalID = in.ClusterServiceClassExternalID
	out.ClusterServicePlanExternalID = in.ClusterServicePlanExternalID
	out.ClusterServiceClassName = in.ClusterServiceClassName
	out.ClusterServicePlanName = in.ClusterServicePlanName
	out.ServiceClassExternalName = in.ServiceClassExternalName
	out.ServicePlanExternalName = in.ServicePlanExternalName
	out.ServiceClassExternalID = in.ServiceClassExternalID
	out.ServicePlanExternalID = in.ServicePlanExternalID
	out.ServiceClassName = in.ServiceClassName
	out.ServicePlanName = in.ServicePlanName
	return nil
}
func Convert_v1beta1_PlanReference_To_servicecatalog_PlanReference(in *PlanReference, out *servicecatalog.PlanReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_PlanReference_To_servicecatalog_PlanReference(in, out, s)
}
func autoConvert_servicecatalog_PlanReference_To_v1beta1_PlanReference(in *servicecatalog.PlanReference, out *PlanReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClusterServiceClassExternalName = in.ClusterServiceClassExternalName
	out.ClusterServicePlanExternalName = in.ClusterServicePlanExternalName
	out.ClusterServiceClassExternalID = in.ClusterServiceClassExternalID
	out.ClusterServicePlanExternalID = in.ClusterServicePlanExternalID
	out.ClusterServiceClassName = in.ClusterServiceClassName
	out.ClusterServicePlanName = in.ClusterServicePlanName
	out.ServiceClassExternalName = in.ServiceClassExternalName
	out.ServicePlanExternalName = in.ServicePlanExternalName
	out.ServiceClassExternalID = in.ServiceClassExternalID
	out.ServicePlanExternalID = in.ServicePlanExternalID
	out.ServiceClassName = in.ServiceClassName
	out.ServicePlanName = in.ServicePlanName
	return nil
}
func Convert_servicecatalog_PlanReference_To_v1beta1_PlanReference(in *servicecatalog.PlanReference, out *PlanReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_PlanReference_To_v1beta1_PlanReference(in, out, s)
}
func autoConvert_v1beta1_RemoveKeyTransform_To_servicecatalog_RemoveKeyTransform(in *RemoveKeyTransform, out *servicecatalog.RemoveKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Key = in.Key
	return nil
}
func Convert_v1beta1_RemoveKeyTransform_To_servicecatalog_RemoveKeyTransform(in *RemoveKeyTransform, out *servicecatalog.RemoveKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_RemoveKeyTransform_To_servicecatalog_RemoveKeyTransform(in, out, s)
}
func autoConvert_servicecatalog_RemoveKeyTransform_To_v1beta1_RemoveKeyTransform(in *servicecatalog.RemoveKeyTransform, out *RemoveKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Key = in.Key
	return nil
}
func Convert_servicecatalog_RemoveKeyTransform_To_v1beta1_RemoveKeyTransform(in *servicecatalog.RemoveKeyTransform, out *RemoveKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_RemoveKeyTransform_To_v1beta1_RemoveKeyTransform(in, out, s)
}
func autoConvert_v1beta1_RenameKeyTransform_To_servicecatalog_RenameKeyTransform(in *RenameKeyTransform, out *servicecatalog.RenameKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.From = in.From
	out.To = in.To
	return nil
}
func Convert_v1beta1_RenameKeyTransform_To_servicecatalog_RenameKeyTransform(in *RenameKeyTransform, out *servicecatalog.RenameKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_RenameKeyTransform_To_servicecatalog_RenameKeyTransform(in, out, s)
}
func autoConvert_servicecatalog_RenameKeyTransform_To_v1beta1_RenameKeyTransform(in *servicecatalog.RenameKeyTransform, out *RenameKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.From = in.From
	out.To = in.To
	return nil
}
func Convert_servicecatalog_RenameKeyTransform_To_v1beta1_RenameKeyTransform(in *servicecatalog.RenameKeyTransform, out *RenameKeyTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_RenameKeyTransform_To_v1beta1_RenameKeyTransform(in, out, s)
}
func autoConvert_v1beta1_SecretKeyReference_To_servicecatalog_SecretKeyReference(in *SecretKeyReference, out *servicecatalog.SecretKeyReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Key = in.Key
	return nil
}
func Convert_v1beta1_SecretKeyReference_To_servicecatalog_SecretKeyReference(in *SecretKeyReference, out *servicecatalog.SecretKeyReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_SecretKeyReference_To_servicecatalog_SecretKeyReference(in, out, s)
}
func autoConvert_servicecatalog_SecretKeyReference_To_v1beta1_SecretKeyReference(in *servicecatalog.SecretKeyReference, out *SecretKeyReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Name = in.Name
	out.Key = in.Key
	return nil
}
func Convert_servicecatalog_SecretKeyReference_To_v1beta1_SecretKeyReference(in *servicecatalog.SecretKeyReference, out *SecretKeyReference, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_SecretKeyReference_To_v1beta1_SecretKeyReference(in, out, s)
}
func autoConvert_v1beta1_SecretTransform_To_servicecatalog_SecretTransform(in *SecretTransform, out *servicecatalog.SecretTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RenameKey = (*servicecatalog.RenameKeyTransform)(unsafe.Pointer(in.RenameKey))
	out.AddKey = (*servicecatalog.AddKeyTransform)(unsafe.Pointer(in.AddKey))
	out.AddKeysFrom = (*servicecatalog.AddKeysFromTransform)(unsafe.Pointer(in.AddKeysFrom))
	out.RemoveKey = (*servicecatalog.RemoveKeyTransform)(unsafe.Pointer(in.RemoveKey))
	return nil
}
func Convert_v1beta1_SecretTransform_To_servicecatalog_SecretTransform(in *SecretTransform, out *servicecatalog.SecretTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_SecretTransform_To_servicecatalog_SecretTransform(in, out, s)
}
func autoConvert_servicecatalog_SecretTransform_To_v1beta1_SecretTransform(in *servicecatalog.SecretTransform, out *SecretTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.RenameKey = (*RenameKeyTransform)(unsafe.Pointer(in.RenameKey))
	out.AddKey = (*AddKeyTransform)(unsafe.Pointer(in.AddKey))
	out.AddKeysFrom = (*AddKeysFromTransform)(unsafe.Pointer(in.AddKeysFrom))
	out.RemoveKey = (*RemoveKeyTransform)(unsafe.Pointer(in.RemoveKey))
	return nil
}
func Convert_servicecatalog_SecretTransform_To_v1beta1_SecretTransform(in *servicecatalog.SecretTransform, out *SecretTransform, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_SecretTransform_To_v1beta1_SecretTransform(in, out, s)
}
func autoConvert_v1beta1_ServiceBinding_To_servicecatalog_ServiceBinding(in *ServiceBinding, out *servicecatalog.ServiceBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ServiceBindingSpec_To_servicecatalog_ServiceBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ServiceBindingStatus_To_servicecatalog_ServiceBindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServiceBinding_To_servicecatalog_ServiceBinding(in *ServiceBinding, out *servicecatalog.ServiceBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBinding_To_servicecatalog_ServiceBinding(in, out, s)
}
func autoConvert_servicecatalog_ServiceBinding_To_v1beta1_ServiceBinding(in *servicecatalog.ServiceBinding, out *ServiceBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ServiceBindingSpec_To_v1beta1_ServiceBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ServiceBindingStatus_To_v1beta1_ServiceBindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServiceBinding_To_v1beta1_ServiceBinding(in *servicecatalog.ServiceBinding, out *ServiceBinding, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBinding_To_v1beta1_ServiceBinding(in, out, s)
}
func autoConvert_v1beta1_ServiceBindingCondition_To_servicecatalog_ServiceBindingCondition(in *ServiceBindingCondition, out *servicecatalog.ServiceBindingCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = servicecatalog.ServiceBindingConditionType(in.Type)
	out.Status = servicecatalog.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_v1beta1_ServiceBindingCondition_To_servicecatalog_ServiceBindingCondition(in *ServiceBindingCondition, out *servicecatalog.ServiceBindingCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBindingCondition_To_servicecatalog_ServiceBindingCondition(in, out, s)
}
func autoConvert_servicecatalog_ServiceBindingCondition_To_v1beta1_ServiceBindingCondition(in *servicecatalog.ServiceBindingCondition, out *ServiceBindingCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = ServiceBindingConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_servicecatalog_ServiceBindingCondition_To_v1beta1_ServiceBindingCondition(in *servicecatalog.ServiceBindingCondition, out *ServiceBindingCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBindingCondition_To_v1beta1_ServiceBindingCondition(in, out, s)
}
func autoConvert_v1beta1_ServiceBindingList_To_servicecatalog_ServiceBindingList(in *ServiceBindingList, out *servicecatalog.ServiceBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ServiceBinding)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ServiceBindingList_To_servicecatalog_ServiceBindingList(in *ServiceBindingList, out *servicecatalog.ServiceBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBindingList_To_servicecatalog_ServiceBindingList(in, out, s)
}
func autoConvert_servicecatalog_ServiceBindingList_To_v1beta1_ServiceBindingList(in *servicecatalog.ServiceBindingList, out *ServiceBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ServiceBinding)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ServiceBindingList_To_v1beta1_ServiceBindingList(in *servicecatalog.ServiceBindingList, out *ServiceBindingList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBindingList_To_v1beta1_ServiceBindingList(in, out, s)
}
func autoConvert_v1beta1_ServiceBindingPropertiesState_To_servicecatalog_ServiceBindingPropertiesState(in *ServiceBindingPropertiesState, out *servicecatalog.ServiceBindingPropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParameterChecksum = in.ParameterChecksum
	out.UserInfo = (*servicecatalog.UserInfo)(unsafe.Pointer(in.UserInfo))
	return nil
}
func Convert_v1beta1_ServiceBindingPropertiesState_To_servicecatalog_ServiceBindingPropertiesState(in *ServiceBindingPropertiesState, out *servicecatalog.ServiceBindingPropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBindingPropertiesState_To_servicecatalog_ServiceBindingPropertiesState(in, out, s)
}
func autoConvert_servicecatalog_ServiceBindingPropertiesState_To_v1beta1_ServiceBindingPropertiesState(in *servicecatalog.ServiceBindingPropertiesState, out *ServiceBindingPropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParameterChecksum = in.ParameterChecksum
	out.UserInfo = (*UserInfo)(unsafe.Pointer(in.UserInfo))
	return nil
}
func Convert_servicecatalog_ServiceBindingPropertiesState_To_v1beta1_ServiceBindingPropertiesState(in *servicecatalog.ServiceBindingPropertiesState, out *ServiceBindingPropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBindingPropertiesState_To_v1beta1_ServiceBindingPropertiesState(in, out, s)
}
func autoConvert_v1beta1_ServiceBindingSpec_To_servicecatalog_ServiceBindingSpec(in *ServiceBindingSpec, out *servicecatalog.ServiceBindingSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_LocalObjectReference_To_servicecatalog_LocalObjectReference(&in.InstanceRef, &out.InstanceRef, s); err != nil {
		return err
	}
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParametersFrom = *(*[]servicecatalog.ParametersFromSource)(unsafe.Pointer(&in.ParametersFrom))
	out.SecretName = in.SecretName
	out.SecretTransforms = *(*[]servicecatalog.SecretTransform)(unsafe.Pointer(&in.SecretTransforms))
	out.ExternalID = in.ExternalID
	out.UserInfo = (*servicecatalog.UserInfo)(unsafe.Pointer(in.UserInfo))
	return nil
}
func Convert_v1beta1_ServiceBindingSpec_To_servicecatalog_ServiceBindingSpec(in *ServiceBindingSpec, out *servicecatalog.ServiceBindingSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBindingSpec_To_servicecatalog_ServiceBindingSpec(in, out, s)
}
func autoConvert_servicecatalog_ServiceBindingSpec_To_v1beta1_ServiceBindingSpec(in *servicecatalog.ServiceBindingSpec, out *ServiceBindingSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_LocalObjectReference_To_v1beta1_LocalObjectReference(&in.InstanceRef, &out.InstanceRef, s); err != nil {
		return err
	}
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParametersFrom = *(*[]ParametersFromSource)(unsafe.Pointer(&in.ParametersFrom))
	out.SecretName = in.SecretName
	out.SecretTransforms = *(*[]SecretTransform)(unsafe.Pointer(&in.SecretTransforms))
	out.ExternalID = in.ExternalID
	out.UserInfo = (*UserInfo)(unsafe.Pointer(in.UserInfo))
	return nil
}
func Convert_servicecatalog_ServiceBindingSpec_To_v1beta1_ServiceBindingSpec(in *servicecatalog.ServiceBindingSpec, out *ServiceBindingSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBindingSpec_To_v1beta1_ServiceBindingSpec(in, out, s)
}
func autoConvert_v1beta1_ServiceBindingStatus_To_servicecatalog_ServiceBindingStatus(in *ServiceBindingStatus, out *servicecatalog.ServiceBindingStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]servicecatalog.ServiceBindingCondition)(unsafe.Pointer(&in.Conditions))
	out.AsyncOpInProgress = in.AsyncOpInProgress
	out.LastOperation = (*string)(unsafe.Pointer(in.LastOperation))
	out.CurrentOperation = servicecatalog.ServiceBindingOperation(in.CurrentOperation)
	out.ReconciledGeneration = in.ReconciledGeneration
	out.OperationStartTime = (*v1.Time)(unsafe.Pointer(in.OperationStartTime))
	out.InProgressProperties = (*servicecatalog.ServiceBindingPropertiesState)(unsafe.Pointer(in.InProgressProperties))
	out.ExternalProperties = (*servicecatalog.ServiceBindingPropertiesState)(unsafe.Pointer(in.ExternalProperties))
	out.OrphanMitigationInProgress = in.OrphanMitigationInProgress
	out.UnbindStatus = servicecatalog.ServiceBindingUnbindStatus(in.UnbindStatus)
	return nil
}
func Convert_v1beta1_ServiceBindingStatus_To_servicecatalog_ServiceBindingStatus(in *ServiceBindingStatus, out *servicecatalog.ServiceBindingStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBindingStatus_To_servicecatalog_ServiceBindingStatus(in, out, s)
}
func autoConvert_servicecatalog_ServiceBindingStatus_To_v1beta1_ServiceBindingStatus(in *servicecatalog.ServiceBindingStatus, out *ServiceBindingStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]ServiceBindingCondition)(unsafe.Pointer(&in.Conditions))
	out.AsyncOpInProgress = in.AsyncOpInProgress
	out.LastOperation = (*string)(unsafe.Pointer(in.LastOperation))
	out.CurrentOperation = ServiceBindingOperation(in.CurrentOperation)
	out.ReconciledGeneration = in.ReconciledGeneration
	out.OperationStartTime = (*v1.Time)(unsafe.Pointer(in.OperationStartTime))
	out.InProgressProperties = (*ServiceBindingPropertiesState)(unsafe.Pointer(in.InProgressProperties))
	out.ExternalProperties = (*ServiceBindingPropertiesState)(unsafe.Pointer(in.ExternalProperties))
	out.OrphanMitigationInProgress = in.OrphanMitigationInProgress
	out.UnbindStatus = ServiceBindingUnbindStatus(in.UnbindStatus)
	return nil
}
func Convert_servicecatalog_ServiceBindingStatus_To_v1beta1_ServiceBindingStatus(in *servicecatalog.ServiceBindingStatus, out *ServiceBindingStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBindingStatus_To_v1beta1_ServiceBindingStatus(in, out, s)
}
func autoConvert_v1beta1_ServiceBroker_To_servicecatalog_ServiceBroker(in *ServiceBroker, out *servicecatalog.ServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ServiceBrokerSpec_To_servicecatalog_ServiceBrokerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ServiceBrokerStatus_To_servicecatalog_ServiceBrokerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServiceBroker_To_servicecatalog_ServiceBroker(in *ServiceBroker, out *servicecatalog.ServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBroker_To_servicecatalog_ServiceBroker(in, out, s)
}
func autoConvert_servicecatalog_ServiceBroker_To_v1beta1_ServiceBroker(in *servicecatalog.ServiceBroker, out *ServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ServiceBrokerSpec_To_v1beta1_ServiceBrokerSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ServiceBrokerStatus_To_v1beta1_ServiceBrokerStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServiceBroker_To_v1beta1_ServiceBroker(in *servicecatalog.ServiceBroker, out *ServiceBroker, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBroker_To_v1beta1_ServiceBroker(in, out, s)
}
func autoConvert_v1beta1_ServiceBrokerAuthInfo_To_servicecatalog_ServiceBrokerAuthInfo(in *ServiceBrokerAuthInfo, out *servicecatalog.ServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Basic = (*servicecatalog.BasicAuthConfig)(unsafe.Pointer(in.Basic))
	out.Bearer = (*servicecatalog.BearerTokenAuthConfig)(unsafe.Pointer(in.Bearer))
	return nil
}
func Convert_v1beta1_ServiceBrokerAuthInfo_To_servicecatalog_ServiceBrokerAuthInfo(in *ServiceBrokerAuthInfo, out *servicecatalog.ServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBrokerAuthInfo_To_servicecatalog_ServiceBrokerAuthInfo(in, out, s)
}
func autoConvert_servicecatalog_ServiceBrokerAuthInfo_To_v1beta1_ServiceBrokerAuthInfo(in *servicecatalog.ServiceBrokerAuthInfo, out *ServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Basic = (*BasicAuthConfig)(unsafe.Pointer(in.Basic))
	out.Bearer = (*BearerTokenAuthConfig)(unsafe.Pointer(in.Bearer))
	return nil
}
func Convert_servicecatalog_ServiceBrokerAuthInfo_To_v1beta1_ServiceBrokerAuthInfo(in *servicecatalog.ServiceBrokerAuthInfo, out *ServiceBrokerAuthInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBrokerAuthInfo_To_v1beta1_ServiceBrokerAuthInfo(in, out, s)
}
func autoConvert_v1beta1_ServiceBrokerCondition_To_servicecatalog_ServiceBrokerCondition(in *ServiceBrokerCondition, out *servicecatalog.ServiceBrokerCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = servicecatalog.ServiceBrokerConditionType(in.Type)
	out.Status = servicecatalog.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_v1beta1_ServiceBrokerCondition_To_servicecatalog_ServiceBrokerCondition(in *ServiceBrokerCondition, out *servicecatalog.ServiceBrokerCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBrokerCondition_To_servicecatalog_ServiceBrokerCondition(in, out, s)
}
func autoConvert_servicecatalog_ServiceBrokerCondition_To_v1beta1_ServiceBrokerCondition(in *servicecatalog.ServiceBrokerCondition, out *ServiceBrokerCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = ServiceBrokerConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_servicecatalog_ServiceBrokerCondition_To_v1beta1_ServiceBrokerCondition(in *servicecatalog.ServiceBrokerCondition, out *ServiceBrokerCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBrokerCondition_To_v1beta1_ServiceBrokerCondition(in, out, s)
}
func autoConvert_v1beta1_ServiceBrokerList_To_servicecatalog_ServiceBrokerList(in *ServiceBrokerList, out *servicecatalog.ServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ServiceBroker)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ServiceBrokerList_To_servicecatalog_ServiceBrokerList(in *ServiceBrokerList, out *servicecatalog.ServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBrokerList_To_servicecatalog_ServiceBrokerList(in, out, s)
}
func autoConvert_servicecatalog_ServiceBrokerList_To_v1beta1_ServiceBrokerList(in *servicecatalog.ServiceBrokerList, out *ServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ServiceBroker)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ServiceBrokerList_To_v1beta1_ServiceBrokerList(in *servicecatalog.ServiceBrokerList, out *ServiceBrokerList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBrokerList_To_v1beta1_ServiceBrokerList(in, out, s)
}
func autoConvert_v1beta1_ServiceBrokerSpec_To_servicecatalog_ServiceBrokerSpec(in *ServiceBrokerSpec, out *servicecatalog.ServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceBrokerSpec_To_servicecatalog_CommonServiceBrokerSpec(&in.CommonServiceBrokerSpec, &out.CommonServiceBrokerSpec, s); err != nil {
		return err
	}
	out.AuthInfo = (*servicecatalog.ServiceBrokerAuthInfo)(unsafe.Pointer(in.AuthInfo))
	return nil
}
func Convert_v1beta1_ServiceBrokerSpec_To_servicecatalog_ServiceBrokerSpec(in *ServiceBrokerSpec, out *servicecatalog.ServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBrokerSpec_To_servicecatalog_ServiceBrokerSpec(in, out, s)
}
func autoConvert_servicecatalog_ServiceBrokerSpec_To_v1beta1_ServiceBrokerSpec(in *servicecatalog.ServiceBrokerSpec, out *ServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceBrokerSpec_To_v1beta1_CommonServiceBrokerSpec(&in.CommonServiceBrokerSpec, &out.CommonServiceBrokerSpec, s); err != nil {
		return err
	}
	out.AuthInfo = (*ServiceBrokerAuthInfo)(unsafe.Pointer(in.AuthInfo))
	return nil
}
func Convert_servicecatalog_ServiceBrokerSpec_To_v1beta1_ServiceBrokerSpec(in *servicecatalog.ServiceBrokerSpec, out *ServiceBrokerSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBrokerSpec_To_v1beta1_ServiceBrokerSpec(in, out, s)
}
func autoConvert_v1beta1_ServiceBrokerStatus_To_servicecatalog_ServiceBrokerStatus(in *ServiceBrokerStatus, out *servicecatalog.ServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceBrokerStatus_To_servicecatalog_CommonServiceBrokerStatus(&in.CommonServiceBrokerStatus, &out.CommonServiceBrokerStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServiceBrokerStatus_To_servicecatalog_ServiceBrokerStatus(in *ServiceBrokerStatus, out *servicecatalog.ServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceBrokerStatus_To_servicecatalog_ServiceBrokerStatus(in, out, s)
}
func autoConvert_servicecatalog_ServiceBrokerStatus_To_v1beta1_ServiceBrokerStatus(in *servicecatalog.ServiceBrokerStatus, out *ServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceBrokerStatus_To_v1beta1_CommonServiceBrokerStatus(&in.CommonServiceBrokerStatus, &out.CommonServiceBrokerStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServiceBrokerStatus_To_v1beta1_ServiceBrokerStatus(in *servicecatalog.ServiceBrokerStatus, out *ServiceBrokerStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceBrokerStatus_To_v1beta1_ServiceBrokerStatus(in, out, s)
}
func autoConvert_v1beta1_ServiceClass_To_servicecatalog_ServiceClass(in *ServiceClass, out *servicecatalog.ServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ServiceClassSpec_To_servicecatalog_ServiceClassSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ServiceClassStatus_To_servicecatalog_ServiceClassStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServiceClass_To_servicecatalog_ServiceClass(in *ServiceClass, out *servicecatalog.ServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceClass_To_servicecatalog_ServiceClass(in, out, s)
}
func autoConvert_servicecatalog_ServiceClass_To_v1beta1_ServiceClass(in *servicecatalog.ServiceClass, out *ServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ServiceClassSpec_To_v1beta1_ServiceClassSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ServiceClassStatus_To_v1beta1_ServiceClassStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServiceClass_To_v1beta1_ServiceClass(in *servicecatalog.ServiceClass, out *ServiceClass, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceClass_To_v1beta1_ServiceClass(in, out, s)
}
func autoConvert_v1beta1_ServiceClassList_To_servicecatalog_ServiceClassList(in *ServiceClassList, out *servicecatalog.ServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ServiceClass)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ServiceClassList_To_servicecatalog_ServiceClassList(in *ServiceClassList, out *servicecatalog.ServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceClassList_To_servicecatalog_ServiceClassList(in, out, s)
}
func autoConvert_servicecatalog_ServiceClassList_To_v1beta1_ServiceClassList(in *servicecatalog.ServiceClassList, out *ServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ServiceClass)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ServiceClassList_To_v1beta1_ServiceClassList(in *servicecatalog.ServiceClassList, out *ServiceClassList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceClassList_To_v1beta1_ServiceClassList(in, out, s)
}
func autoConvert_v1beta1_ServiceClassSpec_To_servicecatalog_ServiceClassSpec(in *ServiceClassSpec, out *servicecatalog.ServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceClassSpec_To_servicecatalog_CommonServiceClassSpec(&in.CommonServiceClassSpec, &out.CommonServiceClassSpec, s); err != nil {
		return err
	}
	out.ServiceBrokerName = in.ServiceBrokerName
	return nil
}
func Convert_v1beta1_ServiceClassSpec_To_servicecatalog_ServiceClassSpec(in *ServiceClassSpec, out *servicecatalog.ServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceClassSpec_To_servicecatalog_ServiceClassSpec(in, out, s)
}
func autoConvert_servicecatalog_ServiceClassSpec_To_v1beta1_ServiceClassSpec(in *servicecatalog.ServiceClassSpec, out *ServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceClassSpec_To_v1beta1_CommonServiceClassSpec(&in.CommonServiceClassSpec, &out.CommonServiceClassSpec, s); err != nil {
		return err
	}
	out.ServiceBrokerName = in.ServiceBrokerName
	return nil
}
func Convert_servicecatalog_ServiceClassSpec_To_v1beta1_ServiceClassSpec(in *servicecatalog.ServiceClassSpec, out *ServiceClassSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceClassSpec_To_v1beta1_ServiceClassSpec(in, out, s)
}
func autoConvert_v1beta1_ServiceClassStatus_To_servicecatalog_ServiceClassStatus(in *ServiceClassStatus, out *servicecatalog.ServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServiceClassStatus_To_servicecatalog_CommonServiceClassStatus(&in.CommonServiceClassStatus, &out.CommonServiceClassStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServiceClassStatus_To_servicecatalog_ServiceClassStatus(in *ServiceClassStatus, out *servicecatalog.ServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceClassStatus_To_servicecatalog_ServiceClassStatus(in, out, s)
}
func autoConvert_servicecatalog_ServiceClassStatus_To_v1beta1_ServiceClassStatus(in *servicecatalog.ServiceClassStatus, out *ServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServiceClassStatus_To_v1beta1_CommonServiceClassStatus(&in.CommonServiceClassStatus, &out.CommonServiceClassStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServiceClassStatus_To_v1beta1_ServiceClassStatus(in *servicecatalog.ServiceClassStatus, out *ServiceClassStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceClassStatus_To_v1beta1_ServiceClassStatus(in, out, s)
}
func autoConvert_v1beta1_ServiceInstance_To_servicecatalog_ServiceInstance(in *ServiceInstance, out *servicecatalog.ServiceInstance, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ServiceInstanceSpec_To_servicecatalog_ServiceInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ServiceInstanceStatus_To_servicecatalog_ServiceInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServiceInstance_To_servicecatalog_ServiceInstance(in *ServiceInstance, out *servicecatalog.ServiceInstance, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceInstance_To_servicecatalog_ServiceInstance(in, out, s)
}
func autoConvert_servicecatalog_ServiceInstance_To_v1beta1_ServiceInstance(in *servicecatalog.ServiceInstance, out *ServiceInstance, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ServiceInstanceSpec_To_v1beta1_ServiceInstanceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ServiceInstanceStatus_To_v1beta1_ServiceInstanceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServiceInstance_To_v1beta1_ServiceInstance(in *servicecatalog.ServiceInstance, out *ServiceInstance, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceInstance_To_v1beta1_ServiceInstance(in, out, s)
}
func autoConvert_v1beta1_ServiceInstanceCondition_To_servicecatalog_ServiceInstanceCondition(in *ServiceInstanceCondition, out *servicecatalog.ServiceInstanceCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = servicecatalog.ServiceInstanceConditionType(in.Type)
	out.Status = servicecatalog.ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_v1beta1_ServiceInstanceCondition_To_servicecatalog_ServiceInstanceCondition(in *ServiceInstanceCondition, out *servicecatalog.ServiceInstanceCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceInstanceCondition_To_servicecatalog_ServiceInstanceCondition(in, out, s)
}
func autoConvert_servicecatalog_ServiceInstanceCondition_To_v1beta1_ServiceInstanceCondition(in *servicecatalog.ServiceInstanceCondition, out *ServiceInstanceCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Type = ServiceInstanceConditionType(in.Type)
	out.Status = ConditionStatus(in.Status)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}
func Convert_servicecatalog_ServiceInstanceCondition_To_v1beta1_ServiceInstanceCondition(in *servicecatalog.ServiceInstanceCondition, out *ServiceInstanceCondition, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceInstanceCondition_To_v1beta1_ServiceInstanceCondition(in, out, s)
}
func autoConvert_v1beta1_ServiceInstanceList_To_servicecatalog_ServiceInstanceList(in *ServiceInstanceList, out *servicecatalog.ServiceInstanceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ServiceInstance)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ServiceInstanceList_To_servicecatalog_ServiceInstanceList(in *ServiceInstanceList, out *servicecatalog.ServiceInstanceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceInstanceList_To_servicecatalog_ServiceInstanceList(in, out, s)
}
func autoConvert_servicecatalog_ServiceInstanceList_To_v1beta1_ServiceInstanceList(in *servicecatalog.ServiceInstanceList, out *ServiceInstanceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ServiceInstance)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ServiceInstanceList_To_v1beta1_ServiceInstanceList(in *servicecatalog.ServiceInstanceList, out *ServiceInstanceList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceInstanceList_To_v1beta1_ServiceInstanceList(in, out, s)
}
func autoConvert_v1beta1_ServiceInstancePropertiesState_To_servicecatalog_ServiceInstancePropertiesState(in *ServiceInstancePropertiesState, out *servicecatalog.ServiceInstancePropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClusterServicePlanExternalName = in.ClusterServicePlanExternalName
	out.ClusterServicePlanExternalID = in.ClusterServicePlanExternalID
	out.ServicePlanExternalName = in.ServicePlanExternalName
	out.ServicePlanExternalID = in.ServicePlanExternalID
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParameterChecksum = in.ParameterChecksum
	out.UserInfo = (*servicecatalog.UserInfo)(unsafe.Pointer(in.UserInfo))
	return nil
}
func Convert_v1beta1_ServiceInstancePropertiesState_To_servicecatalog_ServiceInstancePropertiesState(in *ServiceInstancePropertiesState, out *servicecatalog.ServiceInstancePropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceInstancePropertiesState_To_servicecatalog_ServiceInstancePropertiesState(in, out, s)
}
func autoConvert_servicecatalog_ServiceInstancePropertiesState_To_v1beta1_ServiceInstancePropertiesState(in *servicecatalog.ServiceInstancePropertiesState, out *ServiceInstancePropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ClusterServicePlanExternalName = in.ClusterServicePlanExternalName
	out.ClusterServicePlanExternalID = in.ClusterServicePlanExternalID
	out.ServicePlanExternalName = in.ServicePlanExternalName
	out.ServicePlanExternalID = in.ServicePlanExternalID
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParameterChecksum = in.ParameterChecksum
	out.UserInfo = (*UserInfo)(unsafe.Pointer(in.UserInfo))
	return nil
}
func Convert_servicecatalog_ServiceInstancePropertiesState_To_v1beta1_ServiceInstancePropertiesState(in *servicecatalog.ServiceInstancePropertiesState, out *ServiceInstancePropertiesState, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceInstancePropertiesState_To_v1beta1_ServiceInstancePropertiesState(in, out, s)
}
func autoConvert_v1beta1_ServiceInstanceSpec_To_servicecatalog_ServiceInstanceSpec(in *ServiceInstanceSpec, out *servicecatalog.ServiceInstanceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_PlanReference_To_servicecatalog_PlanReference(&in.PlanReference, &out.PlanReference, s); err != nil {
		return err
	}
	out.ClusterServiceClassRef = (*servicecatalog.ClusterObjectReference)(unsafe.Pointer(in.ClusterServiceClassRef))
	out.ClusterServicePlanRef = (*servicecatalog.ClusterObjectReference)(unsafe.Pointer(in.ClusterServicePlanRef))
	out.ServiceClassRef = (*servicecatalog.LocalObjectReference)(unsafe.Pointer(in.ServiceClassRef))
	out.ServicePlanRef = (*servicecatalog.LocalObjectReference)(unsafe.Pointer(in.ServicePlanRef))
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParametersFrom = *(*[]servicecatalog.ParametersFromSource)(unsafe.Pointer(&in.ParametersFrom))
	out.ExternalID = in.ExternalID
	out.UserInfo = (*servicecatalog.UserInfo)(unsafe.Pointer(in.UserInfo))
	out.UpdateRequests = in.UpdateRequests
	return nil
}
func Convert_v1beta1_ServiceInstanceSpec_To_servicecatalog_ServiceInstanceSpec(in *ServiceInstanceSpec, out *servicecatalog.ServiceInstanceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceInstanceSpec_To_servicecatalog_ServiceInstanceSpec(in, out, s)
}
func autoConvert_servicecatalog_ServiceInstanceSpec_To_v1beta1_ServiceInstanceSpec(in *servicecatalog.ServiceInstanceSpec, out *ServiceInstanceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_PlanReference_To_v1beta1_PlanReference(&in.PlanReference, &out.PlanReference, s); err != nil {
		return err
	}
	out.ClusterServiceClassRef = (*ClusterObjectReference)(unsafe.Pointer(in.ClusterServiceClassRef))
	out.ClusterServicePlanRef = (*ClusterObjectReference)(unsafe.Pointer(in.ClusterServicePlanRef))
	out.ServiceClassRef = (*LocalObjectReference)(unsafe.Pointer(in.ServiceClassRef))
	out.ServicePlanRef = (*LocalObjectReference)(unsafe.Pointer(in.ServicePlanRef))
	out.Parameters = (*runtime.RawExtension)(unsafe.Pointer(in.Parameters))
	out.ParametersFrom = *(*[]ParametersFromSource)(unsafe.Pointer(&in.ParametersFrom))
	out.ExternalID = in.ExternalID
	out.UserInfo = (*UserInfo)(unsafe.Pointer(in.UserInfo))
	out.UpdateRequests = in.UpdateRequests
	return nil
}
func Convert_servicecatalog_ServiceInstanceSpec_To_v1beta1_ServiceInstanceSpec(in *servicecatalog.ServiceInstanceSpec, out *ServiceInstanceSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceInstanceSpec_To_v1beta1_ServiceInstanceSpec(in, out, s)
}
func autoConvert_v1beta1_ServiceInstanceStatus_To_servicecatalog_ServiceInstanceStatus(in *ServiceInstanceStatus, out *servicecatalog.ServiceInstanceStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]servicecatalog.ServiceInstanceCondition)(unsafe.Pointer(&in.Conditions))
	out.AsyncOpInProgress = in.AsyncOpInProgress
	out.OrphanMitigationInProgress = in.OrphanMitigationInProgress
	out.LastOperation = (*string)(unsafe.Pointer(in.LastOperation))
	out.DashboardURL = (*string)(unsafe.Pointer(in.DashboardURL))
	out.CurrentOperation = servicecatalog.ServiceInstanceOperation(in.CurrentOperation)
	out.ReconciledGeneration = in.ReconciledGeneration
	out.ObservedGeneration = in.ObservedGeneration
	out.OperationStartTime = (*v1.Time)(unsafe.Pointer(in.OperationStartTime))
	out.InProgressProperties = (*servicecatalog.ServiceInstancePropertiesState)(unsafe.Pointer(in.InProgressProperties))
	out.ExternalProperties = (*servicecatalog.ServiceInstancePropertiesState)(unsafe.Pointer(in.ExternalProperties))
	out.ProvisionStatus = servicecatalog.ServiceInstanceProvisionStatus(in.ProvisionStatus)
	out.DeprovisionStatus = servicecatalog.ServiceInstanceDeprovisionStatus(in.DeprovisionStatus)
	out.DefaultProvisionParameters = (*runtime.RawExtension)(unsafe.Pointer(in.DefaultProvisionParameters))
	return nil
}
func Convert_v1beta1_ServiceInstanceStatus_To_servicecatalog_ServiceInstanceStatus(in *ServiceInstanceStatus, out *servicecatalog.ServiceInstanceStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServiceInstanceStatus_To_servicecatalog_ServiceInstanceStatus(in, out, s)
}
func autoConvert_servicecatalog_ServiceInstanceStatus_To_v1beta1_ServiceInstanceStatus(in *servicecatalog.ServiceInstanceStatus, out *ServiceInstanceStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Conditions = *(*[]ServiceInstanceCondition)(unsafe.Pointer(&in.Conditions))
	out.AsyncOpInProgress = in.AsyncOpInProgress
	out.OrphanMitigationInProgress = in.OrphanMitigationInProgress
	out.LastOperation = (*string)(unsafe.Pointer(in.LastOperation))
	out.DashboardURL = (*string)(unsafe.Pointer(in.DashboardURL))
	out.CurrentOperation = ServiceInstanceOperation(in.CurrentOperation)
	out.ReconciledGeneration = in.ReconciledGeneration
	out.ObservedGeneration = in.ObservedGeneration
	out.OperationStartTime = (*v1.Time)(unsafe.Pointer(in.OperationStartTime))
	out.InProgressProperties = (*ServiceInstancePropertiesState)(unsafe.Pointer(in.InProgressProperties))
	out.ExternalProperties = (*ServiceInstancePropertiesState)(unsafe.Pointer(in.ExternalProperties))
	out.ProvisionStatus = ServiceInstanceProvisionStatus(in.ProvisionStatus)
	out.DeprovisionStatus = ServiceInstanceDeprovisionStatus(in.DeprovisionStatus)
	out.DefaultProvisionParameters = (*runtime.RawExtension)(unsafe.Pointer(in.DefaultProvisionParameters))
	return nil
}
func Convert_servicecatalog_ServiceInstanceStatus_To_v1beta1_ServiceInstanceStatus(in *servicecatalog.ServiceInstanceStatus, out *ServiceInstanceStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServiceInstanceStatus_To_v1beta1_ServiceInstanceStatus(in, out, s)
}
func autoConvert_v1beta1_ServicePlan_To_servicecatalog_ServicePlan(in *ServicePlan, out *servicecatalog.ServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_ServicePlanSpec_To_servicecatalog_ServicePlanSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ServicePlanStatus_To_servicecatalog_ServicePlanStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServicePlan_To_servicecatalog_ServicePlan(in *ServicePlan, out *servicecatalog.ServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServicePlan_To_servicecatalog_ServicePlan(in, out, s)
}
func autoConvert_servicecatalog_ServicePlan_To_v1beta1_ServicePlan(in *servicecatalog.ServicePlan, out *ServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_servicecatalog_ServicePlanSpec_To_v1beta1_ServicePlanSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_servicecatalog_ServicePlanStatus_To_v1beta1_ServicePlanStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServicePlan_To_v1beta1_ServicePlan(in *servicecatalog.ServicePlan, out *ServicePlan, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServicePlan_To_v1beta1_ServicePlan(in, out, s)
}
func autoConvert_v1beta1_ServicePlanList_To_servicecatalog_ServicePlanList(in *ServicePlanList, out *servicecatalog.ServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]servicecatalog.ServicePlan)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1beta1_ServicePlanList_To_servicecatalog_ServicePlanList(in *ServicePlanList, out *servicecatalog.ServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServicePlanList_To_servicecatalog_ServicePlanList(in, out, s)
}
func autoConvert_servicecatalog_ServicePlanList_To_v1beta1_ServicePlanList(in *servicecatalog.ServicePlanList, out *ServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ServicePlan)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_servicecatalog_ServicePlanList_To_v1beta1_ServicePlanList(in *servicecatalog.ServicePlanList, out *ServicePlanList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServicePlanList_To_v1beta1_ServicePlanList(in, out, s)
}
func autoConvert_v1beta1_ServicePlanSpec_To_servicecatalog_ServicePlanSpec(in *ServicePlanSpec, out *servicecatalog.ServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServicePlanSpec_To_servicecatalog_CommonServicePlanSpec(&in.CommonServicePlanSpec, &out.CommonServicePlanSpec, s); err != nil {
		return err
	}
	out.ServiceBrokerName = in.ServiceBrokerName
	if err := Convert_v1beta1_LocalObjectReference_To_servicecatalog_LocalObjectReference(&in.ServiceClassRef, &out.ServiceClassRef, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServicePlanSpec_To_servicecatalog_ServicePlanSpec(in *ServicePlanSpec, out *servicecatalog.ServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServicePlanSpec_To_servicecatalog_ServicePlanSpec(in, out, s)
}
func autoConvert_servicecatalog_ServicePlanSpec_To_v1beta1_ServicePlanSpec(in *servicecatalog.ServicePlanSpec, out *ServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServicePlanSpec_To_v1beta1_CommonServicePlanSpec(&in.CommonServicePlanSpec, &out.CommonServicePlanSpec, s); err != nil {
		return err
	}
	out.ServiceBrokerName = in.ServiceBrokerName
	if err := Convert_servicecatalog_LocalObjectReference_To_v1beta1_LocalObjectReference(&in.ServiceClassRef, &out.ServiceClassRef, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServicePlanSpec_To_v1beta1_ServicePlanSpec(in *servicecatalog.ServicePlanSpec, out *ServicePlanSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServicePlanSpec_To_v1beta1_ServicePlanSpec(in, out, s)
}
func autoConvert_v1beta1_ServicePlanStatus_To_servicecatalog_ServicePlanStatus(in *ServicePlanStatus, out *servicecatalog.ServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_v1beta1_CommonServicePlanStatus_To_servicecatalog_CommonServicePlanStatus(&in.CommonServicePlanStatus, &out.CommonServicePlanStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1beta1_ServicePlanStatus_To_servicecatalog_ServicePlanStatus(in *ServicePlanStatus, out *servicecatalog.ServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_ServicePlanStatus_To_servicecatalog_ServicePlanStatus(in, out, s)
}
func autoConvert_servicecatalog_ServicePlanStatus_To_v1beta1_ServicePlanStatus(in *servicecatalog.ServicePlanStatus, out *ServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := Convert_servicecatalog_CommonServicePlanStatus_To_v1beta1_CommonServicePlanStatus(&in.CommonServicePlanStatus, &out.CommonServicePlanStatus, s); err != nil {
		return err
	}
	return nil
}
func Convert_servicecatalog_ServicePlanStatus_To_v1beta1_ServicePlanStatus(in *servicecatalog.ServicePlanStatus, out *ServicePlanStatus, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_ServicePlanStatus_To_v1beta1_ServicePlanStatus(in, out, s)
}
func autoConvert_v1beta1_UserInfo_To_servicecatalog_UserInfo(in *UserInfo, out *servicecatalog.UserInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]servicecatalog.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}
func Convert_v1beta1_UserInfo_To_servicecatalog_UserInfo(in *UserInfo, out *servicecatalog.UserInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1beta1_UserInfo_To_servicecatalog_UserInfo(in, out, s)
}
func autoConvert_servicecatalog_UserInfo_To_v1beta1_UserInfo(in *servicecatalog.UserInfo, out *UserInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}
func Convert_servicecatalog_UserInfo_To_v1beta1_UserInfo(in *servicecatalog.UserInfo, out *UserInfo, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_servicecatalog_UserInfo_To_v1beta1_UserInfo(in, out, s)
}
