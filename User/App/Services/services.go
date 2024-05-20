package Services

import (
	UserApp "github.com/TeenBanner/Inventory_system/User/App"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

type Service struct {
	UserApp.UseCase
}

func NewServices(useCase UserApp.UseCase) *Service {
	return &Service{
		useCase,
	}
}

type Services interface {
	Register(user models2.Register) error
	Login(user models2.Login) (string, error)

	UpdateEmail(actualEmail string, form models2.UpdateEmailForm) error
	UpdatePassword(email string, form models2.UpdatePasswordForm) error
	UpdateName(email string, form models2.UpdateNameForm) error

	UpdatePostTitle(email string, PostForm models2.UpdatePost) error
	UpdatePostBody(email string, Postform models2.UpdatePost) error
	DeletePost(title, email string) error

	CreatePost(email string, post models2.CreatePost) (models2.Post, error)

	GetUserByEmail(email string) (models2.UserDTO, error)
	GetUserByName(name string) (models2.UserDTO, error)
	GetAllUsers() ([]models2.UserDTO, error)
	GetAllPostsFromUserEmail(email string) ([]models2.Post, error)
	GetPostByTitleAndEmail(title string, email string) (models2.Post, error)
	GetAllPostsFromName(name string) ([]models2.Post, error)
}
