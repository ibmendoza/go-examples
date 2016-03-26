package main

import (
	"fmt"
	"github.com/influxdata/influxdb/client/v2"
)

func main() {
	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	//NewQuery returns a query object database and precision strings can be
	//empty strings if they are not needed for the query.
	q := client.NewQuery("select * from cpu_usage group by region limit 2", "systemstats", "ns")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}
