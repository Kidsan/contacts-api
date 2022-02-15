package mock

import (
	"context"

	contactsapi "github.com/kidsan/contacts-api"
)

type MockContactHandler struct {
	Data []contactsapi.Contact
}

func (m *MockContactHandler) Get(ctx context.Context) ([]contactsapi.Contact, error) {
	return m.Data, nil
}

func (m *MockContactHandler) Save(ctx context.Context, s contactsapi.Contact) (contactsapi.Contact, error) {
	m.Data = append(m.Data, s)
	return s, nil
}
