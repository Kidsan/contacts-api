package main

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

type ContactsHandler struct {
	data []string
}

func (c *ContactsHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Join(c.data, ",")))
}

func (c *ContactsHandler) Post(w http.ResponseWriter, r *http.Request) {
	c.data = append(c.data, c.data[0]+"a")
	w.Write([]byte(c.data[len(c.data)-1]))
}

func main() {
	r := chi.NewRouter()
	handler := ContactsHandler{}
	handler.data = append(handler.data, "hello")
	r.Get("/", handler.Get)
	r.Post("/", handler.Post)
	http.ListenAndServe(":3000", r)
}
