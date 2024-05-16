package Domain

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func (U *User) AreEqual(email, passwordToVerify string) (bool, error) {
	if passwordToVerify == "" || email == "" {
		return false, errors.New("Please provide a valid email or password")
	}
	emailFromDb, err := U.UserStorage.PsqlLoginGetEmail(email)
	if err != nil {
		return false, err
	}

	if emailFromDb != email {
		return false, errors.New("email does not match or not exists")
	}
	passwordFromDb, err := U.UserStorage.PsqlLoginGetPassword(email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordFromDb), []byte(passwordToVerify))
	if err != nil {
		return false, err
	}

	return true, nil
}
