package main

import (
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	db, err := bolt.Open("bolt.db", 0666, nil)

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	// Execute several commands within a write transaction.
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("widgets"))
		if err != nil {
			return err
		}
		if err := b.Put([]byte("foo"), []byte("bar")); err != nil {
			return err
		}

		c, err := b.CreateBucketIfNotExists([]byte("nested"))
		if err != nil {
			return err
		}
		err = c.Put([]byte("nested.foo"), []byte("nested.foo"))

		return nil
	})

	// If our transactional block didn't return an error then our data is saved.
	if err == nil {
		messages := make(chan string)

		db.View(func(tx *bolt.Tx) error {
			value := tx.Bucket([]byte("widgets")).Get([]byte("foo"))
			//log.Println(value)
			log.Printf("The value of 'foo' within the transaction is: %s\n", value)

			go func() { messages <- string(value) }()

			return nil
		})

		msg := <-messages
		log.Printf("The value of 'foo' outside the transaction is: %s\n", msg)
	}

	//view nested bucket
	db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte("widgets")).Bucket([]byte("nested")).Get([]byte("nested.foo"))
		//log.Println(value)
		log.Printf("The value of 'nested.foo' is: %s\n", value)

		return nil
	})
}
