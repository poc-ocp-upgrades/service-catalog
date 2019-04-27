package e2e

import (
	"flag"
	"testing"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"github.com/kubernetes-incubator/service-catalog/test/e2e/framework"
)

var brokerImageFlag string

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flag.StringVar(&brokerImageFlag, "broker-image", "quay.io/kubernetes-service-catalog/user-broker:latest", "The container image for the broker to test against")
	framework.RegisterParseFlags()
	if "" == framework.TestContext.KubeConfig {
		klog.Fatalf("environment variable %v must be set", clientcmd.RecommendedConfigPathEnvVar)
	}
	if "" == framework.TestContext.ServiceCatalogConfig {
		klog.Fatalf("environment variable %v must be set", framework.RecommendedConfigPathEnvVar)
	}
}
func TestE2E(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RunE2ETests(t)
}
