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

type UserServices interface {
	Register(user models2.User) (models2.User, error)
	UpdateName(email, NewName string) error
	UpdateUserEmail(ActualEmail, ewEmail string) error
	GetUserByEmail(email string) (models2.User, error)
}
