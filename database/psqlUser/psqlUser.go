package psqlUser

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/models"
)

type PsqlUser struct {
	db *sql.DB
}

func NewPsqlUser(db *sql.DB) *PsqlUser {
	return &PsqlUser{}
}

func (u *PsqlUser) CreateUser(user models.User) error {
	stmt, err := u.db.Prepare(CreateUserQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row, err := stmt.Exec(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)
}
