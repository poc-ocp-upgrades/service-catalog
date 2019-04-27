package svcat

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat/service-catalog"
	k8sclient "k8s.io/client-go/kubernetes"
)

type App struct {
	servicecatalog.SvcatClient
	CurrentNamespace	string
}

func NewApp(k8sClient k8sclient.Interface, serviceCatalogClient clientset.Interface, ns string) (*App, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	app := &App{SvcatClient: &servicecatalog.SDK{K8sClient: k8sClient, ServiceCatalogClient: serviceCatalogClient}, CurrentNamespace: ns}
	return app, nil
}
