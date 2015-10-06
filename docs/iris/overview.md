### Understanding Microservices

#### Iris

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

