package plan_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	_ "github.com/kubernetes-incubator/service-catalog/internal/test"
)

func TestPlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plan Suite")
}
