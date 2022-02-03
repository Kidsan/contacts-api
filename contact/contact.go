package contact

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"
	"go.uber.org/zap"
)

func BuildRouter(logger *zap.Logger) chi.Router {
	contactRepository := &adapters.ContactHandler{}
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(logger, contactService)
	router := contactHTTP.GetRouter()

	return router
}
