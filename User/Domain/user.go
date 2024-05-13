package UserDomain

import (
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
)

// userStorage use user methods on DB
type UserStorage interface {
	PsqlCreateUser(user model2.User) error
	PsqlGetUserByEmail(email string) (model2.User, error)
	PsqlGetUserByName(name string) (model2.User, error)
	PsqlUpdateUserName(email, name string) error
	PsqlUpdateUserEmail(ActualEmail, NewEmail string) error
	PsqlGetAllUsers() ([]model2.User, error)

	PsqlGetUserPosts(name string) ([]model.Post, error)
}
