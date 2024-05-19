package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
)

func (S *Service) UpdatePostTitle(email, postForm models2.UpdatePost) error {
	if postForm.Title == "" {
		return errors.New("please provide new title for the post")
	}

	oldPostId := S.UseCase.FindPostId(email, postForm)
}
