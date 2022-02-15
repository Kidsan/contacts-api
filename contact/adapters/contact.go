package adapters

import (
	"context"

	uuid "github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
	"gorm.io/gorm"
)

type ContactHandler struct {
	data       []contactsapi.Contact
	connection *gorm.DB
}

func NewContactRepository(connection *gorm.DB) *ContactHandler {
	return &ContactHandler{connection: connection}
}

func (c *ContactHandler) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	return c.data, nil
}

func (c *ContactHandler) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	s.ID = uuid.New()
	c.data = append(c.data, s)
	return s, nil
}
