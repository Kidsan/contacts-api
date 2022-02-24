package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"

	pb "github.com/kidsan/contacts-api/protobuffer"
)

func (s *HTTPServer) buildContactRouter() func(router chi.Router) {
	collector := adapters.NewAdapterCollector("contacts_api")
	contactRepository := adapters.NewContactRepository(s.connection, collector)
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(s.logger, contactService)

	return func(router chi.Router) {
		router.Get("/", contactHTTP.GetAllHandler())
		router.Post("/", contactHTTP.PostHandler())
		router.With(contactHTTP.CheckIDParam()).Get("/{id}", contactHTTP.FindHandler())
	}
}

func (g *GRPCServer) buildContactServer() pb.ContactsServer {
	collector := adapters.NewAdapterCollector("contacts_api")
	contactRepository := adapters.NewContactRepository(g.connection, collector)
	contactService := domain.NewContactService(contactRepository)
	contactGRPC := ports.NewContactGRPCHandler(g.logger, contactService)

	return *contactGRPC
}
