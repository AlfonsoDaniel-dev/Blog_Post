package psqlPost

import (
	"database/sql"
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	"github.com/TeenBanner/Inventory_system/pkg/database"
)

type psqlPost struct {
	db *sql.DB
}

func NewPsqlPost(DB *sql.DB) *psqlPost {
	return &psqlPost{
		db: DB,
	}
}

func (P *psqlPost) PsqlCreatePost(email string, post model.Post) error {
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

func (P *psqlPost) PsqlFindPostByTitle(title string) (model.Post, error) {
	stmt, err := P.db.Prepare(SqlFindPostByName)
	if err != nil {
		return model.Post{}, err
	}
	defer stmt.Close()
	post := model.Post{}

	row := stmt.QueryRow(title)
	nulltime := sql.NullTime{}

	err = row.Scan(&post.ID, &post.Title, &post.Body, &post.OwnerId, &post.CreatedAt, nulltime)
	post.UpdatedAt = nulltime.Time

	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

/*
func (P *psqlPost) PsqlGetPostsFromUser(email string) ([]model.Post, error) {

} */
