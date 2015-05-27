package main

import (
	"fmt"
	"github.com/otium/queue"
)

func main() {
	q := queue.NewQueue(func(val interface{}) {
		intVal, ok := val.(int)

		if ok {
			fmt.Println(intVal)
		}

	}, 20)
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	q.Wait()

	fmt.Println("Hello World!")
}
