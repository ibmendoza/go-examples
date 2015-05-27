package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/akualab/graph"
	"github.com/akualab/graph/dot"
	"github.com/ibmendoza/remarshal"
	"log"
)

func main() {
	g := graph.New()

	// create nodes with values.
	g.Set("1", 123)
	g.Set("2", 678)
	g.Set("3", "abc")
	g.Set("4", "xyz")
	g.Set("xxx", "yyy")

	// make connections (ignoring errors for clarity.)
	g.Connect("1", "2", 5)
	g.Connect("1", "3", 1)
	g.Connect("2", "3", 9)
	g.Connect("4", "2", 3)
	g.Connect("4", "xxx", 1.11)

	g.WriteJSONGraph("graph.json")
	// to JSON
	jsonEncoded, _ := json.Marshal(g)
	log.Println(string(jsonEncoded))
	log.Println("")

	yamlEncoded, _ := remarshal.Convert(jsonEncoded, "JSON", "YAML")
	log.Println(yamlEncoded)

	// to GOB
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(g)
	log.Println(enc)

	// to DOT (use the dot sub-package.)
	d := dot.DOT(g, "some graph")
	log.Println(d)
}
