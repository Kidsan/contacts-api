package adapters

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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
	sqlQuery := "select * from contacts;"

	tx := c.connection.WithContext(ctx).Raw(sqlQuery).Scan(&result)
	if tx.Error != nil {
		return nil, fmt.Errorf("adapters: could not list all contacts: %w", tx.Error)
	}

	return result, nil
}

func (c *ContactRepository) Save(ctx context.Context, newContact contactsapi.Contact) (contactsapi.Contact, error) {
	tx := c.connection.WithContext(ctx).Create(&newContact)
	if tx.Error != nil {
		return contactsapi.Contact{}, fmt.Errorf("adapters: could not create new contact: %w", tx.Error)
	}

	return newContact, nil
}

func (c *ContactRepository) Find(ctx context.Context, id string) (contactsapi.Contact, error) {
	var result contactsapi.Contact
	sqlQuery := "select * from contacts where id = ?;"

	tx := c.connection.WithContext(ctx).Raw(sqlQuery, id).Scan(&result)
	if tx.Error != nil {
		return contactsapi.Contact{}, fmt.Errorf("adapters: could not find specific contact with id %v: %w", id, tx.Error)
	}

	if result.ID == uuid.Nil {
		return contactsapi.Contact{}, nil
	}

	return result, nil
}
