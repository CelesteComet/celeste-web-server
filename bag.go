package main

import (
  "log"
  "gopkg.in/matryer/respond.v1"
  "net/http"
  "github.com/jmoiron/sqlx"  
  "github.com/gorilla/mux"
  "strconv"
  "encoding/json"
  "github.com/dgrijalva/jwt-go"
)

type Bag struct {
  Id         int `json:"id"`
  Name       string `json:"name"`
  Brand      string `json:"brand"`
  Image_url  string `json:"image_url"`
  Created_by int `json:"created_by"`
  Created_at string `json:"created_at"`  
}

type BagHandler struct {
  DB *sqlx.DB
}

func (h *BagHandler) Index() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Context().Value("ctx").(jwt.MapClaims)["id"])
    bags := []*Bag{}
    err := h.DB.Select(&bags, "SELECT * FROM Bag")
    if err != nil {
      respond.With(w, r, http.StatusInternalServerError, []string{err.Error()}) 
      return; 
    }
    respond.With(w, r, http.StatusOK, bags)
  })
}

func (h *BagHandler) Create() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    bag := Bag{}

    err := json.NewDecoder(r.Body).Decode(&bag); if err != nil {
      respond.With(w, r, http.StatusInternalServerError, []string{err.Error()}) 
    }

    query := `INSERT INTO Bag (name, brand, image_url, created_by) VALUES ($1, $2, $3, $4) RETURNING Bag.id, Bag.created_at`

    rows, err := h.DB.Queryx(query, bag.Name, bag.Brand, bag.Image_url, 1)
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

func (h *BagHandler) Show() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    bag := Bag{}
    vars := mux.Vars(r)
    i, err := strconv.Atoi(vars["n"]) 
    if err != nil {
      respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})  
      return;
    }
    err = h.DB.Get(&bag, "SELECT * FROM Bag WHERE id = $1", i)
    if err != nil {
      respond.With(w, r, http.StatusInternalServerError, []string{err.Error()}) 
      return; 
    }
    respond.With(w, r, http.StatusOK, bag)    
  })
}

func (h *BagHandler) Update() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    bag := Bag{}
    err := json.NewDecoder(r.Body).Decode(&bag); if err != nil {
      respond.With(w, r, http.StatusInternalServerError, []string{err.Error()}) 
    }    
    query := `UPDATE Bag 
SET name = $1, 
brand = $2,
image_url = $3,
created_by = $4
WHERE id = $5`
    rows, err := h.DB.Queryx(query, bag.Name, bag.Brand, bag.Image_url, 1, 145)
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

func (h *BagHandler) Destroy() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println("implement get all bags")
  })
}

