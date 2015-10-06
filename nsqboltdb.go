//https://github.com/gophergala/goloso/blob/master/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitly/go-nsq"
	"github.com/bitly/nsq/util"
	"github.com/boltdb/bolt"
)

var (
	showHelp    = flag.Bool("help", false, "print help")
	showVersion = flag.Bool("version", false, "print version")
	topic       = flag.String("topic", "", "NSQ topic")
	channel     = flag.String("channel", "", "NSQ channel")
)

// bootstrap event struct
/*
type NSQMessage struct {
	Event      []string `json:event`      // event type
	Uuid       []string `json:uuid`       // event uuid
	InstanceId []string `json:instanceid` // instance id
	IpAddress  []string `json:ipaddress`  // ipaddess
	Os         []string `json:os`         // operaring system
}
*/

func main() {

	flag.Parse()

	if *showHelp {
		fmt.Println(`
Usage:
    goloso --help
    goloso --version
    goloso --channel "orc.sys.events" --topic "ec2"
`)
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("Goloso v%s\n", util.BINARY_VERSION)
		os.Exit(0)
	}

	fmt.Println("Goloso.. starting")

	if *channel == "" {
		log.Fatalln("Err: missing channel")
	}

	if *topic == "" {
		log.Fatalln("Err: missing topic. \"--topic is required\"")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var (
		consumer *nsq.Consumer
		err      error
	)

	// connect to database
	fmt.Print("Connecting to bolt...")

	// setup bolt db connection
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("done")

	// create buquete
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Goloso"))
		if err != nil {
			return fmt.Errorf("Create bucket: %s", err)
			fmt.Println("Goloso exists")
		}

		fmt.Println("Goloso bucket created")

		return nil
	})

	lookup := "localhost:4161"

	// setup nsq config
	conf := nsq.NewConfig()
	conf.MaxInFlight = 1000

	// setup nsq consumer
	consumer, err = nsq.NewConsumer(*topic, *channel, conf)
	if err != nil {
		log.Fatalln("Err: can't consume", err)
	}

	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Message: %s", message)

		// json_decode && store in KV

		// db.Update(func(tx *bolt.Tx) error {
		// 	b := tx.Bucket([]byte("Goloso"))
		// 	err := b.Put([]byte("answer"), []byte("42"))
		// 	return err
		// })

		return nil
	}))

	err = consumer.ConnectToNSQLookupd(lookup)
	if err != nil {
		log.Fatalln("Err: can't connect to lookupd", err)
	}

	for {
		select {
		case <-consumer.StopChan:
			return
		case <-sigChan:
			consumer.Stop()
		}
	}
}
