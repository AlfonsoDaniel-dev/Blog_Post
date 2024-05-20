package UserApp

import (
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

type UseCase interface {
	CreateUser(user model2.User) error
	GetUserByEmail(email string) (model2.User, error)
	GetUserByName(name string) (model2.User, error)

	UpdateUserEmail(ActualEmail, NewEmail string) error
	UpdateUserName(email, NewName string) error
	UpdateUserPassword(email, NewPassword string) error

	AreEqual(email, passwordToVerify string) (bool, error)

	FindEmailByName(name string) (string, error)

	CreatePost(email string, post model2.Post) error
	GetUserPosts(email string) ([]model2.Post, error)
	FindPostId(searchEmail, searchTitle string) (uuid.UUID, error)
	FindPostTitle(email string) (string, error)
	GetPostsByEmail(name string) ([]model2.Post, error)
	FindPostsByTitle(title string) ([]model2.Post, error)
	FindPostById(postId uuid.UUID) (model2.Post, error)

	UpdatePostTitle(title, email string) error
	UpdatePostBody(body, email string) error

	DeletePostByTitleAndEmail(title, email string) error
	DeleteAccount(email string) error

	GetAllUsers() ([]model2.User, error)
}
