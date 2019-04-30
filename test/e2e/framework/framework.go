package framework

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-incubator/service-catalog/pkg/util/kube"
)

type Framework struct {
	BaseName		string
	KubeClientSet		kubernetes.Interface
	ServiceCatalogClientSet	clientset.Interface
	Namespace		*corev1.Namespace
	cleanupHandle		CleanupActionHandle
}

func NewDefaultFramework(baseName string) *Framework {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := &Framework{BaseName: baseName}
	BeforeEach(f.BeforeEach)
	AfterEach(f.AfterEach)
	return f
}
func (f *Framework) BeforeEach() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	f.cleanupHandle = AddCleanupAction(f.AfterEach)
	By("Creating a kubernetes client")
	kubeConfig, err := kube.LoadConfig(TestContext.KubeConfig, TestContext.KubeContext)
	Expect(err).NotTo(HaveOccurred())
	kubeConfig.QPS = 50
	kubeConfig.Burst = 100
	f.KubeClientSet, err = kubernetes.NewForConfig(kubeConfig)
	Expect(err).NotTo(HaveOccurred())
	By("Creating a service catalog client")
	serviceCatalogConfig, err := kube.LoadConfig(TestContext.ServiceCatalogConfig, TestContext.ServiceCatalogContext)
	Expect(err).NotTo(HaveOccurred())
	serviceCatalogConfig.QPS = 50
	serviceCatalogConfig.Burst = 100
	f.ServiceCatalogClientSet, err = clientset.NewForConfig(serviceCatalogConfig)
	Expect(err).NotTo(HaveOccurred())
	By("Building a namespace api object")
	namespace, err := CreateKubeNamespace(f.BaseName, f.KubeClientSet)
	Expect(err).NotTo(HaveOccurred())
	f.Namespace = namespace
}
func (f *Framework) AfterEach() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	RemoveCleanupAction(f.cleanupHandle)
	err := DeleteKubeNamespace(f.KubeClientSet, f.Namespace.Name)
	Expect(err).NotTo(HaveOccurred())
}
func ServiceCatalogDescribe(text string, body func()) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return Describe("[service-catalog] "+text, body)
}
