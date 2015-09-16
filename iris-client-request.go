package main

import (
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"time"
)

func main() {
	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}
	defer conn.Close()

	request := []byte("some request binary")
	if reply, err := conn.Request("echo", request, time.Second); err != nil {
		log.Printf("failed to execute request: %v.", err)
	} else {
		fmt.Printf("reply arrived: %v.", string(reply))
	}

	fmt.Scanln()
}
