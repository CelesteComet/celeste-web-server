package main

import (
	"net/http"
)

func SayHello() string  {
	return "HELLO"
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public/"))
	index := http.FileServer(http.Dir("client/dist/"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.Handle("/", index)

	server := &http.Server{
		Addr:			"0.0.0.0:8080",
		Handler: 	mux,
	}

	server.ListenAndServe()
}


