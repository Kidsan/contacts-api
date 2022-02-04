package adapters

import (
	uuid "github.com/google/uuid"
)

type Contact struct {
	Name string
	ID   uuid.UUID
}

type ContactHandler struct {
	data []Contact
}

func (c *ContactHandler) Get() []Contact {
	return c.data
}

func (c *ContactHandler) Save(s Contact) Contact {
	newContact := Contact{
		ID:   uuid.New(),
		Name: s.Name,
	}
	c.data = append(c.data, newContact)
	return newContact
}
