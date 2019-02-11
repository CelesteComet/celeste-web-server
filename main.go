package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "encoding/json"
)

// Declare the database
var (
	host     = "celestecomet.c7bjz8zer8ha.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = os.Getenv("AWS_DB_USERNAME")
	password = os.Getenv("AWS_DB_PASSWORD")
	dbname   = "CelesteComet"
)

var (
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func main() {
	log.Println("Connecting to AWS RDS Postgresql server")
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server connection successful")
	defer db.Close()

	// Create a Router
	router := mux.NewRouter()

	// Create a Server
	server := Server{
		Router: router,
		Port:   ":8080",
		DB:     db,
	}

	// Initiate Routes
	server.Routes()

	http.ListenAndServe(server.Port, server.Router)
}
