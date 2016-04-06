//https://academy.datastax.com/resources/getting-started-apache-cassandra-and-go

package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demo"
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a user
	if err := session.Query("INSERT INTO users (lastname, age, city, email, firstname) VALUES ('Jones', 35, 'Austin', 'bob@example.com', 'Bob')").Exec(); err != nil {
		log.Fatal(err)
	}

	// Use select to get the user we just entered
	var firstname, lastname, city, email string
	var age int

	if err := session.Query("SELECT firstname, age FROM users WHERE lastname='Jones'").Scan(&firstname, &age); err != nil {
		log.Fatal(err)
	}
	fmt.Println(firstname, age)

	// Update the same user with a new age
	if err := session.Query("UPDATE users SET age = 36 WHERE lastname = 'Jones'").Exec(); err != nil {
		log.Fatal(err)
	}

	// Select and show the change
	iter := session.Query("SELECT firstname, age FROM users WHERE lastname='Jones'").Iter()
	for iter.Scan(&firstname, &age) {
		fmt.Println(firstname, age)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	// Delete the user from the users table
	if err := session.Query("DELETE FROM users WHERE lastname = 'Jones'").Exec(); err != nil {
		log.Fatal(err)
	}

	// Show that the user is gone
	session.Query("SELECT * FROM users").Iter()
	for iter.Scan(&lastname, &age, &city, &email, &firstname) {
		fmt.Println(lastname, age, city, email, firstname)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
