package model

import "time"

type Order struct {
	Id          int `gorm:"primaryKey"`
	TotalAmount int64
	CreatedAt   time.Time
}
