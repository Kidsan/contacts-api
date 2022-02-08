package domain

import (
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

func (cs *ContactService) Get() []contactsapi.Contact {
	return cs.repository.Get()
}

func (cs *ContactService) Save(s contactsapi.Contact) contactsapi.Contact {
	result := cs.repository.Save(contactsapi.Contact{
		Name: s.Name,
	})
	return contactsapi.Contact{
		ID:   result.ID,
		Name: result.Name,
	}
}
