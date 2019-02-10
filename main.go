package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// "encoding/json"
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

func authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// If Cookie exists, check the JWT
		tokenString := cookie.Value

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			hmacSampleSecret := []byte("secret")
			return hmacSampleSecret, nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var key interface{} = "ctx"
		ctx := context.WithValue(r.Context(), key, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
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

	// Connect Handlers With Database
	BagHandler := BagHandler{DB: db}
	AuthHandler := AuthHandler{}

	// Create a Router
	router := mux.NewRouter()

	// Public files that are stored on server with static files for React client
	serverFilesHandler := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	// staticFilesHandler := http.FileServer(http.Dir("./client/dist"))

	// Top Namespace API Routes
	apiRoutes := router.PathPrefix("/api").Subrouter()
	authRoutes := router.PathPrefix("/auth").Subrouter()

	// Auth Routes
	authRoutes.Handle("", AuthHandler.Login()).Methods("Post")
	authRoutes.Handle("", AuthHandler.Logout()).Methods("Delete")
	// authRoutes.Handle("", AuthHandler.SignUp()).Methods("Post")

	// Sub API routes
	bagRoutes := apiRoutes.PathPrefix("/bags").Subrouter()

	// API routes
	bagRoutes.Handle("", BagHandler.Index()).Methods("Get")
	bagRoutes.Handle("", BagHandler.Create()).Methods("Post")
	bagRoutes.Handle("/{n}", BagHandler.Show()).Methods("Get")
	bagRoutes.Handle("/{n}", BagHandler.Update()).Methods("Put")
	bagRoutes.Handle("/{n}", BagHandler.Destroy()).Methods("Delete")

	// Server Routes
	router.PathPrefix("/public/").Handler(serverFilesHandler)
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relativeFilePath := "./client/dist" + r.RequestURI
		_, err := os.Stat(relativeFilePath)
		if err != nil {
			http.ServeFile(w, r, "./client/dist/index.html")
		}
		http.ServeFile(w, r, relativeFilePath)
	})

	// Create a Server
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	server.ListenAndServe()
}
