package v1alpha1

import (
	unsafe "unsafe"
	settings "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	localSchemeBuilder.Register(RegisterConversions)
}
func RegisterConversions(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := s.AddGeneratedConversionFunc((*PodPreset)(nil), (*settings.PodPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodPreset_To_settings_PodPreset(a.(*PodPreset), b.(*settings.PodPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.PodPreset)(nil), (*PodPreset)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_PodPreset_To_v1alpha1_PodPreset(a.(*settings.PodPreset), b.(*PodPreset), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodPresetList)(nil), (*settings.PodPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(a.(*PodPresetList), b.(*settings.PodPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.PodPresetList)(nil), (*PodPresetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(a.(*settings.PodPresetList), b.(*PodPresetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PodPresetSpec)(nil), (*settings.PodPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(a.(*PodPresetSpec), b.(*settings.PodPresetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*settings.PodPresetSpec)(nil), (*PodPresetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(a.(*settings.PodPresetSpec), b.(*PodPresetSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}
func autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in *PodPreset, out *settings.PodPreset, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_v1alpha1_PodPreset_To_settings_PodPreset(in *PodPreset, out *settings.PodPreset, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1alpha1_PodPreset_To_settings_PodPreset(in, out, s)
}
func autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *PodPreset, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}
func Convert_settings_PodPreset_To_v1alpha1_PodPreset(in *settings.PodPreset, out *PodPreset, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_settings_PodPreset_To_v1alpha1_PodPreset(in, out, s)
}
func autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]settings.PodPreset)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_v1alpha1_PodPresetList_To_settings_PodPresetList(in *PodPresetList, out *settings.PodPresetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1alpha1_PodPresetList_To_settings_PodPresetList(in, out, s)
}
func autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *PodPresetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PodPreset)(unsafe.Pointer(&in.Items))
	return nil
}
func Convert_settings_PodPresetList_To_v1alpha1_PodPresetList(in *settings.PodPresetList, out *PodPresetList, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_settings_PodPresetList_To_v1alpha1_PodPresetList(in, out, s)
}
func autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Selector = in.Selector
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	out.Volumes = *(*[]v1.Volume)(unsafe.Pointer(&in.Volumes))
	out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
	return nil
}
func Convert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in *PodPresetSpec, out *settings.PodPresetSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_v1alpha1_PodPresetSpec_To_settings_PodPresetSpec(in, out, s)
}
func autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *PodPresetSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	out.Selector = in.Selector
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.EnvFrom = *(*[]v1.EnvFromSource)(unsafe.Pointer(&in.EnvFrom))
	out.Volumes = *(*[]v1.Volume)(unsafe.Pointer(&in.Volumes))
	out.VolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.VolumeMounts))
	return nil
}
func Convert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in *settings.PodPresetSpec, out *PodPresetSpec, s conversion.Scope) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autoConvert_settings_PodPresetSpec_To_v1alpha1_PodPresetSpec(in, out, s)
}
