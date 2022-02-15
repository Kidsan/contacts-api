package contactsapi

import (
	"context"

	"github.com/google/uuid"
)

type Contact struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

type ContactRepository interface {
	Get(context.Context) ([]Contact, error)
	Save(context.Context, Contact) (Contact, error)
	Find(context.Context, string) (Contact, error)
}

type ContactService interface {
	Get(context.Context) ([]Contact, error)
	Save(context.Context, Contact) (Contact, error)
	Find(context.Context, string) (Contact, error)
}
