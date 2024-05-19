package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) UpdatePostTitle(email string, postform models2.UpdatePost) error {
	if postform.Title == "" {
		return errors.New("email and postform are required")
	}

	OldTitle, err := S.UseCase.FindPostTitle(email)
	if err != nil {
		return err
	}

	if OldTitle == postform.Title {
		return errors.New("titles cannot be the same title")
	}

	err = S.UseCase.UpdatePostTitle(email, postform.Title)
	if err != nil {
		return errors.New("failed to update post title")
	}

	return nil
}

func (S *Service) UpdatePostBody(email string, postform models2.UpdatePost) error {
	if postform.Body == "" || postform.Title == "" {
		return errors.New("postform cannot be empty")
	}

	postId, err := S.UseCase.FindPostId(email, postform.Title)
	if err != nil {
		return errors.New("failed to find post ID")
	}

	oldPost, err := S.UseCase.FindPostById(postId)
	if err != nil {
		return errors.New("failed to find The Old post")
	}

	if oldPost.Body == postform.Body {
		return errors.New("post body cannot be the same body")
	}

	err = S.UseCase.UpdatePostBody(postform.Body, email)
	if err != nil {
		return errors.New("failed to update post body")
	}

	return nil
}
