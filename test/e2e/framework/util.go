package framework

import (
	"fmt"
	"time"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	Poll			= 2 * time.Second
	defaultTimeout		= 30 * time.Second
	EndpointRegisterTimeout	= time.Minute
)

func nowStamp() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return time.Now().Format(time.StampMilli)
}
func log(level string, format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(GinkgoWriter, nowStamp()+": "+level+": "+format+"\n", args...)
}
func Logf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	log("INFO", format, args...)
}
func Failf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	msg := fmt.Sprintf(format, args...)
	log("INFO", msg)
	Fail(nowStamp()+": "+msg, 1)
}
func Skipf(format string, args ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	msg := fmt.Sprintf(format, args...)
	log("INFO", msg)
	Skip(nowStamp() + ": " + msg)
}

type ClientConfigGetter func() (*rest.Config, error)

var RunId = uuid.NewUUID()

func CreateKubeNamespace(baseName string, c kubernetes.Interface) (*corev1.Namespace, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: fmt.Sprintf("e2e-tests-%v-", baseName)}}
	Logf("namespace: %v", ns)
	var got *corev1.Namespace
	err := wait.PollImmediate(Poll, defaultTimeout, func() (bool, error) {
		var err error
		got, err = c.CoreV1().Namespaces().Create(ns)
		if err != nil {
			Logf("Unexpected error while creating namespace: %v", err)
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}
	return got, nil
}
func DeleteKubeNamespace(c kubernetes.Interface, namespace string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.CoreV1().Namespaces().Delete(namespace, nil)
}
func ExpectNoError(err error, explain ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err != nil {
		Logf("Unexpected error occurred: %v", err)
	}
	ExpectWithOffset(1, err).NotTo(HaveOccurred(), explain...)
}
func WaitForPodRunningInNamespace(c kubernetes.Interface, pod *corev1.Pod) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if pod.Status.Phase == corev1.PodRunning {
		return nil
	}
	return waitTimeoutForPodRunningInNamespace(c, pod.Name, pod.Namespace, defaultTimeout)
}
func waitTimeoutForPodRunningInNamespace(c kubernetes.Interface, podName, namespace string, timeout time.Duration) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(Poll, defaultTimeout, podRunning(c, podName, namespace))
}
func WaitForEndpoint(c kubernetes.Interface, namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(Poll, EndpointRegisterTimeout, endpointAvailable(c, namespace, name))
}
func endpointAvailable(c kubernetes.Interface, namespace, name string) wait.ConditionFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() (bool, error) {
		endpoint, err := c.CoreV1().Endpoints(namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			if apierrs.IsNotFound(err) {
				return false, nil
			}
			return false, err
		}
		if len(endpoint.Subsets) == 0 || len(endpoint.Subsets[0].Addresses) == 0 {
			return false, nil
		}
		return true, nil
	}
}
func podRunning(c kubernetes.Interface, podName, namespace string) wait.ConditionFunc {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return func() (bool, error) {
		pod, err := c.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		switch pod.Status.Phase {
		case corev1.PodRunning:
			return true, nil
		case corev1.PodFailed, corev1.PodSucceeded:
			return false, fmt.Errorf("pod ran to completion")
		}
		return false, nil
	}
}
