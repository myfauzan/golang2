package models

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `gorm:"not null;type:varchar(10)" json:"itemCode"`
	Description string `gorm:"not null;type:varchar(100)" json:"Description"`
	Quantity    int
	OrderID     uint
}