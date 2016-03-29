//https://gist.github.com/andreagrandi/97263aaf7f9344d3ffe6

//Test it with: curl -X POST -d "{\"test\": \"that\"}" http://localhost:8080

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type test_struct struct {
	Test string
}

func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var t test_struct
	err := decoder.Decode(&t)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Test)
}

func main() {
	http.HandleFunc("/", parseGhPost)
	http.ListenAndServe(":8080", nil)
}
