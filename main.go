package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/time/rate"

	"github.com/CelesteComet/celeste-web-server/config"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "encoding/json"
)

// Declare the database
var (
	host     = "raja.db.elephantsql.com"
	port     = 5432
	user     = os.Getenv("AWS_DB_USERNAME")
	password = os.Getenv("AWS_DB_PASSWORD")
	dbname   = "umtggmht"
)

var (
	connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
)

var limiter = rate.NewLimiter(2, 5)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	log.Println("Connecting to AWS RDS Postgresql server")
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server connection successful")
	defer db.Close()

	// Initialize Configuration
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// Create a Router
	router := mux.NewRouter()
	router.Use(limit)

	// Create a Server
	server := Server{
		Router: router,
		Port:   ":8080",
		DB:     db,
	}

	// Initiate Routes
	server.Routes(config)

	http.ListenAndServe(server.Port, server.Router)
}
