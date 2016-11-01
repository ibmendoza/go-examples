package main

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Items []Item

type Item struct {
	Id     int
	Name   string
	Dosage string
	Price  float32
	VATpct int
}

func main() {
	item := Item{}
	var slcItems Items

	db := sqlx.MustConnect("mysql", "username:password@tcp(127.0.0.1:3306)/databasename")

	rows, err := db.Queryx("SELECT id, name, dosage, sellingprice as price, vatpercentage as vatpct FROM products")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			break
		}
		log.Println(item)

		slcItems = append(slcItems, item)
	}

	b, err := json.Marshal(slcItems)

	if err != nil {
		log.Println(err)
	}
	// Convert bytes to string.
	log.Println(string(b))
}
