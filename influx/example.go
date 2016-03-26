package main

import (
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

const (
	MyDB     = "square_holes"
	username = "bubba"
	password = "bumblebeetuna"
)

func main() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Println(err)
	}
	/*
	   type BatchPoints interface {
	       // AddPoint adds the given point to the Batch of points
	       AddPoint(p *Point)
	       // AddPoints adds the given points to the Batch of points
	       AddPoints(ps []*Point)
	       // Points lists the points in the Batch
	       Points() []*Point

	       // Precision returns the currently set precision of this Batch
	       Precision() string
	       // SetPrecision sets the precision of this batch.
	       SetPrecision(s string) error

	       // Database returns the currently set database of this Batch
	       Database() string
	       // SetDatabase sets the database of this Batch
	       SetDatabase(s string)

	       // WriteConsistency returns the currently set write consistency of this Batch
	       WriteConsistency() string
	       // SetWriteConsistency sets the write consistency of this Batch
	       SetWriteConsistency(s string)

	       // RetentionPolicy returns the currently set retention policy of this Batch
	       RetentionPolicy() string
	       // SetRetentionPolicy sets the retention policy of this Batch
	       SetRetentionPolicy(s string)
	   }
	*/

	/*
	   NewBatchPoints returns a BatchPoints interface based on the given config.
	   BatchPoints is an interface into a batched grouping of points to write into InfluxDB
	   together. BatchPoints is NOT thread-safe, you must create a separate batch for each
	   goroutine.
	*/
	bp, errbp := client.NewBatchPoints(client.BatchPointsConfig{

		// Database is the database to write points to
		Database: MyDB,

		// Precision is the write precision of the points, defaults to "ns"
		Precision: "s",
	})

	if errbp != nil {
		log.Println(errbp)
	}

	// Create a point and add to batch
	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}

	/*
		NewPoint returns a point with the given timestamp. If a timestamp is not
		given, then data is sent to the database without a timestamp, in which case
		the server will assign local time upon reception. NOTE: it is recommended to
		send data with a timestamp.
	*/
	pt, errnp := client.NewPoint("cpu_usage", tags, fields, time.Now())

	if errnp != nil {
		log.Println(errnp)
	}

	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)
}
