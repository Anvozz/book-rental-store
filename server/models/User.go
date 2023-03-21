package models

import "time"

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `                  json:"createdAt"`
	UpdatedAt time.Time `                  json:"updatedAt"`
	Name      string    `                  json:"name"      validate:"required,min=3,max=64"`
	Email     string    `                  json:"email"     validate:"required"`
	Address   string    `                  json:"address"`
	Tel       string    `                  json:"tel"`
	Point     int32     `                  json:"point"`
	Status    int16     `                  json:"status"    validate:"required,number,max=1"`
}
