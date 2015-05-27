// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

//Adapted from https://gobyexample.com/worker-pools

package main

import (
	"fmt"
	"github.com/chuckpreslar/emission"
	"runtime"
)

func worker(id, j int) {
	fmt.Println("worker", id, "processing job", j)
	//time.Sleep(time.Second)
	//results <- j * 2
}

func printResults() {
	fmt.Println("results")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	emitter := emission.NewEmitter()

	emitter.On("worker", worker)

	for w := 1; w <= 3; w++ {
		for j := 1; j <= 9; j++ {
			emitter.Emit("worker", w, j)
		}
	}
}
