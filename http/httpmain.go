package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ibmendoza/salt"
	"github.com/jmoiron/sqlx"
	"github.com/johnnylee/sqlxchain"
	"github.com/vaughan0/go-ini"
)

func main() {
	file, err := ini.LoadFile("license.conf")
	if err != nil {
		log.Fatal("Missing license.conf")
	}

	token, ok := file.Get("LICENSE", "token")
	if !ok {
		log.Fatal("Missing token in license.conf")
	}

	key, ok := file.Get("LICENSE", "thekey")
	if !ok {
		log.Fatal("Missing key in license.conf")
	}

	_, err = salt.Verify(token, key)
	if err != nil {
		log.Println("Invalid license. Please contact your vendor")
		log.Fatal(err)
	}

	dbx, err = sqlxchain.New("mysql", "user:pswd@tcp(127.0.0.1:3306)/db")

	if err != nil {
		log.Fatal(err)
	}

	db = sqlx.MustConnect("mysql", "user:pswd@tcp(127.0.0.1:3306)/db")

	r := mux.NewRouter()

	// This will serve files under http://localhost:8080/static
	fs := http.FileServer(http.Dir("./current"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/vproducts/{name}", ProductsHandler)

	r.HandleFunc("/searchuser/{name}", UserSearchHandler)

	r.HandleFunc("/saveorder", saveOrderHandler)

	r.HandleFunc("/testproducts", ProductsHandler)

	r.HandleFunc("/insertuser", insertUserHandler)

	r.HandleFunc("/updateuser", updateUserHandler)

	r.HandleFunc("/deactivateuser", deactivateUserHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server up and running...")
	log.Fatal(srv.ListenAndServe())
}
