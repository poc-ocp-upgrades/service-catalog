package framework

import (
	"net/http"
	"sync"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/klog"
)

var registerMetrics sync.Once

const (
	promNamespace = "servicecatalog_health"
)

var (
	ExecutionCount				= prometheus.NewCounter(prometheus.CounterOpts{Namespace: promNamespace, Name: "execution_count", Help: "Number of times the health check has run."})
	ErrorCount					= prometheus.NewCounterVec(prometheus.CounterOpts{Namespace: promNamespace, Name: "error_count", Help: "Number of times the health check ended in error, by error."}, []string{"error"})
	eventHandlingTimeSummary	= prometheus.NewSummaryVec(prometheus.SummaryOpts{Namespace: promNamespace, Name: "successful_duration_seconds", Help: "processing time (s) of successfully executed operation, by operation.", Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}}, []string{"operation"})
)

func ReportOperationCompleted(operation string, startTime time.Time) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	eventHandlingTimeSummary.WithLabelValues(operation).Observe(time.Since(startTime).Seconds())
}
func register(registry *prometheus.Registry) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	registerMetrics.Do(func() {
		registry.MustRegister(ExecutionCount)
		registry.MustRegister(ErrorCount)
		registry.MustRegister(eventHandlingTimeSummary)
	})
}
func RegisterMetricsAndInstallHandler(m *http.ServeMux) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	registry := prometheus.NewRegistry()
	register(registry)
	m.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError}))
	klog.V(3).Info("Registered /metrics with prometheus")
}
