package routes

import (
	//"database/sql"
	"github.com/CelesteComet/celeste-auth-server/pkg/auth"
	"github.com/CelesteComet/celeste-web-server/app"
	mhttp "github.com/CelesteComet/celeste-web-server/app/http"
	"github.com/CelesteComet/celeste-web-server/app/postgres"
	_ "github.com/lib/pq"
	"net/http"
  "github.com/markbates/goth"
  "github.com/markbates/goth/gothic"
  "github.com/markbates/goth/providers/google"    	
  "os"
  "log"
  "encoding/json"
	"io/ioutil"
  "bytes"
)

type indexHandler struct {}

func (h indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./client/dist/index.html")
}



func InitRoutes(s *app.Server) {

  goth.UseProviders(
    google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),  
  )  	

	// Public files that are stored on server with static files for React client
	serverFilesHandler := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	staticFilesHandler := http.FileServer(http.Dir("./client/dist"))

	// Create Services
	bagService := postgres.BagService{DB: s.Database}

	// Create Handlers
	bagHandler := mhttp.BagHandler{BagService: bagService}

	// Attach Handlers to Routes
	s.Router.PathPrefix("/public/").Handler(auth.MustAuth(serverFilesHandler))

	// Authentication Routes
	s.Router.Handle("/auth", &auth.CheckLoggedInHandler{})
	s.Router.Handle("/auth/logout", &auth.LogOutHandler{})
	s.Router.Handle("/auth/{provider}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if _ , err := gothic.CompleteUserAuth(w, r); err == nil {
      // t, _ := template.New("foo").Parse(userTemplate)
      // t.Execute(res, gothUser)
    } else {
    	log.Println("DOING WHAT NEEDS TO BE DONE")
      gothic.BeginAuthHandler(w, r)
    }     
  }))

	s.Router.Handle("/auth/{provider}/callback", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if gothUser , err := gothic.CompleteUserAuth(w, r); err == nil {
			url := "http://localhost:1337/oauth"
			jsonStr, err := json.Marshal(gothUser)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	    client := &http.Client{}
	    resp, err := client.Do(req)
	    if err != nil {
	        panic(err)
	    }
	    bodyBytes, _ := ioutil.ReadAll(resp.Body)
	    _ = string(bodyBytes)	 

  // Set HttpOnly To Prevent Future Tampering
	  http.SetCookie(w, &http.Cookie{
	    Name:   "jwt",
	    Value: resp.Header.Get("JWT"),
	    HttpOnly: true,
	    Path: "/",
	  })	    
	   	http.Redirect(w, r, "http://localhost:8080", http.StatusMovedPermanently)
	    defer resp.Body.Close()			
    } 
  }))  

	// API Routes
	s.Router.Handle("/api/bags", bagHandler.GetBags())
	s.Router.Handle("/api/bags/{n}", bagHandler.GetBag())
	s.Router.Handle("/api/bagtags", bagHandler.GetBagWithTag())
	s.Router.Handle("/api/users/{userID}/bags", auth.MustAuth(bagHandler.GetUserBags()))

	// React Application
	s.Router.PathPrefix("/").Handler(staticFilesHandler)
	s.Router.PathPrefix("/").Handler(indexHandler{})
}
