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
	Login(c echo.Context) error

	CreatePost(c echo.Context) error

	GetPostsByTitleAndEmail(c echo.Context) error
	GetAllPostsFromEmail(c echo.Context) error
	GetPostsFromName(c echo.Context) error
	GetAllPosts(c echo.Context) error

	UserUpdateTheirEmail(c echo.Context) error
	UserUpdateTheirName(c echo.Context) error
	UserUpdateTheirPassword(c echo.Context) error

	UserUpdatePostTitle(c echo.Context) error
	UserUpdatePostBody(c echo.Context) error

	UserDeletePost(c echo.Context) error
	UserDeleteTheirAccount(c echo.Context) error

	UserGetTheirInfo(c echo.Context) error

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

	users.PUT("/user/email", h.HanlderServices.UserUpdateTheirEmail)
	users.PUT("/user/name", h.HanlderServices.UserUpdateTheirName)
	users.PUT("/user/password", h.HanlderServices.UserUpdateTheirPassword)
	users.GET("/user/posts", h.HanlderServices.GetAllPostsFromEmail)

	users.GET("/user", h.HanlderServices.UserGetTheirInfo)
	users.GET("/", h.HanlderServices.GetAllPostsFromEmail)
	users.GET("/users", h.HanlderServices.GetAll)

	users.POST("/post", h.HanlderServices.CreatePost)

	users.PUT("/post/update/title", h.HanlderServices.UserUpdatePostTitle)
	users.PUT("/post/update/body", h.HanlderServices.UserUpdatePostBody)
	users.DELETE("/post/delete/:title", h.HanlderServices.UserDeletePost)

	users.POST("/users/account/delete", h.HanlderServices.UserDeleteTheirAccount)
}

func (h *UserController) PublicRoutes(e *echo.Echo) {

	e.Use(middleware.Recover())

	public := e.Group("/api/v1/public")
	public.GET("/:name", h.HanlderServices.GetPostsFromName)
	public.GET("/:title/:email", h.HanlderServices.GetPostsByTitleAndEmail)
	public.GET("/posts/all", h.HanlderServices.GetAllPosts)
	public.POST("/register", h.HanlderServices.Register)
	public.POST("/Login", h.HanlderServices.Login)
}
