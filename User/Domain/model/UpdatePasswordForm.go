package model2

type UpdatePasswordForm struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
