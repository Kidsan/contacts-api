package ports

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type ContactService interface {
	Get() []string
	Save(string) string
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

func (c *ContactHTTP) GetRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", c.List)
	r.Post("/", c.Save)
	return r
}

func (c *ContactHTTP) List(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Listing all contects")
	w.Write([]byte(strings.Join(c.service.Get(), ","))) // send json
}

func (c *ContactHTTP) Save(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("Creating new contact")
	c.service.Save("POST") // marshall body and save, generate id
	w.Write([]byte(strings.Join(c.service.Get(), ",")))
}
