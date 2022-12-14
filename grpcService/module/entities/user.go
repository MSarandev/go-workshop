package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// User a db object.
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   uuid.UUID `bun:"id,pk,notnull" json:"id" yaml:"id"`
	Name string    `bun:"name,nullzero" json:"name" yaml:"name"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// Identity returns the identifier.
func (u *User) Identity() *uuid.UUID {
	return &u.ID
}
