package migration

const (
	createSchema = `
		CREATE SCHEMA IF NOT EXISTS contacts_api;`

	createTable = `
	CREATE TABLE IF NOT EXISTS contacts_api.contacts (
		id uuid PRIMARY KEY,
		name varchar(40) NOT NULL
	);`
)
