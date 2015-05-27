//http://nirbhay.in/2013/03/ajax-with-go/
/* tick-tock.go */
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Content for the main html page..
var page = `<html>
           <head>
             <script type="text/javascript"
               src="http://ajax.googleapis.com/ajax/libs/jquery/1.3.2/jquery.min.js">
             </script>
             <style>
               div {
                 font-family: "Times New Roman", Georgia, Serif;
                 font-size: 2em;
                 width: 13.3em;
                 padding: 8px 8px;
                 border: 2px solid #2B1B17;
                 border-radius: 10px;
                 color: #2B1B17;
                 text-shadow: 1px 1px #E5E4E2;
                 background: #FFFFFF;
               }
             </style>
           </head>
           <body>
             <h2>Go Timer (ticks every second!), the system clock</h2>
             <div id="output"></div>
             <script type="text/javascript">
               $(document).ready(function () {
                 $("#output").append("Waiting for system time..");
                 setInterval("delayedPost()", 1000);
               });
               function delayedPost() {
                 $.post("http://localhost:9999/gettime", "", function(data, status) {
                 $("#output").empty();
                 $("#output").append(data);
                 });
               }
             </script>
           </body>
         </html>`

// handler for the main page.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}

// handler to cater AJAX requests
func handlerGetTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
}

func main() {
	http.HandleFunc("/time", handler)
	http.HandleFunc("/gettime", handlerGetTime)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
