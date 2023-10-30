package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"foreignKey:UserID" json:"user_id"`
	Address    string      `json:"address"`
	Longitude  string      `json:"longitude"`
	Latitude   string      `json:"latitude"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
}

type OrderItem struct {
	gorm.Model
	OrderID  uint `json:"order_id" gorm:"index"`
	FoodID   uint `json:"food_id"`
	Quantity int  `json:"quantity"`
	Foods    Food `gorm:"foreignKey:FoodID"`
}
