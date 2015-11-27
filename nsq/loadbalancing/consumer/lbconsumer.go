package main

import (
	"flag"
	"fmt"
	"github.com/itmarketplace/go-queue"
	"github.com/nsqio/go-nsq"
	"log"
	"runtime"
)

var lkp = flag.String("lkp", "", "IP address of nsqlookupd")

func main() {
	flag.Parse()

	c := queue.NewConsumer("test", "ch")

	c.Set("nsqlookupd", *lkp+":4161")
	c.Set("concurrency", runtime.GOMAXPROCS(runtime.NumCPU()))
	c.Set("max_attempts", 10)
	c.Set("max_in_flight", 150)
	c.Set("default_requeue_delay", "15s")

	c.Start(nsq.HandlerFunc(func(msg *nsq.Message) error {
		log.Println(string(msg.Body))

		return nil
	}))
	fmt.Scanln()
}
