package UserDomain

import (
	"github.com/TeenBanner/Inventory_system/models"
	"github.com/google/uuid"
)

// userStorage use user methods on DB
type UserStorage interface {
	Createuser(user models.User) error
	GetUser(id uuid.UUID) (models.User, error)
}

// AdminStorage extends userStorge for Admin type functions
type AdminStorage interface {
	UserStorage
	GetAllUsers() ([]models.User, error)
}
