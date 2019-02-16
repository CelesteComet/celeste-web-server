package app

import (
	"net/http"
)

// Server has a method Routes which initializes all routes
type Server interface {
	Routes()
}

// AuthHandler interface describes authentication actions
type AuthHandler interface {
	Login() http.Handler
	Logout() http.Handler
	Authenticate() http.Handler
	SignUp() http.Handler
}

// Bag represents a bag
type Bag struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	ImageURL  string `json:"image_url"  db:"image_url"`
	CreatedBy int    `json:"created_by" db:"created_by"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

// BagPage represents a single view of a bag
type BagPage struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Brand           string `json:"brand"`
	ImageURL        string `json:"image_url"  db:"image_url"`
	CreatedBy       int    `json:"created_by" db:"created_by"`
	CreatedByMember string `json:"created_by_member" db:"created_by_member"`
	CreatedAt       string `json:"created_at" db:"created_at"`
}

// BagHandler interface makes HTTP requests for bags
type BagHandler interface {
	Index() http.Handler
	Create() http.Handler
	Show() http.Handler
	ShowBagDetailPage() http.Handler
	Update() http.Handler
	Destroy() http.Handler
}

// Comment represents a comment
type Comment struct {
	ID              int    `json:"id"`
	ItemID          int    `json:"item_id" db:"item_id"`
	Content         string `json:"content" db:"content"`
	CreatedBy       int    `json:"created_by" db:"created_by"`
	CreatedByMember string `json:"created_by_member" db:"created_by_member"`
	GravatarHash    string `json:"gravatar_hash" db:"gravatar_hash"`
	CreatedAt       string `json:"created_at" db:"created_at"`
}

// CommentHandler responsible for making comments
type CommentHandler interface {
	Index() http.Handler
	Create() http.Handler
	Show() http.Handler
	Destroy() http.Handler
	Update() http.Handler
}
