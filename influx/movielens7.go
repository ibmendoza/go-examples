//Generate line protocol file from movielens ratings.dat

package main

import (
	"bufio"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var numFlushed = 0

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		log.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "ML1M",
		Precision: "s",
	})

	start := time.Now()

	f, _ := os.Open("ratings.dat")

	scanner := bufio.NewScanner(f)

	i := 0

	counter := 0.0
	isFlushed := false

	batchsize := 5000.0

	for scanner.Scan() {
		i++

		line := scanner.Text()

		result := strings.Split(line, "::")

		var uid, mid, r, tm string
		cnt := 0

		for _ = range result {

			//Example: 1::1193::5::978300760
			//userid::movieid::rating::time

			if cnt == 3 {

				uid = result[0]
				mid = result[1]
				r = result[2]
				tm = result[3]

				/*
					str = str + "movielens,userid=" + uid + ",movieid=" + mid +
						" rating=" + r + " " + tm + "\n"
				*/

				tags := map[string]string{"userid": uid, "movieid": mid}

				fields := map[string]interface{}{"rating": r}

				var tmField time.Time
				var timeInSeconds int64
				timeInSeconds, err = strconv.ParseInt(tm, 10, 64)

				if err != nil {
					log.Println("Invalid time", tm)
				} else {
					tmField = time.Unix(timeInSeconds, 0)
				}

				pt, err := client.NewPoint(
					"movielens",
					tags,
					fields,
					tmField,
				)
				if err != nil {
					println("Error:", err.Error())
					continue
				}
				bp.AddPoint(pt)

			}
			cnt++
		}

		counter++
		isFlushed = math.Mod(counter, batchsize) == 0

		if isFlushed {
			log.Println(counter)

			err = c.Write(bp)
			if err != nil {
				log.Println(err)
			}

			runtime.Gosched()
		}
	}

	//last write
	//postToDB(str)
	err = c.Write(bp)
	if err != nil {
		log.Println("Error: ", err.Error())
	}

	log.Println("ok")
	log.Println(time.Since(start))
}
