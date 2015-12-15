package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"gopkg.in/tylerb/graceful.v1"
	"log"
	"net/http"
	"time"
)

type EventPayload struct {
	Event   string
	Payload []byte
}

type appHandler struct {
	conn *iris.Connection
}

func (ah appHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var kv EventPayload
	err := decoder.Decode(&kv)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(kv.Event)

	var jsonkv []byte

	jsonkv, err = json.Marshal(kv)
	if err != nil {
		log.Println(err)
		return
	}

	//request/reply load balances among all servers in the cluster
	_, err = ah.conn.Request("cluster", jsonkv, time.Second*60)
	if err != nil {
		log.Println(err)
	}
}

func main() {

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	app_Handler := &appHandler{conn: conn}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.Handle("/hook", app_Handler)

	graceful.Run(":3001", 10*time.Second, mux)
}
