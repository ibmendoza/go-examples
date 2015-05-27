package main

import (
	"fmt"
	"github.com/ibmendoza/salt"
	"time"
)

func main() {

	claims := make(map[string]interface{})
	claims["sub"] = "1234567890"
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = salt.ExpiresInSeconds(1)
	//claims["exp"] = jwt.ExpiresInMinutes(1)
	//claims["exp"] = jwt.ExpiresInHours(1)

	key, _ := salt.GenerateKey()
	fmt.Println(key)

	token, _ := salt.Sign(claims, key)

	fmt.Println(token)

	fmt.Println(salt.Verify(token, key))

	timer1 := time.NewTimer(time.Second * 2) //after 2secs
	<-timer1.C

	fmt.Println(salt.Verify(token, key))
}
