package domain

import (
	"context"

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
	return cs.repository.Get(ctx)
}

func (cs *ContactService) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	result, _ := cs.repository.Save(ctx, contactsapi.Contact{
		Name: s.Name,
	})
	return contactsapi.Contact{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}
