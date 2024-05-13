package UserApp

import (
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

type UseCase interface {
	Create(user models2.User) error
	GetByEmail(email string) (models2.User, error)
	GetByName(name string) (models2.User, error)
	UpdateEmail(ActualEmail, NewEmail string) error
	UpdateUserName(email, NewName string) error

	GetPosts(name string) ([]model.Post, error)

	AdminGetAllUsers() ([]models2.User, error)
}
