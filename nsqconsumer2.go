////http://adampresley.com/2015/02/16/waiting-for-goroutines-to-finish-running-before-exiting.html
package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	/*
	* When SIGINT or SIGTERM is caught write to the quitChannel
	 */
	quitChannel := make(chan os.Signal)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}

	//start of consumer code block
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		wg.Add(1)
		defer wg.Done()
		log.Printf("Got a message: %v", message)
		return nil
	}))
	//end of consumer code block

	//err := q.ConnectToNSQD("127.0.0.1:4150") - hardcoded

	err := q.ConnectToNSQLookupd("127.0.0.1:4161") //better
	if err != nil {
		log.Fatal("Could not connect")
	}

	/*
	* Wait until we get the quit message
	 */
	<-quitChannel

	log.Println("Received quit. Sending shutdown and waiting on goroutines...")

	/*
	* Block until wait group counter gets to zero
	 */
	wg.Wait()
	log.Println("Done.")
}
