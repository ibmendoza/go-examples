package main

import (
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"runtime"
)

type topicEvent struct {
	topic string
}

func (t topicEvent) HandleEvent(event []byte) {

	log.Println("received test event: with payload data as: " + string(event))
}

func main() {
	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	var topicHandler = topicEvent{topic: "test"}

	//Subscribe(topic string, handler TopicHandler, limits *TopicLimits)
	sub := conn.Subscribe("test", topicHandler, &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  1024 * 1024,
	})

	if sub != nil {
		conn.Log.Debug("error in subscribe")
	}

	defer conn.Close()

	fmt.Scanln()
}
