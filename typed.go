package main

import (
	"fmt"
	"github.com/karlseguin/typed"
)

func main() {
	json := `
{
  "log": true,
  "name": "leto",
  "percentiles": [75, 85, 95],
  "server": {
    "port": 9001,
    "host": "localhost"
  },
  "cache": {
    "users": {"ttl": 5}
  },
  "blocked": {
    "10.10.1.1": true
  }
}`
	typed, err := typed.JsonString(json)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(typed.Bool("log"))
	fmt.Println(typed.String("name"))
	fmt.Println(typed.Ints("percentiles"))
	fmt.Println(typed.FloatOr("load", 0.5))

	server := typed.Object("server")
	fmt.Println(server.Int("port"))
	fmt.Println(server.String("host"))

	fmt.Println(typed.Map("server"))

	fmt.Println(typed.StringObject("cache")["users"].Int("ttl"))
	fmt.Println(typed.StringBool("blocked")["10.10.1.1"])

	blocked := typed.Object("blocked")
	fmt.Println(blocked.Bool("10.10.1.1"))
}
