//http://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
package main

import (
	"fmt"
	"math/rand"
)

var num = "0123456789"

// generates a random string of fixed size
func srand(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = num[rand.Intn(len(num))]
	}
	return string(buf)
}

func main() {
	fmt.Println(srand(7))
}
