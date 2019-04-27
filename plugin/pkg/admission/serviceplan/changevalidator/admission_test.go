package changevalidator

import (
	"strings"
	"testing"
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/admission"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scadmission "github.com/kubernetes-incubator/service-catalog/pkg/apiserver/admission"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/fake"
	informers "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion"
	core "k8s.io/client-go/testing"
)

func newHandlerForTest(internalClient internalclientset.Interface) (admission.Interface, informers.SharedInformerFactory, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := informers.NewSharedInformerFactory(internalClient, 5*time.Minute)
	handler, err := NewDenyPlanChangeIfNotUpdatable()
	if err != nil {
		return nil, f, err
	}
	pluginInitializer := scadmission.NewPluginInitializer(internalClient, f, nil, nil)
	pluginInitializer.Initialize(handler)
	err = admission.ValidateInitialization(handler)
	return handler, f, err
}
func newFakeServiceCatalogClientForTest(sc *servicecatalog.ClusterServiceClass) *fake.Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClient := &fake.Clientset{}
	scList := &servicecatalog.ClusterServiceClassList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	scList.Items = append(scList.Items, *sc)
	fakeClient.AddReactor("list", "clusterserviceclasses", func(action core.Action) (bool, runtime.Object, error) {
		return true, scList, nil
	})
	return fakeClient
}
func newServiceInstance(namespace string, serviceClassName string, planName string) servicecatalog.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance := servicecatalog.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: "instance", Namespace: namespace}, Spec: servicecatalog.ServiceInstanceSpec{PlanReference: servicecatalog.PlanReference{ClusterServicePlanExternalName: planName}, ClusterServiceClassRef: &servicecatalog.ClusterObjectReference{Name: serviceClassName}}}
	return instance
}
func newClusterServiceClass(name string, plan string, updateablePlan bool) *servicecatalog.ClusterServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := &servicecatalog.ClusterServiceClass{ObjectMeta: metav1.ObjectMeta{Name: name}, Spec: servicecatalog.ClusterServiceClassSpec{CommonServiceClassSpec: servicecatalog.CommonServiceClassSpec{PlanUpdatable: updateablePlan}}}
	return sc
}
func setupInstanceLister(fakeClient *fake.Clientset) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	instance := newServiceInstance("dummy", "foo", "original-plan-name")
	scList := &servicecatalog.ServiceInstanceList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	scList.Items = append(scList.Items, instance)
	fakeClient.AddReactor("list", "serviceinstances", func(action core.Action) (bool, runtime.Object, error) {
		return true, scList, nil
	})
}
func TestClusterServicePlanChangeBlockedByUpdateablePlanSetting(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := newClusterServiceClass("foo", "bar", false)
	fakeClient := newFakeServiceCatalogClientForTest(sc)
	handler, informerFactory, err := newHandlerForTest(fakeClient)
	if err != nil {
		t.Errorf("unexpected error initializing handler: %v", err)
	}
	setupInstanceLister(fakeClient)
	instance := newServiceInstance("dummy", "foo", "new-plan")
	informerFactory.Start(wait.NeverStop)
	err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Update, false, nil))
	if err != nil {
		if !strings.Contains(err.Error(), "The Service Class foo does not allow plan changes.") {
			t.Errorf("unexpected error %q returned from admission handler.", err.Error())
		}
	} else {
		t.Error("This should have been an error")
	}
}
func TestClusterServicePlanChangePermittedByUpdateablePlanSetting(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := newClusterServiceClass("foo", "bar", true)
	fakeClient := newFakeServiceCatalogClientForTest(sc)
	handler, informerFactory, err := newHandlerForTest(fakeClient)
	if err != nil {
		t.Errorf("unexpected error initializing handler: %v", err)
	}
	setupInstanceLister(fakeClient)
	instance := newServiceInstance("dummy", "foo", "new-plan")
	informerFactory.Start(wait.NeverStop)
	err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Update, false, nil))
	if err != nil {
		t.Errorf("Unexpected error: %v", err.Error())
	}
}
