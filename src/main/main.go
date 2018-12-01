package main

import (
	"fmt"
	"github.com/bpiatek/roommanager/src"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Server will start at http://localhost:8080/")

	src.ConnectDatabse()

	route := mux.NewRouter()
	AddApproutes(route)

	log.Fatal(http.ListenAndServe(":8080", route))

}