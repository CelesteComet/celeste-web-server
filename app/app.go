package app

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

type BagHTTPService interface {
  Index() ([]*Bag, error)
	Show(id int) (*Bag, error)
}
