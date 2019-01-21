 package http

import (
  "fmt"
	"net/http"
  "github.com/CelesteComet/celeste-web-server/app"
  "github.com/CelesteComet/celeste-web-server/app/postgres"
)

var _ app.BagHTTPService = &BagHTTPService{}

type BagHTTPService struct {}

type indexHandler struct {}
type showHandler struct {}

func (mBagHTTPService *BagHTTPService) Index() indexHandler { 
  indexHandler := indexHandler{}
  return indexHandler  
}


func (mBagHTTPService *BagHTTPService) Show(id int) showHandler {
  showHandler := showHandler{}
  return showHandler{}
}

