package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type rowsOneField []OneField

type OneField struct {
	Name string
}

func main() {
	server := "localhost"             
	user := `computername\username` 
	pswd := "password"             

	dsn := "server=" + server + ";user id=" + user + ";password=" + pswd + ";database=dbrma"

	db, err := sql.Open("mssql", dsn)

	dbm := sqlx.MustConnect("mssql", dsn)

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

	row := OneField{}
	var slcRows rowsOneField

	sql := `
		select page as name
		from acl
		where idemployee = ?
	`
	rows, err := dbm.Queryx(sql, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		err = rows.StructScan(&row)
		if err != nil {
			break
		} else {
			slcRows = append(slcRows, row)
		}
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(slcRows)
}
