package output

import (
	"bytes"
	"testing"
	_ "github.com/kubernetes-incubator/service-catalog/internal/test"
	"k8s.io/apimachinery/pkg/runtime"
)

func TestWriteParameters(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	testcases := []struct {
		name		string
		parameters	*runtime.RawExtension
		output		string
	}{{"Nil parameter", nil, "\nParameters:\n  No parameters defined\n"}, {"JSON w/data parameter", &runtime.RawExtension{Raw: []byte(`{"foo":"bar"}`)}, "\nParameters:\n  foo: bar\n"}, {"JSON empty parameter", &runtime.RawExtension{Raw: []byte(`{}`)}, "\nParameters:\n  No parameters defined\n"}, {"String parameter", &runtime.RawExtension{Raw: []byte("param")}, "\nParameters:\nparam\n"}, {"Empty string parameter", &runtime.RawExtension{Raw: []byte("")}, "\nParameters:\n  No parameters defined\n"}}
	for _, tc := range testcases {
		output := &bytes.Buffer{}
		writeParameters(output, tc.parameters)
		if tc.output != output.String() {
			t.Errorf("%v: Output mismatch: expected \"%v\", actual \"%v\"", tc.name, tc.output, output.String())
		}
	}
}
