package fake

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	"k8s.io/apimachinery/pkg/util/rand"
	testing "k8s.io/client-go/testing"
)

func (c *FakeServiceInstances) UpdateReferences(serviceInstance *v1beta1.ServiceInstance) (*v1beta1.ServiceInstance, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instanceCopy := serviceInstance.DeepCopy()
	obj, err := c.Fake.Invokes(testing.NewUpdateSubresourceAction(serviceinstancesResource, "reference", c.ns, instanceCopy), instanceCopy)
	if obj == nil {
		return nil, err
	}
	updatedInstance := (obj.(*v1beta1.ServiceInstance))
	updatedInstance.ResourceVersion = rand.String(10)
	return updatedInstance, err
}
