//http://www.alexedwards.net/blog/a-recap-of-request-handling

package main

import (
  "log"
  "net/http"
  "time"
)

func timeHandler(format string) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
  return http.HandlerFunc(fn)
}

func main() {
  // Note that we skip creating the ServeMux...

  var format string = time.RFC1123
  th := timeHandler(format)

  // We use http.Handle instead of mux.Handle...
  http.Handle("/time", th)

  log.Println("Listening...")
  // And pass nil as the handler to ListenAndServe.
  http.ListenAndServe(":3000", nil)
}
