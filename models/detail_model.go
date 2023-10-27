package models

import "time"

type OrderDetail struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	OrderID    uint    `json:"order_id"`
	FoodID     uint    `json:"food_id"`
	Invoice    string  `json:"invoice"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  time.Time
}
