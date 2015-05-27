////http://play.golang.org/p/HLERcuVnmC

package main

import "fmt"

type Future struct {
	result  interface{}
	error   bool
	c       chan interface{}
}

func (f *Future) Get() interface{} {
	result, ok := <-f.c
	if ok {
		f.result = result
		close(f.c)
	} else {
		f.error = true
	}
	return f.result
}

func (f *Future) isError() bool {
	return f.error
}

// return a completion function
func NewFuture() Future {
	c := make(chan interface{})
	go func() {
		c <- 42
	}()
	return Future{c: c}
}

func main() {

	future := NewFuture()
	fmt.Printf("%v\n", future)
	fmt.Printf("%v\n", future.Get())
	fmt.Printf("%v\n", future.Get())

}
