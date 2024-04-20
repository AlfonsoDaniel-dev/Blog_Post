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

}
