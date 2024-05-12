package UserDomain

import (
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
)

// userStorage use user methods on DB
type UserStorage interface {
	CreateUser(user model2.User) error
	GetUser(email string) (model2.User, error)
	GetUserPosts(email string) ([]model.Post, error)
	UpdateUserName(email string) error
}

// AdminStorage extends userStorge for Admin type functions
type AdminStorage interface {
	UserStorage
	GetAllUsers() ([]model2.User, error)
}
