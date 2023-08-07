package main

import (
	"github.com/adelapazborrero/prom-ler/internal/infra/config"
	"github.com/adelapazborrero/prom-ler/internal/restapi"
)

func main() {
	//METRICS
	// prom := config.InitializePrometheus()
	// prom.Metrics.CpuTemp.Set(65.3)
	// prom.Metrics.HdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	// prom.Metrics.Objectman.With(prometheus.Labels{"MyVal": "Myval"}).Set(1)

	//POSTGRES
	postgres := config.InitializePostgres()
	defer postgres.Close()

	//HTTP SERVER
	httpserver := restapi.NewHTTPServer(postgres)
	httpserver.InitializeServer()

	// http.Handle("/metrics", promhttp.HandlerFor(prom.Registry, promhttp.HandlerOpts{Registry: prom.Registry}))

	httpserver.Serve()
}
