package Services

import (
	"errors"
)

func (S *Service) DeletePost(title, email string) error {
	if title == "" || email == "" {
		return errors.New("please provide a valid searchTitle or valid searchEmail")
	}

	postid, err := S.UseCase.FindPostId(email, title)
	if err != nil {
		return errors.New("post id not found")
	}

	_, err = S.UseCase.FindPostById(postid)
	if err != nil {
		return errors.New("post not found")
	}

	err = S.UseCase.DeletePostByTitleAndEmail(title, email)
	if err != nil {
		return errors.New("somthing went wrong while deleting post")
	}

	return nil
}
