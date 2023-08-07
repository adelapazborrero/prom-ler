package main

import (
	"net/http"

	"github.com/adelapazborrero/prom-ler/internal/infra/config"
	"github.com/adelapazborrero/prom-ler/internal/restapi"
	"github.com/prometheus/client_golang/prometheus"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	//METRICS
	prom := config.InitializePrometheus()
	prom.Metrics.CpuTemp.Set(65.3)
	prom.Metrics.HdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	prom.Metrics.Objectman.With(prometheus.Labels{"MyVal": "Myval"}).Set(1)

	//POSTGRES
	postgres := config.InitializePostgres()
	defer postgres.Close()

	//HTTP SERVER
	httpserver := restapi.NewHTTPServer(postgres, prom.Registry)
	httpserver.InitializeServer()

	httpserver.Serve()
}
