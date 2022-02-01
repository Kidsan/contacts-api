package contact

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"
)

func Init() chi.Router {
	a := &adapters.ContactHandler{}
	b := domain.NewContactService(a)
	c := ports.NewContactRouter(b)
	router := c.GetRouter()

	return router
}
