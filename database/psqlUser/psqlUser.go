package psqlUser

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/database"
	"github.com/TeenBanner/Inventory_system/models"
	"log"
)

type PsqlUser struct {
	db *sql.DB
}

func NewPsqlUser(db *sql.DB) *PsqlUser {
	return &PsqlUser{}
}

func (u *PsqlUser) CreateUser(user models.User) error {
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
