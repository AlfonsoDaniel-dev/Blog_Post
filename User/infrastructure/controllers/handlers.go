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

	data := models2.Register{}

	err := c.Bind(&data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusBadRequest, response)
	}

	err = U.Services.Register(data)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error registering user")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(nil, "Success", "User created")
	return c.JSON(http.StatusCreated, response)
}

func (U *Handler) Login(c echo.Context) error {
	data := models2.Login{}
	if c.Request().Method != http.MethodPost {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}
	err := c.Bind(&data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusBadRequest, response)
	}
	token, err := U.Services.Login(data)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error login")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(token, "Login Success", "User Logged In")
	return c.JSON(http.StatusOK, response)
}

func (H *Handler) GetAll(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	users, err := H.Services.GetAllUsers()
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error getting all users")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(users, "Success", "Users retrieved")
	return c.JSON(http.StatusOK, response)
}
