//Gani Mendoza
//http://itjumpstart.wordpress.com

package main

import (
	"fmt"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("192.168.0.153:11211", "192.168.0.153:11212")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value"),
		Expiration: int32(time.Now().Add(time.Second * 5).Unix())})

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(it.Value))
	}

	time.Sleep(time.Second * 6)
	it, err = mc.Get("foo")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(it.Value))
	}
}
