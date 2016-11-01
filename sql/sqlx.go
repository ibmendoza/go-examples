package main

import (
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Item struct {
	Id     int
	Name   string
	Dosage string
}

func main() {
	item := Item{}

	db := sqlx.MustConnect("mysql", "username:password@tcp(127.0.0.1:3306)/databasename")

	rows, err := db.Queryx("SELECT id, name, dosage FROM products")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			break
		}
		log.Println(item)
	}

}
