package UserDomain

import (
	"fmt"
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UserStorage
}

type Admin struct {
	AdminStorage
}

func NewUser(storage UserStorage) *User {
	return &User{
		UserStorage: storage,
	}
}

func (u *User) CreateUser(user model2.User) error {
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

func (U *User) GetUserByEmail(email string) (model2.User, error) {
	user, err := U.GetUser(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetAllUsers admin function
func (U *Admin) AdminGetAllUsers() ([]model2.User, error) {
	users, err := U.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
