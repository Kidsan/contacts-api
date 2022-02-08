package ports

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	contactsapi "github.com/kidsan/contacts-api"
	"go.uber.org/zap"
)

type ContactHTTP struct {
	logger  *zap.Logger
	service contactsapi.ContactService
}

func NewContactRouter(logger *zap.Logger, s contactsapi.ContactService) *ContactHTTP {
	return &ContactHTTP{
		logger:  logger,
		service: s,
	}
}

func (c *ContactHTTP) GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)
	r.Post("/", c.Save)
	return r
}

func (c *ContactHTTP) List(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Listing all contacts")
	contacts := c.service.Get()
	var allContacts []contactsapi.Contact
	for _, v := range contacts {
		allContacts = append(allContacts, contactsapi.Contact{
			Name: v.Name,
			ID:   v.ID,
		})
	}
	result, err := json.Marshal(allContacts)
	if err != nil {
		return
	}
	w.Write(result)
}

func (c *ContactHTTP) Save(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Creating new contact")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var newContact contactsapi.Contact
	err = json.Unmarshal(body, &newContact)
	if err != nil {
		return
	}
	res := c.service.Save(newContact)
	result, _ := json.Marshal(res)
	w.Write([]byte(result))
}
