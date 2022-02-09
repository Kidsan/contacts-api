package adapters

import (
	"context"

	uuid "github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
)

type ContactHandler struct {
	data []contactsapi.Contact
}

func (c *ContactHandler) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	return c.data, nil
}

func (c *ContactHandler) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	newContact := contactsapi.Contact{
		ID:   uuid.New(),
		Name: s.Name,
	}
	c.data = append(c.data, newContact)
	return newContact, nil
}
