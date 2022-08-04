package db

import (
	"log"

	"github.com/uptrace/bun"
)

type Instance struct {
	db     *bun.DB
	logger *log.Logger
}

// NewInstance - constructor for an Instance.
func NewInstance(db *bun.DB, logger *log.Logger) *Instance {
	return &Instance{
		db:     db,
		logger: logger,
	}
}

// Health - checks if a database connection is present.
func (i *Instance) Health() bool {
	if err := i.db.Ping(); err != nil {
		i.logger.Fatalf("database unreachable, err: %s", err.Error())
		return false
	}

	return true
}
