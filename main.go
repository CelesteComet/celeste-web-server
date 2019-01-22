package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
  "github.com/CelesteComet/celeste-web-server/app/http"
  "github.com/CelesteComet/celeste-web-server/app/postgres"
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

type CelesteWebServer struct {
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

	server := &CelesteWebServer{
		database: db,
		router:   router,
	}

	// Initialize Routes
	server.routes()
  bagService := postgres.BagService{DB: db}
	bagHTTPService := mhttp.BagHTTPService{BagService: bagService}
	server.router.HandleFunc("/bag/", bagHTTPService.Index()).Methods("GET")

	/*
	bagservice := postgres.bagservice{db: db}
	bags, err := bagService.Bags()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bags)

	bag, err := bagService.Bag(2)
	if err != nil {
		fmt.Println(err)
	}
	*/
  

	loggedRouter := handlers.LoggingHandler(os.Stdout, server.router)
	http.ListenAndServe(":8080", loggedRouter)
}
