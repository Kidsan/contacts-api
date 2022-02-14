package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	contactsapi "github.com/kidsan/contacts-api"
	"github.com/kidsan/contacts-api/contact"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	router *chi.Mux
	config contactsapi.ServerConfig
}

func NewServer(config contactsapi.Config, logger *zap.Logger) *Server {
	r := chi.NewRouter()
	contact.RegisterContactRoutes(logger, r)

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
