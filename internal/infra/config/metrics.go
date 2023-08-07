package config

import (
	"github.com/adelapazborrero/prom-ler/internal/infra/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type MetricRegistry struct {
	Registry *prometheus.Registry
	Metrics  *metrics.Metrics
}

func InitializePrometheus() MetricRegistry {
	reg := prometheus.NewRegistry()

	m := &metrics.Metrics{
		CpuTemp: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Current temperature of the CPU.",
		}),
		HdFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "hd_errors_total",
				Help: "Number of hard-disk errors.",
			},
			[]string{"device"},
		),
		Objectman: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "hello_world",
				Help: "The thing that we like",
			},
			[]string{"MyVal"},
		),
	}

	reg.MustRegister(m.CpuTemp)
	reg.MustRegister(m.HdFailures)
	reg.MustRegister(m.Objectman)

	return MetricRegistry{
		Registry: reg,
		Metrics:  m,
	}
}
