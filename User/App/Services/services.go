package Services

import (
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
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

type UserServices interface {
	Register(user models2.User) (models2.User, error)
	UpdateName(email, NewName string) error
	UpdateUserEmail(ActualEmail, ewEmail string) error
	GetUserByEmail(email string) (models2.User, error)
	GetUserByName(name string) (models2.User, error)
	GetPostsFromName(name string) ([]model.Post, error)
}
