package main

import (
	"flag"
	"os"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
		//maxOpenConnections int
		//maxIdleConnections int
		//maxIdleTime int
	}
}

func getConfig(cfg config) *config {
	flag.IntVar(&cfg.port, "port", 4000, "Default API Server Port" )
	flag.StringVar(&cfg.env, "environment", "development", "Default Environment")
	flag.StringVar(&cfg.db.dsn, "database-dsn", os.Getenv("GOBEE_DB_DSN"), "Database DSN")
	flag.Parse()
	return &cfg
}