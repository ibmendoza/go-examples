//courtesy: https://github.com/fatih/set

package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
)

func main() {
	// ... or with some initial values
	s := set.New("istanbul", "frankfurt", 30.123, "san francisco", 1234)
	// add items
	s.Add("istanbul")
	s.Add("istanbul") // nothing happens if you add duplicate item

	// add multiple items
	s.Add("ankara", "san francisco", 3.14)

	// remove item
	s.Remove("frankfurt")
	s.Remove("frankfurt") // nothing happes if you remove a nonexisting item

	// remove multiple items
	s.Remove("barcelona", 3.14, "ankara")

	// removes an arbitary item and return it
	item := s.Pop()

	fmt.Println(item)

	// create a new copy
	other := s.Copy()
	fmt.Println(other)

	// number of items in the set
	len := s.Size()
	fmt.Println(len)

	// return a list of items
	items := s.List()
	fmt.Println(items)

	// string representation of set
	fmt.Printf("set is %s\n", s.String())

	//fmt.Println([]byte(s.String()))
}
