package main

import (
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/ibmendoza/salt"
	"github.com/johnnylee/sqlxchain"
	//"time"
)

func main() {

	dbx, err := sqlxchain.New("mysql", "user:pswd@tcp(127.0.0.1:3306)/database")

	if err != nil {
		log.Fatal(err)
	}

	//INPUT
	client := "Diovel Pharmacy"

	//INPUT
	exp := salt.ExpiresInMonths(2)
	//exp := salt.ExpiresInDays(30)
	//exp := salt.ExpiresInHours(48)
	//exp := 1454000043

	log.Println("exp")
	log.Println(exp)

	claims := make(map[string]interface{})
	claims["client"] = client
	//claims["exp"] = salt.ExpiresInSeconds(5)
	//claims["exp"] = salt.ExpiresInMinutes(1)
	claims["exp"] = exp

	key, _ := salt.GenerateKey()
	//log.Println(len(key))
	log.Println("key")
	log.Println(key)

	token, _ := salt.Sign(claims, key)
	//log.Println(len(token))
	log.Println("token")
	log.Println(token)

	//log.Println(salt.Verify(token, key))

	//timer1 := time.NewTimer(time.Second * 2)
	//<-timer1.C

	//log.Println(salt.Verify(token, key)) //expired after 2secs

	sql := "insert into vault(token, thekey, client, dtstart, dtend, exp) " +
		"values(?, ?, ?, now(), FROM_UNIXTIME(?), ?)"

	i := int(exp)

	err = dbx.Context().Begin().
		Exec(sql, token, key, client, strconv.Itoa(i), exp).
		Commit().
		Err()

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("license.conf")
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString("[LICENSE]" +
		"\ncompany=" + client +
		"\ntoken=" + token +
		"\nthekey=" + key)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()

	/*
		_, e := salt.Verify(token, key)
		if e != nil {
			log.Println(e)
		}
	*/
}
