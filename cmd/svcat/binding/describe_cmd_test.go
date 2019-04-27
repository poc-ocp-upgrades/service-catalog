package binding

import (
	"bytes"
	"strings"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/test"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1beta1"
	svcatfake "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/fake"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	_ "github.com/kubernetes-incubator/service-catalog/internal/test"
)

func TestDescribeCommand(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	const namespace = "default"
	testcases := []struct {
		name		string
		fakeBindings	[]string
		bindingName	string
		expectedError	string
		wantError	bool
	}{{name: "describe non existing binding", fakeBindings: []string{}, bindingName: "mybinding", expectedError: "unable to get binding '" + namespace + ".mybinding'", wantError: true}, {name: "describe existing binding", fakeBindings: []string{"mybinding"}, bindingName: "mybinding", wantError: false}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			k8sClient := k8sfake.NewSimpleClientset()
			var fakes []runtime.Object
			for _, name := range tc.fakeBindings {
				fakes = append(fakes, &v1beta1.ServiceBinding{ObjectMeta: v1.ObjectMeta{Namespace: namespace, Name: name}, Spec: v1beta1.ServiceBindingSpec{}})
			}
			svcatClient := svcatfake.NewSimpleClientset(fakes...)
			fakeApp, _ := svcat.NewApp(k8sClient, svcatClient, namespace)
			output := &bytes.Buffer{}
			cxt := svcattest.NewContext(output, fakeApp)
			cmd := &describeCmd{Namespaced: command.NewNamespaced(cxt)}
			cmd.Namespace = namespace
			cmd.name = tc.bindingName
			err := cmd.Run()
			if tc.wantError {
				if err == nil {
					t.Errorf("expected a non-zero exit code, but the command succeeded")
				}
				errorOutput := err.Error()
				if !strings.Contains(errorOutput, tc.expectedError) {
					t.Errorf("Unexpected output:\n\nExpected:\n%q\n\nActual:\n%q\n", tc.expectedError, errorOutput)
				}
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected the command to succeed but it failed with %q", err)
			}
		})
	}
}
