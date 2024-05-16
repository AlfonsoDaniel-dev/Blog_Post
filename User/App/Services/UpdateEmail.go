package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) UpdateEmail(actualEmail string, form models2.UpdateEmailForm) error {
	if form.Email == "" || form.Password == "" {
		return errors.New("Email or Password is empty")
	}

	ok, err := S.UseCase.AreEqual(form.Email, form.Password)
	if err != nil || !ok {
		return errors.New("email or Password is wrong")
	}

	err = S.UseCase.UpdateUserEmail(actualEmail, form.Email)
	if err != nil {
		return err
	}

	return nil
}
