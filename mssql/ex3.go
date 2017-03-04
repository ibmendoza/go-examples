package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/johnnylee/sqlxchain"
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

	dbx, errx := sqlxchain.New("mssql", dsn)

	if errx != nil {
		log.Fatal(errx)
	}

	if err != nil {
		log.Println("From Open() attempt: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("From Ping() Attempt: " + err.Error())
		return
	}
	log.Println("ok")

	//demo sqlx

	row := OneField{}
	var slcRows rowsOneField

	sql := `
		select page as name
		from acl
		where idemployee = ?
	`
	rows, err := dbm.Queryx(sql, 1)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return
	}

	log.Println(slcRows)

	//demo sqlxchain

	sql = `
			select count(id) as cnt
			from acl
			where idemployee = ?
		`
	var cnt int

	err = dbx.Context().Begin().
		Get(&cnt, sql, 1).
		Commit().
		Err()

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(cnt)
}
