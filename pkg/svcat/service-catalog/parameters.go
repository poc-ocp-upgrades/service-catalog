package servicecatalog

import (
	"encoding/json"
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

func BuildParameters(params interface{}) *runtime.RawExtension {
	_logClusterCodePath()
	defer _logClusterCodePath()
	paramsJSON, err := json.Marshal(params)
	if err != nil {
		panic(fmt.Errorf("unable to marshal the request parameters %v (%s)", params, err))
	}
	return &runtime.RawExtension{Raw: paramsJSON}
}
func BuildParametersFrom(secrets map[string]string) []v1beta1.ParametersFromSource {
	_logClusterCodePath()
	defer _logClusterCodePath()
	params := make([]v1beta1.ParametersFromSource, 0, len(secrets))
	for secret, key := range secrets {
		param := v1beta1.ParametersFromSource{SecretKeyRef: &v1beta1.SecretKeyReference{Name: secret, Key: key}}
		params = append(params, param)
	}
	return params
}
