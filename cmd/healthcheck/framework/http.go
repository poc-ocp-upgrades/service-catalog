package framework

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/klog"
)

func ServeHTTP(healthcheckOptions *HealthCheckServer) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if err := healthcheckOptions.SecureServingOptions.MaybeDefaultWithSelfSignedCerts("", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return fmt.Errorf("failed to establish SecureServingOptions %v", err)
	}
	klog.V(3).Infof("Starting http server and mux on port %v", healthcheckOptions.SecureServingOptions.BindPort)
	go func() {
		mux := http.NewServeMux()
		RegisterMetricsAndInstallHandler(mux)
		healthz.InstallHandler(mux, healthz.PingHealthz)
		server := &http.Server{Addr: net.JoinHostPort(healthcheckOptions.SecureServingOptions.BindAddress.String(), strconv.Itoa(healthcheckOptions.SecureServingOptions.BindPort)), Handler: mux}
		klog.Fatal(server.ListenAndServeTLS(healthcheckOptions.SecureServingOptions.ServerCert.CertKey.CertFile, healthcheckOptions.SecureServingOptions.ServerCert.CertKey.KeyFile))
	}()
	return nil
}
