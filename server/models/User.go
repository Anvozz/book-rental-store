package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string
	Address string
	Tel string
	Point int32
	Status int16
}