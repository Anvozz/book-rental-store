package models

import "gorm.io/gorm"

type Statement struct {
	gorm.Model
	Name   string
	Type   int16
	Price  int64
	Status int16
}
