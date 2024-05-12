package model2

import (
	"github.com/TeenBanner/Inventory_system/Post/domain/model"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Posts     []model.Post `json:"posts"`
	IsAdmin   bool         `json:"isAdmin"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
