package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kidsan/contacts-api/contact"
)

func main() {
	r := chi.NewRouter()
	contactRoutes := contact.Init()
	r.Mount("/contacts", contactRoutes)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	http.ListenAndServe(":3000", r)
}
