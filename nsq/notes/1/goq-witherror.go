package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"github.com/itmarketplace/go-queue"
	"github.com/nsqio/go-nsq"
	"runtime"
	"time"
)

var start = time.Now()
var ops uint64 = 0
var numbPtr = flag.Int("msg", 10000, "number of messages (default: 10000)")
var lkp = flag.String("lkp", "", "IP address of nsqlookupd1")

func main() {

	flag.Parse()

	c := queue.NewConsumer("test", "ch")

	c.Set("nsqlookupd", *lkp+":4161")
	c.Set("concurrency", runtime.GOMAXPROCS(runtime.NumCPU()))
	c.Set("max_attempts", 10)
	c.Set("max_in_flight", 150)
	c.Set("default_requeue_delay", "15s")

	c.Start(nsq.HandlerFunc(func(msg *nsq.Message) error {

		buf := bytes.NewBuffer(msg.Body)
		myInt, err := binary.ReadVarint(buf)
		if err != nil {
			return errors.New("invalid number")
		} else {
			//simulate error
			if myInt%2 == 0 {
				return errors.New("even")
			} else {
				return nil
			}
		}
	}))

	fmt.Scanln()
}
