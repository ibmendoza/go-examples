### Understanding Microservices

Microservices = message queue + message transport + Web services

#### Iris (message transport)

- Get Iris code
- Developer Mode
- Cluster Mode

#### Get Iris code

Download Iris version 0.3.2 [here](https://github.com/project-iris/iris/releases/download/v0.3.2/iris-v0.3.2-linux-amd64) 
and compile using the latest version of Golang

#### Developer Mode

```
iris -dev
```

#### Cluster mode

```
iris -net "clustername" -rsa /path/to/id_rsa
```

#### Publish/subscribe

```go
//publish
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
```

To run,

```
publish -m 100000
```

If there is no corresponding subscriber client in either local or remote machine, Iris daemon will output an error 
message:

```
failed to handle delivered publish (churn?)
```

In order to avoid that, run a corresponding subscriber client first before running a publisher client, otherwise
messages will be lost. Remember that Iris is an in-memory messaging transport (NSQ is the same but messages will be
stored to disk once a threshold is reached).

```go
//subscribe
package main

import (
	"flag"
	"fmt"
	"gopkg.in/project-iris/iris-go.v1"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

type topicEvent struct{}

var start = time.Now()
var ops uint64 = 0
var numbPtr = flag.Int("msg", 10000, "number of messages (default: 10000)")

func (t topicEvent) HandleEvent(event []byte) {

	//log.Println(string(event))

	atomic.AddUint64(&ops, 1)
	if ops == uint64(*numbPtr) {
		elapsed := time.Since(start)
		log.Printf("Time took %s", elapsed)
	}
}

func main() {

	flag.Parse()

	conn, err := iris.Connect(55555)
	if err != nil {
		log.Fatalf("failed to connect to the Iris relay: %v.", err)
	} else {
		log.Println("Connected to port 55555")
	}

	var topicHandler = new(topicEvent)

	conn.Subscribe("test", topicHandler, &iris.TopicLimits{
		EventThreads: runtime.NumCPU(),
		EventMemory:  64 * 1024 * 1024,
	})

	defer conn.Close()

	fmt.Scanln()
}
```

#### Request/reply

The Web is the most famous example of request/reply pattern. Here is an example in Iris.

```go
//request
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
```

```go
//reply
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
```

#### Broadcast

```go
//server broadcast
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
	log.Println(string(msg))
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
```

```go
//client receiving broadcast
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
```
