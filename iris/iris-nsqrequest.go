// http://itjumpstart.wordpress.com

package main

import (
	"flag"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var numbPtr = flag.Int("msg", 100, "number of messages (default: 100)")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	var s []byte

	//simulate producer
	for i := 1; i <= *numbPtr; i++ {
		//conn.Publish("topic", []byte(randSeq(160)))

		//request/reply load balances among all servers in the cluster
		s, err = conn.Request("cluster", []byte(randSeq(160)), time.Second*2)
		if err != nil {
			log.Println(err)
		}
		log.Println("ip addr " + string(s))
	}

	for {
		select {
		case <-sigChan:
			conn.Close()
			os.Exit(1)
		}
	}
}
