package main

import (
	"github.com/antonholmquist/jason"
	"log"
	"time"
)

func saveToInfluxDB(measurement string, mapTags map[string]string,
	mapFields map[string]interface{}, tm time.Time) error {

	return nil
}

func parse(json []byte) {
	var v *jason.Value
	var err error

	v, err = jason.NewValueFromBytes(json)

	if err != nil {
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

		fields, err = obj.GetObject("fields")
		if err != nil {
			log.Println("Error: Parsing fields. Must have at least one field")
		} else {
			//log.Println("FIELDS")
			//log.Println(fields)

			mapFields = make(map[string]interface{})
			for key, value := range fields.Map() {
				//log.Println(key)
				//log.Println(value.Interface())
				mapFields[key] = value.Interface()
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
			saveToInfluxDB(measurement, mapTags, mapFields, tm)
		}
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
		"fieldkey1": "fieldvalue1",
		"fieldkeyN": "fieldvalueN"
	},
	"time": "timeString or timeUNIXepoch"
}, {
	"measurement": "measurement name 2",
	"fields": {
		"idle": 53.3,
		"s1": "chasney",
		"s2": "chelsea",
		"friends": true,
		"duration": 135
	}
}]
`
	parse([]byte(exampleJSON))
}
