//https://gist.github.com/alexedwards/4d20c505f389597c3360
package main

import (
	"fmt"
	"github.com/alexedwards/stack"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	stdStack := stack.New(tokenMiddleware)

	router.GET("/hello/:forename", InjectParams(stdStack.Then(helloHandler)))
	http.ListenAndServe(":3000", router)
}

func InjectParams(hc stack.HandlerChain) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		newHandlerChain := stack.Inject(hc, "params", ps)
		newHandlerChain.ServeHTTP(w, r)
	}
}

func tokenMiddleware(ctx *stack.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx.Put("token", "c9e452805dee5044ba520198628abcaa")
		next.ServeHTTP(w, r)
	})
}

func helloHandler(ctx *stack.Context, w http.ResponseWriter, r *http.Request) {
	params, ok := ctx.Get("params").(httprouter.Params)
	if !ok {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	token, ok := ctx.Get("token").(string)
	if !ok {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	fmt.Fprintf(w, "Hello %s. Your token is %s.", params.ByName("forename"), token)
}
