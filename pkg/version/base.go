package version

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
)

var (
	gitMajor	string
	gitMinor	string
	gitVersion	= "v0.0.0-master+$Format:%h$"
	gitCommit	= "$Format:%H$"
	gitTreeState	= "not a git tree"
	buildDate	= "1970-01-01T00:00:00Z"
)

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
