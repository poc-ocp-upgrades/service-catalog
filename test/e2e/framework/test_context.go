package framework

import (
	"flag"
	"os"
	"github.com/onsi/ginkgo/config"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	RecommendedConfigPathEnvVar = "SERVICECATALOGCONFIG"
)

type TestContextType struct {
	KubeHost		string
	KubeConfig		string
	KubeContext		string
	ServiceCatalogHost	string
	ServiceCatalogConfig	string
	ServiceCatalogContext	string
}

var TestContext TestContextType

func RegisterCommonFlags() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	config.DefaultReporterConfig.Verbose = true
	config.GinkgoConfig.EmitSpecProgress = true
	config.GinkgoConfig.RandomizeAllSpecs = true
	flag.StringVar(&TestContext.KubeHost, "kubernetes-host", "http://127.0.0.1:8080", "The kubernetes host, or apiserver, to connect to")
	flag.StringVar(&TestContext.KubeConfig, "kubernetes-config", os.Getenv(clientcmd.RecommendedConfigPathEnvVar), "Path to config containing embedded authinfo for kubernetes. Default value is from environment variable "+clientcmd.RecommendedConfigPathEnvVar)
	flag.StringVar(&TestContext.KubeContext, "kubernetes-context", "", "config context to use for kubernetes. If unset, will use value from 'current-context'")
	flag.StringVar(&TestContext.ServiceCatalogHost, "service-catalog-host", "http://127.0.0.1:30000", "The service catalog host, or apiserver, to connect to")
	flag.StringVar(&TestContext.ServiceCatalogConfig, "service-catalog-config", os.Getenv(RecommendedConfigPathEnvVar), "Path to config containing embedded authinfo for service catalog. Default value is from environment variable "+RecommendedConfigPathEnvVar)
	flag.StringVar(&TestContext.ServiceCatalogContext, "service-catalog-context", "", "config context to use for service catalog. If unset, will use value from 'current-context'")
}
func RegisterParseFlags() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RegisterCommonFlags()
	flag.Parse()
}
