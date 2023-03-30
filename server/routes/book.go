package routes

import (
	"strconv"
	"time"

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

func (r *restHandler) GetBook(c *fiber.Ctx) error {
	var book []models.Book
	err := r.db.Model(&models.Book{}).Preload("Category").Order("id ASC").Find(&book).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(err)
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func (r *restHandler) GetBookByid(c *fiber.Ctx) error {
	id := c.Params("id")
	bookId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	book := models.Book{ID: uint(bookId)}
	result := r.db.Find(&book)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{Message: "No book found"})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (r *restHandler) PutBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := bookvalidateStruct(*book)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if book.ID == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(ResponseMessage{Message: "Invalid argument field Id"})
	}

	r.db.Save(
		&models.Book{
			ID:           book.ID,
			Name:         book.Name,
			Amount:       book.Amount,
			Status:       book.Status,
			CategoryID:   book.CategoryID,
			Desscription: book.Desscription,
			UpdatedAt:    time.Now(),
		},
	)

	return c.Status(fiber.StatusOK).JSON(ResponseMessage{Message: "Update book successfully."})
}

func (r *restHandler) DeleteBook(c *fiber.Ctx) error {
	idParam, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	book := &models.Book{ID: uint(idParam)}
	r.db.Delete(&book)
	return c.Status(fiber.StatusOK).JSON(ResponseMessage{Message: "Delete book successfully."})
}
