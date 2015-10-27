//Isagani Mendoza (http://itjumpstart.wordpress.com)

//https://blog.golang.org/pipelines
//Pipeline example using event emitter instead (publish/subscribe)

package main

import (
	"fmt"
)

import (
	"github.com/chuckpreslar/emission"
)

/*
//first stage
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

//second stage
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//final stage
func main() {
    // Set up the pipeline.
    c := gen(2, 3)
    out := sq(c)

    // Consume the output.
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}

*/

func main() {
	emitter := emission.NewEmitter()

	//subscriber or worker function
	square := func(nums ...int) {
		var sq int
		for _, n := range nums {
			sq = n * n
			fmt.Println(sq)
		}
	}

	//bind interface to subscriber/worker function
	emitter.On("square", square)

	//signal event with payload data
	emitter.Emit("square", 2, 3)
}
