//For admin use only
//Purpose: Generate token and key for a particular client
//Inputs: client name, expiry

package main

import (
	_ "github.com/go-sql-driver/mysql"

	"bufio"
	"flag"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/ibmendoza/salt"
	"github.com/johnnylee/sqlxchain"
	"time"
)

const chars = "abcdefghjkmnpqrstuvwxyz" +
	"ABCDEFGHJKMNPQRSTUVWXYZ" + "23456789"

func RandomString(strlen int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func main() {

	//dbx, err := sqlxchain.New("mysql", "root:e593bF@tcp(127.0.0.1:3306)/dbadmin")

	dbx, err := sqlxchain.New("mysql", "root:e593bF@tcp(127.0.0.1:3306)/dbadmin")

	if err != nil {
		log.Fatal(err)
	}

	var company = flag.String("c", "", "company")
	var expiry = flag.Int("exp", 1, "expires in months")

	flag.Parse()

	if *company == "" {
		log.Fatal("Please provide company name")
	}

	client := *company
	exp := salt.ExpiresInMonths(time.Duration(*expiry))

	log.Println("exp")
	log.Println(exp)

	key, _ := salt.GenerateKey()
	log.Println("key")
	log.Println(key)

	claims := make(map[string]interface{})

	pswd := RandomString(7)
	claims["pswd"] = pswd

	claims["exp"] = exp

	token, _ := salt.Sign(claims, key)
	log.Println("token")
	log.Println(token)

	sql := "insert into vault(dbpswd, client, dtstart, dtend) " +
		"values(?, ?, now(), FROM_UNIXTIME(?))"

	i := int(exp)

	err = dbx.Context().Begin().
		Exec(sql, pswd, client, strconv.Itoa(i)).
		Commit().
		Err()

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("license_" + strconv.Itoa(int(exp)) + "_" +
		client + ".conf")

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString("[LICENSE]" + "\ncompany=" + client +
		"\nkey=" + key +
		"\ntoken=" + token)
	if err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
