package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func generateDsn() (string, error) {
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", errors.New("failed to load env param")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return "", errors.New("failed to load env param")
	}

	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return "", errors.New("failed to load env param")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		return "", errors.New("failed to load env param")
	}

	database, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return "", errors.New("failed to load env param")
	}

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, database), nil
}

// NewConnection - initialises a db connection.
func NewConnection() (*bun.DB, error) {
	dsn, err := generateDsn()
	if err != nil {
		return nil, err
	}

	conn := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(conn, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db, nil
}
