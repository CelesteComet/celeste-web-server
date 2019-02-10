package http

import (
	"github.com/CelesteComet/celeste-web-server/app"
	"github.com/CelesteComet/celeste-web-server/app/postgres"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"strings"
	// "log"
	"net/http"
	"strconv"
	"gopkg.in/matryer/respond.v1"
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
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}
		respond.With(w, r, http.StatusOK, bags)
	})
}

func (h *BagHandler) GetBag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		i, err := strconv.Atoi(vars["n"])
		bag, err := h.BagService.Bag(i)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}
		respond.With(w, r, http.StatusOK, bag)
	})
}

func (h *BagHandler) GetUserBags() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// To return
		bags := []*app.Bag{}

		// Parse params
		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["userID"])

		// Query DB
		rows, err := h.BagService.DB.Queryx("select * from bag where created_by = $1", userID)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}
		defer rows.Close()

		// Prepare Data
		for rows.Next() {
			bag := app.Bag{}
			if err := rows.StructScan(&bag); err != nil {
				respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			}
			bags = append(bags, &bag)
		}

		respond.With(w, r, http.StatusOK, bags)
	})
}

func (h *BagHandler) GetBagWithTag() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// To return
		bags := []*app.Bag{}

		// Parse Params
		tags := r.URL.Query().Get("tags")
		stringSlice := strings.Split(tags, " ")
		arg := map[string]interface{}{
		    "tagLength": len(stringSlice),
		    "tags": stringSlice,
		}

		// Prepare Query
		queryString := `select bag.* from bag
 where bag.id in (
 select bag.id from bag 
 join bagtags on bag.id = bagtags.bagId 
 join tags on tags.id = bagtags.tagId 
 where tags.name in (:tags) group by bag.id having COUNT(bag.id) = :tagLength);`

		query, args, err := sqlx.Named(queryString, arg)
		query, args, err = sqlx.In(query, args...)
		query = h.BagService.DB.Rebind(query)
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}

		// Query DB
		rows, err := h.BagService.DB.Queryx(query, args...) 
		if err != nil {
			respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
		}
		defer rows.Close()

		// Prepare Data
		for rows.Next() {
			bag := app.Bag{}
			if err := rows.StructScan(&bag); err != nil {
				respond.With(w, r, http.StatusInternalServerError, []string{err.Error()})
			}
			bags = append(bags, &bag)
		}

		respond.With(w, r, http.StatusOK, bags)
	})
}

