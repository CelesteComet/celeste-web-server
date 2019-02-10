package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/matryer/respond.v1"
	"gopkg.in/resty.v1"
	// "fmt"
)

type AuthHandler struct{}

func (h *AuthHandler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging in a user")
		body, err := ioutil.ReadAll(r.Body)

		credentials := make(map[string]interface{})

		if err = json.Unmarshal(body, &credentials); err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			return
		}

		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			SetBody(string(body)).
			Post("http://localhost:1337/login")

		if resp.StatusCode() == 200 {
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
		} else {
			err := []string{}
			json.Unmarshal(resp.Body(), &err)
			respond.With(w, r, resp.StatusCode(), err)
			return
		}
	})
}

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

func (h *AuthHandler) SignUp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("implement Signup")
	})
}
