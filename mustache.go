package main

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/mustache"
)

func main() {
	//for simplicity, json must be an object of key-value pairs of
	//a) primitives (boolean, numbers, strings)
	//b) composite (array or objects)

	//In golang, a struct
	//No need to worry since hoisie/mustache accepts any data type (interface{})

	//By convention, avoid computed function. Instead, compute data beforehand
	//and return it as json accordingly

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	beatles := []byte(

		`{
  "beatles": [
    { "firstName": "John", "lastName": "Lennon" },
    { "firstName": "Paul", "lastName": "McCartney" },
    { "firstName": "George", "lastName": "Harrison" },
    { "firstName": "Ringo", "lastName": "Starr" }
  ]
}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	rendered := mustache.Render("{{num}}, strs {{strs}}", dat)
	fmt.Println(rendered)

	data := mustache.Render("hello {{c}}", map[string]string{"c": "world"})
	fmt.Println(data)

	//beatles
	if err := json.Unmarshal(beatles, &dat); err != nil {
		panic(err)
	}
	data = mustache.Render("{{#beatles}} * {{firstName}} "+"\n {{/beatles}}", dat)

	fmt.Println(data)

}
