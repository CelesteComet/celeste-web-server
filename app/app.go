package app

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Port     string
	Database *sqlx.DB
	Router   *mux.Router
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(email string, password string) (*User, error)
	FindByCredentials(email string, password string) (*User, error)
}

type UserHandler interface {
	CreateUser(email string, password string) (*User, error)
	FindByCredentials(email string, password string) (*User, error)
}

type Bag struct {
	Id         int `json:"id"`
	Name       string `json:"name"`
	Brand      string `json:"brand"`
	Image_url  string `json:"image_url"`
	Created_by int `json:"created_by"`
}

type BagService interface {
	Bag(id int) (*Bag, error)
	Bags() ([]*Bag, error)
}

type BagHandler interface {
	GetBags() http.Handler
	GetBag() http.Handler
	GetUserBags() http.Handler
}
