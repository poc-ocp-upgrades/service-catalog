package e2e

import (
	"time"
	v1beta1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"github.com/kubernetes-incubator/service-catalog/test/e2e/framework"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	instanceDeleteTimeout = 60 * time.Second
)

func newTestInstance(name, serviceClassName, planName string) *v1beta1.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &v1beta1.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: v1beta1.ServiceInstanceSpec{PlanReference: v1beta1.PlanReference{ClusterServicePlanExternalName: planName, ClusterServiceClassExternalName: serviceClassName}}}
}
func createInstance(c clientset.Interface, namespace string, instance *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServicecatalogV1beta1().ServiceInstances(namespace).Create(instance)
}
func deleteInstance(c clientset.Interface, namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return c.ServicecatalogV1beta1().ServiceInstances(namespace).Delete(name, nil)
}
func waitForInstanceToBeDeleted(c clientset.Interface, namespace, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return wait.Poll(framework.Poll, instanceDeleteTimeout, func() (bool, error) {
		_, err := c.ServicecatalogV1beta1().ServiceInstances(namespace).Get(name, metav1.GetOptions{})
		if err == nil {
			framework.Logf("waiting for service instance %s to be deleted", name)
			return false, nil
		}
		if errors.IsNotFound(err) {
			framework.Logf("verified service instance %s is deleted", name)
			return true, nil
		}
		return false, err
	})
}

var _ = framework.ServiceCatalogDescribe("ServiceInstance", func() {
	f := framework.NewDefaultFramework("service-instance")
	It("should verify an Instance can be deleted if referenced service class does not exist.", func() {
		By("Creating an Instance")
		instance := newTestInstance("test-instance", "no-service-class", "no-plan")
		instance, err := createInstance(f.ServiceCatalogClientSet, f.Namespace.Name, instance)
		Expect(err).NotTo(HaveOccurred())
		By("Deleting the Instance")
		err = deleteInstance(f.ServiceCatalogClientSet, f.Namespace.Name, instance.Name)
		Expect(err).NotTo(HaveOccurred())
		err = waitForInstanceToBeDeleted(f.ServiceCatalogClientSet, f.Namespace.Name, instance.Name)
		Expect(err).NotTo(HaveOccurred())
	})
})
