//https://github.com/didip/tollbooth

package main

import (
    "github.com/didip/tollbooth"
    "net/http"
    "time"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    // Create a request limiter per handler.
    http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, time.Second), HelloHandler))
    http.ListenAndServe(":12345", nil)
}
