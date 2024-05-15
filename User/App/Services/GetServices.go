package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) GetByEmail(Email string) (models2.UserDTO, error) {
	if Email == "" {
		return models2.UserDTO{}, errors.New("search email cannot be nil")
	}

	user, err := S.UseCase.GetUserByEmail(Email)
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

func (S *Service) GetByName(name string) (models2.UserDTO, error) {
	if name == "" {
		return models2.UserDTO{}, errors.New("search name cannot be nil")
	}
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

func (S *Service) GetPostsFromName(name string) ([]models2.Post, error) {
	if name == "" {
		return nil, errors.New("search name cannot be nil")
	}

	posts, err := S.UseCase.GetPostsByName(name)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (S *Service) GetAllUsers() ([]models2.UserDTO, error) {
	users, err := S.UseCase.GetAllUsers()
	usersToSend := []models2.UserDTO{}
	for _, user := range users {
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

func (S *Service) GetAllPostsFromUser(email string) ([]models2.Post, error) {
	if email == "" {
		return nil, errors.New("search email cannot be nil")
	}

	posts, err := S.UseCase.GetUserPosts(email)
	if err != nil {
		return nil, err
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

	return post, nil
}
