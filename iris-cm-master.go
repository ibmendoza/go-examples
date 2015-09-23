package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"net/http"
	"runtime"
)

type cmdReplyEvent struct{}

func (c cmdReplyEvent) HandleEvent(event []byte) {
	log.Println(string(event))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))

	conn.Publish("cmd", []byte(ps.ByName("name")))
}

var conn *iris.Connection

func main() {

	topicLimits := &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	conn, err = iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	var cmdreplyHandler = new(cmdReplyEvent)

	conn.Subscribe("reply", cmdreplyHandler, topicLimits)

	//http handlers
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	//https://github.com/julienschmidt/httprouter/issues/7
	router.ServeFiles("/static/*filepath", http.Dir("C:/static"))

	log.Fatal(http.ListenAndServe(":8080", router))

	defer conn.Close()
}
