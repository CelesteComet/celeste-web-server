package main

import (
	"net/http"
	"fmt"
	"log"
	"database/sql"
  _ "github.com/lib/pq"
)

// Declare the database
const (
	connStr = "postgres://cwyfcaka:SBkAFLqMd4G27FYAFhoPy9yYfdfCfWMK@elmer.db.elephantsql.com:5432/cwyfcaka"
)

func SayHello() string  {
	return "HELLO"
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("publc/"))
	index := http.FileServer(http.Dir("client/dist/"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/", index)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from person")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		} 
		names = append(names, name)
	}

	fmt.Println(names)


	server := &http.Server{
		Addr:			"0.0.0.0:8080",
		Handler: 	mux,
	}

	server.ListenAndServe()
}


