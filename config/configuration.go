package config

import (
	"fmt"

	contactsapi "github.com/kidsan/contacts-api"
	"github.com/spf13/viper"
)

func Read(prefix string) (contactsapi.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return contactsapi.Config{}, err
	}

	var config contactsapi.Config
	if err := viper.Unmarshal(&config); err != nil {
		return contactsapi.Config{}, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return config, nil
}
