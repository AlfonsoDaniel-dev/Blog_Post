package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) UpdateName(form models2.UpdateNameForm) error {
	if form.NewName == "" || form.Email == "" {
		return errors.New("please provide a valid new name or email")
	}

	ok, err := S.UseCase.AreEqual(form.Email, form.Password)
	if err != nil || !ok {
		return errors.New("identification failed")
	}

	err = S.UseCase.UpdateUserName(form.Email, form.NewName)
	if err != nil {
		return err
	}

	return err
}
