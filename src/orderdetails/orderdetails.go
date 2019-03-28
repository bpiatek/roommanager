package orderdetails

import "database/sql"

type OrderDetails struct {
	OrderId int				`json:"order_id"`
	RoomId int				`json:"room_id"`
	Minutes sql.NullInt64	`json:"minutes"`

}

type OrderDetailsDTO struct {
	RoomId int				`json:"room_id"`
	Minutes int				`json:"minutes"`
}