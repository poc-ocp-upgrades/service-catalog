package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type SettingsV1alpha1Interface interface {
	RESTClient() rest.Interface
	PodPresetsGetter
}
type SettingsV1alpha1Client struct{ restClient rest.Interface }

func (c *SettingsV1alpha1Client) PodPresets(namespace string) PodPresetInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return newPodPresets(c, namespace)
}
func NewForConfig(c *rest.Config) (*SettingsV1alpha1Client, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SettingsV1alpha1Client{client}, nil
}
func NewForConfigOrDie(c *rest.Config) *SettingsV1alpha1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
func New(c rest.Interface) *SettingsV1alpha1Client {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &SettingsV1alpha1Client{c}
}
func setConfigDefaults(config *rest.Config) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	return nil
}
func (c *SettingsV1alpha1Client) RESTClient() rest.Interface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.restClient
}
