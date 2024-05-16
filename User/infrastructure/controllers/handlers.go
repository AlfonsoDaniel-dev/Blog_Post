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

	response := responses.NewResponse(map[string]any{"token": token, "user": data}, "Success", "User created")
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

	response := responses.NewResponse(map[string]string{"token": token}, "Login Success", "User Logged In")
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

	token := c.Request().Header.Get("Authorization")
	email, err := authorization.GetEmailFromJWT(token)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error Getting userPosts")
		return c.JSON(http.StatusInternalServerError, response)
	}

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

func (H *Handler) GetPostsFromName(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	name := c.Param("name")

	posts, err := H.Services.GetAllPostsFromName(name)
	if err != nil {
		response := responses.NewResponse(err, "ERROR", "Error obteniendo los posts")
		return c.JSON(http.StatusNotFound, response)
	}

	response := responses.NewResponse(posts, "Ok", "Posts Get Succesfully")
	return c.JSON(http.StatusOK, response)
}

func (H *Handler) UserGetTheirInfo(c echo.Context) error {
	if c.Request().Method != http.MethodGet {
		response := responses.NewResponse(nil, "Error", "Mehod not allowe")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	token := c.Request().Header.Get("Authorization")

	email, err := authorization.GetEmailFromJWT(token)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error while getting email from token")
		return c.JSON(http.StatusForbidden, response)
	}
	info, err := H.Services.GetUserByEmail(email)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error getting user Data")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(info, "Success", "User data delivered")
	return c.JSON(http.StatusOK, response)
}

func (H *Handler) UserUpdateTheirName(c echo.Context) error {
	if c.Request().Method != http.MethodPut {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	token := c.Request().Header.Get("Authorization")

	data := models2.UpdateNameForm{}

	err := c.Bind(&data)

	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusInternalServerError, response)
	}

	email, err := authorization.GetEmailFromJWT(token)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error while getting email from token")
		return c.JSON(http.StatusForbidden, response)
	}

	err = H.Services.UpdateName(email, data)

	response := responses.NewResponse(data, "Succes", "Name updated Succesfully")
	return c.JSON(http.StatusOK, response)
}

func (H *Handler) UserUpdateTheirEmail(c echo.Context) error {
	if c.Request().Method != http.MethodPut {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	token := c.Request().Header.Get("Authorization")
	email, err := authorization.GetEmailFromJWT(token)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Error while getting email from token")
		return c.JSON(http.StatusForbidden, response)
	}

	data := models2.UpdateEmailForm{}
	err = c.Bind(&data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusInternalServerError, response)
	}

	err = H.Services.UpdateEmail(email, data)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error While updating email")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(data, "Success", "Email updated Succesfully")
	return c.JSON(http.StatusOK, response)
}

func (H *Handler) UserUpdateTheirPassword(c echo.Context) error {
	if c.Request().Method != http.MethodPut {
		response := responses.NewResponse(nil, "Error", "Method not allowed")
		return c.JSON(http.StatusMethodNotAllowed, response)
	}

	token := c.Request().Header.Get("Authorization")

	email, err := authorization.GetEmailFromJWT(token)
	data := models2.UpdatePasswordForm{}

	err = c.Bind(&data)
	if err != nil {
		response := responses.NewResponse(nil, "Error", "Request bad structured")
		return c.JSON(http.StatusBadRequest, response)
	}

	err = H.Services.UpdatePassword(email, data)
	fmt.Println(err)
	if err != nil {
		response := responses.NewResponse(err, "Error", "Error While updating password")
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := responses.NewResponse(data, "Success", "Password updated Succesfully")
	return c.JSON(http.StatusOK, response)
}
