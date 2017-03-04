package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
  //Windows authentication
	server := "localhost"             
	user := `computername\username`
	pswd := "password"             

	db, err := sql.Open("mssql", "server="+server+";user id="+user+";password="+pswd)
	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("From Ping() Attempt: " + err.Error())
		return
	}
	fmt.Println("ok")
}
