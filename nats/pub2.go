package main

import (
	"flag"
	"github.com/nats-io/nats"
	"log"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()1234567890")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	//gnatsd -m 8222
	flag.Parse()

	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()
	log.Println("Connected to NATS server: " + nats.DefaultURL)

	start := time.Now()

	msg := &nats.Msg{Subject: "foo", Reply: "bar", Data: []byte(randSeq(320))}

	cnt := 0

	for {
		cnt++
		natsConnection.PublishMsg(msg)
		msg = &nats.Msg{Subject: "foo", Reply: "bar", Data: []byte(randSeq(320))}
		log.Println(cnt)
	}

	elapsed := time.Since(start)
	log.Printf("Time took %s", elapsed)
}
