package defaultserviceplan

import (
	"fmt"
	"strings"
	"testing"
	"time"
	"k8s.io/klog"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
	handler, err := NewDefaultClusterServicePlan()
	if err != nil {
		return nil, f, err
	}
	pluginInitializer := scadmission.NewPluginInitializer(internalClient, f, nil, nil)
	pluginInitializer.Initialize(handler)
	err = admission.ValidateInitialization(handler)
	return handler, f, err
}
func newFakeServiceCatalogClientForTest(sc *servicecatalog.ClusterServiceClass, sps []*servicecatalog.ClusterServicePlan, classFilter string) *fake.Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClient := &fake.Clientset{}
	fakeClient.AddReactor("get", "clusterserviceclasses", func(action core.Action) (bool, runtime.Object, error) {
		if sc != nil {
			return true, sc, nil
		}
		return true, nil, apierrors.NewNotFound(schema.GroupResource{}, "")
	})
	scList := &servicecatalog.ClusterServiceClassList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	if sc != nil {
		scList.Items = append(scList.Items, *sc)
	}
	fakeClient.AddReactor("list", "clusterserviceclasses", func(action core.Action) (bool, runtime.Object, error) {
		return true, scList, nil
	})
	spList := &servicecatalog.ClusterServicePlanList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	for _, sp := range sps {
		if classFilter == "" || classFilter == sp.Spec.ClusterServiceClassRef.Name {
			spList.Items = append(spList.Items, *sp)
		}
	}
	fakeClient.AddReactor("list", "clusterserviceplans", func(action core.Action) (bool, runtime.Object, error) {
		return true, spList, nil
	})
	return fakeClient
}
func newFakeServiceCatalogClientForNamespacedTest(sc *servicecatalog.ServiceClass, sps []*servicecatalog.ServicePlan, classFilter string) *fake.Clientset {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClient := &fake.Clientset{}
	fakeClient.AddReactor("get", "serviceclasses", func(action core.Action) (bool, runtime.Object, error) {
		if sc != nil {
			return true, sc, nil
		}
		return true, nil, apierrors.NewNotFound(schema.GroupResource{}, "")
	})
	scList := &servicecatalog.ServiceClassList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	if sc != nil {
		scList.Items = append(scList.Items, *sc)
	}
	fakeClient.AddReactor("list", "serviceclasses", func(action core.Action) (bool, runtime.Object, error) {
		return true, scList, nil
	})
	spList := &servicecatalog.ServicePlanList{ListMeta: metav1.ListMeta{ResourceVersion: "1"}}
	for _, sp := range sps {
		if classFilter == "" || classFilter == sp.Spec.ServiceClassRef.Name {
			spList.Items = append(spList.Items, *sp)
		}
	}
	fakeClient.AddReactor("list", "serviceplans", func(action core.Action) (bool, runtime.Object, error) {
		return true, spList, nil
	})
	return fakeClient
}
func newServiceInstance(namespace string) servicecatalog.ServiceInstance {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return servicecatalog.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Name: "instance", Namespace: namespace}}
}
func newClusterServiceClass(id string, name string) *servicecatalog.ClusterServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := &servicecatalog.ClusterServiceClass{ObjectMeta: metav1.ObjectMeta{Name: id}, Spec: servicecatalog.ClusterServiceClassSpec{CommonServiceClassSpec: servicecatalog.CommonServiceClassSpec{ExternalID: id, ExternalName: name}}}
	return sc
}
func newServiceClass(id string, name string) *servicecatalog.ServiceClass {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	sc := &servicecatalog.ServiceClass{ObjectMeta: metav1.ObjectMeta{Name: id}, Spec: servicecatalog.ServiceClassSpec{CommonServiceClassSpec: servicecatalog.CommonServiceClassSpec{ExternalID: id, ExternalName: name}}}
	return sc
}
func newClusterServicePlans(count uint, useDifferentClasses bool) []*servicecatalog.ClusterServicePlan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	classname := "test-serviceclass"
	sp1 := &servicecatalog.ClusterServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "bar-id"}, Spec: servicecatalog.ClusterServicePlanSpec{CommonServicePlanSpec: servicecatalog.CommonServicePlanSpec{ExternalName: "bar", ExternalID: "12345"}, ClusterServiceClassRef: servicecatalog.ClusterObjectReference{Name: classname}}}
	if useDifferentClasses {
		classname = "different-serviceclass"
	}
	sp2 := &servicecatalog.ClusterServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "baz-id"}, Spec: servicecatalog.ClusterServicePlanSpec{CommonServicePlanSpec: servicecatalog.CommonServicePlanSpec{ExternalName: "baz", ExternalID: "23456"}, ClusterServiceClassRef: servicecatalog.ClusterObjectReference{Name: classname}}}
	if 0 == count {
		return []*servicecatalog.ClusterServicePlan{}
	}
	if 1 == count {
		return []*servicecatalog.ClusterServicePlan{sp1}
	}
	if 2 == count {
		return []*servicecatalog.ClusterServicePlan{sp1, sp2}
	}
	return []*servicecatalog.ClusterServicePlan{}
}
func newServicePlans(count uint, useDifferentClasses bool) []*servicecatalog.ServicePlan {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	classname := "test-serviceclass"
	sp1 := &servicecatalog.ServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "bar-id"}, Spec: servicecatalog.ServicePlanSpec{CommonServicePlanSpec: servicecatalog.CommonServicePlanSpec{ExternalName: "bar", ExternalID: "12345"}, ServiceClassRef: servicecatalog.LocalObjectReference{Name: classname}}}
	if useDifferentClasses {
		classname = "different-serviceclass"
	}
	sp2 := &servicecatalog.ServicePlan{ObjectMeta: metav1.ObjectMeta{Name: "baz-id"}, Spec: servicecatalog.ServicePlanSpec{CommonServicePlanSpec: servicecatalog.CommonServicePlanSpec{ExternalName: "baz", ExternalID: "23456"}, ServiceClassRef: servicecatalog.LocalObjectReference{Name: classname}}}
	if 0 == count {
		return []*servicecatalog.ServicePlan{}
	}
	if 1 == count {
		return []*servicecatalog.ServicePlan{sp1}
	}
	if 2 == count {
		return []*servicecatalog.ServicePlan{sp1, sp2}
	}
	return []*servicecatalog.ServicePlan{}
}
func TestWithListFailure(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fakeClient := &fake.Clientset{}
	fakeClient.AddReactor("list", "clusterserviceclasses", func(action core.Action) (bool, runtime.Object, error) {
		return true, nil, fmt.Errorf("simulated test failure")
	})
	handler, informerFactory, err := newHandlerForTest(fakeClient)
	if err != nil {
		t.Errorf("unexpected error initializing handler: %v", err)
	}
	informerFactory.Start(wait.NeverStop)
	instance := newServiceInstance("dummy")
	instance.Spec.ClusterServiceClassExternalName = "foo"
	err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Create, false, nil))
	if err == nil {
		t.Errorf("unexpected success with no ClusterServiceClasses.List succeeding")
	} else if !strings.Contains(err.Error(), "simulated test failure") {
		t.Errorf("did not find expected error, got %q", err)
	}
	assertPlanReference(t, servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo"}, instance.Spec.PlanReference)
}
func TestWithPlanWorks(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		requestedPlan	servicecatalog.PlanReference
		namespaced	bool
	}{{"cluster external name", servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}, false}, {"cluster external id", servicecatalog.PlanReference{ClusterServiceClassExternalID: "foo", ClusterServicePlanExternalID: "12345"}, false}, {"cluster k8s", servicecatalog.PlanReference{ClusterServiceClassName: "foo-id", ClusterServicePlanName: "bar-id"}, false}, {"ns external name", servicecatalog.PlanReference{ServiceClassExternalName: "foo", ServicePlanExternalName: "bar"}, true}, {"ns external id", servicecatalog.PlanReference{ServiceClassExternalID: "foo", ServicePlanExternalID: "bar"}, true}, {"ns k8s", servicecatalog.PlanReference{ServiceClassName: "foo", ServicePlanName: "bar"}, true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var fakeClient *fake.Clientset
			if tc.namespaced {
				fakeClient = newFakeServiceCatalogClientForNamespacedTest(nil, newServicePlans(1, false), "")
			} else {
				fakeClient = newFakeServiceCatalogClientForTest(nil, newClusterServicePlans(1, false), "")
			}
			handler, informerFactory, err := newHandlerForTest(fakeClient)
			if err != nil {
				t.Errorf("unexpected error initializing handler: %v", err)
			}
			informerFactory.Start(wait.NeverStop)
			instance := newServiceInstance("dummy")
			instance.Spec.PlanReference = tc.requestedPlan
			err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Create, false, nil))
			if err != nil {
				actions := ""
				for _, action := range fakeClient.Actions() {
					actions = actions + action.GetVerb() + ":" + action.GetResource().Resource + ":" + action.GetSubresource() + ", "
				}
				t.Errorf("unexpected error %q returned from admission handler: %v", err, actions)
			}
			assertPlanReference(t, tc.requestedPlan, instance.Spec.PlanReference)
		})
	}
}
func TestWithNoPlanFailsWithNoClass(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		requestedPlan	servicecatalog.PlanReference
		namespaced	bool
	}{{"cluster external name", servicecatalog.PlanReference{ClusterServiceClassExternalName: "bad-class"}, false}, {"cluster external id", servicecatalog.PlanReference{ClusterServiceClassExternalID: "bad-class"}, false}, {"cluster k8s", servicecatalog.PlanReference{ClusterServiceClassName: "bad-class"}, false}, {"ns external name", servicecatalog.PlanReference{ServiceClassExternalName: "bad-class"}, true}, {"ns external id", servicecatalog.PlanReference{ServiceClassExternalID: "bad-class"}, true}, {"ns k8s", servicecatalog.PlanReference{ServiceClassName: "bad-class"}, true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var fakeClient *fake.Clientset
			if tc.namespaced {
				fakeClient = newFakeServiceCatalogClientForNamespacedTest(nil, newServicePlans(1, false), "")
			} else {
				fakeClient = newFakeServiceCatalogClientForTest(nil, newClusterServicePlans(1, false), "")
			}
			handler, informerFactory, err := newHandlerForTest(fakeClient)
			if err != nil {
				t.Errorf("unexpected error initializing handler: %v", err)
			}
			informerFactory.Start(wait.NeverStop)
			instance := newServiceInstance("dummy")
			instance.Spec.PlanReference = tc.requestedPlan
			err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Create, false, nil))
			if err == nil {
				t.Errorf("unexpected success with no plan specified and no serviceclass existing")
			} else if !strings.Contains(err.Error(), "does not exist, can not figure") {
				t.Errorf("did not find expected error, got %q", err)
			}
		})
	}
}
func TestWithNoPlanWorksWithSinglePlan(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		requestedPlan	servicecatalog.PlanReference
		resolvedPlan	servicecatalog.PlanReference
		namespaced	bool
	}{{"cluster external name", servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo"}, servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}, false}, {"cluster external id", servicecatalog.PlanReference{ClusterServiceClassExternalID: "foo"}, servicecatalog.PlanReference{ClusterServiceClassExternalID: "foo", ClusterServicePlanExternalID: "12345"}, false}, {"cluster k8s", servicecatalog.PlanReference{ClusterServiceClassName: "foo-id"}, servicecatalog.PlanReference{ClusterServiceClassName: "foo-id", ClusterServicePlanName: "bar-id"}, false}, {"ns external name", servicecatalog.PlanReference{ServiceClassExternalName: "foo"}, servicecatalog.PlanReference{ServiceClassExternalName: "foo", ServicePlanExternalName: "bar"}, true}, {"ns external id", servicecatalog.PlanReference{ServiceClassExternalID: "foo"}, servicecatalog.PlanReference{ServiceClassExternalID: "foo", ServicePlanExternalID: "12345"}, true}, {"ns k8s", servicecatalog.PlanReference{ServiceClassName: "foo-id"}, servicecatalog.PlanReference{ServiceClassName: "foo-id", ServicePlanName: "bar-id"}, true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var fakeClient *fake.Clientset
			if tc.namespaced {
				sc := newServiceClass("foo-id", "foo")
				sps := newServicePlans(1, false)
				klog.V(4).Infof("Created Service as %+v", sc)
				fakeClient = newFakeServiceCatalogClientForNamespacedTest(sc, sps, "")
			} else {
				csc := newClusterServiceClass("foo-id", "foo")
				csps := newClusterServicePlans(1, false)
				klog.V(4).Infof("Created Service as %+v", csc)
				fakeClient = newFakeServiceCatalogClientForTest(csc, csps, "")
			}
			handler, informerFactory, err := newHandlerForTest(fakeClient)
			if err != nil {
				t.Errorf("unexpected error initializing handler: %v", err)
			}
			informerFactory.Start(wait.NeverStop)
			instance := newServiceInstance("dummy")
			instance.Spec.PlanReference = tc.requestedPlan
			err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Create, false, nil))
			if err != nil {
				actions := ""
				for _, action := range fakeClient.Actions() {
					actions = actions + action.GetVerb() + ":" + action.GetResource().Resource + ":" + action.GetSubresource() + ", "
				}
				t.Errorf("unexpected error %q returned from admission handler: %v", err, actions)
			}
			assertPlanReference(t, tc.resolvedPlan, instance.Spec.PlanReference)
		})
	}
}
func TestWithNoPlanFailsWithMultiplePlans(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		requestedPlan	servicecatalog.PlanReference
		namespaced	bool
	}{{"cluster external name", servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo"}, false}, {"cluster external id", servicecatalog.PlanReference{ClusterServiceClassExternalID: "foo"}, false}, {"cluster k8s", servicecatalog.PlanReference{ClusterServiceClassName: "foo-id"}, false}, {"ns external name", servicecatalog.PlanReference{ServiceClassExternalName: "foo"}, true}, {"ns external id", servicecatalog.PlanReference{ServiceClassExternalID: "foo"}, true}, {"ns k8s", servicecatalog.PlanReference{ServiceClassName: "foo-id"}, true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var fakeClient *fake.Clientset
			if tc.namespaced {
				sc := newServiceClass("foo-id", "foo")
				sps := newServicePlans(2, false)
				klog.V(4).Infof("Created Service as %+v", sc)
				fakeClient = newFakeServiceCatalogClientForNamespacedTest(sc, sps, "")
			} else {
				csc := newClusterServiceClass("foo-id", "foo")
				csps := newClusterServicePlans(2, false)
				klog.V(4).Infof("Created Service as %+v", csc)
				fakeClient = newFakeServiceCatalogClientForTest(csc, csps, "")
			}
			handler, informerFactory, err := newHandlerForTest(fakeClient)
			if err != nil {
				t.Errorf("unexpected error initializing handler: %v", err)
			}
			informerFactory.Start(wait.NeverStop)
			instance := newServiceInstance("dummy")
			instance.Spec.PlanReference = tc.requestedPlan
			err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Create, false, nil))
			if err == nil {
				t.Errorf("unexpected success with no plan specified and no serviceclass existing")
				return
			} else if !strings.Contains(err.Error(), "has more than one plan, PlanName must be") {
				t.Errorf("did not find expected error, got %q", err)
			}
		})
	}
}
func TestWithNoPlanSucceedsWithMultiplePlansFromDifferentClasses(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	cases := []struct {
		name		string
		requestedPlan	servicecatalog.PlanReference
		resolvedPlan	servicecatalog.PlanReference
		namespaced	bool
	}{{"cluster external name", servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo"}, servicecatalog.PlanReference{ClusterServiceClassExternalName: "foo", ClusterServicePlanExternalName: "bar"}, false}, {"cluster external id", servicecatalog.PlanReference{ClusterServiceClassExternalID: "foo"}, servicecatalog.PlanReference{ClusterServiceClassExternalID: "foo", ClusterServicePlanExternalID: "12345"}, false}, {"cluster k8s", servicecatalog.PlanReference{ClusterServiceClassName: "foo-id"}, servicecatalog.PlanReference{ClusterServiceClassName: "foo-id", ClusterServicePlanName: "bar-id"}, false}, {"ns external name", servicecatalog.PlanReference{ServiceClassExternalName: "foo"}, servicecatalog.PlanReference{ServiceClassExternalName: "foo", ServicePlanExternalName: "bar"}, true}, {"ns external id", servicecatalog.PlanReference{ServiceClassExternalID: "foo"}, servicecatalog.PlanReference{ServiceClassExternalID: "foo", ServicePlanExternalID: "12345"}, true}, {"ns k8s", servicecatalog.PlanReference{ServiceClassName: "foo-id"}, servicecatalog.PlanReference{ServiceClassName: "foo-id", ServicePlanName: "bar-id"}, true}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var fakeClient *fake.Clientset
			classFilter := "test-serviceclass"
			if tc.namespaced {
				sc := newServiceClass("foo-id", "foo")
				sps := newServicePlans(2, true)
				klog.V(4).Infof("Created Service as %+v", sc)
				fakeClient = newFakeServiceCatalogClientForNamespacedTest(sc, sps, classFilter)
			} else {
				csc := newClusterServiceClass("foo-id", "foo")
				csps := newClusterServicePlans(2, true)
				klog.V(4).Infof("Created Service as %+v", csc)
				fakeClient = newFakeServiceCatalogClientForTest(csc, csps, classFilter)
			}
			handler, informerFactory, err := newHandlerForTest(fakeClient)
			if err != nil {
				t.Errorf("unexpected error initializing handler: %v", err)
			}
			informerFactory.Start(wait.NeverStop)
			instance := newServiceInstance("dummy")
			instance.Spec.PlanReference = tc.requestedPlan
			err = handler.(admission.MutationInterface).Admit(admission.NewAttributesRecord(&instance, nil, servicecatalog.Kind("ServiceInstance").WithVersion("version"), instance.Namespace, instance.Name, servicecatalog.Resource("serviceinstances").WithVersion("version"), "", admission.Create, false, nil))
			if err != nil {
				actions := ""
				for _, action := range fakeClient.Actions() {
					actions = actions + action.GetVerb() + ":" + action.GetResource().Resource + ":" + action.GetSubresource() + ", "
				}
				t.Errorf("unexpected error %q returned from admission handler: %v", err, actions)
			}
			assertPlanReference(t, tc.resolvedPlan, instance.Spec.PlanReference)
		})
	}
}
func assertPlanReference(t *testing.T, expected servicecatalog.PlanReference, actual servicecatalog.PlanReference) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if expected != actual {
		t.Errorf("PlanReference was not as expected: %+v actual: %+v", expected, actual)
	}
}
