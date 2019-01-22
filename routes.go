package main

import (
  "log"
	"net/http"
)

// Route Handler for Public Files
var serverFileHandler = http.FileServer(http.Dir("./public"))

// Router Handler for Client Static Files
var clientFileHandler = http.FileServer(http.Dir("./client/dist"))

// Index Handler for React Application
var indexHandler = func(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "./client/dist/index.html")
}

func (s *CelesteWebServer) routes() {
  log.Println("Creating server routes")
	s.router.HandleFunc("/", indexHandler)
	//s.router.HandleFunc("/api/bag", handleBag)
  s.router.PathPrefix("/public/").Handler(http.StripPrefix("/public/",serverFileHandler))
	s.router.PathPrefix("/").Handler(clientFileHandler)
}
