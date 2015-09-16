package main

import (
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
)

func main() {
	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}
	defer conn.Close()

	broadcast := []byte("broadcast message")
	if err := conn.Broadcast("echo", broadcast); err != nil {
		log.Printf("failed to execute broadcast: %v.", err)
	} else {
		fmt.Println("broadcast successful")
	}

	fmt.Scanln()
}
