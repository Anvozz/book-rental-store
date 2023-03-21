package routes

import (
	"strconv"

	"github.com/Anvozz/book-rental-shop/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserRestHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	GetUserByid(c *fiber.Ctx) error
	// PutUser(c *fiber.Ctx) error
	// DeleteUser(c *fiber.Ctx) error
}

type UserHandler struct{}

var uservalidate = validator.New()

func uservalidateStruct(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := uservalidate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (r *restHandler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := uservalidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res := r.db.Create(&user)

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.Error)
	}

	return c.Status(fiber.StatusCreated).
		JSON(ResponseMessage{Message: "Create category successfully."})
}

func (r *restHandler) GetUser(c *fiber.Ctx) error {
	var users []models.User
	r.db.Order("id ASC").Find(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}

func (r *restHandler) GetUserByid(c *fiber.Ctx) error {
	id := c.Params("id")
	userId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	user := models.User{ID: uint(userId)}
	result := r.db.Find(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{Message: "No user found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
