package Services

import (
	"errors"
)

func (S *Service) DeletePost(title, email string) error {
	if title == "" || email == "" {
		return errors.New("please provide a valid searchTitle or valid searchEmail")
	}

	err := S.UseCase.DeletePostByTitleAndEmail(title, email)
	if err != nil {
		return errors.New("somthing went wrong while deleting post")
	}

	return nil
}
