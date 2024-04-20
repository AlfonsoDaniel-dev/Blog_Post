package user

import "github.com/TeenBanner/Inventory_system/models"

type input interface {
	CreateUser(user models.User)
	GetAllUsers() ([]models.User, error)
}
