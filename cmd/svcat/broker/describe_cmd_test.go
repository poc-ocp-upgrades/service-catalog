package broker

import (
	"bytes"
	"strings"
	"testing"
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
		fakeBrokers	[]string
		brokerName	string
		expectedError	string
		wantError	bool
	}{{name: "describe non existing broker", fakeBrokers: []string{}, brokerName: "mybroker", expectedError: "unable to get broker 'mybroker'", wantError: true}, {name: "describe existing broker", fakeBrokers: []string{"mybroker"}, brokerName: "mybroker", wantError: false}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			k8sClient := k8sfake.NewSimpleClientset()
			var fakes []runtime.Object
			for _, name := range tc.fakeBrokers {
				fakes = append(fakes, &v1beta1.ClusterServiceBroker{ObjectMeta: v1.ObjectMeta{Name: name}, Spec: v1beta1.ClusterServiceBrokerSpec{}})
			}
			svcatClient := svcatfake.NewSimpleClientset(fakes...)
			fakeApp, _ := svcat.NewApp(k8sClient, svcatClient, namespace)
			output := &bytes.Buffer{}
			cxt := svcattest.NewContext(output, fakeApp)
			cmd := &describeCmd{Context: cxt}
			cmd.name = tc.brokerName
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
