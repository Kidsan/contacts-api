package domain

import (
	"context"
	"fmt"

	contactsapi "github.com/kidsan/contacts-api"
)

type ContactService struct {
	repository contactsapi.ContactRepository
}

func NewContactService(cr contactsapi.ContactRepository) *ContactService {
	return &ContactService{
		repository: cr,
	}
}

func (cs *ContactService) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	contacts, err := cs.repository.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("domain(contact-service): could not get all contacts %w", err)
	}
	return contacts, nil
}

func (cs *ContactService) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	result, err := cs.repository.Save(ctx, s)
	if err != nil {
		return contactsapi.Contact{}, fmt.Errorf("domain(contact-service): could save new contact %w", err)
	}
	return result, nil
}
