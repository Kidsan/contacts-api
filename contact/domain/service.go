package domain

import "github.com/kidsan/contacts-api/contact/adapters"

type ContactRepository interface {
	Get() []adapters.Contact
	Save(string) adapters.Contact
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

func (cs *ContactService) Save(s string) adapters.Contact {
	return cs.repository.Save(s)
}
