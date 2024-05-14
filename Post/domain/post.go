package PostDomain

import "github.com/TeenBanner/Inventory_system/Post/domain/model"

type PostStorage interface {
	PsqlCreatePost(post model.Post) error
	PsqlGetPostFromUser(email string) (model.Post, error)
	PsqlGetPosts() ([]model.Post, error)
}
