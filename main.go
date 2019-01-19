package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"database/sql"
  _ "github.com/lib/pq"
	"os"
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
  fmt.Println(os.Environ())
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("publc/"))
	index := http.FileServer(http.Dir("client/dist/"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/", index)

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
		fmt.Println("DOING")
	}

	bagsJson, err := json.Marshal(bags)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(os.Stdout, "%s", bagsJson)

	server := &http.Server{
		Addr:			"0.0.0.0:8080",
		Handler: 	mux,
	}

	server.ListenAndServe()
}


