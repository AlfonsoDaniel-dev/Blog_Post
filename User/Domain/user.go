package UserDomain

import (
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

// userStorage use user methods on DB
type UserStorage interface {
	Createuser(user model.User) error
	GetUser(id uuid.UUID) (model.User, error)
}

// AdminStorage extends userStorge for Admin type functions
type AdminStorage interface {
	UserStorage
	GetAllUsers() ([]model.User, error)
}
