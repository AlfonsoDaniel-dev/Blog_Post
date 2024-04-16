package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TeenBanner/Inventory_system/models"
)

const sqlCreateUser = `
INSERT INTO USERS(id, name, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)
`

const sqlGetAllUsers = `
	SELECT * FROM USERS
`

type psqlUser struct {
	DB *sql.DB
}

func NewPsqlUser(db *sql.DB) psqlUser {
	return psqlUser{
		DB: db,
	}
}

func (u *psqlUser) CreateUser(user models.User) error {
	stmt, err := u.DB.Prepare(sqlCreateUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	user_creation := time.Now()

	rows, err := stmt.Exec(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user_creation,
		TimeToNull(user.Updated_at),
	)

	RowsAff, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Registro insertado: %v, filas affectadas %s", rows, RowsAff)

	return nil
}

func (u *psqlUser) GetAllUsers() ([]models.User, error) {
	stmt, err := u.DB.Prepare(sqlGetAllUsers)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	users := []models.User{}

	for rows.Next() {
		user := models.User{}

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Created_at,
			&user.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Registros de usuarios obtenidos")

	return users, nil
}
