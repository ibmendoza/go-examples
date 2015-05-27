//http://play.golang.org/p/-3_kf8YfJD
package main

import "fmt"

// return a completion function
func foo() func() int {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	completion := func() int { return <-c }

	return completion
}

// just return the channel
func foo2() <-chan int {
	c := make(chan int)
	go func() {
		c <- 42
	}()

	return c
}

func main() {

	future := foo()
	fmt.Printf("%v\n", future)
	fmt.Printf("%v\n", future())

	resChan := foo2()
	fmt.Printf("%v\n", resChan)
	fmt.Printf("%v\n", <-resChan)

}
