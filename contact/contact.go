package contact

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"
)

func BuildRouter() chi.Router {
	contactRepository := &adapters.ContactHandler{}
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(contactService)
	router := contactHTTP.GetRouter()

	return router
}
