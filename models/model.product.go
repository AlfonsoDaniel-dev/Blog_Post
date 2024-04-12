package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Created_At time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
