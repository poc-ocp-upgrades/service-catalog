package framework

import (
	"fmt"
	"time"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/uuid"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog"
)

const (
	poll		= 2 * time.Second
	defaultTimeout	= 30 * time.Second
)

func RestclientConfig(config, context string) (*api.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if config == "" {
		return nil, fmt.Errorf("Config file must be specified to load client config")
	}
	c, err := clientcmd.LoadFromFile(config)
	if err != nil {
		return nil, fmt.Errorf("error loading config: %v", err.Error())
	}
	if context != "" {
		c.CurrentContext = context
	}
	return c, nil
}
func LoadConfig(config, context string) (*rest.Config, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	c, err := RestclientConfig(config, context)
	if err != nil {
		return nil, err
	}
	return clientcmd.NewDefaultClientConfig(*c, &clientcmd.ConfigOverrides{}).ClientConfig()
}
func CreateKubeNamespace(c kubernetes.Interface) (*corev1.Namespace, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	ns := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{GenerateName: fmt.Sprintf("svc-catalog-health-check-%v-", uuid.NewUUID())}}
	var got *corev1.Namespace
	err := wait.PollImmediate(poll, defaultTimeout, func() (bool, error) {
		var err error
		got, err = c.CoreV1().Namespaces().Create(ns)
		if err != nil {
			klog.Errorf("Unexpected error while creating namespace: %v", err)
			return false, err
		}
		return true, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating test namespace: %v", err.Error())
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
func WaitForEndpoint(c kubernetes.Interface, namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.PollImmediate(poll, defaultTimeout, endpointAvailable(c, namespace, name))
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
