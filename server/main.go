package main

import (
	"github.com/Anvozz/book-rental-shop/database"
	"github.com/Anvozz/book-rental-shop/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.Connect()
	var user []models.User
	// bob := models.User{Name: "Wongsathorn Kanno" , Email: "sankaapb@gmail.com" , Address:  "22" , Status: 1 , Tel: "0649453094" , Point: 500}
	// db.Create(&bob)
	db.Find(&user)
	app := fiber.New()

  app.Get("/users", func(c *fiber.Ctx) error {
    return c.JSON(user)
  })

  app.Listen(":3000")
}