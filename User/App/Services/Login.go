package Services

import (
	"errors"
	"fmt"
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/TeenBanner/Inventory_system/pkg/authorization"
)

func (S *Service) Login(data model2.Login) (string, error) {
	if data.Email == "" || data.Password == "" {
		return "", errors.New("please provide a valid email or password")
	}

	ok, err := S.UseCase.AreEqual(data.Email, data.Password)
	if err != nil || !ok {
		return "", errors.New("invalid email or password")
	}

	fmt.Println("la funcion fallo aqui")

	token, err := authorization.GenerateToken(data)
	if err != nil {
		return "", err
	}

	return token, nil
}
