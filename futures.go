//http://labs.strava.com/blog/futures-in-golang/

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Future(f func() (interface{}, error)) func() (interface{}, error) {
	var result interface{}
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		result, err = f()
	}()

	return func() (interface{}, error) {
		<-c
		return result, err
	}
}

func main() {
	url := "http://labs.strava.com"
	future := Future(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	})

	// do many other things

	b, err := future()
	body, _ := b.([]byte)

	log.Printf("response length: %d", len(body))
	log.Printf("request error: %v", err)
}
