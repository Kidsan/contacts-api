package ports

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (c *ContactHTTP) List(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Listing all contacts")
	contacts, err := c.service.Get(r.Context())
	if err != nil {
		return
	}
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
	res, err := c.service.Save(r.Context(), newContact)
	if err != nil {
		return
	}
	result, err := json.Marshal(res)
	if err != nil {
		return
	}
	w.Write([]byte(result))
}
