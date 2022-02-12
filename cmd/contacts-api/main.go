package main

import (
	"fmt"

	"github.com/kidsan/contacts-api/config"
	"github.com/kidsan/contacts-api/logger"
	"github.com/kidsan/contacts-api/server"
)

func main() {
	logger := logger.NewLogger()
	config, err := config.Read("")
	if err != nil {
		fmt.Println(err)
	}

	server := server.NewServer(config, logger)

	server.Start()
}
