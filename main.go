package main

import (
	"database/sql"
	"fmt"
	"github.com/CelesteComet/celeste-web-server/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"os"
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

type Server struct {
	database *sql.DB
	router   *mux.Router
}

func main() {
	router := mux.NewRouter()

	log.Println("Connecting to AWS RDS Postgresql server")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server connection successful")
	defer db.Close()

	//bagService := &postgres.BagService{DB: db}

	// Initialize Routes
	routes := http.Routes{Router: router}
	routes.Init()
	//loggedRouter := handlers.LoggingHandler(os.Stdout, routes.Tier)
	routes.Listen()
}
