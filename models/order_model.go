package models

import (
	"os/user"
	"time"
)

type Status int

const (
	StatusCreated    Status = 0
	StatusInProgress Status = 1
	StatusDone       Status = 2
	StatusCanceled   Status = 3
)

var Statuses = []Status{StatusCreated, StatusInProgress, StatusDone, StatusCanceled}

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint
	User      user.User `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	Status    Status `gorm:"type:smallint;check:status IN (0,1,2,3)"`
	Total     float64   `gorm:"check:total >= 0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Cost calculates the total cost of order item.
// IsValidStatus checks whether provided status is a valid Status.
// Useful to validate input that comes from external sources, e.g as
// a query parameter.
func IsValidStatus(status int) bool {
	for _, s := range Statuses {
		if status == int(s) {
			return true
		}
	}

	return false
}