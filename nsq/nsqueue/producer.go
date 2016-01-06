// https://github.com/crackcomm/nsqueue/blob/master/examples/producer/producer.go

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/crackcomm/nsqueue/producer"
)

var (
	amount   = flag.Int("amount", 20, "Amount of messages to produce every 100 ms")
	nsqdAddr = flag.String("nsqd", "127.0.0.1:4150", "nsqd tcp address")
)

func main() {
	flag.Parse()
	producer.Connect(*nsqdAddr)

	for _ = range time.Tick(100 * time.Millisecond) {
		fmt.Println("Ping...")
		for i := 0; i < *amount; i++ {
			body, _ := time.Now().MarshalBinary()
			producer.PublishAsync("latency-test", body, nil)
		}
	}
}
