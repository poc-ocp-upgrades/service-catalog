package version

import (
	"fmt"
	"runtime"
	apimachineryversion "k8s.io/apimachinery/pkg/version"
)

func Get() apimachineryversion.Info {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return apimachineryversion.Info{Major: "", Minor: "", GitVersion: gitVersion, GitCommit: gitCommit, GitTreeState: gitTreeState, BuildDate: buildDate, GoVersion: runtime.Version(), Compiler: runtime.Compiler, Platform: fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)}
}
