package server

import (
	"github.com/kubernetes-incubator/service-catalog/cmd/controller-manager/app"
	"github.com/kubernetes-incubator/service-catalog/cmd/controller-manager/app/options"
	"github.com/kubernetes-incubator/service-catalog/pkg/hyperkube"
)

func NewControllerManager() *hyperkube.Server {
	_logClusterCodePath()
	defer _logClusterCodePath()
	s := options.NewControllerManagerServer()
	hks := hyperkube.Server{PrimaryName: "controller-manager", AlternativeName: "service-catalog-controller-manager", SimpleUsage: "controller-manager", Long: `The service-catalog controller manager is a daemon that embeds the core control loops shipped with the service catalog.`, Run: func(_ *hyperkube.Server, args []string, stopCh <-chan struct{}) error {
		return app.Run(s)
	}, RespectsStopCh: false}
	s.AddFlags(hks.Flags())
	return &hks
}
