package main

import (
	"context"

	"awesomeProject1/module/entities"
	"awesomeProject1/support/db"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	conn, err := db.NewConnection()
	if err != nil {
		logger.Fatal(err)
	}

	instance := db.NewInstance(conn, logger)

	ok := instance.Health()

	logger.Log(logrus.InfoLevel, "DB conn: ", ok)

	user := entities.User{
		ID:   uuid.MustParse("1771424e-0000-0000-0000-a6df0d9ba203"),
		Name: "Test Test",
	}

	instance.MigrateOne(context.Background(), &user)
}
