package psqlUser

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/database"
	"github.com/TeenBanner/Inventory_system/models"
	"log"
)

// UserStorage it's used for interact with DB
type UserStorage struct {
	db *sql.DB
}

// NewUserStorage contructure for UserStorage
func NewPsqlUser(db *sql.DB) *UserStorage {
	return &UserStorage{}
}

// User methods
func (u *UserStorage) CreateUser(user models.User) error {
	stmt, err := u.db.Prepare(SqlCreateUserQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	UserNullTime := database.TimeToNull(user.UpdatedAt)

	_, err = stmt.Exec(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		UserNullTime,
	)

	if err != nil {
		return err
	}

	log.Println("Usuario creado")
	return nil
}

// AdminMethods

// AdminGetUser get's info from a user
func (u *UserStorage) AdminGetUser(id int) models.User {
	stmt, err := u.db.Prepare()
}
