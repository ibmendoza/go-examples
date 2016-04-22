//http://www.capykoa.com/articles/11

package main

import (
    "flag"
    "log"
    "time"

    "github.com/dleung/gotail"
)

var fname string

func main() {
    flag.StringVar(&fname, "file", "", "File to tail")
    flag.Parse()

    tail, err := gotail.NewTail(fname, gotail.Config{Timeout: 10})
    if err != nil {
        log.Fatalln(err)
    }

    var count int
    go func() {
        for {
            startTime := time.Now()
            countNow := count
            time.Sleep(5 * time.Second)
            duration := time.Since(startTime).Seconds()
            newCount := count

            log.Printf("%d processed at %0.2f rows/sec\n", count, float64(newCount-countNow)/duration)
        }
    }()
    for _ = range tail.Lines {
        count++
    }
}
