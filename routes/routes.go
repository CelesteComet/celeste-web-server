package routes

import (
	//"database/sql"
	"github.com/CelesteComet/celeste-auth-server/pkg/auth"
	"github.com/CelesteComet/celeste-web-server/app"
	mhttp "github.com/CelesteComet/celeste-web-server/app/http"
	"github.com/CelesteComet/celeste-web-server/app/postgres"
	_ "github.com/lib/pq"
	"net/http"
)

func InitRoutes(s *app.Server) {

	// Public files that are stored on server with static files for React client
	serverFilesHandler := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	staticFilesHandler := http.FileServer(http.Dir("./client/dist"))

	// Create Services
	bagService := postgres.BagService{DB: s.Database}

	// Create Handlers
	bagHandler := mhttp.BagHandler{BagService: bagService}

	// Attach Handlers to Routes
	s.Router.PathPrefix("/public/").Handler(auth.MustAuth(serverFilesHandler))

	s.Router.Handle("/bags", bagHandler.GetBags())
	s.Router.Handle("/bags/{n}", bagHandler.GetBag())
	s.Router.Handle("/users/{userID}/bags", auth.MustAuth(bagHandler.GetUserBags()))
	s.Router.Handle("/logout", &auth.LogOutHandler{})
	s.Router.PathPrefix("/").Handler(staticFilesHandler)
}
