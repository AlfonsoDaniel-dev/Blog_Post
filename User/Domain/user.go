package Domain

import (
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

	PsqlCreatePost(email string, post model2.Post) error
	PsqlGetUserPosts(name string) ([]model2.Post, error)
	PsqlFindPostByTitle(title string) ([]model2.Post, error)

	PsqlLoginGetEmail(email string) (string, error)
	PsqlLoginGetPassword(email string) (string, error)
}
