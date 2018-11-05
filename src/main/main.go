package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
	db, err := sql.Open("mysql", "dupa:dupa123@/roomDb")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("DUPA")

	stmtIns, err := db.Prepare("INSERT INTO person VALUES( ?, ? )") // ? = placeholder

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close() // Close the statement when

	_, err = stmtIns.Exec("Krzysztof", "Nowak") // Insert tuples (i, i^2)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

