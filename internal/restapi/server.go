package restapi

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/adelapazborrero/prom-ler/internal/app/users"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	port         string
	userResource users.UserHttpResource
	Router       *mux.Router
	server       *http.Server
	registry     *prometheus.Registry
}

func NewHTTPServer(db *sql.DB, prometheusRegistry *prometheus.Registry) HTTPServer {
	return HTTPServer{
		port:         ":8080",
		registry:     prometheusRegistry,
		userResource: users.NewHTTPResource(db),
	}
}
func (s *HTTPServer) InitializeServer() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Initializing server")
	r := mux.NewRouter()

	r.Handle("/metrics", promhttp.HandlerFor(s.registry, promhttp.HandlerOpts{Registry: s.registry})).Methods(http.MethodGet)

	r.HandleFunc("/users/{id}", s.userResource.FindUsersById).Methods(http.MethodGet)

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)

	httpServer := &http.Server{
		Addr:         s.port,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s.Router = r
	s.server = httpServer

}

func (s *HTTPServer) Serve() {
	logrus.Info("Server starting on port" + s.port)

	err := s.server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
