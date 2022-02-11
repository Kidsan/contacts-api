package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	Database Database
}

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type Database struct {
	Database        string
	Driver          string
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	Schema          string
	ApplicationName string
	MigrationPath   string
}

func (d *Database) DSN() string {
	const (
		dsnPattern = "host=%v port=%v user=%v password=%v dbname=%v sslmode=%v " +
			"application_name=%v search_path=%v"
		disableSSLMode = "disable"
	)

	return fmt.Sprintf(dsnPattern, d.Host, d.Port, d.User, d.Password, d.DBName, disableSSLMode, d.ApplicationName, d.Schema)
}

func Read(prefix string) (Configuration, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// viper.BindEnv("id")
	// viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return Configuration{}, fmt.Errorf("unable to decode into struct: %w", err)
	}
	fmt.Printf("%v", config)
	return Configuration{}, nil
}
