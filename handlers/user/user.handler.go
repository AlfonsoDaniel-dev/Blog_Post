package user

import (
	"errors"

	"github.com/TeenBanner/Inventory_system/helpers"
	"github.com/TeenBanner/Inventory_system/models"
)

type User struct {
	db Useroutput
}

func NewUser(DB Useroutput) User {
	return User{
		db: DB,
	}
}

func (u *User) CreateUser(user models.User) error {
	password := user.Password

	user.Password = helpers.HashPassword(password)

	if user.Password == "" {
		return errors.New("password cannot be hash")
	}
	err := u.db.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) GetUser() ([]models.User, error) {
	users, err := u.db.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
