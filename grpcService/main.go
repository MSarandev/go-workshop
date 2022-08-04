package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"awesomeProject1/support/db"
)

func main() {
	var buf bytes.Buffer

	logger := log.New(&buf, "logger: ", log.Lmicroseconds)

	conn, err := db.NewConnection()
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	instance := db.NewInstance(conn, logger)

	ok := instance.Health()
	fmt.Println("db ok: ", ok)
}
