package main

import (
	"fmt"
	"github.com/ibmendoza/jwt"
	"time"
)

func main() {

	claims := make(map[string]interface{})
	claims["sub"] = "1234567890"
	claims["name"] = "John Doe"
	claims["admin"] = true
	claims["exp"] = jwt.ExpiresInSeconds(3)
	//claims["exp"] = jwt.ExpiresInMinutes(1)
	//claims["exp"] = jwt.ExpiresInHours(1)

	//fmt.Println(timeNow())

	naclKey, err := jwt.GenerateKey()
	fmt.Println(naclKey)

	token, err := jwt.Sign("HS384", claims, "secret", naclKey)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token, err) //encrypted claims

	token2, err := jwt.Sign("HS256", claims, "secret", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(token2, err) //non-encrypted claims

	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C

	//fmt.Println(timeNow()) //after 2sec

	fmt.Println(jwt.Verify(token, "secret", naclKey)) //true
	fmt.Println(jwt.Verify(token, "scret", ""))       //false

	fmt.Println(jwt.Verify(token2, "secret", ""))     //true
	fmt.Println(jwt.Verify(token2, "scret", ""))      //false
	fmt.Println(jwt.Verify(token2, "secret", "asdf")) //false
	fmt.Println(jwt.Verify(token2, "secret", ""))     //false

	timer1 = time.NewTimer(time.Second * 2) //after 4secs
	<-timer1.C

	fmt.Println(jwt.Verify(token, "secret", naclKey)) //expired
}
