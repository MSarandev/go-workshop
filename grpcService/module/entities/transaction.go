package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Transaction a db object.
type Transaction struct {
	bun.BaseModel `bun:"table:transactions,alias:t"`

	ID     uuid.UUID `bun:"id,pk,notnull" json:"id" yaml:"id"`
	Type   string    `bun:"type,nullzero" json:"type" yaml:"type"`
	UserID uuid.UUID `bun:"user_id,pk,notnull" json:"user_id" yaml:"user_id"`
	User   *User     `bun:"rel:belongs-to,join:user_id=id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// Identity returns the identifier.
func (t *Transaction) Identity() *uuid.UUID {
	return &t.ID
}
