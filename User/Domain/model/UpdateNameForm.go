package model2

type UpdateNameForm struct {
	Login
	NewName string `json:"name"`
}
