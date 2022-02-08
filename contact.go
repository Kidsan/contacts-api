package contactsapi

import "github.com/google/uuid"

type Contact struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

type ContactRepository interface {
	Get() []Contact
	Save(Contact) Contact
}

type ContactService interface {
	Get() []Contact
	Save(Contact) Contact
}
