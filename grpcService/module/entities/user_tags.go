package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// UserTags pivot declarator.
type UserTags struct {
	bun.BaseModel `bun:"table:user_tags,alias:uts"`

	UserID uuid.UUID `bun:"user_id,pk,notnull" json:"user_id" yaml:"user_id"`
	User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	TagID  uuid.UUID `bun:"tag_id,pk,notnull" json:"tag_id" yaml:"tag_id"`
	Tag    *Tag      `bun:"rel:belongs-to,join:tag_id=id"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
