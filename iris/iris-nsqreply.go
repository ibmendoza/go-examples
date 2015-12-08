// http://itjumpstart.wordpress.com

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
