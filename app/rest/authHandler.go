package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CelesteComet/celeste-web-server/app"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/matryer/respond.v1"
	"gopkg.in/resty.v1"
)

// AuthHandler calls Auth microservice API to do authentication things
type AuthHandler struct{}

var _ app.AuthHandler = &AuthHandler{}

// Login logs in a user
func (h *AuthHandler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging in a user")
		body, err := ioutil.ReadAll(r.Body)

		credentials := make(map[string]interface{})

		if err = json.Unmarshal(body, &credentials); err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			return
		}

		resty.SetDebug(true)
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(string(body)).
			Post("http://ec2-54-85-14-41.compute-1.amazonaws.com/login")

		if resp.StatusCode() != 200 {
			myErrors := []string{}
			json.Unmarshal(resp.Body() , &myErrors)
			respond.With(w, r, 500, myErrors)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    resp.Header().Get("jwt"),
			HttpOnly: true,
			Path:     "/",
		})

		user := make(map[string]interface{})
		json.Unmarshal(resp.Body(), &user)
		respond.With(w, r, resp.StatusCode(), user)
		return

	})
}

// Logout logs out a user by killing the cookie
func (h *AuthHandler) Logout() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "jwt",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
	})
}

// Authenticate checks JWT and validates it
func (h *AuthHandler) Authenticate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			respond.With(w, r, http.StatusUnauthorized, []string{err.Error()})
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
			respond.With(w, r, http.StatusUnauthorized, []string{err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respond.With(w, r, http.StatusUnauthorized, []string{err.Error()})
			return
		}

		// var key interface{} = "ctx"
		// ctx := context.WithValue(r.Context(), key, claims)

		respond.With(w, r, http.StatusOK, claims)
	})
}

// Sign up calls the Auth microservice API to create a new user
func (h *AuthHandler) SignUp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("implement Signup")
	})
}
