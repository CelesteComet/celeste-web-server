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

type Bag struct {
	Id int
	Name string
	Brand string
	Image_url string
}

type Bags []Bag


var (
  connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

func SayHello() string  {
	return "HELLO"
}

func main() {
	router := mux.NewRouter()

	// Route Handler for Public Files
	files := http.FileServer(http.Dir("./public"))
	// Files will be served from /public/*
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", files))

	// Router handler for static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./client/dist")))


  // Serve index page through frontend for all unhandled routes
  router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/dist/index.html")
	})

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from bag")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	bags := Bags{}
	for rows.Next() {
		bag := Bag{}
		if err := rows.Scan(&bag.Id, &bag.Name, &bag.Brand, &bag.Image_url); err != nil {
			log.Fatal(err)
		} 
		bags = append(bags, bag)
	}

	//bagsJson, err := json.Marshal(bags)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Fprintf(os.Stdout, "%s", bagsJson)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	server := &http.Server{
		Addr:			"0.0.0.0:8080",
		Handler: 	loggedRouter,
	}

	log.Fatal(server.ListenAndServe())
}


