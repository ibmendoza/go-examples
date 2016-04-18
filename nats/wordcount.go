//count the frequency of each word
//and store it to onecache or cassandra
//Example of separating producer from consumer

package main

import (
	"bufio"
	"fmt"
	"github.com/nats-io/nats"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	f, _ := os.Open("socialcancer.txt")

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanWords)

	cnt := 0

	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("Error connecting to NATS server")
	}

	defer natsConnection.Close()
	fmt.Println("Connected to NATS server: " + nats.DefaultURL)

	start := time.Now()

	var msg *nats.Msg

	for scanner.Scan() {
		cnt++

		//line := scanner.Text()

		r := strings.NewReplacer(".", "", ":", "", ",", "", "'", "", ";", "",
			"[", "", "]", "", "#", "", "?", "")
		newline := r.Replace(scanner.Text())

		fmt.Println(newline)

		/*
		   Output:

		   The
		   Project
		   Gutenberg
		   EBook
		   of
		   The
		   Social
		   Cancer,
		*/

		msg = &nats.Msg{Subject: "pipe1", Data: []byte(newline)}
		natsConnection.PublishMsg(msg)

		if cnt == 300 {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elapsed time: ", elapsed)
}
