//Generate line protocol file from movielens ratings.dat

package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	f, _ := os.Open("ratings.dat")

	scanner := bufio.NewScanner(f)

	i := 0

	var str string

	file, err := os.Create("lp.txt")
	if err != nil {
		log.Println(err)
	}
	file, err = os.OpenFile("lp.txt", os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()

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

				//tags in InfluxDB Go client v2 expects a string
				uid = result[0]
				mid = result[1]
				r = result[2]
				tm = result[3]

				str = str + "movielens,userid=" + uid + ",movieid=" + mid +
					" rating=" + r + " " + tm + "\n"

				//http://stackoverflow.com/questions/7151261/append-to-a-file-in-go
				_, err = file.WriteString(str)

				if err != nil {
					log.Println("Error writing ", str)
				}
			}
			cnt++
		}

		if i == 5000 {
			break
		}
	}

	log.Println("ok")
	log.Println(time.Since(start))
}
