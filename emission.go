package main

import (
	"fmt"
	"github.com/chuckpreslar/emission"
)

func fibo(n int) (result int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		result = x
		x, y = y, x+y
	}

	return result
}

func fibo1(arg int) {
	c := fibo(arg)
	fmt.Println("Fibo result", c)
}

func main() {
	emitter := emission.NewEmitter()
	emitter.On("fibo1", fibo1)
	emitter.Emit("fibo1", 21)
	emitter.Emit("fibo1", 14)
	emitter.Emit("fibo1", 7)

	emitter.Off("fibo1", fibo1) //zero receiver
	emitter.Emit("fibo1", 7)

	fmt.Scanln()
}
