//http://play.golang.org/p/bD8zcinOxX
package main

import (
	"log"
	"sync"
)

type Profile struct {
	Name    string
	Score   float32
	Friends []int
}

func name() string {
	return "Ryan"
}

func score() float32 {
	return 50
}

func friends() []int {
	return []int{1, 2, 3}
}

//Contains all the methods of the original waitGroup.
type waitGroup struct {
	sync.WaitGroup
}

//This is used to launch your functions.
func (w *waitGroup) Launch(fn func()) {
	w.Add(1)
	go func() {
		fn()
		w.Done()
	}()
}

func main() {
	log.Println("Getting things")

	var n string
	var s float32
	var f []int

	var wg waitGroup

	log.Println("Getting name")

	//We are passing an anonymous function to Launch(), and than using closure to
	//assign the return value to n in the main funciton.
	wg.Launch(func() {
		n = name()
	})

	log.Println("Getting score")
	wg.Launch(func() {
		s = score()
	})

	log.Println("Getting friends")
	wg.Launch(func() {
		f = friends()
	})

	wg.Wait()
	p := Profile{Name: n, Score: s, Friends: f}

	log.Println(p)
}
