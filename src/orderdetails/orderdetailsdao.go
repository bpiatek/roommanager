package orderdetails

import (
	"fmt"
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

func GetOrderDetailsByOrderId(orderId int) (o OrderDetailsDTO, r error) {
	var orderDetails OrderDetailsDTO;

	rows, err := src.GetDB().Prepare("SELECT room_id, num_of_minutes FROM orderdetails where order_id=?")
	if err != nil {
		fmt.Println(err)
		return orderDetails, err
	}

	defer rows.Close()

	err = rows.QueryRow(orderId).Scan(&orderDetails.RoomId, &orderDetails.Minutes)
	if err != nil {
		fmt.Println(err.Error())
	}


	return orderDetails, err

}