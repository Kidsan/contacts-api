package contactsapi

import (
	"github.com/google/uuid"
)

type Contact struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}
