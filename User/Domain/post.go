package Domain

import (
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
)

type PostStorage interface {
	PsqlCreatePost(post model2.Post) error
	PsqlGetPostFromUser(email string) (model2.Post, error)
	PsqlGetPosts() ([]model2.Post, error)
}
