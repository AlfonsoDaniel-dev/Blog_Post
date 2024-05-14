package psqlUser

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/TeenBanner/Inventory_system/pkg/database"
	"github.com/google/uuid"
	"time"
)

func (P *userStorage) PsqlCreatePost(email string, post model2.Post) error {
	stmt, err := P.db.Prepare(SqlCreatePost)
	if err != nil {
		return err
	}

	defer stmt.Close()

	PostNullTime := database.TimeToNull(post.UpdatedAt)

	_, err = stmt.Exec(post.ID, post.Title, post.Body, email, post.CreatedAt, PostNullTime)
	if err != nil {
		return err
	}

	return nil
}

func (u *userStorage) PsqlGetUserPosts(name string) ([]model2.Post, error) {
	stmt, err := u.db.Prepare(SqlGetUserPosts)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(name)

	posts := []model2.Post{}
	for rows.Next() {
		post := model2.Post{}

		nullTime := sql.NullTime{}
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.OwnerEmail, &post.CreatedAt, &nullTime)
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

func (P *userStorage) PsqlFindPostByTitle(title string) ([]model2.Post, error) {
	stmt, err := P.db.Prepare(SqlFindPostByName)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	posts := []model2.Post{}

	rows, err := stmt.Query(title)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		post := model2.Post{}

		nullTime := sql.NullTime{}
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.OwnerEmail, &post.CreatedAt, &nullTime)
		post.UpdatedAt = nullTime.Time
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)

	}

	return posts, nil
}

func (P *userStorage) PsqlFindPostsById(id uuid.UUID) ([]model2.Post, error) {
	stmt, err := P.db.Prepare(SqlGetPostsById)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	posts := []model2.Post{}
	for rows.Next() {
		post := model2.Post{}

		nulltime := sql.NullTime{}
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.OwnerEmail, &post.CreatedAt, &nulltime)
		post.UpdatedAt = nulltime.Time
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (P *userStorage) PsqlUpdatePostTitle(email string, title string) error {
	stmt, err := P.db.Prepare(SqlUpdatePostTitle)
	if err != nil {
		return err
	}

	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(title, now, email)
	if err != nil {
		return err
	}

	return nil
}

func (P *userStorage) PsqlUpdatePostBody(email, body string) error {
	stmt, err := P.db.Prepare(SqlUpdatePostBody)
	if err != nil {
		return err
	}

	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(body, now, email)
	if err != nil {
		return err
	}

	return nil
}

func (U *userStorage) PsqlFindPostId(searchTitle, searchEmail string) (uuid.UUID, error) {

}
