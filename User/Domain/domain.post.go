package Domain

import (
	"errors"
	model2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"github.com/google/uuid"
	"time"
)

func (U *User) CreatePost(email string, post model2.Post) error {
	if email == "" {
		return errors.New("please provide an valid email")
	}

	post.CreatedAt = time.Now()
	post.OwnerEmail = email

	err := U.UserStorage.PsqlCreatePost(email, post)
	if err != nil {
		return err
	}

	return nil
}

func (U *User) GetUserPosts(name string) ([]model2.Post, error) {
	if name == "" {
		return nil, errors.New("please provide an valid email")
	}

	posts, err := U.UserStorage.PsqlGetUserPosts(name)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (U *User) GetPostsByName(name string) ([]model2.Post, error) {
	if name == "" {
		return nil, errors.New("search name can't be empty")
	}

	posts, err := U.UserStorage.PsqlGetUserPosts(name)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (U *User) FindPostsByTitle(title string) ([]model2.Post, error) {
	if title == "" {
		return nil, errors.New("search title can't be empty")
	}

	posts, err := U.UserStorage.PsqlGetUserPosts(title)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (U *User) FindPostById(id uuid.UUID) (model2.Post, error) {
	if id == uuid.Nil {
		return model2.Post{}, errors.New("please provide a valid uuid")
	}

	post, err := U.UserStorage.PsqlFindPostById(id)
	if err != nil {
		return model2.Post{}, err
	}

	return post, nil
}

func (U *User) FindPostId(searchTitle, searchEmail string) (uuid.UUID, error) {
	if searchEmail == "" || searchTitle == "" {
		return uuid.Nil, errors.New("please provide a valid search email or title")
	}

	post, err := U.UserStorage.PsqlFindPostId(searchTitle, searchEmail)
	if err != nil {
		return uuid.Nil, errors.New("post Id Does Not Exist.")
	}

	return post, nil
}
