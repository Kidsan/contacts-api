package main

import (
	"fmt"

	contactsapi "github.com/kidsan/contacts-api"
	"github.com/kidsan/contacts-api/config"
	"github.com/kidsan/contacts-api/logger"
	"github.com/kidsan/contacts-api/migration"
	"github.com/kidsan/contacts-api/server"
)

func main() {
	logger := logger.NewLogger()
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	if err := runMigration(config.Database); err != nil {
		panic(err)
	}

	server := server.NewServer(config, logger)

	server.Start()
}

func runMigration(config contactsapi.DatabaseConfig) error {
	migration, err := migration.NewMigration(config)
	if err != nil {
		return err
	}
	defer migration.Close()

	return migration.Up()
}
