package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	contactsapi "github.com/kidsan/contacts-api"
	"github.com/kidsan/contacts-api/contact"
	"go.uber.org/zap"
)

type Server struct {
	logger *zap.Logger
	router *chi.Mux
}

func NewServer(config contactsapi.Config, logger *zap.Logger) *Server {
	r := chi.NewRouter()
	contact.RegisterContactRoutes(logger, r)
	return &Server{
		logger: logger,
		router: r,
	}
}

func (s *Server) Start() {
	s.logger.Info("Application listening on port 3000")
	http.ListenAndServe(":3000", s.router)
}
