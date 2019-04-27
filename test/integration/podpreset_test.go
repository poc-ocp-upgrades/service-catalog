package integration

import (
	"fmt"
	"reflect"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/features"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	settingsapi "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	servicecatalogclient "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"k8s.io/apimachinery/pkg/runtime"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
)

func TestPodPresetClient(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	enablePodPresetFeature()
	defer disablePodPresetFeature()
	const name = "test-podpreset"
	client, _, shutdown := getFreshApiserverAndClient(t, func() runtime.Object {
		return &settingsapi.PodPreset{}
	})
	defer shutdown()
	if err := testPodPresetClient(client, name); err != nil {
		t.Fatal(err)
	}
}
func enablePodPresetFeature() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", features.PodPreset))
}
func disablePodPresetFeature() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	utilfeature.DefaultFeatureGate.Set(fmt.Sprintf("%v=true", features.PodPreset))
}
func testPodPresetClient(client servicecatalogclient.Interface, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testNamespace := "test-namespace"
	podPresetName := "test-podpreset"
	cl := client.Settings().PodPresets(testNamespace)
	tests := []struct{ input *settingsapi.PodPreset }{{input: &settingsapi.PodPreset{ObjectMeta: metav1.ObjectMeta{Name: podPresetName, Namespace: testNamespace}, Spec: settingsapi.PodPresetSpec{Selector: metav1.LabelSelector{MatchExpressions: []metav1.LabelSelectorRequirement{{Key: "security", Operator: metav1.LabelSelectorOpIn, Values: []string{"S2"}}}}, Volumes: []corev1.Volume{{Name: "vol", VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}}}}, VolumeMounts: []corev1.VolumeMount{{Name: "vol", MountPath: "/foo"}}, Env: []corev1.EnvVar{{Name: "abc", Value: "value"}, {Name: "ABC", Value: "value"}}}}}}
	for _, test := range tests {
		podpresets, err := cl.List(metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("error listing podpresets: %v", err)
		}
		if n := len(podpresets.Items); n > 0 {
			return fmt.Errorf("podpresets should not exist on start, found %v podpresets", n)
		}
		in := test.input
		out, err := cl.Create(in)
		if err != nil {
			return fmt.Errorf("error creating podpreset :%v", err)
		}
		if in.Name != out.Name {
			return fmt.Errorf("name doesn't match: %v", err)
		}
		podpresets, err = cl.List(metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("error listing podpreset :%v", err)
		}
		if n := len(podpresets.Items); n != 1 {
			return fmt.Errorf("expected list size to be 1 and got: %d", n)
		}
		got, err := cl.Get(podPresetName, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("error listing podpreset :%v", err)
		}
		if !reflect.DeepEqual(got, out) {
			return fmt.Errorf("objects do not match")
		}
		err = cl.Delete(podPresetName, &metav1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("error deleting podpreset : %v", err)
		}
		podpresets, err = cl.List(metav1.ListOptions{})
		if err != nil {
			return fmt.Errorf("error listing podpreset : %v", err)
		}
		if n := len(podpresets.Items); n != 0 {
			return fmt.Errorf("expected no podpresets, but found: %d", n)
		}
	}
	return nil
}
