package database

import (
	"fmt"
	"log"

	"github.com/Anvozz/book-rental-shop/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=book_db port=5832 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect database")
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.Category{},
		&models.Book{},
		&models.Bookhistory{},
		&models.Statement{},
		&models.User{},
	)
	return db
}
