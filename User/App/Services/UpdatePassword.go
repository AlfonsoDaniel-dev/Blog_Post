package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) UpdatePassword(email string, form models2.UpdatePasswordForm) error {
	if form.OldPassword == form.NewPassword {
		return errors.New("passwords cannot be the same")
	}

	ok, err := S.UseCase.AreEqual(email, form.OldPassword)
	if err != nil || !ok {
		return errors.New("Email or password does not match")
	}

	err = S.UseCase.UpdateUserPassword(email, form.NewPassword)
	if err != nil {
		return err
	}
	return nil
}
