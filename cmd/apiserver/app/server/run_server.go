package server

import (
	"fmt"
	"net/http"
	"github.com/kubernetes-incubator/service-catalog/pkg/api"
	"k8s.io/apiserver/pkg/server/healthz"
	genericapiserverstorage "k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storage/etcd3/preflight"
	"github.com/kubernetes-incubator/service-catalog/pkg/apiserver"
	"github.com/kubernetes-incubator/service-catalog/pkg/apiserver/options"
	"k8s.io/klog"
)

func RunServer(opts *ServiceCatalogServerOptions, stopCh <-chan struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if stopCh == nil {
		stopCh = make(chan struct{})
	}
	err := opts.Validate()
	if nil != err {
		return err
	}
	return runEtcdServer(opts, stopCh)
}
func runEtcdServer(opts *ServiceCatalogServerOptions, stopCh <-chan struct{}) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	etcdOpts := opts.EtcdOptions
	klog.V(4).Infoln("Preparing to run API server")
	genericConfig, scConfig, err := buildGenericConfig(opts)
	if err != nil {
		return err
	}
	klog.V(4).Infoln("Creating storage factory")
	storageGroupsToEncodingVersion, err := options.NewStorageSerializationOptions().StorageGroupsToEncodingVersion()
	if err != nil {
		return fmt.Errorf("error generating storage version map: %s", err)
	}
	storageFactory, err := apiserver.NewStorageFactory(etcdOpts.StorageConfig, etcdOpts.DefaultStorageMediaType, api.Codecs, genericapiserverstorage.NewDefaultResourceEncodingConfig(api.Scheme), storageGroupsToEncodingVersion, nil, apiserver.DefaultAPIResourceConfigSource(), nil)
	if err != nil {
		klog.Errorf("error creating storage factory: %v", err)
		return err
	}
	config := apiserver.NewEtcdConfig(genericConfig, 0, storageFactory)
	completed := config.Complete()
	klog.V(4).Infoln("Completing API server configuration")
	server, err := completed.NewServer(stopCh)
	if err != nil {
		return fmt.Errorf("error completing API server configuration: %v", err)
	}
	addPostStartHooks(server.GenericAPIServer, scConfig, stopCh)
	etcdChecker := checkEtcdConnectable{ServerList: etcdOpts.StorageConfig.ServerList}
	healthz.InstallPathHandler(server.GenericAPIServer.Handler.NonGoRestfulMux, "/healthz/ready", etcdChecker)
	klog.Infoln("Running the API server")
	server.PrepareRun().Run(stopCh)
	return nil
}

type checkEtcdConnectable struct{ ServerList []string }

func (c checkEtcdConnectable) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return "etcd"
}
func (c checkEtcdConnectable) Check(_ *http.Request) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	klog.Info("etcd checker called")
	serverReachable, err := preflight.EtcdConnection{ServerList: c.ServerList}.CheckEtcdServers()
	if err != nil {
		klog.Errorf("etcd checker failed with err: %v", err)
		return err
	}
	if !serverReachable {
		msg := "etcd failed to reach any server"
		klog.Error(msg)
		return fmt.Errorf(msg)
	}
	return nil
}
