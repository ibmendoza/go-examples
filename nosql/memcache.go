package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("192.168.0.153:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(it.Value))
	}
}
