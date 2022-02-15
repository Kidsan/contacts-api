package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	contactsapi "github.com/kidsan/contacts-api"
	"github.com/kidsan/contacts-api/contact"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Server struct {
	logger *zap.Logger
	router *chi.Mux
	config contactsapi.ServerConfig
}

func NewServer(config contactsapi.Config, logger *zap.Logger) *Server {
	r := chi.NewRouter()
	s := &Server{
		config: config.Server,
		logger: logger,
		router: r,
	}
	connection, _ := s.openDBConnection(config.Database.DSN(), config.Database.Database)
	contact.RegisterContactRoutes(logger, r, connection)

	return &Server{
		config: config.Server,
		logger: logger,
		router: r,
	}
}

func (s *Server) Start() {
	s.logger.Info(fmt.Sprintf("Application listening on port %d", s.config.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), s.router)
}

func (s *Server) openDBConnection(dsn, databaseName string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "contacts_api.",
			SingularTable: false,
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("api: could not open database: %w", err)
	}

	// db, err := connection.DB()
	// if err != nil {
	// 	return nil, fmt.Errorf("api: could not get database: %w", err)
	// }

	return connection, nil
}
