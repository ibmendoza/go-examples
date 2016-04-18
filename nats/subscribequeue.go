package main

import (
	"log"
	"runtime"

	"github.com/nats-io/nats"
)

func main() {

	// Create server connection: auth and no auth
	// natsConnection, _ := nats.Connect("nats://foo:bar@localhost:4222")
	natsConnection, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal("Error connecting to NATS server")
	}

	log.Println("Connected to " + nats.DefaultURL)

	// Subscribe to subject
	log.Println("Subscribing to subject ")
	natsConnection.Subscribe("pipe1", func(msg *nats.Msg) {

		// Handle the message
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
	})

	// Keep the connection alive
	runtime.Goexit()
}
