### Unicast (Pipeline)

**Scalability Protocols**

Producer

```go
//pushpipeline.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ibmendoza/msgq"
)

//Example:
//pushpipeline -ip tcp://192.168.0.150:65000 -msg "Hello world"

var ip = flag.String("ip", "", "IP address")
var msg = flag.String("msg", "", "message to send")

func main() {
	flag.Parse()

	if *ip == "" {
		log.Fatal("ip address required")
	}

	if *msg == "" {
		log.Fatal("message required")
	}

	fmt.Println(*ip)

	pushpipeline, err := msgq.NewPushPipeline(*ip)
	if err != nil {
		log.Fatal("Error creating pushpipeline socket")
	}

	err = msgq.Dial(pushpipeline, *ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 1; i <= 1000; i++ {
		err = msgq.Send(pushpipeline, []byte(*msg))
		if err != nil {
			log.Println("error sending msg")
		}
	}
}
```

Consumer

```go
//pullpipeline.go

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ibmendoza/msgq"
)

//Example:
//pullpipeline -ip tcp://192.168.0.150:65000

var ip = flag.String("ip", "", "IP address")

func main() {
	flag.Parse()

	if *ip == "" {
		log.Fatal("ip address required")
	}

	pullpipeline, err := msgq.NewPullPipeline(*ip)
	if err != nil {
		log.Fatal("Error creating pullpipeline socket")
	}

	err = msgq.Listen(pullpipeline, *ip)
	if err != nil {
		log.Fatal("Error listening to pullpipeline socket")
	}

	for {
		msg, e := msgq.Receive(pullpipeline)
		if e != nil {
			log.Println("Error receiving message from pullpipeline socket")
		}

		fmt.Println(string(msg))
	}
}
```
