package main

import (
	"fmt"
	"github.com/ibmendoza/go-lib"
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {

	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	config := nsq.NewConfig()

	ipaddr, _ := lib.GetIPAddress()

	w, err := nsq.NewProducer(ipaddr+":4150", config)

	if err != nil {
		log.Fatal("Could not connect")
	}

	for i := 0; i < 52; i++ {
		w.Publish("test", []byte(string(letters[i])))
	}

	w.Stop()

	fmt.Scanln()
}
