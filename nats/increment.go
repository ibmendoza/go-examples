package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/nats-io/nats"
)

var mc = memcache.New("192.168.0.130:11211")

func msgHandler(msg *nats.Msg) {
	data := string(msg.Data)
	_, err := mc.Get(data)
	if err != nil {
		mc.Set(&memcache.Item{
			Key:        data,
			Value:      []byte("0"),
			Expiration: 0,
		})
	}
	fmt.Println(data)
	var cnt uint64
	cnt, err = mc.Increment(data, 1)
	if err != nil {
		log.Println("Error in increment, ", err)
	}

	fmt.Println(cnt)
}

func main() {

	// Create server connection: auth and no auth
	// natsConnection, _ := nats.Connect("nats://foo:bar@localhost:4222")
	natsConnection, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal("Error connecting to NATS server")
	}

	fmt.Println("Connected to " + nats.DefaultURL)

	// Subscribe to subject
	fmt.Println("Subscribing to subject ")
	natsConnection.QueueSubscribe("pipe1", "socialcancer", msgHandler)

	// Keep the connection alive
	runtime.Goexit()
}
