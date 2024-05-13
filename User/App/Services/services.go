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

type Services interface {
	Register(user models2.User) (models2.User, error)

	GetByEmail(email string) (models2.User, error)
	GetByName(name string) (models2.User, error)
	GetPostsFromName(name string) ([]model.Post, error)

	UpdateEmail(ActualEmail, NewEmail string) error
	UpdateName(email, NewName string) error
}
