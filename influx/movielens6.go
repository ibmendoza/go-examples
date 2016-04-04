//Generate line protocol file from movielens ratings.dat

package main

import (
	"bufio"
	"bytes"
	//"github.com/jpillora/backoff"
	"log"
	"math"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var client = &http.Client{Timeout: timeout}
var r *http.Request
var numFlushed = 0
var resp *http.Response
var timeout = time.Duration(5 * time.Second)

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func postToDB(str string) {
	var err error
	numFlushed++
	log.Println(numFlushed)

	r, err = http.NewRequest("POST", "http://192.168.56.102:8086/write?db=ML1M",
		bytes.NewBufferString(str))

	if err != nil {
		log.Println(err)
	}

	resp, err = client.Do(r)

	if err != nil {
		log.Println(err)
	}

	runtime.Gosched()

	log.Println(resp.Status)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	start := time.Now()

	f, _ := os.Open("ratings.dat")

	scanner := bufio.NewScanner(f)

	i := 0

	var str string

	counter := 0.0
	isFlushed := false

	//batchsize of 5000 points saturated InfluxDB at 102nd insert
	/*
		b := &backoff.Backoff{
			Min:    500 * time.Millisecond,
			Max:    2 * time.Second,
			Factor: 2,
			Jitter: false,
		}
	*/

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

				str = str + "movielens,userid=" + uid + ",movieid=" + mid +
					" rating=" + r + " " + tm + "\n"

			}
			cnt++
		}
		/*
			if i == 5000 {
				break
			}
		*/

		counter++
		isFlushed = math.Mod(counter, batchsize) == 0

		if isFlushed {

			//d := b.Duration()

			postToDB(str)

			runtime.Gosched()

			//time.Sleep(d)

			//b.Reset()

			//reset string
			str = ""
		}
	}

	//last write
	postToDB(str)

	log.Println("ok")
	log.Println(time.Since(start))
}
