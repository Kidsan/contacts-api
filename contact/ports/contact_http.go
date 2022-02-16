package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
	"go.uber.org/zap"
)

type ContactService interface {
	Get(context.Context) ([]contactsapi.Contact, error)
	Save(context.Context, contactsapi.Contact) (contactsapi.Contact, error)
	Find(context.Context, string) (contactsapi.Contact, error)
}

type ContactHTTP struct {
	logger  *zap.Logger
	service ContactService
}

func NewContactRouter(logger *zap.Logger, s ContactService) *ContactHTTP {
	return &ContactHTTP{
		logger:  logger,
		service: s,
	}
}

func (c *ContactHTTP) Init() func(router chi.Router) {
	return func(router chi.Router) {
		router.Get("/", c.GetAllHandler())
		router.Post("/", c.PostHandler())
		router.With(c.checkIDParam()).Get("/{id}", c.FindHandler())
	}
}

func (c *ContactHTTP) GetAllHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.logger.Info("Listing all contacts")
		contacts, err := c.service.Get(r.Context())
		if err != nil {
			c.logger.Error(fmt.Sprintf("ports(contact): could not list all contacts: %v", err))
			return
		}
		var allContacts []contactsapi.Contact
		for _, v := range contacts {
			allContacts = append(allContacts, contactsapi.Contact{
				Name: v.Name,
				ID:   v.ID,
			})
		}
		result, err := json.Marshal(allContacts)
		if err != nil {
			c.logger.Error(fmt.Sprintf("ports(contact): could not list all contacts: %v", err))
			return
		}
		w.Write(result)
	}
}

func (c *ContactHTTP) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.logger.Info("Creating new contact")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			c.logger.Error(fmt.Sprintf("ports(contact): could not create contact: %v", err))
			return
		}
		var newContact contactsapi.Contact
		err = json.Unmarshal(body, &newContact)
		if err != nil {
			c.logger.Error(fmt.Sprintf("ports(contact): could not create contact: %v", err))
			return
		}
		res, err := c.service.Save(r.Context(), newContact)
		if err != nil {
			c.logger.Error(fmt.Sprintf("ports(contact): could not create contact: %v", err))
			return
		}
		result, err := json.Marshal(res)
		if err != nil {
			c.logger.Error(fmt.Sprintf("ports(contact): could not create contact: %v", err))
			return
		}
		w.Write([]byte(result))
	}
}

func (c *ContactHTTP) FindHandler() http.HandlerFunc {
	const msg = "cannot get route"

	return func(w http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")
		contact, err := c.service.Find(req.Context(), id)
		if err != nil {
			c.logger.Error(err.Error())
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(&contact); err != nil {
			c.logger.Error(fmt.Sprintf("ports: %v", err.Error()))
			http.Error(w, msg, http.StatusInternalServerError)
			return
		}
	}
}

func (c *ContactHTTP) checkIDParam() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id := chi.URLParam(req, "id")
			_, err := uuid.Parse(id)
			if err != nil {
				msg := "cannot parse route id"
				http.Error(w, msg, http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
