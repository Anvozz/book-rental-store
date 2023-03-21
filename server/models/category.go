package models

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `                  json:"createdAt"`
	UpdatedAt time.Time `                  json:"updatedAt"`
	Name      string    `                  json:"name"      validate:"required,min=3,max=32"`
	Status    int64     `                  json:"status"    validate:"required,number,max=1"`
}
