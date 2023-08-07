package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UserHttpResource struct {
	db *sql.DB
}

func NewHTTPResource(db *sql.DB) UserHttpResource {
	return UserHttpResource{
		db: db,
	}
}

func (u UserHttpResource) FindUsersById(w http.ResponseWriter, r *http.Request) {
	logrus.Info("FindUsersById")

	vars := mux.Vars(r)
	userId := vars["id"]

	service := NewService(u.db)

	user := service.GetById(userId)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
