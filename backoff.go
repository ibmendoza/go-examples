//http://itjumpstart.wordpress.com

package main

import (
	"fmt"
	"net"
	"time"

	"github.com/jpillora/backoff"
)

func main() {
	b := &backoff.Backoff{
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Jitter: true,
	}

	for {
		_, err := net.Dial("tcp", "192.168.0.153:2015")
		if err != nil {
			d := b.Duration()
			fmt.Printf("%s, reconnecting in %s", err, d)
			time.Sleep(d)
			continue
		}
		//connected
		fmt.Println("connected...")
		break
	}

	//continue your work here
	fmt.Println("doing something...")
}
