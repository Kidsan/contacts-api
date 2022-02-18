package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"

	pb "github.com/kidsan/contacts-api/protobuffer"
)

func (s *HTTPServer) BuildContactRouter() func(router chi.Router) {
	contactRepository := adapters.NewContactRepository(s.connection)
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(s.logger, contactService)

	return contactHTTP.Init()
}

func (g *GRPCServer) BuildContactServer() pb.ContactsServer {
	contactRepository := adapters.NewContactRepository(g.connection)
	contactService := domain.NewContactService(contactRepository)
	contactGRPC := ports.NewContactGRPCHandler(g.logger, contactService)

	return *contactGRPC
}
