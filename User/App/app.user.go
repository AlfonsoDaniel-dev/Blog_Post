package UserApp

import models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"

type UseCase interface {
	Create(user models2.User) error
	GetUser(email string) (models2.User, error)
	UpdateEmail(ActualEmail, NewEmail string) error
	UpdateUserName(email, NewName string) error

	AdminGetAllUsers() ([]models2.User, error)
}
