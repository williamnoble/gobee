package main

import (
	"database/sql"
	"fmt"
	"gobee/internal/data"
	"gobee/internal/jsonlogger"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"

type application struct {
	config *config
	models data.Models
	logger *jsonlogger.Logger

}

func main() {
	var cfg config
	var db *sql.DB

	app := application{
		config: getConfig(cfg),
		logger: jsonlogger.New(os.Stdout, jsonlogger.LevelInfo),
		models: data.NewModels(db),
	}

	svr := &http.Server{
		Addr: fmt.Sprintf(":d", app.config.port),
		Handler: app.routes(),
		WriteTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
	}

	err := svr.ListenAndServe()
	app.logger.PrintFatal(err, nil)
}