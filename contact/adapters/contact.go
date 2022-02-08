package adapters

import (
	uuid "github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
)

type ContactHandler struct {
	data []contactsapi.Contact
}

func (c *ContactHandler) Get() []contactsapi.Contact {
	return c.data
}

func (c *ContactHandler) Save(s contactsapi.Contact) contactsapi.Contact {
	newContact := contactsapi.Contact{
		ID:   uuid.New(),
		Name: s.Name,
	}
	c.data = append(c.data, newContact)
	return newContact
}
