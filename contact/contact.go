package contact

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"
	"go.uber.org/zap"
)

func RegisterContactRoutes(logger *zap.Logger, r *chi.Mux) {
	contactRepository := &adapters.ContactHandler{}
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(logger, contactService)

	r.Get("/contacts", contactHTTP.List)
	r.Post("/contacts", contactHTTP.Save)
}
