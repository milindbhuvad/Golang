package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver for side effects
)

func main() {
	// Data Source Name (DSN) format: username:password@tcp(host:port)/database
	dsn := "root:@tcp(127.0.0.1:3306)/task" // Modify with your own credentials
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ping the database to ensure connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the database!")

	// Perform a SELECT query
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // Make sure to close rows after the query is done

	// Iterate over the result set and print out each row
	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	// Check for any error after iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
