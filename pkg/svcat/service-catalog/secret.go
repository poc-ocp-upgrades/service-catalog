package servicecatalog

import (
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (sdk *SDK) RetrieveSecretByBinding(binding *v1beta1.ServiceBinding) (*corev1.Secret, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	secret, err := sdk.Core().Secrets(binding.Namespace).Get(binding.Spec.SecretName, metav1.GetOptions{})
	if err != nil {
		if !sdk.IsBindingReady(binding) && errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get secret %s/%s (%s)", binding.Namespace, binding.Spec.SecretName, err)
	}
	return secret, nil
}
