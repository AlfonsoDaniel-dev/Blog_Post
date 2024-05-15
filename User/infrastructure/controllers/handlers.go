package controllers

import (
	"fmt"
	"github.com/TeenBanner/Inventory_system/User/App/Services"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/TeenBanner/Inventory_system/User/infrastructure/controllers/responses"
	"github.com/TeenBanner/Inventory_system/pkg/authorization"
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

	login := models2.Login{
		Email:    data.Email,
		Password: data.Password,
	}

	token, err := U.Services.Login(login)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error registering user")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(map[string]string{"token": token}, "Success", "User created")
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

func (H *Handler) CreatePost(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}
	data := models2.CreatePost{}
	token := c.Request().Header.Get("Authorization")
	email, err := authorization.GetEmailFromJWT(token)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error getting user email")
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = c.Bind(&data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusBadRequest, response)
	}

	post, err := H.Services.CreatePost(email, data)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error creating post")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(post, "Success", "Post Created Successfully")
	return c.JSON(http.StatusCreated, response)
}

func (H *Handler) GetPostsByTitleAndEmail(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	title := c.Param("title")
	email := c.Param("email")

	posts, err := H.Services.GetPostByTitleAndEmail(title, email)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error getting all posts")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(posts, "Success", "Posts retrieved")
	return c.JSON(http.StatusOK, response)
}

func (H *Handler) GetAllPostsFromEmail(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	email := c.Param("email")
	fmt.Println(email)
	if email == "" {
		response := responses.NewResponse(email, "Error", "Error getting user")
		return c.JSON(http.StatusInternalServerError, response)
	}

	posts, err := H.Services.GetAllPostsFromUserEmail(email)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error getting all posts")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(posts, "Success", "Posts retrieved")
	return c.JSON(http.StatusOK, response)
}
