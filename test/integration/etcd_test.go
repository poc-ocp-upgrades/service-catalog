package integration

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
	"github.com/coreos/etcd/embed"
	"github.com/coreos/pkg/capnslog"
	"k8s.io/klog"
)

type EtcdContext struct {
	etcd		*embed.Etcd
	dir		string
	Endpoint	string
}

var etcdContext = EtcdContext{}

func startEtcd() error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var err error
	if etcdContext.dir, err = ioutil.TempDir(os.TempDir(), "service_catalog_integration_test"); err != nil {
		return fmt.Errorf("could not create TempDir: %v", err)
	}
	cfg := embed.NewConfig()
	capnslog.SetGlobalLogLevel(capnslog.WARNING)
	cfg.Dir = etcdContext.dir
	if etcdContext.etcd, err = embed.StartEtcd(cfg); err != nil {
		return fmt.Errorf("Failed starting etcd: %+v", err)
	}
	select {
	case <-etcdContext.etcd.Server.ReadyNotify():
		klog.Info("server is ready!")
	case <-time.After(60 * time.Second):
		etcdContext.etcd.Server.Stop()
		klog.Error("server took too long to start!")
	}
	return nil
}
func stopEtcd() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if etcdContext.etcd == nil {
		return
	}
	etcdContext.etcd.Server.Stop()
	os.RemoveAll(etcdContext.dir)
	select {
	case <-etcdContext.etcd.Server.StopNotify():
		klog.Info("server is stopped!")
	case <-time.After(60 * time.Second):
		klog.Error("server took too long to stop!")
	}
}
func TestMain(m *testing.M) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := startEtcd(); err != nil {
		panic(fmt.Sprintf("Failed to start etcd, %v", err))
	}
	result := m.Run()
	stopEtcd()
	os.Exit(result)
}
