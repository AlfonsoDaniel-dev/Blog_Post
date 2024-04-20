package user

import (
	"github.com/TeenBanner/Inventory_system/models"
	"github.com/google/uuid"
)

type Useroutput interface {
	Create(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	UpdateUserName(id int, user models.User) error
	UpdatePassword(id int, password uuid.UUID) error
	DeleteUser(id int) error
	GetHashPassword(id int) (string, error)
}
