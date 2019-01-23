package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

// route handler for public files
var serverFileHandler = http.FileServer(http.Dir("./public"))

// Router Handler for Client Static Files
var clientFileHandler = http.FileServer(http.Dir("./client/dist"))

// Index Handler for React Application
var indexHandler = func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/dist/index.html")
}

type Routes struct {
	Router *mux.Router
}

func (r *Routes) Init() {
	r.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", serverFileHandler))
}

func (r *Routes) Listen() {
	http.ListenAndServe(":8080", r.Router)
}
