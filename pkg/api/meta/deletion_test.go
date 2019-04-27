package meta

import (
	"testing"
	"time"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeletionTimestampExists(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	obj := &sc.ServiceInstance{ObjectMeta: metav1.ObjectMeta{}}
	exists, err := DeletionTimestampExists(obj)
	if err != nil {
		t.Fatal(err)
	}
	if exists {
		t.Fatalf("deletion timestamp reported as exists when it didn't")
	}
	tme := metav1.NewTime(time.Now())
	obj.DeletionTimestamp = &tme
	exists, err = DeletionTimestampExists(obj)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Fatal("deletion timestamp reported as missing when it isn't")
	}
}
func TestRoundTripDeletionTimestamp(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	t1 := metav1.NewTime(time.Now())
	t2 := metav1.NewTime(time.Now().Add(1 * time.Hour))
	obj := &sc.ServiceInstance{ObjectMeta: metav1.ObjectMeta{DeletionTimestamp: &t1}}
	t1Ret, err := GetDeletionTimestamp(obj)
	if err != nil {
		t.Fatalf("error getting 1st deletion timestamp (%s)", err)
	}
	if !t1.Equal(t1Ret) {
		t.Fatalf("expected deletion timestamp %s, got %s", t1, *t1Ret)
	}
	if err := SetDeletionTimestamp(obj, t2.Time); err != nil {
		t.Fatalf("error setting deletion timestamp (%s)", err)
	}
	t2Ret, err := GetDeletionTimestamp(obj)
	if err != nil {
		t.Fatalf("error getting 2nd deletion timestamp (%s)", err)
	}
	if !t2.Equal(t2Ret) {
		t.Fatalf("expected deletion timestamp %s, got %s", t2, *t2Ret)
	}
}
