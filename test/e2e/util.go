package e2e

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func NewUPSBrokerPod(name string) *corev1.Pod {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: name, Labels: map[string]string{"app": name}}, Spec: corev1.PodSpec{Containers: []corev1.Container{{Name: name, Image: brokerImageFlag, Args: []string{"--port", "8080"}, Ports: []corev1.ContainerPort{{ContainerPort: 8080}}}}}}
}
func NewUPSBrokerService(name string) *corev1.Service {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: name, Labels: map[string]string{"app": name}}, Spec: corev1.ServiceSpec{Selector: map[string]string{"app": name}, Ports: []corev1.ServicePort{{Protocol: corev1.ProtocolTCP, Port: 80, TargetPort: intstr.FromInt(8080)}}}}
}
