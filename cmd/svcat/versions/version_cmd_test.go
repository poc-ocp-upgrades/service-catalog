package versions

import (
	"bytes"
	"strings"
	"testing"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/pkg"
	svcatfake "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset/fake"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat"
	k8sfake "k8s.io/client-go/kubernetes/fake"
	_ "github.com/kubernetes-incubator/service-catalog/internal/test"
)

func TestVersionCommand(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pkg.VERSION = "v0.0.0"
	testcases := []struct {
		name		string
		client		bool
		server		bool
		wantOutput	string
		wantError	bool
	}{{name: "show client version only", client: true, server: false, wantOutput: "Client Version: v0.0.0\n", wantError: false}, {name: "show server & client version", client: true, server: true, wantOutput: "Client Version: v0.0.0\nServer Version: v0.0.0-master+$Format:%h$\n", wantError: false}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			k8sClient := k8sfake.NewSimpleClientset()
			svcatClient := svcatfake.NewSimpleClientset()
			output := &bytes.Buffer{}
			fakeApp, _ := svcat.NewApp(k8sClient, svcatClient, "default")
			cxt := &command.Context{Output: output, App: fakeApp}
			versionCommand := &versionCmd{cxt, tc.client, tc.server}
			err := versionCommand.Run()
			if tc.wantError && err == nil {
				t.Errorf("expected a non-zero exit code, but the command succeeded")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected the command to succeed but it failed with %q", err)
			}
			gotOutput := output.String()
			if err != nil {
				gotOutput += err.Error()
			}
			if !strings.Contains(gotOutput, tc.wantOutput) {
				t.Errorf("unexpected output \n\nWANT:\n%q\n\nGOT:\n%q\n", tc.wantOutput, gotOutput)
			}
		})
	}
}
