package models

import (
	"time"

	"gorm.io/gorm"
)

type Bookhistory struct {
	gorm.Model
	BookId uint
	Book Book `gorm:"foreignKey:BookId"`
	Duedate time.Time
	Price float64
	Status int16
	Day int16
}