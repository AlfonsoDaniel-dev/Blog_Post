package Domain

import (
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

// userStorage use user methods on DB
type UserStorage interface {
	PsqlCreateUser(user model2.User) error
	PsqlGetUserByEmail(email string) (model2.User, error)
	PsqlGetUserByName(name string) (model2.User, error)

	PsqlUpdateUserName(email, name string) error
	PsqlUpdateUserEmail(ActualEmail, NewEmail string) error
	PsqlUpdateUserPassword(email, password string) error

	PsqlGetAllUsers() ([]model2.User, error)
	PsqlFindUserEmailByName(name string) (string, error)

	PsqlGetUserName(email string) (string, error)

	PsqlCreatePost(email string, post model2.Post) error
	PsqlGetUserPosts(name string) ([]model2.Post, error)
	PsqlFindPostById(id uuid.UUID) (model2.Post, error)
	PsqlFindPostByTitle(title string) ([]model2.Post, error)

	PsqlUpdatePostBody(email, body string) error
	PsqlUpdatePostTitle(email, title string) error

	PsqlFindPostId(searchTitle, SearchEmail string) (uuid.UUID, error)
	PsqlFindPostTitle(SearchEmail string) (string, error)

	PsqlDeletePost(title, email string) error

	PsqlDeleteAccount(email string) error

	PsqlLoginGetEmail(email string) (string, error)
	PsqlLoginGetPassword(email string) (string, error)
}
