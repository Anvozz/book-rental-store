package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Anvozz/book-rental-shop/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type CategoryRestHandler interface {
	CreateCategory(c *fiber.Ctx) error
	GeteCategory(c *fiber.Ctx) error
	PutCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type CategoryHandler struct{}

var validate = validator.New()

func validateStruct(category models.Category) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(category)
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

func (r *restHandler) CreateCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validateStruct(*category)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	res := r.db.Create(&category)

	if res.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(res.Error)
	}

	return c.Status(fiber.StatusCreated).
		JSON(ResponseMessage{Message: "Create category successfully."})
}

func (r *restHandler) GeteCategory(c *fiber.Ctx) error {
	var category []models.Category
	r.db.Order("id ASC").Find(&category)
	return c.Status(http.StatusOK).JSON(category)
}

func (r *restHandler) PutCategory(c *fiber.Ctx) error {
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := validateStruct(*category)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	if category.ID == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(ResponseMessage{Message: "Invalid argument field Id"})
	}

	r.db.Save(
		&models.Category{
			ID:        category.ID,
			Name:      category.Name,
			Status:    category.Status,
			CreatedAt: time.Now(),
		},
	)

	return c.Status(fiber.StatusOK).JSON(ResponseMessage{Message: "Update category successfully."})
}

func (r *restHandler) DeleteCategory(c *fiber.Ctx) error {
	idParam, err := strconv.ParseUint(c.Params("id"), 10, 32)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	category := &models.Category{ID: uint(idParam)}
	r.db.Delete(&category)
	return c.Status(fiber.StatusOK).JSON(ResponseMessage{Message: "Delete category successfully."})
}
