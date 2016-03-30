package main

import (
	"flag"
	"fmt"
	"github.com/antonholmquist/jason"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"net"
	"net/http"
	"time"
)

var flushNow = false
var c client.Client
var bp client.BatchPoints

var database = flag.String("database", "", "database")

//also can be 192.168.56.101:8086 etc
var url = flag.String("url", "localhost:8086", "URL of InfluxDB")

//func OpenInfluxDB(url string) {
func openInfluxDB(db string) (err error) {

	var url = "http://" + *url

	c, err = client.NewHTTPClient(client.HTTPConfig{Addr: url})

	_, err = queryDB(c, fmt.Sprintf("CREATE DATABASE %s", db), db)
	if err != nil {
		log.Println(err)
	}

	return err
}

//func newBatchPoint(database, precision string) (err error) {
func newBatchPoint() (err error) {

	bp, err = client.NewBatchPoints(client.BatchPointsConfig{
		Database:  *database,
		Precision: "s"})

	return err
}

func addPoint(measurement string, tags map[string]string,
	fields map[string]interface{}, tm time.Time) {

	pt, err := client.NewPoint(measurement, tags, fields, tm)
	if err != nil {
		log.Println("Error in addPoint:", err.Error())
	}
	bp.AddPoint(pt)
}

func saveToInfluxDB(proceed bool) {
	if proceed {
		err := c.Write(bp)
		if err != nil {
			log.Println("Error in saveToInfluxDB: ", err.Error())
		}
	}
}

func queryDB(clnt client.Client, cmd, db string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: db,
	}
	if response, err := clnt.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func parse(v *jason.Value) {
	var batchsize = 1000
	//var v *jason.Value
	var err error

	if err != nil {
		log.Println("Error parsing JSON")
		log.Fatal(err)
	}
	//log.Println(v)

	//slice of obj measurement1 and measurement2
	var slcObj []*jason.Value
	slcObj, err = v.Array()

	if err != nil {
		log.Fatal(err)
	}
	//log.Println(slcObj)

	counter := 0

	//loop through array/slice
	for _, value := range slcObj {
		var measurement string
		var mapTags map[string]string
		var mapFields map[string]interface{}
		var tm time.Time
		var obj, fields, tags *jason.Object
		var tagValue string

		err = nil //initialize

		//log.Println(key)
		//log.Println(value)

		obj, err = value.Object()
		if err != nil {
			log.Println("Error: Expecting object in array of measurement")
		}
		//log.Println("OBJECT")
		//log.Println(obj)

		measurement, err = obj.GetString("measurement")
		if err != nil {
			log.Println("Error: Expecting measurement key name in JSON")
		}

		if measurement == "" {
			log.Println("Error: Blank value of measurement name")
		}
		log.Println(measurement)

		//time
		var timeInSeconds int64
		timeInSeconds, err = obj.Int64()

		//if valid UNIX epoch time
		if err == nil {
			tm = time.Unix(timeInSeconds, 0)

		} else {
			var strTime string
			strTime, err = obj.GetString("time")

			if err != nil {
				log.Println("Time not specified. Assumes time.Now")
				tm = time.Now()
			} else {
				var t1 time.Time
				t1, err = time.Parse(time.RFC3339, strTime)
				if err != nil {
					log.Println("Invalid time", strTime)
				} else {
					secs := t1.Unix()
					tm = time.Unix(secs, 0)
				}

				//log.Println("TIME")
				//log.Println(strTime)
			}
		}

		fields, err = obj.GetObject("fields")
		if err != nil {
			log.Println("Error: Parsing fields. Must have at least one field")
		} else {
			//log.Println("FIELDS")
			//log.Println(fields)

			mapFields = make(map[string]interface{})

			var b bool
			var i int64
			var f float64
			var s string
			var em error

			for key, value := range fields.Map() {
				//log.Println(key)
				//log.Println(value.Interface())
				b, em = value.Boolean()
				if em == nil {
					mapFields[key] = b
				}

				i, em = value.Int64()
				if em == nil {
					mapFields[key] = i
				}

				f, em = value.Float64()
				if em == nil {
					mapFields[key] = f
				}

				s, em = value.String()
				if em == nil {
					mapFields[key] = s
				}

				//mapFields[key] = value.Interface()
			}
			log.Println(mapFields)
		}

		tags, err = obj.GetObject("tags")
		if err != nil {
			//tags are optional so it's ok to be nil
			err = nil
		} else {
			//log.Println("TAGS")
			//log.Println(tags)

			mapTags = make(map[string]string)
			for key, value := range tags.Map() {
				//log.Println(key)
				//log.Println(value.Interface())
				tagValue, err = value.String()
				if err != nil {
					log.Println("Error parsing tag value")
				} else {
					mapTags[key] = tagValue
				}
			}
			log.Println(mapTags)
		}

		if err != nil {
			//don't process this instance if there's any error
			log.Println("Error parsing this instance", value, err)
		} else {
			counter++

			addPoint(measurement, mapTags, mapFields, tm)

			saveToInfluxDB(counter == batchsize)
		}
	} //end for loop

	saveToInfluxDB(true)
}

func parseJSON(rw http.ResponseWriter, request *http.Request) {
	v, err := jason.NewValueFromReader(request.Body)
	if err != nil {
		log.Println(err)
	} else {
		parse(v)
	}
}

func main() {

	exampleJSON :=
		`
	   [{
	   	"measurement": "measurement name",
	   	"tags": {
	   		"tagkey1": "tagvalue1",
	   		"tagkeyN": "tagvalueN"
	   	},
	   	"fields": {
	   		"fieldkey1": "5826",
	   		"fieldkeyN": "159458.98"
	   	},
	   	"time": "timeString or timeUNIXepoch"
	   }, {
	   	"measurement": "measurement name 2",
	   	"fields": {
	   		"idle": 123.45,
	   		"s1": "chasney",
	   		"s2": "chelsea",
	   		"friends": true,
	   		"duration": 589622458,
	   		"string1": "2589363",
	   		"string2": "47895.98"			
	   	}
	   }]
	   `

	flag.Parse()

	if *url == "" {
		log.Fatal("Please specify URL of InfluxDB. Example: 192.168.56.101:8086")
	}

	//servAddr := "192.168.56.101:8086"
	tcpAddr, err := net.ResolveTCPAddr("tcp", *url)
	if err != nil {
		log.Fatal("ResolveTCPAddr failed: ", err)
	}

	_, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed: ", err)
	}

	if *database == "" {
		log.Fatal("Please specify name of database")
	}

	err = openInfluxDB(*database)
	if err != nil {
		log.Fatal("Error creating InfluxDB Client: ", err)
	}

	err = newBatchPoint()
	if err != nil {
		log.Fatal("Error creating batch point: ", err)
	}

	defer c.Close()

	var v *jason.Value
	v, err = jason.NewValueFromBytes([]byte(exampleJSON))
	if err != nil {
		log.Println(err)
	} else {
		parse(v)
	}

	http.HandleFunc("/", parseJSON)
	http.ListenAndServe(":8080", nil)
}
