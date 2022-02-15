package domain

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (cs *ContactService) Save(ctx context.Context, newContact contactsapi.Contact) (contactsapi.Contact, error) {
	newContact.ID = uuid.New()
	result, err := cs.repository.Save(ctx, newContact)
	if err != nil {
		return contactsapi.Contact{}, fmt.Errorf("domain(contact-service): could save new contact %w", err)
	}
	return result, nil
}

func (cs *ContactService) Find(ctx context.Context, id string) (contactsapi.Contact, error) {
	result, err := cs.repository.Find(ctx, id)
	if err != nil {
		return contactsapi.Contact{}, fmt.Errorf("domain(contact-service): could find contact %w", err)
	}
	return result, nil
}
