package domain

import (
	"github.com/kidsan/contacts-api/contact/adapters"
	"github.com/kidsan/contacts-api/contact/ports"
)

type ContactRepository interface {
	Get() []adapters.Contact
	Save(adapters.Contact) adapters.Contact
}

type ContactService struct {
	repository ContactRepository
}

func NewContactService(cr ContactRepository) *ContactService {
	return &ContactService{
		repository: cr,
	}
}

func (cs *ContactService) Get() []adapters.Contact {
	return cs.repository.Get()
}

func (cs *ContactService) Save(s ports.Contact) adapters.Contact {
	return cs.repository.Save(adapters.Contact{
		Name: s.Name,
	})
}
