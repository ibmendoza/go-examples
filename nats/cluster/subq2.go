package main

import (
	"github.com/nats-io/nats"
	"log"
	"runtime"
	"sync/atomic"
)

var ops uint64 = 0

func main() {

	// Create server connection: auth and no auth
	//natsConnection, err := nats.Connect("nats://foo:bar@localhost:4222")
	//natsConnection, err := nats.Connect(nats.DefaultURL)

	//natsConnection, err := nats.Connect("nats://192.168.99.100:4222")

	natsConnection, err := nats.Connect("nats://192.168.99.101:5222")

	if err != nil {
		log.Fatal("NATS server not running")
	}

	//log.Println("Connected to nats://192.168.56.1:4222")

	// Subscribe to subject
	log.Printf("Subscribing to subject 'telegraf'\n")
	natsConnection.QueueSubscribe("telegraf", "queue", func(msg *nats.Msg) {

		// Handle the message
		//log.Printf("Received message '%s\n", string(msg.Data)+"'")
		atomic.AddUint64(&ops, 1)
		log.Println(ops)
	})

	// Keep the connection alive
	runtime.Goexit()
}
