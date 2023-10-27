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

func (f *Food) DecrementFood(quantity int) {
	if f.Stock >= quantity {
		f.Stock -= quantity
	}
}

func (f *Food) IncrementFood(quantity int) {
	f.Stock += quantity
}

func (f *Food) TotalCost(quantity int) float64 {
	return f.Price * float64(quantity)
}