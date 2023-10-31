package models

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Status      string  `json:"status"`
}

type FoodResponse struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Available   bool    `json:"available"`
	Price       float64 `json:"price" validate:"required"`
	Status      string  `json:"status" validate:"required"`
}
