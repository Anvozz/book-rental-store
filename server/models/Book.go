package models

import "time"

type Book struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time `                  json:"createdAt"`
	UpdatedAt    time.Time `                  json:"updatedAt"`
	CategoryID   uint      `                  json:"categoryId"  validate:"required"`
	Category     Category  `                  json:"category"    validate:"-"`
	Name         string    `                  json:"name"        validate:"required,min=3,max=200"`
	Amount       int32     `                  json:"amount"      validate:"required"`
	Desscription *string   `                  json:"description"`
	Status       int16     `                  json:"status"      validate:"required,max=1"`
}
