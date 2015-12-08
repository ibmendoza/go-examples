## NSQ and Iris

This is an example of load balancing two NSQ servers (nsqd) in an Iris cluster. The request-reply pattern in Iris is designed for 
load balancing. Get Iris version 0.3.2 [here](https://github.com/ibmendoza/project-iris/releases)

Structure:

- 192.168.56.101 (Iris request client + Message Producer)

- 192.168.56.102 (Iris reply client + nsqd + NSQ Producer)

- 192.168.56.103 (Iris reply client + nsqd + NSQ Producer)


Content of nsq.sh

```./nsqlookupd & ./nsqd --lookupd-tcp-address=127.0.0.1:4160 & ./nsqadmin --lookupd-http-address=127.0.0.1:4161```

### 192.168.56.101 (Producer)

Run the following:

- iris -net cluster -rsa /home/id_rsa

- request


```go
//request.go

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
```

### 192.168.56.102 and 192.168.56.103 (NSQ Producer)

Run the following:

- iris -net cluster -rsa /home/id_rsa

- nsq.sh

- request

```go
//reply.go

package main

import (
	"flag"
	"github.com/ibmendoza/go-lib"
	"github.com/nsqio/go-nsq"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type topicEvent struct {
	nsqp *nsq.Producer
}

func (t *topicEvent) Init(conn *iris.Connection) error {
	return nil
}

func (t *topicEvent) HandleBroadcast(msg []byte) {
}

func (t *topicEvent) HandleRequest(req []byte) ([]byte, error) {
	t.nsqp.Publish("topic", req)
	ip, _ := lib.GetIPAddress()
	return []byte(ip), nil
}

func (t *topicEvent) HandleDrop(reason error) {
}

func (t *topicEvent) HandleTunnel(tun *iris.Tunnel) {
}

var cluster = flag.String("cluster", "", "Iris cluster name")

func main() {
	flag.Parse()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	}

	config := nsq.NewConfig()
	w, err := nsq.NewProducer(":4150", config)
	if err != nil {
		log.Fatal("Could not connect to local nsqd")
	}

	var topicHandler = &topicEvent{nsqp: w}

	service, err := iris.Register(55555, *cluster, topicHandler, nil)
	if err != nil {
		log.Fatalf("failed to register to the Iris relay: %v.", err)
	}

	for {
		select {
		case <-sigChan:
			w.Stop()
			service.Unregister()
			conn.Close()
			os.Exit(1)
		}
	}
}
```

In this sample, nsqd from 192.168.56.102 and 192.168.56.103 got 57 and 43 messages respectively.
