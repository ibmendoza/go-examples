// http://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	var aww string
	wg.Add(1)
	go func() {
		defer wg.Done()
		aww = fetch("http://www.reddit.com/r/aww.json")
	}()

	var funny string
	wg.Add(1)
	go func() {
		defer wg.Done()
		funny = fetch("http://www.reddit.com/r/funny.json")
	}()

	wg.Wait()

	fmt.Println("aww:", aww)
	fmt.Println("funny:", funny)
}

func fetch(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
