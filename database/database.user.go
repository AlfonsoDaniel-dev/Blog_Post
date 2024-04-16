package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TeenBanner/Inventory_system/models"
	"github.com/google/uuid"
)

const sqlCreateUser = `
INSERT INTO USERS(id, name, email, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)
`

const sqlGetAllUsers = `
	SELECT * FROM USERS
`

const sqlGetUserById = `SELECT (id, name, email, created_at, updated_at) FROM USERS WHERE id = $1`

const sqlUpdateUserName = `UPDATE user SET name= $1 WHERE id = $2`

const sqlUserUpdateEmail = `UPDATE user SET email =$1 WHERE id = $2`

const sqlUserUpdatePassword = `UPDATE user SER password = $1 WHERE id = $2`

const sqlDeleteUser = `DELETE * FROM USERS WHERE id = $1`

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
	if err != nil {
		return err
	}

	RowsAff, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Registro insertado: %v, filas affectadas %v", rows, RowsAff)

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

func (u *psqlUser) GetUserById(id int) (models.User, error) {
	user := models.User{}

	stmt, err := u.DB.Prepare(sqlGetUserById)
	if err != nil {
		return models.User{}, err
	}

	row := stmt.QueryRow(id)

	_ = row.Scan(&user.ID, &user.Name, &user.Email, &user.Created_at, &user.Updated_at)

	fmt.Println("Usuario obtenido por su id")

	return user, nil
}

func (u *psqlUser) UpdateUserName(id int, user models.User) error {
	stmt, err := u.DB.Prepare(sqlUpdateUserName)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, id)
	if err != nil {
		return err
	}

	fmt.Println("Se Actualizo el nombre del usuario")

	return nil
}

func (u *psqlUser) UpdateUserEmail(id int, user models.User) error {
	stmt, err := u.DB.Prepare(sqlUserUpdateEmail)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.Email, id)
	if err != nil {
		return err
	}

	fmt.Println("Se actualizo el Email del usuario")

	return nil
}

func (u *psqlUser) UpdatePassword(password uuid.UUID, id int) error {
	stmt, err := u.DB.Prepare(sqlUserUpdatePassword)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(password, id)
	if err != nil {
		return err
	}

	fmt.Println("Se actualizo el password de un usuario")

	return nil
}

func (u *psqlUser) DeleteUser(id int) error {
	stmt, err := u.DB.Prepare(sqlDeleteUser)
	if err != nil {
		return err
	}

	defer stmt.Close()

	rows, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows_affected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Usuario borrado. Registros afectados: %v", rows_affected)

	return nil
}
