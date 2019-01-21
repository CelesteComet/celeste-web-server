package postgres

import (
	"database/sql"
  "log"
	_ "github.com/lib/pq"
	"github.com/CelesteComet/celeste-web-server/app"
)

var _ app.BagService = &BagService{}

type BagService struct {
	DB *sql.DB
}

func (mBagService *BagService) Bag(id int) (mBag *app.Bag, err error) {
  bag := app.Bag{}
	rows, err := mBagService.DB.Query("select * from bag where id = $1", id)
	if err != nil {
    log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
    if err := rows.Scan(&bag.Id, &bag.Name, &bag.Brand, &bag.Image_url); err != nil {
      log.Fatal(err)
    }
	}
	return &bag, err
}

func (mBagService *BagService) Bags() (mBags []*app.Bag, err error) {
  bags := []*app.Bag{}
	rows, err := mBagService.DB.Query("select * from bag") 
	if err != nil {
    log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		bag := app.Bag{}
    if err := rows.Scan(&bag.Id, &bag.Name, &bag.Brand, &bag.Image_url); err != nil {
      log.Fatal(err)
    }
		bags = append(bags, &bag)
	}
	return bags, err
}




