package order

import (
	"encoding/json"
	"fmt"
	"github.com/bpiatek/roommanager/src"
	"log"
	"time"
)

func makeOrderDB(userId int)  (b bool, err error){
	stmtIns, err := src.GetDB().Prepare("INSERT INTO orders VALUES( ?, ?, ? )")
	if err != nil {
	panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close()

	_, err = stmtIns.Exec(0, time.Now(), userId)
	if err != nil {
		return false, err
	}
	log.Printf("Added order to database.")

	return  true, err
}

func getOrdersByCustomerId(customerId int) []byte {
	var (
		order  Order
		orders []Order
	)
	stmt, err := src.GetDB().Prepare("SELECT * FROM orders where customer_id=?")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rows, e := stmt.Query(customerId)
	for rows.Next() {
		rows.Scan(&order.Id, &order.Date, &order.CustomerId)
		if e != nil {
			panic(e.Error())
		}
		orders = append(orders, order)
	}

	defer stmt.Close()

	jsonResponse, jsonError := json.Marshal(orders)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}

	return jsonResponse
}
