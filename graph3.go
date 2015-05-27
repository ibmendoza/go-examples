package main

import (
	"fmt"
	"github.com/akualab/graph"
)

func main() {
	g := graph.New()

	// create nodes with values.
	g.Set("node1", 123)
	g.Set("node2", 678)
	g.Set("node3", "abc")
	g.Set("node4", "xyz")
	g.Set("nodexxx", "yyy")

	// make connections (ignoring errors for clarity.)
	g.Connect("node1", "node2", 1)
	g.Connect("node1", "node3", 5)
	//g.Connect("node1", "node3", 4)
	g.Connect("node2", "node3", 9)
	g.Connect("node2", "node4", 3)

	g.WriteJSONGraph("graph.json")

	node1, _ := g.Get("node2")
	//fmt.Println(node1.Value())

	slc := g.Predecessors(node1)
	for k, v := range slc {
		fmt.Println("Predecessors")
		fmt.Println("Key= ", k)
		fmt.Println("Value= ", v)
	}

	for i := 0; i < len(slc); i++ {
		fmt.Println("Predecessors")
		fmt.Println("Key= ", slc[i].Key())
		fmt.Println("Value= ", slc[i].Value())
	}

	slc2 := node1.Successors()
	for k, _ := range slc2 {
		fmt.Println("Successors")
		fmt.Println("Key= ", k.Key())
		fmt.Println("Value= ", k.Value())
		//fmt.Println("Value= ", v)
	}
}
