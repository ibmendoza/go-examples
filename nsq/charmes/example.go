//http://blog.charmes.net/2014/10/play-with-nsq-in-go.html

package main

import (
        "fmt"

        nsq "github.com/bitly/go-nsq"
)

func nsqSubscribe(tcpAddr, topicName, channelName string, hdlr nsq.HandlerFunc) error {
        fmt.Printf("Subscribe on %s/%s\n", topicName, channelName)

        // Create the configuration object and set the maxInFlight
        cfg := nsq.NewConfig()
        cfg.MaxInFlight = 8

        // Create the consumer with the given topic and chanel names
        r, err := nsq.NewConsumer(topicName, channelName, cfg)
        if err != nil {
                return err
        }

        // Set the handler
        r.AddHandler(hdlr)

        // Connect to the NSQ daemon
        if err := r.ConnectToNSQD(tcpAddr); err != nil {
                return err
        }

        // Wait for the consumer to stop.
        <-r.StopChan
        return nil
}

func main() {
        nsqSubscribe("localhost:4150", "mytopic", "mychan1", func(msg *nsq.Message) error {
                fmt.Printf("%s\n", msg.Body)
                return nil
        })
}
