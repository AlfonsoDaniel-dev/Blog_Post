package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

func (S *Service) GetUserByEmail(Email string) (models2.User, error) {
	if Email == "" {
		return models2.User{}, errors.New("search email cannot be nil")
	}

	user, err := S.UseCase.GetUser(Email)
	if err != nil {
		return models2.User{}, err
	}

	user.Password = ""
	user.Posts = nil
	user.ID = uuid.Nil

	return user, nil
}

func UserByName(name string) (models2.User, error) {

}
