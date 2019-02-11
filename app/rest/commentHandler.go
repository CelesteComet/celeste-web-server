package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CelesteComet/celeste-web-server/app"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gopkg.in/matryer/respond.v1"
)

// CommentHandler handles http requests to comments
type CommentHandler struct {
	DB *sqlx.DB
}

// ID              int    `json:"id"`
// ItemID          int    `json:"item_id" db:"item_id"`
// Content         string `json:"name"`
// CreatedBy       int    `json:"created_by" db:"created_by"`
// CreatedByMember string `json:"created_by_member" db:"created_by_member"`
// CreatedAt       string `json:"creat

// Comment represents a comment
type Comment app.Comment

// Index shows all comments of an item
func (h *CommentHandler) Index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		comments := []*Comment{}
		itemID := mux.Vars(r)["itemID"]

		query := `select 
comments.id, comments.item_id, comments.created_by, display_name as created_by_member, comments.created_at 
from comments 
join member on member.id = comments.created_by
where comments.item_id = $1;`

		err := h.DB.Select(&comments, query, itemID)
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, []string{err.Error()})
		}
		respond.With(w, r, http.StatusOK, comments)
	})
}

// Create a new comment
func (h *CommentHandler) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value("ctx").(jwt.MapClaims)["id"]
		itemID := mux.Vars(r)["itemID"]
		comment := Comment{}
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, []string{err.Error()})
		}
		json.Unmarshal(bytes, &comment)

		query := `INSERT INTO comments (item_id, content, created_by) VALUES ($1, $2, $3) returning *`
		rows, err := h.DB.Queryx(query, itemID, comment.Content, userID)
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, []string{err.Error()})
		}

		for rows.Next() {
			if err := rows.StructScan(&comment); err != nil {
				respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			}
		}

		respond.With(w, r, http.StatusOK, comment)
	})
}

func (h *CommentHandler) Show() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (h *CommentHandler) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (h *CommentHandler) Destroy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("DELETING A COMMENT")
		comment := Comment{}
		userID := r.Context().Value("ctx").(jwt.MapClaims)["id"].(float64)
		commentID := mux.Vars(r)["id"]

		query := `select * from comments where id = $1`
		err := h.DB.Get(&comment, query, commentID)
		if err != nil {
			respond.With(w, r, http.StatusBadRequest, []string{err.Error()})
			return
		}

		if float64(comment.CreatedBy) != userID {
			respond.With(w, r, http.StatusUnauthorized, []string{"Not Authorized"})
			return
		}

		_, err = h.DB.Queryx(`delete from comments where id = $1`, commentID)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			return
		}

		respond.With(w, r, http.StatusOK, comment)
	})
}
