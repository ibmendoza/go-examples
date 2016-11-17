package main

import (
	//"fmt"
	"log"
	"strconv"

	"github.com/antonholmquist/jason"
	"github.com/dchest/uniuri"
	"github.com/jmoiron/sqlx"
	"github.com/johnnylee/sqlxchain"

	_ "github.com/go-sql-driver/mysql"
)

func checkMaxLevel(idparent string) int {
	var cnt int
	sql := "select max(numlevel) as numlevel from code where idparent = " + idparent

	dbx.Context().Begin().Get(&cnt, sql)

	return cnt
}

func countNumIdIndirect(idparent string, idindirect int) int {
	var cnt int
	sql := "select count(idindirect) as cnt from code " +
		//"where idparent = " + idparent + " numlevel = " + strconv.Itoa(level)
		"where idparent = " + idparent + " " +
		"and idindirect = " + strconv.Itoa(idindirect)

	dbx.Context().Begin().Get(&cnt, sql)

	return cnt
}

func checkNumMembers(idparent string, level int) int {
	var cnt int
	sql := "select count(id) as nummembers from code where idparent = " +
		idparent + " and numlevel = " + strconv.Itoa(level)

	dbx.Context().Begin().Get(&cnt, sql)

	return cnt
}

func genRandomString() string {
	var s string
	var cnt int

	for {
		s = uniuri.NewLen(4)

		sql := "select count(id) as cnt from acct where identifier = " + s

		dbx.Context().Begin().Get(&cnt, sql)

		if cnt == 0 {
			break
		}
	}

	return s
}

func insertDownlineThenCode(v *jason.Object, idparent string, level, idindirect int) error {
	var id int64

	sqlDownline := "insert into acct(identifier, fname, mname, lname," +
		"birthdate, address, gender, civilstatus, cpcountrycode, " +
		"cpnumber, email, jwtpassword) " +

		//sqlDownline := "insert into acct(identifier) values(?)"
		"values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	sqlCode := "INSERT INTO code(dttime, idpkg, idparent, iddownline, " +
		"idindirect, numlevel) VALUES (now(), ?, ?, ?, ?, ?)"

	identifier := genRandomString()
	fname, _ := v.GetString("fname")
	mname, _ := v.GetString("mname")
	lname, _ := v.GetString("lname")
	birthdate, _ := v.GetString("birthdate")
	address, _ := v.GetString("address")
	gender, _ := v.GetString("gender")
	civilstatus, _ := v.GetString("civilstatus")
	cpcountrycode, _ := v.GetString("cpcountrycode")
	cpnumber, _ := v.GetString("cpnumber")
	email, _ := v.GetString("email")
	jwtpassword, _ := v.GetString("jwtpassword")

	idpkg, _ := v.GetString("idpkg")

	err := dbx.Context().Begin().
		Exec(sqlDownline, identifier, fname, mname, lname,
			birthdate, address, gender, civilstatus,
			cpcountrycode, cpnumber, email, jwtpassword).
		LastInsertId(&id).
		Exec(sqlCode, idpkg, idparent, id, idindirect, level).
		Commit().Err()

	return err
}

type Records []Record

type Record struct {
	IdDownline int
}

func getIdIndirect(idparent string, level int) int {
	rec := Record{}

	rows, err := db.Queryx("select iddownline " +
		"from code " +
		"where idparent = " + idparent + " " +
		"order by iddownline")

	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err = rows.StructScan(&rec)
		if err != nil {
			break
		}
		//log.Println(rec)

		if countNumIdIndirect(idparent, rec.IdDownline) == 3 {
			continue
		} else {
			break
			//log.Println(rec.IdDownline)
			return rec.IdDownline
		}
	}
	return rec.IdDownline
}

func main() {
	var err error
	dbx, err = sqlxchain.New("mysql", "root:root1234@tcp(127.0.0.1:3306)/3x3")

	if err != nil {
		log.Fatal(err)
	}

	db = sqlx.MustConnect("mysql", "root:root1234@tcp(127.0.0.1:3306)/3x3")

	lkp := Lookup()
	//log.Println(lkp.Get(1))

	exampleJSON := `{
    			"fname": "Walter",
				"mname": "Volks",
				"lname": "Wagen",
				"birthdate": "2000-10-02",
				"address": "white plains",
				"gender": "M",
				"civilstatus": "single",
				"cpcountrycode": "63",
				"cpnumber": "9981234567",
				"email": "name@name.com",
				"jwtpassword": "1234",
				
				"idpkg": "1",
				"idparent": "1"
  			}`

	v, err := jason.NewObjectFromBytes([]byte(exampleJSON))
	if err != nil {
		return
	}

	idparent, _ := v.GetString("idparent")

	var level int

	/*	table code (dttime, idpkg)
		id 	idparent	iddownline	idindirect	numlevel
		1		1			2			1			1
		2		1			3			1			1
		3		1			4			1			1

		4		1			5			2			2
		5		1			6			2			2
		6		1			7			2			2

		7		1			8			3			2
		8		1			9			3			2
		9		1			10			3			2

		10		1			11			4			2
		11		1			12			4			2
		12		1			13			4			2
	*/

	level = checkMaxLevel(idparent)
	log.Println(level)

	if level >= 1 && level <= 10 {
		if checkNumMembers(idparent, level) < lkp.Get(level) {

			//insert record
			if level == 1 {
				insertDownlineThenCode(v, idparent, level, 1)
			} else {
				idindirect := getIdIndirect(idparent, level)
				insertDownlineThenCode(v, idparent, level, idindirect)
			}

		} else {
			//auto-fill to next level

			idindirect := getIdIndirect(idparent, level)

			if level == 10 {
				insertDownlineThenCode(v, idparent, 10, idindirect)
			} else {
				insertDownlineThenCode(v, idparent, level+1, idindirect)
			}

		}
	} else {

		//first record

		insertDownlineThenCode(v, idparent, 1, 1)

		if err != nil {
			log.Println(err)

			//fmt.Fprintf(w, "Error in saving transaction to database")
		} else {
			//fmt.Fprintf(w, "")
		}

	}
}
