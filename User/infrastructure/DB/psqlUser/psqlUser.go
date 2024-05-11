package psqlUser

import (
	"database/sql"
	"errors"
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/TeenBanner/Inventory_system/pkg/database"
	"github.com/google/uuid"
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
func (u *UserStorage) CreateUser(user models2.User) error {
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
func (u *UserStorage) GetUser(id uuid.UUID) (models2.User, error) {
	stmt, err := u.db.Prepare(SqlGetUser)
	if err != nil {
		return models2.User{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
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

func (u *UserStorage) GetUserPosts(id uuid.UUID) ([]model.Post, error) {
	stmt, err := u.db.Prepare(SqlGetUserPosts)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	posts := []model.Post{}
	for rows.Next() {
		post := model.Post{}

		nullTime := sql.NullTime{}
		err := rows.Scan(&post.ID, &post.OwnerId, &post.Title, &post.Body, &post.CreatedAt, &nullTime)
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

// AdminMethods
