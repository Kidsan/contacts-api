package contact

import (
	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/domain"
	"github.com/kidsan/contacts-api/contact/ports"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func RegisterContactRoutes(logger *zap.Logger, r *chi.Mux, connection *gorm.DB) {
	contactRepository := adapters.NewContactRepository(connection)
	contactService := domain.NewContactService(contactRepository)
	contactHTTP := ports.NewContactRouter(logger, contactService)

	r.Get("/contacts", contactHTTP.GetAllHandler())
	r.Post("/contacts", contactHTTP.PostHandler())
	// maybe return a http handler that server package can mount?
}
