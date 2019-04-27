package lifecycle

import (
	"testing"
	"time"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/admission"
	core "k8s.io/client-go/testing"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	scadmission "github.com/kubernetes-incubator/service-catalog/pkg/apiserver/admission"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset"
	"github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/internalclientset/fake"
	informers "github.com/kubernetes-incubator/service-catalog/pkg/client/informers_generated/internalversion"
)

func newHandlerForTest(internalClient internalclientset.Interface) (admission.Interface, informers.SharedInformerFactory, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	f := informers.NewSharedInformerFactory(internalClient, 5*time.Minute)
	handler, err := NewCredentialsBlocker()
	if err != nil {
		return nil, f, err
	}
	pluginInitializer := scadmission.NewPluginInitializer(internalClient, f, nil, nil)
	pluginInitializer.Initialize(handler)
	err = admission.ValidateInitialization(handler)
	return handler, f, err
}
func newServiceInstance() servicecatalog.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return servicecatalog.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: "test-instance", Namespace: "test-ns"}}
}
func newServiceBinding() servicecatalog.ServiceBinding {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return servicecatalog.ServiceBinding{ObjectMeta: metav1.ObjectMeta{Name: "test-cred", Namespace: "test-ns"}, Spec: servicecatalog.ServiceBindingSpec{InstanceRef: servicecatalog.LocalObjectReference{Name: "test-instance"}, SecretName: "test-secret"}}
}
func TestBlockNewCredentialsForDeletedInstance(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClient := &fake.Clientset{}
	handler, informerFactory, err := newHandlerForTest(fakeClient)
	if err != nil {
		t.Errorf("unexpected error initializing handler: %v", err)
	}
	instance := newServiceInstance()
	instance.DeletionTimestamp = &metav1.Time{}
	scList := &servicecatalog.ServiceInstanceList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	scList.Items = append(scList.Items, instance)
	fakeClient.AddReactor("list", "serviceinstances", func(action core.Action) (bool, runtime.Object, error) {
		return true, scList, nil
	})
	credential := newServiceBinding()
	informerFactory.Start(wait.NeverStop)
	err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&credential, nil, servicecatalog.Kind("ServiceBindings").WithVersion("version"), "test-ns", "test-cred", servicecatalog.Resource("servicebindings").WithVersion("version"), "", admission.Create, false, nil))
	if err == nil {
		t.Error("Unexpected error: admission controller failed blocking the request")
	} else {
		if err.Error() != "servicebindings.servicecatalog.k8s.io \"test-cred\" is forbidden: ServiceBinding test-ns/test-cred references a ServiceInstance that is being deleted: test-ns/test-instance" {
			t.Fatalf("admission controller blocked the request but not with expected error, expected a forbidden error, got %q", err.Error())
		}
	}
}
func TestAllowNewCredentialsForNonDeletedInstance(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClient := &fake.Clientset{}
	handler, informerFactory, err := newHandlerForTest(fakeClient)
	if err != nil {
		t.Errorf("unexpected error initializing handler: %v", err)
	}
	informerFactory.Start(wait.NeverStop)
	credential := newServiceBinding()
	err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&credential, nil, servicecatalog.Kind("ServiceBindings").WithVersion("version"), "test-ns", "test-cred", servicecatalog.Resource("servicebindings").WithVersion("version"), "", admission.Create, false, nil))
	if err != nil {
		t.Errorf("Error, admission controller should not block this test")
	}
}
