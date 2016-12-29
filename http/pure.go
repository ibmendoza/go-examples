package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/pure"
	
	//mw "github.com/go-playground/pure/examples/middleware/logging-recovery"
)

func main() {

	p := pure.New()
	//p.Use(mw.LoggingAndRecovery(true))

	p.Get("/user/:id/:id2", userHandler)

	http.ListenAndServe(":3007", p.Serve())
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// extract params like so
	rv := pure.RequestVars(r) // done this way so only have to extract from context once, read above

	fmt.Fprintf(w, rv.URLParam("id")+" - "+rv.URLParam("id2"))
}
