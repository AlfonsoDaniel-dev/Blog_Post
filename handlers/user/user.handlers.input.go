package user

import "github.com/TeenBanner/Inventory_system/models"

type input interface {
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
}
