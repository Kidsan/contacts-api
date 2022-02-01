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

// func (c *ContactsHandler) Get(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(strings.Join(c.data, ",")))
// }

// func (c *ContactsHandler) Post(w http.ResponseWriter, r *http.Request) {
// 	c.data = append(c.data, c.data[0]+"a")
// 	w.Write([]byte(c.data[len(c.data)-1]))
// }
