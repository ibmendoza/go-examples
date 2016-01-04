package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func createBucket(db *bolt.DB, value string) (err error) {
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(value))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func insertAcct(db *bolt.DB, value string) (err error) {
	// Execute several commands within a write transaction.
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("acct"))

		if err := b.Put([]byte("cp"), []byte(value)); err != nil {
			return err
		}

		return nil
	})

	return err
}

func insertAccts(db *bolt.DB, prefix string) {
	var value, suffix string
	var err error
	for i := 0; i != 10000; i++ {

		suffix = fmt.Sprintf("%04d", i)
		value = prefix + suffix

		fmt.Println(value)

		err = insertAcct(db, value)

		if err != nil {
			log.Println("error in boltdb insert")
		}
	}
}

func main() {
	db, err := bolt.Open("sms.db", 0666, nil)

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	err = createBucket(db, "acct")
	if err != nil {
		log.Fatalln(err)
	}

	insertAccts(db, "408956")
	insertAccts(db, "212547")

	fmt.Println("Done")
	fmt.Scanln()
}
