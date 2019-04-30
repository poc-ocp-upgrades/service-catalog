package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodPreset struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	Spec			PodPresetSpec	`json:"spec,omitempty"`
}
type PodPresetSpec struct {
	Selector	metav1.LabelSelector	`json:"selector"`
	Env		[]v1.EnvVar		`json:"env,omitempty"`
	EnvFrom		[]v1.EnvFromSource	`json:"envFrom,omitempty"`
	Volumes		[]v1.Volume		`json:"volumes,omitempty"`
	VolumeMounts	[]v1.VolumeMount	`json:"volumeMounts,omitempty"`
}
type PodPresetList struct {
	metav1.TypeMeta	`json:",inline"`
	metav1.ListMeta	`json:"metadata,omitempty"`
	Items		[]PodPreset	`json:"items"`
}
