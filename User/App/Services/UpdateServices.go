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
