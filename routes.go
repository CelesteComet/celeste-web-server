package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CelesteComet/celeste-web-server/app/rest"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Server is a server with DB, Router, Port and maybe email...
type Server struct {
	Router *mux.Router
	DB     *sqlx.DB
	Port   string
}

// Routes method initailizes routes on the router
func (s *Server) Routes() {

	// Connect Handlers With Database
	BagHandler := rest.BagHandler{DB: s.DB}
	AuthHandler := rest.AuthHandler{}
	CommentHandler := rest.CommentHandler{DB: s.DB}

	// Public files that are stored on server with static files for React client
	serverFilesHandler := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	// staticFilesHandler := http.FileServer(http.Dir("./client/dist"))

	apiRoutes := s.Router.PathPrefix("/api").Subrouter()
	authRoutes := s.Router.PathPrefix("/auth").Subrouter()

	// Auth Routes
	authRoutes.Handle("", AuthHandler.Authenticate()).Methods("Get")
	authRoutes.Handle("", AuthHandler.Login()).Methods("Post")
	authRoutes.Handle("", AuthHandler.Logout()).Methods("Delete")
	// authRoutes.Handle("", AuthHandler.SignUp()).Methods("Post")

	// Sub API routes
	bagRoutes := apiRoutes.PathPrefix("/bags").Subrouter()
	commentRoutes := apiRoutes.PathPrefix("/{itemID}/comments").Subrouter()

	// API routes
	bagRoutes.Handle("", BagHandler.Index()).Methods("Get")
	bagRoutes.Handle("", BagHandler.Create()).Methods("Post")
	bagRoutes.Handle("/{n}", BagHandler.ShowBagDetailPage()).Methods("Get")
	bagRoutes.Handle("/{n}", BagHandler.Update()).Methods("Put")
	bagRoutes.Handle("/{n}", BagHandler.Destroy()).Methods("Delete")

	commentRoutes.Handle("", authenticateUser(CommentHandler.Index())).Methods("Get")
	commentRoutes.Handle("/{id}", authenticateUser(CommentHandler.Create())).Methods("Post")

	// Server Routes
	s.Router.PathPrefix("/public/").Handler(serverFilesHandler)
	s.Router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		relativeFilePath := "./client/dist" + r.RequestURI
		log.Println(relativeFilePath)
		_, err := os.Stat(relativeFilePath)
		if err != nil {
			http.ServeFile(w, r, "./client/dist/index.html")
		}
		http.ServeFile(w, r, relativeFilePath)
	})
}

// MIDDLEWARE

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
