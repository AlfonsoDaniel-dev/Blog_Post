package user

import (
	"net/http"

	"github.com/TeenBanner/Inventory_system/helpers"
	"github.com/TeenBanner/Inventory_system/models"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	i input
}

func NewUserHandler(I input) userHandler {
	return userHandler{
		i: I,
	}
}

func (uh *userHandler) Create(c echo.Context) error {
	user := models.User{}

	err := c.Bind(&user)
	if err != nil {
		response := helpers.NewResponse("error", "Peticion mal estructurada", err)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = uh.i.CreateUser(user)
	if err != nil {
		response := helpers.NewResponse("Error", "Error Creando el usuario", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helpers.NewResponse("OK", "Usuario Creado", user)

	return c.JSON(http.StatusCreated, response)
}

func (uh *userHandler) GetAll(c echo.Context) error {
	users, err := uh.i.GetAllUsers()
	if err != nil {
		response := helpers.NewResponse("Error", "Error Obteninedo todos los usuarios", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := helpers.NewResponse("OK", "usuarios obtenidos correctamente", users)
	return c.JSON(http.StatusOK, response)
}

func 
