package Services

import (
	"errors"
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
	"time"
)

func (S *Service) CreatePost(email string, post model2.CreatePost) (model2.Post, error) {
	if email == "" {
		return model2.Post{}, errors.New("please provide a valid email")
	}
	IdPost, _ := uuid.NewUUID()

	now := time.Now()
	ModelPost := model2.Post{
		ID:         IdPost,
		Title:      post.Title,
		Body:       post.Body,
		OwnerEmail: email,
		CreatedAt:  now,
	}

	err := S.UseCase.CreatePost(email, ModelPost)
	if err != nil {
		return ModelPost, err
	}

	return ModelPost, nil
}
