package controller

import (
	"sync"
	"testing"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	clientgotesting "k8s.io/client-go/testing"
)

func TestGetClusterIDDefaulting(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, _, tc, _ := newTestController(t, noFakeActions())
	tc.setClusterID("")
	if tc.getClusterID() == "" {
		t.Fatalf("cluster id should have been generated and filled in upon request")
	}
	t.Log(tc.getClusterID())
}
func TestGetClusterIDConcurrently(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, _, tc, _ := newTestController(t, noFakeActions())
	tc.setClusterID("")
	var a, b string
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		a = tc.getClusterID()
		wg.Done()
	}()
	go func() {
		b = tc.getClusterID()
		wg.Done()
	}()
	wg.Wait()
	if a != b {
		t.Fatal("a and b should match", a, b)
	}
	if tc.getClusterID() == "" {
		t.Fatalf("cluster id should have been generated and filled in upon request")
	}
	t.Log(tc.getClusterID())
}
func TestGetClusterIDRoundTrip(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, _, _, tc, _ := newTestController(t, noFakeActions())
	tc.setClusterID("")
	tc.setClusterID(testClusterID)
	if tc.getClusterID() != testClusterID {
		t.Fatalf("should have got the same string out that we put in")
	}
}
func TestMonitorConfigMapNoConfigmap(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kc, _, _, tc, _ := newTestController(t, noFakeActions())
	kc.AddReactor("get", "configmaps", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		m := make(map[string]string)
		m["id"] = testClusterID
		return true, nil, errors.NewNotFound(schema.GroupResource{Group: "core", Resource: "configmap"}, DefaultClusterIDConfigMapName)
	})
	tc.setClusterID(testClusterID)
	tc.monitorConfigMap()
	if tc.getClusterID() != testClusterID {
		t.Fatalf("should have got the same string out that we put in")
	}
	if kc.Actions()[1].Matches("create", "configmaps") {
		createdCM := kc.Actions()[1].(clientgotesting.CreateAction).GetObject().(*corev1.ConfigMap)
		if id, ok := createdCM.Data["id"]; !(ok && id == testClusterID) {
			t.Fatalf("new configmap should have id as existing testClusterID. Had id %q", id)
		}
	} else {
		t.Fatalf("should have created a new configmap")
	}
}
func TestMonitorConfigMapNoConfigmapNoExistingClusterID(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kc, _, _, tc, _ := newTestController(t, noFakeActions())
	kc.AddReactor("get", "configmaps", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		m := make(map[string]string)
		m["id"] = testClusterID
		return true, nil, errors.NewNotFound(schema.GroupResource{Group: "core", Resource: "configmap"}, DefaultClusterIDConfigMapName)
	})
	tc.setClusterID("")
	tc.monitorConfigMap()
	if tc.getClusterID() == "" {
		t.Fatalf("cluster id should have been generated and filled in upon request")
	}
	if kc.Actions()[1].Matches("create", "configmaps") {
		createdCM := kc.Actions()[1].(clientgotesting.CreateAction).GetObject().(*corev1.ConfigMap)
		if id, ok := createdCM.Data["id"]; !(ok && id != "") {
			t.Fatalf("new configmap should have a non-blank id")
		}
	} else {
		t.Fatalf("should have created a new configmap")
	}
}
func TestMonitorConfigMapConfigmapOverride(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kc, _, _, tc, _ := newTestController(t, noFakeActions())
	kc.AddReactor("get", "configmaps", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		m := make(map[string]string)
		m["id"] = testClusterID
		return true, &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: DefaultClusterIDConfigMapName}, Data: m}, nil
	})
	tc.setClusterID("non-cluster-id")
	tc.monitorConfigMap()
	if tc.getClusterID() != testClusterID {
		t.Fatalf("should have got the override id from the configmap")
	}
}
func TestMonitorConfigMapConfigmapWithNoData(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kc, _, _, tc, _ := newTestController(t, noFakeActions())
	blankcm := &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: DefaultClusterIDConfigMapName}}
	kc.AddReactor("get", "configmaps", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		return true, blankcm, nil
	})
	tc.setClusterID(testClusterID)
	tc.monitorConfigMap()
	if tc.getClusterID() != testClusterID {
		t.Fatalf("should have got the set cluster id")
	}
	if expectedCMget := kc.Actions()[0]; expectedCMget.GetVerb() != "get" {
		t.Fatalf("get configmap is first")
	}
	if expectedCMupdate := kc.Actions()[1]; expectedCMupdate.GetVerb() == "update" {
		updatedCM := expectedCMupdate.(clientgotesting.UpdateAction).GetObject().(*corev1.ConfigMap)
		if id := updatedCM.Data["id"]; id != testClusterID {
			t.Fatalf("configmap should have been updated with the existing clusterid, was %q, expected %q", id, testClusterID)
		}
	} else {
		t.Fatalf("configmap should have been updated with the existing clusterid")
	}
}
func TestMonitorConfigMapConfigmapWithOtherData(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	kc, _, _, tc, _ := newTestController(t, noFakeActions())
	kc.AddReactor("get", "configmaps", func(action clientgotesting.Action) (bool, runtime.Object, error) {
		m := make(map[string]string)
		m["notid"] = "other-non-id-stuff-that-needs-to-be-perserved"
		return true, &corev1.ConfigMap{ObjectMeta: metav1.ObjectMeta{Name: DefaultClusterIDConfigMapName}, Data: m}, nil
	})
	tc.setClusterID(testClusterID)
	tc.monitorConfigMap()
	if tc.getClusterID() != testClusterID {
		t.Fatalf("should have got the set cluster id")
	}
	if expectedCMget := kc.Actions()[0]; expectedCMget.GetVerb() != "get" {
		t.Fatalf("get configmap is first")
	}
	if expectedCMupdate := kc.Actions()[1]; expectedCMupdate.GetVerb() == "update" {
		updatedCM := expectedCMupdate.(clientgotesting.UpdateAction).GetObject().(*corev1.ConfigMap)
		if notid := updatedCM.Data["notid"]; notid == "" {
			t.Fatalf("configmap should have another key")
		}
		if id := updatedCM.Data["id"]; id != testClusterID {
			t.Fatalf("configmap should have been updated with the existing clusterid")
		}
	} else {
		t.Fatalf("configmap should have been updated with the existing clusterid")
	}
}
