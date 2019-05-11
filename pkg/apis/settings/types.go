package settings

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodPreset struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec	PodPresetSpec
}
type PodPresetSpec struct {
	Selector		metav1.LabelSelector
	Env				[]v1.EnvVar
	EnvFrom			[]v1.EnvFromSource
	Volumes			[]v1.Volume
	VolumeMounts	[]v1.VolumeMount
}
type PodPresetList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items	[]PodPreset
}
