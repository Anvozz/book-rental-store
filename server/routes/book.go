package routes

import (
	"github.com/Anvozz/book-rental-shop/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type BookRestHandler interface {
	CreateBook(c *fiber.Ctx) error
	GetBook(c *fiber.Ctx) error
	GetBookByid(c *fiber.Ctx) error
	PutBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
}

type BookHandler struct{}

var bookvalidate = validator.New()

func bookvalidateStruct(book models.Book) []*ErrorResponse {
	var errors []*ErrorResponse
	err := bookvalidate.Struct(book)
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

func (r *restHandler) CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := bookvalidateStruct(*book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	res := r.db.Create(&book)

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.Error)
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseMessage{Message: "Create book successfully."})
}