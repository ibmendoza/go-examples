package main

import (
	"flag"
	"fmt"
	"github.com/gdamore/mangos"
	"github.com/gdamore/mangos/protocol/pub"
	"github.com/gdamore/mangos/protocol/sub"
	"github.com/gdamore/mangos/transport/ipc"
	"github.com/gdamore/mangos/transport/tcp"
	"log"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()1234567890")
var numbPtr = flag.Int("msg", 10, "number of messages (default: 10000)")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func die(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func main() {
	flag.Parse()

	var sock mangos.Socket
	var err error
	if sock, err = pub.NewSocket(); err != nil {
		die("can't get new pub socket: %s", err)
	}
	sock.AddTransport(ipc.NewTransport())
	sock.AddTransport(tcp.NewTransport())
	if err = sock.Listen("tcp://127.0.0.1:55555"); err != nil {
		die("can't listen on pub socket: %s", err.Error())
	}

	//sub
	var subscriber mangos.Socket

	if subscriber, err = sub.NewSocket(); err != nil {
		die("can't get new sub socket: %s", err.Error())
	}
	subscriber.AddTransport(ipc.NewTransport())
	subscriber.AddTransport(tcp.NewTransport())
	if err = subscriber.Dial("tcp://127.0.0.1:55555"); err != nil {
		die("can't dial on sub socket: %s", err.Error())
	}
	// Empty byte array effectively subscribes to everything
	err = subscriber.SetOption(mangos.OptionSubscribe, []byte(""))

	start := time.Now()

	//var msg []byte
	for i := 1; i <= *numbPtr; i++ {
		//conn.Publish("test", []byte(randSeq(320)))
		sock.Send([]byte(randSeq(320)))

		if _, err = subscriber.Recv(); err != nil {
			die("Cannot recv: %s", err.Error())
		}

		//log.Println(string(msg))
		/*
			runtime.Gosched()

			atomic.AddUint64(&ops, 1)
			if ops == uint64(*numbPtr) {
				elapsed := time.Since(start)
				log.Printf("Time took %s", elapsed)
			}
		*/
	}

	elapsed := time.Since(start)
	log.Printf("Time took %s", elapsed)

	defer sock.Close()

	fmt.Scanln()
}
