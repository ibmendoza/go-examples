//https://itjumpstart.wordpress.com

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"username:password@tcp(127.0.0.1:3306)/databasename")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	defer db.Close()

	var (
		id     int
		name   string
		dosage string
		price  float32
		vat    int
	)

	sql := "select id, name, dosage, sellingprice, vatpercentage " +
		"from products order by name, dosage"

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &dosage, &price, &vat)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, " ", name, " ", dosage, " ", price, " ", vat)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Scanln()
}
