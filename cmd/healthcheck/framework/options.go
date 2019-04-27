package framework

import (
	"os"
	"time"
	"github.com/spf13/pflag"
	"k8s.io/client-go/tools/clientcmd"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type HealthCheckServer struct {
	KubeHost		string
	KubeConfig		string
	KubeContext		string
	HealthCheckInterval	time.Duration
	SecureServingOptions	*genericoptions.SecureServingOptions
	TestBrokerName		string
}

const (
	defaultHealthCheckInterval	= 2 * time.Minute
	defaultSecurePort		= 443
	defaultCertDirectory		= "/var/run/service-catalog-healthcheck"
)

func NewHealthCheckServer() *HealthCheckServer {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := HealthCheckServer{HealthCheckInterval: defaultHealthCheckInterval, SecureServingOptions: genericoptions.NewSecureServingOptions()}
	s.SecureServingOptions.BindPort = defaultSecurePort
	s.SecureServingOptions.ServerCert.CertDirectory = defaultCertDirectory
	return &s
}
func (s *HealthCheckServer) AddFlags(fs *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	fs.StringVar(&s.KubeHost, "kubernetes-host", "http://127.0.0.1:8080", "The kubernetes host, or apiserver, to connect to")
	fs.StringVar(&s.KubeConfig, "kubernetes-config", os.Getenv(clientcmd.RecommendedConfigPathEnvVar), "Path to config containing embedded authinfo for kubernetes. Default value is from environment variable "+clientcmd.RecommendedConfigPathEnvVar)
	fs.StringVar(&s.KubeContext, "kubernetes-context", "", "config context to use for kubernetes. If unset, will use value from 'current-context'")
	fs.DurationVar(&s.HealthCheckInterval, "healthcheck-interval", s.HealthCheckInterval, "How frequently the end to end health check should be performed")
	fs.StringVar(&s.TestBrokerName, "broker-name", "ups-broker", "Broker Name to test against - can only be ups-broker or osb-stub.  You must ensure the specified broker is deployed.")
	s.SecureServingOptions.AddFlags(fs)
}
