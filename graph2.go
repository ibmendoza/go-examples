package main

import (
	"bytes"
	"encoding/gob"
	"github.com/akualab/graph"
	"log"
)

func main() {

	g := graph.New()

	// set key → value pairs
	g.Set("1", 123)
	g.Set("2", 678)
	g.Set("3", "abc")
	g.Set("4", "xyz")

	// connect vertexes/nodes
	g.Connect("1", "2", 5)
	g.Connect("1", "3", 1)
	g.Connect("2", "3", 9)
	g.Connect("4", "2", 3)

	// delete a node, and all connections to it
	g.Delete("1")

	// encode into buffer
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	err := enc.Encode(g)
	if err != nil {
		log.Println(err)
	}

	log.Println(enc)

	// now decode into new graph
	dec := gob.NewDecoder(buf)
	newG := graph.New()
	err = dec.Decode(newG)
	if err != nil {
		log.Println(err)
	}

	log.Println(newG)
}
