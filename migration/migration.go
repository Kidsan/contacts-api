package migration

import (
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	contactsapi "github.com/kidsan/contacts-api"
	_ "github.com/lib/pq"
)

type Migration struct {
	migrationPath string
	dbName        string
	schema        string
	connection    *sql.DB
}

func NewMigration(database contactsapi.DatabaseConfig) (*Migration, error) {
	connection, err := sql.Open(database.Driver, database.DSN())
	if err != nil {
		return nil, fmt.Errorf("database: could not open database connection: %w", err)
	}

	return &Migration{
		migrationPath: database.MigrationPath,
		dbName:        database.DBName,
		schema:        database.Schema,
		connection:    connection,
	}, nil
}

func (m *Migration) Up() error {
	_, err := m.connection.Exec(createSchema)
	if err != nil {
		return fmt.Errorf("database: could not create schema: %w", err)
	}

	_, err = m.connection.Exec(createTable)
	if err != nil {
		return fmt.Errorf("database: could not create table: %w", err)
	}

	return nil
}

func (m *Migration) Close() error {
	return m.connection.Close()
}
