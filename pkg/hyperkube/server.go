package hyperkube

import (
	"io/ioutil"
	"strings"
	utiltemplate "github.com/kubernetes-incubator/service-catalog/pkg/kubernetes/pkg/util/template"
	"k8s.io/apiserver/pkg/util/flag"
	"github.com/spf13/pflag"
)

type serverRunFunc func(s *Server, args []string, stopCh <-chan struct{}) error
type Server struct {
	SimpleUsage		string
	Long			string
	Run				serverRunFunc
	PrimaryName		string
	AlternativeName	string
	RespectsStopCh	bool
	flags			*pflag.FlagSet
	hk				*HyperKube
}

func (s *Server) Usage() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := `{{if .Long}}{{.Long | trim | wrap ""}}
{{end}}Usage:
  {{.SimpleUsage}} [flags]

Available Flags:
{{.Flags.FlagUsages}}`
	return utiltemplate.ExecuteTemplate(s.hk.Out(), tt, s)
}
func (s *Server) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.PrimaryName != "" {
		return s.PrimaryName
	}
	name := s.SimpleUsage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}
func (s *Server) Flags() *pflag.FlagSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if s.flags == nil {
		s.flags = pflag.NewFlagSet(s.Name(), pflag.ContinueOnError)
		s.flags.SetOutput(ioutil.Discard)
		s.flags.SetNormalizeFunc(flag.WordSepNormalizeFunc)
	}
	return s.flags
}
