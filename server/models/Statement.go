package models

import "time"

type Statement struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `                  json:"createdAt"`
	UpdatedAt   time.Time `                  json:"updatedAt"`
	Description string    `                  json:"description" validate:"require"`
	Type        int16     `                  json:"type"        validate:"require"`
	Price       int64     `                  json:"price"       validate:"require"`
	Status      int16     `                  json:"status"      validate:"require"`
}
