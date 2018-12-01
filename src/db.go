package src

import (
	"database/sql"
	"fmt"
)

var db *sql.DB
var err error

func ConnectDatabse() {
	db, err = sql.Open("mysql", "dupa:dupa123@/roomDb?parseTime=true")
	fmt.Println("Database connected.")
}

func GetDB() *sql.DB {
	return db
}
