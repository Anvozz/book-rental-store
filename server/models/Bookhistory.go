package models

import (
	"time"
)

type Bookhistory struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `                  json:"createdAt"`
	UpdatedAt time.Time `                  json:"updatedAt"`
	BookId    uint
	Book      Book
	Duedate   time.Time
	Price     int64
	Status    int16
	Day       int16
}
