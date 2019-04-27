package validation

import (
	"strings"
	"testing"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/settings"
)

func TestValidateEmptyPodPreset(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	emptyPodPreset := &settings.PodPreset{Spec: settings.PodPresetSpec{}}
	errList := ValidatePodPreset(emptyPodPreset)
	if errList == nil {
		t.Fatal("empty pod preset should return an error")
	}
}
func TestValidateEmptyPodPresetItems(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	emptyPodPreset := &settings.PodPreset{ObjectMeta: metav1.ObjectMeta{Name: "hello", Namespace: "sample"}, Spec: settings.PodPresetSpec{Selector: metav1.LabelSelector{MatchExpressions: []metav1.LabelSelectorRequirement{{Key: "security", Operator: metav1.LabelSelectorOpIn, Values: []string{"S2"}}}}}}
	errList := ValidatePodPreset(emptyPodPreset)
	if !strings.Contains(errList.ToAggregate().Error(), "must specify at least one") {
		t.Fatal("empty pod preset with label selector should return an error")
	}
}
func TestValidatePodPresets(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	p := &settings.PodPreset{ObjectMeta: metav1.ObjectMeta{Name: "hello", Namespace: "sample"}, Spec: settings.PodPresetSpec{Selector: metav1.LabelSelector{MatchExpressions: []metav1.LabelSelectorRequirement{{Key: "security", Operator: metav1.LabelSelectorOpIn, Values: []string{"S2"}}}}, Volumes: []v1.Volume{{Name: "vol", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}}}, Env: []v1.EnvVar{{Name: "abc", Value: "value"}, {Name: "ABC", Value: "value"}}, EnvFrom: []v1.EnvFromSource{{ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "abc"}}}, {Prefix: "pre_", ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "abc"}}}}}}
	errList := ValidatePodPreset(p)
	if errList != nil {
		if errList.ToAggregate() != nil {
			t.Fatalf("errors: %#v", errList.ToAggregate().Error())
		}
	}
	p = &settings.PodPreset{ObjectMeta: metav1.ObjectMeta{Name: "hello", Namespace: "sample"}, Spec: settings.PodPresetSpec{Selector: metav1.LabelSelector{MatchExpressions: []metav1.LabelSelectorRequirement{{Key: "security", Operator: metav1.LabelSelectorOpIn, Values: []string{"S2"}}}}, Volumes: []v1.Volume{{Name: "vol", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}}}, Env: []v1.EnvVar{{Name: "abc", Value: "value"}, {Name: "ABC", Value: "value"}}, VolumeMounts: []v1.VolumeMount{{Name: "vol", MountPath: "/foo"}}, EnvFrom: []v1.EnvFromSource{{ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "abc"}}}, {Prefix: "pre_", ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "abc"}}}}}}
	errList = ValidatePodPreset(p)
	if errList != nil {
		if errList.ToAggregate() != nil {
			t.Fatalf("errors: %#v", errList.ToAggregate().Error())
		}
	}
}
func TestValidatePodPresetsiVolumeMountError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t.Skipf("skipping this test till validation for volume is in place")
	p := &settings.PodPreset{ObjectMeta: metav1.ObjectMeta{Name: "hello", Namespace: "sample"}, Spec: settings.PodPresetSpec{Selector: metav1.LabelSelector{MatchExpressions: []metav1.LabelSelectorRequirement{{Key: "security", Operator: metav1.LabelSelectorOpIn, Values: []string{"S2"}}}}, Volumes: []v1.Volume{{Name: "vol", VolumeSource: v1.VolumeSource{EmptyDir: &v1.EmptyDirVolumeSource{}}}}, VolumeMounts: []v1.VolumeMount{{Name: "dne", MountPath: "/foo"}}, Env: []v1.EnvVar{{Name: "abc", Value: "value"}, {Name: "ABC", Value: "value"}}, EnvFrom: []v1.EnvFromSource{{ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "abc"}}}, {Prefix: "pre_", ConfigMapRef: &v1.ConfigMapEnvSource{LocalObjectReference: v1.LocalObjectReference{Name: "abc"}}}}}}
	errList := ValidatePodPreset(p)
	if !strings.Contains(errList.ToAggregate().Error(), "spec.volumeMounts[0].name: Not found") {
		t.Fatal("should have returned error for volume that does not exist")
	}
}
