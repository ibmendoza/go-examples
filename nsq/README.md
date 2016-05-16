For NSQ primer, read

Traun Leyden's post (http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go)

Guillaume Charmes (http://blog.charmes.net/2014/10/first-look-at-nsq.html)

IrisMQ Guide (http://github.com/irismq)

#### Decoupling nsqd, nsqlookupd, nsq producer, nsq consumer

On VirtualBox, run three Turnkey Linux VMs (using host-only adapter)

```bash
192.168.56.101 - nsqlookupd
192.168.56.102 - nsqd and nsq producer client
192.168.56.103 - nsq consumer client
```

**192.168.56.101**

```
nsqlookupd
```

**192.168.56.102**

```bash
--lookupd-tcp-address = IP address of nsqlookupd
--broadcast-address = IP address of machine where nsqd is running
```

```
nsqd --lookupd-tcp-address=192.168.56.101:4160 --broadcast-address=192.168.56.102
```

On the same node as nsqd is the nsq producer client

```go
//NSQ Producer Client

package main

import (
	"flag"
	"fmt"
	"github.com/ibmendoza/go-lib"
	"github.com/nsqio/go-nsq"
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
	/*
	   Below are the default port settings
	   nsqd listens at port 4150 (for TCP clients), 4151 (for HTTP clients)

	   nsqlookupd listens at port 4160 (for TCP clients), 4161 (for HTTP clients)

	   nsqadmin listens at port 4171 (for HTTP clients) or
	     to be specified (for go-nsq clients) q.ConnectToNSQLookupd("127.0.0.1:4161")

	   http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go/
	   $ nsqlookupd &
	   $ nsqd --lookupd-tcp-address=127.0.0.1:4160 &
	   $ nsqadmin --lookupd-http-address=127.0.0.1:4161 &
	*/

	config := nsq.NewConfig()

	ipaddr, _ := lib.GetIPAddress()

	w, err := nsq.NewProducer(ipaddr+":4150", config)

	if err != nil {
		log.Fatal("Could not connect")
	}

	flag.Parse()

	start := time.Now()

	for i := 1; i <= *numbPtr; i++ {
		w.Publish("test", []byte(randSeq(320)))
	}

	elapsed := time.Since(start)
	log.Printf("Time took %s", elapsed)

	w.Stop()

	fmt.Scanln()
}
```

**NSQ Consumer Client**

Test nsq consumer client at **192.168.56.103**

```go
//NSQ Consumer Client
package main

import (
	"flag"
	"fmt"
	"github.com/itmarketplace/go-queue"
	"github.com/nsqio/go-nsq"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

var start = time.Now()
var ops uint64 = 0
var numbPtr = flag.Int("msg", 10000, "number of messages (default: 10000)")
var lkp = flag.String("lkp", "", "IP address of nsqlookupd")

func main() {

	/*
	   Below are the default port settings
	   nsqd listens at port 4150 (for TCP clients), 4151 (for HTTP clients)

	   nsqlookupd listens at port 4160 (for TCP clients), 4161 (for HTTP clients)

	   nsqadmin listens at port 4171 (for HTTP clients) or
	     to be specified (for go-nsq clients) q.ConnectToNSQLookupd("127.0.0.1:4161")

	   http://tleyden.github.io/blog/2014/11/12/an-example-of-using-nsq-from-go/
	   $ nsqlookupd &
	   $ nsqd --lookupd-tcp-address=127.0.0.1:4160 &
	   $ nsqadmin --lookupd-http-address=127.0.0.1:4161 &
	*/

	flag.Parse()

	/*
		wg := &sync.WaitGroup{}

		config := nsq.NewConfig()
		q, _ := nsq.NewConsumer("test", "ch", config)

		q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
			wg.Add(1)

			//log.Printf("Got a message: %v", string(message.Body))

			atomic.AddUint64(&ops, 1)
			if ops == uint64(*numbPtr) {
				elapsed := time.Since(start)
				log.Printf("Time took %s", elapsed)
			}

			wg.Done()
			return nil
		}))

		//err := q.ConnectToNSQD("127.0.0.1:4150") - hardcoded

		//err := q.ConnectToNSQLookupd("127.0.0.1:4161") //better

		err := q.ConnectToNSQLookupd(*ipnsqlookupd + ":4161") //much better
		if err != nil {
			log.Panic("Could not connect")
		}
		wg.Wait()
	*/

	c := queue.NewConsumer("test", "ch")

	c.Set("nsqlookupd", *lkp+":4161")
	c.Set("concurrency", runtime.GOMAXPROCS(runtime.NumCPU()))
	c.Set("max_attempts", 10)
	c.Set("max_in_flight", 150)
	c.Set("default_requeue_delay", "15s")

	c.Start(nsq.HandlerFunc(func(msg *nsq.Message) error {
		atomic.AddUint64(&ops, 1)
		if ops == uint64(*numbPtr) {
			elapsed := time.Since(start)
			log.Printf("Time took %s", elapsed)
		}

		//log.Println(string(msg.Body))
		return nil
	}))
	fmt.Scanln()
}
```
