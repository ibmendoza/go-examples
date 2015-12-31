//https://itjumpstart.wordpress.com/2015/12/22/reverse-proxy-using-project-iris/

package main

import (
	"flag"
	"github.com/ibmendoza/go-lib"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ipEvent struct {
}

func (t *ipEvent) Init(conn *iris.Connection) error {
	return nil
}

func (t *ipEvent) HandleBroadcast(msg []byte) {
}

func (t *ipEvent) HandleRequest(req []byte) ([]byte, error) {
	ip, _ := lib.GetIPAddress()
	return []byte(ip), nil
}

func (t *ipEvent) HandleDrop(reason error) {
}

func (t *ipEvent) HandleTunnel(tun *iris.Tunnel) {
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

	if err != nil {
		log.Fatal("Could not connect to local nsqd")
	}

	var ipHandler = &ipEvent{}

	service, err := iris.Register(55555, *cluster, ipHandler, nil)
	if err != nil {
		log.Fatalf("failed to register to the Iris relay: %v.", err)
	}

	for {
		select {
		case <-sigChan:
			service.Unregister()
			conn.Close()
			os.Exit(1)
		}
	}
}
