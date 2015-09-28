package main

import (
	"github.com/nsqio/go-nsq"
	"log"
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

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
