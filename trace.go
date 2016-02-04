//https://github.com/VividCortex/trace

package main

import (
    "fmt"
    "time"

    "github.com/VividCortex/trace"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        trace.Trace(j)

        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 9; a++ {
        <-results
    }
}
