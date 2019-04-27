package fake

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/typed/settings/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSettingsV1alpha1 struct{ *testing.Fake }

func (c *FakeSettingsV1alpha1) PodPresets(namespace string) v1alpha1.PodPresetInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &FakePodPresets{c, namespace}
}
func (c *FakeSettingsV1alpha1) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var ret *rest.RESTClient
	return ret
}
