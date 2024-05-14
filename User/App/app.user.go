package UserApp

import (
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

type UseCase interface {
	CreateUser(user models2.User) error
	GetUserByEmail(email string) (models2.User, error)
	GetUserByName(name string) (models2.User, error)
	UpdateUserEmail(ActualEmail, NewEmail string) error
	UpdateUserName(email, NewName string) error
	AreEqual(email, passwordToVerify string) (bool, error)

	GetPostsByName(name string) ([]models2.Post, error)

	GetAllUsers() ([]models2.User, error)
}
