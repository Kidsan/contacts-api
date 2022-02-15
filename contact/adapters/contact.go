package adapters

import (
	"context"

	contactsapi "github.com/kidsan/contacts-api"
)

type ContactHandler struct {
	data []contactsapi.Contact
}

func (c *ContactHandler) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	return c.data, nil
}

func (c *ContactHandler) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	c.data = append(c.data, s)
	return s, nil
}
