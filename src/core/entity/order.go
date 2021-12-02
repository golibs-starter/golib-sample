package entity

import "time"

type Order struct {
	Id          int
	UserId      string
	TotalAmount int64
	CreatedAt   time.Time
}
