package Services

import (
	"errors"
	models2 "github.com/TeenBanner/Inventory_system/User/Domain/model"
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

func (S *Service) DeleteAccount(form models2.DeleteAccountForm) error {
	if form.Password == "" || form.Email == "" {
		return errors.New("please provide a valid password or email for delete account")
	}

	ok, err := S.UseCase.AreEqual(form.Email, form.Password)
	if err != nil || !ok {
		return errors.New("email or password is required to delete account")
	}

	err = S.UseCase.DeleteAccount(form.Email)
	if err != nil {
		return errors.New("somthing went wrong while deleting account")
	}

	return nil
}
