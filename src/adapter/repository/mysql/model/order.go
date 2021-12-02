package model

import "time"

type Order struct {
	Id          int `gorm:"primaryKey"`
	UserId      string
	TotalAmount int64
	CreatedAt   time.Time
}
