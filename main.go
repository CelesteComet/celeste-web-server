package main

import (
	"fmt"
	"github.com/CelesteComet/celeste-web-server/app"
	"github.com/CelesteComet/celeste-web-server/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"github.com/jmoiron/sqlx"
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

func SayHello() string {
	return "HELLO"
}

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
	s := &app.Server{Database: db, Router: router, Port: ":8080"}

	// Initialize Server Routes
	routes.InitRoutes(s)

	// Start the Server
	http.ListenAndServe(s.Port, s.Router)
}
