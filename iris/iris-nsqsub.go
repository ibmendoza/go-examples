package main

import (
	"github.com/ibmendoza/msgq"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type topicEvent struct {
	nsqp *nsq.Producer
}

func (t topicEvent) HandleEvent(event []byte) {
	t.nsqp.Publish("topic", event)
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var iris = new(msgq.Iris)
	err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(":4150", config)
	if err != nil {
		log.Fatal("Could not connect to local nsqd")
	}

	var topicHandler = &topicEvent{nsqp: w}

	iris.Subscribe("topic", topicHandler)

	for {
		select {
		case <-sigChan:
			w.Stop()
			iris.Close()
			os.Exit(1)
		}
	}
}
