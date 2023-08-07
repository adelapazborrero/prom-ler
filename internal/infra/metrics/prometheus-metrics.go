package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	CpuTemp    prometheus.Gauge
	HdFailures *prometheus.CounterVec
	Objectman  *prometheus.GaugeVec
}
