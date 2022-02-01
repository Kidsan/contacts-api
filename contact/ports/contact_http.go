package ports

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type ContactService interface {
	Get() []string
	Save(string) string
}

type ContactHTTP struct {
	service ContactService
}

func NewContactRouter(s ContactService) *ContactHTTP {
	return &ContactHTTP{
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
	w.Write([]byte(strings.Join(c.service.Get(), ",")))
}

func (c *ContactHTTP) Save(w http.ResponseWriter, r *http.Request) {
	c.service.Save("POST")
	w.Write([]byte(strings.Join(c.service.Get(), ",")))
}
