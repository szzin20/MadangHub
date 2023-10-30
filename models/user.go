package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string  `json:"username" form:"username" gorm:"unique;not null"`
	Email    string  `json:"email" form:"email" gorm:"unique;not null"`
	Password string  `json:"password" form:"password" gorm:"not null"`
	Address  string  `json:"address" form:"address"`
	Role     string  `gorm:"default:user" json:"role"`
	Order    []Order `gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
