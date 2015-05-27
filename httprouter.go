package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"runtime"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	//https://github.com/julienschmidt/httprouter/issues/7
	router.ServeFiles("/static/*filepath", http.Dir("C:/static"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
