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
	Id         int
	Name       string
	Brand      string
	Image_url  string
	Created_by int
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
