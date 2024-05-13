package psqlUser

import (
	"database/sql"
	"errors"
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/TeenBanner/Inventory_system/pkg/database"
	"log"
)

// UserStorage it's used for interact with DB
type userStorage struct {
	db *sql.DB
}

// NewUserStorage contructure for UserStorage
func NewPsqlUser(DB *sql.DB) *userStorage {
	return &userStorage{
		db: DB,
	}
}

// User methods
func (u *userStorage) CreateUser(user models2.User) error {
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

// GetUser get info from a user
func (u *userStorage) GetUserByEmail(email string) (models2.User, error) {
	stmt, err := u.db.Prepare(SqlGetUser)
	if err != nil {
		return models2.User{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)
	if row == nil {
		return models2.User{}, errors.New("user does not exist")
	}
	user := models2.User{}
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models2.User{}, err
	}

	return user, nil
}

func (u *userStorage) GetUserPosts(name string) ([]model.Post, error) {
	stmt, err := u.db.Prepare(SqlGetUserPosts)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(name)

	posts := []model.Post{}
	for rows.Next() {
		post := model.Post{}

		nullTime := sql.NullTime{}
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.OwnerId, &post.CreatedAt, &nullTime)
		post.UpdatedAt = nullTime.Time
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (u *userStorage) UpdateUserName(email, name string) error {
	stmt, err := u.db.Prepare(SqlUpdateUserName)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(name, email)
	if err != nil {
		return err
	}
	return nil
}

func (u *userStorage) UpdateUserEmail(ActualEmail, NewEmail string) error {
	stmt, err := u.db.Prepare(SqlUpdateUserEmail)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(NewEmail, ActualEmail)
	if err != nil {
		return err
	}

	return nil
}

// AdminMethods

func (u *userStorage) GetAllUsers() ([]models2.User, error) {
	stmt, err := u.db.Prepare(SqlAdminGetAllUsers)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	users := []models2.User{}
	for rows.Next() {
		user := models2.User{}
		nullTime := sql.NullTime{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &nullTime)
		user.UpdatedAt = nullTime.Time
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}

	return users, nil
}

func (U *userStorage) GetUserByName(name string) (models2.User, error) {
	stmt, err := U.db.Prepare(SqlGetUserByName)
	if err != nil {
		return models2.User{}, err
	}

	user := models2.User{}

	row := stmt.QueryRow(name)

	if row == nil {
		return models2.User{}, err
	}

	nulltime := sql.NullTime{}

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &nulltime)
	if err != nil {
		return models2.User{}, err
	}

	user.Password = ""
	user.UpdatedAt = nulltime.Time

	return user, nil
}
