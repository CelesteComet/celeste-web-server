package main

import (
	"net/http"
	"fmt"
	"log"
  "database/sql"
  _ "github.com/lib/pq"
	"os"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// Declare the database
var (
	host = "celestecomet.c7bjz8zer8ha.us-east-1.rds.amazonaws.com"
	port = 5432
	user = os.Getenv("AWS_DB_USERNAME")
	password = os.Getenv("AWS_DB_PASSWORD")
	dbname = "CelesteComet"
)

var (
  connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func SayHello() string  {
	return "HELLO"
}

type CelesteWebServer struct {
  database *sql.DB
	router *mux.Router
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

	server := &CelesteWebServer{
	  database: db,
		router: router,
	}


  // Initialize Routes
	server.routes()


	bagService := postgres.BagService{DB: db}
	bags, err := bagService.Bags()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bags)

	bag, err := bagService.Bag(2)
	if err != nil {
	  fmt.Println(err)
	}

  loggedRouter := handlers.LoggingHandler(os.Stdout, server.router)
  http.ListenAndServe(":8080", loggedRouter) 
}
