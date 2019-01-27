package routes

import (
	//"database/sql"
	"github.com/CelesteComet/celeste-auth-server/pkg/auth"
	"github.com/CelesteComet/celeste-web-server/app"
	"github.com/CelesteComet/celeste-web-server/app/http"
	"github.com/CelesteComet/celeste-web-server/app/postgres"
	_ "github.com/lib/pq"
)

func InitRoutes(s *app.Server) {

	// Create Services
	bagService := postgres.BagService{DB: s.Database}

	// Create Handlers
	bagHandler := http.BagHandler{BagService: bagService}

	// Attach Handlers to Routes
	s.Router.Handle("/bags", bagHandler.GetBags())
	s.Router.Handle("/bags/{n}", bagHandler.GetBag())
	s.Router.Handle("/users/{userID}/bags", auth.MustAuth(bagHandler.GetUserBags()))
	s.Router.Handle("/logout", &auth.LogOutHandler{})

}
