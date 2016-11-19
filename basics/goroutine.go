package main

import (
	"fmt"
	"log"

	"github.com/dchest/uniuri"
)

func f(s string, x, y int) {
	log.Println("string")
	log.Println(s)
	log.Println("x ", x)
	log.Println("y ", y)
	log.Println("x + y ", x+y)
	go g(s, 5, 6)
}

func g(s string, x, y int) {
	log.Println("string")
	log.Println(s)
	log.Println("x ", x)
	log.Println("y ", y)
	log.Println("x + y ", x+y)
}

func main() {

	var s string
	var i int
	for {
		i++
		s = uniuri.NewLen(4)
		go f(s, 1, 3)
		if i == 5 {
			break
		}
	}
	fmt.Scanln()
}
