package Services

import (
	"errors"
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

func (S *Service) GetByEmail(Email string) (models2.User, error) {
	if Email == "" {
		return models2.User{}, errors.New("search email cannot be nil")
	}

	user, err := S.UseCase.GetUserByEmail(Email)
	if err != nil {
		return models2.User{}, err
	}

	user.Password = ""
	user.Posts = nil
	user.ID = uuid.Nil

	return user, nil
}

func (S *Service) GetByName(name string) (models2.User, error) {
	if name == "" {
		return models2.User{}, errors.New("search name cannot be nil")
	}
	user, err := S.UseCase.GetUserByName(name)
	if err != nil {
		return models2.User{}, err
	}

	user.Password = ""

	return user, nil
}

func (S *Service) GetPostsFromName(name string) ([]model.Post, error) {
	if name == "" {
		return nil, errors.New("search name cannot be nil")
	}

	posts, err := S.UseCase.GetPostsByName(name)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
