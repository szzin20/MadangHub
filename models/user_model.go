package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint      `gorm:"type:varchar;primaryKey;not null" json:"id"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email" gorm:"unique"`
	Password  string    `json:"password" form:"password"`
	Phone     string    `json:"phone" form:"phone"`
	Address   string    `gorm:"type:longtext" json:"address"`
	Role      string    `gorm:"type:enum('admin', 'user');default:'user'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAT time.Time `json:"deleted_at"`
}