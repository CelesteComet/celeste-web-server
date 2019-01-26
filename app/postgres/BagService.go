package postgres

import (
	"github.com/CelesteComet/celeste-web-server/app"
	_ "github.com/lib/pq"
	"log"
	"github.com/jmoiron/sqlx"
)

var _ app.BagService = &BagService{}

type BagService struct {
	DB *sqlx.DB
}

func (b *BagService) Bag(id int) (Bag *app.Bag, err error) {
	bag := app.Bag{}
	rows, err := b.DB.Queryx("select * from bag where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.StructScan(&bag); err != nil {
			log.Fatal(err)
		}
	}
	return &bag, err
}

func (b *BagService) Bags() (Bags []*app.Bag, err error) {
	bags := []*app.Bag{}
	rows, err := b.DB.Queryx("select * from bag")
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
	return bags, err
}
