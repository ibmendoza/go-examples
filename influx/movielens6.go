//Generate line protocol input from movielens ratings.dat

package main

import (
	"bufio"
	"bytes"
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
var timeout = time.Duration(10 * time.Second)

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func postToDB(str string) {
	var err error
	numFlushed++
	log.Println(numFlushed)

	r, err = http.NewRequest("POST", "http://192.168.56.101:8086/write?db=ML1M",
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

		counter++
		isFlushed = math.Mod(counter, batchsize) == 0

		if isFlushed {

			postToDB(str)

			runtime.Gosched()

			str = ""
		}
	}

	//last write
	postToDB(str)

	log.Println("ok")
	log.Println(time.Since(start))
}
