package models

import "time"

type Food struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time 
	DeletedAT   time.Time 
}

func (e *Food) DecrementFood(quantity int) {
	if e.Stock >= quantity {
		e.Stock -= quantity
	}
}
