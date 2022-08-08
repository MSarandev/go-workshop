package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Tag a tag db model.
type Tag struct {
	bun.BaseModel `bun:"table:tags,alias:t"`

	ID   uuid.UUID `bun:"id,pk,notnull" json:"id" yaml:"id"`
	Name string    `bun:"name,notnull" json:"name" yaml:"name"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// Identity returns the identifier.
func (t *Tag) Identity() *uuid.UUID {
	return &t.ID
}
