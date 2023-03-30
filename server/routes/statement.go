package routes

import (
	"strconv"

	"github.com/Anvozz/book-rental-shop/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type StatementRestHandler interface {
	CreateStatement(c *fiber.Ctx) error
	GetStatement(c *fiber.Ctx) error
	GetStatementByid(c *fiber.Ctx) error
	PutStatement(c *fiber.Ctx) error
	DeleteStatement(c *fiber.Ctx) error
}

type StatementHandler struct{}

var statementvalidate = validator.New()

func statementvalidateStruct(statement models.Statement) []*ErrorResponse {
	var errors []*ErrorResponse
	err := bookvalidate.Struct(statement)
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

func (r *restHandler) GetStatement(c *fiber.Ctx) error {
	var statement []models.Statement
	r.db.Order("id ASC").Find(&statement)
	return c.Status(fiber.StatusOK).JSON(statement)
}

func (r *restHandler) GetStatementByid(c *fiber.Ctx) error {
	id := c.Params("id")
	statementId, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	statement := models.Statement{ID: uint(statementId)}
	result := r.db.Find(&statement)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(ResponseMessage{Message: "Statement not found"})
	}

	return c.Status(fiber.StatusOK).JSON(statement)
}
