package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

func (S Service) Register(user models2.User) (models2.User, error) {
	if user.ID == uuid.Nil {
		return models2.User{}, errors.New("user Id is Nil")
	}

	if user.Name == "" {
		return models2.User{}, errors.New("user Name cannot be empty")
	}

	if user.Password == "" {
		return models2.User{}, errors.New("user Password cannot be empty")
	}

	if user.Email == "" {
		return models2.User{}, errors.New("email cannot be empty")
	}

	user.IsAdmin = false

	err := S.UseCase.CreateUser(user)
	if err != nil {
		return models2.User{}, err

	}

	return user, nil
}
