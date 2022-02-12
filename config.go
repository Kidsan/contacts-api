package contactsapi

import "fmt"

type Config struct {
	Database Database `mapstructure:"DATABASE"`
}

type Database struct {
	Database        string `mapstructure:"DATABASE"`
	Host            string `mapstructure:"HOST"`
	Port            int    `mapstructure:"PORT"`
	User            string `mapstructure:"USER"`
	Password        string `mapstructure:"PASS"`
	DBName          string `mapstructure:"NAME"`
	Schema          string `mapstructure:"SCHEMA"`
	ApplicationName string `mapstructure:"APP_NAME"`
	MigrationPath   string `mapstructure:"MIGRATION_PATH"`
}

func (d *Database) DSN() string {
	const (
		dsnPattern = "host=%v port=%v user=%v password=%v dbname=%v sslmode=%v " +
			"application_name=%v search_path=%v"
		disableSSLMode = "disable"
	)

	return fmt.Sprintf(dsnPattern, d.Host, d.Port, d.User, d.Password, d.DBName, disableSSLMode, d.ApplicationName, d.Schema)
}
