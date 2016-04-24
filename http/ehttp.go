//https://github.com/creack/ehttp
//Due to http limitation, we can send the headers only once. If some data has been sent prior to the error, 
//then nothing gets send to the client, the error gets logged on the server side.

package main

import (
    "log"
    "net/http"

    "github.com/creack/ehttp"
    "github.com/creack/ehttp/ehttprouter"
    "github.com/julienschmidt/httprouter"
)

func hdlr(w http.ResponseWriter, req *http.Request, p httprouter.Params) error {
    return ehttp.NewErrorf(418, "fail")
}

func main() {
    router := httprouter.New()
    router.GET("/", ehttprouter.MWError(hdlr))
    log.Fatal(http.ListenAndServe(":8080", router))
}
