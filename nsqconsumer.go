package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {

	/*
	   Below are the default port settings
	   nsqd listens at port 4150 (for TCP clients), 4151 (for HTTP clients)

	   nsqlookupd listens at port 4160 (for TCP clients), 4161 (for HTTP clients)

	   nsqadmin listens at port 4171 (for HTTP clients) or
	     to be specified (for go-nsq clients) q.ConnectToNSQLookupd("127.0.0.1:4161")

	   http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go/
	   $ nsqlookupd &
	   $ nsqd --lookupd-tcp-address=127.0.0.1:4160 &
	   $ nsqadmin --lookupd-http-address=127.0.0.1:4161 &
	*/

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)
		wg.Done()
		return nil
	}))

	//err := q.ConnectToNSQD("127.0.0.1:4150") - hardcoded

	err := q.ConnectToNSQLookupd("127.0.0.1:4161") //better
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}
