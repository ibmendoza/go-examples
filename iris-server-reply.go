package main

import (
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
)

type EchoHandler struct{}

func (b *EchoHandler) Init(conn *iris.Connection) error {
	return nil
}

func (b *EchoHandler) HandleBroadcast(msg []byte) {
}

func (b *EchoHandler) HandleRequest(req []byte) ([]byte, error) {
	return req, nil
}

func (b *EchoHandler) HandleDrop(reason error) {
}

func (b *EchoHandler) HandleTunnel(tun *iris.Tunnel) {
}

func main() {
	service, err := iris.Register(55555, "echo", new(EchoHandler), nil)
	if err != nil {
		log.Fatalf("failed to register to the Iris relay: %v.", err)
	}
	defer service.Unregister()

	fmt.Scanln()
}
