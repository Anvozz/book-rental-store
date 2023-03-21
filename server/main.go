package main

import (
	"log"

	"github.com/Anvozz/book-rental-shop/database"
	"github.com/Anvozz/book-rental-shop/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func setupRoutes(c *fiber.App, db *gorm.DB) {
	r := routes.New(db)

	// Category
	c.Get("/category", r.GeteCategory)
	c.Post("/category", r.CreateCategory)
	c.Put("/category", r.PutCategory)
	c.Delete("/category/:id", r.DeleteCategory)

	// User
	c.Post("/users", r.CreateUser)
	c.Get("/users", r.GetUser)
	c.Get("/users/:id", r.GetUserByid)
}

func main() {
	log.Println("Starting book api server")
	// Setup Databse
	db := database.Connect()
	// Create fiber instance
	app := fiber.New()
	app.Use(logger.New())
	// Setup routes
	setupRoutes(app, db)
	// Listen
	app.Listen(":3000")
}
