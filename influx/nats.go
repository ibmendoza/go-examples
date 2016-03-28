package main

import (
	"flag"
	"github.com/antonholmquist/jason"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/nats-io/nats"
	"log"
	"net/http"
	"time"
)

func main() {
	//gnatsd -m 8222
	flag.Parse()

	start := time.Now()

	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal("NATS server not running")
	}
	defer natsConnection.Close()
	log.Println("Connected to NATS server: " + nats.DefaultURL)

	// Make InfluxDB client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://192.168.56.101:8086",
	})
	if err != nil {
		log.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "telegraf",
		Precision: "s",
	})
	//make influxdb client

	var response *http.Response
	//scrape NATS statistics

	response, err = http.Get("http://localhost:8222/varz")
	v, _ := jason.NewObjectFromReader(response.Body)

	now, _ := v.GetString("now")
	in_msgs, _ := v.GetNumber("in_msgs")
	out_msgs, _ := v.GetNumber("out_msgs")
	in_bytes, _ := v.GetNumber("in_bytes")
	out_bytes, _ := v.GetNumber("out_bytes")

	t1, _ := time.Parse(time.RFC3339, now)
	secs := t1.Unix()
	t := time.Unix(secs, 0)

	tags := map[string]string{
		"in_msgs":  string(in_msgs),
		"out_msgs": string(out_msgs),
	}

	fields := map[string]interface{}{
		"in_bytes":  in_bytes,
		"out_bytes": out_bytes,
		"now":       t1}

	pt, err := client.NewPoint(
		"natsMeasurement",
		tags,
		fields,
		t)
	if err != nil {
		log.Println("Error:", err.Error())
	}
	bp.AddPoint(pt)

	err = c.Write(bp)
	if err != nil {
		log.Println("Error: ", err.Error())
	}

	elapsed := time.Since(start)
	log.Printf("Time took %s", elapsed)
}
