package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/antonholmquist/jason"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/johnnylee/sqlxchain"

	_ "github.com/go-sql-driver/mysql"
)

type OrderDetail struct {
	IDOrder   int
	IDProduct int
	Price     float32
	Quantity  float32
	ItemTotal float32
}

func insertUserHandler(w http.ResponseWriter, r *http.Request) {
	v, err := jason.NewObjectFromReader(r.Body)

	if err != nil {
		log.Println("error in jason parse")
		log.Println(err)

		fmt.Fprintf(w, "Error in receiving JSON input for user")
		return
	}
	/*
		var obj = {
			id: 8,
			fname: 'John',
			mname: 'Paul',
			lname: 'McCartney',
			bdate: '2009-05-30',
			dtstarted: '2016-11-01'
		}
	*/

	sql := "insert into users(firstname, middlename, lastname, birthdate, datestarted) " +
		"values(?, ?, ?, ?, ?)"

	fname, _ := v.GetString("fname")
	mname, _ := v.GetString("mname")
	lname, _ := v.GetString("lname")
	bdate, _ := v.GetString("bdate")
	dtstarted, _ := v.GetString("dtstarted")

	err = dbx.Context().Begin().
		Exec(sql, fname, mname, lname, bdate, dtstarted).
		Commit().
		Err()

	if err != nil {
		log.Println("error in db tx")
		log.Println(err)

		fmt.Fprintf(w, "Error in saving user to database")
	} else {
		fmt.Fprintf(w, "")
	}
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {

	v, err := jason.NewObjectFromReader(r.Body)

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "Error in receiving JSON input for user")
		return
	}

	/*
		var obj = {
			id: 8,
			fname: 'John',
			mname: 'Paul',
			lname: 'McCartney',
			bdate: '2009-05-30',
			dtstarted: '2016-11-01'
		}
	*/

	sql := `update users 
			set firstname=?, 
			middlename=?, 
			lastname=?, 
			birthdate=?, 
			datestarted=?
			where id=?
			`
	id, _ := v.GetInt64("id")
	fname, _ := v.GetString("fname")
	mname, _ := v.GetString("mname")
	lname, _ := v.GetString("lname")
	bdate, _ := v.GetString("bdate")
	dtstarted, _ := v.GetString("dtstarted")

	err = dbx.Context().Begin().
		Exec(sql, fname, mname, lname, bdate, dtstarted, id).
		Commit().
		Err()

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "Error in updating user to database")
	} else {
		fmt.Fprintf(w, "")
	}
}

func deactivateUserHandler(w http.ResponseWriter, r *http.Request) {

	v, err := jason.NewObjectFromReader(r.Body)

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "Error in receiving JSON input for user")
		return
	}

	/*
		var obj = {
			id: 8
		}
	*/

	sql := `update users 
			set isactive=0
			where id=?
			`
	id, _ := v.GetInt64("id")

	err = dbx.Context().Begin().
		Exec(sql, id).
		Commit().
		Err()

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "Error in deactivating user to database")
	} else {
		fmt.Fprintf(w, "")
	}
}

func saveOrderHandler(w http.ResponseWriter, r *http.Request) {
	v, err := jason.NewObjectFromReader(r.Body)

	total, e := v.GetString("Total")
	if e != nil {
		log.Println(e)
		fmt.Fprintf(w, "Error in reading order total price")
		return
	}

	itemsObj, e := v.GetObjectArray("items")
	if e != nil {
		log.Println(e)
		fmt.Fprintf(w, "Error in reading order items")
		return
	}

	var id int64

	dbxContext := dbx.Context().Begin().
		Exec("INSERT INTO pos.order (iduser, dt, total) VALUES (1, now(), ?)", total).
		LastInsertId(&id)

	sql := "insert into pos.orderdetails (idorder, idproduct, price, quantity, itemtotal) " +
		"values (?, ?, ?, ?, ?)"

	var idproduct int64
	var price, qty, itemtotal string

	for _, item := range itemsObj {
		idproduct, _ = item.GetInt64("idproduct")
		price, _ = item.GetString("price")
		qty, _ = item.GetString("qty")
		itemtotal, _ = item.GetString("total")

		dbxContext.Exec(sql, id, idproduct, price, qty, itemtotal)
	}
	err = dbxContext.Commit().Err()

	if err != nil {
		log.Println(err)

		fmt.Fprintf(w, "Error in saving order to database")
	} else {
		fmt.Fprintf(w, "")
	}
}

type Items []Item

type Item struct {
	Id    int
	Price float32
	Name  string
	//Attr string
	//VATpct int
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := "%" + vars["name"] + "%"

	item := Item{}
	var slcItems Items

	rows, err := db.Queryx("SELECT id, sellingprice as price, "+
		"concat(name, ' - ', attr) as name FROM products "+
		"WHERE concat(name, ' - ', attr) like ?", name)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			break
		}
		//log.Println(item)

		slcItems = append(slcItems, item)
	}

	js, err := json.Marshal(slcItems)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert bytes to string.
	//log.Println(string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type Searches []Search

type Search struct {
	Id   int
	Name string
}

func UserSearchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := "%" + vars["name"] + "%"

	search := Search{}
	var slcSearch Searches

	rows, err := db.Queryx("SELECT id, "+
		"concat(identifier, ' ', firstname, ' ', middlename, ' ', lastname) as name FROM users "+
		"WHERE concat(identifier, ' ', firstname, ' ', middlename, ' ', lastname) like ?", name)

	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.StructScan(&search)
		if err != nil {
			break
		}

		//log.Println(search)

		slcSearch = append(slcSearch, search)
	}

	js, err := json.Marshal(slcSearch)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	item := Item{}
	var slcItems Items

	rows, err := db.Queryx("SELECT id, sellingprice as price, " +
		"concat(name, ' - ', attr) as name FROM products")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			break
		}
		//log.Println(item)

		slcItems = append(slcItems, item)
	}

	js, err := json.Marshal(slcItems)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Convert bytes to string.
	//log.Println(string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

var dbx *sqlxchain.SqlxChain

var db *sqlx.DB
