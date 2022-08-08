package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// UserAcls a pivot declarator.
type UserAcls struct {
	bun.BaseModel `bun:"table:user_acls,alias:u_acl"`

	UserID uuid.UUID `bun:"user_id,pk,notnull" json:"user_id" yaml:"user_id"`
	User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	AclID  uuid.UUID `bun:"acl_id,pk,notnull" json:"acl_id" yaml:"acl_id"`
	Acl    *Acl      `bun:"rel:belongs-to,join:acl_id=id" json:"acl" yaml:"acl"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
