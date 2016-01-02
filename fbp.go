//https://groups.google.com/forum/#!msg/golang-nuts/vgj_d-MjUHA/T9sE64Yrcq0J
//https://github.com/egonelbre/exp/blob/master/fbp/example/main.go

package main

import (
	"fmt"
	"strings"

	"github.com/egonelbre/exp/fbp"
)

type Comm struct{ In, Out chan string }

func main() {
	comm := &Comm{}
	graph := fbp.New(comm)

	graph.Registry = fbp.Registry{
		"Split": NewSplit,
		"Lower": NewLower,
		"Upper": NewUpper,
	}

	graph.Setup(`
		: s Split
		: l Lower
		: u Upper
		$.In    -> s.In
		s.Left  -> l.In
		s.Right -> u.In
		l.Out -> $.Out
		u.Out -> $.Out
	`)

	graph.Start()

	for i := range []int{1, 2, 3, 4, 5} {
		comm.In <- fmt.Sprintf("Hello %v", i)
	}
	close(comm.In)

	for v := range comm.Out {
		fmt.Printf("%v\n", v)
	}
}

type Split struct{ In, Left, Right chan string }

func NewSplit() fbp.Node { return &Split{} }
func (node *Split) Run() error {
	defer close(node.Left)
	defer close(node.Right)
	for v := range node.In {
		m := len(v) / 2
		node.Left <- v[:m]
		node.Right <- v[m:]
	}
	return nil
}

type Lower struct{ In, Out chan string }

func NewLower() fbp.Node { return &Lower{} }
func (node *Lower) Run() error {
	defer close(node.Out)
	for v := range node.In {
		node.Out <- strings.ToLower(v)
	}
	return nil
}

type Upper struct{ In, Out chan string }

func NewUpper() fbp.Node { return &Upper{} }
func (node *Upper) Run() error {
	defer close(node.Out)
	for v := range node.In {
		node.Out <- strings.ToUpper(v)
	}
	return nil
}
