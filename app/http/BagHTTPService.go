package mhttp 

import (
	"log"
	"net/http"
	"encoding/json"
  "github.com/CelesteComet/celeste-web-server/app"
  "github.com/CelesteComet/celeste-web-server/app/postgres"
)

type BagHTTPService struct {
  BagService postgres.BagService
}

// Make sure that struct BagHTTPService implements app.BagHTTPService
var _ app.BagHTTPService = &BagHTTPService{}

func (b *BagHTTPService) Index() http.HandlerFunc {

	bags, err := b.BagService.Bags()
	if err != nil {
	  log.Println(err)
	}

	log.Println(bags)

	bagsJson, err := json.Marshal(bags)
	if err != nil {
	  log.Println(err)
	}


  return func(w http.ResponseWriter, r *http.Request) {
		log.Println("EH")
    w.Header().Set("Content-Type", "application/json") 
		w.WriteHeader(http.StatusOK)
		w.Write(bagsJson)
  }
}

func (b *BagHTTPService) Show(id int) {
  log.Println("INDEXING BAGS")
}
