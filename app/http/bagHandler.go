package http

import (
	"github.com/CelesteComet/celeste-web-server/app"
	"github.com/CelesteComet/celeste-web-server/app/postgres"
	"github.com/CelesteComet/celeste-web-server/pkg/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

// BAG HANDLER
type BagHandler struct {
	BagService postgres.BagService
}

var _ app.BagHandler = &BagHandler{}

func (h *BagHandler) GetBags() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bags, err := h.BagService.Bags()
		if err != nil {
			log.Println(err)
		}
		json.Respond(w, r, bags, 200)
	})
}

func (h *BagHandler) GetBag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["n"])
		bag, err := h.BagService.Bag(i)
		if err != nil {
			log.Println(err)
		}
		json.Respond(w, r, bag, 200)
	})
}

func (h *BagHandler) GetUserBags() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userID"])
		rows, err := h.BagService.DB.Queryx("select * from bag where created_by = $1", userID)
		bags := []*app.Bag{}
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			bag := app.Bag{}
			if err := rows.StructScan(&bag); err != nil {
				log.Fatal(err)
			}
			bags = append(bags, &bag)
		}

		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		json.Respond(w, r, bags, 200)
	})
}

