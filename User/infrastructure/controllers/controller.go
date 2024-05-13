package controllers

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/User/App/Services"
	UserDomain "github.com/TeenBanner/Inventory_system/User/Domain"
	"github.com/TeenBanner/Inventory_system/User/infrastructure/DB/psqlUser"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HanlderServices interface {
	Register(c echo.Context) error
	GetAll(c echo.Context) error
}

type UserController struct {
	Services.Services
	HanlderServices
}

func NewUserRouter(e *echo.Echo, db *sql.DB) {
	controller := BuildUserController(e, db)

	controller.PublicRoutes(e)
}

func BuildUserController(e *echo.Echo, DB *sql.DB) *UserController {
	UserStorage := psqlUser.NewPsqlUser(DB)
	user := UserDomain.NewUser(UserStorage)
	service := Services.NewServices(user)

	handler := NewHandler(service)

	return &UserController{
		service,
		handler,
	}
}

func (h *UserController) PublicRoutes(e *echo.Echo) {

	e.Use(middleware.Recover())

	e.POST("/", h.HanlderServices.Register)
	e.GET("/users", h.HanlderServices.GetAll)
}
