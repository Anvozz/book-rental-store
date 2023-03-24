package models

import "time"

type Book struct {
	ID           uint      `gorm:"primarykey"            json:"id"`
	CreatedAt    time.Time `                             json:"createdAt"`
	UpdatedAt    time.Time `                             json:"updatedAt"`
	CategoryId   uint      `                             json:"categoryId"  validate:"required"`
	Category     Category  `gorm:"foreignKey:CategoryId" json:"category"`
	Name         string    `                             json:"name"        validate:"required,min=3,max=200"`
	Amount       int32     `                             json:"amount"      validate:"required"`
	Desscription *string   `                             json:"description"`
	Status       int16     `                             json:"status"      validate:"required,max=1"`
}
