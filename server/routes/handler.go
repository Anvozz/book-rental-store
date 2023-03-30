package routes

import "gorm.io/gorm"

type RestHandler interface {
	CategoryRestHandler
	UserRestHandler
	BookRestHandler
}

type restHandler struct {
	db *gorm.DB
	CategoryHandler
	UserHandler
	BookHandler
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func New(db *gorm.DB) RestHandler {
	return &restHandler{db: db}
}
