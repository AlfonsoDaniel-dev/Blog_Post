package Services

import (
	"errors"
	"fmt"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
)

func (S Service) Register(user models2.Register) error {
	if user.Name == "" {
		return errors.New("user Name cannot be empty")
	}

	if user.Password == "" {
		return errors.New("user Password cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	Id := uuid.New()

	userRecord := models2.User{
		ID:       Id,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	err := S.UseCase.CreateUser(userRecord)
	if err != nil {
		return err
	}

	fmt.Println("registro fallo en app")

	return nil
}
