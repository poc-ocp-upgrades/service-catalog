package e2e

import (
	"testing"
	"k8s.io/apiserver/pkg/util/logs"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"github.com/kubernetes-incubator/service-catalog/test/e2e/framework"
	"k8s.io/klog"
	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/gomega"
)

func RunE2ETests(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	logs.InitLogs()
	defer logs.FlushLogs()
	gomega.RegisterFailHandler(ginkgo.Fail)
	if config.GinkgoConfig.FocusString == "" && config.GinkgoConfig.SkipString == "" {
		config.GinkgoConfig.SkipString = `\[Flaky\]|\[Feature:.+\]`
	}
	klog.Infof("Starting e2e run %q on Ginkgo node %d", framework.RunId, config.GinkgoConfig.ParallelNode)
	ginkgo.RunSpecs(t, "Service Catalog e2e suite")
}
