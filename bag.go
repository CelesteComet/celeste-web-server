package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gopkg.in/matryer/respond.v1"
)

// Bag represents a bag
type Bag struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	ImageURL  string `json:"image_url"  db:"image_url"`
	CreatedBy int    `json:"created_by" db:"created_by"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

// BagHandler handles API
type BagHandler struct {
	DB *sqlx.DB
}

// Index returns all bags
func (h *BagHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bags := []*Bag{}
		err := h.DB.Select(&bags, "SELECT * FROM Bag")
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			return
		}
		respond.With(w, r, http.StatusOK, bags)
	})
}

// Create one bag
func (h *BagHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bag := Bag{}

		err := json.NewDecoder(r.Body).Decode(&bag)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}

		query := `INSERT INTO Bag (name, brand, image_url, created_by) VALUES ($1, $2, $3, $4) RETURNING Bag.id, Bag.created_at`

		rows, err := h.DB.Queryx(query, bag.Name, bag.Brand, bag.ImageURL, 1)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}

		for rows.Next() {
			if err := rows.StructScan(&bag); err != nil {
				respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			}
		}

		respond.With(w, r, http.StatusOK, bag)
	})
}

// Show one bag
func (h *BagHandler) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bag := Bag{}
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["n"])
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			return
		}
		err = h.DB.Get(&bag, "SELECT * FROM Bag WHERE id = $1", i)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			return
		}
		respond.With(w, r, http.StatusOK, bag)
	})
}

// Update a bag
func (h *BagHandler) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bag := Bag{}
		err := json.NewDecoder(r.Body).Decode(&bag)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}
		query := `UPDATE Bag 
SET name = $1, 
brand = $2,
image_url = $3,
created_by = $4
WHERE id = $5`
		rows, err := h.DB.Queryx(query, bag.Name, bag.Brand, bag.ImageURL, 1, 145)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}

		for rows.Next() {
			if err := rows.StructScan(&bag); err != nil {
				respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			}
		}

		respond.With(w, r, http.StatusOK, bag)
	})
}

// Destroy a bag
func (h *BagHandler) Destroy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("implement get all bags")
	})
}

// Helper Methods

// Gets the user ID from the JWT that is passed in the request through context
func getCurrentUserID(r *http.Request) int {
	context := r.Context().Value("ctx")
	if !context.(bool) {
		return 0
	}
	return context.(jwt.MapClaims)["id"].(int)
}
