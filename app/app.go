package app

import (
  "net/http"
)

type Bag struct {
	Id int
	Name string
	Brand string
	Image_url string
}

type BagService interface {
	Bag(id int) (*Bag, error)
	Bags() ([]*Bag, error)
}
/*
type BagHTTPService interface {
  bagService *BagService
  Index() ([]*Bag, error)
	Show(id int) (*Bag, error)
}
*/

type BagHTTPService interface {
  Index() http.HandlerFunc
  Show(id int) 
}

