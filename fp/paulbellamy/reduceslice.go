//https://github.com/paulbellamy/pipe

package main

import (
	"fmt"
	"github.com/paulbellamy/pipe"
)

func main() {
	// Declare a slice of some things
	chars := []string{"a", "b", "c"}

	// Function to apply
	concat := func(a, b string) string {
		return fmt.Sprintf("%s%s", a, b)
	}

	sum := pipe.ReduceSlice(concat, "", chars).(string)

	fmt.Println(sum)

	// Output:
	// abc
}
