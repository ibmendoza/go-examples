package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/set"
	"github.com/nats-io/nats"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var s = set.New()

	f, _ := os.Open("socialcancer.txt")
	defer f.Close()

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
			"[", "", "]", "", "#", "", "?", "", "(", "", ")", "", "\"", "")
		newline := r.Replace(scanner.Text())

		fmt.Println(newline)
		s.Add(newline)

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

		msg = &nats.Msg{Subject: "pipe", Data: []byte(newline)}
		natsConnection.PublishMsg(msg)

		if cnt == 300 {
			break
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Elapsed time: ", elapsed)

	msg = &nats.Msg{Subject: "end", Data: []byte("end")}
	natsConnection.PublishMsg(msg)

	fmt.Println("Count: ", cnt)

	fmt.Println("set")
	fmt.Println(s)
	fmt.Println("set size")
	fmt.Println(s.Size())

	fmt.Scanln()
}
