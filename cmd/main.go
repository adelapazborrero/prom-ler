package main

import (
	"net/http"

	"github.com/adelapazborrero/prom-ler/internal/app/users"
	"github.com/adelapazborrero/prom-ler/internal/infra/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func main() {
	prom := config.InitializePrometheus()
	postgres := config.InitializePostgres()

	userService := users.NewService(postgres)
	userService.Insert(&users.User{Username: "Roberto", Email: "rob@bol.com"})
	user := userService.GetById("1")
	log.Info("Got user", "username", user.Username, "email", user.Email)

	prom.Metrics.CpuTemp.Set(65.3)
	prom.Metrics.HdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	prom.Metrics.Objectman.With(prometheus.Labels{"MyVal": "Myval"}).Set(1)

	http.Handle("/metrics", promhttp.HandlerFor(prom.Registry, promhttp.HandlerOpts{Registry: prom.Registry}))

	log.Info("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
