//https://gist.github.com/tsenart/5fc18c659814c078378d
package main

import (
	"net/http"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func helloHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    		var name string
    		// Execute the query.
    		row := db.QueryRow("SELECT myname FROM mytable")
    		if err := row.Scan(&name); err != nil {
        		http.Error(w, err.Error(), 500)
        		return
    		}
    		// Write it back to the client.
    		fmt.Fprintf(w, "hi %s!\n", name)
    	})
}

func withMetrics(l *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		l.Printf("%s %s took %s", r.Method, r.URL, time.Since(began))
	})
}

func main() {
	// Open our database connection.
	db, err := sql.Open("postgres", "…")
	if err != nil {
		log.Fatal(err)
	}
	// Create our logger
	logger := log.New(os.Stdout, "", 0)
	// Register our handler.
	http.Handle("/hello", helloHandler(db))
	// Register our handler with metrics logging
	http.Handle("/hello_again", withMetrics(logger, helloHandler(db)))
	http.ListenAndServe(":8080", nil)
}
