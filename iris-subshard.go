package main

import (
	"flag"
	"fmt"
	//"gopkg.in/project-iris/iris-go.v1"
	"github.com/ibmendoza/iris-go"
	"log"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

type topicEvent struct{}
type doneEvent struct{}

var start = time.Now()
var numbPtr = flag.Int("msg", 10000, "number of messages (default: 10000)")
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()1234567890")
var ops uint64 = 0

func (t topicEvent) HandleEvent(event []byte) {
	b := make([]rune, 4092)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	atomic.AddUint64(&ops, 1)
}

func (t doneEvent) HandleEvent(event []byte) {
	elapsed := time.Since(start)
	log.Printf("Time took %s", elapsed)

	opsFinal := atomic.LoadUint64(&ops)

	if opsFinal == uint64(*numbPtr) {
		log.Println("Subscribers match number of publishers")
	} else {
		log.Println("Number of subscribers processed: ", opsFinal)
	}
}

func main() {

	topicLimits := &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	}

	flag.Parse()

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	var topicHandler = new(topicEvent)
	var doneHandler = new(doneEvent)

	conn.Subscribe("test", topicHandler, topicLimits)
	//conn.Subscribe("odd", topicHandler, topicLimits)
	conn.Subscribe("done", doneHandler, topicLimits)

	defer conn.Close()

	fmt.Scanln()
}
