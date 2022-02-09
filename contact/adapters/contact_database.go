package adapters

import (
	"context"
	"fmt"

	uuid "github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
	"gorm.io/gorm"
)

type ContactDatabase struct {
	connection *gorm.DB
}

func (c *ContactDatabase) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	contacts := make([]contactsapi.Contact, 0)
	sqlQuery := "select * from contacts"

	tx := c.connection.WithContext(ctx).Raw(sqlQuery).Scan(&contacts)
	if tx.Error != nil {
		return nil, fmt.Errorf("adapters: could not list all contacts: %w", tx.Error)
	}

	return contacts, nil
}

func (c *ContactDatabase) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	newContact := contactsapi.Contact{
		ID:   uuid.New(),
		Name: s.Name,
	}
	// c.data = append(c.data, newContact)
	return newContact, nil
}
