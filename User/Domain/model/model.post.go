package model2

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	OwnerEmail string    `json:"ownerId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
