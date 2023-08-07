package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	host     = "prom-ler-db"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "prom-ler-db"
)

func InitializePostgres() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logrus.Error("Could not connect to postgres", "error", err.Error())
		return nil
	}
	logrus.Info("Successfully connected to Postgres")

	return db
}
