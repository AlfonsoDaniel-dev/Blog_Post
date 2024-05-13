package UserDomain

import (
	"errors"
	"fmt"
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UserStorage
}

func NewUser(storage UserStorage) *User {
	return &User{
		UserStorage: storage,
	}
}

func (u *User) Create(user model2.User) error {
	user.CreatedAt = time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.IsAdmin = false

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = u.UserStorage.CreateUser(user)
	if err != nil {
		return err
	}

	fmt.Println("User created")

	return nil
}

func (u *User) GetUser(email string) (model2.User, error) {
	user, err := u.GetUser(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) GetUserByName(name string) (model2.User, error) {
	user, err := u.GetUserByName(name)

	if err != nil {
		return model2.User{}, err
	}

	return user, nil
}

func (u *User) UpdateEmail(ActualEmail, NewEmail string) error {
	if NewEmail == ActualEmail {
		return errors.New("user email can't be equal")
	}

	if NewEmail == "" {
		return errors.New("user email can't be empty")
	}

	err := u.UpdateUserEmail(ActualEmail, NewEmail)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUserName(email, NewName string) error {
	if email == "" {
		return errors.New("user email can't be empty")
	}

	if NewName == "" {
		return errors.New("user name can't be empty")
	}

	err := u.UpdateUserName(email, NewName)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUsers admin functiond
func (u *User) AdminGetAllUsers() ([]model2.User, error) {
	users, err := u.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
