//InfluxDB Ingestion Tool

/*

Notes: All bindings expect JSON schema

- flags:
	user: string (optional)
	password: string (optional)
	database name: string (required)
	connection: udp or http (optional) (default: udp)
	influx: string (required) - URL of InfluxDB (http://192.168.56.101:8086)

	nats: string (optional) - URL of NATS server (http only)
	subject: string (required if used in conjunction with NATS)

	nsqd: string (optional) - URL of nsqd (http only)
	nsqlookupd: string (optional) - URL of nsqlookupd (http only)
	topic: string (optional) - NSQ topic
	channel: string (optional) - NSQ channel

	nanomsg: string (optional) - Bind address
	nanotopic: string (optional) - Used in conjunction with nanomsg

	mangos: string (optional) - Ex: tcp://127.0.0.1:55555

	zmq3: string (optional) - tcp://localhost:5563
	zmq4: string (optional) - tcp://localhost:5563

	mqtt: https://www.cloudmqtt.com/docs-go.html

- JSON schema is an array of objects. Each object must have at least one field.
  If time is not specified, assumes current time

[   {
		"measurement" : "measurement name",
		"tags" : {
			"tagkey1" : "tagvalue1",
			"tagkeyN" : "tagvalueN"
		},
		"fields" : {
			"fieldkey1" : "fieldvalue1",
			"fieldkeyN" : "fieldvalueN"
		},
		"time" : "timeString or timeUNIXepoch"
	}, {
		"measurement2" : "measurement name 2",
		"tags" : {
			"tagkey2" : "tagvalue2",
			"tagkey2" : "tagvalueN"
		},
		"fields" : {
			"fieldkey2" : "fieldvalue2",
			"fieldkeyN" : "fieldvalueN"
		},
		"time" : "timeString or timeUNIXepoch"
	}
]


measurement - string (required)
tag keys and tag values are strings
fields can be integers, floats, strings and boolean (true, false)

- ingest from JSON file
- exposes /hook endpoint for HTTP POST from clients (assumes JSON schema)
*/
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
