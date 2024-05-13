package controllers

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/User/App/Services"
	UserDomain "github.com/TeenBanner/Inventory_system/User/Domain"
	"github.com/TeenBanner/Inventory_system/User/infrastructure/DB/psqlUser"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Services.Services
}

func BuildUserController(e echo.Echo, DB *sql.DB) *UserController {
	UserStorage := psqlUser.NewPsqlUser(DB)
	user := UserDomain.NewUser(UserStorage)
	service := Services.NewServices(user)

	return &UserController{
		service,
	}
}

func PublicRoutes(e *echo.Echo) {

}
