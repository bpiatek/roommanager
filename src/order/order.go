package order

import "time"

type Order struct {
	Id int
	Date time.Time
	CustomerId int
}
