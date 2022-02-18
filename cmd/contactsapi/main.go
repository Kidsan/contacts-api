package main

import (
	"fmt"

	contactsapi "github.com/kidsan/contacts-api"
	"github.com/kidsan/contacts-api/config"
	"github.com/kidsan/contacts-api/http"
	"github.com/kidsan/contacts-api/logger"
	"github.com/kidsan/contacts-api/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

type PortListener interface {
	Start()
}

func main() {
	logger := logger.NewLogger()
	config, err := config.Read()
	if err != nil {
		panic(err)
	}

	if err := runMigration(config.Database); err != nil {
		panic(err)
	}

	dbConnection, err := openDBConnection(config.Database.DSN(), config.Database.Database)
	if err != nil {
		panic(err)
	}

	var server PortListener
	switch config.Server.Type {
	case "http":
		server = http.NewHTTPServer(config, logger, dbConnection)
	default:
		logger.Warn("No server type chosen, defaulting to http")
		server = http.NewHTTPServer(config, logger, dbConnection)
	}
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

func openDBConnection(dsn, databaseName string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})

	if err != nil {
		return nil, fmt.Errorf("api: could not open database: %w", err)
	}

	return connection, nil
}
