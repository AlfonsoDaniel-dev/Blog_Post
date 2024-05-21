package Services

import (
	"errors"
	"fmt"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
	"strings"
)

func (S *Service) GetUserByEmail(email string) (models2.UserDTO, error) {
	if email == "" {
		return models2.UserDTO{}, errors.New("search email cannot be nil")
	}

	user, err := S.UseCase.GetUserByEmail(email)
	if err != nil {
		return models2.UserDTO{}, err
	}

	userToSend := models2.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return userToSend, nil
}

func (S *Service) GetUserByName(name string) (models2.UserDTO, error) {
	if name == "" {
		return models2.UserDTO{}, errors.New("search name cannot be nil")
	}

	name = strings.ReplaceAll(name, "_", " ")

	user, err := S.UseCase.GetUserByName(name)
	if err != nil {
		return models2.UserDTO{}, err
	}

	UserToSend := models2.UserDTO{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return UserToSend, nil
}

func (S *Service) GetAllUsers() ([]models2.UserDTO, error) {
	users, err := S.UseCase.GetAllUsers()
	usersToSend := []models2.UserDTO{}
	for i, user := range users {
		users[i].Name = strings.ReplaceAll(users[i].Name, "_", " ")
		userToSend := models2.UserDTO{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
		usersToSend = append(usersToSend, userToSend)
	}
	if err != nil {
		return nil, err
	}

	return usersToSend, nil
}

func (S *Service) GetAllPostsFromUserEmail(email string) ([]models2.Post, error) {
	if email == "" {
		return nil, errors.New("search email cannot be nil")
	}

	posts, err := S.UseCase.GetUserPosts(email)
	if err != nil {
		return nil, err
	}
	for i := range posts {
		title := posts[i].Title

		posts[i].Title = strings.ReplaceAll(title, "_", " ")
	}

	return posts, nil
}

func (S *Service) GetAllPostsByTitle(title string) ([]models2.Post, error) {
	if title == "" {
		return nil, errors.New("search title cannot be nil")
	}

	posts, err := S.UseCase.FindPostsByTitle(title)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(posts); i++ {
		title := posts[i].Title

		posts[i].Title = strings.ReplaceAll(title, "_", " ")
	}

	return posts, nil
}

func (S *Service) GetPostByTitleAndEmail(searchTitle, searchEmail string) (models2.Post, error) {
	if searchTitle == "" || searchEmail == "" {
		return models2.Post{}, errors.New("please provide a valid search title or email")
	}

	id, err := S.UseCase.FindPostId(searchEmail, searchTitle)
	if err != nil {
		return models2.Post{}, err
	}

	post, err := S.UseCase.FindPostById(id)
	if err != nil {
		return models2.Post{}, err
	}
	title := strings.ReplaceAll(post.Title, "_", " ")
	post.Title = title

	return post, nil
}

func (S *Service) GetAllPostsFromName(name string) ([]models2.Post, error) {
	if name == "" {
		return nil, errors.New("please provide a valid SearchName")
	}

	name = strings.ReplaceAll(name, " ", "_")
	email, err := S.UseCase.FindEmailByName(name)
	if err != nil {
		return nil, err
	}

	posts, err := S.UseCase.GetPostsByEmail(email)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(posts); i++ {
		title := posts[i].Title

		posts[i].Title = strings.ReplaceAll(title, "_", " ")
	}

	return posts, nil
}

func (S *Service) UserGetAllPosts() ([]models2.Post, error) {
	posts, err := S.UseCase.GetAllPosts()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("somthing went wrong while getting all posts")
	}

	for i := range posts {
		title := posts[i].Title

		posts[i].Title = strings.ReplaceAll(title, "_", " ")
	}

	return posts, nil
}
