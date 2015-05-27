//Adapted from https://gist.github.com/ryanfitz/4191392
package main

import (
	"fmt"
	"time"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)

	}
}

func hello2(t time.Time) {
	fmt.Printf("%v: every 2 seconds!\n", t)
}

func hello1(t time.Time) {
	fmt.Printf("%v: every 1 second!\n", t)
}

func main() {
	doEvery(2*time.Second, hello2)

}
