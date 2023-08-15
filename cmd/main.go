package main

import (
	"net/http"

	"github.com/adelapazborrero/ds"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/adelapazborrero/prom-ler/internal/infra/config"
	"github.com/adelapazborrero/prom-ler/internal/restapi"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func main() {
	//METRICS
	stack := ds.NewStack[string]()
	q := ds.NewQueue[string]()
	stack.Add("hi")
	q.Add("hello")
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
