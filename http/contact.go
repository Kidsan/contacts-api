package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"
)

func (s *Server) BuildContactRouter() func(router chi.Router) {
	contactRepository := adapters.NewContactRepository(s.connection)
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(s.logger, contactService)

	return contactHTTP.Init()
}
