package main

import (
	"github.com/nats-io/nats"
	"log"
	"runtime"
	"sync/atomic"
)

var ops uint64 = 0

func main() {

	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// Subscribe to subject
	log.Printf("Subscribing to subject 'telegraf'\n")
	natsConnection.Subscribe("telegraf", func(msg *nats.Msg) {

		// Handle the message
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
		atomic.AddUint64(&ops, 1)
		log.Println(ops)
	})

	// Keep the connection alive
	runtime.Goexit()
}
