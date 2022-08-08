package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// Acl control mechanism declarator.
type Acl struct {
	bun.BaseModel `bun:"table:acl,alias:acl"`

	ID        uuid.UUID `bun:"id,pk,notnull" json:"id" yaml:"id"`
	Name      string    `bun:"name,notnull" json:"name" yaml:"name"`
	RoleFlag  *AclFlag  `bun:"role_flag,notnull" json:"role_flag" yaml:"role_flag"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

// AclFlag fine-grain acl control flag declarator.
type AclFlag string

var (
	AclCanUpload AclFlag = "CAN_UPLOAD"
	AclCanBuy    AclFlag = "CAN_BUY"
	AclAdmin     AclFlag = "ADMIN"
	AclBanned    AclFlag = "BANNED"
	AclLocked    AclFlag = "LOCKED"
	AclFraud     AclFlag = "FRAUD"
)

// Identity returns the identifier.
func (a *Acl) Identity() *uuid.UUID {
	return &a.ID
}
