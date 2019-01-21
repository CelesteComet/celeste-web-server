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
  "github.com/CelesteComet/celeste-web-server/app/postgres"
  mHttp "github.com/CelesteComet/celeste-web-server/app/http"
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

func main() {
	router := mux.NewRouter()


	// Route Handler for Public Files
	fileHandler := http.FileServer(http.Dir("./public"))
	// Files will be served from /public/*
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileHandler))

	// Router handler for static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/dist")))

  // Serve index page through frontend for all unhandled routes
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/dist/index.html")
	})

	// Connect to database.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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

	fmt.Println(bag)



	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Addr:			"0.0.0.0:8080",
		Handler: 	loggedRouter,
	}

	log.Fatal(server.ListenAndServe())
}


