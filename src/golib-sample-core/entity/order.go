package entity

import "time"

type Order struct {
	Id          int
	TotalAmount int64
	CreatedAt   time.Time
}
