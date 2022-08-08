package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"grpcService/module/entities"
	"grpcService/support/db"
)

func models() map[string]interface{} {
	return map[string]interface{}{
		"acl":          new(entities.Acl),
		"tags":         new(entities.Tag),
		"users":        new(entities.User),
		"user_tags":    new(entities.UserTags),
		"user_acls":    new(entities.UserAcls),
		"songs":        new(entities.Song),
		"transactions": new(entities.Transaction),
	}
}

func instantiate(logger *logrus.Logger) (*db.Migrator, error) {
	conn, err := db.NewConnection()
	if err != nil {
		return nil, err
	}

	instance := db.NewInstance(conn, logger)

	return db.NewMigrator(logger, instance), nil
}

func migrateAll(
	ctx context.Context,
	logger *logrus.Logger,
	migrator *db.Migrator,
	models map[string]interface{},
) error {
	for f, v := range models {
		if err := migrator.MigrateOne(ctx, v, f); err != nil {
			logger.Error("Failed to migrate: ", err)
			return err
		}
	}

	return nil
}

func dropAll(
	ctx context.Context,
	logger *logrus.Logger,
	migrator *db.Migrator,
	models map[string]interface{},
) error {
	for f, v := range models {
		if err := migrator.DropTable(ctx, v, f); err != nil {
			logger.Error("Failed to drop table: ", err)
			return err
		}
	}

	return nil
}

func main() {
	logger := logrus.New()
	ctx := context.Background()

	args := map[string]int64{
		"migrate": 1,
		"cleanup": 2,
	}

	op, ok := args[os.Args[1]]
	if !ok {
		logger.Fatal("Operation not supported. Exiting.")
	}

	migrator, err := instantiate(logger)
	if err != nil {
		logger.Fatal("Failed to instantiate migrator: ", err)
	}

	switch op {
	case 1:
		migrateAll(ctx, logger, migrator, models())
	case 2:
		dropAll(ctx, logger, migrator, models())
	}
}
