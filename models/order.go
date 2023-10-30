package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID     uint    `gorm:"foreignKey:UserID" json:"user_id"`
	FoodsID    uint    `json:"foods_id"`
	Address    string  `json:"address"`
	Longitude  string  `json:"longitude"`
	Latitude   string  `json:"latitude"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Foods      Food    `gorm:"foreignKey:FoodsID"`
}

type OrderInput struct {
	UserID   uint    `json:"user_id" validate:"required"`
	UserName string  `json:"user_name" validate:"required"`
	FoodID   uint    `json:"food_id" validate:"required"`
	Quantity int     `json:"quantity" validate:"required, min:1"`
	Total    float64 `json:"total"`
}
