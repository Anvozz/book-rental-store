package models

import "time"

type Statement struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `                  json:"createdAt"`
	UpdatedAt   time.Time `                  json:"updatedAt"`
	Description string    `                  json:"description" validator:"require"`
	Type        int16     `                  json:"type"        validator:"require"`
	Price       int64     `                  json:"price"       validator:"require"`
	Status      int16     `                  json:"status"      validator:"require"`
}
