package ports

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kidsan/contacts-api/contact/adapters"
	"go.uber.org/zap"
)

type ContactService interface {
	Get() []adapters.Contact
	Save(Contact) adapters.Contact
}

type ContactHTTP struct {
	logger  *zap.Logger
	service ContactService
}

type Contact struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func NewContactRouter(logger *zap.Logger, s ContactService) *ContactHTTP {
	return &ContactHTTP{
		logger:  logger,
		service: s,
	}
}

func (c *ContactHTTP) GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)
	r.Post("/", c.Save)
	return r
}

func (c *ContactHTTP) List(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Listing all contacts")
	contacts := c.service.Get()
	var allContacts []Contact
	for _, v := range contacts {
		allContacts = append(allContacts, Contact{
			Name: v.Name,
			ID:   v.ID,
		})
	}
	result, err := json.Marshal(allContacts)
	if err != nil {
		return
	}
	w.Write(result)
}

func (c *ContactHTTP) Save(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Creating new contact")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var newContact Contact
	err = json.Unmarshal(body, &newContact)
	if err != nil {
		return
	}
	c.service.Save(newContact) // marshall body and save, generate id
	//w.Write([]byte(strings.Join(c.service.Get(), ",")))
}
