package controllers

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/User/App/Services"
	Domain "github.com/TeenBanner/Inventory_system/User/Domain"
	"github.com/TeenBanner/Inventory_system/User/infrastructure/DB/psqlUser"
	"github.com/TeenBanner/Inventory_system/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HanlderServices interface {
	Register(c echo.Context) error
	GetAll(c echo.Context) error
	Login(c echo.Context) error
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
	user := Domain.NewUser(UserStorage)
	service := Services.NewServices(user)

	handler := NewHandler(service)

	return &UserController{
		service,
		handler,
	}
}

func (h *UserController) PrivateRoutes(e *echo.Echo) {

	users := e.Group("/api/v1/private")
	users.Use(middlewares.AuthMiddleware)
	users.GET("/users", h.HanlderServices.GetAll)
}
func (h *UserController) PublicRoutes(e *echo.Echo) {

	e.Use(middleware.Recover())

	public := e.Group("/api/v1/public")
	public.POST("/register", h.HanlderServices.Register)
	public.POST("/Login", h.HanlderServices.Login)
}
