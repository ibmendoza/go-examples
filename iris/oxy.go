// Reverse Proxy backed by Iris (no need for etcd, Consul, Zookeeper, etc)

// https://itjumpstart.wordpress.com/2015/12/22/reverse-proxy-using-project-iris/

// nginx -> :8080 (reverse proxy) -> :63450 (Web app servers)

package main

import (
	"flag"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"net/http"
	"time"
)

var cluster = flag.String("cluster", "", "Iris cluster name")

//https://medium.com/@matryer/the-http-handlerfunc-wrapper-technique-in-golang-c60bf76e6124
func wrapHandlerFunc(c *iris.Connection, f *forward.Forwarder) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		ip := []byte("ip")
		var ipaddr string
		if reply, err := c.Request(*cluster, ip, time.Second*5); err != nil {
			log.Printf("failed to execute request: %v.", err)
		} else {
			ipaddr = string(reply)
		}

		// let us forward this request to another server
		req.URL = testutils.ParseURI("http://" + ipaddr + ":63450")
		f.ServeHTTP(w, req)
	}
}

func main() {
	flag.Parse()

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}
	defer conn.Close()

	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()

	redirect := wrapHandlerFunc(conn, fwd)

	// that's it! our reverse proxy is ready!

        // nginx -> :8080 (reverse proxy) -> :63450 (Web app servers)

	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	s.ListenAndServe()
}
