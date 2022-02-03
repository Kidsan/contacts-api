package adapters

import (
	uuid "github.com/google/uuid"
)

type Contact struct {
	name string
	id   uuid.UUID
}

type ContactHandler struct {
	data []Contact
}

func (c *ContactHandler) Get() []Contact {
	return c.data
}

func (c *ContactHandler) Save(s string) Contact {
	newContact := Contact{
		id:   uuid.New(),
		name: s,
	}
	c.data = append(c.data, newContact)
	return newContact
}
