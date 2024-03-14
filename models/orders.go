package models

import "time"

type Order struct {
	OrderID			uint	`gorm:"primaryKey"`
	CustomerName	string	`gorm:"not null;type:varchar(100)"`
	Items			[]Item
	OrderedAt		time.Time
}