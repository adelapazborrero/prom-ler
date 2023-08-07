package restapi

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/adelapazborrero/prom-ler/internal/app/users"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	port         string
	userResource users.UserHttpResource
	Router       *mux.Router
	server       *http.Server
}

func NewHTTPServer(db *sql.DB) HTTPServer {
	return HTTPServer{
		port:         ":8080",
		userResource: users.NewHTTPResource(db),
	}
}
func (s *HTTPServer) InitializeServer() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Initializing server")
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", s.userResource.FindUsersById).Methods(http.MethodGet)

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	s.Router = r

	s.server = &http.Server{
		Handler: r,
		Addr:    "localhost" + s.port,
		// WriteTimeout: 15 * time.Second,
		// ReadTimeout:  15 * time.Second,
	}

}

func (s *HTTPServer) Serve() {
	logrus.Info("Server starting on port" + s.port)
	log.Fatal(s.server.ListenAndServe())
}
