package main

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

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

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	item := Item{}
	var slcItems Items

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

	js, err := json.Marshal(slcItems)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert bytes to string.
	log.Println(string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var db *sqlx.DB

func main() {

	db = sqlx.MustConnect("mysql", "username:pswd@tcp(127.0.0.1:3306)/dbname")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
