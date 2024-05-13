package Services

import (
	"errors"
)

func (S Service) UpdateName(email, NewName string) error {
	if email == "" {
		return errors.New("email cannot be nil")
	}

	if NewName == "" {
		return errors.New("newName cannot be nil")
	}

	err := S.UseCase.UpdateUserName(email, NewName)
	if err != nil {
		return err
	}

	return nil
}

func (S Service) UpdateEmail(ActualEmail, NewEmail string) error {
	if ActualEmail == "" {
		return errors.New("ActualEmail cannot be nil")
	}

	if NewEmail == "" {
		return errors.New("NewEmail cannot be nil")
	}

	err := S.UseCase.UpdateUserEmail(ActualEmail, NewEmail)
	if err != nil {
		return err
	}

	return nil
}
