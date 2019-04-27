package pretty

import (
	"testing"
)

func TestPrettyContextBuilderKind(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetKind(ServiceInstance)
	e := "ServiceInstance"
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderNamespace(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetNamespace("Namespace")
	e := `"Namespace"`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetName("Name")
	e := `"Name"`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderKindAndNamespace(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetKind(ServiceInstance).SetNamespace("Namespace")
	e := `ServiceInstance "Namespace"`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderKindAndName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetKind(ServiceInstance).SetName("Name")
	e := `ServiceInstance "Name"`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderKindNamespaceName(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetKind(ServiceInstance).SetNamespace("Namespace").SetName("Name")
	e := `ServiceInstance "Namespace/Name"`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderKindNamespaceNameNew(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := NewContextBuilder(ServiceInstance, "Namespace", "Name", "")
	e := `ServiceInstance "Namespace/Name"`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderMessage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	e := `Msg`
	g := pcb.Message("Msg")
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderContextAndMessage(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetKind(ServiceInstance).SetNamespace("Namespace").SetName("Name")
	e := `ServiceInstance "Namespace/Name": Msg`
	g := pcb.Message("Msg")
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderMessagef(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	e := `This was built.`
	g := pcb.Messagef("This %s built.", "was")
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderMessagefMany(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	e := `One 2 three 4 "five" 6`
	g := pcb.Messagef("%s %d %s %v %q %d", "One", 2, "three", 4, "five", 6)
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderContextMessagefAndContext(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := ContextBuilder{}
	pcb.SetKind(ServiceInstance).SetNamespace("Namespace").SetName("Name")
	e := `ServiceInstance "Namespace/Name": This was the message: Msg`
	g := pcb.Messagef("This was the message: %s", "Msg")
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}
func TestPrettyContextBuilderNamespaceNameAndResourceVersion(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := NewContextBuilder(ServiceInstance, "Namespace", "Name", "877")
	e := `ServiceInstance "Namespace/Name" v877`
	g := pcb.String()
	if g != e {
		t.Fatalf("Unexpected value of ContextBuilder String; expected %v, got %v", e, g)
	}
}

var bResult string

func BenchmarkPCB(b *testing.B) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pcb := NewContextBuilder(ServiceInstance, "Namespace", "Name", "877")
	b.ResetTimer()
	var s string
	for i := 0; i <= b.N; i++ {
		s = pcb.String()
	}
	bResult = s
}
