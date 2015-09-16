package main

import (
	"flag"
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()1234567890")
var numbPtr = flag.Int("msg", 10000, "number of messages (default: 10000)")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	flag.Parse()

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	start := time.Now()

	for i := 1; i <= *numbPtr; i++ {
		conn.Publish("test", []byte(randSeq(320)))
	}

	elapsed := time.Since(start)
	log.Printf("Time took %s", elapsed)

	defer conn.Close()

	fmt.Scanln()
}
