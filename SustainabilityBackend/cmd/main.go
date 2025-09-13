package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Loads env variables into enviroment
	godotenv.Load()

	// Get env variables
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	appDB := os.Getenv("APP_DB")

	// Database init
	var dsn string = dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + appDB

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Router
	mux := http.NewServeMux()

	server := NewServer(serverHost, serverPort, mux, db)

	server.Run()
}
