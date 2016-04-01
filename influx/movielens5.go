//Generate line protocol file from movielens ratings.dat

package main

import (
	"bufio"
	"bytes"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func main() {

	start := time.Now()

	f, _ := os.Open("ratings2k.dat")

	scanner := bufio.NewScanner(f)

	i := 0

	var str string

	counter := 0.0
	isFlushed := false
	batchsize := 500.0

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
			/*
				instead of write to file, post to HTTP

				filename := "lp" + FloatToString(counter) + ".txt"

				file, err := os.Create(filename)
				if err != nil {
					log.Println(err)
				}
				file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0666)

				_, err = file.WriteString(str)

				if err != nil {
					log.Println("Error WriteString ", err)
				}

				err = file.Sync()
				if err != nil {
					log.Println("Error Sync() ", err)
				}

				file.Close()
			*/

			client := &http.Client{}
			r, _ := http.NewRequest("POST", "http://192.168.56.101:8086/write?db=ml5k",
				bytes.NewBufferString(str))

			resp, _ := client.Do(r)
			log.Println(resp.Status)

			//reset string
			str = ""
		}
	}

	log.Println("ok")
	log.Println(time.Since(start))
}
