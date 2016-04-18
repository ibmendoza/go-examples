package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("192.168.0.130:11211")

	it, err := mc.Get("foo")
	if err != nil {
		mc.Set(&memcache.Item{
			Key:        "foo",
			Value:      []byte("0"),
			Expiration: 0,
		})
	}
	it, err = mc.Get("foo")
	fmt.Println(string(it.Value))

	var cnt uint64
	cnt, err = mc.Increment("foo", 1)
	fmt.Println("cnt")
	fmt.Println(cnt)
}
