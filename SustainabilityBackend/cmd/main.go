package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ascribner/sustainabilityapp/env"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database init
	var dsn string = env.EnvConfig.DbUser + ":" + env.EnvConfig.DbPassword + "@tcp(" + env.EnvConfig.DbHost + ":" + env.EnvConfig.DbPort + ")/" + env.EnvConfig.AppDB

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to DB")

	// Router
	mux := http.NewServeMux()

	// Server
	server := NewServer(env.EnvConfig.ServerHost, env.EnvConfig.ServerPort, mux, db)

	log.Print("Running on " + env.EnvConfig.ServerHost + ":" + env.EnvConfig.ServerPort)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
