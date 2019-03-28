package orderdetails

import (
	"github.com/bpiatek/roommanager/src"
	"log"
)

func AddOrderDetailsDB(orderId int, details OrderDetailsDTO) (b bool, err error) {
	stmtIns, err := src.GetDB().Prepare("INSERT INTO orderdetails VALUES( ?, ?, ? )")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close()
	_, err = stmtIns.Exec(orderId, details.RoomId, details.Minutes)

	if err != nil {
		return false, err
	}
	log.Printf("Added orderdetails to database.")

	return  true, err
}