package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) UpdateName(email string, form models2.UpdateNameForm) error {
	if form.NewName == "" || email == "" {
		return errors.New("please provide a valid new name or email")
	}

	ok, err := S.UseCase.AreEqual(email, form.Password)
	if err != nil || !ok {
		return errors.New("identification failed")
	}

	err = S.UseCase.UpdateUserName(email, form.NewName)
	if err != nil {
		return err
	}

	return err
}
