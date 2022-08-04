package entities

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// User a db object.
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   uuid.UUID `bun:"id" json:"id" yaml:"id"`
	Name string    `bun:"name" json:"name" yaml:"name"`
}
