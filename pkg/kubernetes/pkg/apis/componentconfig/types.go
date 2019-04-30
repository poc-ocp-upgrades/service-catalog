package componentconfig

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LeaderElectionConfiguration struct {
	LeaderElect	bool
	LeaseDuration	metav1.Duration
	RenewDeadline	metav1.Duration
	RetryPeriod	metav1.Duration
	ResourceLock	string
}
