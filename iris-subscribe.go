package main

import (
	"flag"
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

type topicEvent struct{}

var start = time.Now()
var ops uint64 = 0
var numbPtr = flag.Int("msg", 10000, "number of messages (default: 10000)")

func (t topicEvent) HandleEvent(event []byte) {

	log.Println(string(event))

	atomic.AddUint64(&ops, 1)
	if ops == uint64(*numbPtr) {
		elapsed := time.Since(start)
		log.Printf("Time took %s", elapsed)
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

	var topicHandler = new(topicEvent)

	conn.Subscribe("test", topicHandler, &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	})

	defer conn.Close()

	fmt.Scanln()
}
