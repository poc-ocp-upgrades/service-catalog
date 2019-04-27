package meta

import (
	"testing"
	sc "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetNamespace(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	const namespace = "testns"
	obj := &sc.ServiceInstance{ObjectMeta: metav1.ObjectMeta{Namespace: namespace}}
	ns, err := GetNamespace(obj)
	if err != nil {
		t.Fatalf("error getting namespace (%s)", err)
	}
	if ns != namespace {
		t.Fatalf("actual namespace (%s) wasn't expected (%s)", ns, namespace)
	}
}
