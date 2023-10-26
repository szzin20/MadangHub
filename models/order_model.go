package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	FoodID    uint      `json:"food_id"`
	Quantity  int       `json:"quantity"`
	Invoice   string    `json:"invoice"`
	TotalCost int       `json:"total_cost"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time
}
