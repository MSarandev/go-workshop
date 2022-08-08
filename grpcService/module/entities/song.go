package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Song a db object.
type Song struct {
	bun.BaseModel `bun:"table:songs,alias:t"`

	ID     uuid.UUID `bun:"id,pk,notnull" json:"id" yaml:"id"`
	Name   string    `bun:"name,notnull" json:"name" yaml:"name"`
	Artist string    `bun:"artist,notnull" json:"artist" yaml:"artist"`
	Length int64     `bun:"length,notnull,default:0" json:"length" yaml:"length"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// Identity returns the identifier.
func (s *Song) Identity() *uuid.UUID {
	return &s.ID
}
