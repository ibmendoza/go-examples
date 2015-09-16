package main

import (
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

//applicable when there's one publisher and one subscriber only
func (t topicEvent) HandleEvent(event []byte) {
	//log.Println("received test event: with payload data as: " + string(event))
	atomic.AddUint64(&ops, 1)
	if ops == 1000000 {
		elapsed := time.Since(start)
		log.Printf("Time took %s", elapsed)
	}
}

func main() {
	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	var topicHandler = new(topicEvent)

	//Subscribe(topic string, handler TopicHandler, limits *TopicLimits)
	conn.Subscribe("test", topicHandler, &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	})

	defer conn.Close()

	fmt.Scanln()
}
