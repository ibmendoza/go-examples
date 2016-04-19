package main

import (
	"fmt"
	"github.com/nats-io/nats"
	"log"
	"runtime"
	"strconv"
	"sync/atomic"

	"github.com/fatih/set"
	"github.com/joyrexus/buckets"
)

var s = set.New()
var ops uint64 = 0
var bx, _ = buckets.Open("bolt")

func msgHandler(msg *nats.Msg) {
	data := string(msg.Data)

	bucket, _ := bx.New([]byte("bucket"))

	got, err := bucket.Get([]byte(data))
	if err != nil {
		// Put key/value into the bucket
		key, value := []byte(data), []byte("1")
		if err := bucket.Put(key, value); err != nil {
			fmt.Println("Error insert item: ", err)
		} else {
			fmt.Println(data, " 1")
		}
	} else {
		//increment
		key := []byte(data)
		v, _ := strconv.Atoi(string(got))
		v = v + 1
		value := []byte(strconv.Itoa(v))

		if err := bucket.Put(key, value); err != nil {
			fmt.Println("Error increment: ", err)
		} else {
			fmt.Println(data, " ", v)
		}
	}

	//add to set
	s.Add(data)

	atomic.AddUint64(&ops, 1)
	if ops == 300 {
		fmt.Println("set")
		fmt.Println(s)
		fmt.Println("set size")
		fmt.Println(s.Size())
		fmt.Println("ops")
		fmt.Println(ops)
		s.Clear()
	}
}

func endHandler(msg *nats.Msg) {
	fmt.Println("set")
	fmt.Println(s)
	fmt.Println("set size")
	fmt.Println(s.Size())
	fmt.Println("ops")
	fmt.Println(ops)
	s.Clear()
}

func main() {
	natsConnection, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal("Error connecting to NATS server")
	}

	defer bx.Close()

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Connected to " + nats.DefaultURL)

	fmt.Println("Subscribing to subject ")

	natsConnection.Subscribe("pipe", msgHandler)

	natsConnection.Subscribe("end", endHandler)

	fmt.Scanln()
}
