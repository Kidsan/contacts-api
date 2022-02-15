package mock

import (
	"context"
	"fmt"

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

func (m *MockContactHandler) Find(ctx context.Context, id string) (contactsapi.Contact, error) {
	var result contactsapi.Contact
	for _, v := range m.Data {
		if v.ID.String() == id {
			return result, nil
		}
	}
	return result, fmt.Errorf("adapters(mock)")
}
