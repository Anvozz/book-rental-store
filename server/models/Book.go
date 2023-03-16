package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	CategoryId uint
	Category Category `gorm:"foreignKey:CategoryId"`
	Name string
	Amount int32
	Desscription *string
	Status int16
}