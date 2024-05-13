package controllers

import (
	"github.com/TeenBanner/Inventory_system/User/App/Services"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/TeenBanner/Inventory_system/User/infrastructure/controllers/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	Services.Services
}

func NewHandler(services Services.Services) *Handler {
	return &Handler{
		Services: services,
	}
}

func (U *Handler) Register(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	data := models2.User{}

	err := c.Bind(&data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := U.Services.Register(data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error registering user")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(user, "Success", "User created")
	return c.JSON(http.StatusCreated, response)
}
