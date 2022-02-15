package adapters

import (
	"context"
	"fmt"

	uuid "github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
	"gorm.io/gorm"
)

type ContactRepository struct {
	connection *gorm.DB
}

func NewContactRepository(connection *gorm.DB) *ContactRepository {
	return &ContactRepository{connection: connection}
}

func (c *ContactRepository) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	result := make([]contactsapi.Contact, 0)
	sqlQuery := "select * from contacts_api.contacts;"

	tx := c.connection.WithContext(ctx).Raw(sqlQuery).Scan(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("adapters: could not list all contacts: %w", tx.Error)
	}

	return result, nil
}

func (c *ContactRepository) Save(ctx context.Context, newContact contactsapi.Contact) (contactsapi.Contact, error) {
	newContact.ID = uuid.New()
	tx := c.connection.WithContext(ctx).Create(&newContact)
	if tx.Error != nil {
		return contactsapi.Contact{}, fmt.Errorf("adapters: could not create new contact: %w", tx.Error)
	}

	return newContact, nil
}
